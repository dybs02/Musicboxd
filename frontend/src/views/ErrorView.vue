<script setup lang="ts">
import { useAuthStore } from '@/services/authStore';
import { useRoute } from 'vue-router';


const store = useAuthStore();
const route = useRoute();
const message = route.params.message;
const cause = route.params.cause;

const isJWTExpired = message.includes('token is expired') || message.includes('authorization header is missing');
if (isJWTExpired) {
  setTimeout(() => {
    store.logout();
    // @ts-ignore
    window.location.href = import.meta.env.VITE_BACKEND_URL+'/v1/api/auth/login';
  }, 5000);
}

</script>

<template>
  <div v-if="isJWTExpired" class="flex flex-col items-center h-screen">
    <h1 class="mx-auto pt-10 text-2xl">
      {{ $t('sessionExpiredMessage') }}
    </h1>
  </div>

  <div v-else class="flex flex-col items-center h-screen">
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
