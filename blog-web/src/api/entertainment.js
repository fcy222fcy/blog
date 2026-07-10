import request from './request'

export function getEntertainmentPublic(params) {
  return request({
    url: '/entertainment',
    method: 'get',
    params
  })
}
