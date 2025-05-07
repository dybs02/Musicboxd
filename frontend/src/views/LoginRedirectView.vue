<script setup lang="ts">
import { useAuthStore } from "@/services/authStore";
import { emptyUser, type UserType } from '@/types/user';
import { handleGqlError } from '@/utils/error';
import { gql } from "@apollo/client/core";
import { useQuery } from '@vue/apollo-composable';
import { onMounted, ref, watch } from 'vue';
import { useRoute, useRouter } from 'vue-router';


const store = useAuthStore();
const route = useRoute();
const router = useRouter();


const fetch_user_avatar = async () => {
  const { loading, error, result } = useQuery(gql`
    query UserById($id: String!) {
      userById(id: $id) {
        images {
          url
        }
      }
    }`,
    {
      id: route.query.userId,
    }
  );

  watch(error, (err) => {
    handleGqlError(router, err);
  });

  watch(result, (newresult) => {
    console.log("User avatar result: ", newresult);
    const newUser = newresult?.userById;
    if (newUser) {
      store.setAvatarUrl(newUser.images[0]?.url ?? "");
    }
  });
};

onMounted(async () => {
  const jwt = route.query.jwt;
  const userId = route.query.userId;
  const userRole = route.query.role;

  if (jwt && userId && userRole) {
    store.setId(userId as string);
    store.setJWT(jwt as string);
    store.setRole(userRole as string);
  }
  
  await fetch_user_avatar();

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
