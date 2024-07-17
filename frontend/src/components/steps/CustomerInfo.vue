<template>
   <FormKit type="step" name="customerInfo" :before-step-change="beforeStepChange">
      <FormKit label="First Name" type="text" v-model="orderStore.customer.firstName" validation="required"  id="fname" autofocus/>
      <FormKit label="Last Name" type="text" v-model="orderStore.customer.lastName" validation="required"  id="lname"/>
      <FormKit label="Email" type="email" v-model="orderStore.customer.email" validation="required"  id="email" :disabled="orderStore.computeID.length > 0"/>
      <FormKit v-if="orderStore.computeID.length > 0" type="select"
         label="Academic Status" v-model="orderStore.customer.academicStatusID"
         placeholder="Select an academic status" :options="academicStatuses" validation="required"/>
   </FormKit>
</template>

<script setup>
import { computed } from 'vue'
import { useOrderStore } from '@/stores/order'

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
</style>
