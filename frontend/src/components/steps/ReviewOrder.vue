<template>
   <FormKit type="step" name="reviewOrder">
      <div class="request">
         <dl>
            <dt>Date Due:</dt>
            <dd>{{orderStore.dateDue}}</dd>
            <template v-if="orderStore.specialInstructions">
               <dt>Special Instructions:</dt>
               <dd>{{orderStore.specialInstructions}}</dd>
            </template>
            <dt>Intended Use:</dt>
            <dd>{{orderStore.intendedUse(orderStore.intendedUseID)}}</dd>
         </dl>
      </div>
      <div class="items">
         <div class="item" v-for="(item,idx) in orderStore.items" :key="`item-${idx}`">
            <div class="item-bar">
               <span>Item #{{idx+1}}</span>
               <span class="buttons" v-if="orderStore.items.length > 1">
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
      <template #stepNext>
         <span class="next-btns">
            <uva-button @click="submitClicked" class="pad-left">Submit Order</uva-button>
         </span>
      </template>
   </FormKit>
</template>

<script setup>
import {useOrderStore} from '@/stores/order'

const orderStore = useOrderStore()

function deleteClicked(idx) {
   orderStore.removeItem(idx)
}
function submitClicked() {
   orderStore.submitOrder()
}
</script>

<style scoped lang="scss">
.pad-left {
   margin-left: 5px;
}
.request {
   text-align: left;
}
.items {
   text-align: left;
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
}
</style>
