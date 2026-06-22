import request from './request'

export const login = (data) => request.post('/auth/login', data)
export const getUserInfo = () => request.get('/user/info')
export const updateUserInfo = (data) => request.put('/user/info', data)
