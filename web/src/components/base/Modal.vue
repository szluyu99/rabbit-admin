<script setup lang="ts">
import { ref, watch } from 'vue'
import {
  Dialog,
  DialogPanel,
  DialogTitle,
  TransitionChild,
  TransitionRoot,
} from '@headlessui/vue'

import { XMarkIcon } from '@heroicons/vue/24/outline'

interface Props {
  show: boolean
  showClose?: boolean
  width?: string
}

const props = withDefaults(defineProps<Props>(), {
  showClose: true,
  width: 'w-xl',
})

const emits = defineEmits(['update:show'])

const open = ref(props.show)
watch(() => props.show, val => open.value = val)

function closeModal() {
  open.value = false
  emits('update:show', open.value)
}
</script>

<template>
  <TransitionRoot appear :show="open" as="template">
    <Dialog as="div" class="relative z-10" @close="closeModal">
      <TransitionChild
        as="template"
        enter="duration-300 ease-out" enter-from="opacity-0" enter-to="opacity-100"
        leave="duration-200 ease-in" leave-from="opacity-100" leave-to="opacity-0"
      >
        <div class="fixed inset-0 bg-black bg-opacity-25" />
      </TransitionChild>

      <div class="fixed inset-0 overflow-y-auto">
        <div class="flex min-h-full items-center justify-center p-4 text-center ">
          <TransitionChild
            as="template"
            enter="ease-out duration-300" enter-from="opacity-0 translate-y-4 sm:translate-y-0 sm:scale-95" enter-to="opacity-100 translate-y-0 sm:scale-100"
            leave="ease-in duration-200" leave-from="opacity-100 translate-y-0 sm:scale-100" leave-to="opacity-0 translate-y-4 sm:translate-y-0 sm:scale-95"
          >
            <DialogPanel
              class="relative px-4 pt-5 pb-4 transform overflow-hidden
              rounded-lg bg-white text-left shadow-xl transition-all
              sm:my-8 sm:p-6"
              :class="width"
            >
              <div v-if="showClose" class="absolute top-0 right-0 hidden pt-4 pr-4 sm:block">
                <button
                  type="button"
                  class="rounded-md bg-white text-gray-400
                  hover:text-gray-500 focus:outline-none focus:ring-2 focus:ring-indigo-500 focus:ring-offset-2"
                  @click="closeModal"
                >
                  <span class="sr-only">Close</span>
                  <XMarkIcon class="h-6 w-6" aria-hidden="true" />
                </button>
              </div>

              <template v-if="$slots?.header">
                <DialogTitle as="h3" class="mb-5 text-base font-semibold leading-6 text-gray-900">
                  <slot name="header" />
                </DialogTitle>
              </template>
              <slot />
              <template v-if="$slots?.footer">
                <!-- <div class="mt-5 sm:mt-6 sm:grid sm:grid-flow-row-dense sm:grid-cols-2 sm:gap-3"> -->
                <slot name="footer" />
                <!-- </div> -->
              </template>
            </DialogPanel>
          </TransitionChild>
        </div>
      </div>
    </Dialog>
  </TransitionRoot>
</template>
