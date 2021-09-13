import { createStore } from 'vuex'
import axios from 'axios'
import { getField, updateField } from 'vuex-map-fields'

export default createStore({
   state: {
      working: false,
      version: "unknown",
      error: "",
      termsAgree: false,
      isUVA: "yes"
   },
   getters: {
      getField,
   },
   mutations: {
      updateField,
      setVersion(state, data) {
         state.version = `${data.version}-${data.build}`
      },
      setError(state, msg) {
         state.error = msg
      },
      clearError(state) {
         state.error = ""
      },
      setWorking(state, flag) {
         state.working = flag
      }
   },
   actions: {
      getVersion(ctx) {
         axios.get("/version").then(response => {
            ctx.commit('setVersion', response.data)
         }).catch( _e => {
            // NO-OP
         })
      },
   }
})
