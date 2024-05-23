<template>
  <ElForm v-model="item">
    <md-outlined-text-field class="item-field" label="名称" v-model="item.name" required />
    <ElFormItem label="类型">
      <ElTreeSelect
        v-model="item.catId"
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
    <md-outlined-text-field
      class="item-field"
      label="金额"
      type="number"
      prefix-text="￥"
      v-model="item.money"
    />
    <md-outlined-text-field
      class="item-field"
      label="原价"
      type="number"
      prefix-text="￥"
      v-model="item.original"
    />
    <md-outlined-text-field class="item-field" label="数量" type="number" v-model="item.amount" />
    <md-outlined-text-field class="item-field" label="单位" v-model="item.unit" />
    <ElFormItem label="时间">
      <ElDatePicker
        v-model="item.time"
        type="datetime"
        format="YYYY-MM-DD HH:mm:ss"
        value-format="YYYY-MM-DDTHH:mm:ss+08:00"
      />
    </ElFormItem>
    <md-outlined-select class="item-field" label="用户" v-model="item.userId" required>
      <md-select-option :value="ROOT_ID">
        <div slot="headline">张天培</div>
      </md-select-option>
      <md-select-option value="11111111-1111-1111-1111-111111111111">
        <div slot="headline">陈慧芳</div>
      </md-select-option>
    </md-outlined-select>
    <ElFormItem>
      <ElButton type="primary" @click="submit">提交</ElButton>
      <ElButton @click="cancel">取消</ElButton>
    </ElFormItem>
  </ElForm>
</template>
<script setup lang="ts">
import '@/utils/Date';
import { ref, reactive, onActivated } from 'vue';
import { useRoute, useRouter } from 'vue-router';
import {
  ElForm,
  ElFormItem,
  ElTreeSelect,
  ElDatePicker,
  ElButton,
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
import { Outlay } from './model';
import { OutlayCat, ROOT, ROOT_ID } from '@/views/OutlayCat/model';
import '@material/web/textfield/outlined-text-field';
import '@material/web/select/outlined-select';
import '@material/web/select/select-option';
const route = useRoute();
const router = useRouter();
type Query = { mode: 'create' } | { id: string; mode: 'update' };
const query = ref(route.query as Query);
const treeProps: TreeOptionProps = {
  label: 'name',
  isLeaf: (data, _node) => !(data as OutlayCat)?.hasChildren,
  disabled: (_data, node) => node?.level <= 1,
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
      resolve(items);
      break;
  }
};
const expandedIds = reactive(new Set([ROOT_ID]));
const item = ref<Outlay>(
  new Outlay({
    id: undefined,
    name: '',
    catId: '',
    money: 0,
    original: 0,
    amount: 0,
    unit: '',
    time: new Date().format('YYYY-MM-DDTHH:mm:ss+08:00'),
    userId: ROOT_ID,
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
      if (res.result) {
        emit('updated', item.value);
        ElMessage.success('更新成功');
        router.back();
      } else {
        let tip = '更新失败';
        switch (res.code) {
          default:
            break;
          // case 5:
          //   tip += '，名称重复';
          //   break;
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
      expandedIds.clear();
      expandedIds.add(ROOT_ID);
      item.value = new Outlay({
        id: undefined,
        name: '',
        catId: '',
        money: 0,
        original: 0,
        amount: 0,
        unit: '',
        time: new Date().format('YYYY-MM-DDTHH:mm:ss+08:00'),
        userId: '',
        cat: {
          parentId: '',
          name: '',
          unit: '',
          sort: 0,
          stable: false,
          remark: '',
        },
      });
      break;
    case 'update':
      const res = await Outlay.get(query.value.id);
      if (res) {
        if (res.cat.parentId) {
          const ids = await OutlayCat.getParentIds(res.cat.parentId);
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
  created: [item: Outlay];
  updated: [item: Outlay];
};
const emit = defineEmits<Emits>();
onActivated(init);
// init();
</script>
<style scoped lang="scss">
.item-field {
  font-family: 'Courier New', Courier, monospace;
  & + & {
    margin-top: 20px;
  }
}
</style>
