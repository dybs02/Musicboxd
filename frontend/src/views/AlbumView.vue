<script setup lang="ts">
import { gql } from "@apollo/client/core";
import { useQuery } from '@vue/apollo-composable';
import Divider from 'primevue/divider';
import Image from 'primevue/image';
import ProgressSpinner from 'primevue/progressspinner';
import Accordion from 'primevue/accordion';
import AccordionPanel from 'primevue/accordionpanel';
import AccordionHeader from 'primevue/accordionheader';
import AccordionContent from 'primevue/accordioncontent';
import Card from 'primevue/card';
import { computed, ref, watch } from 'vue';
import { useRoute, useRouter } from 'vue-router';

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
  album.value = computed(() => result.value.album ?? {});
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
    <div class="w-1/2 px-4 pt-4">
      <!-- <Card>
        <template #content> -->
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
          <div class="text-sm sm:text-xl">
            <!-- TODO add track list (lazy load?) -->
          </div>
          <Accordion value="0" unstyled>
          <AccordionPanel value="0" unstyled>
            <AccordionHeader unstyled>Track List </AccordionHeader>
            <AccordionContent unstyled>
              <div v-for="track in album.value.track_list" class="m-0">
                <a @click="router.push({ name: 'track', params: { id: track.id } });" class="cursor-pointer">
                  {{ track.name }}
                </a>
              </div>
            </AccordionContent>
          </AccordionPanel>
        </Accordion>
        <!-- </template>
      </Card> -->
    </div>
  </div>

</template>

<style scoped>

</style>
