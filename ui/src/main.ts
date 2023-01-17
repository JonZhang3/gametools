import { createApp } from "vue";
import ArcoVue from "@arco-design/web-vue";
import ArcoVueIcon from "@arco-design/web-vue/es/icon";

import App from "@/App.vue";
import router from "@/router";
import store from "@/stores";

import "@/styles/index.less";

const app = createApp(App);
app.use(store);
app.use(router);
app.use(ArcoVue);
app.use(ArcoVueIcon);
app.mount("#app");
