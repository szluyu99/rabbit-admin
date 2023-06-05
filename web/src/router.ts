import type { RouteRecordRaw } from 'vue-router'
import { createRouter, createWebHistory } from 'vue-router'

import api from './api'
import { useUserStore } from './store'
import { getToken } from './utils/token'

const Layout = () => import('@/layout/Layout.vue')

// TODO: backend routes
export const layoutDynamicRoutes: Array<RouteRecordRaw> = [
  {
    name: 'Home',
    path: '/',
    redirect: '/home',
    children: [
      {
        name: 'Home', path: 'home', component: () => import('@/views/home/Home.vue'), meta: { icon: 'i-mdi:lightning-bolt' },
      },
    ],
  },
  {
    name: 'Article',
    path: '/article',
    redirect: '/article/list',
    children: [
      { name: 'Write', path: 'write/:articleId?', component: () => import('@/views/article/Write.vue'), meta: { icon: 'i-mdi:pen', hidden: true } },
      { name: 'Article', path: 'list', component: () => import('@/views/article/ArticleList.vue'), meta: { icon: 'i-mdi:blogger' } },
      { name: 'Category', path: 'category', component: () => import('@/views/article/Category.vue'), meta: { icon: 'i-mdi:menu' } },
      { name: 'Tag', path: 'tag', component: () => import('@/views/article/Tag.vue'), meta: { icon: 'i-mdi:tag' } },
    ],
  },
  {
    name: 'Auth',
    path: '/power',
    redirect: '/power/role',
    children: [
      { name: 'Role', path: 'role', component: () => import('@/views/auth/Role.vue'), meta: { icon: 'i-mdi:power' } },
      { name: 'Permission', path: 'permission', component: () => import('@/views/auth/Permission.vue'), meta: { icon: 'i-mdi:ab-testing' } },
      { name: 'Group', path: 'group', component: () => import('@/views/auth/Group.vue'), meta: { icon: 'i-mdi:group', hidden: true } },
      { name: 'User', path: 'user', component: () => import('@/views/auth/User.vue'), meta: { icon: 'i-mdi:account' } },
    ],
  },
  {
    name: 'System',
    path: '/system',
    redirect: '/system/config',
    children: [
      { name: 'Config', path: 'config', component: () => import('@/views/system/Config.vue'), meta: { icon: 'i-mdi:cog-transfer' } },
    ],
  },
]

const routes: Array<RouteRecordRaw> = [
  {
    name: 'General',
    path: '/',
    redirect: '/home',
    component: Layout,
    children: layoutDynamicRoutes,
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
  if (to.path === '/signin' || to.path === '/signup')
    return true

  // 1. default use session
  // 2. if session invalid, try to use token
  try {
    const userInfo = await api.userInfo()

    const userStore = useUserStore()
    userStore.signin(userInfo)

    if (to.path === '/signin')
      return { path: '/' }
    // TODO: refreshToken
    return true
  }
  catch {
    try {
      const token = getToken()
      if (token)
        await api.login({ token })
      else return { path: '/signin' }
    }
    catch {
      return { path: '/signin' }
    }
  }
})

export default router
