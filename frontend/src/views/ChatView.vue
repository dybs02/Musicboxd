<script setup lang="ts">
import Chat from '@/components/chat/Chat.vue';
import { useAuthStore } from '@/services/authStore';
import { CHAT_BY_PARTICIPANT } from '@/services/queries';
import { emptyChat, type ChatType } from '@/types/message';
import { handleGqlError } from '@/utils/error';
import { useQuery } from '@vue/apollo-composable';
import Button from 'primevue/button';
import { ref, watch } from 'vue';
import { useRoute, useRouter } from 'vue-router';


const store = useAuthStore();
const route = useRoute();
const router = useRouter();


const chat = ref<ChatType>(emptyChat);
let chatLoading = ref(true);


const openChat = () => {
  if (!route.params.userId) {
    // TODO navigate to error page
    console.error("No user ID provided in route parameters.");
    return;
  }

  const { loading, onError, onResult } = useQuery(
    CHAT_BY_PARTICIPANT,
    {
      participantId: route.params.userId,
    }
  );

  
  onError((err) => {
    handleGqlError(router, err);
  });
  
  onResult((res: any) => {
    if (res.loading) {
      return;
    }
    
    chatLoading.value = false;
    chat.value = res?.data?.chat;
  })

};


watch(() => route.params, openChat, { immediate: true })


</script>

<template>

  <div v-if="!chatLoading">
    <Chat
      :chat="chat"
    />
  </div>

</template>

<style scoped>

</style>
