<script setup lang="ts">
import { useQuery } from '@vue/apollo-composable';
import Divider from 'primevue/divider';
import Image from 'primevue/image';
import ProgressSpinner from 'primevue/progressspinner';
import Card from 'primevue/card';
import { computed, ref, watch } from 'vue';
import { useRoute, useRouter } from 'vue-router';
import { GET_TRACK_BY_ID } from "@/services/queries";

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

const route = useRoute();
const router = useRouter();
const track = ref<any>(emptyTrack);

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


const fetch_data = async () => {
  fetch_track();
  // fetch_rewiew();
};

watch(() => route.params.id, fetch_data, { immediate: true })

</script>

<template>
  <!-- TODO v-if for error -->
  <div v-if="trackLoading" class="flex justify-center pt-12"> // TODO add loading OR
    <ProgressSpinner />
  </div>

  <div v-else class="flex">
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
    </div>
  </div>

</template>

<style scoped>

</style>
