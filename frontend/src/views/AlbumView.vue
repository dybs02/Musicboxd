<script setup lang="ts">
import AlbumInfo from '@/components/album/AlbumInfo.vue';
import TrackList from "@/components/album/TrackList.vue";
import Review from "@/components/Review.vue";
import { useAuthStore } from "@/services/authStore";
import { GET_ALBUM_BY_ID, GET_REWIEW_ID_BY_ITEM_ID_USER_ID } from "@/services/queries";
import { emptyReview, type ReviewType } from '@/types/review';
import { emptyAlbum, type AlbumType } from '@/types/spotify';
import { handleGqlError } from '@/utils/error';
import { useQuery } from '@vue/apollo-composable';
import { useMediaQuery } from '@vueuse/core';
import Image from 'primevue/image';
import ProgressSpinner from 'primevue/progressspinner';
import { computed, onMounted, ref, watch } from 'vue';
import { useRoute, useRouter } from 'vue-router';


const store = useAuthStore();
const route = useRoute();
const router = useRouter();
const isMdScreen = useMediaQuery('(max-width: 550px)') // Tailwind md breakpoint

const props = defineProps<{
  itemId: string;
  hideAddReview: boolean;
}>();


let album = ref<AlbumType>(emptyAlbum);
let albumLoading = ref(true);
let reviewLoading = ref(true);

const albumId = computed(() => {
  return props.itemId ?? route.params.id;
});


const fetch_album = async () => {
  const { onError, onResult } = useQuery(
    GET_ALBUM_BY_ID,
    {
      id: albumId.value,
    }
  );

  albumLoading.value = true;

  onError((err) => {
    handleGqlError(router, err);
  });

  onResult((res) => {
    if (res.loading) {
      return;
    }

    album.value = res?.data?.album;
    albumLoading.value = false;
  });
};

const fetch_rewiew = async () => {
  if (route.fullPath.includes('review')) {
    // Already in a review, do not redirect
    return;
  }

  const { onError, onResult } = useQuery(
    GET_REWIEW_ID_BY_ITEM_ID_USER_ID,
    {
      itemId: albumId.value,
      userId: store.getId(),
    }
  );

  reviewLoading.value = true;

  onError((err) => {
    handleGqlError(router, err, ["mongo: no documents in result"]);
  });
  
  onResult((res) => {
    if (res.loading) {
      return;
    }
    
    if (res?.data?.review) {
      router.push({
        name: 'review',
        params: {
          id: res?.data?.review._id,
        },
      });
      return;
    }
  });
};


const fetch_data = async () => {
  fetch_album();
  fetch_rewiew();
};


watch(() => route.params, fetch_data, { immediate: true })

</script>

<template>
  <div v-if="albumLoading" class="flex justify-center pt-12">
    <ProgressSpinner />
  </div>

  <div v-else>
    <div class="sm:flex">
      <div class="sm:w-1/3 m-4 sm:m-0">
        <Image
          :src="album.images[0].url ?? ''"
          alt="Album Cover"
          class="sm:p-4 drop-shadow-xl"
          preview
          />
      </div>
      <div class="sm:w-2/3 sm:pl-4 sm:pt-4">
        <AlbumInfo
          :album="album"
        />
        <div v-if="!isMdScreen" >
          <TrackList
            :track_list="album.tracks.items"
            class="mt-4"
          />
        </div>
      </div>
    </div>
    <div v-if="isMdScreen" class="sm:px-4 sm:pt-4">
      <TrackList
        :track_list="album.tracks.items"
        class="mt-4"
      />
    </div>

    <Review
      v-if="!props.hideAddReview"
      class="pt-4"
      :review="emptyReview"
      :item-id="albumId"
      :item-type="'album'"
    />
  </div>

</template>

<style scoped>

</style>
