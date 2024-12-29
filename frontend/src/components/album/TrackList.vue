<script setup lang="ts">
import Accordion from 'primevue/accordion';
import AccordionPanel from 'primevue/accordionpanel';
import AccordionHeader from 'primevue/accordionheader';
import AccordionContent from 'primevue/accordioncontent';
import { useRouter } from 'vue-router';


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
  <Accordion unstyled>
    <AccordionPanel value="0" unstyled>
      <AccordionHeader unstyled toggleicon="null" class="pb-1">
        Track List 
      </AccordionHeader>
      <AccordionContent unstyled class="bg-neutral-800 rounded-md p-2">
        <div v-for="track in props.track_list" class="m-0">
          <a @click="router.push({ name: 'track', params: { id: track.id } });" class="cursor-pointer">
            {{ track.name }}
          </a>
        </div>
      </AccordionContent>
    </AccordionPanel>
  </Accordion>
</template>

<style scoped>

</style>
