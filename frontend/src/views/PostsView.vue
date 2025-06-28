<script setup lang="ts">
import LinkedReview from '@/components/posts/LinkedReview.vue';
import Post from '@/components/posts/Post.vue';
import { useAuthStore } from "@/services/authStore";
import { ADD_POST, GET_RECENT_POSTS, GET_REWIEW_BY_ID_POST_LINK } from "@/services/queries";
import { emptyRecentPosts, type RecentPostsType } from '@/types/posts';
import { emptyReview, type ReviewType } from '@/types/review';
import { handleGqlError } from '@/utils/error';
import { Form } from '@primevue/forms';
import { useMutation, useQuery } from '@vue/apollo-composable';
import Avatar from 'primevue/avatar';
import Button from 'primevue/button';
import Card from 'primevue/card';
import FloatLabel from "primevue/floatlabel";
import ProgressSpinner from 'primevue/progressspinner';
import SelectButton from 'primevue/selectbutton';
import Textarea from 'primevue/textarea';
import { useToast } from 'primevue/usetoast';
import { ref, watch } from 'vue';
import { useRoute, useRouter } from 'vue-router';


const store = useAuthStore();
const route = useRoute();
const router = useRouter();
const toast = useToast();


const pageSize = 10;
const filter = ref({ name: 'All', value: '' });
const filterOptions = ref([
  { name: 'All',       value: '' },
  { name: 'Following', value: 'following' },
  { name: 'User',      value: 'user' },
]);

let posts = ref<RecentPostsType>(emptyRecentPosts);
let postsLoading = ref(true);

const newPostContent = ref<string>('');
let linkedReview = ref<ReviewType>(emptyReview);
let linkedReviewLoading = ref(true);
let showLinkedReview = ref(false);

const fetch_posts = async () => {
  const { loading, onError, onResult } = useQuery(
    GET_RECENT_POSTS,
    {
      pageSize: pageSize,
      page: 1,
      type: filter.value.value,
    }
  );

  postsLoading = loading;

  onError((err) => {
    handleGqlError(router, err);
  });

  onResult((res: any) => {
    if (res.loading) {
      return;
    }

    posts.value = res?.data?.getRecentPost;
  })
};


const submitPost = async () => {
  if (newPostContent.value == '') {
    // TODO add toast component
    return;
  }

  let vars: { content: string; linkedReviewId?: string } = {
    content: newPostContent.value,
  }
  if (linkedReview.value._id !== '' && linkedReview.value._id) {
    vars.linkedReviewId = linkedReview.value._id
  }

  const { mutate: addPost, error: addPostError, onDone: addPostOnDone } = useMutation(
    ADD_POST,
    () => ({
      variables: vars,
    }
  ));

  watch(addPostError, (err) => {
    handleGqlError(router, err);
  });

  addPostOnDone(async (result) => {
    newPostContent.value = '';
    removeLinkedReview();

    if (route.params.reviewId) {
      router.push({name: 'posts'});
    }

    toast.add({ severity: 'success', summary: `Added new post`, life: 2000 });

    // TODO have a secondary posts component before the public ones?
  });

  addPost();
}

const fetch_rewiew = async (id: string) => {
  const { onError, onResult } = useQuery(
    GET_REWIEW_BY_ID_POST_LINK,
    {
      reviewId: id,
    }
  );

  linkedReviewLoading.value = true;

  onError((err) => {
    handleGqlError(router, err, ["mongo: no documents in result"]);
    toast.add({ severity: 'warn', summary: `This is not a valid review link`, life: 2000 });
  });
  
  onResult((res) => {
    if (res.loading) {
      return;
    }

    linkedReview.value = res?.data?.reviewById;
    linkedReviewLoading.value = false;
  });
};

watch(newPostContent, async (content) => {
  const match = content.match(new RegExp(`${window.location.origin}/review/[a-f\\d]{24}`, 'i'))
  if (!match) {
    return;
  }
  
  const reviewLink = match[0];
  const reviewId = reviewLink.split('/').pop();

  if (!reviewId) {
    return;
  }

  await fetch_rewiew(reviewId);

  newPostContent.value = content.replace(reviewLink, ``);
  toast.add({ severity: 'info', summary: `Review link added`, life: 1000 });
  showLinkedReview.value = true;
});

const removeLinkedReview = () => {
  linkedReview.value = emptyReview;
  showLinkedReview.value = false;
};


const fetch_data = async () => {
  fetch_posts();

  if (route.params.reviewId) {
    await fetch_rewiew(route.params.reviewId as string);
    showLinkedReview.value = true;
  }
};

// TODO better waty to do this?
watch(() => route.params, fetch_data, { immediate: true })
</script>

<template>

  <div class="pb-2 w-full items-center flex justify-center">
    <SelectButton
    v-model="filter"
    :options="filterOptions"
    optionLabel="name"
    @change="fetch_posts"
    />
  </div>
  
  <div v-if="postsLoading" class="flex justify-center pt-12">
    <ProgressSpinner />
  </div>
  
  <div v-else>
    <Card>
      <template #header>
        <div class="px-6 pt-4">
          <h2 class="text-2xl">Write your post!</h2>
        </div>
      </template>
      <template #content>
        <Form class="">
          <FloatLabel variant="on" class="w-full">
            <Textarea v-model="newPostContent" id="over_label" rows="4" class="w-full" />
            <label for="on_label">Your post</label>
          </FloatLabel>
        </Form>
        <div v-if="showLinkedReview && linkedReview._id !== ''">
          <LinkedReview
            :review="linkedReview"
            :showCloseButton="true"
            @closeLinkedReview="removeLinkedReview"
          />
        </div>
        <div class="flex justify-end mt-4">
          <Button type="submit" severity="secondary" label="Add Post" class="ml-2" @click="submitPost"/>
        </div>
      </template>
    </Card>
  
    <div v-for="post in posts.posts" :key="post._id" class="mt-4">
      <Post
        :post="post"
      />
    </div>
  </div>


</template>

<style scoped>

</style>
