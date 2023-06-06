import { createApp } from 'vue'
import { UToast } from 'unocss-ui'

type ToastCtx = InstanceType<typeof UToast>
type ToastProps = ToastCtx['$props']

// Create a mount node for the toast component
const mountNode = document.createElement('div')
document.body.appendChild(mountNode)

let instance = createToast({
  position: 'top',
  align: 'center',
  closeable: true,
})
let mounted: ToastCtx = instance.mount(mountNode) as ToastCtx

// Cache the toast position and align
// When the props change, unmount the current toast instance
const _cache: ToastProps = {
  position: 'top',
  align: 'center',
  closeable: true,
}

export function useToast(options: ToastProps = {
  position: 'top',
  align: 'center',
  closeable: true,
}) {
  if (!compareObjects(_cache, options)) {
    _cache.position = options.position
    _cache.align = options.align

    instance.unmount()
    instance = createToast(options)
    mounted = instance.mount(mountNode) as ToastCtx
  }

  return {
    success: (msg: string) => mounted.success(msg),
    info: (msg: string) => mounted.info(msg),
    warning: (msg: string) => mounted.warning(msg),
    error: (msg: string) => mounted.error(msg),
  }
}

function createToast(o: ToastProps) {
  return createApp(UToast, o)
}

function compareObjects(o1: Record<string, unknown>, o2: Record<string, unknown>): boolean {
  return ['position', 'align', 'closeable'].every(key => o1[key] === o2[key])
}

export const toast = useToast()
