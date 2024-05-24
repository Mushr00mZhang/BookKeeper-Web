import { defineConfig } from 'vite';
import vue from '@vitejs/plugin-vue';
import ElementPlus from 'unplugin-element-plus/vite';
import { visualizer } from 'rollup-plugin-visualizer';

const chunkPackages = [
  'axios',
  'echarts',
  'element-plus',
  '@element-plus',
  '@material/web',
  'lodash-es',
  'vue',
  '@vue',
  'vue-router',
  'zrender',
];
// https://vitejs.dev/config/
export default defineConfig({
  base: '',
  plugins: [
    ElementPlus({}),
    vue({
      template: {
        compilerOptions: {
          isCustomElement: (tag) => tag.startsWith('md'),
        },
      },
    }),
    // visualizer({
    //   open: true, // 注意这里要设置为true，否则无效
    //   gzipSize: true, // 分析图生成的文件名
    //   brotliSize: true, // 收集 brotli 大小并将其显示
    // }),
  ],
  resolve: {
    alias: {
      '@': '/src/',
    },
  },
  server: {
    port: 4001,
    proxy: {
      '/api': 'http://localhost:4000',
      // '/api': {
      //   target: 'http://localhost:8080',
      //   changeOrigin: true,
      //   rewrite: (path) => path.replace(/^\/api/, ''),
      // },
    },
  },
  build: {
    rollupOptions: {
      output: {
        manualChunks(id, meta) {
          const info = meta.getModuleInfo(id);
          if (!info.isIncluded) return;
          const pkg = chunkPackages.find((i) => new RegExp(`/${i}/`).test(id));
          if (pkg) return pkg.replace(/^[@/]+|[@/]+$/g, '').replace(/[@/]+/g, '-');
          return;
        },
      },
    },
  },
});
