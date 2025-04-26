<script setup lang="ts">
import Comment from '@/components/comments/Comment.vue';
import { useAuthStore } from '@/services/authStore';
import { GET_REPORTED_COMMENTS } from '@/services/queries_admin';
import type { ReportedCommentType } from '@/types/comments';
import { useQuery } from '@vue/apollo-composable';
import Button from 'primevue/button';
import Card from 'primevue/card';
import ProgressSpinner from 'primevue/progressspinner';
import { computed, ref, watch } from 'vue';
import { useRoute, useRouter } from 'vue-router';



const store = useAuthStore();
const route = useRoute();
const router = useRouter();

// TODO use this approach in other components
let reportedComments = ref<ReportedCommentType[]>([]);
let reportedCommentsLoading = ref(true);



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
    console.error(err);
    router.push({ 
      name: 'error'
    });
  });

  reportedComments = computed<ReportedCommentType[]>(() => result?.value?.reportedComments ?? []);
}


// TODO move to separate file - used bt Comments.vue as well
const navigateToUser = (userId: string) => {
  router.push({ 
    name: 'user', 
    params: { id: userId }
  });
}




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
                <span class="cursor-pointer text-neutral-200" @click="navigateToUser(c.reportedByUser._id)">
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
            <div>
              <Button label="Delete" icon="pi pi-trash" class="p-button-danger mr-4" />
              <Button label="Ignore" icon="pi pi-check" class="p-button-success" />
            </div>
          </div>
        </template>
      </Card>
    </div>
  </div>

</template>

<style scoped>

</style>
