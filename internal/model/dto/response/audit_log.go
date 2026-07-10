package response

import "time"

// AuditLogResponse 审计日志响应
type AuditLogResponse struct {
	ID           uint      `json:"id"`
	OperatorID   uint      `json:"operator_id"`
	OperatorName string    `json:"operator_name"`
	Action       string    `json:"action"`
	TargetType   string    `json:"target_type"`
	TargetID     uint      `json:"target_id"`
	TargetTitle  string    `json:"target_title"`
	Detail       string    `json:"detail"`
	IP           string    `json:"ip"`
	CreatedAt    time.Time `json:"created_at"`
}
