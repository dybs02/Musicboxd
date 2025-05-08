<script setup lang="ts">
import DiaryReviewsList from '@/components/diary/DiaryReviewsList.vue';
import { useAuthStore } from "@/services/authStore";
import { GET_RECENT_USER_REVIEWS_PAGINATION } from "@/services/queries";
import { emptyRecentUserReviews, type RecentUserReviewsType } from '@/types/review';
import { handleGqlError } from '@/utils/error';
import { useQuery } from '@vue/apollo-composable';
import Card from 'primevue/card';
import ProgressSpinner from 'primevue/progressspinner';
import { computed, ref, watch } from 'vue';
import { useRoute, useRouter } from 'vue-router';



const store = useAuthStore();
const route = useRoute();
const router = useRouter();


let userReviews = ref<RecentUserReviewsType>(emptyRecentUserReviews);
let userReviewsLoading = ref(true);



const fetch_user_reviews = async () => {
  const { loading, error, result } = useQuery(
    GET_RECENT_USER_REVIEWS_PAGINATION,
    {
      pageSize: 10,
      page: 1,
      itemType: '',
      userId: route.params.userId ?? store.getId(),
    }
  );

  userReviewsLoading = loading;

  watch(error, (err) => {
    handleGqlError(router, err);
  });

  userReviews = computed<RecentUserReviewsType>(() => result?.value?.recentUserReviews ?? emptyRecentUserReviews);
};



const fetch_data = async () => {
  fetch_user_reviews();
};

watch(() => route.params, fetch_data, { immediate: true })

</script>

<template>
  <div v-if="userReviewsLoading" class="flex justify-center pt-12">
    <ProgressSpinner />
  </div>

  <div v-else>
    <Card>
      <template #content>
        <DiaryReviewsList
          :reviews="userReviews.reviews"
        />
      </template>
    </Card>
  </div>

</template>

<style scoped>

</style>
