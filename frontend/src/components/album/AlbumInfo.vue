<script setup lang="ts">
import type { AlbumType } from '@/types/spotify';
import Card from 'primevue/card';
import Divider from 'primevue/divider';

const props = defineProps<{
  album: AlbumType;
}>();

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
      <div class="sm:text-2xl pb-1">
        <div v-for="(artist, index) in props.album.artists" class="inline">
          <a :href="artist.external_urls.spotify" target="_blank">
            {{ artist.name }}{{ index < props.album.artists.length - 1 ? ',  ' : '' }}
          </a>
        </div>
      </div>
    </template>
  </Card>

</template>

<style scoped>

</style>
