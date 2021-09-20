import { createStore } from 'vuex'
import axios from 'axios'
import { getField, updateField } from 'vuex-map-fields'

export default createStore({
   state: {
      working: false,
      version: "unknown",
      error: "",
      constants: {},
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
      currAddressIdx: 0,
      addresses: [],
      request: {
         dueDate: "",
         specialInstructions: "",
         intendedUseID: 0
      },
      currItemIdx: -1,
      items: [],
      origItem: {},
      itemMode: "add",
      requestID: ""
   },
   getters: {
      getField,
      currStep: state => {
         return state.steps[state.currStepIdx]
      },
      numSteps: state => {
         return state.steps.length
      },
      academicStatuses: state => {
         return state.constants['academicStatus']
      },
      academicStatusName: state => (id) => {
         let vals = state.constants['academicStatus']
         let v = vals.find( item => item.id == id)
         if (v) {
            return v.name
         }
         return "Unknown"
      },
      intendedUses: state => {
         return state.constants['intendedUse']
      },
      intendedUse: state => (id) => {
         let uses = state.constants['intendedUse']
         let v = uses.find( item => item.id == id)
         if (v) {
            return v.name
         }
         return "Unknown"
      }
   },
   mutations: {
      updateField,
      clearItems(state) {
         state.items.splice(0, state.items.length)
      },
      editItem(state, idx) {
         state.currItemIdx = idx
         state.currStepIdx = 3
         state.itemMode =  "edit"
         state.origItem = Object.assign({}, state.items[state.currItemIdx])
      },
      itemEditCanceled(state) {
         state.items.splice(state.currItemIdx, 1,  state.origItem )
         state.currStepIdx = 4
         state.itemMode =  "add"
         state.origItem = {}
      },
      itemEditDone(state) {
         state.itemMode =  "add"
         state.currStepIdx = 4
         state.origItem = {}
      },
      addItem(state) {
         state.items.push( {title: "", pages: "", callNumber: "", author: "", published: "", location: "", link: "", description: ""} )
         state.currItemIdx = state.items.length -1
      },
      addMoreItems(state) {
         state.items.push( {title: "", pages: "", callNumber: "", author: "", published: "", location: "", link: "", description: ""} )
         state.currItemIdx = state.items.length -1
         state.itemMode =  "add"
         state.currStepIdx = 3
      },
      removeItem(state, idx) {
         if (idx < 0 || idx > state.items.length-1) return
         state.items.splice(idx, 1)
      },
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
         state.addresses.splice(0, state.addresses.length)
         state.currAddressIdx = 0
         state.items.splice(0, state.items.length)
         state.currItemIdx = 0
         state.error = ""
      },
      nextStep(state) {
         state.currStepIdx++
      },
      resetAddresses(state) {
         state.currAddressIdx = 0
         state.addresses.splice(0, state.addresses.length)
      },
      setAddresses(state, data) {
         data.forEach( addr => state.addresses.push( Object.assign({}, addr) ))
         if (state.addresses.length == 0) {
            state.addresses.push( {addressType: "primary", address1: "", address2: "", city: "", state: "", zip: "", country: "", phone: ""})
         }
      },
      nextAddress(state) {
         state.currAddressIdx++
      },
      setComputeID(state, cid) {
         state.computeID = cid
         state.customer.email = `${cid}@virginia.edu`
      },
      setConstants(state, data) {
         state.constants = data
      },
      initRequest(state) {
         let sd = new Date()
         sd.setDate(sd.getDate() + 29)
         state.request.dueDate = sd
         state.items.splice(0, state.items.length)
         state.items.push( {title: "", pages: "", callNumber: "", author: "", published: "", location: "", link: "", description: ""} )
         state.currItemIdx = 0
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
      getConstants(ctx) {
         axios.get(`/api/constants`).then(response => {
            ctx.commit("setConstants", response.data)
         }).catch( _e => {
            // NO-OP, there is just no constants
         })
      },
      startRequest(ctx) {
         ctx.dispatch("getConstants")
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
      getAddressInfo(ctx) {
         ctx.commit("setWorking", true)
         ctx.commit("resetAddresses")
         axios.get(`/api/users/${ctx.state.customer.id}/address`).then(response => {
            ctx.commit("setAddresses", response.data)
            ctx.commit("setWorking", false)
            ctx.commit("nextStep")
         }).catch( _e => {
            // NO-OP, there is just no user data pre-populated
            ctx.commit("setWorking", false)
         })
      },
      updateCustomer(ctx) {
         ctx.commit("setWorking", true)
         axios.post(`/api/users`, ctx.state.customer).then(response => {
            ctx.commit("setUserData", response.data)
            ctx.dispatch("getAddressInfo")
         }).catch( err => {
            ctx.commit("setError", err)
            ctx.commit("setWorking", false)
         })
      },
      updateAddress(ctx) {
         ctx.commit("setWorking", true)
         let currAddr = ctx.state.addresses[ ctx.state.currAddressIdx ]
         axios.post(`/api/users/${ctx.state.customer.id}/address`, currAddr).then( _response => {
            if (ctx.state.differentBillingAddress && ctx.state.currAddressIdx == 0) {
               ctx.commit("nextAddress")
               ctx.commit("setWorking", false)
            } else {
               ctx.commit("initRequest")
               ctx.commit("nextStep")
               ctx.commit("setWorking", false)
            }
         }).catch( err => {
            ctx.commit("setError", err)
            ctx.commit("setWorking", false)
         })
      },
      submitOrder(_ctx) {
         // ctx.commit("setWorking", true)
         // let req = {
         //    userID: ctx.state.customer.id,
         //    request: ctx.state.request,
         //    items: ctx.state.items
         // }
         // axios.post(`/api/submit`, req).then( _response => {
         //    // TODO
         //    ctx.commit("setWorking", false)
         // }).catch( err => {
         //    ctx.commit("setError", err)
         //    ctx.commit("setWorking", false)
         // })
         this.router.push("/thanks")
      }
   }
})
