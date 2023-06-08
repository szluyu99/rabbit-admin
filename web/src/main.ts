import '@unocss/reset/tailwind.css'
import 'unocss-ui/style.css'
import 'uno.css'
import './style.scss'

import { createApp } from 'vue'
import App from './App.vue'
import store from './store'
import router from './router'

const app = createApp(App)
app.use(store)
app.use(router)
app.mount('#app')
