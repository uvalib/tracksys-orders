import { createStore } from 'vuex'
import axios from 'axios'
import { getField, updateField } from 'vuex-map-fields'

export default createStore({
   state: {
      working: false,
      version: "unknown",
      error: "",
      termsAgree: false,
      termsError: false,
      isUVA: "yes",
      currStepIdx: 0,
      steps: [
         {name: "Customer Information", component: "CustomerInfo", page: 1},
         {name: "Address Information", component: "AddressInfo", page: 2},
         {name: "General Request Information", component: "RequestInfo", page: 3},
         {name: "Item Information", component: "ItemInfo", page: 4},
         {name: "Review Order", component: "ReviewOrder", page: 5},
      ],
      computeID: "",
      customer: {
         firstName: "",
         lastName: "",
         email: "",
         academicStatusID: 0,
      }
   },
   getters: {
      getField,
      currStep: state => {
         return state.steps[state.currStepIdx]
      },
      numSteps: state => {
         return state.steps.length
      }
   },
   mutations: {
      updateField,
      clearError(state) {
         state.error = ""
      },
      clearRequest(state) {
         state.currStepIdx = 0
         state.computeID = ""
         state.customer.firstName = ""
         state.customer.lastName = ""
         state.customer.email = ""
         state.customer.academicStatusID = 0
      },
      setComputeID(state, cid) {
         state.computeID = cid
      },
      setError(state, msg) {
         state.error = msg
      },
      setVersion(state, data) {
         state.version = `${data.version}-${data.build}`
      },
      setWorking(state, flag) {
         state.working = flag
      },
   },
   actions: {
      getVersion(ctx) {
         axios.get("/version").then(response => {
            ctx.commit('setVersion', response.data)
         }).catch( _e => {
            // NO-OP
         })
      },
      startRequest(ctx) {
         ctx.commit("clearRequest")
         // ctx.commit("setWorking", true)
      }
   }
})
