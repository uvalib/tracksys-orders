<template>
   <div class="request">
      <div class="form-row">
         <label for="due">Date Due (required)</label>
         <div class="note">
            Normal delivery time is 4 weeks from today. We make every effort to honor earlier delivery if requested, but we cannot guarantee it.
            <span class="important">Starting mid-November through mid-January additional turnaround time is required due to the holiday season.</span>
         </div>
         <Datepicker v-model="dueDate" :enable-time-picker="false"
            placeholder="YYYY-MM-DD" :format="formatDate"
            :auto-apply="true" :weekStart="0">
         </Datepicker>
      </div>
      <div class="form-row">
         <label for="instruct">Special Instructions</label>
         <div class="note">
            Include any additional information required to fulfill this request.
         </div>
         <textarea rows="5" id="instruct" v-model="specialInstructions"></textarea>
      </div>
      <div class="form-row">
         <label for="intended-use">How do you intend to use items in this order? (required)</label>
         <div class="note">
            This helps us determine a digital format and watermark appropriate to your ultimate intended use.
            We do not grant permissions for specific uses, nor offer legal advice. Copyright and other legal restrictions may apply to use for
            purposes other than private study, scholarship, or research. Please see our
            <a href="https://small.library.virginia.edu/services/publishing/" target="_blank">Permissions and Publishing</a>
            page for more information.
         </div>
         <select id="intended-use" v-model="intendedUseID">
            <option disabled value="0">Select an intended use</option>
            <option v-for="opt in intendedUses" :key="`u-${opt.id}`" :value="opt.id">{{opt.name}}</option>
         </select>
      </div>

      <div class="intended-use-info" v-if="intendedUseID=='100' || intendedUseID=='104' || intendedUseID=='106'">
         <p>
            For this use Digital Production Group will deliver <strong>300dpi JPEG</strong> images suitable for your needs.
            These images will include a copyright statment within a surrounding gray border that states:
         </p>
         <blockquote class="copyright-note" v-if="intendedUseID=='100'" >
            This single copy was produced for the purposes of classroom teaching pursuant to 17 USC § 107 (fair use).
            Copyright and other legal restrictions may apply to further uses. Special Collections, University of Virginia Library.
         </blockquote>
         <blockquote class="copyright-note" v-if="intendedUseID=='104' || intendedUseID=='106'" >
            This single copy was produced for the purposes of private study, scholarship, or research pursuant to 17 USC § 107 and/or 108.
            Copyright and other legal restrictions may apply to further uses. Special Collections, University of Virginia Library.
         </blockquote>
         <p>If you require special formats such as TIFF or higher resolutions, please add a note in the Special Instructions field (above.)</p>
      </div>

      <div class="intended-use-info" v-if="intendedUseID=='103' || intendedUseID=='109'" >
         <p>For this use Digital Production Group will deliver <strong>300dpi JPEG</strong> images suitable for your needs.</p>
      </div>

      <div class="intended-use-info" v-if="intendedUseID=='113'" >
         <p>For this use Digital Production Group will deliver a <strong>PDF Document</strong> containing images suitable for your needs.</p>
      </div>

      <div class="intended-use-info"  v-if="intendedUseID=='102' || intendedUseID=='105' || intendedUseID=='112' " >
         <p>For this use Digital Production Group will deliver <strong>Highest possible dpi TIFF</strong>
            images suitable for your needs.  The resolution of the images is dependent on the size of the physical object.
            Generally, if the item is smaller than 11" on the longest side, the resolution will be 600dpi.
            If the item is between 11" and 14" on the longest side, the resolution will be 400dpi.
            Any item larger than 14" on the long side will generally have a resolution of 300dpi.
         </p>
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
         'request.dueDate', 'request.specialInstructions', 'request.intendedUseID'
      ]),
      ...mapState({
         error: state => state.error,
      }),
      ...mapGetters([
        'intendedUses'
      ])
   },
   methods: {
      formatDate(date) {
         const day = `${date.getDate()}`
         const month = `${date.getMonth()+1}`
         const year = date.getFullYear()
         return `${year}-${month.padStart(2,"0")}-${day.padStart(2,"0")}`
      },
      cancelClicked() {
         this.$store.commit("clearRequest")
         this.$router.push("/")
      },
      nextClicked() {
         if (this.dueDate == "" || this.dueDate == null) {
            this.$store.commit("setError", "Due date is required")
            return
         }
         if (this.intendedUseID == 0) {
            this.$store.commit("setError", "Intended use is required")
            return
         }
         this.$store.commit("nextStep")
      }
   }
};
</script>

<style scoped lang="scss">
.request {
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
      textarea {
         box-sizing: border-box;
         width: 100%;
         border: 1px solid var(--uvalib-grey-light);
         border-radius: 5px;
         padding: 5px 10px;
         margin: 5px 0;
         font-family: "franklin-gothic-urw", arial, sans-serif;
         -webkit-font-smoothing: antialiased;
         -moz-osx-font-smoothing: grayscale;
      }
      .note {
         font-style: italic;
         margin: 5px 0;
         font-size: 0.9em;
         .important {
            color: var(--uvalib-red);
            font-weight: bold;
         }
      }
   }
   .intended-use-info {
      font-size: 0.9em;
      border: 1px solid var(--uvalib-grey-light);
      padding: 5px 10px;
      border-radius: 5px;
      margin-top: 0px;
      p {
         margin: 5px 0;
      }
      blockquote {
         font-size: 1.1em;
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