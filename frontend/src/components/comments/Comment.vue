<script setup lang="ts">
import { REPORT_COMMENT } from '@/services/queries';
import type { CommentType } from '@/types/comments';
import { navigateToUser } from '@/utils/navigate';
import { useMutation } from '@vue/apollo-composable';
import Avatar from 'primevue/avatar';
import Button from 'primevue/button';
import Card from 'primevue/card';
import ConfirmPopup from 'primevue/confirmpopup';
import { useConfirm } from 'primevue/useconfirm';
import { useToast } from 'primevue/usetoast';
import { useRouter } from 'vue-router';

const props = defineProps({
  comment: {
    type: Object as () => CommentType,
    required: true
  },
  showReportButton: {
    type: Boolean,
    default: true,
    required: false
  },
});



const router = useRouter();
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
      // TODO confirm report with a reason
      const { mutate: reportComment, onDone: successfulMutation} = useMutation(
        REPORT_COMMENT,
        () => ({
          variables: {
            id: commentID,
          },
        }
      ));

      successfulMutation((result: any) => {
        console.log(result);
      });

      reportComment()
      toast.add({ severity: 'info', summary: 'Comment reported', life: 2000 });
    }
  });
};

</script>

<template>

  <Card class="bg-comment">
    <template #title>
      <div class="flex mb-1">
        <div class="my-auto cursor-pointer" @click="navigateToUser(router, props.comment.user._id)">
          <Avatar :image="props.comment.user.images[0].url" class="mr-2" size="normal" shape="circle" />
        </div>
        <div class="block">
          <span class="cursor-pointer" @click="navigateToUser(router, props.comment.user._id)">
            {{ props.comment.user.displayName }}
          </span>
          <br>
          <span class="text-xs text-neutral-500">
            {{ new Intl.DateTimeFormat('en-US', {
              year: 'numeric',
              month: 'short',
              day: 'numeric',
              hour: 'numeric',
              minute: 'numeric'
            }).format(new Date(props.comment.updatedAt)) }}
          </span>
        </div>
        <div class="ml-auto" v-if="props.showReportButton">
          <ConfirmPopup></ConfirmPopup>
          <Button @click="confirmReport($event, props.comment._id)" icon="pi pi-flag-fill" aria-label="Save" severity="secondary" size="small" />
        </div>
      </div>
    </template>
    <template #content>
      <div style="white-space: pre;">
        {{ props.comment.text }}
      </div>
    </template>
  </Card>

</template>

<style scoped>

.bg-comment {
  background: var(--color-primary-light);
}

</style>
