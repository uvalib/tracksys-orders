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

import vueCountryRegionSelect from 'vue3-country-region-select'
app.use(vueCountryRegionSelect)

import Datepicker from 'vue3-date-time-picker'
import 'vue3-date-time-picker/dist/main.css'
app.component('Datepicker', Datepicker)

import '@fortawesome/fontawesome-free/css/all.css'

// actually mount to DOM
app.mount('#app')
