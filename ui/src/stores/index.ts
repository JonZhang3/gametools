import { createPinia } from "pinia";
import useAppStore from "@/stores/app";
import useUserStore from "@/stores/user";

const store = createPinia();

export { useAppStore, useUserStore };

export default store;
