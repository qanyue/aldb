import {createApp} from "vue";
import App from "./App.vue";
import router from "./router";

import "~/styles/index.scss";
import "element-plus/dist/index.css";
import "element-plus/theme-chalk/src/message.scss";
import {createPinia} from "pinia";

const pinia = createPinia();

const app = createApp(App);
// app.use(ElementPlus);
app.use(pinia)
app.use(router);
app.mount("#app");
