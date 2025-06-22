<script setup lang="ts">
import { ADD_LIKE_OR_DISLIKE_COMMENT, ADD_LIKE_OR_DISLIKE_POST, ADD_LIKE_OR_DISLIKE_REVIEW } from '@/services/queries';
import { handleGqlError } from '@/utils/error';
import { useMutation } from '@vue/apollo-composable';
import { ref, watch } from 'vue';
import { useRoute, useRouter } from 'vue-router';


const route = useRoute();
const router = useRouter();


const props = defineProps<{
  id: string;
  idType: string;
  userReaction: string;
  likesCount: number;
  dislikesCount: number;
  fontRemSize: number;
}>();


var userReaction = ref(props.userReaction);
var likesCount = ref(props.likesCount);
var dislikesCount = ref(props.dislikesCount);
var fontRemSize = ref(props.fontRemSize + 'rem');


const addReactionReview = (reaction: string) => {
  const { mutate: addLikeDislike, onError: onErrorAddLikeDislike, onDone: onDoneAddLikeDislike } = useMutation(
    ADD_LIKE_OR_DISLIKE_REVIEW,
    () => ({
      variables: {
        reviewId: props.id,
        action: reaction,
      },
    }
  ));
  
  onErrorAddLikeDislike((err) => {
    handleGqlError(router, err);
  });
  
  onDoneAddLikeDislike((res: any) => {
    if (res.loading) {
      return;
    }
  
    likesCount.value = res.data.addLikeDislikeReview.likesCount;
    dislikesCount.value = res.data.addLikeDislikeReview.dislikesCount;
    userReaction.value = res.data.addLikeDislikeReview.userReaction;
  })

  addLikeDislike();
};

const addReactionComment = (reaction: string) => {
  const { mutate: addLikeDislike, onError: onErrorAddLikeDislike, onDone: onDoneAddLikeDislike } = useMutation(
    ADD_LIKE_OR_DISLIKE_COMMENT,
    () => ({
      variables: {
        commentId: props.id,
        action: reaction,
      },
    }
  ));
  
  onErrorAddLikeDislike((err) => {
    handleGqlError(router, err);
  });
  
  onDoneAddLikeDislike((res: any) => {
    if (res.loading) {
      return;
    }
  
    likesCount.value = res.data.addLikeDislikeComment.likesCount;
    dislikesCount.value = res.data.addLikeDislikeComment.dislikesCount;
    userReaction.value = res.data.addLikeDislikeComment.userReaction;
  })

  addLikeDislike();
};

const addReactionPost = (reaction: string) => {
  const { mutate: addLikeDislike, onError: onErrorAddLikeDislike, onDone: onDoneAddLikeDislike } = useMutation(
    ADD_LIKE_OR_DISLIKE_POST,
    () => ({
      variables: {
        postId: props.id,
        action: reaction,
      },
    }
  ));
  
  onErrorAddLikeDislike((err) => {
    handleGqlError(router, err);
  });
  
  onDoneAddLikeDislike((res: any) => {
    if (res.loading) {
      return;
    }
  
    likesCount.value = res.data.addLikeDislikePost.likesCount;
    dislikesCount.value = res.data.addLikeDislikePost.dislikesCount;
    userReaction.value = res.data.addLikeDislikePost.userReaction;
  })

  addLikeDislike();
};


const addReaction = (reaction: string) => {
  if (props.idType === 'review') {
    addReactionReview(reaction);
  }
  if (props.idType === 'comment') {
    addReactionComment(reaction);
  }
  if (props.idType === 'post') {
    addReactionPost(reaction);
  }
};

watch(props, (newProps) => {
  userReaction.value = newProps.userReaction;
  likesCount.value = newProps.likesCount;
  dislikesCount.value = newProps.dislikesCount;
  fontRemSize.value = newProps.fontRemSize + 'rem';
}, { immediate: true });

</script>

<template>

  <div class="flex my-auto text-neutral-500">
    <div class="cursor-pointer" @click="addReaction('like')">
      <i v-if="userReaction==='like'" class="pi pi-thumbs-up-fill like-dislike-selected"></i>
      <i v-else class="pi pi-thumbs-up like-dislike"></i>
      <span class="ml-1 like-dislike">{{ likesCount }}</span>
    </div>
    <div class="ml-4 cursor-pointer" @click="addReaction('dislike')">
      <i v-if="userReaction==='dislike'" class="pi pi-thumbs-down-fill like-dislike-selected"></i>
      <i v-else class="pi pi-thumbs-down like-dislike"></i>
      <span class="ml-1 like-dislike">{{ dislikesCount }}</span>
    </div>
  </div>

</template>

<style>

.like-dislike {
  font-size: v-bind('fontRemSize');
}

.like-dislike-selected {
  color: var(--p-rating-icon-active-color);;
  font-size: v-bind('fontRemSize');
}

</style>
