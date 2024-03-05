import { createRouter, createWebHistory, RouteRecordRaw } from 'vue-router';
import OutlayCatList from '@/views/OutlayCat/List.vue';

const routes: RouteRecordRaw[] = [
  // { path: '/', component: Home },
  { path: '/outlaycat', children: [{ path: '/list', component: OutlayCatList }] },
];

export const router = createRouter({
  history: createWebHistory(),
  routes,
});
