<script lang="ts" setup>
const currentYear = new Date().getFullYear()
</script>

<template>
  <div class="w-full items-center justify-between border-t-2 px-8 py-3 lg:flex">
    <p class="mb-2 lg:mb-0">
      © Copyright {{ currentYear }}
    </p>

    <div class="flex">
      <a href="https://github.com/cherryful" target="_blank" class="mr-6 hover:text-gray-900">Our Organization</a>
      <a href="https://github.com/szluyu99" target="_blank" class="hover:text-gray-900">About Me</a>
    </div>
  </div>
</template>
