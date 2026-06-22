import request from './request'

export const getLinkList = () => {
  return request.get('/links')
}
