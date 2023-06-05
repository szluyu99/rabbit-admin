import {
  defineConfig,
  presetIcons,
  presetTypography,
  presetUno,
  presetWebFonts,
  transformerDirectives,
  transformerVariantGroup,
} from 'unocss'

import { colors } from 'unocss/preset-mini'
import { presetUnocssUI } from 'unocss-ui'

export default defineConfig({
  safelist: ['i-mdi:pen', 'i-mdi:blogger', 'i-mdi:menu', 'i-mdi:tag', 'i-mdi:power', 'i-mdi:ab-testing', 'i-mdi:group', 'i-mdi:account', 'i-mdi:cog-transfer'],
  theme: {
    colors: {
      primary: colors.violet,
      secondary: colors.teal,
      accent: colors.pink,
      success: colors.green,
      info: colors.blue,
      warning: colors.yellow,
      error: colors.red,
    },
  },
  presets: [
    presetUno(),
    presetTypography(),
    presetIcons(),
    presetWebFonts({
      provider: 'bunny',
      fonts: { sans: 'Inter' },
    }),
    presetUnocssUI(),
  ],
  transformers: [
    transformerDirectives(),
    transformerVariantGroup(),
  ],
})
