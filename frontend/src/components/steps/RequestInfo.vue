<template>
   <FormKit type="step" name="requestInfo">
      <FormKit type="date" label="Date Due" v-model="orderStore.dateDue" :min="minDueDate" validation="required"
         help="Normal delivery time is 4 weeks from today. We make every effort to honor earlier delivery if requested, but we cannot guarantee it."
      />
      <FormKit label="Special Instructions" type="textarea" v-model="orderStore.specialInstructions" id="instruct" :rows="5"
         help="Include any additional information required to fulfill this request."
      />
      <div class="form-row">
         <label for="intended-use">How do you intend to use items in this order?</label>
         <div class="note">
            This helps us determine a digital format and watermark appropriate to your ultimate intended use.
            We do not grant permissions for specific uses, nor offer legal advice. Copyright and other legal restrictions may apply to use for
            purposes other than private study, scholarship, or research. Please see our
            <a href="https://small.library.virginia.edu/services/publishing/" target="_blank">Permissions and Publishing</a>
            page for more information.
         </div>
         <FormKit type="select" label="Intended Use" v-model="orderStore.intendedUseID" placeholder="Select an intended use" :options="intendedUses" validation="required"/>
      </div>

      <div class="intended-use-info" v-if="orderStore.intendedUseID=='100' || orderStore.intendedUseID=='104' || orderStore.intendedUseID=='106'">
         <p>
            For this use Digital Production Group will deliver <strong>300dpi JPEG</strong> images suitable for your needs.
            These images will include a copyright statment within a surrounding gray border that states:
         </p>
         <blockquote class="copyright-note" v-if="orderStore.intendedUseID=='100'" >
            This single copy was produced for the purposes of classroom teaching pursuant to 17 USC § 107 (fair use).
            Copyright and other legal restrictions may apply to further uses. Special Collections, University of Virginia Library.
         </blockquote>
         <blockquote class="copyright-note" v-if="orderStore.intendedUseID=='104' || orderStore.intendedUseID=='106'" >
            This single copy was produced for the purposes of private study, scholarship, or research pursuant to 17 USC § 107 and/or 108.
            Copyright and other legal restrictions may apply to further uses. Special Collections, University of Virginia Library.
         </blockquote>
         <p>If you require special formats such as TIFF or higher resolutions, please add a note in the Special Instructions field (above.)</p>
      </div>

      <div class="intended-use-info" v-if="orderStore.intendedUseID=='103' || orderStore.intendedUseID=='109'" >
         <p>For this use Digital Production Group will deliver <strong>300dpi JPEG</strong> images suitable for your needs.</p>
      </div>

      <div class="intended-use-info" v-if="orderStore.intendedUseID=='113'" >
         <p>For this use Digital Production Group will deliver a <strong>PDF Document</strong> containing images suitable for your needs.</p>
      </div>

      <div class="intended-use-info"  v-if="orderStore.intendedUseID=='102' || orderStore.intendedUseID=='105' || orderStore.intendedUseID=='112' " >
         <p>For this use Digital Production Group will deliver <strong>Highest possible dpi TIFF</strong>
            images suitable for your needs.  The resolution of the images is dependent on the size of the physical object.
            Generally, if the item is smaller than 11" on the longest side, the resolution will be 600dpi.
            If the item is between 11" and 14" on the longest side, the resolution will be 400dpi.
            Any item larger than 14" on the long side will generally have a resolution of 300dpi.
         </p>
      </div>
   </FormKit>
</template>

<script setup>
import { computed } from 'vue'
import {useOrderStore} from '@/stores/order'
import dayjs from 'dayjs'

const orderStore = useOrderStore()

const intendedUses = computed(()=>{
   let out = []
   if (orderStore.working) return out
   orderStore.intendedUses.forEach(l => {
      out.push( {label:  l.name, value: l.id})
   })
   return out
})

const minDueDate = computed(() => {
   let minDate =  dayjs(new Date(new Date().getTime()+(29*24*60*60*1000))).format("YYYY-MM-DD")
   return minDate
})
</script>

<style scoped lang="scss">
.form-row {
   label {
      font-weight: bold;
   }
   .note {
      margin: 5px 0 20px 0;
      .important {
         color: var(--uvalib-red);
         font-weight: bold;
      }
      p {
         margin:5px 0 0 0;
      }
   }
}
.intended-use-info {
   border: 1px solid var(--uvalib-grey-light);
   padding: 5px 10px;
   border-radius: 5px;
   margin-top: 0px;
   p {
      margin: 5px 0;
   }
}
</style>
