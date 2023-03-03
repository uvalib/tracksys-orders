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
               <dl class="fields">
                  <dt>First Name</dt>
                  <dd>{{orderStore.customer.firstName}}</dd>
                  <dt>Last Name</dt>
                  <dd>{{orderStore.customer.lastName}}</dd>
                  <dt>Email</dt>
                  <dd>{{orderStore.customer.email}}</dd>
                  <dt>Academic Status</dt>
                  <dd>{{academicStatus}}</dd>
               </dl>
            </div>
            <div class="subsection">
               <h3>Primary Address</h3>
               <dl class="fields">
                  <dt>Address Line 1</dt>
                  <dd>{{orderStore.primaryAddress.address1}}</dd>
                  <dt>Address Line 2</dt>
                  <dd>{{orderStore.primaryAddress.address2}}</dd>
                  <dt>City</dt>
                  <dd>{{orderStore.primaryAddress.city}}</dd>
                  <dt>State</dt>
                  <dd>{{orderStore.primaryAddress.state}}</dd>
                  <dt>Zip Code</dt>
                  <dd>{{orderStore.primaryAddress.zip}}</dd>
                  <dt>Country</dt>
                  <dd>{{orderStore.primaryAddress.country}}</dd>
                  <dt>Phone</dt>
                  <dd>{{orderStore.primaryAddress.phone}}</dd>
               </dl>
            </div>
            <div class="subsection" v-if="orderStore.sameBillingAddress == false">
               <h3>Billing Address</h3>
               <dl class="fields">
                  <dt>Address Line 1</dt>
                  <dd>{{orderStore.billingAddress.address1}}</dd>
                  <dt>Address Line 2</dt>
                  <dd>{{orderStore.billingAddress.address2}}</dd>
                  <dt>City</dt>
                  <dd>{{orderStore.billingAddress.city}}</dd>
                  <dt>State</dt>
                  <dd>{{orderStore.billingAddress.state}}</dd>
                  <dt>Zip Code</dt>
                  <dd>{{orderStore.billingAddress.zip}}</dd>
                  <dt>Country</dt>
                  <dd>{{orderStore.billingAddress.country}}</dd>
                  <dt>Phone</dt>
                  <dd>{{orderStore.billingAddress.phone}}</dd>
               </dl>
            </div>
            <div class="subsection">
               <h3>Item Information</h3>
               <div v-for="item,idx in orderStore.items" class="item" :key="`item-${idx}`">
                  <div class="item-num">Item #{{idx+1}}</div>
                  <dl class="fields indent">
                     <dt>Title</dt>
                     <dd>{{item.title}}</dd>
                     <dt>Pages</dt>
                     <dd>{{item.pages}}</dd>
                     <template v-if="item.callNumber">
                        <dt>Call Number</dt>
                        <dd>{{item.callNumber}}</dd>
                     </template>
                     <template v-if="item.author">
                        <dt>Author</dt>
                        <dd>{{item.author}}</dd>
                     </template>
                     <template v-if="item.published">
                        <dt>Year Published</dt>
                        <dd>{{item.published}}</dd>
                     </template>
                     <template v-if="item.location">
                        <dt>Location</dt>
                        <dd>{{item.location}}</dd>
                     </template>
                     <template v-if="item.link">
                        <dt>Web Link</dt>
                        <dd>{{item.link}}</dd>
                     </template>
                     <template v-if="item.description">
                        <dt>Description</dt>
                        <dd>{{item.description}}</dd>
                     </template>
                  </dl>
               </div>
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
                  <a href="digitalservices@virginia.edu" target="_blank">digitalservices@virginia.edu</a> and include the following request number: <strong>{{orderStore.requestID}}</strong>
               </dd>
            </dl>
         </div>
      </div>
      <div class="button-bar">
         <uva-button @click="createAnother">Submit another request</uva-button>
      </div>
   </div>
</template>

<script setup>
import { useOrderStore } from '@/stores/order'
import { computed } from 'vue'
import { useRouter } from 'vue-router'

const router = useRouter()
const orderStore = useOrderStore()

const academicStatus = computed(() => {
   return orderStore.academicStatusName(orderStore.customer.academicStatusID)
})

function createAnother() {
   router.push("/")
}
</script>

<style lang="scss" scoped>
div.content {
   text-align: left;
   margin: 0 auto;
   p {
      margin: 5px 0;
   }
   p.thanks {
      font-size: 1.4em;
      font-weight: bold;
      margin-bottom: 15px;
   }
   .subsection {
      padding-left: 25px;
      h3 {
         font-weight: normal;
         font-size: 1.15em;
         border-bottom: 1px solid var(--uvalib-grey-light);
      }
      .item-num {
         font-weight: bold;
         margin: 20px 0 0px 25px;
      }
   }
   h2 {
      font-weight: normal;
      padding: 5px 10px;
      margin: 25px 0;
      background: #efefef;
   }
   dl {
      color: var(--uvalib-text);
      margin: 10px 0 0 15px;
      dt {
         font-weight: bold;
      }
      dd {
         margin-bottom: 10px;
      }
   }
   dl.fields {
      display: inline-grid;
      grid-template-columns: max-content 2fr;
      grid-column-gap: 15px;
      margin: 10px 0 0 15px;
      dt {
         font-weight: bold;
         text-align: right;
      }
      dd {
         margin: 0 0 5px 0;
         word-break: break-word;
         -webkit-hyphens: auto;
         -moz-hyphens: auto;
         hyphens: auto;
      }
   }
   dl.fields.indent {
      margin-left: 25px;
   }
}
.button-bar {
   max-width: 60%;
   margin: 25px auto;
   text-align: right;
}
@media only screen and (min-width: 768px) {
   div.content  {
      width: 60%;
   }
 }
 @media only screen and (max-width: 768px) {
   div.content  {
      width: 95%;
      font-size: 0.9em;
   }
 }
</style>
