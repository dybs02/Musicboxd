<script setup lang="ts">
import LikesDislikes from '@/components/likes-dislikes/LikesDislikes.vue';
import { GET_COMMENT_REPLIES, REPORT_COMMENT } from '@/services/queries';
import { type CommentType } from '@/types/comments';
import { handleGqlError } from '@/utils/error';
import { navigateToUser } from '@/utils/navigate';
import { useMutation, useQuery } from '@vue/apollo-composable';
import mitt from 'mitt';
import Avatar from 'primevue/avatar';
import Button from 'primevue/button';
import Card from 'primevue/card';
import ConfirmPopup from 'primevue/confirmpopup';
import { useConfirm } from 'primevue/useconfirm';
import { useToast } from 'primevue/usetoast';
import { inject, ref } from 'vue';
import { useI18n } from 'vue-i18n';
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
  showLikes: {
    type: Boolean,
    default: true,
    required: false
  },
  nestedLevel: {
    type: Number,
    default: 0,
    required: false
  }
});

const { t } = useI18n()
const router = useRouter();
const confirm = useConfirm();
const toast = useToast();
const emitter = inject<ReturnType<typeof mitt>>('emitter')

let replies = ref<CommentType[]>([]);
const commentText = ref(props.comment.text);
const isTranslated = ref(false);

const translateComment = async () => {
  if (isTranslated.value) {
    commentText.value = props.comment.text;
    isTranslated.value = false;
    return;
  }

  try {
    const targetLanguage = navigator.language.split('-')[0];
    // @ts-ignore 
    const apiKey = import.meta.env.VITE_TRANSLATE_API_KEY;
    // @ts-ignore 
    const location = import.meta.env.VITE_TRANSLATE_API_LOCATION;
    const endpoint = 'https://api.cognitive.microsofttranslator.com/translate?api-version=3.0&to=' + targetLanguage;

    const response = await fetch(endpoint, {
      method: 'POST',
      headers: {
        'Ocp-Apim-Subscription-Key': apiKey,
        'Ocp-Apim-Subscription-Region': location,
        'Content-Type': 'application/json',
      },
      body: JSON.stringify([{ 'Text': props.comment.text }]),
    });

    if (!response.ok) {
      throw new Error('Translation failed');
    }

    const data = await response.json();
    if (data && data[0] && data[0].translations && data[0].translations[0] && data[0].translations[0].text) {
      commentText.value = data[0].translations[0].text;
      isTranslated.value = true;
    } else {
      throw new Error('Translation failed');
    }
  } catch (error) {
    toast.add({ severity: 'error', summary: t('error'), detail: t('translationError'), life: 3000 });
    console.error(error);
  }
};



const confirmReport = (event: any, commentID: string) => {
  confirm.require({
    target: event.currentTarget,
    message: t('reportCommentConfirm'),
    icon: 'pi pi-exclamation-triangle',
    rejectProps: {
      label: t('cancel'),
      severity: 'secondary',
      outlined: true
    },
    acceptProps: {
      label: t('report')
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
      toast.add({ severity: 'info', summary: t('commentReported'), life: 2000 });
    }
  });
};

const reply = () => {
  emitter?.emit('replyToComment', { comment: props.comment });
};

const fetchReplies = () => {
  const { onError, onResult } = useQuery(
    GET_COMMENT_REPLIES,
    {
      commentId: props.comment._id,
    }
  );

  onError((err) => {
    handleGqlError(router, err);
  });

  onResult((res: any) => {
    if (res.loading) {
      return;
    } 

    replies.value = res?.data?.replies;
  })
};

</script>

<template>

  <Card :class="(props.nestedLevel % 2 === 0 ? 'custom-card' : '')">
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
          <Button
            class="mr-2"
            @click="translateComment()"
            v-tooltip.bottom="isTranslated ? $t('showOriginalComment') : $t('translateComment')" 
            icon="pi pi-language"
            aria-label="Translate"
            severity="secondary"
            size="small"
          />
          <Button
            class="mr-2"
            @click="reply()"
            v-tooltip.bottom="$t('replyToComment')" 
            icon="pi pi-reply"
            aria-label="Save"
            severity="secondary"
            size="small"
          />
          <Button
            @click="confirmReport($event, props.comment._id)"
            v-tooltip.bottom="$t('reportComment')" 
            icon="pi pi-flag-fill"
            aria-label="Save"
            severity="secondary"
            size="small"
          />
        </div>
      </div>
    </template>
    <template #content>
      <div class="flex-1 break-words">
        {{ commentText }}
      </div>
      <div v-if="props.showLikes" class="flex justify-content-between pt-4">
        <LikesDislikes
          :id="props.comment._id"
          :idType="'comment'"
          :userReaction="props.comment.userReaction"
          :likesCount="props.comment.likesCount"
          :dislikesCount="props.comment.dislikesCount"
          :fontRemSize="0.7"
        />
        <div class="text-xs text-neutral-500 ml-4">
          {{ props.comment.repliesCount }} {{ props.comment.repliesCount === 1 ? $t('reply') : $t('replies') }}
        </div>
      </div>
    </template>
    <template #footer>
      <div v-if="props.comment.repliesCount > 0 && replies.length == 0" class="mt-2">
        <span class="text-neutral-500 cursor-pointer" @click="fetchReplies">
          {{ $t('view') }} {{ props.comment.repliesCount }} {{ props.comment.repliesCount === 1 ? $t('reply') : $t('replies') }}
        </span>
      </div>
      <div v-if="replies.length > 0" class="mt-2">
        <div v-for="reply in replies" :key="reply._id" class="nested-comment mb-2">
          <Comment
            :comment="reply"
            :nested-level="props.nestedLevel + 1"
          />
        </div>
      </div>
    </template>
  </Card>

</template>

<style scoped>

.custom-card  {
  background: var(--p-darker);
}

</style>