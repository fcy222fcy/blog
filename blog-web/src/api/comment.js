import request from './request'

export const getCommentsByArticle = (articleId, params) => {
  return request.get(`/comments/article/${articleId}`, { params })
}

export const createComment = (data) => {
  return request.post('/comments', data)
}

export const likeComment = (commentId) => {
  return request.post(`/comments/${commentId}/like`)
}

export const unlikeComment = (commentId) => {
  return request.delete(`/comments/${commentId}/like`)
}
