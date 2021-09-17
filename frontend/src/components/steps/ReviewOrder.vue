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
      <div class="items">
         <div class="item" v-for="(item,idx) in items" :key="`item-${idx}`">
            <div class="item-bar">
               <span>Item #{{idx+1}}</span>
               <span class="buttons">
                  <button @click="editClicked(idx)">Edit</button>
                  <button @click="deleteClicked(idx)">Delete</button>
               </span>
            </div>
            <dl>
               <dt>Title:</dt>
               <dd>{{item.title}}</dd>
               <dt>Pages:</dt>
               <dd>{{item.pages}}</dd>
               <template v-if="item.callNumber">
                  <dt>Call Number:</dt>
                  <dd>{{item.callNumber}}</dd>
               </template>
               <template v-if="item.author">
                  <dt>Author:</dt>
                  <dd>{{item.author}}</dd>
               </template>
               <template v-if="item.published">
                  <dt>Year Published:</dt>
                  <dd>{{item.published}}</dd>
               </template>
               <template v-if="item.location">
                  <dt>Location:</dt>
                  <dd>{{item.location}}</dd>
               </template>
               <template v-if="item.link">
                  <dt>Web Link:</dt>
                  <dd>{{item.link}}</dd>
               </template>
               <template v-if="item.description">
                  <dt>Additional Info:</dt>
                  <dd>{{item.description}}</dd>
               </template>
            </dl>
         </div>
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
      editClicked(idx) {
         console.log("edit "+idx)
      },
      deleteClicked(idx) {
         this.$store.commit("removeItem", idx)
      },
      addClicked() {
         // TODO
      },
      submitClicked() {
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

   .item-bar {
      font-size: 1.15em;
      display: flex;
      flex-flow: row nowrap;
      justify-content: space-between;
      border-top: 1px solid var(--uvalib-grey-light);
      padding: 5px;
      border-bottom: 1px solid var(--uvalib-grey-light);
      button {
         margin-left: 5px;
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
