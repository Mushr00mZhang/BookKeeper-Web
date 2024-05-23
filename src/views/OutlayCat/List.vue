<template>
  <ElTree
    :load="load"
    :props="treeProps"
    :default-expanded-keys="[...expandedIds]"
    node-key="id"
    lazy
    ref="treeRef"
    @node-expand="expand"
    @node-collapse="collapse"
  >
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
import { onActivated, reactive, ref } from 'vue';
import { useRouter } from 'vue-router';
import { ElTree, ElButton, ElMessage, ElMessageBox, TreeInstance } from 'element-plus';
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
const router = useRouter();
const treeProps: TreeOptionProps = {
  label: 'name',
  isLeaf: (data, _node) => !(data as OutlayCat).hasChildren,
};
const treeRef = ref<TreeInstance>();
const expandedIds = reactive(new Set([ROOT_ID]));
const load: LoadFunction = async (node: Node, resolve: (data: TreeData) => void) => {
  switch (node.level) {
    case 0:
      resolve([ROOT]);
      break;
    case 1:
      resolve([]);
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
      const tree = treeRef.value;
      if (tree) {
        const node = tree.getNode(item?.id || '');
        tree.remove(node);
      }
    } else {
      ElMessage.error('删除失败');
    }
  }
};
const expand = (item: OutlayCat, _node: Node) => {
  if (item.id) {
    expandedIds.add(item.id);
  }
};
const collapse = (item: OutlayCat, _node: Node) => {
  if (item.id && item.id !== ROOT_ID) {
    expandedIds.delete(item.id);
  }
};
const init = async () => {
  const node = treeRef?.value?.getNode(ROOT_ID);
  if (node) {
    node.loaded = false;
    node.loading = true;
    const items = await OutlayCat.list({ parentId: ROOT_ID });
    node.childNodes.splice(0);
    node.doCreateChildren(items);
    node.loaded = true;
    node.loading = false;
  }
};
onActivated(init);
</script>
