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
      differentBillingAddress: false,
      customer: {
         id: 0,
         firstName: "",
         lastName: "",
         email: "",
         academicStatusID: 0,
      },
      address: {
         addressType: "primary",
         id: 0,
         address1: "",
         address2: "",
         city: "",
         state: "",
         zip: "",
         country: "",
         phone: ""
      },
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
         state.address.id = 0
         state.address.type = ""
         state.address.address1 = ""
         state.address.address2 = ""
         state.address.city = ""
         state.address.state = ""
         state.address.zip = ""
         state.address.country = ""
         state.address.phone = ""
      },
      nextStep(state) {
         state.currStepIdx++
      },
      resetAddress(state, addrType) {
         state.address.id = 0
         state.address.addressType = addrType
         state.address.address1 = ""
         state.address.address2 = ""
         state.address.city = ""
         state.address.state = ""
         state.address.zip = ""
         state.address.country = ""
         state.address.phone = ""
      },
      setAddress(state, data) {
         state.address.id = data.id
         state.address.addressType = data.addressType
         state.address.address1 = data.address1
         state.address.address2 = data.address2
         state.address.city = data.city
         state.address.state = data.state
         state.address.zip = data.zip
         state.address.country = data.country
         state.address.phone = data.phone
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
      getAddressInfo(ctx, type) {
         ctx.commit("setWorking", true)
         axios.get(`/api/users/${ctx.state.customer.id}/address?type=${type}`).then(response => {
            ctx.commit("setAddress", response.data)
            ctx.commit("setWorking", false)
         }).catch( _e => {
            // NO-OP, there is just no user data pre-populated
            ctx.commit("setWorking", false)
            ctx.commit("resetAddress", type)
         })
      },
      updateCustomer(ctx) {
         ctx.commit("setWorking", true)
         axios.post(`/api/users`, ctx.state.customer).then(response => {
            ctx.commit("setUserData", response.data)
            ctx.commit("setWorking", false)
            ctx.commit("nextStep")
            ctx.dispatch("getAddressInfo", "primary")
         }).catch( err => {
            ctx.commit("setError", err)
            ctx.commit("setWorking", false)
         })
      },
      updateAddress(ctx) {
         ctx.commit("setWorking", true)
         axios.post(`/api/users/${ctx.state.customer.id}/address`, ctx.state.address).then( _response => {
            if (ctx.state.differentBillingAddress && ctx.state.address.addressType == 'primary') {
               ctx.dispatch("getAddressInfo", "billing")
            } else {
               ctx.commit("nextStep")
               ctx.commit("setWorking", false)
            }
         }).catch( err => {
            ctx.commit("setError", err)
            ctx.commit("setWorking", false)
         })
      }
   }
})
