import { createRouter, createWebHistory } from 'vue-router'
import Home from '../views/Home.vue'
import Fridge from '../views/Fridge.vue'
import Register from '../views/auth/Register.vue'
import Login from '../views/auth/Login.vue'

const routes = [
  {
    path: '/',
    name: 'Home',
    component: Home
  },
  {
    path: '/register',
    name: 'Register',
    component: Register,
  },
  {
    path: '/login',
    name: 'Login',
    component: Login,
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
