<script setup lang="ts">
defineProps<{
  title?: string
}>()
</script>

<template>
  <p class="text-2xl font-bold">
    {{ title }}
  </p>
  <p class="mt-3 mb-5">
    <slot name="description" />
  </p>
  <div class="bg-white rounded shadow p-6 max-w-6xl space-y-6">
    <slot />
  </div>
</template>
