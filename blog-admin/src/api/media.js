import request from './request'

/**
 * 上传文件
 * @param {File} file 文件对象
 * @returns {Promise}
 */
export function uploadFile(file) {
  const formData = new FormData()
  formData.append('file', file)
  return request({
    url: '/media/upload',
    method: 'post',
    data: formData,
    headers: {
      'Content-Type': 'multipart/form-data'
    }
  })
}
