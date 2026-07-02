import request from './request'

// 获取关于页面（公开）
export const getAboutPage = () => request.get('/about')

// 更新关于页面（需要登录）
export const updateAboutPage = (data) => request.put('/admin/about', data)
