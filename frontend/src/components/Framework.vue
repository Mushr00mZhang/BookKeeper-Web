<template>
  <ElContainer>
    <ElHeader>
      <slot name="header"></slot>
    </ElHeader>
    <ElContainer>
      <ElAside width="200px">
        <slot name="aside"></slot>
      </ElAside>
      <ElMain>
        <slot name="main"></slot>
      </ElMain>
    </ElContainer>
  </ElContainer>
</template>
<script setup lang="ts">
import { ElContainer, ElHeader, ElAside, ElMain } from 'element-plus';
</script>
