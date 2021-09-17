<template>
   <div class="thanks">
      <h1>Digital Production Group Request Form</h1>
      <div class="content">
         <div class="section">
            <p class="thanks">
               Thanks! We've received your request!
            </p>
            <p>You will soon receive an email containing all the information you provided.</p>
         </div>
         <div class="section">
            <h2>Order Summary</h2>
            <div class="subsection">
               <h3>Customer Information</h3>
               <div class="fields">
               </div>
            </div>
            <div class="subsection">
               <h3>Primary Address</h3>
            </div>
            <div class="subsection">
               <h3>Item Information</h3>
            </div>
         </div>
         <div class="section">
            <h2>Guideline Reminder</h2>
            <dl>
               <dt>Processing Time</dt>
               <dd>Orders generally require at least 20 business days to process.  Rush requests will be honored if possible.</dd>
               <dt>Fees</dt>
               <dd><em><strong>(Non-UVA only)</strong></em>  $50 for the first hour, $25 each additional hour.
                  Special Collections staff will send a fee estimate that must be paid before any digitization work is done.</dd>
               <dt>Deliverables</dt>
               <dd>300dpi JPEG or highest resolution TIFF, depending on your intended use.</dd>
               <dt>Delivery</dt>
               <dd>All requests are delivered as online downloads. Special delivery needs will be honored at Digital Production Group’s discretion.</dd>
               <dt>Copyright</dt>
               <dd>
                  All copies are provided for private study, scholarship, or research pursuant to 17 USC § 107 and/or 108.
                  Copyright and other legal restrictions may apply to further uses, and the Library does not grant permissions
                  or give legal advice regarding reuse of items in its collections.
                  Please see our <a targte="_blank" href="https://small.library.virginia.edu/services/publishing/">Permissions and Publishing</a> page for more information.
               </dd>
               <dt>Further Questions?</dt>
               <dd>
                  If you have questions about your request, contact
                  <a href="digitalservices@virginia.edu" target="_blank">digitalservices@virginia.edu</a> and include the following request number: <strong>{{requestID}}</strong>
               </dd>
            </dl>
         </div>
      </div>
      <div class="button-bar">
         <uva-button @click="createAnother">Submit another request</uva-button>
      </div>
   </div>
</template>

<script>
import {mapState} from "vuex"
export default {
   name: 'Thanks',
   computed: {
      ...mapState({
         error: state => state.error,
         items: state => state.items,
         request: state => state.request,
         customer: state => state.customer,
         requestID: state => state.requestID
      }),
   },
   methods: {
      createAnother() {
         this.$router.push("/")
      }
   },
}
</script>

<style lang="scss" scoped>
div.content {
   text-align: left;
   max-width: 60%;
   margin: 0 auto;
   p {
      margin: 5px 0;
   }
   p.thanks {
      font-size: 1.4em;
      font-weight: bold;
      margin-bottom: 15px;
   }
   h3 {
      font-weight: normal;
      font-size: 1.15em;
   }
   h2 {
      font-weight: normal;
      border-bottom: 1px solid var(--uvalib-grey-light);
      padding-bottom: 15px;
      margin-top: 35px;
   }
   dl {
      color: var(--uvalib-text);
      dt {
         font-weight: bold;
      }
   }
}
.button-bar {
   max-width: 60%;
   margin: 25px auto;
   text-align: right;
}
</style>
