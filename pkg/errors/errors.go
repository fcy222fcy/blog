package errors

import "errors"

// 预定义错误
var (
	ErrNotFound          = errors.New("资源不存在")
	ErrInvalidParam      = errors.New("参数错误")
	ErrUnauthorized      = errors.New("未授权")
	ErrForbidden         = errors.New("禁止访问")
	ErrInternalServer    = errors.New("服务器内部错误")
	ErrDailyQuestionNotFound = errors.New("每日一问不存在")
	ErrCommentNotFound   = errors.New("评论不存在")
	ErrNoAvailableQuestion = errors.New("暂无可用问题")
)

// BizError 业务错误
type BizError struct {
	Code    int    `json:"code"`
	Message string `json:"message"`
}

// Error 实现 error 接口
func (e *BizError) Error() string {
	return e.Message
}

// New 创建新的业务错误
func New(code int, message string) *BizError {
	return &BizError{
		Code:    code,
		Message: message,
	}
}

// 错误码常量
const (
	CodeSuccess           = 0
	CodeBadRequest        = 40001
	CodeUnauthorized      = 40101
	CodeForbidden         = 40301
	CodeNotFound          = 40401
	CodeInternalServer    = 50001
	CodeDailyQuestionNotFound = 40402
	CodeCommentNotFound   = 40403
	CodeNoAvailableQuestion = 40404
)
