<script setup lang="ts">
import { GET_USER_BY_ID } from '@/services/queries';
import { emptyUser, type UserType } from '@/types/user';
import { handleGqlError } from '@/utils/error';
import { useQuery } from '@vue/apollo-composable';
import { computed, ref, watch } from 'vue';
import { useRoute, useRouter } from 'vue-router';
import Panel from 'primevue/panel';
import Avatar from 'primevue/avatar';
import Card from 'primevue/card';
import { getCountryName } from '@/utils/country';
import FavouriteAlbums from '@/components/favouriteAlbums/FavouriteAlbums.vue';
import Divider from 'primevue/divider';


const route = useRoute();
const router = useRouter();


let user = ref<UserType>(emptyUser);
let userLoading = ref(true);


const fetch_user = async () => {
  const { loading, error, result } = useQuery(
    GET_USER_BY_ID,
    {
      id: route.params.id,
    }
  );

  userLoading = loading;

  watch(error, (err) => {
    handleGqlError(router, err);
  });

  user = computed<UserType>(() => result?.value?.userById ?? emptyUser);
};


const fetch_data = async () => {
  fetch_user();
};

watch(() => route.params, fetch_data, { immediate: true })

</script>



<template>

  <Card>
    <template #header>
      <div class="flex items-center gap-2 p-4">
        <Avatar :image="user.images[0].url" size="xlarge" shape="circle" />
        <div class="flex flex-col">
          <span class="text-3xl font-bold">{{ user.displayName }}</span>
          <span class="text-sm text-neutral-500">{{ getCountryName(user.country) }}</span>
        </div>

        <div class="mx-auto">
        </div>

        <div class="flex flex-col items-center">
          <span class="text-neutral-500">Total reviews:</span>
          <!-- TODO add number of total reviews -->
          <span class="text-3xl">12</span> 
        </div>
      </div>
    </template>
    <template #content>
      <h2 class="text-2xl font-bold mb-4">Favourite Albums</h2>
      <Divider />
      <FavouriteAlbums
        :favourite-albums="user.favouriteAlbums"
      />
    </template>
  </Card>




</template>

<style scoped>

</style>
