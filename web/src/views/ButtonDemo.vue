<script setup lang="ts">
import Button from '@/components/base/Button.vue'
</script>

<template>
  <div class="space-y-4">
    Button Demo:
    <div class="space-x-3">
      <span class="inline-block w-20">
        size:
      </span>
      <Button size="sm">
        button
      </Button>
      <Button>
        button
      </Button>
      <Button size="lg">
        button
      </Button>
    </div>
    <div class="space-x-3">
      <span class="inline-block w-20">
        type:
      </span>
      <Button type="default">
        default
      </Button>
      <Button type="success">
        success
      </Button>
      <Button type="info">
        info
      </Button>
      <Button type="warning">
        warning
      </Button>
      <Button type="error">
        error
      </Button>
      <Button type="primary">
        primary
      </Button>
    </div>
    <div class="space-x-3">
      <span class="inline-block w-20">
        disabled:
      </span>
      <Button disabled>
        disabled
      </Button>
    </div>
    <div class="space-x-3 flex items-center">
      <span class="inline-block w-20">
        icon:
      </span>
      <Button>
        <div class="f-c-c">
          <span class="i-mdi:lightning-bolt text-4" />
          <span> Lightning </span>
        </div>
      </Button>
    </div>
    <div class="space-x-3">
      <span class="inline-block w-20">
        custom:
      </span>
      <Button
        type="empty"
        custom="bg-pink-500 text-white hover:bg-pink-700 focus:ring-pink-500 p-1"
      >
        Pink Button
      </Button>
      <Button
        type="empty" size="empty"
        custom="bg-orange-500 text-white hover:bg-orange-700 focus:ring-orange-500 border-rounded p-3"
      >
        Orange Huge Button
      </Button>
    </div>
  </div>
</template>
