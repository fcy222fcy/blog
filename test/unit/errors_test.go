package unit

import (
	"testing"

	bizerrors "blog/pkg/errors"
)

func TestNew(t *testing.T) {
	err := bizerrors.New(1001, "用户不存在")
	if err == nil {
		t.Fatal("New() should not return nil")
	}
	if err.Code != 1001 {
		t.Errorf("Code = %v, want 1001", err.Code)
	}
	if err.Message != "用户不存在" {
		t.Errorf("Message = %v, want 用户不存在", err.Message)
	}
}

func TestNewWithDetail(t *testing.T) {
	err := bizerrors.NewWithDetail(1001, "用户不存在", "用户名 admin 不存在")
	if err == nil {
		t.Fatal("NewWithDetail() should not return nil")
	}
	if err.Code != 1001 {
		t.Errorf("Code = %v, want 1001", err.Code)
	}
	if err.Message != "用户不存在" {
		t.Errorf("Message = %v, want 用户不存在", err.Message)
	}
	if err.Detail != "用户名 admin 不存在" {
		t.Errorf("Detail = %v, want 用户名 admin 不存在", err.Detail)
	}
}

func TestBizError_Error(t *testing.T) {
	tests := []struct {
		name     string
		err      *bizerrors.BizError
		expected string
	}{
		{
			name:     "有详情",
			err:      bizerrors.NewWithDetail(1001, "用户不存在", "admin not found"),
			expected: "admin not found",
		},
		{
			name:     "无详情",
			err:      bizerrors.New(1001, "用户不存在"),
			expected: "用户不存在",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if tt.err.Error() != tt.expected {
				t.Errorf("Error() = %v, want %v", tt.err.Error(), tt.expected)
			}
		})
	}
}

func TestIsBizError(t *testing.T) {
	bizErr := bizerrors.New(1001, "用户不存在")
	if !bizerrors.IsBizError(bizErr) {
		t.Fatal("IsBizError() should return true for BizError")
	}

	normalErr := bizerrors.ErrNotFound
	if bizerrors.IsBizError(normalErr) {
		t.Fatal("IsBizError() should return false for normal error")
	}
}

func TestGetMessage(t *testing.T) {
	tests := []struct {
		code     int
		expected string
	}{
		{0, "成功"},
		{400, "请求错误"},
		{401, "未授权"},
		{403, "禁止访问"},
		{404, "资源不存在"},
		{500, "服务器内部错误"},
		{1001, "用户不存在"},
		{1002, "用户已存在"},
		{1003, "密码错误"},
		{1004, "Token 已过期"},
		{1005, "Token 无效"},
		{2001, "参数错误"},
		{3001, "文章不存在"},
		{3002, "分类不存在"},
		{3003, "标签不存在"},
		{3004, "评论不存在"},
		{3005, "友链不存在"},
		{3006, "每日一问不存在"},
		{9999, "未知错误"},
	}

	for _, tt := range tests {
		msg := bizerrors.GetMessage(tt.code)
		if msg != tt.expected {
			t.Errorf("GetMessage(%d) = %v, want %v", tt.code, msg, tt.expected)
		}
	}
}

func TestPredefinedErrors(t *testing.T) {
	tests := []struct {
		name     string
		err      error
		expected string
	}{
		{"ErrNotFound", bizerrors.ErrNotFound, "资源不存在"},
		{"ErrInvalidParam", bizerrors.ErrInvalidParam, "参数错误"},
		{"ErrUnauthorized", bizerrors.ErrUnauthorized, "未授权"},
		{"ErrForbidden", bizerrors.ErrForbidden, "禁止访问"},
		{"ErrInternalServer", bizerrors.ErrInternalServer, "服务器内部错误"},
		{"ErrDailyQuestionNotFound", bizerrors.ErrDailyQuestionNotFound, "每日一问不存在"},
		{"ErrCommentNotFound", bizerrors.ErrCommentNotFound, "评论不存在"},
		{"ErrNoAvailableQuestion", bizerrors.ErrNoAvailableQuestion, "暂无可用问题"},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if tt.err.Error() != tt.expected {
				t.Errorf("Error() = %v, want %v", tt.err.Error(), tt.expected)
			}
		})
	}
}
