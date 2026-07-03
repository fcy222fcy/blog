import request from './request'

export const login = (data) => request.post('/auth/login', data)
export const getUserInfo = () => request.get('/user/profile')
export const updateUserInfo = (data) => request.put('/user/profile', data)
