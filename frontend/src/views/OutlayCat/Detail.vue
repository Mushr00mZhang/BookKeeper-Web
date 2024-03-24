<template>
  <ElForm v-model="item">
    <ElFormItem label="父级">
      <ElTreeSelect
        v-model="item.parentId"
        :load="load"
        :props="treeProps"
        :default-expanded-keys="[...expandedIds]"
        node-key="id"
        value-key="id"
        lazy
        check-strictly
        filterable
        ref="treeRef"
      />
    </ElFormItem>
    <ElFormItem label="名称">
      <ElInput v-model="item.name" />
    </ElFormItem>
    <ElFormItem label="单位">
      <ElInput v-model="item.unit" />
    </ElFormItem>
    <ElFormItem label="排序">
      <ElInputNumber v-model="item.sort" />
    </ElFormItem>
    <ElFormItem label="固定">
      <ElCheckbox v-model="item.stable" />
    </ElFormItem>
    <ElFormItem label="备注">
      <ElInput v-model="item.remark" />
    </ElFormItem>
    <ElFormItem>
      <ElButton type="primary" @click="submit">提交</ElButton>
      <ElButton @click="cancel">取消</ElButton>
    </ElFormItem>
  </ElForm>
</template>
<script setup lang="ts">
import { ref, reactive, onActivated } from 'vue';
import { useRoute, useRouter } from 'vue-router';
import {
  ElForm,
  ElFormItem,
  ElTreeSelect,
  ElInput,
  ElInputNumber,
  ElButton,
  ElCheckbox,
  ElMessage,
  TreeInstance,
} from 'element-plus';
import {
  LoadFunction,
  TreeData,
  TreeOptionProps,
} from 'element-plus/es/components/tree/src/tree.type.mjs';
import Node from 'element-plus/es/components/tree/src/model/node';
import 'element-plus/es/components/tree/style/css';
import 'element-plus/es/components/form/style/css';
import 'element-plus/es/components/form-item/style/css';
import 'element-plus/es/components/input/style/css';
import 'element-plus/es/components/input-number/style/css';
import 'element-plus/es/components/button/style/css';
import { OutlayCat, ROOT, ROOT_ID } from './model';
const route = useRoute();
const router = useRouter();
type Query = { mode: 'create'; parentId: string } | { id: string; mode: 'update' };
const query = ref(route.query as Query);
const treeProps: TreeOptionProps = {
  label: 'name',
  isLeaf: (data, _node) => !(data as OutlayCat).hasChildren,
};
const treeRef = ref<TreeInstance>();
const load: LoadFunction = async (node: Node, resolve: (data: TreeData) => void) => {
  switch (node.level) {
    case 0:
      resolve([ROOT]);
      break;
    default:
      const parentId = (node.data as OutlayCat).id;
      const items = await OutlayCat.list({ parentId });
      if (query.value.mode === 'create') {
        resolve(items);
      } else if (query.value.mode === 'update') {
        resolve(items.filter((i) => i.id !== item.value.id));
      } else {
        resolve([]);
      }
      break;
  }
};
const expandedIds = reactive(new Set([ROOT_ID]));
const item = ref<OutlayCat>(
  new OutlayCat({
    id: undefined,
    parentId: '',
    name: '',
    unit: '',
    sort: 0,
    stable: false,
    remark: '',
    // outlays: [],
    // children: [],
    hasChildren: false,
  })
);
const submit = async () => {
  switch (query.value.mode) {
    case 'create':
      const id = await item.value.create();
      if (id) {
        emit('created', item.value);
        ElMessage.success('创建成功');
        router.back();
      } else {
        ElMessage.error('创建失败');
      }
      break;
    case 'update':
      const res = await item.value.update();
      if (res?.result) {
        emit('updated', item.value);
        ElMessage.success('更新成功');
        router.back();
      } else {
        let tip = '更新失败';
        switch (res?.code) {
          default:
            break;
          case 5:
            tip += '，名称重复';
            break;
        }
        ElMessage.error(tip);
      }
      break;
  }
};
const cancel = () => {
  router.back();
};
const init = async () => {
  query.value = route.query as Query;
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
  switch (query.value.mode) {
    case 'create':
      if (query.value.parentId) {
        const ids = await OutlayCat.getParentIds(query.value.parentId);
        expandedIds.clear();
        for (const id of ids) {
          expandedIds.add(id);
        }
      }
      item.value.parentId = query.value.parentId;
      break;
    case 'update':
      const res = await OutlayCat.get(query.value.id);
      if (res) {
        if (res.parentId) {
          const ids = await OutlayCat.getParentIds(res.parentId);
          expandedIds.clear();
          for (const id of ids) {
            expandedIds.add(id);
          }
        }
        item.value = res;
      }
      break;
  }
};
type Emits = {
  created: [item: OutlayCat];
  updated: [item: OutlayCat];
};
const emit = defineEmits<Emits>();
onActivated(init);
</script>
