<script setup lang="ts">
import ReviewPanel from '@/components/ReviewPanel.vue';
import { RECENT_REVIEWS } from '@/services/queries';
import { emptyReview, type Review } from '@/types/reviews';
import { useQuery } from '@vue/apollo-composable';
import ProgressSpinner from 'primevue/progressspinner';
import { ref, watch } from 'vue';
import { useRoute, useRouter } from 'vue-router';


const route = useRoute();
const router = useRouter();

const recentReviewsNumber = 8;
const emptyRecentReviews = Array(recentReviewsNumber).fill(emptyReview);
const recentReviews = ref<Review[]>(emptyRecentReviews);

let recentReviewsLoading = ref(true);


const fetch_recent_reviews = async () => {
  const { loading, error, result } = useQuery(
    RECENT_REVIEWS,
    { number: recentReviewsNumber },
    { fetchPolicy: 'network-only' }
  );

  recentReviewsLoading = loading;

  watch(error, (err) => {
    console.error(err);
    router.push({ 
      name: 'error'
    });
  });

  watch(result, (data) => {
    if (data) {
      recentReviews.value = data.recentReviews ?? [];
    }
  });
};

const fetch_data = async () => {
  fetch_recent_reviews();
};


watch(() => route.params, fetch_data, { immediate: true })

</script>

<template>
  <div v-if="recentReviewsLoading" class="flex justify-center pt-12">
    <ProgressSpinner />
  </div>

  <div v-else>
    <ReviewPanel
      :reviews="recentReviews"
    />
  </div>

</template>

<style scoped>

</style>
