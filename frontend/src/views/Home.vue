<template>
   <h1>Digital Production Group Request Form</h1>
   <div class="content">
      <p>
         Because preservation of its materials is of paramount concern, the Special Collections Department reserves the right to
         decline to digitize materials, or to limit the amount of material that can be safely digitized. If you have questions,
         please contact Digital Production Group at <a href="mailto:digitalservices@virginia.edu">digitalservices@virginia.edu</a>.
      </p>
      <p>
         <span class="warn">Notice: Warning Concerning Copyright</span>
         <span class="dash">—</span>
         <span> copyright law of the United States (Title 17, United States Code) governs the making of photocopies or other
            reproductions of copyrighted material. Under certain conditions specified in the law, libraries and archives are authorized
            to furnish a photocopy or other reproduction. One of these specified conditions is that the photocopy or reproduction is not to
            be "used for any purpose other than private study, scholarship, or research." If a user makes a request for, or later uses a
            photocopy or reproduction for purposes in excess of "fair use," that user may be liable for copyright infringement.
            The Library's provision of a copy does not constitute permission by any copyright holder and is not a substitute for the user's diligent
            compliance with all applicable copyright requirements. This institution reserves the right to refuse to accept a copying order if, in its judgment,
            fulfillment of the order would involve violation of the copyright law.
         </span>
      </p>
      <p>
         <span class="warn">Notice: Sales Tax Collection</span>
         <span class="dash">—</span>
         The University of Virginia is required to collect sales tax for online sales delivered to the following states:
         District of Colombia, North Carolina, New Jersey, Ohio, Pennsylvania, and Virginia. If you are in one of these localities your payment
         total will include the original asking price plus the total of any sales tax owed on the purchase. If you are not located in one of these states,
         you will not have sales tax added to your purchase.
      </p>
      <p>
         <span class="warn">Notice: Sales Tax Exemption</span>
         <span class="dash">—</span>
         If you represent a tax-exempt organization, please contact <a href="mailto:digitalservices@virginia.edu">digitalservices@virginia.edu</a>
         to provide a sales tax exemption certificate. Upon reception of a valid exemption certificate, the Library will provide instructions for payment.
      </p>
      <p>
         <span class="warn">Notice: Digitizing Personal Items</span>
         <span class="dash">—</span>
         <span>
            We are not responsible for any damage to materials you provide to us for digitization.
         </span>
      </p>
      <h2>What you need to know, before you submit your request</h2>
      <div class="sect">
         <ul>
            <li>
               We normally require at least 20 working days to complete a digitization request. Rush requests will be honored if possible.
               <span class="warn">
                   Starting mid-November through mid-January additional turnaround time is required due to the holiday season.
               </span>
            </li>
            <li>
               <strong>Non-UVa Patrons:</strong>
               We charge a fee of $100 for the first hour of work (which includes a $50 processing fee) and $50 per hour for each additional hour
               (or partial hour) of expected work. Costs are calculated based on the length of time required for retrieval, identification, imaging,
               quality assurance, metadata creation, and delivery. For redelivery of images from our archive, we will only charge the processing fee ($50),
               unless it takes more than an hour to retrieve the items from our system.
               Library staff will contact you soon after your request is received with a quote based on our time estimate.
               If you accept our fee, pre-payment is required before we will proceed with your request.
               Payment information will be included with the fee estimate. Estimates are good for 30 days.
            </li>
            <li>
               We provide images at various resolutions and formats depending on the intended use of the request.
               Generally, we can provide 300dpi JPEG images, TIF images at the highest resolution possible, or PDFs.
            </li>
            <li>
               All orders will be delivered via the web.
            </li>
         </ul>
      </div>
      <h2>Terms of Agreement</h2>
      <div class="sect">
         <label class="terms" id="terms-agree" >
            <input type="checkbox" v-model="orderStore.termsAgreed">
            <span>
               I understand and agree to the above copyright statement and will only use the digitized resources as I indicate in my online request.
               I also understand it is my responsibility to determine who owns the copyright of these materials unless otherwise noted.
            </span>
         </label>
         <div v-if="termsError" class="terms-note">You must agree to the terms and conditions to continue.</div>
      </div>
      <h2>UVA Status</h2>
      <fieldset>
         <legend>>UVA Status</legend>
         <div class="opt">
            <input type="radio" name="is_uva" id="is_uva_yes" :value="true" v-model="orderStore.isUVA" class="uva-radio">
            <label for="is_uva_yes" class="radio">
               I am UVA faculty, staff, or student.  <span style="font-style: italic;">(You will be asked to verify your identity using
               <a target="_blank" href="https://virginia.service-now.com/its?id=itsweb_kb_article&sys_id=4f98738bdbf6c744f032f1f51d9619ed">NetBadge</a>.)</span>
            </label>
         </div>
         <div class="opt">
            <input type="radio" name="is_uva" id="is_uva_no" :value="false" v-model="orderStore.isUVA" class="uva-radio">
            <label for="is_uva_no" class="radio">I am not affiliated with UVA.</label>
         </div>
      </fieldset>
         <div class="button-bar">
         <uva-button @click="createRequestClicked">Create Request</uva-button>
      </div>
   </div>
</template>

<script setup>
import {useOrderStore} from '@/stores/order'
import { useRouter } from 'vue-router'
import { ref } from 'vue'

const router = useRouter()
const orderStore = useOrderStore()
const termsError = ref(false)

function createRequestClicked() {
   termsError.value = false
   if ( orderStore.termsAgreed == false) {
      termsError.value = true
      return
   }

   orderStore.clearRequest()
   if (orderStore.isUVA ) {
      window.location.href = "/authenticate"
   } else {
      orderStore.startRequest()
      router.push("/request")
   }
}
</script>

<style lang="scss" scoped>
@media only screen and (min-width: 768px) {
   div.content {
      max-width: 75%;
   }
   h2 {
      font-size: 1.5em;
   }
   .sect {
      padding-left: 25px;
   }
   fieldset {
      margin: 0;
      padding: 0 0 0 25px;
   }

}
@media only screen and (max-width: 768px) {
   div.content {
      max-width: 90%;
   }
   h2 {
      font-size: 1em;
   }
   .sect {
      padding-left: 10px;
   }
   fieldset {
      margin: 0;
      padding: 0 0 0 10px;
   }
}
div.content {
   text-align: left;
   margin: 0 auto;
   .terms-note {
      text-align: left;
      font-weight: normal;
      margin: 10px 0 20px 0;
      text-align: right;
      color: var(--uvalib-red-dark);
   }
   .warn {
      display: inline-block;
      font-weight: bold;
      color: var(--uvalib-red-dark);
      margin-right: 5px;
   }
   .dash {
      display: inline-block;
      margin: 0 5px;
   }
   h2 {
      font-weight: bold;
      border-bottom: 1px solid var(--uvalib-grey-light);
      padding-bottom: 15px;
      margin-top: 35px;
      color: var(--uvalib-brand-blue);
   }
   .terms  {
      display: flex;
      flex-flow: row nowrap;
      justify-content: flex-start;
      align-items: flex-start;
      gap: 15px;
   }
   ul {
      margin: 0;
      padding: 0 0 0 15px;
   }

   fieldset {
      display: flex;
      flex-direction: column;
      gap: 10px;
      border: none;
      outline: none;
      legend {
         display:none;
      }
      .opt {
         display: flex;
         flex-flow: row nowrap;
         gap: 10px;
         justify-content: flex-start;
         align-items: baseline;
         .uva-radio {
            width: 16px;
            height: 16px;
         }
      }
   }
}
.button-bar {
   margin: 25px 0;
   text-align: right;
}
</style>
