<script setup lang="ts">
import { useAuthStore } from '@/services/authStore';
import { navigateToTrack } from '@/utils/navigate';
import Accordion from 'primevue/accordion';
import AccordionContent from 'primevue/accordioncontent';
import AccordionHeader from 'primevue/accordionheader';
import AccordionPanel from 'primevue/accordionpanel';
import Card from 'primevue/card';
import { useRouter } from 'vue-router';


const store = useAuthStore();
const router = useRouter();

const props = defineProps<{
  track_list: {
    name: string;
    duration_ms: number;
    id: string;
    external_urls: { spotify: string };
  }[];
}>();

</script>

<template>
  <Card>
    <template #content>
      <Accordion >
        <AccordionPanel value="0" >
          <AccordionHeader>
            {{ $t('trackList') }} 
          </AccordionHeader>
          <AccordionContent class="bg-darker">
            <div v-for="(track, index) in props.track_list" :key="track.id">
              <div @click="navigateToTrack(router, track.id)" class="cursor-pointer mb-2">
                <span class="font-bold">
                  {{ index + 1 }}. 
                </span>
                <span>
                  {{ track.name }} - 
                </span>
                <span class="text-sm text-neutral-600">
                  {{ Math.floor(track.duration_ms / 60000) }}:{{ Math.floor((track.duration_ms % 60000) / 1000).toString().padStart(2, '0') }}
                </span>
              </div>
            </div>
          </AccordionContent>
        </AccordionPanel>
      </Accordion>
    </template>
  </Card>
</template>

<style scoped>

</style>
