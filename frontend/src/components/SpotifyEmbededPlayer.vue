<script setup lang="ts">

const props = defineProps({
  trackId: {
    type: String,
    required: true
  }
});

</script>

<template>
  <iframe
    data-testid="embed-iframe"
    style="border-radius:12px"
    :src="`https://open.spotify.com/embed/track/${props.trackId}`"
    height="152"
    frameBorder="0"
    allowfullscreen
    allow="autoplay; clipboard-write; encrypted-media; fullscreen; picture-in-picture"
    loading="lazy"
  ></iframe>

</template>

<style scoped>

</style>
