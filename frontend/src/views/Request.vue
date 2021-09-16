<template>
   <h1>Digital Production Group Request Form</h1>
   <div class="content">
      <div class="step">
         <span class="name" v-if="currStep.component != 'ItemInfo'">{{currStep.name}}</span>
         <span class="name" v-else>Item {{currItemIdx+1}} Information</span>
         <span class="sequence">Section {{currStep.page}} of {{numSteps}}</span>
      </div>
      <component :is="currStep.component"></component>
   </div>
   <wait-spinner message="Processing..." :overlay="true" v-if="working"/>
</template>

<script>
import { mapState, mapGetters } from "vuex"
import CustomerInfo from "@/components/CustomerInfo"
import AddressInfo from "@/components/AddressInfo"
import RequestInfo from "@/components/RequestInfo"
import ItemInfo from "@/components/ItemInfo"
export default {
   name: 'Request',
   components: {
       CustomerInfo, AddressInfo, RequestInfo, ItemInfo
   },
   computed: {
      ...mapState({
         working : state => state.working,
         currItemIdx: state => state.currItemIdx
      }),
      ...mapGetters([
        'currStep', 'numSteps'
      ])
   },
   methods: {
   }
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
