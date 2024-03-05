<template>
  <RouterView #="{ Component }">
    <Transition>
      <KeepAlive>
        <Component :is="Component" />
      </KeepAlive>
    </Transition>
  </RouterView>
</template>
<script setup lang="ts">
import { RouterView, useRouter } from 'vue-router';
const router = useRouter();
console.log(router);
</script>
