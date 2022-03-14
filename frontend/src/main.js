import { createApp, markRaw } from 'vue'
import App from './App.vue'
import router from './router'
import { createPinia } from 'pinia'

const pinia = createPinia()
pinia.use(({ store }) => {
   store.router = markRaw(router)
})

const app = createApp(App)

app.use( pinia )
app.use( router )

import UvaButton from "@/components/UvaButton.vue"
app.component('UvaButton', UvaButton)

import vueCountryRegionSelect from 'vue3-country-region-select'
app.use(vueCountryRegionSelect)

import Datepicker from 'vue3-datepicker'
import 'vue3-datepicker/dist/vue3-datepicker.css'
app.component('Datepicker', Datepicker)

import '@fortawesome/fontawesome-free/css/all.css'

// actually mount to DOM
app.mount('#app')
