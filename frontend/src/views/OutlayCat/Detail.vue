<template>
  <ElForm>
    <ElFormItem label="名称">
      <ElInput v-model="outlaycat.name" />
    </ElFormItem>
    <ElFormItem label="单位">
      <ElInput v-model="outlaycat.unit" />
    </ElFormItem>
    <ElFormItem label="排序">
      <ElInputNumber v-model="outlaycat.sort" />
    </ElFormItem>
    <ElFormItem label="备注">
      <ElInput v-model="outlaycat.remark" />
    </ElFormItem>
    <ElFormItem>
      <ElButton type="primary" @click="submit">提交</ElButton>
      <ElButton @click="cancel">取消</ElButton>
    </ElFormItem>
  </ElForm>
</template>
<script setup lang="ts">
import { ref } from 'vue';
import { ElForm, ElFormItem, ElInput, ElInputNumber, ElButton } from 'element-plus';
import 'element-plus/es/components/form/style/css';
import 'element-plus/es/components/form-item/style/css';
import 'element-plus/es/components/input/style/css';
import 'element-plus/es/components/input-number/style/css';
import 'element-plus/es/components/button/style/css';
import { OutlayCat } from './model';
import { useRoute, useRouter } from 'vue-router';
const route = useRoute();
const router = useRouter();
type Query = { mode: 'create'; parentId: string | null } | { id: string; mode: 'update' };
// const { id = null, mode, parentId = null } = route.query as Query;
const query = route.query as Query;
const outlaycat = ref<OutlayCat>(
  new OutlayCat({
    id: null,
    children: [],
    parentId: null,
    name: '',
    unit: '',
    sort: 0,
    stable: false,
    remark: '',
  })
);
const submit = async () => {
  switch (query.mode) {
    case 'create':
      await outlaycat.value.create();
      break;
    case 'update':
      await outlaycat.value.update();
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
      outlaycat.value.parentId = query.parentId;
      break;
    case 'update':
      const res = await OutlayCat.get(query.id);
      if (res) {
        outlaycat.value = res;
      }
      break;
  }
};
init();
</script>
