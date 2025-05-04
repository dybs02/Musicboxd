<script setup lang="ts">
import FavouriteItem from '@/components/favouriteAlbums/FavouriteItem.vue';
import type { FavouriteAlbumEntryType } from '@/types/user';


const props = defineProps({
  favouriteAlbums: {
    type: Object as () => FavouriteAlbumEntryType[],
    required: true
  },
});


</script>

<template>
  <div class="flex flex-row gap-4">
    <div v-for="album in favouriteAlbums" :key="album.key" class="w-1/4 aspect-square">
      <FavouriteItem :album="album" />
    </div>
  </div>
</template>

<style scoped>

</style>
