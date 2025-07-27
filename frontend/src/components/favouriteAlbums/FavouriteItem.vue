<script setup lang="ts">
import { useAuthStore } from '@/services/authStore';
import { UPDATE_CURRENT_USER_FAVOURITE_ALBUM } from '@/services/queries';
import type { AlbumType } from '@/types/spotify';
import type { FavouriteAlbumEntryType } from '@/types/user';
import { handleGqlError } from '@/utils/error';
import { navigateToAlbum } from '@/utils/navigate';
import { useMutation } from '@vue/apollo-composable';
import Button from 'primevue/button';
import Dialog from 'primevue/dialog';
import { defineProps, ref, watch } from 'vue';
import { useRoute, useRouter } from 'vue-router';
import SearchItem from './SearchItem.vue';



const props = defineProps({
  album: {
    type: Object as () => FavouriteAlbumEntryType,
    required: true,
  },
});


const store = useAuthStore();
const route = useRoute();
const router = useRouter();

const visible = ref(false);
let selectedAlbum = ref<AlbumType | null>(null);


const select_option = async (album: AlbumType) => {
  selectedAlbum.value = album;
};

const saveFavouriteAlbum = async () => {
  if (!selectedAlbum.value) {
    return;
  }

  const { mutate: updateFavouriteAlbum, error, onDone: successfulMutation} = useMutation(
    UPDATE_CURRENT_USER_FAVOURITE_ALBUM,
    () => ({
      variables: {
        favouriteAlbum: {
          key: props.album.key,
          albumId: selectedAlbum?.value?.id,
        },
      },
    }
  ));

  watch(error, (err) => {
    visible.value = false;
    selectedAlbum.value = null;
    handleGqlError(router, err);
  });

  successfulMutation((result: any) => {
    console.log(result);
    visible.value = false;
    selectedAlbum.value = null;

    router.go(0);
    // TODO or instead of reloading the page, update the album in the parent component by emit
    // props.album.album = result.data.updateCurrentUser.favouriteAlbums.find(
    //   (entry: FavouriteAlbumEntryType) => entry.key === props.album.key
    // );
  });

  updateFavouriteAlbum();
}


const isLoggedUserProfile = () => {
  return route.params.id === store.getId();
};

</script>

<template>
  <div class="border rounded shadow">
    <div class="relative">
      <div v-if="album.album">
        <img :src="album.album.images[0].url as string" alt="Album Cover" />
      </div>
      <div v-else class="bg-darker aspect-square justify-center items-center flex flex-col">
      </div>
      <div class="opacity-0 hover:opacity-100 duration-100 absolute inset-0 z-10 flex justify-center items-center text-9xl text-neutral-100 font-semibold">
        <i v-if="isLoggedUserProfile()" class="pi pi-plus overlay-icon" @click="visible = true"></i>
        <!-- TODO maybe navigate to review if exists ? -->
        <i v-else class="pi pi-link overlay-icon" @click="navigateToAlbum(router, album.album.albumId)"></i>
      </div>
    </div>
  </div>

  <Dialog v-model:visible="visible" modal :header="$t('searchForAlbum')" :style="{ width: '25rem' }">
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
    <div class="flex justify-end gap-2 pt-4">
      <Button type="button" :label="$t('cancel')" severity="secondary" @click="visible = false"></Button>
      <Button type="button" :label="$t('save')" @click="saveFavouriteAlbum"></Button>
    </div>
</Dialog>



</template>

<style scoped>

.overlay-icon {
  font-size: 4rem;
  color: white; /* TOOD fix color */
}

</style>
