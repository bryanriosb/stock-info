import { fileURLToPath, URL } from 'node:url'

import { defineConfig } from 'vite'
import vue from '@vitejs/plugin-vue'
import vueDevTools from 'vite-plugin-vue-devtools'

// https://vite.dev/config/
export default defineConfig({
  plugins: [
    vue(),
    vueDevTools(),
  ],
  resolve: {
    alias: {
      '@': fileURLToPath(new URL('./src', import.meta.url))
    },
  },
  server: {
    host: '0.0.0.0',
    watch: {
      usePolling: true, // Required for Docker file watching
    },
    proxy: {
      '/api': {
        target: process.env.DOCKER_ENV ? 'http://backend:5000' : 'http://localhost:5000',
        changeOrigin: true,
      }
    }
  }
})
