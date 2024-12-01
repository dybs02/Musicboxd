<script setup lang="ts">
import { gql } from "@apollo/client/core";
import { useQuery } from '@vue/apollo-composable';
import AutoComplete from 'primevue/autocomplete';
import Button from 'primevue/button';
import Menubar from 'primevue/menubar';
import { computed, ref } from 'vue';
import { useRouter } from 'vue-router';



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
          router.push({name: 'user', params: {id: 123}});
      }
  },
  {
      label: 'About',
      icon: 'pi pi-envelope',
      command: () => {
          router.push({name: 'about'});
      }
  }
]);
const suggest_items = ref([]);

const suggest_search = async (event: any) => {
  // TODO query more fields and pass selected item to View so as not to make another query?
  const { loading, result } = useQuery(gql`
    query Search($_query: String!, $_type: String) {
      search(query: $_query, type: $_type) {
        tracks {
          items {
            id
            name
          }
        }
      }
    }
    `, {
      _query: search_value.value,
      _type: 'track'
    }
  )

  while (loading.value) {
    await new Promise(resolve => setTimeout(resolve, 100))
  }

  const search_res = computed(() => result.value?.search.tracks.items ?? [])
  suggest_items.value = search_res.value
}

const login = () => {
  window.location.href = 'http://localhost:8080/v1/api/auth/login'
}

</script>

<template>
  <Menubar :model="nav_items">
    <template #start>
      <div class="w-10 h-10">
        <img alt="Vue logo" src="https://picsum.photos/640" />
      </div>
    </template>
    <template #end>
      <div class="flex items-center gap-2">
        <AutoComplete placeholder="Search"
                      v-model="search_value"
                      :suggestions="suggest_items"
                      @complete="suggest_search"
        >
          <template #option="slotProps">
            <div class="flex items-center">
                <img :alt="slotProps.option.name" src="https://picsum.photos/640" class="w-6 mr-2" />
                <div>{{ slotProps.option.name }}</div>
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
