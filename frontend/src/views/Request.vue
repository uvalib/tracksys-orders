<template>
   <h1>Digital Production Group Request Form</h1>
   <div class="content">
      <div class="step">
         <span class="name">{{currStep.name}}</span>
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
export default {
   name: 'Request',
   components: {
       CustomerInfo, AddressInfo, RequestInfo
   },
   computed: {
      ...mapState({
         working : state => state.working,
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
   max-width: 60%;
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
   }
}
</style>
