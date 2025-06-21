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


const getFirstFavAlbumImageUrl = () => {
  return user.value.favouriteAlbums.find(entry => entry.key === 1)?.album?.images[0]?.url ?? "";
};


const fetch_data = async () => {
  fetch_user();
};

watch(() => route.params, fetch_data, { immediate: true })

</script>



<template>
  
  <div class="backgroud-image-wrap">
    <img
    :src="getFirstFavAlbumImageUrl()"
    class="absolute w-full overflow-hidden object-cover opacity-30 -z-10 asdasasdas"
    />
    <div class="left-shadow"></div>
    <div class="right-shadow"></div>
  </div>
    
  <Card class="relative z-10 -mt-4"> 
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
          <span class="text-neutral-500 text-xl">Total reviews:</span>
          <div class="text-xl">
            <span class="text-2xl"> {{ user.albumReviewsNumber }} </span> 
            <span class="text-neutral-500 pl-2">
              {{ user.albumReviewsNumber === 1 ? 'album' : 'albums' }}
            </span> 
          </div>
          <div class="text-xl">
            <span class="text-2xl"> {{ user.trackReviewsNumber }} </span> 
            <span class="text-neutral-500 pl-2">
              {{ user.trackReviewsNumber === 1 ? 'track' : 'tracks' }}
            </span> 
          </div>
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
.backgroud-image-wrap {
  margin: -20px 0px 0px 0px;
  position: relative;
  overflow: hidden;
  height: 300px;
}

.left-shadow {
  position: absolute;
  top: 0;
  left: 0;
  width: 100px;
  height: 100%;
  background: linear-gradient(to right, var(--color-midnight-dark-background), transparent);
  z-index: -5; /* Above the image but below other content */
}

.right-shadow {
  position: absolute;
  top: 0;
  right: 0;
  width: 100px;
  height: 100%;
  background: linear-gradient(to left, var(--color-midnight-dark-background), transparent);
  z-index: -5; /* Above the image but below other content */
}
</style>
