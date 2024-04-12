<template>
  <ElContainer class="framework">
    <ElHeader>
      <ElButton class="menu-switch" @click="toggleAside" :icon="Menu" />
      <slot name="header"></slot>
    </ElHeader>
    <ElContainer>
      <Transition>
        <ElAside :width="showAside ? '1.2rem' : '0rem'">
          <slot name="aside"></slot>
        </ElAside>
      </Transition>
      <ElMain>
        <slot name="main"></slot>
      </ElMain>
    </ElContainer>
  </ElContainer>
</template>
<script setup lang="ts">
import { ref } from 'vue';
import { useRouter } from 'vue-router';
import { ElButton, ElContainer, ElHeader, ElAside, ElMain } from 'element-plus';
import { Menu } from '@element-plus/icons-vue';
import 'element-plus/es/components/button/style/css';
// const useRoute
const router = useRouter();
router.afterEach(() => {
  showAside.value = false;
});
const showAside = ref(false);
const toggleAside = () => {
  showAside.value = !showAside.value;
};
</script>
<style lang="scss">
.framework {
  width: 100%;
  height: 100%;
  > .el-header {
    height: 0.33rem;
    border-bottom: 0.01rem solid;
    padding: 0;
    display: flex;
    .menu-switch {
      border: 0;
      padding: 0;
      height: 0.32rem;
      width: 0.32rem;
    }
  }
  > .el-container {
    overflow: hidden;
    > .el-aside {
      transition: width 0.1s;
    }
  }
}
</style>
