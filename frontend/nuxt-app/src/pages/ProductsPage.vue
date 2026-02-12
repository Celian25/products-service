<template>
  <div>
    <UPage class="my-5">
      <UBreadcrumb :items="breadcrumb" />
      <UInputMenu
        placeholder="Поиск по продуктам..."
        :items="items"
        :arrow="false"
        :open-on-click="true"
        :multiple="false"
        class="w-full"
        clear
      >
        <template #empty>Ничего не найдено &#128532</template>
      </UInputMenu>

      <USeparator class="my-5" />
      <app-cards title="Продукты" :items="cardProps" />
    </UPage>
  </div>
</template>

<script setup lang="ts">
import { ref } from 'vue';
import { PageCard } from "../types/types";
import type { BreadcrumbItem } from '@nuxt/ui'

const breadcrumb = ref<BreadcrumbItem[]>([
  {
    label: 'Категории',
    icon: 'i-lucide-book-open',
    to: '/'
  },
  {
    label: 'Название категории',
    icon: 'i-lucide-box',
    to: '/products-slug'
  },
])

const items = ref([
  "Продукт1",
  "Продукт2",
  "Продукт3",
  "Продукт4",
  "Продукт5",
  "Продукт6",
  "Продукт7",
  "Продукт8",
  "Продукт9",
]);

const cardProps: PageCard[] = [
  {
    title: "Продукт 1",
    description:
      "Nuxt UI integrates with latest Tailwind CSS, bringing significant improvements.",
    to: "/products-slug/product-slug",
    icon: "i-heroicons-home",
    img: "src/assets/imgs/img-placeholder.png",
    price: 2222,
  },
  {
    title: "Продукт 2",
    description: "Go to dashboard",
    to: "/products-slug/product-slug",
    icon: "i-heroicons-home",
    img: "src/assets/imgs/img-placeholder.png",
    price: 2223,
  },
  {
    title: "Продукт 2",
    description: "Go to dashboard",
    to: "/products-slug/product-slug",
    icon: "i-heroicons-home",
    img: "src/assets/imgs/img-placeholder.png",
    price: 2223,
  },
  {
    title: "Продукт 2",
    description: "Go to dashboard",
    to: "/products-slug/product-slug",
    icon: "i-heroicons-home",
    img: "src/assets/imgs/img-placeholder.png",
    price: 2223,
  },
];

</script>

<style scoped></style>
