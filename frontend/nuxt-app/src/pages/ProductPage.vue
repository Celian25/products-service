<template>
  <UBreadcrumb :items="breadcrumb" />
  <UCard variant="subtle">
    <template #header>
      <Placeholder class="h-8" />
      Изображение
    </template>

    <Placeholder class="h-32" />
    Описание
    <template #footer>
      <Placeholder class="h-8" />
      Цена
    </template>
  </UCard>
</template>

<script setup lang="ts">
import { ref } from "vue";
import type { BreadcrumbItem } from "@nuxt/ui";

const breadcrumb = ref<BreadcrumbItem[]>([
  {
    label: "Категории",
    icon: "i-lucide-book-open",
    to: "/",
  },
  {
    label: "Название категории",
    icon: "i-lucide-box",
    to: "/products-slug",
  },
  {
    label: "Название продукта",
    icon: "i-lucide-link",
    to: "/products-slug/product-slug",
  },
]);
</script>

<style scoped></style>
