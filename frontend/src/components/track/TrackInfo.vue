<script setup lang="ts">
import SpotifyLogo from '@/assets/icons/SpotifyLogo.vue';
import type { AlbumType, TrackType } from '@/types/spotify';
import Card from 'primevue/card';
import Divider from 'primevue/divider';
import { navigateToAlbum, navigateToTrack } from "@/utils/navigate";
import Rating from 'primevue/rating';
import { useRouter } from 'vue-router';
import SpotifyEmbededPlayer from '../SpotifyEmbededPlayer.vue';


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
      <div class="text-3xl sm:text-5xl font-bold" @click="navigateToTrack(router, track.id)">
        <div class="cursor-pointer w-fit">
          {{ track.name }}
          <i class="pi pi-external-link text-gray-400"></i>
        </div>
      </div>
      <Divider />
      <div class="flex flex-col sm:flex-row">
        <div class="sm:w-2/3">
          <div v-for="(artist, index) in props.track.artists" class="flex">
            <div :href="artist.external_urls.spotify" target="_blank" class="cursor-pointer flex pt-1">
              <SpotifyLogo class="h-5 pr-2 my-auto text-gray-400"/>
              {{ artist.name }}{{ index < props.track.artists.length - 1 ? ',  ' : '' }}
            </div>
          </div>
          <div class="text-sm sm:text-xl pt-2">
            <div @click="navigateToAlbum(router, track.album.id)" class="cursor-pointer">
              {{ track.album.name }}
              <span class="text-slate-500">
                ({{ track.album.release_date.split("-")[0] }})
              </span>
              <i class="pi pi-external-link text-gray-400 pl-1"></i>
            </div>
          </div>
          <div class="pt-4">
            <SpotifyEmbededPlayer :track-id="track.id" class="h-[95px] sm:h-[152px]"/>
          </div>
        </div>
        <div class="sm:w-1/3 sm:text-center flex flex-col items-center pt-2">
          <span class="text-slate-500">
            {{ $t('averageUserRating') }}
          </span>
          <div class="pt-2">
            <Rating v-tooltip.bottom="`Average rating: ${avgRatingFormated}`" v-model="avgRating" :stars="5" readonly></Rating>
          </div>
          <span class="text-xl">
            {{ avgRatingFormated }}
          </span>
          <span class="text-slate-500 pt-2">
            {{ $t('ratedBy') }}
          </span>
          <span class="text-xl">
            {{ props.track.ratingCount }} {{ $t('users') }}
          </span>
        </div>
      </div>
    </template>
  </Card>

</template>

<style scoped>

</style>
