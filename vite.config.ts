import { defineConfig } from 'vite';
import vue from '@vitejs/plugin-vue';
import ElementPlus from 'unplugin-element-plus/vite';

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
        manualChunks: {
          echarts: ['echarts'],
          'element-plus': ['element-plus'],
          'material-web': ['@material/web'],
          vue: ['vue']
        }
      }
    }
  }
});
