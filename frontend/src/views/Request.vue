<template>
   <h1>Digital Production Group Request Form</h1>
   <div class="content">
      <div v-if="orderStore.computeID.length == 0 && showPrompt" class="customer">
         <h3>Patron Account</h3>
         <div class="help">
            If you have placed an order with us before, enter your email address and click 
            'Prior Patron' to lookup your account; otherwise, click 'New Patron'.
         </div>
         <div class="lookup">
            <input v-model="emailLookup" placeholder="Patron email" @keydown="orderStore.lookupFailed = false" autofocus />
            <button @click="lookupPatron">Prior Patron</button>
            <button @click="showPrompt=false">New Patron</button>
            <button @click="router.push('/')" class="cancel">Cancel Request</button>
         </div>
         <div class="fail" v-if="orderStore.lookupFailed">Unable to find an existing patron with email {{ emailLookup }}.</div>
      </div>
      <FormKit v-else type="form" :actions="false">
         <FormKit type="multi-step" tab-style="progress" :allow-incomplete="false">
            <CustomerInfo />
            <AddressInfo />
            <RequestInfo />
            <ItemInfo />
            <ReviewOrder />
         </FormKit>
      </FormKit>
   </div>
</template>

<script setup>
import CustomerInfo from "@/components/steps/CustomerInfo.vue"
import AddressInfo from "@/components/steps/AddressInfo.vue"
import RequestInfo from "@/components/steps/RequestInfo.vue"
import ItemInfo from "@/components/steps/ItemInfo.vue"
import ReviewOrder from "@/components/steps/ReviewOrder.vue"
import { computed, ref } from 'vue'
import { useOrderStore } from '@/stores/order'
import { useRouter } from 'vue-router'

const orderStore = useOrderStore()
const showPrompt = ref(true)
const emailLookup = ref("")
const router = useRouter()


const lookupPatron = (async () => {
   await orderStore.lookupPatron( emailLookup.value ) 
   if (orderStore.lookupFailed == false ) {
      showPrompt.value = false
   }
})

</script>

<style lang="scss" scoped>
div.content {
   margin: 0 auto;
}
.fail {
   font-weight: bold;
   color: var(--uvalib-red-dark);
}
.customer {
   border: 1px solid #ccccd7;
   border-radius: 0.3rem;
   padding: 20px;
   margin: 0 0 50px 0;
   display: flex;
   flex-direction: column;
   gap: 15px;
   h3 {
      margin: 0;
   }
   .lookup {
      display: flex;
      flex-wrap: row wrap;
      gap: 10px;
      justify-content: flex-start;
      align-items: stretch;
      .cancel {
         margin-left: auto;
      }
      input {
         max-width: 300px;
         flex-grow:1;
         padding: 5px 10px;
         border-radius: 0.3rem;
         border:1px solid #ccccd7;
         &:focus {
            outline: 2px dotted var( --uvalib-accessibility-highlight );
            outline-offset: 3px;
         }
      }
      button {
         white-space: nowrap;
         border-radius: 0.3rem;
         border:1px solid #ccccd7;
         cursor: pointer;
         transition: background-color 0.25s ease-in-out;
         &:hover {
            background-color: #ddddd7;
         }
         &:focus {
            outline: 2px dotted var( --uvalib-accessibility-highlight );
            outline-offset: 3px;
         }
      }
   }
}
@media only screen and (min-width: 768px) {
   div.content  {
      width: 60%;
   }
 }
 @media only screen and (max-width: 768px) {
   div.content  {
      width: 95%;
   }
 }
</style>
