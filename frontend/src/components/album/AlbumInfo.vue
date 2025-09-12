<script setup lang="ts">
import SpotifyLogo from '@/assets/icons/SpotifyLogo.vue';
import type { AlbumType } from '@/types/spotify';
import { navigateToAlbum } from '@/utils/navigate';
import Card from 'primevue/card';
import Divider from 'primevue/divider';
import Rating from 'primevue/rating';
import { useRouter } from 'vue-router';


const router = useRouter();


const props = defineProps<{
  album: AlbumType;
}>();


const avgRating = Math.round(props.album.averageRating);
const avgRatingFormated = props.album.averageRating.toFixed(2);

</script>

<template>

  <Card>
    <template #content>
      <div class="text-3xl sm:text-5xl">
        <div class="cursor-pointer w-fit" @click="navigateToAlbum(router, props.album.id)">
          {{ props.album.name }}
          <span class="text-slate-500">
            ({{ props.album.release_date.split("-")[0] }})
          </span>
          <i class="pi pi-external-link text-gray-400"></i>
        </div>
      </div>
      <Divider />
      <div class="flex flex-col sm:flex-row">
        <div class="sm:text-2xl pb-1 sm:w-2/3">
          <div v-for="(artist, index) in props.album.artists" class="flex">
            <div :href="artist.external_urls.spotify" target="_blank" class="cursor-pointer flex pt-1">
              <SpotifyLogo class="h-5 pr-2 my-auto text-gray-400"/>
              {{ artist.name }}{{ index < props.album.artists.length - 1 ? ',  ' : '' }}
            </div>
          </div>
        </div>
        <div class="sm:w-1/3 sm:text-center flex flex-col items-center pt-2">
          <span class="text-slate-500">
            {{ $t('averageUserRating') }}
          </span>
          <div class="pt-2">
            <Rating v-model="avgRating" :stars="5" readonly disabled></Rating>
          </div>
          <span class="text-xl">
            {{ avgRatingFormated }}
          </span>
          <span class="text-slate-500 pt-2">
            {{ $t('ratedBy') }}
          </span>
          <span class="text-xl">
            {{ $t('ratedByUser', props.album.ratingCount || 0) }}
          </span>
        </div>
      </div>
      <div class="pt-8">
      </div>
      <div class="flex">
        

      </div>
    </template>
  </Card>

</template>

<style scoped>

</style>
