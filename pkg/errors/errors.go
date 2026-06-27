package errors

import "errors"

// 预定义错误
var (
	ErrNotFound              = errors.New("资源不存在")
	ErrInvalidParam          = errors.New("参数错误")
	ErrUnauthorized          = errors.New("未授权")
	ErrForbidden             = errors.New("禁止访问")
	ErrInternalServer        = errors.New("服务器内部错误")
	ErrDailyQuestionNotFound = errors.New("每日一问不存在")
	ErrCommentNotFound       = errors.New("评论不存在")
	ErrNoAvailableQuestion   = errors.New("暂无可用问题")
)

// BizError 业务错误
type BizError struct {
	Code    int    `json:"code"`
	Message string `json:"message"`
	Detail  string `json:"detail,omitempty"`
}

// Error 实现 error 接口
func (e *BizError) Error() string {
	if e.Detail == "" {
		return e.Message
	}
	return e.Detail
}

// New 创建新的业务错误
func New(code int, message string) *BizError {
	return &BizError{
		Code:    code,
		Message: message,
	}
}

// NewWithDetail 创建带详细信息的业务错误
func NewWithDetail(code int, message string, detail string) *BizError {
	return &BizError{
		Code:    code,
		Message: message,
		Detail:  detail,
	}
}

// IsBizError 判断错误是否是业务错误
func IsBizError(err error) bool {
	_, ok := err.(*BizError)
	return ok
}
