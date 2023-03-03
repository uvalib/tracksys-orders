import { createRouter, createWebHistory } from 'vue-router'
import Home from '../views/Home.vue'
import Request from '../views/Request.vue'
import Thanks from '../views/Thanks.vue'
import Forbidden from '../views/Forbidden.vue'
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
      path: '/forbidden',
      name: 'Forbidden',
      component: Forbidden
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

router.beforeEach((to, _from, next) => {
   const orderStore = useOrderStore()
   if (to.path === '/granted') {
      console.log(`User ${to.query.user} authenticated`)
      orderStore.clearRequest()
      orderStore.termsAgreed = true
      orderStore.setComputeID(to.query.user)
      orderStore.startRequest()
      next("/request")
   } else if (to.path != "/") {
      if (orderStore.termsAgreed == false) {
         next("/")
      } else {
         next()
      }
   } else {
      next()
   }
})

export default router
