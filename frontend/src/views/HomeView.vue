<script setup lang="ts">
import ReviewPanel from '@/components/ReviewPanel.vue';
import { RECENT_REVIEWS } from '@/services/queries';
import { emptyReview, type ReviewType } from '@/types/review';
import { useQuery } from '@vue/apollo-composable';
import ProgressSpinner from 'primevue/progressspinner';
import { computed, ref, watch } from 'vue';
import { useRoute, useRouter } from 'vue-router';
import { handleGqlError } from '@/utils/error';


const route = useRoute();
const router = useRouter();

const recentReviewsNumber = 8;
const emptyRecentReviews = Array(recentReviewsNumber).fill(emptyReview);

let recentAlbumReviews = ref<ReviewType[]>(emptyRecentReviews);
let recentAlbumReviewsLoading = ref(true);

let recentTrackReviews = ref<ReviewType[]>(emptyRecentReviews);
let recentTrackReviewsLoading = ref(true);

// TODO should be moved to the ReviewPanel
const fetch_recent_album_reviews = async () => {
  const { onError, onResult } = useQuery(
    RECENT_REVIEWS,
    { 
      number: recentReviewsNumber,
      itemType: 'album',
    },
    {
      fetchPolicy: 'network-only'
    }
  );

  recentAlbumReviewsLoading.value = true;

  onError((err) => {
    handleGqlError(router, err);
  });

  onResult((res) => {
    if (res.loading) {
      return;
    }

    recentAlbumReviews.value = res?.data?.recentReviews;
    recentAlbumReviewsLoading.value = false;
  });
};

const fetch_recent_track_reviews = async () => {
  const { onError, onResult } = useQuery(
    RECENT_REVIEWS,
    { 
      number: recentReviewsNumber,
      itemType: 'track',
    },
    {
      fetchPolicy: 'network-only'
    }
  );

  recentTrackReviewsLoading.value = true;

  onError((err) => {
    handleGqlError(router, err);
  });

  onResult((res) => {
    if (res.loading) {
      return;
    }

    recentTrackReviews.value = res?.data?.recentReviews;
    recentTrackReviewsLoading.value = false;
  });
};

const fetch_data = async () => {
  fetch_recent_album_reviews();
  fetch_recent_track_reviews();
};


watch(() => route.params, fetch_data, { immediate: true })

</script>

<template>
  <div v-if="recentAlbumReviewsLoading" class="flex justify-center pt-12">
    <ProgressSpinner />
  </div>

    <div v-else>
    <ReviewPanel
      :reviews="recentAlbumReviews"
      title="Recently added album reviews"
    />
    <ReviewPanel
      :reviews="recentTrackReviews"
      title="Recently added track reviews"
      class="mt-4"
    />
  </div>

</template>

<style scoped>

</style>
