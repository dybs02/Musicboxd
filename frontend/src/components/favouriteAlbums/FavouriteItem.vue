<script setup lang="ts">
import type { AlbumType } from '@/types/spotify';
import Button from 'primevue/button';
import Dialog from 'primevue/dialog';
import { defineProps, ref } from 'vue';
import SearchItem from './SearchItem.vue';



const props = defineProps({
  album: {
    type: Object as () => ({
      id: String,
      title: String,
      artist: String,
      cover: String,
    }),
    required: true,
  },
});


const visible = ref(false);
let selectedAlbum = ref<AlbumType | null>(null);


const select_option = async (album: AlbumType) => {
  selectedAlbum.value = album;
};

const saveFavouriteAlbum = async () => {
  if (!selectedAlbum.value) {
    return;
  }

  // TODO save the selected album to the user's favourite albums
}


</script>

<template>
  
  
  <div v-if="album.id == '1'" class="border rounded">
    <div
      class="bg-gray-800 aspect-square justify-center items-center flex flex-col cursor-pointer"
      @click="visible = true"
    >
      <i class="pi pi-plus plus-icon" style="font-size: 2rem"></i>
    </div>
  </div>

  <div v-else class="border rounded shadow">
    <img :src="album.cover as string" alt="Album Cover" />
  </div>

  <Dialog v-model:visible="visible" modal header="Search for album" :style="{ width: '25rem' }">
    <SearchItem
      @select-new-favourite-album="select_option"
    />
    <div v-if="selectedAlbum" class="flex justify-center py-4">
      <img :src="selectedAlbum.images[0].url as string" alt="Album Cover" class="w-1/3" />
      <div class="flex flex-col justify-center items-center w-2/3">
        <span class="text-lg font-bold text-center">
          {{ selectedAlbum.name }}
        </span>
        <span class="text-sm text-center">
          {{ selectedAlbum.artists[0].name }}
        </span>
      </div>
    </div>
    <div class="flex justify-end gap-2">
      <Button type="button" label="Cancel" severity="secondary" @click="visible = false"></Button>
      <Button type="button" label="Save" @click="saveFavouriteAlbum"></Button>
    </div>
</Dialog>



</template>

<style scoped>

.plus-icon {
  font-size: 2rem;
  color: white; /* TOOD fix color */
}

</style>
