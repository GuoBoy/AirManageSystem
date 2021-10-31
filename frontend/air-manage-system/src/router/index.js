import { createRouter, createWebHistory } from 'vue-router'
import store from '../store'
import Index from '../views/Index.vue'
import FlightManage from '../views/FlightManage.vue'
import Login from '../views/Login.vue'

const routes = [
  {
    path: '/',
    component: Login,
  },
  {
    path: "/index",
    component: Index,
    children: [{
      path: 'flightmanage',
      component: FlightManage,
    }, {
      path: 'usermanage',
      component: () => import("../views/UserManage.vue")
    }, {
      path: 'ordermanage',
      component: () => import("../views/OrderManage.vue")
    },]
  },
  {
    path: "/login",
    component: Login
  }, {
    path: "/register",
    component: () => import("../views/Register.vue")
  }
]

const router = createRouter({
  history: createWebHistory(),
  routes
})

function isAuth() {
  if (store.getters.getToken !== "") return true
  return false
}

router.beforeEach((to, from, next) => {
  if (to.path === "/login" || to.path === "/register") next()
  else if (!isAuth()) {
    console.log("没有登录");
    next("login")
  }
  else {
    switch (to.path) {
      case "":
        if (!isAuth()) {
          next("login")
        }
        break;
    }
    next()
  }
})

export default router
