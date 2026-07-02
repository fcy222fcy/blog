import request from './request'

// 获取关于页面（公开）
export const getAboutPage = () => request.get('/about')
