<template>
  列表页
  <div v-for="item in items">
    {{ item.name }}
  </div>
</template>
<script setup lang="ts">
import { reactive } from 'vue';
import { OutlayCat, IListDto as OutlayCatListDto } from './model';
const dto = reactive<OutlayCatListDto>({ parentId: null });
const items = reactive<OutlayCat[]>([]);
const init = async () => {
  const list = await OutlayCat.list(dto);
  items.push(...list);
};
init();
</script>
