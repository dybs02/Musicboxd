<script setup lang="ts">
import { onMounted } from 'vue';
import { useRoute, useRouter } from 'vue-router';
import { useAuthStore } from "@/services/authStore";


const store = useAuthStore();
const route = useRoute();
const router = useRouter();


onMounted(() => {
  const jwt = route.query.jwt;
  const userId = route.query.userId;
  const userRole = route.query.role;

  if (jwt && userId && userRole) {
    store.setId(userId as string);
    store.setJWT(jwt as string);
    store.setRole(userRole as string);
  }
  
  router.push({ name: 'home' });
});

</script>

<template>
  <div class="flex">
    <h1 class="mx-auto pt-10 text-5xl">Logged in succesfully</h1>
  </div>
</template>

<style scoped>

</style>
