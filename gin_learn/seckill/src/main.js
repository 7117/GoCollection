import {createApp} from 'vue'
import App from './App.vue'

import ant from 'ant-design-vue';
import 'ant-design-vue/dist/antd.css';
import router from "./router";
import axios from "./axios";


const app = createApp(App);
app.use(router)
app.use(ant)
app.mount('#app')


app.config.globalProperties.$axios = axios
app.config.globalProperties.$router = router



