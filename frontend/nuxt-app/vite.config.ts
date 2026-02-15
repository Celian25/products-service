import { fileURLToPath, URL } from "node:url";
import { defineConfig } from "vite";
import vue from "@vitejs/plugin-vue";
import ui from "@nuxt/ui/vite";

// https://vitejs.dev/config/
export default defineConfig({
  plugins: [
    vue(),
    ui({
      ui: {
        page: {
          slots: {
            root: "my-3",
          },
        },
        separator: {
          slots: {
            root: "my-5",
          },
        },
        breadcrumb: {
          slots: {
            root: "mb-3",
          },
        },
        blogPosts: {
          base: "lg:gap-y-8",
        },
        colors: {
          primary: "cyan",
          neutral: "zinc",
        },
        header: {
          slots: {
            toggle: "lg:block",
            content: "lg:block",
            overlay: "lg:block",
          },
        },
      },
    }),
  ],
  resolve: {
    alias: {
      "@": fileURLToPath(new URL("./src", import.meta.url)),
    },
  },
});
