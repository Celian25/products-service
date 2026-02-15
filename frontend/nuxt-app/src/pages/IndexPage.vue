<template>
  <div>
    <UPage>
      <UBreadcrumb :items="breadcrumb" />
      <UInputMenu
        placeholder="Поиск по категории..."
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
      <app-cards :items="cardProps" />
    </UPage>
  </div>
</template>

<script setup lang="ts">
import { ref } from "vue";
import type { BreadcrumbItem } from '@nuxt/ui'
import AppCards from "@/components/AppCards.vue";
import { PageCard } from "@/types/types";
// import AppCategories from "../components/AppCategories.vue";

const breadcrumb = ref<BreadcrumbItem[]>([
  {
    label: "Категории",
    icon: "i-heroicons-squares-2x2",
    to: "/",
  },
]);

const items = ref([
  "Категория 1",
  "Категория 2",
  "Категория 3",
  "Категория 4",
  "Категория 5",
  "Категория 6",
  "Категория 7",
  "Категория 8",
  "Категория 9",
  "Категория 10",
]);

const cardProps: PageCard[] = [
  {
    title: "Категория 1",
    description:
      "Nuxt UI integrates with latest Tailwind CSS, bringing significant improvements.",
    to: "/products-slug",
    icon: "i-heroicons-home",
    img: "src/assets/imgs/img-placeholder.png",
    price: 2222,
  },
  {
    title: "Категория 2",
    description: "Go to dashboard",
    to: "/products-slug",
    icon: "i-heroicons-home",
    img: "src/assets/imgs/img-placeholder.png",
    price: 2223,
  },
  {
    title: "Категория 2",
    description: "Go to dashboard",
    to: "/products-slug",
    icon: "i-heroicons-home",
    img: "src/assets/imgs/img-placeholder.png",
    price: 2223,
  },
  {
    title: "Категория 2",
    description: "Go to dashboard",
    to: "/products-slug",
    icon: "i-heroicons-home",
    img: "src/assets/imgs/img-placeholder.png",
    price: 2223,
  },
];

</script>

<style scoped></style>
