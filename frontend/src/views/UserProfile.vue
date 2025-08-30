<script setup lang="ts">
import FavouriteAlbums from '@/components/favouriteAlbums/FavouriteAlbums.vue';
import { FOLLOW_USER, GET_USER_BY_ID, UNFOLLOW_USER } from '@/services/queries';
import { emptyUser, type UserType } from '@/types/user';
import { getCountryName } from '@/utils/country';
import { handleGqlError } from '@/utils/error';
import { navigateToChat } from '@/utils/navigate';
import { useMutation, useQuery } from '@vue/apollo-composable';
import Avatar from 'primevue/avatar';
import Button from 'primevue/button';
import Card from 'primevue/card';
import Divider from 'primevue/divider';
import { ref, watch } from 'vue';
import { useRoute, useRouter } from 'vue-router';


const route = useRoute();
const router = useRouter();


let user = ref<UserType>(emptyUser);
let userLoading = ref(true);

let unfollowButtonText = ref("Unfollow");

const fetch_user = async () => {
  const { loading, error, result } = useQuery(
    GET_USER_BY_ID,
    {
      id: route.params.id,
    },
    {
      fetchPolicy: 'cache-and-network'
    }
  );

  userLoading = loading;

  watch(error, (err) => {
    handleGqlError(router, err);
  });

  watch(result, (newResult) => {
    console.log("User data fetched:", newResult);
    if (newResult?.userById) {
      user.value = newResult.userById;
    }

    unfollowButtonText.value = user.value.isFollowing ? "Following" : "Unfollow";
  });
};


const followUser = async () => {
  const { mutate: follow, error: followError, onDone: followOnDone } = useMutation(
    FOLLOW_USER,
    () => ({
      variables: {
        userId: route.params.id,
      },
    }
  ));

  watch(followError, (err) => {
    handleGqlError(router, err);
  });

  followOnDone(async (result) => {
    user.value = {
      ...user.value,
      isFollowing: result.data.followUser.isFollowing,
      isFollower: result.data.followUser.isFollower
    };
  });

  follow();
}

const unfollowUser = async () => {
  const { mutate: unfollow, error: unfollowError, onDone: unfollowOnDone } = useMutation(
    UNFOLLOW_USER,
    () => ({
      variables: {
        userId: route.params.id,
      },
    }
  ));

  watch(unfollowError, (err) => {
    handleGqlError(router, err);
  });

  unfollowOnDone(async (result) => {
    user.value = {
      ...user.value,
      isFollowing: result.data.unfollowUser.isFollowing,
      isFollower: result.data.unfollowUser.isFollower
    };
  });

  unfollow();
}

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
      <div class="sm:flex items-center gap-2 px-8 pt-8">
        <div class="flex items-center">
          <Avatar :image="user.images[0].url" size="xlarge" shape="circle" />
          <div class="flex flex-col pl-4">
            <div class="flex items-center gap-2">
              <span class="text-3xl font-bold">{{ user.displayName }}</span>

            </div>
            <span class="text-sm text-neutral-500">{{ getCountryName(user.country) }}</span>
          </div>
          <div class="pl-4">
            <Button v-if="!user.isFollower" @click="followUser" severity="primary" outlined>{{ $t('follow') }}</Button>
            <Button 
              v-else
              @click="unfollowUser" 
              :severity="unfollowButtonText === 'Unfollow' ? 'danger' : 'primary'" 
              outlined
              @mouseenter="unfollowButtonText = $t('unfollow')"
              @mouseleave="unfollowButtonText = $t('following')"
            >
              {{ unfollowButtonText }}
            </Button>

            <Button @click="() => navigateToChat(router, user._id)" severity="primary" class="sm:ml-2 mt-2 sm:mt-0">
              {{ $t('chat') }}
              <i class="pi pi-external-link ml-1 cursor-pointer"></i>
            </Button>
          </div>
        </div>


        <div class="mx-auto">
        </div>

        <div class="flex sm:flex-col items-center">
          <div class="sm:text-xl">
            <span class=""> {{ user.albumReviewsNumber }} </span> 
            <span class="text-neutral-500 pl-2">
              {{ user.albumReviewsNumber === 1 ? $t('albumReviewSingular') : $t('albumReviewPlural') }}
            </span> 
          </div>
          <div class="sm:text-xl">
            <span class=""> {{ user.trackReviewsNumber }} </span> 
            <span class="text-neutral-500 pl-2">
              {{ user.trackReviewsNumber === 1 ? $t('trackReviewSingular') : $t('trackReviewPlural') }}
            </span> 
          </div>
          <div class="pt-2">
            <Button severity="secondary" @click="() => router.push(`/diary/${user._id}`)">
              {{ user.displayName }}{{ $t('sDiary') }}
              <i class="pi pi-external-link ml-1 cursor-pointer"></i>
            </Button>
          </div>
        </div>
      </div>
    </template>
    <template #content>
      <h2 class="text-2xl font-bold mb-4">{{ $t('favouriteAlbums') }}</h2>
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
  background: linear-gradient(to right, var(--p-darker), transparent);
  z-index: -5; /* Above the image but below other content */
}

.right-shadow {
  position: absolute;
  top: 0;
  right: 0;
  width: 100px;
  height: 100%;
  background: linear-gradient(to left, var(--p-darker), transparent);
  z-index: -5; /* Above the image but below other content */
}
</style>
