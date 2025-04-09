<script setup lang="ts">
import AlbumInfo from '@/components/album/AlbumInfo.vue';
import TrackList from "@/components/album/TrackList.vue";
import Review from "@/components/Review.vue";
import ReviewComments from '@/components/ReviewComments.vue';
import { useAuthStore } from "@/services/authStore";
import { GET_ALBUM_BY_ID, GET_REWIEW_BY_ITEM_ID_USER_ID } from "@/services/queries";
import { useQuery } from '@vue/apollo-composable';
import { useMediaQuery } from '@vueuse/core';
import Image from 'primevue/image';
import ProgressSpinner from 'primevue/progressspinner';
import { computed, reactive, ref, watch } from 'vue';
import { useRoute, useRouter } from 'vue-router';


const emptyAlbum = {
  total_tracks: 0,
  name: "",
  release_date: "",
  images: [
    {
      url: null
    }
  ],
  external_urls: {
    spotify: ""
  },
  artists: [
    {
      external_urls: {
        spotify: ""
      },
      name: "",
      id: "",
      href: "",
      type: ""
    }
  ],
  track_list: []
};
const emptyReview = {
  value: 0,
  title: '',
  description: '',
  comments: [],
};


const store = useAuthStore();
const route = useRoute();
const router = useRouter();
const album = ref<any>(emptyAlbum);
const review = ref<any>(emptyReview);
const isMdScreen = useMediaQuery('(max-width: 767px)') // Tailwind md breakpoint

let albumLoading = ref(true);
let reviewLoading = ref(true);

const fetch_album = async () => {
  const { loading, error, result } = useQuery(GET_ALBUM_BY_ID, {
    id: route.params.albumId
  });

  albumLoading = loading;

  watch(error, (err) => {
    console.error(err);
    router.push({ 
      name: 'error'
    });
  });

  album.value = computed(() => result?.value?.album ?? emptyAlbum);
};

const fetch_rewiew = async () => {

  const { loading, error, result } = useQuery(GET_REWIEW_BY_ITEM_ID_USER_ID, {
    itemId: route.params.albumId,
    userId: store.getId(),
  });

  reviewLoading = loading;

  watch(error, (err) => {
    if (err?.message !== 'mongo: no documents in result') {
      console.error(err);
      router.push({ 
        name: 'error'
      });
    }
  });
  
  review.value = computed(() => result.value?.review ?? emptyReview);
};


const fetch_data = async () => {
  fetch_album();
  fetch_rewiew();
};


watch(() => route.params, fetch_data, { immediate: true })

const updateComments = (comments: Comments) => {
  // idk if there is a better way to update comments & keep them reactive
  let newReview = {}
  Object.assign(newReview, review.value.value);
  Object.assign(newReview, { comments: comments });
  review.value = computed(() => newReview);
};

</script>

<template>
  <div v-if="albumLoading || reviewLoading" class="flex justify-center pt-12">
    <ProgressSpinner />
  </div>

  <div v-else>
    <div class="flex">
      <div class="w-1/2">
        <Image
          :src="album.value.images[0].url ?? ''"
          alt="Album Cover"
          class="sm:p-4 drop-shadow-xl"
          preview
          />
      </div>
      <div class="w-1/2 pl-4 sm:px-4 sm:pt-4">
        <AlbumInfo
          :album="album.value"
        />
        <div v-if="!isMdScreen" >
          <TrackList
            :track_list="album.value.track_list"
            class="mt-4"
          />
          <Review
            :item-id="route.params.albumId as string"
            :item-type="'album'"
            :rating="review.value.value"
            :title="review.value.title"
            :description="review.value.description"
            class="pt-4"
          />
        </div>
      </div>
    </div>
    <div v-if="isMdScreen" class="sm:px-4 sm:pt-4">
      <TrackList
        :track_list="album.value.track_list"
        class="mt-4"
      />
      <Review
        :item-id="route.params.albumId as string"
        :item-type="'album'"
        :rating="review.value.value"
        :title="review.value.title"
        :description="review.value.description"
        class="pt-4"
      />
    </div>

    <ReviewComments
      v-if="review.value !== emptyReview"
      :item-id="route.params.albumId as string"
      :comments="review.value.comments"
      @update-comments="updateComments"
      class="sm:px-4 mt-4"
    />
  </div>

</template>

<style scoped>

</style>
