import {
  defineConfig,
  presetIcons,
  presetUno,
  presetWebFonts,
} from 'unocss'
import { presetForms } from '@julr/unocss-preset-forms'

export default defineConfig({
  shortcuts: [
    ['f-c-c', 'flex items-center justify-center'],
    ['wh-full', 'w-full h-full'],
    ['btn', `px-4 py-1 rounded inline-block bg-teal-600 text-white cursor-pointer 
    hover:bg-teal-700 disabled:cursor-default disabled:bg-gray-600 disabled:opacity-50`],
  ],
  presets: [
    presetUno(),
    presetForms(),
    // presetAttributify(),
    presetIcons({
      // scale: 1.2,
      warn: true,
      extraProperties: {
        cursor: 'pointer',
        // display: 'inline-block',
      },
    }),
    presetWebFonts({
      fonts: {
        sans: 'DM Sans',
        serif: 'DM Serif Display',
        mono: 'DM Mono',
      },
    }),
  ],
})
