<script setup>
import { onMounted, ref } from 'vue'

import ModalDemo from './ModalDemo.vue'
import ButtonDemo from '@/views/ButtonDemo.vue'

import { useStore } from '@/store'
import useMouseClick from '@/hooks/useMouse'
import api from '@/api'

const store = useStore()

const mouse = useMouseClick()
const sentence = ref('君子不器。')

onMounted(() => {
  fetchSentence()
})

async function fetchSentence() {
  const resp = await api.oneSentence()
  sentence.value = resp.hitokoto
}
</script>

<template>
  <div class="space-y-5 m-5">
    <div class="text-3xl i-twemoji-grinning-face-with-smiling-eyes hover:i-twemoji-face-with-tears-of-joy" />

    <div>
      <button class="btn bg-amber-500 hover:bg-amber-600 mr-5" @click="fetchSentence">
        Fetch API
      </button>
      {{ sentence }}
    </div>

    <div>
      <div class="btn bg-blue-500 hover:bg-blue-600 mr-5" @click="store.increment">
        Increase Store Count
      </div>
      <span class="text-xl font-bold text-green-400">
        {{ store.count }}
      </span>
    </div>

    <div>
      <RouterLink to="/hello">
        <div class="btn bg-red-500 hover:bg-red-600">
          To Hello Page
        </div>
      </RouterLink>
    </div>

    <div class="inline-block border p-4">
      useMouseClick hook
      <p>
        Mouse click x:
        <span class="text-lg text-red-500 font-bold">
          {{ mouse.x }}
        </span>
      </p>
      <p>
        Mouse click y:
        <span class="text-lg text-blue-500 font-bold">
          {{ mouse.y }}
        </span>
      </p>
    </div>

    <ButtonDemo />
    <ModalDemo />
  </div>
</template>
