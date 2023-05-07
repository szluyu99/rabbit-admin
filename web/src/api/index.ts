import request from './request'

export default {
  oneSentence: () => request.get('https://v1.hitokoto.cn?c=i'),
  login: (data: any) => request.post('/auth/login', data),
  logout: () => request.get('/auth/logout'),
  register: (data: any) => request.post('/auth/register', data),
  userInfo: () => request.get('/auth/info'), // do auth check
}
