<script setup lang="ts">
import { CREATE_UPDATE_REWIEW_BY_ITEM_ID } from '@/services/queries';
import type { ReviewType } from '@/types/review';
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
import LikesDislikes from '@/components/likes-dislikes/LikesDislikes.vue';


const route = useRoute();
const router = useRouter();
const props = defineProps<{
  review: ReviewType;
  itemId: string;
  itemType: string;
}>();

const reviewEditable = ref(props.review.title == '');
let review = ref({
  itemId: props.itemId,
  itemType: props.itemType,
  value: props.review.value,
  title: props.review.title,
  description: props.review.description,
});


const { mutate: updateReview, error: updateReviewError, onDone: updateReviewOnDone } = useMutation(
  CREATE_UPDATE_REWIEW_BY_ITEM_ID,
  () => ({
    variables: {
      itemId: review.value.itemId,
      itemType: review.value.itemType,
      value: review.value.value,
      title: review.value.title,
      description: review.value.description,
    },
  }
));

watch(updateReviewError, (err) => {
  handleGqlError(router, err);
});

updateReviewOnDone((res) => {
  router.push({
    name: 'review',
    params: {
      id: res?.data?.createOrUpdateReview._id,
    },
  });
});

const submitReview = () => {
  if (review.value.title == '' || review.value.description == '' || review.value.value == 0) {
    // TODO add toast component - not updated
    return;
  }

  if (
    props.review.title == review.value.title &&
    props.review.description == review.value.description &&
    props.review.value == review.value.value
  ) {
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
          <a @click="navigateToUser(router, props.review.user._id)" class="cursor-pointer">
            {{ props.review.user.displayName }}
          </a>
        </span>
      </template>
      <template #subtitle>
        <div class="flex mb-1">
          <div class="mr-2">
            <Avatar
            :image="props.review.user.images[0].url"
            @click="navigateToUser(router, props.review.user._id)"
            class="mr-2 cursor-pointer"
            size="large"
            shape="circle"
            />
          </div>
          <div>
            <span class="text-2xl font-bold">
              {{ props.review.title }}
            </span>
            <Rating v-model="props.review.value" :stars="5" readonly></Rating>
          </div>
        </div>
      </template>
      <template #content>
        <p class="m-0">
            {{ props.review.description }}
        </p>
      </template>
      <template #footer>
        <div class="flex justify-between mt-2">
          <LikesDislikes
            :id="props.review._id"
            :idType="'review'"
            :userReaction="props.review.userReaction"
            :likesCount="props.review.likesCount"
            :dislikesCount="props.review.dislikesCount"
            :fontRemSize="1.3"
          />
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
          <Rating v-model="review.value" :stars="5" class="mt-3"></Rating>
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
