<script setup lang="ts">
import { gql } from "@apollo/client/core";
import { useQuery } from '@vue/apollo-composable';
import { computed, ref, watch } from 'vue';
import { useRoute } from 'vue-router';

const route = useRoute();
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
  <div v-if="_loading">
    <!-- TODO loading animation -->
    <p>Loading...</p>
  </div>
  <div v-else>
    <img :src="track.value.album.images[0].url ?? ''" alt="Album Cover" />
    <p>{{ track.value.name }}</p>
  </div>

</template>

<style scoped>

</style>
