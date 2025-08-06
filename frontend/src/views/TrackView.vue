<script setup lang="ts">
import Review from "@/components/Review.vue";
import TrackInfo from "@/components/track/TrackInfo.vue";
import { useAuthStore } from '@/services/authStore';
import { GET_REWIEW_ID_BY_ITEM_ID_USER_ID, GET_TRACK_BY_ID } from "@/services/queries";
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
const isMdScreen = useMediaQuery('(max-width: 550px)') // Tailwind md breakpoint

const props = defineProps<{
  itemId: string;
  hideAddReview: boolean;
}>();


let track = ref<TrackType>(emptyTrack);
let trackLoading = ref(true);
let reviewLoading = ref(true);

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

const fetch_rewiew = async () => {
  const { onError, onResult } = useQuery(
    GET_REWIEW_ID_BY_ITEM_ID_USER_ID,
    {
      itemId: trackId.value,
      userId: store.getId(),
    }
  );

  reviewLoading.value = true;

  onError((err) => {
    handleGqlError(router, err, ["mongo: no documents in result"]);
  });
  
  onResult((res) => {
    if (res.loading) {
      return;
    }
    
    if (res?.data?.review) {
      router.push({
        name: 'review',
        params: {
          id: res?.data?.review._id,
        },
      });
      return;
    }
  });
};


const fetch_data = async () => {
  fetch_track();
  fetch_rewiew();
};


watch(() => route.params.id, fetch_data, { immediate: true })

</script>

<template>
  <!-- TODO v-if for error -->
  <div v-if="trackLoading" class="flex justify-center pt-12">
    <ProgressSpinner />
  </div>

  <div v-else>
    <div class="sm:flex">
      <div class="sm:w-1/3 m-4 sm:m-0">
        <Image
          :src="track.album.images[0].url ?? ''"
          alt="Album Cover"
          class="sm:p-4 drop-shadow-xl"
          preview
          />
      </div>
      <div class="sm:w-2/3 sm:pl-4 sm:pl-4 sm:pt-4">
        <TrackInfo
          :track="track"
        />
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
