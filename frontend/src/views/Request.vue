<template>
   <h1>Digital Production Group Request Form</h1>
   <div class="content">
      <FormKit type="form" :actions="false">
         <FormKit type="multi-step" tab-style="progress" :allow-incomplete="false">
            <CustomerInfo />
            <AddressInfo />
            <RequestInfo />
            <ItemInfo />
            <ReviewOrder />
         </FormKit>
      </FormKit>
   </div>
   <WaitSpinner message="Processing..." :overlay="true" v-if="orderStore.working" />
</template>

<script setup>
import {useOrderStore} from '@/stores/order'
import CustomerInfo from "@/components/steps/CustomerInfo.vue"
import AddressInfo from "@/components/steps/AddressInfo.vue"
import RequestInfo from "@/components/steps/RequestInfo.vue"
import ItemInfo from "@/components/steps/ItemInfo.vue"
import ReviewOrder from "@/components/steps/ReviewOrder.vue"
import WaitSpinner from '../components/WaitSpinner.vue'

const orderStore = useOrderStore()

</script>

<style lang="scss" scoped>
div.content {
   margin: 0 auto;
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
