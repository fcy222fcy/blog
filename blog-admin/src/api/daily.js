import request from './request'

// 管理后台接口
export const getDailyQuestionList = (params) => request.get('/admin/daily-questions', { params })
export const createDailyQuestion = (data) => request.post('/admin/daily-questions', data)
export const updateDailyQuestion = (id, data) => request.put(`/admin/daily-questions/${id}`, data)
export const deleteDailyQuestion = (id) => request.delete(`/admin/daily-questions/${id}`)
export const updateDailyQuestionStatus = (id, status) => request.put(`/admin/daily-questions/${id}/status`, { status })

// 前台展示接口（供 DailyQuestion 组件使用）
export const getLatestQuestion = () => request.get('/daily-questions/latest')
export const getQuestionByDate = (date) => request.get(`/daily-questions/date/${date}`)
export const likeQuestion = (id) => request.post(`/daily-questions/${id}/like`)
