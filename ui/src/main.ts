import { createApp } from "vue";
import ArcoVue from "@arco-design/web-vue";
// 额外引入图标库
import ArcoVueIcon from "@arco-design/web-vue/es/icon";

import App from "@/App.vue";
import router from "@/router";
import store from "@/stores";

import "@/styles/index.less";

createApp(App)
    .use(store)
    .use(router)
    .use(ArcoVue)
    .use(ArcoVueIcon)
    .mount("#app");
