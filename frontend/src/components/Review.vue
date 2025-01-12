<script setup lang="ts">
import Rating from 'primevue/rating';
import Textarea from 'primevue/textarea';
import FloatLabel from "primevue/floatlabel";
import Button from 'primevue/button';
import { Form } from '@primevue/forms';

const props = defineProps<{
  rating:  number;
  title: string;
  description: string;
}>();

// TODO update review on submit

</script>

<template>

  <!-- TODO add non editable version and edit button -->
  <div>
    <Form>
      <Rating v-model="props.rating" :stars="5"></Rating>
      <FloatLabel variant="on" class="mt-4">
        <Textarea v-model="props.title" id="over_label" rows="1" class="w-full" />
        <label for="on_label">Rewiew title</label>
      </FloatLabel>
      <FloatLabel variant="on" class="mt-4">
        <Textarea v-model="props.description" id="over_label" rows="5" class="w-full" />
        <label for="on_label">Rewiew description</label>
      </FloatLabel>
      <div class="flex mt-2">
        <Button type="submit" severity="secondary" label="Submit" class="ml-auto" />
      </div>
    </Form>
  </div>

</template>

<style>

</style>
