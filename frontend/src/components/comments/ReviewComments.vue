<script setup lang="ts">
import Comment from '@/components/comments/Comment.vue';
import { useAuthStore } from '@/services/authStore';
import { ADD_COMMENT, GET_COMMENTS_BY_ITEM_ID_REVIEW_ID } from '@/services/queries';
import { emptyCommentsPage, type CommentsPageType, type CommentType } from '@/types/comments';
import { handleGqlError } from '@/utils/error';
import { Form } from '@primevue/forms';
import { useMutation, useQuery } from '@vue/apollo-composable';
import Button from 'primevue/button';
import Card from 'primevue/card';
import FloatLabel from "primevue/floatlabel";
import Paginator, { type PageState } from 'primevue/paginator';
import Textarea from 'primevue/textarea';
import { ref, watch } from 'vue';
import { useRoute, useRouter } from 'vue-router';

const store = useAuthStore();
const route = useRoute();
const router = useRouter();
const newComment = ref({
  text: '',
});

const props = defineProps<{
  itemId: string;
}>();


const pageSize = 10;
let commentsPage = ref<CommentsPageType>(emptyCommentsPage);
let commentsPageLoading = ref(true);



const fetch_comments = async () => {
  const { loading, onError, onResult } = useQuery(
    GET_COMMENTS_BY_ITEM_ID_REVIEW_ID,
    {
      itemId: props.itemId,
      reviewId: route.params.userId,
      pageSize: pageSize,
      page: 1,
    }
  );

  commentsPageLoading = loading;

  onError((err) => {
    handleGqlError(router, err);
  });

  onResult((res: any) => {
    if (res.loading) {
      return;
    }

    commentsPage.value = res?.data?.commentsPage;
  })
};

const change_page = async (event: PageState) => {
  const { onError, onResult } = useQuery(
    GET_COMMENTS_BY_ITEM_ID_REVIEW_ID,
    {
      itemId: props.itemId,
      reviewId: route.params.userId,
      pageSize: pageSize,
      page: event.page + 1,
    }
  );

  onError((err) => {
    handleGqlError(router, err);
  });

  onResult((res: any) => {
    if (res.loading) {
      return;
    }

    commentsPage.value = res?.data?.commentsPage;
  })
};





const submitComment = () => {
  if (newComment.value.text == '') {
    // TODO add toast component
    return;
  }

  const { mutate: addComment, error: addCommentError, onDone: addCommentOnDone } = useMutation(
    ADD_COMMENT,
    () => ({
      variables: {
        itemId: props.itemId,
        reviewId: route.params.userId,
        text: newComment.value.text,
      },
    }
  ));

  watch(addCommentError, (err) => {
    handleGqlError(router, err);
  });

  addCommentOnDone((result) => {
    // TODO find and replace the comment in the list
    newComment.value.text = '';
  });

  addComment();
}


const fetch_data = async () => {
  fetch_comments();
};

// TODO better waty to do this?
watch(() => route.params, fetch_data, { immediate: true })
</script>

<template>
  <!-- TODO enable editing, delete, reply -->
  <Toast />
  <div>
    <Card>
      <template #title>
        <div class="pb-2">
          Comments
        </div>
      </template>
      <template #content>
        <div v-for="c in commentsPage.comments" class="mb-2">
          <Comment :comment="c" />
        </div>

        <div>
          <Paginator
            :rows="pageSize"
            :totalRecords="commentsPage.totalComments"
            :first="0"
            @page="change_page($event)"
            class="mt-4"
          >
          </Paginator>
        </div>

      </template>
      <template #footer>
        <div>
          <Card class="bg-comment">
            <template #content>
              <Form class="flex">
                <FloatLabel variant="on" class="w-full">
                  <Textarea v-model="newComment.text" id="over_label" rows="2" class="w-full" />
                  <label for="on_label">Your comment</label>
                </FloatLabel>
                <div class="">
                  <Button type="submit" severity="secondary" label="Add comment" class="ml-2" @click="submitComment"/>
                </div>
              </Form>
            </template>
          </Card>
        </div>
      </template>
    </Card>
  </div>




</template>

<style scoped>

.bg-comment {
  background: var(--color-primary-light);
}

</style>
