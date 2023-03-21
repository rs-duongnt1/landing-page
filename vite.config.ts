import { defineConfig } from 'vite'
import vue from '@vitejs/plugin-vue'
import { quasar, transformAssetUrls } from '@quasar/vite-plugin'
import path from 'path'
// https://vitejs.dev/config/
export default defineConfig({
  plugins: [
    vue({
      template: { transformAssetUrls }
    }),
    quasar({
      sassVariables: 'quasar-variables.sass'
    })
  ],
  server: {
    watch: {
     usePolling: true,
    },
    host: true, // Here
    strictPort: true,
    port: 4040, 
  },
  // build: {
  //   outDir: 'dist',
  //   rollupOptions: {
  //     input: {
  //       index: "./index.html",
  //       // "main": "./src/main.ts"
  //     },
  //     output: {

  //     }
  //   }
  // },
  resolve: {
    alias: {
      '@': path.resolve(__dirname, './webui/src')
    }
  },
})
