<template>
  列表页
  <div v-for="item in items">
    <div>
      {{ item.name }}
    </div>
    <ElButton type="warning" @click="remove(item)">删除</ElButton>
  </div>
  <!-- <ElTree></ElTree> -->
  <ElButton type="primary" @click="create">添加</ElButton>
</template>
<script setup lang="ts">
import { reactive } from 'vue';
import { onBeforeRouteUpdate, useRoute, useRouter } from 'vue-router';
import { OutlayCat, IListDto as OutlayCatListDto } from './model';
import { ElTree, ElButton } from 'element-plus';
import 'element-plus/es/components/tree/style/css';
import 'element-plus/es/components/button/style/css';
const route = useRoute();
const router = useRouter();
const dto = reactive<OutlayCatListDto>({ parentId: null });
const items = reactive<OutlayCat[]>([]);
const create = () => {
  router.push('detail');
};
const remove = async (item: OutlayCat) => {
  await item.delete();
  const i = items.indexOf(item);
  items.splice(i, 1);
};
const init = async () => {
  items.splice(0);
  const list = await OutlayCat.list(dto);
  items.push(...list);
};
init();
// onBeforeRouteUpdate(async (to, from) => {
//   console.log(to, from);
//   await init();
// });
</script>
