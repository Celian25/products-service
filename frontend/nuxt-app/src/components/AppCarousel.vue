<template>
  <UCarousel
    v-slot="{ item }"
    :autoplay="{ delay: 2000 }"
    wheel-gestures
    :prev="{ variant: 'solid' }"
    :next="{ variant: 'solid' }"
    :items="items"
    drag-free
    :arrows="false"
    :ui="{
      item: 'basis-1/3 ps-0',
      prev: 'sm:start-8',
      next: 'sm:end-8',
      container: 'ms-0',
    }"
  >
    <img :src="item" width="320" height="320" />
  </UCarousel>
</template>

<script setup lang="ts">
const items = [
  "src/assets/imgs/img-placeholder.png",
  "src/assets/imgs/img-placeholder.png",
  "src/assets/imgs/img-placeholder.png",
  "src/assets/imgs/img-placeholder.png",
  "src/assets/imgs/img-placeholder.png",
  "src/assets/imgs/img-placeholder.png",
];
</script>

<style scoped></style>
