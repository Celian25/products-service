import { createRouter, createWebHistory } from "vue-router";

export const router = createRouter({
  routes: [
    { path: "/", component: () => import("./pages/index.vue") },
    { path: "/not-found", component: () => import("./pages/not-found.vue") },
  ],
  history: createWebHistory(),
});

router.beforeEach((to, from, next) => {
  console.log(to, from);
  next();
});
