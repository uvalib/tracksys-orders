<template>
   <div class="address">
      <div class="form-row">
         <label for="address1">Address Line 1</label>
         <input id="address1" type="text" v-model="address1">
      </div>
      <div class="form-row">
         <label for="address2">Address Line 2</label>
         <input id="address2" type="text" v-model="address2">
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
import UvaButton from './UvaButton.vue';
export default {
   components: { UvaButton },
   computed: {
      ...mapFields([
         'address.addressType', 'address.address1', 'address.address2', 'address.city', 'address.state',
         'address.zip', 'address.country', 'address.phone', 'differentBillingAddress'
      ]),
      ...mapState({
         error: state => state.error,
         computeID: state => state.computeID
      }),
   },
   methods: {
      cancelClicked() {
         this.$store.commit("clearRequest")
         this.$router.push("/")
      },
      nextClicked() {
         // TODO
      }
   },
};
</script>

<style scoped lang="scss">
.address {
   text-align: left;
   padding: 15px 10%;

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
