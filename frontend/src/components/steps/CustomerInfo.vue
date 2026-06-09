<template>
   <FormKit type="step" name="customerInfo" :before-step-change="beforeStepChange">
      <FormKit label="First Name" type="text" v-model="orderStore.customer.firstName" validation="required"  id="fname" autofocus/>
      <FormKit label="Last Name" type="text" v-model="orderStore.customer.lastName" validation="required"  id="lname"/>
      <template  v-if="orderStore.computeID.length > 0">
         <div>
            <FormKit label="Email" type="email" v-model="orderStore.customer.email" :disabled="true"/>
            <div class="help">If multiple email recipients are needed, please note them in the 'Special Instructions' section of the 'Request Info' step.</div>
         </div>
         <FormKit type="select" label="Academic Status" v-model="orderStore.customer.academicStatusID"
            placeholder="Select an academic status" :options="academicStatuses" validation="required"
         />
      </template>
      <div v-else >
         <div class="email-entry">
            <FormKit label="Email" type="email" v-model="orderStore.customer.email" 
               validation="required|email" validation-visibility="live" id="email" :disabled="orderStore.customer.id >0 && !editEmail"
            />
            <button v-if="orderStore.customer.id >0 && !editEmail" @click="editEmail=true">Change</button>
         </div>
         <div class="help">If multiple email recipients are needed, please note them in the 'Special Instructions' section of the 'Request Info' step.</div>
      </div>
   </FormKit>
</template>

<script setup>
import { computed, ref } from 'vue'
import { useOrderStore } from '@/stores/order'


const editEmail = ref(false)
const orderStore = useOrderStore()

const academicStatuses = computed(()=>{
   let out = []
   if (orderStore.working) return out
   orderStore.academicStatuses.forEach(l => {
      out.push( {label:  l.name, value: l.id})
   })
   return out
})

async function beforeStepChange({currentStep}) {
   if ( currentStep.blockingCount == 0) {
      await orderStore.updateCustomer()
      return orderStore.error == ""
   }
}
</script>

<style scoped lang="scss">
.help {
   margin-top: 10px;
   font-size: 0.95em;
}
.email-entry {
   display: flex;
   flex-flow: row nowrap;
   gap: 10px;
   justify-content: flex-start;
   align-items: flex-end;
   .formkit-outer {
      flex-grow: 1;
   }
   button {
      white-space: nowrap;
      border-radius: 0.3rem;
      border:1px solid #ccccd7;
      cursor: pointer;
      padding: 7px 10px;
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
</style>
