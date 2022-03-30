import { createApp } from 'vue'
import App from './App.vue'
import './assets/tailwind.css'
import '@fortawesome/fontawesome-free/css/all.css'
import '@fortawesome/fontawesome-free/js/all.js'
import axios from 'axios'
import VueAxios from 'vue-axios'
import router from './router'

let app = createApp(App)

app.config.globalProperties.myKey = '49b11a04577b4b818c848eef087c4315';

app.use(router)
.use(VueAxios, axios)
.mount('#app')

app.provide('myKey', '49b11a04577b4b818c848eef087c4315')