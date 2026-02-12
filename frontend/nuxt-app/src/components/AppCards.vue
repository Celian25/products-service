<template>
  <div class="">{{ title }}</div>
  <UBlogPosts>
    <UPageCard
      v-for="{ title, description, icon, to, img } in items"
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
      <template #footer>
        <div>2222</div>
      </template>
    </UPageCard>
  </UBlogPosts>
</template>

<script setup lang="ts">
import { ref } from "vue";
import { PageCard } from "../types/types";
// import { User, UserRole } from "@/types/user";
// import type { User, UserRole } from "@/types/user";

const props = defineProps<{
  title: string;
  items: PageCard[];
}>();
</script>

<style scoped></style>
