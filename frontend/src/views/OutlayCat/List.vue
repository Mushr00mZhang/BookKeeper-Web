<template>
  列表页
  <div v-for="item in items">
    {{ item.name }}
  </div>
  <!-- <ElTree></ElTree> -->
  <ElButton type="primary" @click="create">添加</ElButton>
</template>
<script setup lang="ts">
import { reactive } from 'vue';
import { useRouter } from 'vue-router';
import { OutlayCat, IListDto as OutlayCatListDto } from './model';
import { ElTree, ElButton } from 'element-plus';
import 'element-plus/es/components/tree/style/css';
import 'element-plus/es/components/button/style/css';
const router = useRouter();
const dto = reactive<OutlayCatListDto>({ parentId: null });
const items = reactive<OutlayCat[]>([]);
const create = () => {
  router.push('detail');
};
const init = async () => {
  const list = await OutlayCat.list(dto);
  items.push(...list);
};
init();
</script>
