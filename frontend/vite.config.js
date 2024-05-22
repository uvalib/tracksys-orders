 /*global process */

import { fileURLToPath, URL } from 'url'
import { defineConfig } from 'vite'
import vue from '@vitejs/plugin-vue'

// https://vitejs.dev/config/
export default defineConfig({
   define: {
      // enable hydration mismatch details in production build
      __VUE_PROD_HYDRATION_MISMATCH_DETAILS__: 'true'
    },
   plugins: [vue()],
   resolve: {
      alias: {
         '@': fileURLToPath(new URL('./src', import.meta.url))
      }
   },
   server: { // this is used in dev mode only
      port: 8080,
      proxy: {
         '/version': {
            target: process.env.TS_ORDER_SRV,  // export TS_ORDER_SRV=http://localhost:8085
            changeOrigin: true
         },
         '/healthcheck': {
            target: process.env.TS_ORDER_SRV,
            changeOrigin: true
         },
         '/api': {
            target: process.env.TS_ORDER_SRV,
            changeOrigin: true
         },
         '/authenticate': {
            target: process.env.TS_ORDER_SRV,
            changeOrigin: true
         },
      }
   },
   css: {
      preprocessorOptions: {
        scss: {
           // example : additionalData: `@import "./src/design/styles/variables";`
           // dont need include file extend .scss
           additionalData: `
             @import "@/assets/variables.scss";
             @import "@/assets/mixins.scss";
          `
       },
      },
    },
})


