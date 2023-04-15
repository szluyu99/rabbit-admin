<script setup lang="ts">
interface Props {
  type?: 'default' | 'success' | 'info' | 'warning' | 'error' | 'primary' | 'empty'
  size?: 'sm' | 'base' | 'lg' | 'empty'
  disabled?: boolean
  custom?: string
}

withDefaults(defineProps<Props>(), {
  type: 'default',
  size: 'base',
  disabled: false,
  custom: '',
})

const colors = {
  empty: '',
  default: 'text-gray-700 bg-white border border-gray-300 hover:bg-gray-50 focus:ring-indigo-500',
  success: 'text-white bg-green-500 hover:bg-green-700 focus:ring-green-500',
  info: 'text-white bg-blue-500 hover:bg-blue-700 focus:ring-blue-500',
  warning: 'text-white bg-yellow-500 hover:bg-yellow-700 focus:ring-yellow-500',
  error: 'text-white bg-red-500 hover:bg-red-700 focus:ring-red-500',
  primary: 'text-white bg-indigo-500 hover:bg-indigo-700 focus:ring-indigo-500',
}

const sizes = {
  empty: '',
  sm: 'px-2 py-1',
  base: 'px-3 py-2.5 sm:py-2',
  lg: 'px-10 py-4',
}
</script>

<template>
  <button
    :disabled="disabled"
    type="button"
    class="whitespace-nowrap inline-flex w-full items-center justify-center
    rounded-md shadow-sm text-xs font-sans font-semibold leading-4
    focus:outline-none focus:ring-2 focus:ring-offset-2
    disabled:cursor-not-allowed disabled:opacity-30
    sm:w-auto"
    :class="[colors[type], sizes[size], custom]"
  >
    <slot />
  </button>
</template>
