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
  <div class="flex min-h-full flex-col justify-center mt-10 py-20 sm:px-6 lg:px-8">
    <div class="sm:mx-auto sm:w-full sm:max-w-md">
      <h2 class="mt-6 text-center text-3xl font-bold tracking-tight text-gray-900">
        Free Sign Up
      </h2>
      <p class="mt-2 text-center text-sm text-gray-600">
        Create your account, it takes less than a mintue
      </p>
    </div>

    <div class="mt-8 sm:mx-auto sm:w-full sm:max-w-md">
      <div class="bg-white py-8 px-4 shadow sm:rounded-lg sm:px-10">
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
                  class="block w-full appearance-none rounded-md border border-gray-300 px-3 py-1 placeholder-gray-400 shadow-sm focus:border-indigo-500 focus:outline-none focus:ring-indigo-500 sm:text-sm"
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
                  class="block w-full appearance-none rounded-md border border-gray-300 px-3 py-1 placeholder-gray-400 shadow-sm focus:border-indigo-500 focus:outline-none focus:ring-indigo-500 sm:text-sm"
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
                class="block w-full appearance-none rounded-md border border-gray-300 px-3 py-1 placeholder-gray-400 shadow-sm focus:border-indigo-500 focus:outline-none focus:ring-indigo-500 sm:text-sm"
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
                class="block w-full appearance-none rounded-md border border-gray-300 px-3 py-1 placeholder-gray-400 shadow-sm focus:border-indigo-500 focus:outline-none focus:ring-indigo-500 sm:text-sm"
              >
            </div>
          </div>

          <button
            type="submit"
            class="flex w-full justify-center rounded-md border cursor-pointer border-transparent bg-indigo-600 py-2 px-4 text-sm font-medium text-white shadow-sm hover:bg-indigo-700 focus:outline-none focus:ring-2 focus:ring-indigo-500 focus:ring-offset-2"
          >
            Sign up
          </button>
        </form>
      </div>
    </div>

    <div class="text-gray-400 text-sm mt-6 text-center">
      Already have account?
      <div
        class="ml-1 underline cursor-pointer"
        @click="$router.push('/signin')"
      >
        Sign In
      </div>
    </div>
  </div>
</template>
