<template>
  <UHeader
    title="Products-service"
    mode="slideover"
    @update:open="isOpen = !isOpen"
    to="/"
    :toggle="{
      icon: isOpen ? 'i-heroicons-x-mark' : 'i-ph-sliders-horizontal',
    }"
  >
    <template #body>
      <h2 class="text-2xl font-bold mb-2">Поиск продуктов</h2>
      <UForm class="space-y-4">
        <UInputMenu
        placeholder="Продукт..."
        :items="items"
        :arrow="false"
        :open-on-click="true"
        :multiple="false"
        class="w-full"
        clear
      >
        <template #empty>Ничего не найдено &#128532</template>
      </UInputMenu>
      <div class="text-lg font-bold">Фильтры</div>
      </UForm>

    </template>
  </UHeader>
</template>

<script setup lang="ts">
import { ref } from "vue";
import AppProductsSearch from '"./AppProductsSearch.vue"';

let isOpen = ref(false);

const items = ref([
  "Категория 1",
  "Категория 2",
  "Категория 3",
  "Категория 4",
  "Категория 5",
]);
</script>

<style scoped></style>
