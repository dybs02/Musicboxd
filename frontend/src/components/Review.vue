<script setup lang="ts">
import { CREATE_UPDATE_REWIEW_BY_ITEM_ID } from '@/services/queries';
import { Form } from '@primevue/forms';
import { useMutation } from '@vue/apollo-composable';
import Button from 'primevue/button';
import Card from 'primevue/card';
import FloatLabel from "primevue/floatlabel";
import Rating from 'primevue/rating';
import Textarea from 'primevue/textarea';
import { ref, watch } from 'vue';
import { useRoute, useRouter } from 'vue-router';


const route = useRoute();
const router = useRouter();
const props = defineProps<{
  rating:  number;
  title: string;
  description: string;
}>();
const reviewEditable = ref(props.title == '');

let review = ref({
  rating: props.rating,
  title: props.title,
  description: props.description,
});


const { mutate: updateReview, error: updateReviewError, onDone: updateReviewOnDone } = useMutation(
  CREATE_UPDATE_REWIEW_BY_ITEM_ID,
  () => ({
    variables: {
      itemId: route.params.albumId,
      value: review.value.rating,
      title: review.value.title,
      description: review.value.description,
    },
  }
));

watch(updateReviewError, (err) => {
  console.error(err);
  router.push({ 
    name: 'error' 
  });
});

updateReviewOnDone(() => {
  // TODO refetch review data in parent, reload component and add toast component
  // or just update values here
  router.go(0);
});

const submitReview = () => {
  console.log(review.value);
  if (review.value.title == '' || review.value.description == '' || review.value.rating == 0) {
    // TODO add toast component - not updated
    return;
  }

  if (props.title == review.value.title && props.description == review.value.description && props.rating == review.value.rating) {
    toggleEdit();
    return;
  }

  updateReview();
};

const toggleEdit = () => {
  reviewEditable.value = !reviewEditable.value;
};

</script>

<template>

  <div v-if="!reviewEditable">
    <Card>
      <template #title>
        {{ props.title }}
      </template>
      <template #subtitle>
        <Rating v-model="props.rating" :stars="5" readonly></Rating>
      </template>
      <template #content>
        <p class="m-0">
            {{ props.description }}
        </p>
      </template>
      <template #footer>
        <Button type="button" severity="secondary" label="Edit" @click="toggleEdit"/>
      </template>
    </Card>
  </div>

  <div v-else>
    <Card class="">
      <template #content>
        <Form>
          <FloatLabel variant="on">
            <Textarea v-model="review.title" id="over_label" rows="1" class="w-full" />
            <label for="on_label">Rewiew title</label>
          </FloatLabel>
          <Rating v-model="review.rating" :stars="5" class="mt-3"></Rating>
          <FloatLabel variant="on" class="mt-4">
            <Textarea v-model="review.description" id="over_label" rows="5" class="w-full" />
            <label for="on_label">Rewiew description</label>
          </FloatLabel>
          <div class="flex mt-2">
            <Button type="button" severity="secondary" label="Cancel" class="ml-auto" @click="toggleEdit"/>
            <Button type="submit" severity="secondary" label="Submit" class="ml-2" @click="submitReview"/>
          </div>
        </Form>
      </template>
    </Card>
  </div>

</template>

<style>

</style>
