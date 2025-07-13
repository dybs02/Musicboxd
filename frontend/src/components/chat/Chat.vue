<script setup lang="ts">
import { useAuthStore } from "@/services/authStore";
import { CHAT_SUBSCRIPTION, SEND_MESSAGE } from "@/services/queries";
import { emptyMessage } from "@/types/message";
import { handleGqlError } from "@/utils/error";
import { gql } from "@apollo/client/core";
import { Form } from '@primevue/forms';
import { useMutation, useSubscription } from "@vue/apollo-composable";
import Button from 'primevue/button';
import FloatLabel from "primevue/floatlabel";
import Textarea from 'primevue/textarea';
import { ref, watch } from "vue";
import { useRoute, useRouter } from "vue-router";


const store = useAuthStore();
const route = useRoute();
const router = useRouter();

const props = defineProps<{
  chatId: string;
}>();


const newMessage = ref<string>('')
const messages = ref<any[]>([])
console.log("Chat component initialized with chatId:", props.chatId);
const { result, error: errorHandler } = useSubscription(
  CHAT_SUBSCRIPTION,
  {
    chatId: props.chatId,
  }
)

watch(
  result,
  data => {
    if (data.loading) {
      console.log("Loading messages...");
      return;
    }

    messages.value.push(data.messageAdded)
    console.log("New message received:", data.messageAdded)
  }
)

watch(errorHandler, (err) => {
  console.error("Subscription error:", err);
  console.log(props.chatId)
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
        chatId: props.chatId,
        content: newMessage.value,
      },
    }
  ));

  watch(addMessageError, (err) => {
    handleGqlError(router, err);
  });

  addMessageOnDone(async (result) => {
    newMessage.value = '';
    console.log("Message sent successfully:", result);
  });

  addMessage();
}


</script>

<template>
  <h1>Chat</h1>

  <div>
    {{ messages.length }} messages received
  </div>

  <div v-for="message in messages">
    <p>{{ message.senderId }}: {{ message.content }}</p>
  </div>

  <Form class="flex">
    <FloatLabel variant="on" class="w-full">
      <Textarea v-model="newMessage" id="over_label" rows="1" class="w-full" />
    </FloatLabel>
    <div class="">
      <Button type="submit" severity="secondary" label="Send" class="ml-2" @click="sendMessage"/>
    </div>
  </Form>
</template>

<style scoped>

</style>
