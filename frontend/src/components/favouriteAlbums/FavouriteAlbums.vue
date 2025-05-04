<script setup lang="ts">
import FavouriteItem from '@/components/favouriteAlbums/FavouriteItem.vue';



const albums = [
  {
    id: '1',
    title: 'Album 1',
    artist: 'Artist 1',
    cover: 'https://picsum.photos/640',
  },
  {
    id: '2',
    title: 'Album 2',
    artist: 'Artist 2',
    cover: 'https://picsum.photos/640',
  },
  {
    id: '3',
    title: 'Album 3',
    artist: 'Artist 3',
    cover: 'https://picsum.photos/640',
  },
  {
    id: '4',
    title: 'Album 4',
    artist: 'Artist 4',
    cover: 'https://picsum.photos/640',
  },
];



</script>

<template>
  <h1>Favourite Albums</h1>
  <div class="flex flex-row gap-4">
    <div v-for="album in albums" :key="album.id" class="w-1/4 aspect-square">
      <FavouriteItem :album="album" />
    </div>
  </div>
</template>

<style scoped>

</style>
