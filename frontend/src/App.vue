<template>
  <div class="common-layout">
    <Framework>
      <template #header>
        <Header></Header>
      </template>
      <template #aside>
        <Aside></Aside>
      </template>
      <template #main>
        <Main></Main>
      </template>
    </Framework>
  </div>
</template>
<script setup lang="ts">
import Framework from '@/components/Framework.vue';
import Header from '@/components/Header.vue';
import Aside from '@/components/Aside.vue';
import Main from '@/components/Main.vue';
</script>
<style scoped>
.logo {
  height: 6em;
  padding: 1.5em;
  will-change: filter;
  transition: filter 300ms;
}
.logo:hover {
  filter: drop-shadow(0 0 2em #646cffaa);
}
.logo.vue:hover {
  filter: drop-shadow(0 0 2em #42b883aa);
}
</style>
