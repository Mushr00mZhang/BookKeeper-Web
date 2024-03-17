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
import { useRouter } from 'vue-router';
const router = useRouter();
type Props = {
  parentId: string | null;
};
const props = withDefaults(defineProps<Props>(), {
  parentId: null,
});
const outlaycat = ref<OutlayCat>(
  new OutlayCat({
    id: null,
    children: [],
    parentId: props.parentId,
    name: '',
    unit: '',
    sort: 0,
    stable: false,
    remark: '',
  })
);
const submit = async () => {
  await outlaycat.value.create();
  router.back();
};
</script>
