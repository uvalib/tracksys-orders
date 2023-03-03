<template>
   <FormKit type="step" name="items">
      <div class="item-label">
         <span>Item {{ orderStore.currItemIdx+1 }} of {{ orderStore.items.length }}</span>
         <span class="paging">
            <uva-button @click="deleteItem" :disabled="orderStore.items.length==1" class="trash" title="Delete item"><i class="far fa-trash-alt"></i></uva-button>
            <uva-button @click="prevItem" class="pad-left big" :disabled="orderStore.currItemIdx==0" title="Prior item"><i class="fas fa-chevron-left"></i></uva-button>
            <uva-button @click="nextItem" class="pad-left" :disabled="orderStore.currItemIdx == orderStore.items.length-1" title="Next item"><i class="fas fa-chevron-right"></i></uva-button>
         </span>
      </div>
      <div class="help pad-bottom">
         In the boxes below, include the item, image, or page numbers to be scanned. If possible, include any additional
         descriptive information such as, item call number, box number, and folder or item description. Please add one item for each individual collection or book you would like digitized.
      </div>
      <div class="item">
         <FormKit label="Title" type="text" v-model="item.title" validation="required" id="title" autofocus/>
         <FormKit label="Image or page numbers" type="text" v-model="item.pages" validation="required" id="pages"/>
         <FormKit label="Call Number" type="text" v-model="item.callNumber" id="callnum"/>
         <FormKit label="Author" type="text" v-model="item.author" id="author"/>
         <FormKit label="Year Published" type="text" v-model="item.published" id="published"/>
         <FormKit label="Location" type="text" v-model="item.location" id="location"/>
         <FormKit label="Web link for item (if available)" type="text" validation="url" v-model="item.link" id="link"/>
         <FormKit label="Additional Description" type="textarea" :rows="2" v-model="item.description" id="description"
            help="If the provided fields were insufficient to describe the item, include supplemental information above."
         />
         <div class="help pad-top">
            Use the Add Item button to add addtional items to your order. When complete, click Complete Order to review and submit your order.
         </div>
      </div>
      <template #stepNext="{ handlers, node }">
         <span class="next-btns">
            <uva-button @click="addClicked(node.context)" class="pad-left">Add Item</uva-button>
            <uva-button @click="handlers.incrementStep(1, node.context)()" class="pad-left" data-next="true">Complete Order</uva-button>
         </span>
      </template>
   </FormKit>
</template>

<script setup>
import { useOrderStore } from '@/stores/order'
import { computed, nextTick } from 'vue'

const orderStore = useOrderStore()

const item = computed(() => {
   return orderStore.items[orderStore.currItemIdx]
})

function deleteItem() {
   orderStore.removeItem(orderStore.currItemIdx)
}
function prevItem() {
   if ( orderStore.currItemIdx > 0) {
      orderStore.currItemIdx--
   }
}
function nextItem() {
   if ( orderStore.currItemIdx < orderStore.items.length) {
      orderStore.currItemIdx++
   }
}
function addClicked( context ) {
   if ( context.blockingCount == 0) {
      orderStore.addItem()
      nextTick( () => {
         let titleInput = document.getElementById("title")
         titleInput.focus()
      })
   }
}
</script>

<style scoped lang="scss">
.help {
   text-align: left;
   background: white;
}
.pad-bottom {
   margin-bottom: 10px;
   margin-left: 5px;
}
.pad-top {
   margin-top: 25px
}
.pad-left {
   margin-left: 5px;
}
.pad-left.big {
   margin-left: 10px;
}
.item-label {
   text-align: left;
   font-size: 1.15em;
   font-weight: bold;
   background: var(--uvalib-grey-lightest);
   padding: 5px 10px;
   border-radius: 5px;
   margin-bottom: 10px;
   display: flex;
   flex-flow: row nowrap;
   justify-content: space-between;
   .paging {
      padding: 0;
      button.uva-button {
         padding: 4px 10px;
      }
   }
}
.item {
   padding-left: 5px;;
}
.next-btns {
   margin-left: auto;
}
</style>
