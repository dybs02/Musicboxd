<script setup lang="ts">
import DiaryReviewsList from '@/components/diary/DiaryReviewsList.vue';
import { useAuthStore } from "@/services/authStore";
import { ADD_POST, GET_RECENT_POSTS, GET_RECENT_USER_REVIEWS_PAGINATION } from "@/services/queries";
import { emptyRecentUserReviews, type RecentUserReviewsType } from '@/types/review';
import { handleGqlError } from '@/utils/error';
import { useMutation, useQuery } from '@vue/apollo-composable';
import Card from 'primevue/card';
import Paginator, { type PageState } from 'primevue/paginator';
import ProgressSpinner from 'primevue/progressspinner';
import { ref, watch } from 'vue';
import SelectButton from 'primevue/selectbutton';
import { Form } from '@primevue/forms';
import Button from 'primevue/button';
import FloatLabel from "primevue/floatlabel";
import Textarea from 'primevue/textarea';
import { useRoute, useRouter } from 'vue-router';
import Divider from 'primevue/divider';
import { emptyRecentPosts, type RecentPostsType } from '@/types/posts';
import Post from '@/components/posts/Post.vue';


const store = useAuthStore();
const route = useRoute();
const router = useRouter();


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

  const { mutate: addPost, error: addPostError, onDone: addPostOnDone } = useMutation(
    ADD_POST,
    () => ({
      variables: {
        content: newPostContent.value,
      },
    }
  ));

  watch(addPostError, (err) => {
    handleGqlError(router, err);
  });

  addPostOnDone(async (result) => {
    newPostContent.value = '';
    // TODO have a secondary posts component before the public ones?
  });

  addPost();
}


const fetch_data = async () => {
  fetch_posts();
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
          <div class="flex justify-end mt-4">
            <Button type="submit" severity="secondary" label="Add Post" class="ml-2" @click="submitPost"/>
          </div>
        </Form>
      </template>
    </Card>
  
    <div v-for="post in posts.posts" :key="post._id" class="mt-4">

      <Post
        :post="post"
      />
      <!-- {{ post.content }} -->
      <!-- {{ post.content }} - {{ post.user.displayName }} -->
    </div>
  </div>


</template>

<style scoped>

</style>
