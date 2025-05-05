<script setup lang="ts">
import Review from "@/components/Review.vue";
import ReviewComments from '@/components/comments/ReviewComments.vue';
import { useAuthStore } from '@/services/authStore';
import { GET_REWIEW_BY_ITEM_ID_USER_ID, GET_TRACK_BY_ID } from "@/services/queries";
import { emptyReview, type CommentType, type ReviewType } from "@/types/review";
import { emptyTrack, type TrackType } from "@/types/spotify";
import { handleGqlError } from "@/utils/error";
import { navigateToAlbum } from "@/utils/navigate";
import { useQuery } from '@vue/apollo-composable';
import { useMediaQuery } from '@vueuse/core';
import Card from 'primevue/card';
import Divider from 'primevue/divider';
import Image from 'primevue/image';
import ProgressSpinner from 'primevue/progressspinner';
import { computed, ref, watch } from 'vue';
import { useRoute, useRouter } from 'vue-router';


const store = useAuthStore();
const route = useRoute();
const router = useRouter();
const isMdScreen = useMediaQuery('(max-width: 767px)') // Tailwind md breakpoint


let track = ref<TrackType>(emptyTrack);
let trackLoading = ref(true);

let review = ref<ReviewType>(emptyReview);
let reviewLoading = ref(true);



const fetch_track = async () => {
  const { loading, error, result } = useQuery(
    GET_TRACK_BY_ID,
    {
      id: route.params.trackId
    }
  );

  trackLoading = loading;

  watch(error, (err) => {
    handleGqlError(router, err);
  });

  track = computed<TrackType>(() => result?.value?.track ?? emptyTrack);
};

const fetch_rewiew = async () => {
  const userId = route.params.userId ?? store.getId();

  const { loading, error, result } = useQuery(
    GET_REWIEW_BY_ITEM_ID_USER_ID,
    {
      itemId: route.params.trackId,
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
  fetch_track();
  fetch_rewiew();
};


watch(() => route.params.id, fetch_data, { immediate: true })


// TODO: move to separate file
const updateComments = (comments: CommentType[]) => {
  // idk if there is a better way to update comments & keep them reactive
  let newReview = emptyReview
  Object.assign(newReview, review.value.value);
  Object.assign(newReview, { comments: comments });
  review = computed<ReviewType>(() => newReview);
};

</script>

<template>
  <!-- TODO v-if for error -->
  <div v-if="trackLoading || reviewLoading" class="flex justify-center pt-12">
    <ProgressSpinner />
  </div>

  <div v-else>
    <div class="flex">
      <div class="w-1/2">
        <Image
          :src="track.album.images[0].url ?? ''"
          alt="Album Cover"
          class="sm:p-4 drop-shadow-xl"
          preview
          />
      </div>
      <div class="w-1/2 pl-4 sm:px-4 sm:pt-4">
        <Card>
          <template #content>
            <div class="text-3xl sm:text-5xl font-bold">
              <a :href="track.external_urls.spotify" target="_blank">
                {{ track.name }}
              </a>
            </div>
            <Divider />
            <div class="sm:text-2xl pb-1">
              <a :href="track.artists[0].external_urls.spotify" target="_blank">
                {{ track.artists[0].name }}
              </a>
            </div>
            <div class="text-sm sm:text-xl">
              <a @click="navigateToAlbum(router, track.album.id, store.getId())" class="cursor-pointer">
                {{ track.album.name }}
              </a>
              <a class="text-slate-500">
                ({{ track.album.release_date.split("-")[0] }})
              </a>
            </div>
          </template>
        </Card>
        <div v-if="!isMdScreen" >
          <Review
            :item-id="route.params.trackId as string"
            :item-type="'track'"
            :rating="review.value"
            :title="review.title"
            :description="review.description"
            class="pt-4"
          />
        </div>
      </div>
    </div>

    <div v-if="isMdScreen" class="sm:px-4 sm:pt-4">
      <Review
        :item-id="route.params.trackId as string"
        :item-type="'track'"
        :rating="review.value"
        :title="review.title"
        :description="review.description"
        class="pt-4"
      />
    </div>

    <ReviewComments
      v-if="review !== emptyReview"
      :item-id="route.params.trackId as string"
      :comments="review.comments"
      @update-comments="updateComments"
      class="sm:px-4 mt-4"
    />
  </div>

</template>

<style scoped>

</style>
