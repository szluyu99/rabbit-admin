<script setup lang="ts">
import { reactive } from 'vue'
import { useRouter } from 'vue-router'
import api from '@/api'

const router = useRouter()

const loginForm = reactive({
  email: '',
  password: '',
})

async function doSignin(e: Event) {
  e.preventDefault()

  try {
    await api.login(loginForm)
    router.push('/')
  }
  catch (e) {
    console.error(e)
  }
}
</script>

<template>
  <div class="flex min-h-full flex-col justify-center py-16 sm:px-6 lg:px-8">
    <div class="sm:mx-auto sm:w-full sm:max-w-md">
      <img class="mx-auto h-12 w-auto" src="https://tailwindui.com/img/logos/mark.svg?color=indigo&shade=600" alt="Your Company">
      <h2 class="mt-6 text-center text-3xl font-bold tracking-tight text-gray-900">
        Sign in to your account
      </h2>
      <div class="mt-2 text-center text-sm text-gray-600">
        Or
        {{ ' ' }}
        <div class="font-medium text-indigo-600 hover:text-indigo-500 cursor-pointer" @click="$router.push('/signup')">
          register a new account
        </div>
      </div>
    </div>

    <div class="mt-8 sm:mx-auto sm:w-full sm:max-w-md">
      <div class="bg-white px-4 py-8 shadow sm:rounded-lg sm:px-10">
        <form class="space-y-6" @submit="doSignin">
          <div>
            <label for="email" class="block text-sm font-medium leading-6 text-gray-900">
              Email address
            </label>
            <div class="mt-2">
              <input
                v-model="loginForm.email"
                type="email" autocomplete="email" required
                class="block w-full rounded-md border-0 py-1.5 text-gray-900 shadow-sm ring-1 ring-inset ring-gray-300 placeholder:text-gray-400 focus:ring-2 focus:ring-inset focus:ring-indigo-600 sm:text-sm sm:leading-6"
              >
            </div>
          </div>

          <div>
            <label for="password" class="block text-sm font-medium leading-6 text-gray-900">
              Password
            </label>
            <div class="mt-2">
              <input
                v-model="loginForm.password"
                type="password" autocomplete="current-password" required
                class="block w-full rounded-md border-0 py-1.5 text-gray-900 shadow-sm ring-1 ring-inset ring-gray-300 placeholder:text-gray-400 focus:ring-2 focus:ring-inset focus:ring-indigo-600 sm:text-sm sm:leading-6"
              >
            </div>
          </div>

          <div class="flex items-center justify-between">
            <div class="flex items-center">
              <input id="remember-me" name="remember-me" type="checkbox" class="h-4 w-4 rounded border-gray-300 text-indigo-600 focus:ring-indigo-600">
              <label for="remember-me" class="ml-2 block text-sm text-gray-900">
                Remember me
              </label>
            </div>

            <div class="text-sm">
              <a href="#" class="font-medium text-indigo-600 hover:text-indigo-500">
                Forgot your password?
              </a>
            </div>
          </div>

          <div>
            <button
              type="submit"
              class="flex w-full justify-center rounded-md bg-indigo-600 px-3 py-2 text-sm font-semibold text-white shadow-sm hover:bg-indigo-500 focus-visible:outline focus-visible:outline-2 focus-visible:outline-offset-2 focus-visible:outline-indigo-600"
            >
              Sign in
            </button>
          </div>
        </form>
      </div>
    </div>
  </div>
</template>