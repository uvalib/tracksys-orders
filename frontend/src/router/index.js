import { createRouter, createWebHistory } from 'vue-router'
import Home from '../views/Home.vue'
import Request from '../views/Request.vue'
import Thanks from '../views/Thanks.vue'
import Forbidden from '../views/Forbidden.vue'
import NotFound from '../views/NotFound.vue'

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
   history: createWebHistory(process.env.BASE_URL),
   routes
})

router.beforeEach((to, _from) => {
   if (to.path === '/granted') {
      console.log(`User ${to.query.user} authenticated`)
      router.store.commit("setComputeID", to.query.user)
      router.store.dispatch("startRequest")
      return  "/request"
   }
})

export default router
