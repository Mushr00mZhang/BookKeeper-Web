<template>
  <ElTree :load="load" :props="treeProps" :default-expanded-keys="[ROOT_ID]" node-key="id" lazy>
    <template #default="{ node, data: item }: { node: Node, data: OutlayCat }">
      <label>{{ item.name }}</label>
      <ElButton type="success" @click="create(item)" :icon="Plus" circle />
      <template v-if="node.level > 1">
        <ElButton type="primary" @click="update(item)" :icon="Edit" circle />
        <ElButton type="danger" @click="remove(item)" :icon="Delete" circle />
      </template>
    </template>
  </ElTree>
</template>
<script setup lang="ts">
import { reactive } from 'vue';
import { useRoute, useRouter } from 'vue-router';
import { ElTree, ElButton, ElMessage, ElMessageBox } from 'element-plus';
import { Plus, Edit, Delete } from '@element-plus/icons-vue';
import {
  LoadFunction,
  TreeData,
  TreeOptionProps,
} from 'element-plus/es/components/tree/src/tree.type.mjs';
import Node from 'element-plus/es/components/tree/src/model/node';
import 'element-plus/es/components/tree/style/css';
import 'element-plus/es/components/button/style/css';
import { OutlayCat, ROOT, ROOT_ID } from './model';
// const route = useRoute();
const router = useRouter();
const treeProps: TreeOptionProps = {
  label: 'name',
  isLeaf: (data, _node) => !(data as OutlayCat).hasChildren,
};
const load: LoadFunction = async (node: Node, resolve: (data: TreeData) => void) => {
  switch (node.level) {
    case 0:
      resolve([ROOT]);
      break;
    default:
      const parentId = (node.data as OutlayCat).id;
      const items = await OutlayCat.list({ parentId });
      resolve(items);
      break;
  }
};
const create = (item: OutlayCat) => {
  router.push({ path: 'detail', query: { mode: 'create', parentId: item.id } });
};
const update = (item: OutlayCat) => {
  router.push({ path: 'detail', query: { id: item.id, mode: 'update' } });
};
const remove = async (item: OutlayCat) => {
  const confirmed = await ElMessageBox.confirm(`确认删除${item.name}？`, '', {
    confirmButtonText: '确认',
    cancelButtonText: '取消',
    type: 'warning',
  })
    .then(() => true)
    .catch(() => false);
  if (confirmed) {
    const res = await item.delete();
    if (res) {
      ElMessage.success('删除成功');
    } else {
      ElMessage.error('删除失败');
    }
  }
};
const init = async () => {};
init();
</script>
