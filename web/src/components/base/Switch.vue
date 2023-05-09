<script setup lang="ts">
import { computed } from 'vue'
import { uniqueId } from 'lodash-es'

const props = withDefaults(defineProps<{
  type?: 'success' | 'info' | 'warning' | 'error' | 'primary' | 'secondary' | 'accent'
  size?: 'sm' | 'base' | 'lg'
  modelValue?: boolean | string | number
  checked?: boolean
  disabled?: boolean
  icon?: boolean
  readonly?: boolean
  checkedValue?: any
  uncheckedValue?: any
}>(), {
  type: 'primary',
  modelValue: false,
  size: 'base',
  disabled: false,
  icon: false,
  readonly: false,
  checkedValue: undefined,
  uncheckedValue: undefined,
})

const emits = defineEmits<{
  (evt: 'update:modelValue', val: any): void
}>()

const uid = uniqueId('switch-')

const switchValue = computed({
  get: () => {
    if (props.checkedValue && props.modelValue === props.checkedValue)
      return true
    if (props.uncheckedValue && props.modelValue === props.uncheckedValue)
      return false
    return props.modelValue
  },
  set: (val) => {
    if (props.checkedValue && val)
      val = props.checkedValue
    if (props.uncheckedValue && !val)
      val = props.uncheckedValue
    emits('update:modelValue', val)
  },
})

const themes = {
  success: {
    color: 'bg-success-500',
    text: 'text-success-500',
    focus: 'focus:ring-success-500',
  },
  info: {
    color: 'bg-info-500',
    text: 'text-info-500',
    focus: 'focus:ring-info-500',
  },
  warning: {
    color: 'bg-warning-500',
    text: 'text-warning-500',
    focus: 'focus:ring-warning-500',
  },
  error: {
    color: 'bg-error-500',
    text: 'text-error-500',
    focus: 'focus:ring-error-500',
  },
  primary: {
    color: 'bg-primary-500',
    text: 'text-primary-500',
    focus: 'focus:ring-primary-500',
  },
  secondary: {
    color: 'bg-secondary-500',
    text: 'text-secondary-500',
    focus: 'focus:ring-secondary-500',
  },
  accent: {
    color: 'bg-accent-500',
    text: 'text-accent-500',
    focus: 'focus:ring-accent-500',
  },
}
</script>

<script lang="ts">
export default {
  inheritAttrs: false,
}
</script>

<template>
  <div class="flex items-center">
    <label
      :for="uid"
      class="flex items-center"
      :class="[disabled ? 'cursor-not-allowed opacity-60' : readonly ? '' : 'cursor-pointer']"
    >
      <div class="mr-2 text-sm font-medium leading-6 text-gray-900 empty:hidden">
        <slot />
      </div>
      <div
        tabindex="0"
        class="relative focus:rounded-full focus:outline-none focus:ring-2 focus:ring-offset-2"
        :class="[themes[type].focus]"
      >
        <input
          :id="uid"
          v-model="switchValue"
          v-bind="$attrs"
          type="checkbox"
          :disabled="disabled"
          class="sr-only"
        >
        <div
          class="block rounded-full"
          :class="[
            {
              'w-9 h-4': size === 'sm',
              'w-11 h-6': size === 'base',
              'w-14 h-9': size === 'lg',
            },
            (switchValue || checked) ? themes[type].color : 'bg-gray-200',
          ]"
        />

        <div
          class="absolute left-0 top-0.5 rounded-full bg-white transition"
          :class="[
            {
              'w-3 h-3': size === 'sm',
              'h-5 w-5': size === 'base',
              'w-8 h-8': size === 'lg',
            },
            (switchValue || checked) ? 'translate-x-5' : 'translate-x-0',
          ]"
        >
          <svg
            v-if="icon"
            fill="currentColor" viewBox="0 0 12 12"
            :class="[
              (checked || switchValue) ? themes[type].text : 'text-gray-200',
              {
                'w-3 h-3': size === 'sm',
                'h-5 w-5': size === 'base',
                'w-8 h-8': size === 'lg',
              }]"
          >
            <path v-if="checked || switchValue" d="M3.707 5.293a1 1 0 00-1.414 1.414l1.414-1.414zM5 8l-.707.707a1 1 0 001.414 0L5 8zm4.707-3.293a1 1 0 00-1.414-1.414l1.414 1.414zm-7.414 2l2 2 1.414-1.414-2-2-1.414 1.414zm3.414 2l4-4-1.414-1.414-4 4 1.414 1.414z" />
            <path v-else d="M4 8l2-2m0 0l2-2M6 6L4 4m2 2l2 2" stroke="currentColor" stroke-width="2" stroke-linecap="round" stroke-linejoin="round" />
          </svg>
        </div>
      </div>
    </label>
  </div>
</template>
