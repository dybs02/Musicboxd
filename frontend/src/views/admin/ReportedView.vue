<script setup lang="ts">
import Comment from '@/components/comments/Comment.vue';
import { useAuthStore } from '@/services/authStore';
import { GET_REPORTED_COMMENTS, RESOLE_REPORTED_COMMENT } from '@/services/queries_admin';
import type { ReportedCommentType } from '@/types/moderator';
import { navigateToUser } from '@/utils/navigate';
import { useMutation, useQuery } from '@vue/apollo-composable';
import Button from 'primevue/button';
import Card from 'primevue/card';
import ConfirmPopup from 'primevue/confirmpopup';
import FloatLabel from "primevue/floatlabel";
import ProgressSpinner from 'primevue/progressspinner';
import Textarea from 'primevue/textarea';
import { useConfirm } from 'primevue/useconfirm';
import { useToast } from 'primevue/usetoast';
import { computed, ref, watch } from 'vue';
import { useRoute, useRouter } from 'vue-router';



const store = useAuthStore();
const route = useRoute();
const router = useRouter();
const confirm = useConfirm();
const toast = useToast();


let reportedComments = ref<ReportedCommentType[]>([]);
let reportedCommentsLoading = ref(true);


// TODO refresh comments after mutation also do this in reviewcomments?? - or it might be just caching issue
const fetch_reported_comments = async () => {
  const { loading, error, result } = useQuery(
    GET_REPORTED_COMMENTS,
    {
      // TODO add pagination
      number: 10,
    }
  );

  reportedCommentsLoading = loading;

  watch(error, (err) => {
    if (err?.message === 'no reported comments found') {
      console.log('no reported comments in DB');
      return;
    }

    console.error(err);
    router.push({ 
      name: 'error'
    });
  });

  reportedComments = computed<ReportedCommentType[]>(() => result?.value?.reportedComments ?? []);
}



const resolveComment = (event: any, id: string, status: string, notes: string) => {
  confirm.require({
    target: event.currentTarget,
    message: `Are you sure you want to mark this comment as ${status}?`,
    icon: 'pi pi-exclamation-triangle',
    rejectProps: {
      label: 'Cancel',
      severity: 'secondary',
      outlined: true
    },
    acceptProps: {
      label: 'Confirm'
    },
    accept: () => {
      const { mutate: resolveComment, error: onError, onDone: successfulMutation} = useMutation(
        RESOLE_REPORTED_COMMENT,
        () => ({
          variables: {
            id: id,
            status: status,
            notes: notes ?? '',
          },
        }
      ));

      watch(onError, (err) => {
        console.error(err);
        router.push({ 
          name: 'error'
        });
      });

      successfulMutation((result: any) => {
        toast.add({ severity: 'info', summary: `Comment marked as ${status}`, life: 2000 });
      });

      resolveComment()
    }
  });
};



const fetch_data = async () => {
  fetch_reported_comments();
};

watch(() => route.params, fetch_data, { immediate: true })

</script>

<template>

  <div v-if="reportedCommentsLoading" class="flex justify-center pt-12">
    <ProgressSpinner />
  </div>

  <div v-else>
    <div v-if="reportedComments.length === 0" class="flex justify-center pt-12">
      <h1 class="text-2xl">No reported comments</h1>
    </div>
    <div v-for="c in reportedComments">
      <Card class="mb-4">
        <template #content>
          <div>
            <Comment :comment="c.comment" :show-report-button="false"/>
          </div>

          <div class="flex justify-between mt-4">
            <div>

              <div class="text-sm text-neutral-500">
                Reporter:
                <span class="cursor-pointer text-neutral-200" @click="navigateToUser(router, c.reportedByUser._id)">
                  {{ c.reportedByUser.displayName }}
                </span>
              </div>

              <div class="text-sm text-neutral-500">
                Date of report: 
                <span class="text-xs text-neutral-500">
                  {{ new Intl.DateTimeFormat('en-US', {
                    year: 'numeric',
                    month: 'short',
                    day: 'numeric',
                    hour: 'numeric',
                    minute: 'numeric'
                  }).format(new Date(c.createdAt)) }}
                </span>
              </div>

            </div>
            
            <div class="mx-auto">
              <FloatLabel variant="on">
                
                <Textarea v-model="c.moderatorNotes" id="over_label" rows="2" class="w-full flex-grow block width-full" />
                <label for="on_label">Moderator notes</label>
              </FloatLabel>
            </div>

            <div>
              <ConfirmPopup></ConfirmPopup>
              <Button @click="resolveComment($event, c._id, 'deleted', c.moderatorNotes)" label="Delete" icon="pi pi-trash" class="p-button-danger mr-4" />
              <Button @click="resolveComment($event, c._id, 'ignored', c.moderatorNotes)" label="Ignore" icon="pi pi-check" class="p-button-success" />
            </div>
          </div>
        </template>
      </Card>
    </div>
  </div>

</template>

<style scoped>

</style>
