<script setup lang="ts">
import { RouterLink, RouterView } from 'vue-router'
import Navbar from './components/Navbar.vue'
import Toast from 'primevue/toast';
</script>

<template>
  <Toast />

  <div class="flex sm:justify-center">
    <div class="max-w-6xl w-full inline-block">
      <Navbar />
      <div class="pt-4 px-4">
        <RouterView/>
      </div>
    </div>
  </div>
</template>

<style scoped>

</style>
