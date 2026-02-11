<template>
  <UModal>
    <UButton
      label="Поиск товаров..."
      color="neutral"
      variant="subtle"
      icon="i-lucide-search"
    />

    <template #content>
      <UCommandPalette
        v-model:search-term="searchTerm"
        :groups="groups"
        placeholder="Поиск товаров..."
        class="h-80"
      />
    </template>
  </UModal>
</template>

<script setup lang="ts">
import { ref, computed } from "vue";
import type { CommandPaletteGroup } from "@nuxt/ui";

const searchTerm = ref("");

const groups = ref<CommandPaletteGroup[]>([
  {
    id: "users",
    label: "Users",
    items: [
      {
        label: "Benjamin Canac",
        suffix: "benjamincanac",
        avatar: {
          src: "https://github.com/benjamincanac.png",
        },
      },
      {
        label: "Romain Hamel",
        suffix: "romhml",
        avatar: {
          src: "https://github.com/romhml.png",
        },
      },
      {
        label: "Sébastien Chopin",
        suffix: "atinux",
        avatar: {
          src: "https://github.com/atinux.png",
        },
      },
      {
        label: "Hugo Richard",
        suffix: "HugoRCD",
        avatar: {
          src: "https://github.com/HugoRCD.png",
        },
      },
      {
        label: "Sandro Circi",
        suffix: "sandros94",
        avatar: {
          src: "https://github.com/sandros94.png",
        },
      },
      {
        label: "Daniel Roe",
        suffix: "danielroe",
        avatar: {
          src: "https://github.com/danielroe.png",
        },
      },
      {
        label: "Jakub Michálek",
        suffix: "J-Michalek",
        avatar: {
          src: "https://github.com/J-Michalek.png",
        },
      },
      {
        label: "Eugen Istoc",
        suffix: "genu",
        avatar: {
          src: "https://github.com/genu.png",
        },
      },
    ],
  },
]);

// const { data: users, status } = await useFetch(
//   "https://jsonplaceholder.typicode.com/users",
//   {
//     key: "command-palette-users",
//     params: { q: searchTerm },
//     transform: (data: { id: number; name: string; email: string }[]) => {
//       return (
//         data?.map((user) => ({
//           id: user.id,
//           label: user.name,
//           suffix: user.email,
//           avatar: { src: `https://i.pravatar.cc/120?img=${user.id}` },
//         })) || []
//       );
//     },
//     lazy: true,
//   },
// );
</script>

<style scoped></style>
