import { createRouter, createWebHistory } from 'vue-router'
import Index from '../views/Index.vue'
import Login from '../views/Login.vue'

const routes = [
  {
    path: "/",
    component: Login
  },
  {
    path: "/index",
    component: Index
  },
  {
    path: "/cart",
    component: () => import("../views/Cart.vue")
  },
  {
    path: '/login',
    component: Login
  },
  {
    path: "/register",
    component: () => import("../views/Register.vue")
  },

]

const router = createRouter({
  history: createWebHistory(),
  routes
})

export default router
