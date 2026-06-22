import request from './request'

export const getTagList = () => {
  return request.get('/tags')
}
