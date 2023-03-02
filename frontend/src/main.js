import { createApp, markRaw } from 'vue'
import App from './App.vue'
import router from './router'
import { createPinia } from 'pinia'
import { createI18n } from 'vue-i18n'
import { plugin, defaultConfig } from '@formkit/vue'
import { createMultiStepPlugin } from '@formkit/addons'
import '@formkit/addons/css/multistep'

import './assets/forms.scss'
import './assets/uva-colors.css'

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

app.component('VueDatePicker', VueDatePicker)
app.use( pinia )
app.use( router )
app.use( i18n )

const dc = defaultConfig({
   plugins: [createMultiStepPlugin()],
   config: {
      classes: {
         input: '$reset dpg-form-input',
         label: '$reset dpg-form-label',
         messages: '$reset dpg-form-invalid',
      },
   }
})
app.use(plugin, dc)

import UvaButton from "@/components/UvaButton.vue"
app.component('UvaButton', UvaButton)

import vueCountryRegionSelect from 'vue3-country-region-select'
app.use(vueCountryRegionSelect)

import '@fortawesome/fontawesome-free/css/all.css'

// actually mount to DOM
app.mount('#app')
