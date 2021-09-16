<template>
   <div class="item">
      <div class="note pad-bottom">
         In the boxes below, include the item, image, or page numbers to be scanned. If possible, include any additional
         descriptive information such as, item call number, box number, and folder or item description. Please add one item for each individual collection or book you would like digitized.
      </div>
      <div class="form-row">
         <label for="title">Title (required)</label>
         <input id="title" type="text" v-model="item.title">
      </div>
      <div class="form-row">
         <label for="pages">Image or page numbers (required)</label>
         <textarea rows="2" id="pages" v-model="item.pages"></textarea>
      </div>
      <div class="form-row">
         <label for="callnum">Call Number</label>
         <input id="callnum" type="text" v-model="item.callNumber">
      </div>
      <div class="form-row">
         <label for="author">Author</label>
         <input id="author" type="text" v-model="item.author">
      </div>
      <div class="form-row">
         <label for="published">Year Published</label>
         <input id="published" type="text" v-model="item.published">
      </div>
      <div class="form-row">
         <label for="location">Location</label>
         <input id="location" type="text" v-model="item.location">
      </div>
      <div class="form-row">
         <label for="link">Web link for item (if available)</label>
         <input id="link" type="text" v-model="item.link">
      </div>
      <div class="form-row">
         <label for="pages">Additional Description</label>
         <div class="note">
            If the provided fields were insufficient to describe the item, include supplemental information above.
         </div>
         <textarea rows="2" id="pages" v-model="item.pages"></textarea>
      </div>
      <p class="error">{{error}}</p>
      <div class="button-bar">
         <uva-button @click="cancelClicked">Cancel</uva-button>
         <uva-button @click="addClicked" class="pad-left">Add Item</uva-button>
         <uva-button @click="completeClicked" class="pad-left">Complete Order</uva-button>
      </div>
   </div>
</template>

<script>
import { mapState } from 'vuex'
export default {
   data() {
      return {
         item: {title: "", pages: "", callNumber: "", author: "", published: "", location: "", link: "", description: ""}
      }
   },
   computed: {
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
      addClicked() {
         // TODO
      },
      completeClicked() {
         // TODO
      }
   }
};
</script>

<style scoped lang="scss">
.item {
   text-align: left;
   padding: 15px 10%;

   .note {
      font-style: italic;
      font-size: 0.9em;
      margin: 5px 0;
   }
   .pad-bottom {
      margin-bottom: 25px;
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
      textarea {
         box-sizing: border-box;
         width: 100%;
         border: 1px solid var(--uvalib-grey-light);
         border-radius: 5px;
         padding: 5px 10px;
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
