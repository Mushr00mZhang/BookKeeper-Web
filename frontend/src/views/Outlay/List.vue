<template>
  <ElButton type="success" @click="create" :icon="Plus" circle />
  <div v-for="item in items">
    <label>{{ item.name }}</label>
    <ElButton type="primary" @click="update(item)" :icon="Edit" circle />
    <ElButton type="danger" @click="remove(item)" :icon="Delete" circle />
  </div>
</template>
<script setup lang="ts">
import { reactive } from 'vue';
import { useRouter } from 'vue-router';
import { ElButton, ElMessage, ElMessageBox } from 'element-plus';
import { Plus, Edit, Delete } from '@element-plus/icons-vue';
import 'element-plus/es/components/button/style/css';
import { Outlay } from './model';
const router = useRouter();
const items = reactive<Outlay[]>([]);
const create = () => {
  router.push({ path: 'detail', query: { mode: 'create' } });
};
const update = (item: Outlay) => {
  router.push({ path: 'detail', query: { id: item.id, mode: 'update' } });
};
const remove = async (item: Outlay) => {
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
      const index = items.indexOf(item);
      items.splice(index, 1);
    } else {
      ElMessage.error('删除失败');
    }
  }
};
const init = async () => {
  const res = await Outlay.list({});
  items.push(...res);
};
init();
</script>
