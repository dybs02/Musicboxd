<script setup lang="ts">
import { CREATE_UPDATE_REWIEW_BY_ITEM_ID } from '@/services/queries';
import type { UserType } from '@/types/user';
import { handleGqlError } from '@/utils/error';
import { navigateToUser } from '@/utils/navigate';
import { Form } from '@primevue/forms';
import { useMutation } from '@vue/apollo-composable';
import Avatar from 'primevue/avatar';
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
  itemId: string;
  itemType: string;
  rating:  number;
  title: string;
  description: string;
  user: UserType;
}>();

const reviewEditable = ref(props.title == '');
let review = ref({
  itemId: props.itemId,
  itemType: props.itemType,
  rating: props.rating,
  title: props.title,
  description: props.description,
});


const { mutate: updateReview, error: updateReviewError, onDone: updateReviewOnDone } = useMutation(
  CREATE_UPDATE_REWIEW_BY_ITEM_ID,
  () => ({
    variables: {
      itemId: review.value.itemId,
      itemType: review.value.itemType,
      value: review.value.rating,
      title: review.value.title,
      description: review.value.description,
    },
  }
));

watch(updateReviewError, (err) => {
  handleGqlError(router, err);
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
        <span class="text-sm text-neutral-500"> 
          Review by 
          <a @click="navigateToUser(router, props.user._id)" class="cursor-pointer">
            {{ props.user.displayName }}
          </a>
        </span>
      </template>
      <template #subtitle>
        <div class="flex mb-1">
          <div class="mr-2">
            <Avatar
            :image="props.user.images[0].url"
            @click="navigateToUser(router, props.user._id)"
            class="mr-2 cursor-pointer"
            size="large"
            shape="circle"
            />
          </div>
          <div>
            <span class="text-2xl font-bold">
              {{ props.title }}
            </span>
            <Rating v-model="props.rating" :stars="5" readonly></Rating>
          </div>
        </div>
      </template>
      <template #content>
        <p class="m-0">
            {{ props.description }}
        </p>
      </template>
      <template #footer>
        <div class="float-right">
          <Button type="button" severity="secondary" label="Edit" @click="toggleEdit"/>
        </div>
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
