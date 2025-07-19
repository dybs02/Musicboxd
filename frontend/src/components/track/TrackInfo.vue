<script setup lang="ts">
import type { AlbumType, TrackType } from '@/types/spotify';
import Card from 'primevue/card';
import Divider from 'primevue/divider';
import { navigateToAlbum } from "@/utils/navigate";
import Rating from 'primevue/rating';
import { useRouter } from 'vue-router';


const router = useRouter();


const props = defineProps<{
  track: TrackType;
}>();


const avgRating = Math.round(props.track.averageRating);
const avgRatingFormated = props.track.averageRating.toFixed(2);

</script>

<template>

  <Card>
    <template #content>
      <div class="text-3xl sm:text-5xl font-bold">
        <a :href="track.external_urls.spotify" target="_blank">
          {{ track.name }}
        </a>
      </div>
      <Divider />
      <div class="flex flex-col sm:flex-row">
        <div class="sm:w-2/3">
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
        </div>
        <div class="sm:w-1/3 sm:text-center flex flex-col items-center pt-2">
          <span class="text-slate-500">
            Average user rating:
          </span>
          <div class="pt-2">
            <Rating v-tooltip.bottom="`Average rating: ${avgRatingFormated}`" v-model="avgRating" :stars="5" readonly></Rating>
          </div>
          <span class="text-xl">
            {{ avgRatingFormated }}
          </span>
          <span class="text-slate-500 pt-2">
            Rated by:
          </span>
          <span class="text-xl">
            {{ props.track.ratingCount }} users
          </span>
        </div>
      </div>
    </template>
  </Card>

</template>

<style scoped>

</style>
