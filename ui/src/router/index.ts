import type { RouteRecordRaw } from "vue-router";
import { createRouter, createWebHistory } from "vue-router";
import createRouteGuard from "@/router/guard";
import DefaultLayout from "@/layouts/default.vue";
import ProjectLayout from "@/layouts/project.vue";

const routes: Array<RouteRecordRaw> = [
    {
        path: "/login",
        name: "login",
        component: () => import("@/pages/login.vue"),
    },
    {
        path: "/",
        name: "root",
        component: DefaultLayout,
        redirect: {
            name: "home",
        },
        meta: {
            layout: true,
        },
        children: [
            {
                path: "",
                name: "home",
                component: () => import("@/pages/home.vue"),
                meta: {
                    title: "首页",
                    icon: "icon-home",
                },
            },
            {
                path: "/projects",
                name: "projects",
                component: () => import("@/pages/projects.vue"),
                meta: {
                    title: "项目",
                    icon: "icon-apps",
                },
            },
            {
                path: "/settings",
                name: "settings",
                component: () => import("@/pages/settings.vue"),
                meta: {
                    title: "设置",
                    icon: "icon-settings",
                },
            },
        ],
    },
    {
        path: "/project/:id",
        name: "project-root",
        component: ProjectLayout,
        redirect: {
            name: "documen",
        },
        meta: {
            layout: tre,
        },
        children: [
            {
                path: "",
                name: "document",
                component: () => import("@/pages/project/document.vue"),
                meta: {
                    title: "文档列表",
                    icon: "icon-app",
                },
            },
        ],
    },
];

const router = createRouter({
    history: createWebHistory(import.meta.env.BASE_URL),
    routes,
    scrollBehavior() {
        return { top: 0 };
    },
});

createRouteGuard(router);

export default router;
