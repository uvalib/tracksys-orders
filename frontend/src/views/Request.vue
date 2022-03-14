<template>
   <h1>Digital Production Group Request Form</h1>
   <div class="content">
      <div class="step">
         <span class="name" v-if="orderStore.currStep.component != 'ItemInfo'">{{orderStore.currStep.name}}</span>
         <span class="name" v-else>Item {{orderStore.items.length+1}} Information</span>
         <span class="sequence">Section {{orderStore.currStep.page}} of {{orderStore.numSteps}}</span>
      </div>
      <component :is="orderStore.currStep.component"></component>
   </div>
   <wait-spinner message="Processing..." :overlay="true" v-if="orderStore.working"/>
</template>

<script>
import {useOrderStore} from '@/stores/order'
import CustomerInfo from "@/components/steps/CustomerInfo.vue"
import AddressInfo from "@/components/steps/AddressInfo.vue"
import RequestInfo from "@/components/steps/RequestInfo.vue"
import ItemInfo from "@/components/steps/ItemInfo.vue"
import ReviewOrder from "@/components/steps/ReviewOrder.vue"
export default {
   name: 'Request',
   components: {
       CustomerInfo, AddressInfo, RequestInfo, ItemInfo, ReviewOrder
   },
   setup() {
      const orderStore = useOrderStore()
      return { orderStore }
   },
}
</script>

<style lang="scss" scoped>
div.content {
   text-align: center;
   margin: 0 auto;
   .step {
      display: flex;
      flex-flow: row nowrap;
      justify-content: space-between;
      background: var(--uvalib-grey);
      color: white;
      border: 0;
      padding: 5px 10px;
      border-radius: 5px;
      width: 100%;
      box-sizing: border-box;
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
