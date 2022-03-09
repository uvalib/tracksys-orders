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

import UvaButton from "@/components/UvaButton.vue"
app.component('UvaButton', UvaButton)

import WaitSpinner from "@/components/WaitSpinner.vue"
app.component('WaitSpinner', WaitSpinner)

import vueCountryRegionSelect from 'vue3-country-region-select'
app.use(vueCountryRegionSelect)

import Datepicker from 'vue3-datepicker'
import 'vue3-datepicker/dist/vue3-datepicker.css'
app.component('Datepicker', Datepicker)

import '@fortawesome/fontawesome-free/css/all.css'

// actually mount to DOM
app.mount('#app')
