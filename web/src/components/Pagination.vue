<script setup lang="ts">
import { ref } from 'vue'
import Select from '@/components/base/Select.vue'

const props = defineProps({
  total: { type: Number, default: 0 },
  pos: { type: Number, default: 0 },
  limit: { type: Number, default: 10 },
})

const emit = defineEmits(['query', 'update:pos', 'update:limit'])

const pos = ref(props.pos)
const limit = ref(props.limit)

function handlePrev() {
  if (pos.value === 0)
    return
  pos.value -= limit.value
  emit('update:pos', pos.value)
  emit('query')
}

function handleNext() {
  if (pos.value + limit.value >= props.total)
    return
  pos.value += limit.value
  emit('update:pos', pos.value)
  emit('query')
}

function updateLimit() {
  pos.value = 0

  emit('update:limit', limit.value)
  emit('update:pos', pos.value)
  emit('query')
}

const options = [
  { label: '5', value: 5 },
  { label: '10', value: 10 },
  { label: '20', value: 20 },
  { label: '50', value: 50 },
  { label: '100', value: 100 },
]
</script>

<template>
  <div class="flex items-center justify-between">
    <div class="flex items-center">
      <span class="font-medium mr-1">Total:</span>
      <span class="text-gray-600"> {{ total }}</span>
      <span class="text-gray-700 ml-4 hidden xs:block"> View </span>
      <div class="w-20 ml-2 ">
        <Select
          v-model="limit"
          :options="options"
          :placeholder="String(limit)"
          height="py-1.5"
          @change="updateLimit"
        />
      </div>
    </div>
    <nav
      class="flex items-center justify-between py-3"
      aria-label="Pagination"
    >
      <div class="flex flex-1 justify-between space-x-3 sm:justify-end">
        <button
          class="relative inline-flex items-center px-2 py-1 text-sm font-medium text-gray-700
          rounded-md border border-gray-300 bg-white hover:bg-gray-50 p-1 hover:text-gray-500
          focus:outline-none focus:ring-2 focus:ring-indigo-500 focus:ring-offset-2"
          @click="handlePrev"
        >
          <span class="i-mdi:chevron-left w-5 h-5" />
        </button>
        <button
          class="relative inline-flex items-center px-2 py-1 text-sm font-medium text-gray-700
          rounded-md border border-gray-300 bg-white hover:bg-gray-50 p-1 hover:text-gray-500
          focus:outline-none focus:ring-2 focus:ring-indigo-500 focus:ring-offset-2"
          @click="handleNext"
        >
          <span class="i-mdi:chevron-right w-5 h-5" />
        </button>
      </div>
    </nav>
  </div>
</template>
