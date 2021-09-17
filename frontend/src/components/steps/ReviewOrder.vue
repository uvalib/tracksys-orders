<template>
   <div class="review">
      <div class="request">
         <dl>
            <dt>Date Due:</dt>
            <dd>{{request.dueDate}}</dd>
            <template v-if="request.specialInstructions">
               <dt>Special Instructions:</dt>
               <dd>{{request.specialInstructions}}</dd>
            </template>
            <dt>Intended Use:</dt>
            <dd>{{intendedUse(request.intendedUseID)}}</dd>
         </dl>
      </div>
      <div class="button-bar">
         <uva-button @click="cancelClicked">Cancel</uva-button>
         <uva-button @click="addClicked" class="pad-left">Add Item</uva-button>
         <uva-button @click="submitClicked" class="pad-left">Submit Order</uva-button>
      </div>
   </div>
</template>

<script>
import { mapState, mapGetters } from 'vuex'
export default {
   computed: {
      ...mapState({
         error: state => state.error,
         items: state => state.items,
         request: state => state.request
      }),
      ...mapGetters([
         'intendedUse'
      ])
   },
   methods: {
      cancelClicked() {
         this.$store.commit("clearRequest")
         this.$router.push("/")
      },
      addClicked() {
         // TODO
      },
      sibmitClicked() {
         // TODO
      }
   }
};
</script>

<style scoped lang="scss">
.review {
   text-align: left;
   padding: 15px 10%;

   dl {
      margin-left: 25px;
      display: inline-grid;
      grid-template-columns: max-content 2fr;
      grid-column-gap: 15px;
      dt {
         font-weight: bold;
         text-align: right;
      }
      dd {
         margin: 0 0 10px 0;
         word-break: break-word;
         -webkit-hyphens: auto;
         -moz-hyphens: auto;
         hyphens: auto;
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
