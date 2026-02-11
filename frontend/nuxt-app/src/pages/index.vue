<template>
  <div>
    <UPage class="my-5">
      <app-categories />
      <USeparator class="my-5" />
      <app-cards />
    </UPage>
  </div>
</template>

<script setup lang="ts">
import AppCards from "../components/AppCards.vue";
import AppCategories from "../components/AppCategories.vue";
</script>

<style scoped></style>
