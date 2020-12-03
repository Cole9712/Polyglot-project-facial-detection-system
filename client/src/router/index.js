import Vue from 'vue'
import VueRouter from 'vue-router'
import Home from '../views/Home.vue'

Vue.use(VueRouter)

const routes = [
  {
    path: '/',
    name: 'Home',
    component: Home
  },
  {
    path: '/UploadMine',
    name: 'UploadMine',
    // route level code-splitting
    // this generates a separate chunk (about.[hash].js) for this route
    // which is lazy-loaded when the route is visited.
    component: () => import(/* webpackChunkName: "about" */ '../views/UploadMine.vue')
  },
  {
    path: '/Swap1',
    name: 'Swapping',
    component: () => import(/* webpackChunkName: "about" */ '../views/Swap1.vue')
  }
]

const router = new VueRouter({
  routes
})

export default router
