<script setup lang="ts">
import { useAuthStore } from '@/services/authStore';
import { ADD_COMMENT } from '@/services/queries';
import { Form } from '@primevue/forms';
import { useMutation } from '@vue/apollo-composable';
import Avatar from 'primevue/avatar';
import Button from 'primevue/button';
import Card from 'primevue/card';
import FloatLabel from "primevue/floatlabel";
import Textarea from 'primevue/textarea';
import { ref, watch } from 'vue';
import { useRoute, useRouter } from 'vue-router';
import ConfirmPopup from 'primevue/confirmpopup';
import { useConfirm } from 'primevue/useconfirm';
import { useToast } from 'primevue/usetoast';


const store = useAuthStore();
const route = useRoute();
const router = useRouter();
const newComment = ref({
  text: '',
});


const confirm = useConfirm();
const toast = useToast();


const confirmReport = (event: any, commentID: string) => {
  confirm.require({
    target: event.currentTarget,
    message: 'Are you sure you want report this comment?',
    icon: 'pi pi-exclamation-triangle',
    rejectProps: {
      label: 'Cancel',
      severity: 'secondary',
      outlined: true
    },
    acceptProps: {
      label: 'Report'
    },
    accept: () => {
      // TODO add report comment mutation
      console.log('report comment', commentID);

      toast.add({ severity: 'info', summary: 'Comment reported', life: 2000 });
    }
  });
};


const emit = defineEmits<{
  updateComments: [comments: Comments]
}>();
const props = defineProps<{
  itemId: string;
  comments: Comments;
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


const navigateToUser = (userId: string) => {
  router.push({ 
    name: 'user', 
    params: { id: userId }
  });
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
          <Card class="bg-comment">
            <template #title>
              <div class="flex mb-1">
                <div class="my-auto cursor-pointer" @click="navigateToUser(c.user._id)">
                  <Avatar :image="c.user.images[0].url" class="mr-2" size="normal" shape="circle" />
                </div>
                <div class="block">
                  <span class="cursor-pointer" @click="navigateToUser(c.user._id)">
                    {{ c.user.displayName }}
                  </span>
                  <br>
                  <span class="text-xs text-neutral-500">
                    {{ new Intl.DateTimeFormat('en-US', {
                      year: 'numeric',
                      month: 'short',
                      day: 'numeric',
                      hour: 'numeric',
                      minute: 'numeric'
                    }).format(new Date(c.updatedAt)) }}
                  </span>
                </div>
                <div class="ml-auto">
                  <ConfirmPopup></ConfirmPopup>
                  <Button @click="confirmReport($event, c._id)" icon="pi pi-flag-fill" aria-label="Save" severity="secondary" size="small" />
                </div>
              </div>
            </template>
            <template #content>
              <div style="white-space: pre;">
                {{ c.text }}
              </div>
            </template>
          </Card>
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
