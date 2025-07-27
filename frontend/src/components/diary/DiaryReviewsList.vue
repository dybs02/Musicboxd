<script setup lang="ts">
import type { ReviewType } from '@/types/review';
import DataTable from 'primevue/datatable';
import Column from 'primevue/column';
import Tag from 'primevue/tag';
import Rating from 'primevue/rating';
import Card from 'primevue/card';
import { useMediaQuery } from '@vueuse/core';
import { navigateToReview } from '@/utils/navigate';
import { useRouter } from 'vue-router';
import Button from 'primevue/button';

const props = defineProps({
  reviews: {
    type: Object as () => ReviewType[],
    required: true,
  },
});


const router = useRouter();


const isMdScreen = useMediaQuery('(max-width: 550px)') // Tailwind md breakpoint

const getSeverity = (itemType: string) => {
  switch (itemType) {
    case 'album':
      return 'success';

    case 'track':
      return 'warn';

    default:
        return 'info';
  }
};


</script>

<template>

  <DataTable :value="reviews">
    <template #header>
    </template>
    <Column :header="$t('date')" class="w-24">
      <template #body="slotProps">
        <div class="flex flex-col">
          <span class="text-xs text-neutral-500 mx-auto">
            {{ new Intl.DateTimeFormat('en-US', {
              month: 'short',
              day: 'numeric',
            }).format(new Date(slotProps.data.createdAt)) }}
          </span>
          <span class="text-sm text-neutral-500 mx-auto">
            {{ new Intl.DateTimeFormat('en-US', {
              year: 'numeric',
            }).format(new Date(slotProps.data.createdAt)) }}
          </span>
        </div>
      </template>
    </Column>
    <Column :header="$t('cover')" class="w-24 h-24">
      <template #body="slotProps">
        <img :src="slotProps.data.album.images[0].url" class="rounded" />
      </template>
    </Column>
    <Column v-if="!isMdScreen" :header="$t('type')" class="w-2">
      <template #body="slotProps">
        <Tag :value="$t(slotProps.data.itemType)" :severity="getSeverity(slotProps.data.itemType)" />
      </template>
    </Column>
    <Column v-if="!isMdScreen" :header="$t('albumName')" class="">
      <template #body="slotProps">
        <span>
          {{ slotProps.data.album.name }}
        </span>
      </template>
    </Column>
    <Column v-if="!isMdScreen" field="rating" :header="$t('reviews')" class="w-2">
      <template #body="slotProps">
        <div class="flex" @click="navigateToReview(router, slotProps.data._id)">
          <Rating :modelValue="slotProps.data.value" readonly/>
          <i class="pi pi-external-link ml-4 cursor-pointer"></i>
        </div>
      </template>
    </Column>
    <Column v-if="isMdScreen" :header="$t('review')" class="">
      <template #body="slotProps">
        <div class="flex flex-col" @click="navigateToReview(router, slotProps.data._id)">
          <div class="mx-auto">
            <div class="flex">
              <Tag :value="slotProps.data.itemType" :severity="getSeverity(slotProps.data.itemType)" />
              <Rating :modelValue="slotProps.data.value" readonly class="ml-2"/>
            </div>
          </div>
          <span class="mt-2">
            {{ slotProps.data.album.name }}
          </span>
        </div>
      </template>
    </Column>

    <template #footer> {{ $t('showingRecords', { count: reviews.length }) }} </template>
  </DataTable>

</template>

<style scoped>

</style>
