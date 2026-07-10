const qqEmailRegexp = /^(\d+)@(?:vip\.)?qq\.com$/i

const QQ_SUPPORTED_SIZES = [40, 100, 140, 240, 640]

function normalizeQQSize(size) {
  const s = Number(size)
  if (!s || s <= 0) return 100
  for (const valid of QQ_SUPPORTED_SIZES) {
    if (valid >= s) return valid
  }
  return QQ_SUPPORTED_SIZES[QQ_SUPPORTED_SIZES.length - 1]
}

/**
 * 根据 QQ 号获取 QQ 官方头像
 * @param {string} qq QQ 号
 * @param {number} size 尺寸（自动映射到 qlogo.cn 支持的 40/100/140/240/640）
 * @returns {string} 头像 URL
 */
export function getQQAvatarUrl(qq, size = 80) {
  if (!qq) return ''
  const normalized = normalizeQQSize(size)
  return `https://q1.qlogo.cn/g?b=qq&nk=${encodeURIComponent(qq)}&s=${normalized}`
}

/**
 * 根据邮箱获取头像 URL（智能判断）
 * 1. QQ 邮箱 → 取 QQ 官方真实头像（qlogo.cn）
 * 2. 其他邮箱 → 使用 ui-avatars.com 生成首字母头像兜底
 * @param {string} email 邮箱地址
 * @param {number} size 头像尺寸（默认 80）
 * @returns {string} 头像 URL
 */
export function getAvatarUrl(email, size = 80) {
  if (!email) return ''
  const lower = String(email).trim().toLowerCase()
  const m = lower.match(qqEmailRegexp)
  if (m && m[1]) {
    return getQQAvatarUrl(m[1], size)
  }
  const name = email.split('@')[0]
  return `https://ui-avatars.com/api/?name=${encodeURIComponent(name)}&size=${size}&background=random&color=fff&bold=true`
}
