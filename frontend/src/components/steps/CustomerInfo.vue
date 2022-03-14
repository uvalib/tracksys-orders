<template>
   <div class="customer">
      <div class="form-row">
         <label for="fname">First Name</label>
         <input id="fname" type="text" v-model="orderStore.customer.firstName">
      </div>
      <div class="form-row">
         <label for="lname">Last Name</label>
         <input id="lname" type="text" v-model="orderStore.customer.lastName">
      </div>
      <div class="form-row">
         <label for="email">Email</label>
         <input id="email" type="text" v-model="orderStore.customer.email" :disabled="orderStore.computeID.length > 0">
      </div>
      <div class="form-row" v-if="orderStore.computeID.length > 0">
         <label for="academic-status">Academic Status</label>
         <select id="academic-status" v-model="orderStore.customer.academicStatusID">
            <option disabled value="0">Select an academic status</option>
            <option v-for="opt in orderStore.academicStatuses" :key="`as-${opt.id}`" :value="opt.id">{{opt.name}}</option>
         </select>
      </div>
      <p class="error">{{orderStore.error}}</p>
      <div class="button-bar">
         <uva-button @click="cancelClicked">Cancel</uva-button>
         <uva-button @click="nextClicked" class="pad-left">Next</uva-button>
      </div>
   </div>
</template>

<script setup>
import {useOrderStore} from '@/stores/order'
import { useRouter } from 'vue-router'

const router = useRouter()
const orderStore = useOrderStore()

function cancelClicked() {
   orderStore.clearRequest()
   router.push("/")
}

function nextClicked() {
   if (orderStore.computeID.length == 0) {
      orderStore.customer.academicStatusID = 1
   }
   if ( orderStore.customer.email.length == 0) {
      orderStore.setError("Customer email is required")
      return
   }
   if ( orderStore.customer.academicStatusID == 0) {
      orderStore.setError("Academic status is required")
      return
   }
   orderStore.updateCustomer()
}
</script>

<style scoped lang="scss">
.customer {
   text-align: left;
   padding: 15px 10%;

   .form-row {
      margin: 10px 0;
      label {
         font-weight: bold;
         display: block;
      }
      input, select {
         width: 100%;
         margin: 5px 0;
      }
   }
   .error {
      font-style: italic;
      color: var(--uvalib-red);
      margin-bottom: 0;
   }
   .button-bar {
      text-align: right;
      padding: 15px 0;
      .pad-left {
         margin-left: 10px;
      }
   }
}
</style>
