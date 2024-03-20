import { createApp } from 'vue';
import '@/styles/common.scss';
import App from './App.vue';
import 'element-plus/es/components/message/style/css';
import 'element-plus/es/components/tooltip/style/css';
import { router } from './router';

const app = createApp(App);
app.use(router);
app.mount('#app');
