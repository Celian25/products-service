<template>
  <UPage>
    <UBreadcrumb class="mb-4" :items="breadcrumb" />
    <UCard variant="subtle">
      <template #header>
        <Placeholder class="h-8" />
        Изображение
      </template>

      <Placeholder class="h-32" />
      Описание
      <template #footer>
        <Placeholder class="h-8" />
        Цена: {{ formatCurrency(2999) }}
      </template>
    </UCard>
  </UPage>
</template>

<script setup lang="ts">
import { ref } from "vue";
import type { BreadcrumbItem } from "@nuxt/ui";
import { formatCurrency } from "@/utils/currency";

const breadcrumb = ref<BreadcrumbItem[]>([
  {
    label: "Категории",
    icon: "i-heroicons-squares-2x2",
    to: "/",
  },
  {
    label: "Название категории",
    icon: "i-heroicons-folder",
    to: "/products-slug",
  },
  {
    label: "Название продукта",
    icon: "i-heroicons-tag",
    to: "/products-slug/product-slug",
  },
]);
</script>

<style scoped></style>
