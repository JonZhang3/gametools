import ArcoVue from "@arco-design/web-vue";
import "@arco-design/web-vue/dist/arco.css";
import { defineNuxtPlugin } from "#imports";

export default defineNuxtPlugin(nuxtApp => {
    nuxtApp.vueApp.use(ArcoVue);
});