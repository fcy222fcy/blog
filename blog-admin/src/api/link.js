import request from './request'

export const getLinkList = (params) => request.get('/admin/links', { params })
export const createLink = (data) => request.post('/admin/links', data)
export const updateLink = (id, data) => request.put(`/admin/links/${id}`, data)
export const deleteLink = (id) => request.delete(`/admin/links/${id}`)
export const updateLinkStatus = (id, status) => request.put(`/admin/links/${id}/status`, { status })
