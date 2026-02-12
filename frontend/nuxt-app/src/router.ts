import { createRouter, createWebHistory } from "vue-router";

export const router = createRouter({
  routes: [
    { path: "/", component: () => import("./pages/IndexPage.vue") },
    { path: "/not-found", component: () => import("./pages/NotFoundPage.vue") },
    {
      path: "/products-slug",
      component: () => import("./pages/ProductsPage.vue"),
    },
    {
      path: "/products-slug/product-slug",
      component: () => import("./pages/ProductPage.vue"),
    },
  ],
  history: createWebHistory(),
});

router.beforeEach((to, from, next) => {
  console.log(to, from);
  next();
});
