<script setup lang="ts">
import Comment from '@/components/comments/Comment.vue';
import LikesDislikes from '@/components/likes-dislikes/LikesDislikes.vue';
import LinkedReview from '@/components/posts/LinkedReview.vue';
import { ADD_COMMENT, GET_COMMENTS_BY_ITEM_ID } from '@/services/queries';
import { emptyComment, emptyCommentsPage, type CommentsPageType, type CommentType } from '@/types/comments';
import type { PostType } from '@/types/posts';
import { handleGqlError } from '@/utils/error';
import { navigateToUser } from '@/utils/navigate';
import { Form } from '@primevue/forms';
import { useMutation, useQuery } from '@vue/apollo-composable';
import type mitt from 'mitt';
import Avatar from 'primevue/avatar';
import Button from 'primevue/button';
import Card from 'primevue/card';
import FloatLabel from "primevue/floatlabel";
import Textarea from 'primevue/textarea';
import { inject, nextTick, ref, watch } from 'vue';
import { useRoute, useRouter } from 'vue-router';


const route = useRoute();
const router = useRouter();
const emitter = inject<ReturnType<typeof mitt>>('emitter')

const props = defineProps<{
  post: PostType;
}>();

const newComment = ref({
  text: '',
});

const showAddComment = ref(false);
const showComments = ref(false);
let replyingToCommentId = ref('');
let replyingToComment = ref<CommentType>(emptyComment);
const pageSize = 15;
let commentsPage = ref<CommentsPageType>(emptyCommentsPage);

// Basically duplicated code from ReviewComments.vue
const submitComment = async () => {
  if (newComment.value.text == '') {
    // TODO add toast component
    return;
  }

  const { mutate: addComment, error: addCommentError, onDone: addCommentOnDone } = useMutation(
    ADD_COMMENT,
    () => ({
      variables: {
        itemId: props.post._id,
        itemType: 'posts',
        text: newComment.value.text,
        replyingToId: replyingToCommentId.value,
      },
    }
  ));

  watch(addCommentError, (err) => {
    handleGqlError(router, err);
  });

  addCommentOnDone(async (result) => {
    replyingToCommentId.value = '';
    replyingToComment.value = emptyComment;
    showAddComment.value = false
  });

  addComment();
}

const fetch_comments = async () => {
  const { loading, onError, onResult } = useQuery(
    GET_COMMENTS_BY_ITEM_ID,
    {
      itemId: props.post._id,
      pageSize: pageSize,
      page: 1,
    }
  );

  // commentsPageLoading = loading;

  onError((err) => {
    handleGqlError(router, err);
  });

  onResult((res: any) => {
    if (res.loading) {
      return;
    }

    commentsPage.value = res?.data?.commentsPage;
    showComments.value = true;
  })
};

const toggle_comments = () => {
  showComments.value = !showComments.value;
  if (showComments.value && commentsPage.value.comments.length === 0) {
    fetch_comments();
  }
};

emitter?.on('replyToComment', async (e: any) => {
  const comment = e.comment;

  if (comment && props.post._id === comment.itemId) {
    replyingToCommentId.value = comment._id;
    showAddComment.value = true;
    replyingToComment.value = comment;
    newComment.value.text = `@${comment?.user.displayName} `;

    await nextTick();
    const element = document.getElementById('repply-form');
    if (element) {
      element.scrollIntoView({ behavior: 'smooth' });
    }
  } else {
    replyingToCommentId.value = '';
  }
})

</script>

<template>

  <Card class="mb-3">
    <template #header>
      <div class="flex align-items-center mx-4 mt-4">
        <div class="my-auto cursor-pointer" @click="navigateToUser(router, props.post.user._id)">
          <Avatar :image="props.post.user.images[0].url" class="mr-2" size="normal" shape="circle" />
        </div>
        <div class="block">
          <span class="cursor-pointer" @click="navigateToUser(router, props.post.user._id)">
            {{ props.post.user.displayName }}
          </span>
          <br>
          <span class="text-xs text-neutral-500">
            {{ new Intl.DateTimeFormat('en-US', {
              year: 'numeric',
              month: 'short',
              day: 'numeric',
              hour: 'numeric',
              minute: 'numeric'
            }).format(new Date(props.post.updatedAt)) }}
          </span>
        </div>
      </div>
    </template>
    <template #content>
      <div class="flex-1 break-words">
        {{ props.post.content }}
      </div>
      <!-- TODO -->
      <div v-if="props.post.linkedReview">
        <LinkedReview
          :review="props.post.linkedReview"
          :showCloseButton="false"
        />
      </div>
    </template>
    <template #footer>
      <div class="flex justify-between mt-4">
        <div>
          <LikesDislikes
            :id="props.post._id"
            idType="post"
            :userReaction="props.post.userReaction"
            :likesCount="props.post.likesCount"
            :dislikesCount="props.post.dislikesCount"
            :fontRemSize="1"
          />
        </div>
        <div>
          <div class="flex align-items-center cursor-pointer">
            <i class="pi pi-comment"></i>
            <span class="ml-2" @click="toggle_comments">
              {{ props.post.commentsNumber }} 
              {{ props.post.commentsNumber === 1 ? 'Comment' : 'Comments' }}
            </span>
          </div>
        </div>
        <div>
          <div class="flex align-items-center cursor-pointer" @click="showAddComment = true">
            <i class="pi pi-pen-to-square"></i>
            <span class="ml-2">Add comment</span>
          </div>
        </div>
      </div>
      
      <div v-if="showComments" v-for="comment in commentsPage.comments" class="mt-4">
        <Comment
          :comment="comment"
          :showReportButton="true"
          :showLikes="true"
        />
      </div>

      <div class="mt-4" v-if="showAddComment" id="repply-form">
        <Form class="sm:flex">
          <FloatLabel variant="on" class="w-full">
            <Textarea v-model="newComment.text" id="over_label" rows="4" class="w-full" />
            <label for="on_label">Your comment</label>
          </FloatLabel>
          <div class="sm:ml-2">
            <Button type="submit" severity="secondary" label="Add comment" class="my-2 sm:my-0 w-full whitespace-nowrap" @click="submitComment"/>
            <Button severity="secondary" label="Cancel" class="w-full sm:mt-3 whitespace-nowrap" @click="showAddComment = false"/>
          </div>
        </Form>
      </div>
    </template>
  </Card>

</template>

<style>

</style>
