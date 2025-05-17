<script setup lang="ts">
import AlbumInfo from '@/components/album/AlbumInfo.vue';
import TrackList from "@/components/album/TrackList.vue";
import ReviewComments from '@/components/comments/ReviewComments.vue';
import Review from "@/components/Review.vue";
import { useAuthStore } from "@/services/authStore";
import { GET_ALBUM_BY_ID, GET_REWIEW_BY_ID, GET_REWIEW_BY_ITEM_ID_USER_ID } from "@/services/queries";
import { type CommentType } from "@/types/comments";
import { emptyReview, type ReviewType } from "@/types/review";
import { emptyAlbum, type AlbumType } from '@/types/spotify';
import { handleGqlError } from '@/utils/error';
import { navigateToAlbum } from '@/utils/navigate';
import { useQuery } from '@vue/apollo-composable';
import { useMediaQuery } from '@vueuse/core';
import Image from 'primevue/image';
import ProgressSpinner from 'primevue/progressspinner';
import { computed, ref, watch } from 'vue';
import { useRoute, useRouter } from 'vue-router';
import AlbumView from './AlbumView.vue';
import TrackView from './TrackView.vue';


const store = useAuthStore();
const route = useRoute();
const router = useRouter();


let review = ref<ReviewType>(emptyReview);
let reviewLoading = ref(true);



const fetch_rewiew = async () => {
  const { onError, onResult } = useQuery(
    GET_REWIEW_BY_ID,
    {
      reviewId: route.params.id as string,
    }
  );

  reviewLoading.value = true;

  onError((err) => {
    handleGqlError(router, err);
  });
  
  onResult((res) => {
    if (res.loading) {
      return;
    }

    review.value = res?.data?.reviewById;
    reviewLoading.value = false;
  });
};


const fetch_data = async () => {
  fetch_rewiew();
};


watch(() => route.params, fetch_data, { immediate: true })

</script>

<template>
  <div v-if="reviewLoading" class="flex justify-center pt-12">
    <ProgressSpinner />
  </div>

  <div v-else>
    <AlbumView
      v-if="review.itemType === 'album'"
      :item-id="review.itemId"
      :hide-add-review="true"
    />

    <TrackView
      v-if="review.itemType === 'track'"
      :item-id="review.itemId"
      :hide-add-review="true"
    />


    <Review
      class="py-4"
      :review="review"
      :item-id="review.itemId"
      :item-type="review.itemType"
    />


    <ReviewComments
      class="sm:px-4"
      :item-id="route.params.albumId as string"
    />
  </div>

</template>

<style scoped>

</style>
