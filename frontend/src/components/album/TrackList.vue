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
      <Accordion unstyled>
        <AccordionPanel value="0" unstyled>
          <AccordionHeader unstyled toggleicon="null" class="pb-1">
            Track List 
          </AccordionHeader>
          <AccordionContent unstyled class="bg-neutral-800 rounded-md p-2">
            <div v-for="track in props.track_list" class="m-0">
              <a @click="navigateToTrack(router, track.id, store.getId())" class="cursor-pointer">
                {{ track.name }}
              </a>
            </div>
          </AccordionContent>
        </AccordionPanel>
      </Accordion>
    </template>
  </Card>
</template>

<style scoped>

</style>
