import { createPinia, defineStore } from 'pinia'

const pinia = createPinia()
export default pinia

interface User {
  id: number
  email: string
  roles: []
  enabled: boolean
  activated: boolean
  timezone: string
  isSuper: boolean
}

export const useUserStore = defineStore('user', {
  state: () => ({
    isAuthenticated: false,
    user: {} as User,
  }),
  getters: {
    roles: state => state.user.roles.map((role: any) => role.name),
  },
  actions: {
    signin(user: User) {
      this.isAuthenticated = true
      this.user = user
    },
    signout() {
      this.isAuthenticated = false
      this.user = {} as User
      localStorage.clear()
    },
  },
})
