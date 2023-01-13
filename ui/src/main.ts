import { createApp } from "vue";
import ArcoVue from "@arco-design/web-vue";
import ArcoVueIcon from "@arco-design/web-vue/es/icon";

import App from "@/App.vue";
import router from "@/router";
import store from "@/stores";

import "@/styles/index.less";

// eslint-disable-next-line prettier/prettier
createApp(App)
    .use(store)
    .use(router)
    .use(ArcoVue)
    .use(ArcoVueIcon)
    .mount("#app");
