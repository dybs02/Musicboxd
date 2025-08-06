<script setup lang="ts">
import DiaryReviewsList from '@/components/diary/DiaryReviewsList.vue';
import { useAuthStore } from "@/services/authStore";
import { GET_RECENT_USER_REVIEWS_PAGINATION } from "@/services/queries";
import { emptyRecentUserReviews, type RecentUserReviewsType } from '@/types/review';
import { handleGqlError } from '@/utils/error';
import { useQuery } from '@vue/apollo-composable';
import Card from 'primevue/card';
import Paginator, { type PageState } from 'primevue/paginator';
import ProgressSpinner from 'primevue/progressspinner';
import { computed, ref, watch } from 'vue';
import { useRoute, useRouter } from 'vue-router';
import SelectButton from 'primevue/selectbutton';
import { useI18n } from 'vue-i18n';



const { t } = useI18n()
const store = useAuthStore();
const route = useRoute();
const router = useRouter();


const pageSize = 10;
const filter = ref({ name: t('filterAll'), value: '' });
const filterOptions = computed(() => [
  { name: t('filterAll'), value: '' },
  { name: t('filterAlbums'), value: 'album' },
  { name: t('filterTracks'), value: 'track' },
]);

let userReviews = ref<RecentUserReviewsType>(emptyRecentUserReviews);
let userReviewsLoading = ref(true);



const fetch_user_reviews = async () => {
  // TODO use this approach everywhere
  const { loading, onError, onResult } = useQuery(
    GET_RECENT_USER_REVIEWS_PAGINATION,
    {
      pageSize: pageSize,
      page: 1,
      itemType: filter.value.value,
      userId: route.params.userId ?? store.getId(),
    }
  );

  userReviewsLoading = loading;

  onError((err) => {
    handleGqlError(router, err);
  });

  onResult((res: any) => {
    if (res.loading) {
      return;
    }

    userReviews.value = res?.data?.recentUserReviews;
  })
};



const change_page = async (event: PageState) => {
  const { onError, onResult } = useQuery(
    GET_RECENT_USER_REVIEWS_PAGINATION,
    {
      pageSize: pageSize,
      page: event.page + 1,
      itemType: filter.value.value,
      userId: route.params.userId ?? store.getId(),
    },
  );

  onError((err) => {
    handleGqlError(router, err);
  });

  onResult((res: any) => {
    if (res.loading) {
      return;
    }

    userReviews.value = res?.data?.recentUserReviews;
  })
};



const fetch_data = async () => {
  fetch_user_reviews();
};

watch(() => route.params, fetch_data, { immediate: true })
watch(filter, (newValue) => {fetch_user_reviews();}, { immediate: true });

</script>

<template>
  <div v-if="userReviewsLoading" class="flex justify-center pt-12">
    <ProgressSpinner />
  </div>

  <div v-else>
    <Card>
      <template #header>
        <div class="flex flex-wrap items-center justify-between mx-6 mt-6">
          <span class="text-xl font-bold">{{ $t('yourReviews') }}</span>
          <SelectButton
            v-model="filter"
            :options="filterOptions"
            optionLabel="name"
          />
        </div>
      </template>
      <template #content>
        <DiaryReviewsList
          :reviews="userReviews.reviews"
        />
      </template>
      <template #footer>
        <Paginator
          :rows="pageSize"
          :totalRecords="userReviews.totalReviews"
          :first="0"
          @page="change_page($event)"
          class="mt-4"
        >
        </Paginator>
      </template>
    </Card>
  </div>

</template>

<style scoped>

</style>
