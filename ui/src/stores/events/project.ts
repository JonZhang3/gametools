import emitter from "@/stores/events/emitter";

const CREATE_OK = Symbol("CREATE_PROJECT_OK");
const EDIT_OK = Symbol("EDIT_PROJECT_OK");

export function setCreateOk() {
    emitter.emit(CREATE_OK);
}

export function listenCreateOk(handler: () => void) {
    emitter.on(CREATE_OK, handler);
}

export function setEditOk() {
    emitter.emit(EDIT_OK);
}

export function listenEditOk(handler: () => void) {
    emitter.on(EDIT_OK, handler);
}
