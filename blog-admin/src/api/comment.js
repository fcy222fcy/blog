import request from './request'

export const getCommentList = (params) => request.get('/admin/comments', { params })
export const updateCommentStatus = (id, status) => request.put(`/admin/comments/${id}/status`, { status })
export const deleteComment = (id) => request.delete(`/admin/comments/${id}`)
export const batchDeleteComments = (ids) => request.post('/admin/comments/batch-delete', { ids })
