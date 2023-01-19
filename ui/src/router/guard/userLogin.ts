import type { Router, LocationQueryRaw } from "vue-router";
import { isLogin } from "@/auth";

export default function setupUserLoginGuard(router: Router) {
    router.beforeEach((to, from, next) => {
        if (isLogin()) {
            next();
        } else {
            if (to.name === "login") {
                next();
                return;
            }
            next({
                name: "login",
                query: {
                    redirect: to.name,
                    ...to.query,
                } as LocationQueryRaw,
            });
        }
    });
}
