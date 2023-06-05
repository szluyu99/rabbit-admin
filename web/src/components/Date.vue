<script setup lang="ts">
import { useDateFormat } from '@vueuse/core'

defineProps({
  value: {
    type: String,
    required: true,
  },
  format: {
    type: String,
    default: 'YYYY-MM-DD',
  },
})
</script>

<template>
  <div class="flex items-center gap-1.5">
    <span class="i-heroicons:clock text-xl text-gray-400" />
    {{ useDateFormat(value, format).value }}
  </div>
</template>
