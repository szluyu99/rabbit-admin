<script setup lang="ts">
import { ref } from 'vue'
import Modal from '@/components/base/Modal.vue'

const show = ref(false)
</script>

<template>
  <button class="btn" @click="show = true">
    TailwindUI Modal
  </button>
  <Modal v-model:show="show">
    <template #header>
      Header
    </template>
    Hello World!
    <template #footer>
      <div class="mt-5 sm:mt-4 sm:flex sm:flex-row-reverse">
        <button
          type="button"
          class="inline-flex w-full justify-center rounded-md bg-red-600 px-3 py-2 text-sm font-semibold text-white shadow-sm hover:bg-red-500 focus:outline-none focus:ring-2 focus:ring-indigo-500 focus:ring-offset-2 sm:ml-3 sm:w-auto"
          @click="show = false"
        >
          Deactivate
        </button>
        <button
          class="mt-3 inline-flex w-full justify-center rounded-md bg-white px-3 py-2 text-sm font-semibold text-gray-900 shadow-sm ring-1 ring-gray-300 hover:bg-gray-50  focus:outline-none focus:ring-2 focus:ring-indigo-500 focus:ring-offset-2 sm:mt-0 sm:w-auto"
          @click="show = false"
        >
          Cancel
        </button>
      </div>
    </template>
  </Modal>
</template>
