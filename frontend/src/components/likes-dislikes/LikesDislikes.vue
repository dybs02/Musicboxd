<script setup lang="ts">
import { ADD_LIKE_OR_DISLIKE, CREATE_UPDATE_REWIEW_BY_ITEM_ID } from '@/services/queries';
import type { ReviewType } from '@/types/review';
import type { UserType } from '@/types/user';
import { handleGqlError } from '@/utils/error';
import { navigateToUser } from '@/utils/navigate';
import { Form } from '@primevue/forms';
import { useMutation } from '@vue/apollo-composable';
import Avatar from 'primevue/avatar';
import Button from 'primevue/button';
import Card from 'primevue/card';
import FloatLabel from "primevue/floatlabel";
import Rating from 'primevue/rating';
import Textarea from 'primevue/textarea';
import { ref, watch } from 'vue';
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
        itemId: props.itemId,
        userId: route.params.userId,
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
