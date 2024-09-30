<template>
   <Transition>
      <div v-if="showMessage" class="added-overlay">
         <div class="message">
            <div class="msg-title">Success</div>
            <div class="body">
               An item has successfully been added to your order.
            </div>
            <div class="bar">
               <uva-button @click="toggleAddMessage()" id="close-msg">OK</uva-button>
            </div>
         </div>
      </div>
   </Transition>
   <FormKit type="step" name="items">
      <div class="item-label">
         <span>Item {{ orderStore.currItemIdx+1 }} of {{ orderStore.items.length }}</span>
         <span class="btn-group">
            <uva-button @click="deleteItem" :disabled="orderStore.items.length==1" class="trash" title="Delete item"><i class="far fa-trash-alt"></i></uva-button>
            <uva-button @click="prevItem" class="pad-left big" :disabled="orderStore.currItemIdx==0" title="Prior item"><i class="fas fa-chevron-left"></i></uva-button>
            <uva-button @click="nextItem" class="pad-left" :disabled="orderStore.currItemIdx == orderStore.items.length-1" title="Next item"><i class="fas fa-chevron-right"></i></uva-button>
         </span>
      </div>
      <div class="help">
         In the boxes below, include the item, image, or page numbers to be scanned. If possible, include any additional
         descriptive information such as, item call number, box number, and folder or item description. Please add one item for each individual collection or book you would like digitized.
      </div>
      <div class="item">
         <FormKit label="Title" type="text" v-model="item.title" id="title" @node="titleCreated()" />
         <FormKit label="Image or page numbers" type="text" v-model="item.pages" id="pages"/>
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
         <span class="btn-group">
            <uva-button @click="addClicked()" class="pad-left">Add Item</uva-button>
            <uva-button @click="completeClicked(handlers, node.context)" class="pad-left" data-next="true">Complete Order</uva-button>
         </span>
      </template>
   </FormKit>
</template>

<script setup>
import { useOrderStore } from '@/stores/order'
import { computed, ref } from 'vue'
import { getNode } from '@formkit/core'

const orderStore = useOrderStore()
const hasErrors = ref(false)
const showMessage = ref(false)

const item = computed(() => {
   return orderStore.items[orderStore.currItemIdx]
})

const toggleAddMessage = (()=>{
   showMessage.value = !showMessage.value
   if ( showMessage.value) {
      setTimeout( () => {document.getElementById("close-msg").focus()}, 300)
      setTimeout( () => {
         showMessage.value = false
         document.getElementById("title").focus()
      }, 3000)
   } else {
      document.getElementById("title").focus()
   }
})

const titleCreated = (() => {
   hasErrors.value = false
   setTimeout( () => {
      let titleInput = document.getElementById("title")
      titleInput.focus()
   }, 250)
})

const deleteItem = (() => {
   orderStore.removeItem(orderStore.currItemIdx)
})

const prevItem = (() => {
   if ( orderStore.currItemIdx > 0) {
      orderStore.currItemIdx--
   }
})

const nextItem = (() => {
   if ( orderStore.currItemIdx < orderStore.items.length) {
      orderStore.currItemIdx++
   }
})

const validateForm = (() => {
   hasErrors.value = false
   if (item.value.title == "") {
      let titleNode = getNode('title')
      titleNode.setErrors("A title is required")
      document.getElementById("title").focus()
      hasErrors.value = true
   }
   if (item.value.pages == "") {
      let pagesNode = getNode('pages')
      pagesNode.setErrors("Image or page numbers are required")
      if (  hasErrors.value == false) {
         document.getElementById("pages").focus()
         hasErrors.value = true
      }
   }
})

const addClicked = ( () => {
   validateForm()
   if ( hasErrors.value == false) {
      orderStore.addItem()
      toggleAddMessage()
   }
})

const completeClicked = ( (handlers, context) => {
   validateForm()
   if ( hasErrors.value == false) {
      handlers.incrementStep(1, context)()
   }
})
</script>

<style scoped lang="scss">
div.added-overlay {
   position: fixed;
   left: 0;
   top: 0;
   width: 100%;
   height: 100%;
   z-index: 1000;
   background: rgba(100,100,100,0.3);
   .message {
      box-shadow: 0 3px 10px rgba(0, 0, 0, 0.16), 0 3px 10px rgba(0, 0, 0, 0.23);
      width: 25%;
      margin: 100px auto;
      background: white;
      border-radius: 5px;
   }
   .msg-title {
      background: var(--uvalib-teal-lightest);
      font-weight: bold;
      padding: 5px;
      border-radius: 5px 5px 0 0;
      border-bottom: 1px solid var(--uvalib-teal);
   }
   .body {
      padding: 25px 10px 10px 10px;
   }
   .bar {
      text-align: right;
      padding: 5px 10px 10px 10px;
      button.uva-button {
         font-size: 0.8em;
         background-color: #eee;
         border-color: #666;
         color: #666;
         &:focus {
            outline: 3px dotted #999;
            outline-offset: 3px;
         }
      }
   }
}
.v-enter-active,
.v-leave-active {
  transition: opacity 0.3s ease;
}
.v-enter-from,
.v-leave-to {
  opacity: 0;
}

.item-label {
   text-align: left;
   font-weight: bold;
   background: var(--uvalib-grey-lightest);
   padding: 10px;
   border-radius: 5px;
   display: flex;
   flex-flow: row nowrap;
   justify-content: space-between;
   align-items: center;
}
.btn-group {
   display: flex;
   flex-flow: row nowrap;
   gap: 5px;
   padding: 0;
}
.item {
   display: flex;
   flex-direction: column;
   gap: 20px;
}
</style>
