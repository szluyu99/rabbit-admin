import { onBeforeMount, onUnmounted, reactive } from 'vue'

export default function useMouseClick() {
  const state = reactive({
    x: 0,
    y: 0,
  })

  function onClick(event: MouseEvent) {
    state.x = event.pageX
    state.y = event.pageY
  }

  onBeforeMount(() => {
    window.addEventListener('click', onClick)
  })

  onUnmounted(() => {
    window.removeEventListener('click', onClick)
  })

  return state
}
