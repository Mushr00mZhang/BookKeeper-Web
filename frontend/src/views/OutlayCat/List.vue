<template>
  <!-- 列表页
  <div v-for="item in items">
    <div>
      {{ item.name }}
    </div>
    <ElButton type="primary" @click="create(item)">添加</ElButton>
    <ElButton type="primary" @click="update(item)">编辑</ElButton>
    <ElButton type="warning" @click="remove(item)">删除</ElButton>
  </div> -->
  <ElTree lazy :load="load" :props="treeProps" node-key="id">
    <template #default="{ data: item }: { node: Node, data: OutlayCat }">
      <label>{{ item.name }}</label>
      <ElButton type="success" @click="create(item)" :icon="Plus" size="small" />
      <ElButton type="primary" @click="update(item)" :icon="Edit" size="small" />
      <ElButton type="danger" @click="remove(item)" :icon="Delete" size="small" />
    </template>
  </ElTree>
</template>
<script setup lang="ts">
import { reactive } from 'vue';
import { onBeforeRouteUpdate, useRoute, useRouter } from 'vue-router';
import { OutlayCat, IListDto as OutlayCatListDto } from './model';
import { ElTree, ElButton } from 'element-plus';
import { Plus, Edit, Delete } from '@element-plus/icons-vue';
import 'element-plus/es/components/tree/style/css';
import 'element-plus/es/components/button/style/css';
import {
  LoadFunction,
  TreeData,
  TreeOptionProps,
} from 'element-plus/es/components/tree/src/tree.type.mjs';
import Node from 'element-plus/es/components/tree/src/model/node';
const route = useRoute();
const router = useRouter();
const dto = reactive<OutlayCatListDto>({ parentId: null });
const items = reactive<OutlayCat[]>([]);
const treeProps: TreeOptionProps = {
  label: 'name',
  isLeaf: (data, _node) => !(data as OutlayCat).hasChildren,
};
const load: LoadFunction = async (node: Node, resolve: (data: TreeData) => void) => {
  const parentId = node.level === 0 ? null : (node.data as OutlayCat).id;
  const items = await OutlayCat.list({ parentId });
  resolve(items);
};
const create = (item: OutlayCat) => {
  router.push({ path: 'detail', query: { mode: 'create', parentId: item.id } });
};
const update = (item: OutlayCat) => {
  router.push({ path: 'detail', query: { id: item.id, mode: 'update' } });
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
