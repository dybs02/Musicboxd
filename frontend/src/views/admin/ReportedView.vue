<script setup lang="ts">
import { GET_REPORTED_COMMENTS } from '@/services/queries_admin';
import { useQuery } from '@vue/apollo-composable';
import Button from 'primevue/button';
import { watch } from 'vue';


const fetch_reported_comments = async () => {
  const { loading, error, result } = useQuery(
    GET_REPORTED_COMMENTS,
    {
      number: 10,
    }
  );


  watch(result, (res) => {
    console.log(res);
  });
};






</script>

<template>

<Button @click="fetch_reported_comments">
  Fetch Reported Comments
</Button>

</template>

<style scoped>

</style>
