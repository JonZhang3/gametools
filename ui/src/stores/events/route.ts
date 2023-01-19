import type { Handler } from "mitt";
import type { RouteLocationNormalized } from "vue-router";
import emitter from "@/stores/events/emitter";

const ROUTE_CHANGE = Symbol("ROUTE_CHANGE");

let latestRoute: RouteLocationNormalized;

export function setRouteEmitter(to: RouteLocationNormalized) {
    emitter.emit(ROUTE_CHANGE, to);
    latestRoute = to;
}

export function listenerRouteChange(handler: (route: RouteLocationNormalized) => void, immediate = true) {
    emitter.on(ROUTE_CHANGE, handler as Handler);
    if (immediate && latestRoute) {
        handler(latestRoute);
    }
}

export function removeRouteListener() {
    emitter.off(ROUTE_CHANGE);
}
