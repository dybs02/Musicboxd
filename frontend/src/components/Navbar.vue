<script setup lang="ts">
import { ref } from 'vue';
import { useRouter } from 'vue-router';
import AutoComplete from 'primevue/autocomplete';
import Button from 'primevue/button';
import Menubar from 'primevue/menubar';



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
const suggest_items = ref(['']);


const suggest_search = (event: any) => {
  console.log(event.query);
  suggest_items.value = [
    'Apple',
    'Banana',
    'Orange'
  ];
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
                      />
        <Button label="Login" severity="secondary" @click="login" />
      </div>
    </template>
  </Menubar>
</template>

<style scoped>

</style>
