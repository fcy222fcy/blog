const pad = (n) => (n < 10 ? `0${n}` : `${n}`)

export function formatDate(dateStr, opts = {}) {
  if (!dateStr) return ''
  const { withTime = false } = opts
  const date = new Date(dateStr)
  if (isNaN(date.getTime())) return ''

  const now = Date.now()
  const diff = (now - date.getTime()) / 1000
  const THRESHOLD_72H = 72 * 3600

  if (diff < 0) {
    return formatAbsolute(date, withTime)
  }
  if (diff < 60) {
    return '刚刚'
  }
  if (diff < 3600) {
    return `${Math.floor(diff / 60)}分钟前`
  }
  if (diff < 86400) {
    return `${Math.floor(diff / 3600)}小时前`
  }
  if (diff < THRESHOLD_72H) {
    return `${Math.floor(diff / 86400)}天前`
  }
  return formatAbsolute(date, withTime)
}

function formatAbsolute(date, withTime) {
  const y = date.getFullYear()
  const m = pad(date.getMonth() + 1)
  const d = pad(date.getDate())
  const base = `${y}-${m}-${d}`
  if (!withTime) return base
  const hh = pad(date.getHours())
  const mm = pad(date.getMinutes())
  return `${base} ${hh}:${mm}`
}
