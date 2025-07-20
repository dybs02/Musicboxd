<script setup lang="ts">
import { useAuthStore } from '@/services/authStore';
import type { ReviewType } from '@/types/review';
import { navigateToReview } from '@/utils/navigate';
import Avatar from 'primevue/avatar';
import Card from 'primevue/card';
import Carousel from 'primevue/carousel';
import Rating from 'primevue/rating';
import { ref } from 'vue';
import { useRouter } from 'vue-router';


const store = useAuthStore();
const router = useRouter();
const props = defineProps<{
  reviews:  ReviewType[];
  title: string;
}>();
const responsiveOptions = ref([
    {
        breakpoint: '960px',
        numVisible: 3,
        numScroll: 1
    },
    {
        breakpoint: '767px',
        numVisible: 2,
        numScroll: 1
    },
    {
        breakpoint: '550px',
        numVisible: 1,
        numScroll: 1
    }
]);

</script>

<template>
  <Card class="bg-comment">
    <template #title>
      <div class="text-2xl font-bold mb-4">
        {{ props.title }}
      </div>
    </template>
    <template #content>
      <Carousel :value="props.reviews" :numVisible="4" :numScroll="2" :responsiveOptions="responsiveOptions">
        <template #item="slotProps">
          <div class="bg-darker rounded-lg mx-auto m-2 p-4 w-[200px] h-[320px]">
            <div class="cursor-pointer" @click="navigateToReview(router, slotProps.data._id)">
              <div class="mb-4">
                <div class="relative mx-auto">
                  <img :src="slotProps.data.album.images[0].url" :alt="slotProps.data.album.name" class="w-full" />
                </div>
              </div>
              <div v-if="slotProps.data.itemType === 'album'" class="mb-1 font-bold text-lg truncate">{{ slotProps.data.album.name }}</div>
              <div v-if="slotProps.data.itemType === 'track'" class="mb-1 font-bold text-lg truncate">{{ slotProps.data.track.name }}</div>
              <div class="mb-1 font-small">{{ slotProps.data.album.artists[0].name }}</div>
            </div>
            <div class="flex mt-4 cursor-pointer" @click="navigateToReview(router, slotProps.data._id)">
              <div class="my-auto mr-2">
                <Avatar
                  :image="slotProps.data.user.images[0].url"
                  class=""
                  size="normal"
                  shape="circle"
                />
              </div>
              <div>
                <Rating
                  v-model="slotProps.data.value"
                  :stars="5"
                  readonly
                  class="pb-1"
                />
                <div class="font-medium italic w-[130px] truncate pb-1">{{ slotProps.data.title }}</div>
              </div>
            </div>
          </div>
        </template>
      </Carousel>
    </template>
  </Card>

</template>

<style scoped>

</style>
