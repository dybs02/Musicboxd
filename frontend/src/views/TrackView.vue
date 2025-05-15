<script setup lang="ts">
import Review from "@/components/Review.vue";
import { useAuthStore } from '@/services/authStore';
import { GET_TRACK_BY_ID } from "@/services/queries";
import { emptyReview } from "@/types/review";
import { emptyTrack, type TrackType } from "@/types/spotify";
import { handleGqlError } from "@/utils/error";
import { navigateToAlbum } from "@/utils/navigate";
import { useQuery } from '@vue/apollo-composable';
import { useMediaQuery } from '@vueuse/core';
import Card from 'primevue/card';
import Divider from 'primevue/divider';
import Image from 'primevue/image';
import ProgressSpinner from 'primevue/progressspinner';
import { computed, ref, watch } from 'vue';
import { useRoute, useRouter } from 'vue-router';


const store = useAuthStore();
const route = useRoute();
const router = useRouter();
const isMdScreen = useMediaQuery('(max-width: 767px)') // Tailwind md breakpoint

const props = defineProps<{
  itemId: string;
  hideAddReview: boolean;
}>();


let track = ref<TrackType>(emptyTrack);
let trackLoading = ref(true);

const trackId = computed(() => {
  console.log('trackId', props.itemId, route.params.id);
  return props.itemId ?? route.params.id;
});


const fetch_track = async () => {
  const { onError, onResult } = useQuery(
    GET_TRACK_BY_ID,
    {
      id: trackId.value,
    }
  );

  trackLoading.value = true;

  onError((err) => {
    handleGqlError(router, err);
  });

  onResult((res) => {
    if (res.loading) {
      return;
    }

    track.value = res?.data?.track;
    trackLoading.value = false;
  });
};


const fetch_data = async () => {
  fetch_track();
};


watch(() => route.params.id, fetch_data, { immediate: true })

</script>

<template>
  <!-- TODO v-if for error -->
  <div v-if="trackLoading" class="flex justify-center pt-12">
    <ProgressSpinner />
  </div>

  <div v-else>
    <div class="flex">
      <div class="w-1/3">
        <Image
          :src="track.album.images[0].url ?? ''"
          alt="Album Cover"
          class="sm:p-4 drop-shadow-xl"
          preview
          />
      </div>
      <div class="w-2/3 pl-4 sm:px-4 sm:pt-4">
        <Card>
          <template #content>
            <div class="text-3xl sm:text-5xl font-bold">
              <a :href="track.external_urls.spotify" target="_blank">
                {{ track.name }}
              </a>
            </div>
            <Divider />
            <div class="sm:text-2xl pb-1">
              <a :href="track.artists[0].external_urls.spotify" target="_blank">
                {{ track.artists[0].name }}
              </a>
            </div>
            <div class="text-sm sm:text-xl">
              <a @click="navigateToAlbum(router, track.album.id)" class="cursor-pointer">
                {{ track.album.name }}
              </a>
              <a class="text-slate-500">
                ({{ track.album.release_date.split("-")[0] }})
              </a>
            </div>
          </template>
        </Card>
      </div>
    </div>

    <Review
      v-if="!props.hideAddReview"
      class="pt-4"
      :review="emptyReview"
      :item-id="trackId"
      :item-type="'track'"
    />
  </div>

</template>

<style scoped>

</style>
