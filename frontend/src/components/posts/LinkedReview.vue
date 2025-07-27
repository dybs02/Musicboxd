<script setup lang="ts">
import { type ReviewType } from '@/types/review';
import { navigateToReview, navigateToUser } from '@/utils/navigate';
import Avatar from 'primevue/avatar';
import Button from 'primevue/button';
import Card from 'primevue/card';
import { defineEmits } from 'vue';
import { useRouter } from 'vue-router';


const router = useRouter();

const props = defineProps({
  review: {
    type: Object as () => ReviewType,
    required: true
  },
  showCloseButton: {
    type: Boolean,
    default: true,
    required: false
  },
});

const emit = defineEmits<{
  closeLinkedReview: [];
}>()

</script>

<template>

  <Card class="mt-4" style="background: var(--p-darker);">
    <template #header>
      <div class="flex items-center ml-4 mt-4">
        <i class="pi pi-bookmark mr-2"></i>
        <span class="text-lg">{{ $t('linkedReviewBy') }}</span>
        <div class="ml-2 flex items-center cursor-pointer" @click="navigateToUser(router, props.review.user._id)">
          <Avatar :image="props.review.user.images[0].url" class="mr-2" size="normal" shape="circle" />
          <span class="font-bold">{{ props.review.user.displayName }}</span>
        </div>
        <div v-if="props.showCloseButton" class="ml-auto mr-4">
          <Button
            @click="emit('closeLinkedReview');"
            v-tooltip.bottom="$t('removeLinkedReview')" 
            icon="pi pi-times"
            aria-label="remove-review"
            severity="secondary"
            size="small"
          />
        </div>
      </div>
    </template>
    <template #content>
      <div class="flex items-center cursor-pointer" @click="navigateToReview(router, props.review._id)">
        <img :src="props.review.album.images[0].url" alt="Review Image" class="w-16 h-16 mr-4 rounded">
        <div class="flex-1 min-w-0 break-words">
          <h3 class="text-xl font-bold block break-words">{{ props.review.title }}</h3>
          <div class="text-sm block">{{ props.review.description }}</div>
        </div>
      </div>
    </template>
  </Card>

</template>

<style scoped>

</style>
