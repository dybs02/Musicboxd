<script setup lang="ts">
import { gql } from "@apollo/client/core";
import { useQuery } from '@vue/apollo-composable';
import Divider from 'primevue/divider';
import Image from 'primevue/image';
import ProgressSpinner from 'primevue/progressspinner';
import Card from 'primevue/card';
import { computed, ref, watch } from 'vue';
import { useRoute, useRouter } from 'vue-router';
import TrackList from "@/components/album/TrackList.vue";

const route = useRoute();
const router = useRouter();
const album = ref<any>(null);
var _loading = ref(true);


const fetch_album = async () => {
  const { loading, result } = useQuery(gql`
    query Search {
      album(id: "${route.params.id}") {
        total_tracks
        name
        release_date
        images {
          url
        }
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
        track_list {
          name
          duration_ms
          id
          external_urls {
            spotify
          }
        }
      }
    }`
  )

  _loading = loading;
  album.value = computed(() => result?.value?.album ?? {});
};


watch(() => route.params.id, fetch_album, { immediate: true })

</script>

<template>
  <!-- TODO v-if for error -->
  <div v-if="_loading" class="flex justify-center pt-12">
    <ProgressSpinner />
  </div>

  <div v-else class="flex">
    <div class="w-1/2">
      <Image
        :src="album.value.images[0].url ?? ''"
        alt="Album Cover"
        class="sm:p-4 drop-shadow-xl"
        preview
        />
    </div>
    <div class="w-1/2 pl-4 sm:px-4 sm:pt-4">
      <Card>
        <template #content>
          <div class="text-3xl sm:text-5xl font-bold">
            <a :href="album.value.external_urls.spotify" target="_blank">
              {{ album.value.name }}
            </a>
            <a class="text-slate-500">
              ({{ album.value.release_date.split("-")[0] }})
            </a>
          </div>
          <Divider />
          <div class="sm:text-2xl pb-1">
            <div v-for="(artist, index) in album.value.artists" class="inline">
              <a :href="artist.external_urls.spotify" target="_blank">
                {{ artist.name }}{{ index < album.value.artists.length - 1 ? ',  ' : '' }}
              </a>
            </div>
          </div>
          <div class="hidden sm:block">
            <TrackList :track_list="album.value.track_list" />
          </div>
        </template>
      </Card>
    </div>
  </div>

  <div class="block sm:hidden pt-4">
    <Card>
      <template #content>
        <TrackList :track_list="album.value.track_list" />
      </template>
    </Card>
  </div>

</template>

<style scoped>

</style>
