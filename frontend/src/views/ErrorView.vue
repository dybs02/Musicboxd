<script setup lang="ts">
import { useRoute } from 'vue-router';



const route = useRoute();
const message = route.params.message;
const cause = route.params.cause;



</script>

<template>
  <div class="flex flex-col items-center h-screen">
    <h1 class="mx-auto pt-10 text-5xl">❗Error❗</h1>
    <div class="mx-auto pt-10 text-2xl">
      <p class="text-center">Message: {{ message }}</p>
      <p class="text-center">Cause: {{ cause }}</p>
      <br />
      <p class="text-center">Please try again later.</p>
    </div>
  </div>
</template>

<style scoped>

</style>
