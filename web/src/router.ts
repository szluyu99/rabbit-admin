import type { RouteRecordRaw } from 'vue-router'
import { createRouter, createWebHistory } from 'vue-router'

const Layout = () => import('@/layout/Layout.vue')

const routes: Array<RouteRecordRaw> = [
  {
    name: '',
    path: '/',
    redirect: '/home',
    component: Layout,
    children: [
      { path: '/home', component: () => import('@/views/home/Home.vue') },
      {
        path: '/power',
        redirect: '/power/role',
        children: [
          { path: 'role', component: () => import('@/views/auth/Role.vue') },
          { path: 'permission', component: () => import('@/views/auth/Permission.vue') },
        ],
      },
    ],
  },
  { path: '/signin', component: () => import('@/views/Signin.vue') },
  { path: '/signup', component: () => import('@/views/Signup.vue') },
  { path: '/:path(.*)', component: () => import('@/views/NotFound.vue') },
]

const router = createRouter({
  history: createWebHistory(),
  routes,
})

export default router
