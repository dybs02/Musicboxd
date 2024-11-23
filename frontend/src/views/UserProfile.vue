<script setup lang="ts">
import { useQuery } from '@vue/apollo-composable'
import { gql } from "@apollo/client/core";
import { computed } from 'vue';


const userName = "dybs"

const { result } = useQuery(gql`
  query UserByDisplayName($_displayName: String!) {
    userByDisplayName(displayName: $_displayName) {
      country
      displayName
      email
      href
      spotifyId
      product
      type
      uri
      explicitContent {
        filterEnabled
        filterLocked
      }
      externalUrls {
        spotify
      }
      followers {
        href
        total
      }
      images {
        url
        height
        width
      }
    }
  }
`, {
  _displayName: userName
})
const user = computed(() => result.value?.userByDisplayName ?? {})


</script>

<template>
  <h1>User Profile</h1>
  {{ user }}
</template>

<style scoped>

</style>
