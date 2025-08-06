<script setup lang="ts">
import { useRoute } from 'vue-router';



const route = useRoute();
const message = route.params.message;
const cause = route.params.cause;



</script>

<template>
  <div class="flex flex-col items-center h-screen">
    <h1 class="mx-auto pt-10 text-5xl">❗{{ $t('error') }}❗</h1>
    <div class="mx-auto pt-10 text-2xl">
      <p class="text-center">{{ $t('message') }}: {{ message }}</p>
      <p class="text-center">{{ $t('cause') }}: {{ cause }}</p>
      <br />
      <p class="text-center">{{ $t('tryAgainLater') }}</p>
    </div>
  </div>
</template>

<style scoped>

</style>
