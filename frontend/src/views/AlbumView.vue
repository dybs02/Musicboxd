<script setup lang="ts">
import AlbumInfo from '@/components/album/AlbumInfo.vue';
import TrackList from "@/components/album/TrackList.vue";
import ReviewComments from '@/components/comments/ReviewComments.vue';
import Review from "@/components/Review.vue";
import { useAuthStore } from "@/services/authStore";
import { GET_ALBUM_BY_ID, GET_REWIEW_BY_ITEM_ID_USER_ID } from "@/services/queries";
import { emptyReview, type CommentType, type ReviewType } from '@/types/review';
import { emptyAlbum, type AlbumType } from '@/types/spotify';
import { handleGqlError } from '@/utils/error';
import { navigateToAlbum } from '@/utils/navigate';
import { useQuery } from '@vue/apollo-composable';
import { useMediaQuery } from '@vueuse/core';
import Image from 'primevue/image';
import ProgressSpinner from 'primevue/progressspinner';
import { computed, ref, watch } from 'vue';
import { useRoute, useRouter } from 'vue-router';


const store = useAuthStore();
const route = useRoute();
const router = useRouter();
const isMdScreen = useMediaQuery('(max-width: 767px)') // Tailwind md breakpoint



let album = ref<AlbumType>(emptyAlbum);
let albumLoading = ref(true);

let review = ref<ReviewType>(emptyReview);
let reviewLoading = ref(true);



const fetch_album = async () => {
  const { loading, error, result } = useQuery(
    GET_ALBUM_BY_ID,
    {
      id: route.params.albumId
    }
  );

  albumLoading = loading;

  watch(error, (err) => {
    handleGqlError(router, err);
  });

  album = computed<AlbumType>(() => result?.value?.album ?? emptyAlbum);
};

const fetch_rewiew = async () => {
  const userId = route.params.userId ?? store.getId();

  const { loading, error, result } = useQuery(
    GET_REWIEW_BY_ITEM_ID_USER_ID,
    {
      itemId: route.params.albumId,
      userId: userId,
    }
  );

  reviewLoading = loading;

  watch(error, (err) => {
    handleGqlError(router, err, ['mongo: no documents in result']);

    if (route.params.userId !== store.getId()) {
      navigateToAlbum(
        router,
        route.params.albumId as string,
        store.getId()
      );
    }

  });
  
  review = computed<ReviewType>(() => result.value?.review ?? emptyReview);
};


const fetch_data = async () => {
  fetch_album();
  fetch_rewiew();
};


watch(() => route.params, fetch_data, { immediate: true })


const updateComments = (comments: CommentType[]) => {
  // idk if there is a better way to update comments & keep them reactive
  let newReview = emptyReview
  Object.assign(newReview, review.value.value);
  Object.assign(newReview, { comments: comments });
  review = computed<ReviewType>(() => newReview);
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
          :src="album.images[0].url ?? ''"
          alt="Album Cover"
          class="sm:p-4 drop-shadow-xl"
          preview
          />
      </div>
      <div class="w-1/2 pl-4 sm:px-4 sm:pt-4">
        <AlbumInfo
          :album="album"
        />
        <div v-if="!isMdScreen" >
          <TrackList
            :track_list="album.tracks.items"
            class="mt-4"
          />
          <Review
            :item-id="route.params.albumId as string"
            :item-type="'album'"
            :rating="review.value"
            :title="review.title"
            :description="review.description"
            class="pt-4"
          />
        </div>
      </div>
    </div>
    <div v-if="isMdScreen" class="sm:px-4 sm:pt-4">
      <TrackList
        :track_list="album.tracks.items"
        class="mt-4"
      />
      <Review
        :item-id="route.params.albumId as string"
        :item-type="'album'"
        :rating="review.value"
        :title="review.title"
        :description="review.description"
        class="pt-4"
      />
    </div>

    <ReviewComments
      v-if="review !== emptyReview"
      :item-id="route.params.albumId as string"
      :comments="review.comments"
      @update-comments="updateComments"
      class="sm:px-4 mt-4"
    />
  </div>

</template>

<style scoped>

</style>
