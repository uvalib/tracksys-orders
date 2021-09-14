<template>
   <h1>Digital Production Group Request Form</h1>
   <div class="content">
      <div class="step">
         <span class="name">{{currStep.name}}</span>
         <span class="sequence">Section {{currStep.page}} of {{numSteps}}</span>
      </div>
      <component :is="currStep.component"></component>
   </div>
   <WaitSpinner message="Processing..." :overlay="true" v-if="working"/>
</template>

<script>
import { mapState, mapGetters } from "vuex"
import CustomerInfo from "@/components/CustomerInfo"
export default {
   name: 'Request',
   components: {
       CustomerInfo
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
      font-size: 1.15em;
      display: flex;
      flex-flow: row nowrap;
      justify-content: space-between;
      padding: 0 5px 5px 5px;
      margin-bottom: 10px;
      border-bottom: 1px solid var(--uvalib-grey);
      font-weight: bold;
      color: var(--uvalib-text);
   }
}
</style>
