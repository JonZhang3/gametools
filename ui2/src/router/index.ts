import type { RouteRecordRaw } from "vue-router";
import { createRouter, createWebHistory } from "vue-router";
import DefaultLayout from "@/layouts/default.vue";

const routes: Array<RouteRecordRaw> = [
    {
        path: "/login",
        name: "login",
        component: () => import("@/pages/login.vue")
    },
    {
        path: "/",
        name: "root",
        component: DefaultLayout,
        children: [
            {
                path: "",
                name: "home",
                component: () => import("@/pages/login.vue")
            }
        ]
    }
];

const router = createRouter({
    history: createWebHistory(import.meta.env.BASE_URL),
    routes,
    scrollBehavior() {
        return { top: 0 };
    }
});

export default router;
