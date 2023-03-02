<template>
   <FormKit type="step" name="addressInfo" :before-step-change="beforeStepChange">
      <h2>Primary Address</h2>
      <FormKit label="Address Line 1" type="text" v-model="orderStore.primaryAddress.address1" validation="required" id="address1"/>
      <FormKit label="Address Line 2" type="text" v-model="orderStore.primaryAddress.address2" id="address2"/>
      <FormKit label="City" type="text" v-model="orderStore.primaryAddress.city" id="city" validation="required" />
      <FormKit label="State" type="text" v-model="orderStore.primaryAddress.state" id="state" validation="required" />
      <FormKit label="Zip Code" type="text" v-model="orderStore.primaryAddress.zip" id="zip" validation="required" />
      <div class="form-row">
         <label for="country">Country</label>
         <country-select id="country" v-model="orderStore.primaryAddress.country" :country="orderStore.primaryAddress.country" topCountry="US" :countryName="true" :usei18n="false" />
      </div>
      <FormKit label="Phone" type="text" v-model="orderStore.primaryAddress.phone" id="city"/>
      <div class="billing">
         <h2>
            <span>Billing Address</span>
            <label>
               Same as primary addrress
               <input type="checkbox" v-model="orderStore.sameBillingAddress">
            </label>
         </h2>
         <template v-if="orderStore.sameBillingAddress == false">
            <FormKit label="Address Line 1" type="text" v-model="orderStore.billingAddress.address1" validation="required" id="address1"/>
            <FormKit label="Address Line 2" type="text" v-model="orderStore.billingAddress.address2" id="address2"/>
            <FormKit label="City" type="text" v-model="orderStore.billingAddress.city" id="city" validation="required" />
            <FormKit label="State" type="text" v-model="orderStore.billingAddress.state" id="state" validation="required" />
            <FormKit label="Zip Code" type="text" v-model="orderStore.billingAddress.zip" id="zip" validation="required" />
            <div class="form-row">
               <label for="country">Country</label>
               <country-select id="country" v-model="orderStore.billingAddress.country" :country="orderStore.billingAddress.country" topCountry="US" :countryName="true" :usei18n="false" />
            </div>
            <FormKit label="Phone" type="text" v-model="orderStore.billingAddress.phone" id="city"/>
         </template>
      </div>
   </FormKit>
</template>

<script setup>
import {useOrderStore} from '@/stores/order'

const orderStore = useOrderStore()

async function beforeStepChange({currentStep}) {
   if ( currentStep.blockingCount == 0) {
      await orderStore.updateAddress()
      return orderStore.error == ""
   }
}
</script>

<style scoped lang="scss">
h2 {
   font-size: 1em;
   margin: 0 0 15px 0;
   text-align: left;
   display: flex;
   flex-flow: row nowrap;
   justify-content: space-between;
   align-items: center;
   label {
      font-weight: 500;
      display: flex;
      flex-flow: row nowrap;
      justify-content: space-between;
      align-items: center;

      input[type=checkbox] {
         height: 20px;
         width: 20px;
         display: inline-block;
         margin-left: 10px;
      }
   }
}
.billing {
   border-top: 1px solid var(--uvalib-grey-light);
   margin: 20px 0;
   padding-top: 20px;
}
.form-row {
   margin-top: 20px;
   text-align: left;
   label {
      display:block;
   }
   select {
      width: 100%;
   }
}
</style>
