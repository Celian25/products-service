<template>
  <Suspense>
    <UApp>
      <app-header />

      <UMain>
        <UContainer>
          <RouterView />
        </UContainer>
      </UMain>

      <UFooter>
        <template #left>
          <p class="text-sm text-muted">
            Built with Nuxt UI • © {{ new Date().getFullYear() }}
          </p>
        </template>

        <template #right>
          <UButton
            to="https://github.com/nuxt-ui-templates/starter-vue"
            target="_blank"
            icon="simple-icons:github"
            aria-label="GitHub"
            color="neutral"
            variant="ghost"
          />
        </template>
      </UFooter>
    </UApp>
  </Suspense>
</template>

<script setup lang="ts">
import AppHeader from "./components/AppHeader.vue";
</script>

<style scoped></style>
