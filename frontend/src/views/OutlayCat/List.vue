<template>
  <div v-for="item in items">
    {{ item.name }}
  </div>
</template>
<script setup lang="ts">
import { reactive } from 'vue';
import { OutlayCat } from './model';
const items = reactive<OutlayCat[]>([]);
</script>
