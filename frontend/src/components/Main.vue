<template>
  <RouterView></RouterView>
</template>
<script setup lang="ts">
import { RouterView, useRouter } from 'vue-router';
const router = useRouter();
console.log(router);
</script>
