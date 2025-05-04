<script setup lang="ts">
import { computed, defineProps, ref, watch } from 'vue';
import Dialog from 'primevue/dialog';
import InputText from 'primevue/inputtext';
import AutoComplete, { type AutoCompleteOptionSelectEvent } from 'primevue/autocomplete';

import Button from 'primevue/button';
import type { ReviewAlbum } from '@/types/review';
import { useQuery } from '@vue/apollo-composable';
import { SEARCH_FOR_ALBUMS } from '@/services/queries';
import { handleGqlError } from '@/utils/error';
import { useRouter } from 'vue-router';
import type { AlbumsType, AlbumType } from '@/types/spotify';



const router = useRouter();


const search_value = ref('');
let suggest_items = ref<AlbumType[]>([]);


const emit = defineEmits<{
  (e: 'selectNewFavouriteAlbum', album: AlbumType): void
}>();


const suggest_search = async (event: any) => {
  const { loading, error, result } = useQuery(
    SEARCH_FOR_ALBUMS,
    {
      query: search_value,
    },
  );

  watch(error, (err) => {
    handleGqlError(router, err, ['no reported comments found']);
  });

  suggest_items = computed<AlbumType[]>(() => result?.value?.search?.albums?.items ?? suggest_items.value);
};

const select_option = async (event: any) => {
  emit('selectNewFavouriteAlbum', event.value as AlbumType);
  search_value.value = '';
};


</script>

<template>

  <AutoComplete placeholder="Search"
                v-model="search_value"
                :suggestions="suggest_items"
                @complete="suggest_search"
                @focus="suggest_search"
                @option-select="select_option"
                class=""
                fluid
                panelClass=""
  >
    <template #option="slotProps">
      <div class="flex items-center">
          <img :alt="slotProps.option.name" :src="slotProps.option.images[0].url" class="w-6 mr-2" />
          <div class="w-64">
            {{ slotProps.option.name }}
            <a class="text-slate-500">
              - {{ slotProps.option.artists[0].name }}
            </a>
          </div>
      </div>
    </template>
  </AutoComplete>

</template>

<style scoped>

.plus-icon {
  font-size: 2rem;
  color: white; /* TOOD fix color */
}

</style>
