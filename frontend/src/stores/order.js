import { defineStore } from 'pinia'
import axios from 'axios'
import moment from 'moment'

export const useOrderStore = defineStore('order', {
	state: () => ({
      working: false,
      version: "unknown",
      error: "",
      academicStatuses: [],
      intendedUses: [],
      computeID: "",
      customer: {
         id: 0,
         firstName: "",
         lastName: "",
         email: "",
         academicStatusID: 1,
      },
      sameBillingAddress: true,
      primaryAddress: {
         addressType: "primary",
         address1: "",
         address2: "",
         city: "",
         state: "",
         zip: "",
         country: "",
         phone: "",
      },
      billingAddress: {
         addressType: "billable_address",
         address1: "",
         address2: "",
         city: "",
         state: "",
         zip: "",
         country: "",
         phone: "",
      },
      currItemIdx: -1,
      dateDue: "",
      specialInstructions: "",
      intendedUseID: null,
      items: [],
      origItem: {},
      requestID: "",
      termsAgreed: false,
      isUVA: false
   }),
   getters: {
      academicStatusName: state => {
         // NOTE: getters cannot accept params. Instead have the getter return a function that
         // takes a param. The result will not me cached like other getters tho
         return (id) => {
            let v = state.academicStatuses.find((item) => item.id === id)
            if (v) {
               return v.name
            }
            return "Unknown"
         }
      },
      intendedUse: state => (id) => {
         let v = state.intendedUses.find(item => item.id == id)
         if (v) {
            return v.name
         }
         return "Unknown"
      },
   },
   actions: {
      getVersion() {
         axios.get("/version").then(response => {
            this.version = `${response.data.version}-${response.data.build}`
         }).catch(_e => {})
      },
      getConstants() {
         axios.get(`/api/constants`).then(response => {
            this.academicStatuses = response.data.academicStatus
            this.intendedUses = response.data.intendedUse
         }).catch(_e => {
            // NO-OP, there is just no constants
         })
      },
      async startRequest() {
         this.getConstants()
         if (this.computeID != "") {
            this.working = true
            return axios.get(`/api/users/${this.computeID}`).then(response => {
               this.setUserData(response.data)
               this.working = false
            }).catch(_e => {
               // NO-OP, there is just no user data pre-populated
               this.working = false
            })
         }
      },
      async updateCustomer() {
         this.working = true
         return axios.post(`/api/users`, this.customer).then(response => {
            this.setUserData(response.data.customer)
            this.setAddresses(response.data.addresses)
            this.working = false
         }).catch(err => {
            this.setError(err)
            this.working = false
         })
      },
      async updateAddress() {
         this.working = true
         var data = {
            primary: this.primaryAddress,
            differentBilling: false
         }
         if (this.sameBillingAddress == false ) {
            data.differentBilling = true
            data.billing = this.billingAddress
         }
         return axios.post(`/api/users/${this.customer.id}/address`, data).then(response => {
            this.setAddresses(response.data)
            this.working = false
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
      addItem() {
         this.items.push({ title: "", pages: "", callNumber: "", author: "", published: "", location: "", link: "", description: "" })
         this.currItemIdx = this.items.length - 1
      },
      removeItem(idx) {
         if (idx < 0 || idx > this.items.length - 1) return
         this.currItemIdx = 0
         this.items.splice(idx, 1)
      },
      clearRequest() {
         this.computeID = ""
         this.customer.id = 0
         this.customer.firstName = ""
         this.customer.lastName = ""
         this.customer.email = ""
         this.customer.academicStatusID = 1
         this.sameBillingAddress = true
         this.primaryAddress = { addressType: "primary", address1: "", address2: "", city: "", state: "", zip: "", country: "", phone: "" }
         this.billingAddress = { addressType: "billable_address", address1: "", address2: "", city: "", state: "", zip: "", country: "", phone: "" }
         this.error = ""
         let sd = new Date()
         sd.setDate(sd.getDate() + 29)
         this.dateDue = moment(sd).format("YYYY-MM-DD")
         this.items.splice(0, this.items.length)
         this.items.push({ title: "", pages: "", callNumber: "", author: "", published: "", location: "", link: "", description: "" })
         this.currItemIdx = 0
      },

      setAddresses(data) {
         if ( data.length == 0 ) {
            this.sameBillingAddress = true
            this.primaryAddress = { addressType: "primary", address1: "", address2: "", city: "", state: "", zip: "", country: "", phone: "" }
            this.billingAddress = { addressType: "billable_address", address1: "", address2: "", city: "", state: "", zip: "", country: "", phone: "" }
         } else {
            this.primaryAddress = data[0]
            this.sameBillingAddress = true
            if (data.length == 2) {
               this.sameBillingAddress = false
               this.billingAddress = data[1]
            }
         }
      },
      setComputeID(cid) {
         this.computeID = cid
         this.customer.email = `${cid}@virginia.edu`
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
