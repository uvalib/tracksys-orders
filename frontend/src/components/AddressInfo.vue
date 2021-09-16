<template>
   <div class="address">
      <h2 v-if="addressType=='primary'">Primary Address Information</h2>
      <h2 v-else>Billing Address Information</h2>
      <div class="form-row">
         <label for="address1">Address Line 1</label>
         <input id="address1" type="text" v-model="address1">
      </div>
      <div class="form-row">
         <label for="address2">Address Line 2</label>
         <input id="address2" type="text" v-model="address2">
      </div>
      <div class="form-row">
         <label for="city">City</label>
         <input id="city" type="text" v-model="city">
      </div>
      <div class="form-row">
         <label for="state">State</label>
         <input id="state" type="text" v-model="state">
      </div>
      <div class="form-row">
         <label for="zip">Zip Code</label>
         <input id="zip" type="text" v-model="zip">
      </div>
      <div class="form-row">
         <label for="country">Country</label>
         <country-select id="country" v-model="country" :country="country" topCountry="US" :countryName="true" :usei18n="false" />
      </div>
      <div class="form-row">
         <label for="phone">Phone</label>
         <input id="phone" type="text" v-model="phone">
      </div>
      <div class="form-row bill" v-if="addressType=='primary'">
         <label for="billing">Do you have a different billing address?</label>
         <span class="checkbox"><input id="billing" type="checkbox" v-model="differentBillingAddress"><span>Yes</span></span>
      </div>
      <p class="error">{{error}}</p>
      <div class="button-bar">
         <uva-button @click="cancelClicked">Cancel</uva-button>
         <uva-button @click="nextClicked" class="pad-left">Next</uva-button>
      </div>
   </div>
</template>

<script>
import { mapFields } from 'vuex-map-fields'
import { mapState } from 'vuex'
export default {
   computed: {
      ...mapFields([
         'address.addressType', 'address.address1', 'address.address2', 'address.city', 'address.state',
         'address.zip', 'address.country', 'address.phone', 'differentBillingAddress'
      ]),
      ...mapState({
         error: state => state.error,
      }),
   },
   methods: {
      cancelClicked() {
         this.$store.commit("clearRequest")
         this.$router.push("/")
      },
      nextClicked() {
         this.$store.dispatch("updateAddress")
      }
   },
};
</script>

<style scoped lang="scss">
.address {
   text-align: left;
   padding: 15px 10%;

   h2 {
      text-align: center;
      margin: 0 0 30px 0;
      font-size: 1.15em;
   }

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
      input[type=checkbox] {
         display: inline-block;
         width: 20px;
         padding: 0;
         height: 20px;
         margin-right: 10px;
      }
      .checkbox {
         display: flex;
         flex-flow: row nowrap;
         justify-content: flex-start;
         align-items: center;
         margin-top: 5px;
      }
   }
   .form-row.bill {
      margin-top: 20px;
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
