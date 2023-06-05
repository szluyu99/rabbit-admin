<script setup lang="ts">
import api from '@/api'
import { useUserStore } from '@/store'

const userStore = useUserStore()

api.logout().then(() => {
  userStore.signout()
})
</script>

<template>
  <div class="h-screen flex items-center justify-center">
    <div class="sm:border-l sm:border-gray-200 sm:pl-6">
      <h1 class="text-4xl font-bold tracking-tight text-gray-900 sm:text-5xl">
        See You Again !
      </h1>
      <p class="mt-1 text-base text-gray-500">
        You are now successfully sign out.
      </p>
    </div>
    <div class="mt-10 flex space-x-3 sm:border-l sm:border-transparent sm:pl-6">
      <a
        href="/signin"
        class="inline-flex items-center border border-transparent rounded-md bg-indigo-600 px-4 py-2 text-sm font-medium text-white shadow-sm hover:bg-indigo-700 focus:outline-none focus:ring-2 focus:ring-offset-2 focus:ring-indigo-500"
      >
        Sign In
      </a>
    </div>
  </div>
</template>
