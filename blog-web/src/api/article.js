import request from './request'

export const getArticleList = (params) => {
  return request.get('/articles', { params })
}

export const getArticleDetail = (slug) => {
  return request.get(`/articles/${slug}`)
}

export const getArchives = () => {
  return request.get('/articles/archives')
}

export const searchArticles = (params) => {
  return request.get('/articles/search', { params })
}
