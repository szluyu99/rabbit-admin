import request from './request'

export default {
  oneSentence: () => request.get('https://v1.hitokoto.cn?c=i'),

  // auth
  login: (data: any) => request.post('/auth/login', data),
  logout: () => request.get('/auth/logout'),
  register: (data: any) => request.post('/auth/register', data),
  userInfo: async () => await request.get('/auth/info'), // do auth check

  // permission
  initDefaultRoles: () => request.get('/api/role/default'),
  initDefaultPermission: () => request.get('/api/permission/default'),
  createWebobjectPermissions: (name: string) => request.get(`/api/permission/object/${name}`),

  // article
  getCategoryOptions: async () => {
    const result = await request.get('/api/category/all')
    return (result.items || []).map((item: any) => ({
      label: item.name,
      value: item.id,
    }))
  },
  getTagOptions: async () => {
    try {
      const result = await request.get('/api/tag/all')
      return (result.items || []).map((item: any) => ({
        label: item.name,
        value: item.id,
      }))
    }
    catch (err) {
      return []
    }
  },
  // permission options (tree structure)
  getPermissionOptions: async () => {
    try {
      const result = await request.post('/api/permission')
      return (result.items || []).map((item: any) => ({
        label: item.name,
        value: item.id,
        children: item.children.map((sub: any) => ({
          label: sub.name,
          value: sub.id,
        })),
      }))
    }
    catch (err) {
      return []
    }
  },
}
