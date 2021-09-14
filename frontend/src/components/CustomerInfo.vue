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
         <input id="email" type="text" v-model="email">
      </div>
      <div class="form-row">
         <label for="academic-status">Academic Status</label>
         <select id="academic-status" v-model="academicStatusID">
            <option disabled value="0">Select an academic status</option>
            <option value="1">Non-UVA</option>
            <option value="4">Staff</option>
            <option value="5">Faculty</option>
            <option value="6">Undergraduate Student</option>
            <option value="7">Graduate Student</option>
            <option value="8">Continuing Education</option>
         </select>
      </div>
      <div class="button-bar">
         <UvaButton @click="cancelClicked">Cancel</UvaButton>
         <UvaButton @click="nextClicked" class="pad-left">Next</UvaButton>
      </div>
   </div>
</template>

<script>
import { mapFields } from 'vuex-map-fields'
export default {
   computed: {
      ...mapFields([
         'customer.firstName', 'customer.lastName', 'customer.email', 'customer.academicStatusID'
      ]),
   },
   methods: {
      cancelClicked() {
         this.$store.commit("clearRequest")
         this.$router.push("/")
      },
      nextClicked() {
         alert("next!")
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
   .button-bar {
      text-align: right;
      padding: 15px 0;
      .pad-left {
         margin-left: 10px;
      }
   }
}
</style>
