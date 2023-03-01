import { createApp, markRaw } from 'vue'
import App from './App.vue'
import router from './router'
import { createPinia } from 'pinia'
import { createI18n } from 'vue-i18n'

import VueDatePicker from '@vuepic/vue-datepicker'
import '@vuepic/vue-datepicker/dist/main.css'

const i18n = createI18n({
   legacy: false,
   locale: 'en',
})

const pinia = createPinia()
pinia.use(({ store }) => {
   store.router = markRaw(router)
})

const app = createApp(App)

app.use( pinia )
app.use( router )
app.use( i18n )
app.component('VueDatePicker', VueDatePicker)

import UvaButton from "@/components/UvaButton.vue"
app.component('UvaButton', UvaButton)

import vueCountryRegionSelect from 'vue3-country-region-select'
app.use(vueCountryRegionSelect)

import '@fortawesome/fontawesome-free/css/all.css'

// actually mount to DOM
app.mount('#app')
