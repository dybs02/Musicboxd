<script setup lang="ts">
import { ADD_LIKE_OR_DISLIKE } from '@/services/queries';
import { handleGqlError } from '@/utils/error';
import { useMutation } from '@vue/apollo-composable';
import { ref } from 'vue';
import { useRoute, useRouter } from 'vue-router';


const route = useRoute();
const router = useRouter();


const props = defineProps<{
  itemId: string;
  userReaction: string;
  likesCount: number;
  dislikesCount: number;
  fontRemSize: number;
}>();


var userReaction = ref(props.userReaction);
var likesCount = ref(props.likesCount);
var dislikesCount = ref(props.dislikesCount);
var fontRemSize = ref(props.fontRemSize + 'rem');


const addReaction = (reaction: string) => {
  const { mutate: addLikeDislike, onError: onErrorAddLikeDislike, onDone: onDoneAddLikeDislike } = useMutation(
    ADD_LIKE_OR_DISLIKE,
    () => ({
      variables: {
        reviewId: route.params.id,
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
