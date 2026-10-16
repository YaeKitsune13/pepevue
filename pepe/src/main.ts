import { createApp } from 'vue'
import { createPinia } from 'pinia'

import App from './App.vue'
import router from './router/index'

// Подключение CSS из public
import './styles/main.css'
import './styles/header.css'
import './styles/body.css'
import './styles/buyhints.css'
import './styles/forms.css'
import './styles/profile.css'

const app = createApp(App)
const pinia = createPinia()

app.use(pinia)
app.use(router)

app.mount('#app')
