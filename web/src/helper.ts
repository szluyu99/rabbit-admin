import { useDateFormat } from '@vueuse/core'

export function formatDate(s: string, format = 'YYYY-MM-DD') {
  return useDateFormat(s, format).value
}
