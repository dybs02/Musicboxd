<script setup lang="ts">
import { ref, watch } from "vue"
import { gql } from "@apollo/client/core";
import { useSubscription } from "@vue/apollo-composable"

const messages = ref<any[]>([])

const { result } = useSubscription(gql`
  subscription Msg {
    msg {
      timeStamp
    }
  }
`)

watch(
  result,
  data => {
    messages.value.push(data.msg)
    console.log("New message received:", data)
  }
)
</script>

<template>
  <h1>Chat</h1>
  <div>
    {{ messages.length }} messages received
  </div>

  <div>
    {{ messages }}
  </div>
</template>

<style scoped>

</style>
