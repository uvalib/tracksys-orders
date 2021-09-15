import { createStore } from 'vuex'
import axios from 'axios'
import { getField, updateField } from 'vuex-map-fields'

export default createStore({
   state: {
      working: false,
      version: "unknown",
      error: "",
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
         id: 0,
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
         state.computeID = ""
         state.currStepIdx = 0
         state.customer.id = 0
         state.customer.firstName = ""
         state.customer.lastName = ""
         state.customer.email = ""
         state.customer.academicStatusID = 0
      },
      nextStep(state) {
         state.currStepIdx++
      },
      setComputeID(state, cid) {
         state.computeID = cid
         state.customer.email = `${cid}@virginia.edu`
      },
      setError(state, err) {
         if (err.message) {
            state.error = err.message
         } else if (err.response) {
            state.error = err.response.data
         } else {
            state.error = err
         }
      },
      setVersion(state, data) {
         state.version = `${data.version}-${data.build}`
      },
      setUserData(state, data) {
         state.customer.id = data.id
         state.customer.firstName = data.firstName
         state.customer.lastName = data.lastName
         state.customer.email = data.email
         state.customer.academicStatusID = data.academicStatusID
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
         if (ctx.state.computeID != "") {
            ctx.commit("setWorking", true)
            axios.get(`/api/users/${ctx.state.computeID}`).then(response => {
               ctx.commit("setUserData", response.data)
               ctx.commit("setWorking", false)
            }).catch( _e => {
               // NO-OP, there is just no user data pre-populated
               ctx.commit("setWorking", false)
            })
         }
      },
      async updateCustomer(ctx) {
         ctx.commit("setWorking", true)
         try {
            await axios.post(`/api/users`, ctx.state.customer)
            ctx.commit("setWorking", false)
            ctx.commit("nextStep")
         } catch (err) {
            ctx.commit("setError", err)
            ctx.commit("setWorking", false)
         }
      }
   }
})
