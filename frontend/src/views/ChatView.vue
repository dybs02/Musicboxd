<script setup lang="ts">
import Chat from '@/components/chat/Chat.vue';
import { useAuthStore } from '@/services/authStore';
import { CHAT_BY_PARTICIPANT } from '@/services/queries';
import { emptyChat, type ChatType } from '@/types/message';
import { handleGqlError } from '@/utils/error';
import { useQuery } from '@vue/apollo-composable';
import Button from 'primevue/button';
import { ref } from 'vue';
import { useRoute, useRouter } from 'vue-router';


const store = useAuthStore();
const route = useRoute();
const router = useRouter();


const participantID = "685699fca516abe06da5d798"
const chat = ref<ChatType>(emptyChat);
let chatLoading = ref(true);


const openChat = () => {
  // This is a placeholder for the chat opening logic
  const { loading, onError, onResult } = useQuery(
    CHAT_BY_PARTICIPANT,
    {
      participantId: participantID,
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
    console.log("Open chat with participant ID:", participantID);
    console.log("Chat data:", chat.value);
  })

};


</script>

<template>
  <h1>About</h1>

  <div v-if="!chatLoading">
    <p>Chat with participant ID: {{ participantID }}</p> 
    <Chat
      :chat-id="chat._id"
    />
  </div>


  <Button label="Open chat" @click="openChat" />
</template>

<style scoped>

</style>
