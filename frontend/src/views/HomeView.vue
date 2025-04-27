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

let recentReviews = ref<ReviewType[]>(emptyRecentReviews);
let recentReviewsLoading = ref(true);


const fetch_recent_reviews = async () => {
  const { loading, error, result } = useQuery(
    RECENT_REVIEWS,
    { 
      number: recentReviewsNumber
    },
    {
      fetchPolicy: 'network-only'
    }
  );

  recentReviewsLoading = loading;

  watch(error, (err) => {
    handleGqlError(router, err);
    // console.error(err);
    // router.push({ 
    //   name: 'error'
    // });
  });

  recentReviews = computed<ReviewType[]>(() => result?.value?.recentReviews ?? emptyRecentReviews);
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
