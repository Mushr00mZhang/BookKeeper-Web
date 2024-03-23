<template>
  <ElForm v-model="item">
    <ElFormItem label="名称">
      <ElInput v-model="item.name" />
    </ElFormItem>
    <ElFormItem label="类型">
      <ElTreeSelect
        v-model="item.catId"
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
    <ElFormItem label="金额">
      <ElInputNumber v-model="item.money" />
    </ElFormItem>
    <ElFormItem label="原价">
      <ElInputNumber v-model="item.original" />
    </ElFormItem>
    <ElFormItem label="数量">
      <ElInputNumber v-model="item.amount" />
    </ElFormItem>
    <ElFormItem label="单位">
      <ElInput v-model="item.unit" />
    </ElFormItem>
    <ElFormItem label="时间">
      <ElDatePicker v-model="item.time" type="datetime" />
    </ElFormItem>
    <ElFormItem label="用户">
      <ElSelect v-model="item.userId">
        <ElOption label="空" :value="ROOT_ID" />
      </ElSelect>
    </ElFormItem>
    <ElFormItem>
      <ElButton type="primary" @click="submit">提交</ElButton>
      <ElButton @click="cancel">取消</ElButton>
    </ElFormItem>
  </ElForm>
</template>
<script setup lang="ts">
import '@/utils/Date';
import { ref, reactive } from 'vue';
import { useRoute, useRouter } from 'vue-router';
import {
  ElForm,
  ElFormItem,
  ElTreeSelect,
  ElInput,
  ElInputNumber,
  ElDatePicker,
  ElButton,
  ElMessage,
  ElSelect,
  ElOption,
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
import { Outlay } from './model';
import { OutlayCat, ROOT, ROOT_ID } from '@/views/OutlayCat/model';
const route = useRoute();
const router = useRouter();
type Query = { mode: 'create' } | { id: string; mode: 'update' };
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
      resolve(items);
      break;
  }
};
const parentIds = reactive<string[]>([ROOT_ID]);
const item = ref<Outlay>(
  new Outlay({
    id: undefined,
    name: '',
    catId: '',
    money: 0,
    original: 0,
    amount: 0,
    unit: '',
    time: new Date().format('YYYY-MM-DD HH:mm:ss'),
    userId: '',
    cat: {
      parentId: '',
      name: '',
      unit: '',
      sort: 0,
      stable: false,
      remark: '',
    },
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
      break;
    case 'update':
      const res = await Outlay.get(query.id);
      if (res) {
        item.value = res;
      }
      break;
  }
};
type Emits = {
  created: [item: Outlay];
  updated: [item: Outlay];
};
const emit = defineEmits<Emits>();
init();
</script>
