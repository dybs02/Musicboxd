<script setup lang="ts">
import { useTheme } from '@/primevue/theme';
import Toast from 'primevue/toast';
import { RouterView } from 'vue-router';
import Navbar from './components/Navbar.vue';

const { setTheme } = useTheme();
setTheme()

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

  <div class="pt-16">
  </div>

</template>

<style scoped>

</style>
