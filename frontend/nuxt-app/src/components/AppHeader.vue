<template>
  <UHeader title="Products-service" to="/">
    <template #right>
      <app-products-search />
    </template>
  </UHeader>
</template>

<script setup lang="ts">
import AppProductsSearch from "./AppProductsSearch.vue";
</script>

<style scoped></style>
