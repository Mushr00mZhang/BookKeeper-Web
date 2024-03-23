import { createRouter, createWebHistory, RouteRecordRaw } from 'vue-router';
import OutlayList from '@/views/Outlay/List.vue';
import OutlayDetail from '@/views/Outlay/Detail.vue';
import OutlayCatList from '@/views/OutlayCat/List.vue';
import OutlayCatDetail from '@/views/OutlayCat/Detail.vue';
import Index from '@/views/Index/Index.vue';
const routes: RouteRecordRaw[] = [
  { path: '/', component: Index },
  {
    name: 'Outlay',
    path: '/outlay',
    children: [
      { name: 'OutlayList', path: '/outlay/list', component: OutlayList },
      { name: 'OutlayDetail', path: '/outlay/detail', component: OutlayDetail },
    ],
  },
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
