import request from './request'

export const MEDIA_CATEGORIES = {
  AVATAR: 'avatar',
  ARTICLE: 'article',
  DAILY: 'daily',
  ENTERTAINMENT: 'entertainment',
  LINK: 'link',
  COMMON: 'common'
}

/**
 * 上传文件
 * @param {File} file 文件对象
 * @param {string} [category='common'] 分类：avatar | article | daily | entertainment | link | common
 * @returns {Promise}
 */
export function uploadFile(file, category = MEDIA_CATEGORIES.COMMON) {
  const formData = new FormData()
  formData.append('file', file)
  formData.append('category', category)
  return request({
    url: '/media/upload',
    method: 'post',
    data: formData,
    headers: {
      'Content-Type': 'multipart/form-data'
    }
  })
}

/**
 * 获取媒体列表
 * @param {object} params { page, page_size, keyword, category }
 * @returns {Promise}
 */
export function getMediaList(params = {}) {
  return request({
    url: '/media',
    method: 'get',
    params
  })
}

/**
 * 删除媒体文件
 * @param {string} filename 文件名（纯文件名，或 "分类/文件名"）
 * @param {string} [category] 分类，不传则从 filename 路径解析
 * @returns {Promise}
 */
export function deleteMedia(filename, category) {
  return request({
    url: `/media/${encodeURIComponent(filename)}`,
    method: 'delete',
    params: category ? { category } : {}
  })
}
