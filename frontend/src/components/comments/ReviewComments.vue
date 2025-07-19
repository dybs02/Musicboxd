<script setup lang="ts">
import Comment from '@/components/comments/Comment.vue';
import { useAuthStore } from '@/services/authStore';
import { ADD_COMMENT, GET_COMMENTS_BY_ITEM_ID } from '@/services/queries';
import { emptyComment, emptyCommentsPage, type CommentsPageType, type CommentType } from '@/types/comments';
import { handleGqlError } from '@/utils/error';
import { Form } from '@primevue/forms';
import { useMutation, useQuery } from '@vue/apollo-composable';
import Button from 'primevue/button';
import Card from 'primevue/card';
import FloatLabel from "primevue/floatlabel";
import Paginator, { type PageState } from 'primevue/paginator';
import Textarea from 'primevue/textarea';
import { inject, ref, watch } from 'vue';
import { useRoute, useRouter } from 'vue-router';
import Divider from 'primevue/divider';
import mitt from 'mitt';


const store = useAuthStore();
const route = useRoute();
const router = useRouter();
const emitter = inject<ReturnType<typeof mitt>>('emitter')

const newComment = ref({
  text: '',
});

const props = defineProps<{
  itemId: string;
}>();


const paginatorFirst = ref(0);
const pageSize = 3;
let commentsPage = ref<CommentsPageType>(emptyCommentsPage);
let commentsPageLoading = ref(true);
let replyingToCommentId = ref('');
let replyingToComment = ref<CommentType>(emptyComment);


const fetch_comments = async () => {
  const { loading, onError, onResult } = useQuery(
    GET_COMMENTS_BY_ITEM_ID,
    {
      itemId: route.params.id,
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
    GET_COMMENTS_BY_ITEM_ID,
    {
      itemId: route.params.id,
      pageSize: pageSize,
      page: event.page + 1,
    },
    {
      fetchPolicy: 'no-cache',
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





const submitComment = async () => {
  if (newComment.value.text == '') {
    // TODO add toast component
    return;
  }

  const { mutate: addComment, error: addCommentError, onDone: addCommentOnDone } = useMutation(
    ADD_COMMENT,
    () => ({
      variables: {
        itemId: route.params.id,
        itemType: 'reviews',
        text: newComment.value.text,
        replyingToId: replyingToCommentId.value,
      },
    }
  ));

  watch(addCommentError, (err) => {
    handleGqlError(router, err);
  });

  addCommentOnDone(async (result) => {
    const newTotalPages = Math.ceil((commentsPage.value.totalComments + 1) / pageSize)
    const fakePageState: PageState = {
      page: newTotalPages - 1,
      first: 0,
      rows: 0,
      pageCount: 0,
    }

    await change_page(fakePageState);

    paginatorFirst.value = pageSize * Math.floor(commentsPage.value.totalComments/pageSize);
    newComment.value.text = '';
    replyingToCommentId.value = '';
    replyingToComment.value = emptyComment;
  });

  addComment();
}


emitter?.on('replyToComment', (e: any) => {
  const comment = e.comment;
  console.log('replyToComment', comment);

  replyingToCommentId.value = comment._id;

  if (comment) {
    replyingToComment.value = comment;
    newComment.value.text = `@${comment?.user.displayName} `;
  } else {
    console.error('Comment not found');
    replyingToCommentId.value = '';
  }
})

const closeReply = () => {
  replyingToCommentId.value = '';
  replyingToComment.value = emptyComment;
};


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
          <Comment
            :comment="c"
          />
        </div>

        <div>
          <Paginator
            v-if="commentsPage.totalComments > 0"
            :rows="pageSize"
            :totalRecords="commentsPage.totalComments"
            :first="paginatorFirst"
            @page="change_page($event)"
            class="mt-4"
          >
          </Paginator>
        </div>

      </template>
      <template #footer>
        <div>
          <Card class="bg-comment">
            <template #header>
              <div v-if="replyingToCommentId !== ''" class="px-4 pt-4">
                <div class="flex">
                  <div class="text-slate-500 ml-2">
                    Replying to:
                  </div>
                  <Button
                    class="ml-auto"
                    @click="closeReply"
                    v-tooltip.bottom="`Close reply`" 
                    icon="pi pi-times"
                    aria-label="Save"
                    severity="secondary"
                    size="small"
                  />
                </div>
                <Comment
                  :comment="replyingToComment"
                  :showReportButton="false"
                  :showLikes="false"
                />
                <Divider />
              </div>
            </template>
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
