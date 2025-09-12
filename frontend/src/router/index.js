import { createRouter, createWebHistory } from 'vue-router'
import Home from '../views/Home.vue'
import Request from '../views/Request.vue'
import Thanks from '../views/Thanks.vue'
import NotFound from '../views/NotFound.vue'

import { useOrderStore } from '@/stores/order'

const routes = [
   {
      path: '/',
      name: 'Home',
      component: Home
   },
   {
      path: '/request',
      name: 'Request',
      component: Request
   },
   {
      path: '/thanks',
      name: 'Thanks',
      component: Thanks
   },
   {
      path: '/:pathMatch(.*)*',
      name: "not_found",
      component: NotFound
   }
]

const router = createRouter({
   history: createWebHistory(import.meta.env.BASE_URL),
   routes
})

router.beforeEach((to) => {
   console.log("BEFORE ROUTE "+to.path)
   const orderStore = useOrderStore()

   if (to.path === '/granted') {
      console.log(`User ${to.query.user} authenticated`)
      orderStore.clearRequest()
      orderStore.termsAgreed = true
      orderStore.setComputeID(to.query.user)
      orderStore.startRequest()
      return "/request"
   }

   if (to.path == "/") {
      console.log("NOT A PROTECTED PAGE")
   } else if (to.path == "/request") {
      // terms must be agreed to or go back to home page
      if (orderStore.termsAgreed == false) {
         return "/"
      }
      console.log(`AUTHENTICATED PAGE REQUEST WITH TERMS AGREED`)
   }
})

export default router
