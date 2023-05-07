<script setup lang="ts">
import type { InputHTMLAttributes } from 'vue'
import { computed, reactive } from 'vue'

interface Props extends InputHTMLAttributes {
  modelValue?: string
  label?: string
  type?: string
  disabled?: boolean
  required?: boolean
  preIcon?: string
  suffIcon?: string
  errorMessage?: string
  placeholder?: string
}

const props = withDefaults(defineProps<Props>(), {
  label: '',
  modelValue: '',
  type: 'text',
  disabled: false,
  required: false,
  preIcon: '',
  suffIcon: '',
  errorMessage: '',
  placeholder: 'Please input',
})

const emit = defineEmits<{
  (evt: 'update:modelValue', val: string): void
  (evt: 'prepend'): void
  (evt: 'append'): void
}>()

const inputValue = computed({
  get: () => props.modelValue,
  set: val => emit('update:modelValue', val),
})

const flux = reactive({
  focused: false,
})
</script>

<template>
  <div class="flex flex-col w-full relative" :class="[disabled && 'opacity-60']">
    <label class="text-sm font-bold mb-2 empty:hidden">
      <template v-if="label">{{ label }}</template>
      <slot v-else />
      <span v-if="required" class="text-error-500">*</span>
    </label>

    <div class="flex w-full items-center">
      <div
        v-if="preIcon"
        class="p-2 border border-slate-400 border-r-0 rounded rounded-r-none bg-white dark:bg-slate-800 z-1"
        :class="{
          'important:border-red-500 important:ring-red-500 mb-1': errorMessage,
        }"
        @click.stop="emit('prepend')"
      >
        <div :class="preIcon" class="w-5 h-5" />
      </div>
      <span
        v-if="$slots.prepend && !preIcon"
        class="inline-flex items-center rounded-l-md border border-r-0 border-gray-300 px-3 py-2 text-gray-500 sm:text-sm"
      >
        <slot name="prepend" />
      </span>
      <!-- TODO: v-bind="$attrs" -->
      <input
        v-model="inputValue"
        :type="type"
        :disabled="disabled"
        :placeholder="placeholder"
        class="text-field-input"
        :class=" {
          danger: errorMessage,
          prepend: preIcon || $slots.prepend,
          append: suffIcon || $slots.append,
        }"
        @focus="flux.focused = true"
        @blur="flux.focused = false"
      >
      <div
        v-if="suffIcon"
        class="p-2 border border-slate-400 border-l-0 rounded rounded-l-none bg-white dark:bg-slate-800 z-1;"
        :class="{
          'important:border-red-500 important:ring-red-500 mb-1': errorMessage,
        }"
        @click.stop="emit('append')"
      >
        <div :class="suffIcon" class="w-5 h-5" />
      </div>
      <span
        v-if="$slots.append && !suffIcon"
        class="inline-flex items-center rounded-r-md border border-l-0 border-gray-300 px-3 py-2 text-gray-500 sm:text-sm"
      >
        <slot name="append" />
      </span>
    </div>

    <div v-if="errorMessage" class="text-error-500 text-xs">
      {{ errorMessage }}
    </div>
  </div>
</template>

<style lang="scss" scoped>
.text-field-input {
  @apply w-full border border-slate-400 rounded px-3 py-2 z-2;
  @apply bg-white dark:bg-slate-800 leading-tight;
  @apply focus:outline-none focus:ring-1 focus:ring-primary-400 focus:border-primary-400;
  @apply placeholder-gray-400 placeholder:text-sm;
  &.danger {
    @apply border-red-500 mb-1;
    @apply focus:ring-red-500 focus:border-red-500;
  }
  &.prepend {
    @apply rounded-l-none;
  }
  &.append {
    @apply rounded-r-none;
  }
}
</style>
