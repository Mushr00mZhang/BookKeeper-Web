<template>
  <section class="aside-links">
    <RouterLink to="/">首页</RouterLink>
    <RouterLink to="/outlay/list">支出列表</RouterLink>
    <RouterLink to="/outlaycat/list">支出类型列表</RouterLink>
  </section>
</template>
<script setup lang="ts">
import { RouterLink } from 'vue-router';
</script>
<style lang="scss" scoped>
.aside-links {
  display: flex;
  flex-direction: column;
}
</style>
