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
import type { PostType } from '@/types/posts';


const route = useRoute();
const router = useRouter();

const props = defineProps<{
  post: PostType;
}>();



</script>

<template>

  <Card class="mb-3">
    <template #header>
      <div class="flex align-items-center mx-4 mt-4">
        <div class="my-auto cursor-pointer" @click="navigateToUser(router, props.post.user._id)">
          <Avatar :image="props.post.user.images[0].url" class="mr-2" size="normal" shape="circle" />
        </div>
        <div class="block">
          <span class="cursor-pointer" @click="navigateToUser(router, props.post.user._id)">
            {{ props.post.user.displayName }}
          </span>
          <br>
          <span class="text-xs text-neutral-500">
            {{ new Intl.DateTimeFormat('en-US', {
              year: 'numeric',
              month: 'short',
              day: 'numeric',
              hour: 'numeric',
              minute: 'numeric'
            }).format(new Date(props.post.updatedAt)) }}
          </span>
        </div>
      </div>
    </template>
    <template #content>
      <p>{{ props.post.content }}</p>
    </template>
    <template #footer>
      <LikesDislikes
        :id="props.post._id"
        idType="post"
        :userReaction="props.post.userReaction"
        :likesCount="props.post.likesCount"
        :dislikesCount="props.post.dislikesCount"
        :fontRemSize="1.2"
      />
    </template>
  </Card>

</template>

<style>

</style>
