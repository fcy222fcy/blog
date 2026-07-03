const errorMap = {
  400: '请求参数错误',
  401: '未授权，请登录',
  403: '拒绝访问',
  404: '请求的资源不存在',
  500: '服务器内部错误',
  502: '网关错误',
  503: '服务暂时不可用'
}

/**
 * 轻量级 Toast 提示（不依赖 UI 库）
 * @param {string} message - 提示消息
 * @param {string} type - 类型：error | success | warning
 */
const showToast = (message, type = 'error') => {
  const toast = document.createElement('div')
  toast.textContent = message
  toast.style.cssText = `
    position: fixed; top: 20px; left: 50%; transform: translateX(-50%);
    padding: 12px 24px; border-radius: 8px; color: #fff; font-size: 14px;
    z-index: 10000; animation: toast-in 0.3s ease;
    background: ${type === 'error' ? '#ef4444' : type === 'success' ? '#22c55e' : '#f59e0b'};
    box-shadow: 0 4px 12px rgba(0,0,0,0.15);
  `
  document.body.appendChild(toast)
  setTimeout(() => {
    toast.style.opacity = '0'
    toast.style.transition = 'opacity 0.3s'
    setTimeout(() => toast.remove(), 300)
  }, 3000)
}

/**
 * 统一错误处理函数
 * @param {Error} error - 错误对象
 * @param {Object} options - 配置项
 * @param {boolean} options.showMessage - 是否显示错误消息，默认 true
 * @param {boolean} options.autoRedirect - 401 时是否自动跳转登录页，默认 true
 * @returns {string} 错误消息
 */
export const handleError = (error, options = {}) => {
  const { showMessage = true, autoRedirect = true } = options

  let message = '操作失败，请稍后重试'

  if (error.response) {
    const status = error.response.status
    message = error.response.data?.message || errorMap[status] || message

    // 401 未授权，跳转登录页
    if (status === 401 && autoRedirect) {
      localStorage.removeItem('token')
      window.location.href = '#/login'
    }
  } else if (error.request) {
    message = '网络连接失败，请检查网络'
  } else {
    message = error.message || message
  }

  if (showMessage) {
    showToast(message)
  }

  return message
}

/**
 * 获取友好的错误消息
 * @param {Error} error - 错误对象
 * @returns {string} 错误消息
 */
export const getErrorMessage = (error) => {
  if (error.response) {
    const status = error.response.status
    return error.response.data?.message || errorMap[status] || '操作失败'
  }
  if (error.request) {
    return '网络连接失败，请检查网络'
  }
  return error.message || '操作失败'
}
