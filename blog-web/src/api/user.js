import request from './request'

export const getUserProfile = () => {
  return request.get('/user/profile')
}

export const updateUserProfile = (data) => {
  return request.put('/user/profile', data)
}
