import { createApp } from 'vue'
import App from './App.vue'
import router from './router'
import store from './store'

const app = createApp(App)

// provide store access to the rouer
store.router = router
router.store = store

// bind store and router to all componens as $store and $router
app.use(store)
app.use(router)

import UvaButton from "@/components/UvaButton"
app.component('UvaButton', UvaButton)

import WaitSpinner from "@/components/WaitSpinner"
app.component('WaitSpinner', WaitSpinner)

import '@fortawesome/fontawesome-free/css/all.css'

// actually mount to DOM
app.mount('#app')
