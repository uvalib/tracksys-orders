<template>
   <div id="app">
      <div class="header" role="banner" id="uva-header">
         <div class="library-link">
            <a target="_blank" href="https://library.virginia.edu">
               <UvaLibraryLogo />
            </a>
         </div>
         <div class="site-link">
            <router-link to="/">Digitization Request Form</router-link>
            <p class="version">v{{ orderStore.version }}</p>
         </div>
      </div>
      <router-view />
      <UvaLibraryFooter />
   </div>
</template>

<script setup>
import UvaLibraryLogo from "@/components/UvaLibraryLogo.vue"
import UvaLibraryFooter from "@/components/UvaLibraryFooter.vue"
import {useOrderStore} from '@/stores/order'
import { onMounted } from 'vue'

const orderStore = useOrderStore()

onMounted( () => {
   orderStore.getVersion()
})
</script>

<style lang="scss">
#app {
   font-family: "franklin-gothic-urw", arial, sans-serif;
   -webkit-font-smoothing: antialiased;
   -moz-osx-font-smoothing: grayscale;
   text-align: center;
   color: var(--color-primary-text);
   margin: 0;
   padding: 0;
   background: white;
   font-size: 16px;
   line-height: 1.3;

   h1 {
      color: var(--uvalib-brand-orange);
      margin: 25px 0;
      font-weight: bold;
      position: relative;
   }
   a {
      color: var(--color-link);
      font-weight: 500;
      text-decoration: none;
      &:hover {
         text-decoration: underline;
      }
   }
   input[type=text], select {
      padding: 5px;
      border-radius: 5px;
      border: 1px solid var(--uvalib-grey-light);
      box-sizing: border-box;
   }
}

body {
   margin: 0;
   padding: 0;
   background-color: var(--uvalib-blue-alt-darkest);
}

div.header {
   background-color: var(--uvalib-brand-blue);
   color: white;
   padding: 1vw 20px;
   text-align: left;
   position: relative;
   box-sizing: border-box;
   display: flex;
   flex-direction: row;
   flex-wrap: nowrap;
   justify-content: space-between;
   align-content: stretch;
   align-items: center;
   div.library-link {
      height: 45px;
      width: 220px;
      order: 0;
      flex: 0 1 auto;
      align-self: flex-start;
   }
   div.site-link {
      order: 0;
      font-size: 1.5em;
      a {
         color: white !important;
         text-decoration: none;
         &:hover {
            text-decoration: underline;
         }
      }
      p.version {
         margin: 5px 0 0 0;
         font-size: 0.5em;
         text-align: right;
      }
   }
}
</style>
