<template>
  <ElContainer class="framework">
    <ElHeader>
      <ElButton class="menu-switch" @click="toggleAside" :icon="Menu" />
      <slot name="header"></slot>
    </ElHeader>
    <ElContainer>
      <Transition>
        <ElAside :width="showAside ? '120px' : '0'">
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
    height: 33px;
    border-bottom: 1px solid;
    padding: 0;
    display: flex;
    .menu-switch {
      border: 0;
      padding: 0;
      height: 32px;
      width: 32px;
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
