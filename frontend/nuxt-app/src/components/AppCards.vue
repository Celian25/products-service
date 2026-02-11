<template>
  <UBlogPosts>
    <UPageCard
      v-for="{ title, description, icon, to, img } in cardProps"
      :title="title"
      :description="description"
      :icon="icon"
      :to="to"
      orientation="vertical"
      spotlight
      spotlight-color="primary"
    >
      <template #header>
        <img :src="img" alt="Tailwind CSS" class="w-full" />
      </template>
    </UPageCard>
  </UBlogPosts>
</template>

<script setup lang="ts">
import { ref } from "vue";
// import { User, UserRole } from "@/types/user";
// import type { User, UserRole } from "@/types/user";

const cardProps = [
  {
    title: "Dashboard1",
    description:
      "Nuxt UI integrates with latest Tailwind CSS, bringing significant improvements.",
    to: "/not-found",
    icon: "i-heroicons-home",
    img: "src/assets/imgs/img-placeholder.png",
  },
  {
    title: "Dashboard2",
    description: "Go to dashboard",
    to: "/not-found",
    icon: "i-heroicons-home",
    img: "src/assets/imgs/img-placeholder.png",
  },
];
</script>

<style scoped></style>
