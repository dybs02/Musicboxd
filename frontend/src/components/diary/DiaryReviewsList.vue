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


const props = defineProps({
  reviews: {
    type: Object as () => ReviewType[],
    required: true,
  },
});


const router = useRouter();


const isMdScreen = useMediaQuery('(max-width: 767px)') // Tailwind md breakpoint

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
      <div class="flex flex-wrap items-center justify-between gap-2">
        <span class="text-xl font-bold">Your reviews</span>
      </div>
    </template>
    <Column header="Date" class="w-24">
      <template #body="slotProps">
        <div class="text-xs text-neutral-500 flex flex-col">
          <span class="mx-auto">
            2nd May
          </span>
          <span class="mx-auto">
            2023
          </span>
        </div>
      </template>
    </Column>
    <Column header="Cover" class="w-24 h-24">
      <template #body="slotProps">
        <img :src="slotProps.data.album.images[0].url" class="rounded" />
      </template>
    </Column>
    <Column v-if="!isMdScreen" header="Type" class="w-2">
      <template #body="slotProps">
        <Tag :value="slotProps.data.itemType" :severity="getSeverity(slotProps.data.itemType)" />
      </template>
    </Column>
    <Column v-if="!isMdScreen" header="Album name" class="">
      <template #body="slotProps">
        <span>
          {{ slotProps.data.album.name }}
        </span>
      </template>
    </Column>
    <Column v-if="!isMdScreen" field="rating" header="Reviews" class="w-2">
      <template #body="slotProps">
        <div class="flex" @click="navigateToReview(router, slotProps.data.itemType, slotProps.data.itemId, slotProps.data.userId)">
          <Rating :modelValue="slotProps.data.value" readonly/>
          <i class="pi pi-external-link ml-4 cursor-pointer"></i>
        </div>
      </template>
    </Column>
    <Column v-if="isMdScreen" header="Review" class="">
      <template #body="slotProps">
        <div class="flex flex-col" @click="navigateToReview(router, slotProps.data.itemType, slotProps.data.itemId, slotProps.data.userId)">
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

    <template #footer> Showing {{ reviews.length }} records </template>
  </DataTable>

</template>

<style scoped>

</style>
