<template>
  <USelectMenu
    placeholder="Выбрать категорию"
    :items="items"
    class="w-52 cursor-pointer"
    trailing-icon="i-lucide-arrow-down"
    selected-icon="i-lucide-flame"
    :search-input="false"
    :ui="{
      trailingIcon:
        'group-data-[state=open]:rotate-180 transition-transform duration-300',
    }"
    clear
    size="xl"
  />
</template>

<script setup lang="ts">
import { ref } from "vue";

import type { SelectMenuItem } from "@nuxt/ui";

const items = ref<SelectMenuItem[]>([
  {
    label: "Backlog",
    icon: "i-lucide-circle-plus",
  },
  {
    label: "Todo",
    icon: "i-lucide-circle-plus",
  },
  {
    label: "In Progress",
    icon: "i-lucide-circle-plus",
  },
  {
    label: "Done",
    icon: "i-lucide-circle-plus",
  },
]);
</script>

<style scoped></style>
