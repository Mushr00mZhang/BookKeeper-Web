import { createRouter, createWebHistory, RouteRecordRaw } from 'vue-router';
import OutlayCatList from '@/views/OutlayCat/List.vue';
import Index from '@/views/Index/Index.vue';
const routes: RouteRecordRaw[] = [
  { path: '/', component: Index },
  { path: '/outlaycat', children: [{ path: '/outlaycat/list', component: OutlayCatList }] },
];

export const router = createRouter({
  history: createWebHistory(),
  routes,
});
