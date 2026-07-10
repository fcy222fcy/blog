import request from './request'

export function getEntertainmentList(params) {
  return request({
    url: '/admin/entertainment',
    method: 'get',
    params
  })
}

export function getEntertainmentById(id) {
  return request({
    url: `/admin/entertainment/${id}`,
    method: 'get'
  })
}

export function createEntertainment(data) {
  return request({
    url: '/admin/entertainment',
    method: 'post',
    data
  })
}

export function updateEntertainment(id, data) {
  return request({
    url: `/admin/entertainment/${id}`,
    method: 'put',
    data
  })
}

export function deleteEntertainment(id) {
  return request({
    url: `/admin/entertainment/${id}`,
    method: 'delete'
  })
}
