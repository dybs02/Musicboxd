<script setup lang="ts">
import Comment from '@/components/comments/Comment.vue';
import { useAuthStore } from '@/services/authStore';
import { ADD_COMMENT } from '@/services/queries';
import type { CommentType } from '@/types/comments';
import { Form } from '@primevue/forms';
import { useMutation } from '@vue/apollo-composable';
import Button from 'primevue/button';
import Card from 'primevue/card';
import FloatLabel from "primevue/floatlabel";
import Textarea from 'primevue/textarea';
import { ref, watch } from 'vue';
import { useRoute, useRouter } from 'vue-router';

const store = useAuthStore();
const route = useRoute();
const router = useRouter();
const newComment = ref({
  text: '',
});




const emit = defineEmits<{
  updateComments: [comments: CommentType[]]
}>();
const props = defineProps<{
  itemId: string;
  comments: CommentType[];
}>();



const { mutate: addComment, error: addCommentError, onDone: addCommentOnDone } = useMutation(
  ADD_COMMENT,
  () => ({
    variables: {
      itemId: props.itemId,
      reviewId: route.params.userId,
      text: newComment.value.text,
    },
  }
));

watch(addCommentError, (err) => {
  console.error(err);
  router.push({ 
    name: 'error' 
  });
});

addCommentOnDone((result) => {
  emit('updateComments', result.data?.addComment ?? props.comments);
  newComment.value.text = '';
});


const submitComment = () => {
  if (newComment.value.text == '') {
    // TODO add toast component
    return;
  }

  addComment();
}


</script>

<template>
  <!-- TODO enable editing, delete, reply -->
  <Toast />
  <div>
    <Card>
      <template #title>
        <div class="pb-2">
          Comments
        </div>
      </template>
      <template #content>
        <div v-for="c in props.comments" class="mb-2">
          <Comment :comment="c" />
        </div>

        <div>
          <Card class="bg-comment">
            <template #content>
              <Form class="flex">
                <FloatLabel variant="on" class="w-full">
                  <Textarea v-model="newComment.text" id="over_label" rows="2" class="w-full" />
                  <label for="on_label">Your comment</label>
                </FloatLabel>
                <div class="">
                  <Button type="submit" severity="secondary" label="Add comment" class="ml-2" @click="submitComment"/>
                </div>
              </Form>
            </template>
          </Card>
        </div>
      </template>
    </Card>
  </div>




</template>

<style scoped>

.bg-comment {
  background: var(--color-primary-light);
}

</style>
