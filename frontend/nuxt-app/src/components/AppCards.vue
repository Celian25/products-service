<template>
  <UBlogPosts>
    <UPageCard
      v-for="{ title, description, to, img } in items"
      :title="title"
      :description="description"
      :to="to"
      orientation="vertical"
      spotlight
      spotlight-color="primary"
    >
      <template #header>
        <img :src="img" alt="Tailwind CSS" class="w-full" />
      </template>
      <template #footer>
        <div>{{ formatCurrency(2222) }}</div>
      </template>
    </UPageCard>
  </UBlogPosts>
</template>

<script setup lang="ts">
import { ref } from "vue";
import { PageCard } from "@/types/types";
import { formatCurrency } from "@/utils/currency";
// import { User, UserRole } from "@/types/user";
// import type { User, UserRole } from "@/types/user";

const props = defineProps<{
  items: PageCard[];
}>();
</script>

<style scoped></style>
