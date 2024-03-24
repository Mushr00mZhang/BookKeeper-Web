<template>
  <div>
    <ElForm>
      <ElFormItem label="类型">
        <ElTreeSelect
          v-model="filter.catId"
          :load="load"
          :props="treeProps"
          :default-expanded-keys="[...expandedIds]"
          node-key="id"
          lazy
          check-strictly
          filterable
          clearable
          ref="treeRef"
          @node-expand="expand"
          @node-collapse="collapse"
        />
      </ElFormItem>
      <ElFormItem label="金额">
        <ElInputNumber v-model="filter.lowMoney" :min="0" :max="filter.topMoney ?? Infinity" />
        <span>-</span>
        <ElInputNumber v-model="filter.topMoney" :min="filter.lowMoney ?? 0" />
      </ElFormItem>
      <ElFormItem label="时间">
        <ElDatePicker
          v-model="filter.sTime"
          type="date"
          format="YYYY-MM-DD"
          value-format="YYYY-MM-DDT00:00:00+08:00"
        />
        <span>-</span>
        <ElDatePicker
          v-model="filter.eTime"
          type="date"
          format="YYYY-MM-DD"
          value-format="YYYY-MM-DDT00:00:00+08:00"
        />
      </ElFormItem>
    </ElForm>
    <ElButton type="success" @click="create" :icon="Plus" circle />
    <div v-for="item in pagedList.items">
      <label>{{ item.name }}</label>
      <ElButton type="primary" @click="update(item)" :icon="Edit" circle />
      <ElButton type="danger" @click="remove(item)" :icon="Delete" circle />
    </div>
    <ElPagination
      :total="pagedList.total"
      v-model:page-size="filter.size"
      :current-page="filter.index + 1"
      @update:current-page="(val) => (filter.index = val - 1)"
      layout="prev, pager, next"
    />
  </div>
</template>
<script setup lang="ts">
import { reactive, onActivated, ref, watchEffect } from 'vue';
import { useRouter } from 'vue-router';
import {
  ElForm,
  ElFormItem,
  ElPagination,
  ElInputNumber,
  ElButton,
  ElMessage,
  ElMessageBox,
  ElTreeSelect,
  ElDatePicker,
  TreeInstance,
} from 'element-plus';
import { Plus, Edit, Delete } from '@element-plus/icons-vue';
import {
  LoadFunction,
  TreeData,
  TreeOptionProps,
} from 'element-plus/es/components/tree/src/tree.type.mjs';
import Node from 'element-plus/es/components/tree/src/model/node';
import 'element-plus/es/components/tree/style/css';
import 'element-plus/es/components/button/style/css';
import { IDto as PagedListDto } from '@/utils/PagedList';
import { Outlay, IPagedListDto as OutlayPagedListDto } from './model';
import { OutlayCat, ROOT, ROOT_ID } from '@/views/OutlayCat/model';
const router = useRouter();
const treeProps: TreeOptionProps = {
  label: 'name',
  isLeaf: (data, _node) => !(data as OutlayCat)?.hasChildren,
  disabled: (_data, node) => node?.level <= 1,
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
const filter = reactive<OutlayPagedListDto>({
  catId: undefined,
  lowMoney: undefined,
  topMoney: undefined,
  sTime: undefined,
  eTime: undefined,
  userId: undefined,
  index: 0,
  size: 20,
});
const pagedList = reactive({
  total: 0,
  items: [],
} as PagedListDto<Outlay>);
const list = async () => {
  const res = await Outlay.pagedlist(filter);
  pagedList.total = res.total;
  pagedList.items = res.items;
};
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
      await list();
    } else {
      ElMessage.error('删除失败');
    }
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
  await list();
};
onActivated(init);
watchEffect(list);
</script>
