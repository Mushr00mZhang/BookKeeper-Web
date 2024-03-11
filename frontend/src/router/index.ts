import { createRouter, createWebHistory, RouteRecordRaw } from 'vue-router';
import OutlayCatList from '@/views/OutlayCat/List.vue';
import OutlayCatDetail from '@/views/OutlayCat/Detail.vue';
import Index from '@/views/Index/Index.vue';
const routes: RouteRecordRaw[] = [
  { path: '/', component: Index },
  {
    name: 'OutlayCat',
    path: '/outlaycat',
    children: [
      { name: 'OutlayCatList', path: '/outlaycat/list', component: OutlayCatList },
      { name: 'OutlayCatDetail', path: '/outlaycat/detail', component: OutlayCatDetail },
    ],
  },
];

export const router = createRouter({
  history: createWebHistory(),
  routes,
});
