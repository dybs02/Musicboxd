<script setup lang="ts">
import { useAuthStore } from '@/services/authStore';
import { CHAT_SUBSCRIPTION, GET_MESSAGES_BY_CHAT_ID_PAGE, SEND_MESSAGE } from '@/services/queries';
import { emptyMessagesPage, type ChatType, type MessagesPage, type MessageType } from '@/types/message';
import { handleGqlError } from '@/utils/error';
import { Form } from '@primevue/forms';
import { useMutation, useQuery, useSubscription } from '@vue/apollo-composable';
import { useIntersectionObserver } from '@vueuse/core';
import Button from 'primevue/button';
import FloatLabel from "primevue/floatlabel";
import Textarea from 'primevue/textarea';
import { nextTick, onMounted, ref, useTemplateRef, watch } from 'vue';
import { useRoute, useRouter } from 'vue-router';
import Card from 'primevue/card';


const store = useAuthStore();
const route = useRoute();
const router = useRouter();

const props = defineProps<{
  chat: ChatType;
}>();

const root = useTemplateRef<HTMLElement>('root')
const target = useTemplateRef<HTMLDivElement>('loadingMessagesDiv')

const newMessage = ref<string>('')
const messages = ref<MessageType[]>([])
const pageSize = 10;
let page = 1;
let messagesPage = ref<MessagesPage>(emptyMessagesPage);
let messagesPageLoading = ref(true);


const { result, error: errorHandler } = useSubscription(
  CHAT_SUBSCRIPTION,
  {
    chatId: props.chat._id,
  }
)

watch(
  result,
  data => {
    if (data.loading) {
      return;
    }

    messages.value.push(data.messageAdded)
    scrollToBottom();
  }
)

watch(errorHandler, (err) => {
  console.error("Subscription error:", err);
});

const sendMessage = () => {
  if (newMessage.value.trim() === '') {
    return;
  }

  // TODO This should be outside of the function just called here
  const { mutate: addMessage, error: addMessageError, onDone: addMessageOnDone } = useMutation(
    SEND_MESSAGE,
    () => ({
      variables: {
        chatId: props.chat._id,
        content: newMessage.value,
      },
    }
  ));

  watch(addMessageError, (err) => {
    handleGqlError(router, err);
  });

  addMessageOnDone(async (result) => {
    newMessage.value = '';
  });

  addMessage();
}

const change_page = async () => {
  if (page > messagesPage.value.totalPages && messagesPage.value.totalPages !== 0) {
    return;
  }

  const { loading, onError, onResult } = useQuery(
    GET_MESSAGES_BY_CHAT_ID_PAGE,
    {
      chatId: props.chat._id,
      pageSize: pageSize,
      page: page,
    },
    {
      fetchPolicy: 'no-cache',
    }
  );

  messagesPageLoading = loading;

  onError((err) => {
    handleGqlError(router, err);
  });

  onResult((res: any) => {
    if (res.loading) {
      return;
    }

    page += 1;

    const scrollHeight = root.value?.scrollHeight || 0;
    const scrollTop = root.value?.scrollTop || 0;
    const isAtBottom = scrollTop + root.value!.clientHeight >= scrollHeight - 10;
    
    messagesPage.value = res?.data?.messagesPage;
    messages.value.unshift(...messagesPage.value.messages.reverse());

    nextTick(() => {
      if (!root.value) return;
      
      const newScrollHeight = root.value.scrollHeight;
      const heightDifference = newScrollHeight - scrollHeight;
      
      if (isAtBottom) {
        root.value.scrollTop = newScrollHeight;
      } else {
        root.value.scrollTop = scrollTop + heightDifference;
      }
    });
  })
};

const scrollToBottom = () => {
  if (!root.value) {
    return;
  }

  nextTick(() => {
    root.value!.scrollTop = root.value!.scrollHeight;
  });
};

const { isActive, pause, resume } = useIntersectionObserver(
  target,
  ([entry]) => {
    // TODO maybe fetch while entry.isIntersecting is true, when initally fetched messages are smaller then scroll window height
    if (entry.isIntersecting) {
      change_page();
    }
  },
)

const onLoad = async () => {
  await nextTick(() => {
    scrollToBottom();
  });
};

watch(() => props.chat._id, () => {
  page = 1;
  onLoad();
});

onMounted(() => {
  onLoad();
});

</script>

<template>

  <Card>
    <template #header>
      <div class="pl-6 pt-4">
        <h2 class="text-lg font-semibold">{{ props.chat.name }}</h2>
      </div>
    </template>

    <template #content>
      <div class="bg-darker rounded-md p-4">
        <div ref="root" id="chatBox" class="overflow-auto h-96 scroll-snap-y-container">
          <div ref="loadingMessagesDiv" class="flex h-5 w-full justify-center">No more messages</div>
            
          <div v-for="message in messages" :key="message._id" class="flex items-center justify-between p-2">
            <div class="flex w-full" :class="message.senderId === store.getId() ? 'justify-end' : 'justify-start'">
              <div class="max-w-xs lg:max-w-md px-4 py-2 rounded-lg shadow-md" 
                :class="message.senderId === store.getId() ? 'bg-emerald-600 text-white' : 'bg-gray-200 text-gray-800'">
                <p class="text-sm">{{ message.content }}</p>
                <p class="text-xs opacity-70 mt-1">
                  {{ new Date(message.createdAt).toLocaleString() }}
                </p>
              </div>
            </div>
          </div>
        </div>
      </div>
    </template>

    <template #footer> 
      <div class="pt-4">
        <Form class="flex">
          <FloatLabel variant="on" class="w-full">
            <Textarea v-model="newMessage" id="over_label" rows="2" class="w-full" @keydown.enter.prevent="sendMessage" />
          </FloatLabel>
          <div class="">
            <Button type="submit" severity="secondary" label="Send" class="ml-2 h-full" @click="sendMessage"/>
          </div>
        </Form>
      </div>
    </template>
  </Card>

</template>

<style scoped>

</style>
