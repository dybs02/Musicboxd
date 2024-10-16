<script setup lang="ts">
import { RouterLink, useRouter } from 'vue-router'
import HamburgerMenu from '@/assets/icons/HamburgerMenu.vue';
import MagnifyingGlass from '@/assets/icons/MagnifyingGlass.vue';


const router = useRouter();

const goHome = () => {
  router.push('/');
}

</script>

<template>
  <div class="flex h-10 py-2 px-4 bg-midnight">
    <HamburgerMenu class="fill-snow md:hidden"/>

    <div @click="goHome" class="px-2 my-auto cursor-pointer">
      Home
    </div>
    
    <div class="w-full"></div>
    
    <MagnifyingGlass class="fill-snow"/>
  </div>
</template>

<style scoped>

</style>
