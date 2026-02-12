<template>
  <div class=""></div>
  <UAccordion :items="accordeon">
    <template #body="{ item }">
      <USelect v-model="value" size="xl" :items="items" class="w-48" />
      <USelect v-model="value" size="xl" :items="items" class="w-48" />
      <USelect v-model="value" size="xl" :items="items" class="w-48" />
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
    content:
      "Yes! Nuxt UI is completely free and open source under the MIT license. All 125+ components are available to everyone.",
  },
]);
</script>

<style scoped></style>
