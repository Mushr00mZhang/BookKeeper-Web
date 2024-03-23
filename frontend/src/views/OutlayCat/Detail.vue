<template>
  <ElForm v-model="item">
    <ElFormItem label="父级">
      <ElTreeSelect
        v-model="item.parentId"
        :load="load"
        :props="treeProps"
        :default-expanded-keys="parentIds"
        node-key="id"
        value-key="id"
        lazy
        check-strictly
        filterable
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
import { ref, reactive } from 'vue';
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
const query = route.query as Query;
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
      resolve(items.filter((i) => i.id !== item.value.id));
      break;
  }
};
const parentIds = reactive<string[]>([ROOT_ID]);
const item = ref<OutlayCat>(
  new OutlayCat({
    id: undefined,
    // children: [],
    parentId: '',
    name: '',
    unit: '',
    sort: 0,
    stable: false,
    remark: '',
    hasChildren: false,
  })
);
const submit = async () => {
  switch (query.mode) {
    case 'create':
      const id = await item.value.create();
      if (id) {
        emit('created', item.value);
        ElMessage.success('创建成功');
      } else {
        ElMessage.error('创建失败');
      }
      break;
    case 'update':
      const res = await item.value.update();
      if (res) {
        emit('updated', item.value);
        ElMessage.success('更新成功');
      } else {
        ElMessage.error('更新失败');
      }
      break;
  }
  router.back();
};
const cancel = () => {
  router.back();
};
const init = async () => {
  switch (query.mode) {
    case 'create':
      if (query.parentId) {
        const ids = await OutlayCat.getParentIds(query.parentId);
        parentIds.splice(1);
        parentIds.push(...ids.slice(1));
      }
      item.value.parentId = query.parentId;
      break;
    case 'update':
      const res = await OutlayCat.get(query.id);
      if (res) {
        if (res.parentId) {
          const ids = await OutlayCat.getParentIds(res.parentId);
          parentIds.splice(1);
          parentIds.push(...ids.slice(1));
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
init();
</script>
