import type { RouteRecordRaw } from 'vue-router'
import { createRouter, createWebHistory } from 'vue-router'

import api from './api'

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
        path: '/article',
        redirect: '/article/list',
        children: [
          { path: 'write', component: () => import('@/views/article/Write.vue') },
          { path: 'list', component: () => import('@/views/article/ArticleList.vue') },
          { path: 'category', component: () => import('@/views/article/Category.vue') },
          { path: 'tag', component: () => import('@/views/article/Tag.vue') },
        ],
      },
      {
        path: '/power',
        redirect: '/power/role',
        children: [
          { path: 'role', component: () => import('@/views/auth/Role.vue') },
          { path: 'permission', component: () => import('@/views/auth/Permission.vue') },
          { path: 'group', component: () => import('@/views/auth/Group.vue') },
          { path: 'user', component: () => import('@/views/auth/User.vue') },
        ],
      },
    ],
  },
  { path: '/signin', component: () => import('@/views/Signin.vue') },
  { path: '/signup', component: () => import('@/views/Signup.vue') },
  { path: '/logout', component: () => import('@/views/Logout.vue') },
  { path: '/:path(.*)', component: () => import('@/views/NotFound.vue') },
]

const router = createRouter({
  history: createWebHistory(import.meta.env.VITE_BASE_PUBLIC_PATH),
  routes,
})

// login guard
router.beforeEach(async (to) => {
  if (to.path === '/signin')
    return true

  try {
    await api.userInfo()
    if (to.path === '/signin')
      return { path: '/' }
    // TODO: refreshToken
    return true
  }
  catch {
    return { path: 'signin' }
  }
})

export default router
