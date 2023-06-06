<script setup lang="ts">
import { reactive } from 'vue'
import { useRouter } from 'vue-router'
import api from '@/api'

const router = useRouter()

const signupForm = reactive({
  firstname: '',
  lastname: '',
  email: '',
  password: '',
})

async function doSignup(e: Event) {
  e.preventDefault()

  if (signupForm.password.length <= 5) {
    alert('Your password must be more than 5 characters.')
    return
  }

  try {
    await api.register({
      email: signupForm.email,
      password: signupForm.password,
      locale: navigator.language,
      timezone: Intl.DateTimeFormat().resolvedOptions().timeZone,
    })
    router.push('/')
  }
  catch (e) {
    console.error(e)
  }
}
</script>

<template>
  <div class="mt-10 min-h-full flex flex-col justify-center py-20 lg:px-8 sm:px-6">
    <div class="sm:mx-auto sm:max-w-md sm:w-full">
      <h2 class="mt-6 text-center text-3xl font-bold tracking-tight text-gray-900">
        Free Sign Up
      </h2>
      <p class="mt-2 text-center text-sm text-gray-600">
        Create your account, it takes less than a minute
      </p>
    </div>

    <div class="mt-8 sm:mx-auto sm:max-w-md sm:w-full">
      <div class="bg-white px-4 py-8 shadow sm:rounded-lg sm:px-10">
        <form class="space-y-5" method="POST" @submit="doSignup">
          <div class="flex">
            <div>
              <label for="first-name" class="block text-sm font-medium text-gray-700">
                First Name
              </label>
              <div class="mt-1">
                <input
                  id="first-name"
                  v-model="signupForm.firstname"
                  name="first-name" type="text" autocomplete="first-name"
                  class="block w-full appearance-none border border-gray-300 rounded-md px-3 py-1 shadow-sm focus:border-indigo-500 sm:text-sm focus:outline-none focus:ring-indigo-500 placeholder-gray-400"
                >
              </div>
            </div>

            <div class="ml-6">
              <label for="last-name" class="block text-sm font-medium text-gray-700">
                Last Name
              </label>
              <div class="mt-1">
                <input
                  id="last-name"
                  v-model="signupForm.lastname"
                  name="last-name" type="text" autocomplete="last-name"
                  class="block w-full appearance-none border border-gray-300 rounded-md px-3 py-1 shadow-sm focus:border-indigo-500 sm:text-sm focus:outline-none focus:ring-indigo-500 placeholder-gray-400"
                >
              </div>
            </div>
          </div>

          <div>
            <label for="email" class="block text-sm font-medium text-gray-700">
              Email address
            </label>
            <div class="mt-1">
              <input
                id="email"
                v-model="signupForm.email" name="email" type="email" autocomplete="email" required
                class="block w-full appearance-none border border-gray-300 rounded-md px-3 py-1 shadow-sm focus:border-indigo-500 sm:text-sm focus:outline-none focus:ring-indigo-500 placeholder-gray-400"
              >
            </div>
          </div>

          <div>
            <label for="password" class="block text-sm font-medium text-gray-700">Password</label>
            <div class="mt-1">
              <input
                id="password"
                v-model="signupForm.password"
                name="password" type="password" required
                class="block w-full appearance-none border border-gray-300 rounded-md px-3 py-1 shadow-sm focus:border-indigo-500 sm:text-sm focus:outline-none focus:ring-indigo-500 placeholder-gray-400"
              >
            </div>
          </div>

          <button
            type="submit"
            class="w-full flex cursor-pointer justify-center border border-transparent rounded-md bg-indigo-600 px-4 py-2 text-sm font-medium text-white shadow-sm hover:bg-indigo-700 focus:outline-none focus:ring-2 focus:ring-offset-2 focus:ring-indigo-500"
          >
            Sign up
          </button>
        </form>
      </div>
    </div>

    <div class="mt-6 text-center text-sm text-gray-400">
      Already have account?
      <div
        class="ml-1 cursor-pointer underline"
        @click="$router.push('/signin')"
      >
        Sign In
      </div>
    </div>
  </div>
</template>
