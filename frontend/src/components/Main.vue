<template>
  <RouterView #="{ Component }">
    <KeepAlive>
      <Transition>
        <Component :is="Component" />
      </Transition>
    </KeepAlive>
  </RouterView>
</template>
<script setup lang="ts">
import { RouterView, useRouter } from 'vue-router';
const router = useRouter();
console.log(router);
</script>
