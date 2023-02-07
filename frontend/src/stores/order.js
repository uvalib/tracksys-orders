import { defineStore } from 'pinia'
import axios from 'axios'

export const useOrderStore = defineStore('order', {
	state: () => ({
      working: false,
      version: "unknown",
      error: "",
      constants: {},
      currStepIdx: 0,
      steps: [
         { name: "Customer Information", component: "CustomerInfo", page: 1 },
         { name: "Address Information", component: "AddressInfo", page: 2 },
         { name: "General Request Information", component: "RequestInfo", page: 3 },
         { name: "Item Information", component: "ItemInfo", page: 4 },
         { name: "Review Order", component: "ReviewOrder", page: 5 },
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
      currItemIdx: -1,
      dateDue: "",
      specialInstructions: "",
      intendedUseID: 0,
      items: [],
      origItem: {},
      itemMode: "add",
      requestID: "",
      termsAgreed: false,
      isUVA: false
   }),
   getters: {
      currAddress: state => {
         return state.addresses[state.currAddressIdx]
      },
      currStep: state => {
         return state.steps[state.currStepIdx]
      },
      numSteps: state => {
         return state.steps.length
      },
      academicStatuses: state => {
         return state.constants['academicStatus']
      },
      academicStatusName: state => {
         // NOTE: getters cannot accept params. Instead have the getter return a function that
         // takes a param. The result will not me cached like other getters tho
         return (id) => {
            let v = state.constants['academicStatus'].find((item) => item.id === id)
            if (v) {
               return v.name
            }
            return "Unknown"
         }
      },
      intendedUses: state => {
         return state.constants['intendedUse']
      },
      intendedUse: state => (id) => {
         let uses = state.constants['intendedUse']
         let v = uses.find(item => item.id == id)
         if (v) {
            return v.name
         }
         return "Unknown"
      }
   },
   actions: {
      getVersion() {
         axios.get("/version").then(response => {
            this.version = `${response.data.version}-${response.data.build}`
         }).catch(_e => {})
      },
      getConstants() {
         axios.get(`/api/constants`).then(response => {
            this.constants = response.data
         }).catch(_e => {
            // NO-OP, there is just no constants
         })
      },
      startRequest() {
         this.getConstants()
         if (this.computeID != "") {
            this.working = true
            axios.get(`/api/users/${this.computeID}`).then(response => {
               this.setUserData(response.data)
               this.working = false
            }).catch(_e => {
               // NO-OP, there is just no user data pre-populated
               this.working = false
            })
         }
      },
      getAddressInfo() {
         this.working = true
         this.resetAddresses()
         axios.get(`/api/users/${this.customer.id}/address`).then(response => {
            this.setAddresses(response.data)
            this.working = false
            this.nextStep()
         }).catch(_e => {
            // NO-OP, there is just no user data pre-populated
            this.working = false
         })
      },
      updateCustomer() {
         this.working = true
         axios.post(`/api/users`, this.customer).then(response => {
            this.setUserData(response.data)
            this.getAddressInfo()
         }).catch(err => {
            this.setError(err)
            this.working = false
         })
      },
      updateAddress() {
         this.working = true
         let currAddr = this.addresses[this.currAddressIdx]
         axios.post(`/api/users/${this.customer.id}/address`, currAddr).then(_response => {
            if (this.differentBillingAddress && this.currAddressIdx == 0) {
               this.nextAddress()
               this.working = false
            } else {
               this.initRequest()
               this.nextStep()
               this.working = false
            }
         }).catch(err => {
            this.setError(err)
            this.working = false
         })
      },
      submitOrder() {
         this.working = true
         let req = {
            dateDue: this.dateDue,
            intendedUseID: this.intendedUseID,
            specialInstructions: this.specialInstructions,
            items: this.items,
         }
         axios.post(`/api/users/${this.customer.id}/submit`, req).then(response => {
            this.requestID = parseInt(response.data, 10)
            this.router.push("/thanks")
            this.working = false
         }).catch(err => {
            this.setError(err)
            this.working = false
         })
      },
      clearItems() {
         this.items.splice(0, this.items.length)
      },
      editItem(idx) {
         this.currItemIdx = idx
         this.currStepIdx = 3
         this.itemMode = "edit"
         this.origItem = Object.assign({}, this.items[this.currItemIdx])
      },
      itemEditCanceled() {
         this.items.splice(this.currItemIdx, 1, this.origItem)
         this.currStepIdx = 4
         this.itemMode = "add"
         this.origItem = {}
      },
      itemEditDone() {
         this.itemMode = "add"
         this.currStepIdx = 4
         this.origItem = {}
      },
      addItem() {
         this.items.push({ title: "", pages: "", callNumber: "", author: "", published: "", location: "", link: "", description: "" })
         this.currItemIdx = this.items.length - 1
      },
      addMoreItems() {
         this.items.push({ title: "", pages: "", callNumber: "", author: "", published: "", location: "", link: "", description: "" })
         this.currItemIdx = this.items.length - 1
         this.itemMode = "add"
         this.currStepIdx = 3
      },
      removeItem(idx) {
         if (idx < 0 || idx > this.items.length - 1) return
         this.items.splice(idx, 1)
      },
      clearError() {
         this.error = ""
      },
      clearRequest() {
         this.computeID = ""
         this.currStepIdx = 0
         this.customer.id = 0
         this.customer.firstName = ""
         this.customer.lastName = ""
         this.customer.email = ""
         this.customer.academicStatusID = 0
         this.addresses.splice(0, this.addresses.length)
         this.currAddressIdx = 0
         this.items.splice(0, this.items.length)
         this.currItemIdx = 0
         this.error = ""
      },
      nextStep() {
         this.currStepIdx++
         this.error = ""
      },
      resetAddresses() {
         this.currAddressIdx = 0
         this.addresses.splice(0, this.addresses.length)
      },
      setAddresses(data) {
         data.forEach(addr => this.addresses.push(Object.assign({}, addr)))
         if (this.addresses.length == 0) {
            this.addresses.push({ addressType: "primary", address1: "", address2: "", city: "", state: "", zip: "", country: "", phone: "" })
         }
      },
      nextAddress() {
         this.currAddressIdx++
         if (this.addresses.length < this.currAddressIdx + 1) {
            this.addresses.push({ addressType: "business", address1: "", address2: "", city: "", state: "", zip: "", country: "", phone: "" })
         }
      },
      setComputeID(cid) {
         this.computeID = cid
         this.customer.email = `${cid}@virginia.edu`
      },
      initRequest() {
         let sd = new Date()
         sd.setDate(sd.getDate() + 29)
         this.dateDue = sd
         this.items.splice(0, this.items.length)
         this.items.push({ title: "", pages: "", callNumber: "", author: "", published: "", location: "", link: "", description: "" })
         this.currItemIdx = 0
      },
      setError(err) {
         if (err.message) {
            this.error = err.message
         } else if (err.response) {
            this.error = err.response.data
         } else {
            this.error = err
         }
      },
      setUserData(data) {
         this.customer.id = data.id
         this.customer.firstName = data.firstName
         this.customer.lastName = data.lastName
         this.customer.email = data.email
         this.customer.academicStatusID = data.academicStatusID
      },
   },
})
