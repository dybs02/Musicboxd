<script setup lang="ts">
import type { AlbumType } from '@/types/spotify';
import Card from 'primevue/card';
import Divider from 'primevue/divider';
import Rating from 'primevue/rating';


const props = defineProps<{
  album: AlbumType;
}>();


const avgRating = Math.round(props.album.averageRating);
const avgRatingFormated = props.album.averageRating.toFixed(2);

</script>

<template>

  <Card>
    <template #content>
      <div class="text-3xl sm:text-5xl font-bold">
        <a :href="props.album.external_urls.spotify" target="_blank">
          {{ props.album.name }}
        </a>
        <a class="text-slate-500">
          ({{ props.album.release_date.split("-")[0] }})
        </a>
      </div>
      <Divider />
      <div class="flex flex-col sm:flex-row">
        <div class="sm:text-2xl pb-1 sm:w-2/3">
          <div v-for="(artist, index) in props.album.artists" class="inline">
            <a :href="artist.external_urls.spotify" target="_blank">
              {{ artist.name }}{{ index < props.album.artists.length - 1 ? ',  ' : '' }}
            </a>
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
