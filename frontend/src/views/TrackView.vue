<script setup lang="ts">
import Review from "@/components/Review.vue";
import ReviewComments from '@/components/ReviewComments.vue';
import { useAuthStore } from '@/services/authStore';
import { GET_REWIEW_BY_ITEM_ID_USER_ID, GET_TRACK_BY_ID } from "@/services/queries";
import { useQuery } from '@vue/apollo-composable';
import { useMediaQuery } from '@vueuse/core';
import Card from 'primevue/card';
import Divider from 'primevue/divider';
import Image from 'primevue/image';
import ProgressSpinner from 'primevue/progressspinner';
import { computed, ref, watch } from 'vue';
import { useRoute, useRouter } from 'vue-router';

const emptyTrack = {
  duration_ms: 0,
  href: "",
  id: "",
  name: "",
  popularity: 0,
  track_number: 0,
  external_urls: {
    spotify: ""
  },
  artists: [
    {
      external_urls: {
        spotify: ""
      },
      name: "",
      id: "",
    }
  ],
  album: {
    images: {
      url: ""
    },
    name: "",
    id: "",
    total_tracks: 0,
    release_date: "",
    external_urls: {
      spotify: ""
    }
  }
};
// TODO separate file for empty objects
const emptyReview = {
  value: 0,
  title: '',
  description: '',
  comments: [],
};


const store = useAuthStore();
const route = useRoute();
const router = useRouter();
const track = ref<any>(emptyTrack);
const review = ref<any>(emptyReview);
const isMdScreen = useMediaQuery('(max-width: 767px)') // Tailwind md breakpoint

let trackLoading = ref(true);
let reviewLoading = ref(true);

const fetch_track = async () => {
  const { loading, error, result } = useQuery(GET_TRACK_BY_ID, {
    id: route.params.trackId
  });

  trackLoading = loading;

  watch(error, (err) => {
    console.error(err);
    router.push({ 
      name: 'error'
    });
  });

  track.value = computed(() => result?.value?.track ?? emptyTrack);
};

const fetch_rewiew = async () => {
  const { loading, error, result } = useQuery(GET_REWIEW_BY_ITEM_ID_USER_ID, {
    itemId: route.params.trackId,
    userId: store.getId(),
  });

  reviewLoading = loading;

  watch(error, (err) => {
    if (err?.message !== 'mongo: no documents in result') {
      console.error(err);
      router.push({ 
        name: 'error'
      });
    }
  });

  review.value = computed(() => result.value?.review ?? emptyReview);
};

const fetch_data = async () => {
  fetch_track();
  fetch_rewiew();
};

// TODO: move to separate file
const updateComments = (comments: Comments) => {
  // idk if there is a better way to update comments & keep them reactive
  let newReview = {}
  Object.assign(newReview, review.value.value);
  Object.assign(newReview, { comments: comments });
  review.value = computed(() => newReview);
};

watch(() => route.params.id, fetch_data, { immediate: true })

</script>

<template>
  <!-- TODO v-if for error -->
  <div v-if="trackLoading || reviewLoading" class="flex justify-center pt-12">
    <ProgressSpinner />
  </div>

  <div v-else>
    <div class="flex">
      <div class="w-1/2">
        <Image
          :src="track.value.album.images[0].url ?? ''"
          alt="Album Cover"
          class="sm:p-4 drop-shadow-xl"
          preview
          />
      </div>
      <div class="w-1/2 pl-4 sm:px-4 sm:pt-4">
        <Card>
          <template #content>
            <div class="text-3xl sm:text-5xl font-bold">
              <a :href="track.value.external_urls.spotify" target="_blank">
                {{ track.value.name }}
              </a>
            </div>
            <Divider />
            <div class="sm:text-2xl pb-1">
              <a :href="track.value.artists[0].external_urls.spotify" target="_blank">
                {{ track.value.artists[0].name }}
              </a>
            </div>
            <div class="text-sm sm:text-xl">
              <a @click="router.push({ name: 'album', params: { id: track.value.album.id } });" class="cursor-pointer" target="_blank">
                {{ track.value.album.name }}
              </a>
              <a class="text-slate-500">
                ({{ track.value.album.release_date.split("-")[0] }})
              </a>
            </div>
          </template>
        </Card>
        <div v-if="!isMdScreen" >
          <Review
            :item-id="route.params.trackId as string"
            :item-type="'track'"
            :rating="review.value.value"
            :title="review.value.title"
            :description="review.value.description"
            class="pt-4"
          />
        </div>
      </div>
    </div>

    <div v-if="isMdScreen" class="sm:px-4 sm:pt-4">
      <Review
        :item-id="route.params.trackId as string"
        :item-type="'track'"
        :rating="review.value.value"
        :title="review.value.title"
        :description="review.value.description"
        class="pt-4"
      />
    </div>

    <ReviewComments
      v-if="review.value !== emptyReview"
      :item-id="route.params.trackId as string"
      :comments="review.value.comments"
      @update-comments="updateComments"
      class="sm:px-4 mt-4"
    />
  </div>

</template>

<style scoped>

</style>
