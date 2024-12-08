<script setup lang="ts">
import { gql } from "@apollo/client/core";
import { useQuery } from '@vue/apollo-composable';
import Divider from 'primevue/divider';
import Image from 'primevue/image';
import ProgressSpinner from 'primevue/progressspinner';
import Card from 'primevue/card';
import { computed, ref, watch } from 'vue';
import { useRoute, useRouter } from 'vue-router';

const route = useRoute();
const router = useRouter();
const track = ref<any>(null);
var _loading = ref(true);


const fetch_track = async () => {
  const { loading, result } = useQuery(gql`
    query Search {
      track(id: "${route.params.id}") {
        duration_ms
        explicit
        href
        id
        name
        popularity
        track_number
        type
        uri
        external_urls {
          spotify
        }
        artists {
            external_urls {
              spotify
            }
            name
            id
            href
            type
        }
        album {
            images {
                url
                height
                width
            }
            release_date
            name
            type
            id
            href
            total_tracks
            external_urls {
              spotify
            }
        }
      }
    }`
  )

  _loading = loading;
  track.value = computed(() => result.value.track ?? {});
};


watch(() => route.params.id, fetch_track, { immediate: true })

</script>

<template>
  <!-- TODO v-if for error -->
  <div v-if="_loading" class="flex justify-center pt-12">
    <ProgressSpinner />
  </div>

  <div v-else class="flex">
    <div class="w-1/2">
      <Image
        :src="track.value.album.images[0].url ?? ''"
        alt="Album Cover"
        class="sm:p-4 drop-shadow-xl"
        preview
        />
    </div>
    <div class="w-1/2 px-4 pt-4">
      <!-- <Card>
        <template #content> -->
          <div class="text-3xl sm:text-5xl font-bold">
            <a :href="track.value.external_urls.spotify" target="_blank">
              {{ track.value.name }}
            </a>
          </div>
          <Divider />
          <div class="sm:text-2xl pb-1">
            <a :href="track.value.artists[0].external_urls.spotify" target="_blank">
              {{ track.value.artists[0].name }}
            </a>
          </div>
          <div class="text-sm sm:text-xl">
            <a @click="router.push({ name: 'album', params: { id: track.value.album.id } });" class="cursor-pointer" target="_blank">
              {{ track.value.album.name }}
            </a>
            <a class="text-slate-500">
              ({{ track.value.album.release_date.split("-")[0] }})
            </a>
          </div>
        <!-- </template>
      </Card> -->
    </div>
  </div>

</template>

<style scoped>

</style>
