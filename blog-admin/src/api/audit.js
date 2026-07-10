import request from './request'

// 获取审计日志列表
export function getAuditLogList(params) {
  return request.get('/admin/audit-logs', { params })
}
