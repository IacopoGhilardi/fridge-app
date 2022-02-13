import { createRouter, createWebHistory } from 'vue-router'
import Home from '../views/Home.vue'
import Fridge from '../views/Fridge.vue'

const routes = [
  {
    path: '/',
    name: 'Home',
    component: Home
  },
  {
    path: '/your-fridge',
    name: 'Fridge',
    component: Fridge,
  }
]

const router = createRouter({
  history: createWebHistory(process.env.BASE_URL),
  routes
})

export default router
