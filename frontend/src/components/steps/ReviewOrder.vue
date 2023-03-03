<template>
   <FormKit type="step" name="reviewOrder">
      <template #stepNext>
         <span class="next-btns">
            <uva-button @click="submitClicked" class="pad-left">Submit Order</uva-button>
         </span>
      </template>
   </FormKit>
   <!-- <div class="review">
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
      <p class="error">{{orderStore.error}}</p>
      <div class="button-bar">
         <uva-button @click="cancelClicked">Cancel</uva-button>
         <uva-button @click="addClicked" class="pad-left">Add Items</uva-button>
         <uva-button @click="submitClicked" class="pad-left">Submit Order</uva-button>
      </div>
   </div> -->
</template>

<script setup>
import {useOrderStore} from '@/stores/order'
import { useRouter } from 'vue-router'

const router = useRouter()
const orderStore = useOrderStore()

function cancelClicked() {
   orderStore.clearRequest()
   router.push("/")
}
function editClicked(idx) {
   orderStore.editItem(idx)
}
function deleteClicked(idx) {
   orderStore.removeItem(idx)
}
function addClicked() {
   orderStore.addMoreItems()
}
function submitClicked() {
   orderStore.submitOrder()
}
</script>

<style scoped lang="scss">
.pad-left {
   margin-left: 5px;
}
</style>
