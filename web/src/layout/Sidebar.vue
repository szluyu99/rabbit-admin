<script setup lang="ts">
import { ref } from 'vue'
import { useRoute } from 'vue-router'
import { useSidebar } from '@/composables/useSidebar'
import { layoutDynamicRoutes } from '@/router'

interface MenuOption {
  name: string
  children: Array<{
    name: string
    href: string
    icon?: string
    hidden?: boolean
    externalLink?: boolean
  }>
  _hidden: boolean
}

const { isOpen } = useSidebar()
const route = useRoute()

// Dynamically loading routing to generate sidebar menu
const menus = ref<MenuOption[]>(layoutDynamicRoutes.map(e => ({
  name: e.name,
  children: e.children?.map(ee => ({
    name: ee.name,
    href: e.path.endsWith('/') ? e.path + ee.path : `${e.path}/${ee.path}`,
    icon: ee.meta?.icon || 'i-mdi:lightning-bolt',
    externalLink: ee.meta?.externalLink || false,
    hidden: ee.meta?.hidden || false,
  })),
  _hidden: e.meta?.hidden || false,
}) as MenuOption,
))
</script>

<template>
  <div
    class="fixed inset-0 z-20 bg-black opacity-50 transition-opacity lg:hidden"
    :class="isOpen ? 'block' : 'hidden'"
    @click="isOpen = false"
  />
  <div class="flex lg:w-64">
    <div
      class="fixed inset-y-0 left-0 z-30 w-64 transform overflow-y-auto bg-gray-50 transition duration-300 lg:static lg:inset-0 lg:translate-x-0"
      :class="isOpen ? 'translate-x-0 ease-out' : '-translate-x-full ease-in'"
    >
      <div class="mt-8 flex items-center justify-center">
        <div class="flex items-center">
          <img class="ml-2 h-10 w-10" src="@/assets/rabbit.svg" alt="rabbit">
          <span class="mx-2 text-2xl font-semibold text-primary-500">
            Rabbit Admin
          </span>
        </div>
      </div>
      <nav class="mt-10">
        <template v-for="menu of menus" :key="menu.name">
          <p class="my-2 mb-4 flex items-center justify-between pl-4 text-sm font-semibold text-gray-400">
            {{ menu.name }}
            <span
              class="i-mdi:lightning-bolt mr-2 h-5 w-5 cursor-pointer duration-200"
              :class="{ 'transform rotate-180': !menu._hidden }"
              @click="menu._hidden = !menu._hidden"
            />
          </p>
          <div
            class="duration-200 ease-in-out" :class="{
              'opacity-0 h-0': menu._hidden,
              'opacity-100 h-full': !menu._hidden,
            }"
          >
            <template v-for="child of menu.children" :key="child.name">
              <template v-if="!child.hidden">
                <a
                  v-if="child.externalLink"
                  :href="child.href" target="_blank"
                  class="inactive mt-4 flex items-center border-l-4 px-6 py-2 duration-200"
                >
                  <span v-if="child.icon" class="h-5 w-5" :class="child.icon" />
                  <span class="mx-4"> {{ child.name }} </span>
                </a>
                <RouterLink
                  v-else
                  class="mt-4 flex items-center border-l-4 px-6 py-2 duration-200"
                  :class="[route.path === child.href ? 'active' : 'inactive']"
                  :to="child.href"
                >
                  <span v-if="child.icon" class="h-5 w-5" :class="child.icon" />
                  <span class="mx-4"> {{ child.name }} </span>
                </RouterLink>
              </template>
            </template>
          </div>
        </template>
      </nav>
    </div>
  </div>
</template>

<style scoped>
.active {
  @apply bg-primary-500 text-gray-100 border-gray-600;
}
.inactive {
  @apply text-gray-700 hover:bg-primary-500 hover:text-gray-100 hover:border-gray-600;
}
</style>
