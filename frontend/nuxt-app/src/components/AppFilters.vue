<template>
  <UAccordion :items="accordeon" trailing-icon="i-heroicons-cog-6-tooth">
    <template #body="{ item }">
      <UCarousel
        wheel-gestures
        :prev="{ variant: 'solid' }"
        :next="{ variant: 'solid' }"
        drag-free
        :items="carouselItems"
        :arrows="false"
        :ui="{
          item: 'basis-auto ps-0 pr-3',
          prev: 'sm:start-8',
          next: 'sm:end-8',
          container: 'ms-0',
        }"
      >
        <div class="flex gap-3.75">
          <USelectMenu
            v-model="value"
            size="xl"
            :items="items"
            :search-input="false"
            class="w-48"
          />
        </div>
      </UCarousel>
    </template>
  </UAccordion>
</template>

<script setup lang="ts">
import { ref } from "vue";
import type { AccordionItem } from "@nuxt/ui";
const items = ref(["Backlog", "Todo", "In Progress", "Done"]);
const value = ref("Backlog");

const accordeon = ref<AccordionItem[]>([
  {
    label: "Фильтрация",
  },
]);

const carouselItems = [
  { id: 1, label: "One" },
  { id: 2, label: "One" },
  { id: 3, label: "One" },
  { id: 1, label: "One" },
  { id: 1, label: "One" },
];
</script>

<style scoped></style>
