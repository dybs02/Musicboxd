<script setup lang="ts">
import { useAuthStore } from "@/services/authStore";
import { SEARCH } from "@/services/queries";
import { handleGqlError } from "@/utils/error";
import { navigateToAlbum, navigateToTrack } from "@/utils/navigate";
import { extractSearchResults } from "@/utils/searchResults";
import { gql } from "@apollo/client/core";
import { useQuery } from '@vue/apollo-composable';
import Cookies from "js-cookie";
import AutoComplete, { type AutoCompleteOptionSelectEvent } from 'primevue/autocomplete';
import Button from 'primevue/button';
import Menubar from 'primevue/menubar';
import SelectButton from 'primevue/selectbutton';
import { ref } from 'vue';
import { useRouter } from 'vue-router';


const store = useAuthStore();
const router = useRouter();
const search_value = ref('');
const nav_items = ref([
  {
      label: 'Home',
      icon: 'pi pi-home',
      command: () => {
        router.push({name: 'home'});
      }
  },
  {
      label: 'Profile',
      icon: 'pi pi-star',
      command: () => {
        router.push({name: 'user', params: {id: store.getId()}});
      }
  },
  {
      label: 'Diary',
      icon: 'pi pi-book',
      command: () => {
        router.push({name: 'diary', params: {userId: store.getId()}});
      }
  },
  {
      label: 'About',
      icon: 'pi pi-envelope',
      command: () => {
        router.push({name: 'about'});
      }
  },
]);
const admin_nav_items = ref([
  {
      label: 'Admin',
      icon: 'pi pi-pen-to-square',
      command: () => {
          router.push({name: 'reported'});
      }
  }
])
const suggest_items = ref([]);
const search_type = ref('Track');
const search_options = ref(['Track', 'Album']);

const suggest_search = async () => {
  // TODO query more fields and pass selected item to View so as not to make another query?
  const { loading, onError, onResult } = useQuery(
    SEARCH,
    {
      type: search_type.value.toLowerCase(),
      query: search_value.value,
    }
  )

  onError((err) => {
    handleGqlError(router, err);
  });

  onResult((res: any) => {
    if (res.loading) {
      return;
    }

    suggest_items.value = extractSearchResults(search_type.value, res.data);
  })
}

const login = () => {
  // @ts-ignore
  window.location.href = import.meta.env.VITE_BACKEND_URL+'/v1/api/auth/login'

  const userId = Cookies.get('userId');
  if (userId) {
    store.setId(userId);
  }
}

const select_track = (event: AutoCompleteOptionSelectEvent) => {
  if (search_type.value === 'Track')
    navigateToTrack(router, event.value.id, store.getId())
  else if (search_type.value === 'Album')
    navigateToAlbum(router, event.value.id, store.getId())

  search_value.value = ''
}

</script>

<template>
  <Menubar :model="store.isModerator() ? nav_items.concat(admin_nav_items) : nav_items">
    <template #start>
      <div class="w-10 h-10">
        <img :src="store.getAvatarUrl() as string" class="rounded-full"/>
      </div>
    </template>
    <template #end>
      <div class="flex items-center gap-2">
        <AutoComplete placeholder="Search"
                      v-model="search_value"
                      :suggestions="suggest_items"
                      @complete="suggest_search"
                      @focus="suggest_search"
                      @option-select="select_track"
                      class="sm:w-96 block"
                      fluid
                      panelClass="sm:w-96 w-full"
        >
          <template #header>
            <div class="rounded-md" style="background-color: #09090b;">
              <SelectButton v-model="search_type"
                            :options="search_options"
                            @change="suggest_search"
                            class="my-1 pl-2"
              />
            </div>
          </template>
          <template #option="slotProps">
            <div class="flex items-center">
                <img :alt="slotProps.option.name" :src="slotProps.option.images[0].url" class="w-6 mr-2" />
                <div >
                  {{ slotProps.option.name }}
                  <a class="text-slate-500">
                    - {{ slotProps.option.artists[0].name }}
                  </a>
                </div>
            </div>
          </template>
        </AutoComplete>
        <Button label="Login" severity="secondary" @click="login" />
      </div>
    </template>
  </Menubar>
</template>

<style scoped>

</style>
