import type { Router } from "vue-router";
import events from "@/stores/events";
import setupUserLoginGuard from "@/router/guard/userLogin";

function setupPageGuard(router: Router) {
    router.beforeEach(async (to) => {
        // emit route change
        events.route.setRouteEmitter(to);
    });
}

export default function createRouteGuard(router: Router) {
    setupPageGuard(router);
    setupUserLoginGuard(router);
}
