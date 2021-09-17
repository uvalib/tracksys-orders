<template>
   <div class="customer">
      <div class="form-row">
         <label for="fname">First Name</label>
         <input id="fname" type="text" v-model="firstName">
      </div>
      <div class="form-row">
         <label for="lname">Last Name</label>
         <input id="lname" type="text" v-model="lastName">
      </div>
      <div class="form-row">
         <label for="email">Email</label>
         <input id="email" type="text" v-model="email" :disabled="computeID.length > 0">
      </div>
      <div class="form-row" v-if="computeID.length > 0">
         <label for="academic-status">Academic Status</label>
         <select id="academic-status" v-model="academicStatusID">
            <option disabled value="0">Select an academic status</option>
            <option v-for="opt in academicStatuses" :key="`as-${opt.id}`" :value="opt.id">{{opt.name}}</option>
         </select>
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
import { mapState, mapGetters } from 'vuex'
export default {
   computed: {
      ...mapFields([
         'customer.firstName', 'customer.lastName', 'customer.email', 'customer.academicStatusID'
      ]),
      ...mapState({
         error: state => state.error,
         computeID: state => state.computeID
      }),
      ...mapGetters([
        'academicStatuses'
      ])
   },
   methods: {
      cancelClicked() {
         this.$store.commit("clearRequest")
         this.$router.push("/")
      },
      nextClicked() {
         if (this.computeID.length == 0) {
            this.academicStatusID = "1"
         }
         if ( this.email.length == 0) {
            this.$store.commit("setError", "Customer email is required")
            return
         }
         if ( this.academicStatusID == 0) {
            this.$store.commit("setError", "Academic status is required")
            return
         }
         this.$store.dispatch("updateCustomer")
      }
   }
};
</script>

<style scoped lang="scss">
.customer {
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
