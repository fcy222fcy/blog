/**
 * 根据邮箱获取头像 URL（使用 ui-avatars.com）
 * @param {string} email 邮箱地址
 * @param {number} size 头像尺寸（默认 80）
 * @returns {string} 头像 URL
 */
export function getAvatarUrl(email, size = 80) {
  if (!email) return ''
  const name = email.split('@')[0]
  return `https://ui-avatars.com/api/?name=${encodeURIComponent(name)}&size=${size}&background=random&color=fff&bold=true`
}
