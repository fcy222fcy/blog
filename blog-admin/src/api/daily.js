import request from './request'

export const getDailyQuestionList = (params) => request.get('/admin/daily-questions', { params })
export const createDailyQuestion = (data) => request.post('/admin/daily-questions', data)
export const updateDailyQuestion = (id, data) => request.put(`/admin/daily-questions/${id}`, data)
export const deleteDailyQuestion = (id) => request.delete(`/admin/daily-questions/${id}`)
export const updateDailyQuestionStatus = (id, status) => request.put(`/admin/daily-questions/${id}/status`, { status })
