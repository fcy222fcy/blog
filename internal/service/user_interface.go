package service

import (
	"blog/internal/model/dto/request"
	"blog/internal/model/entity"
)

// UserService 用户服务接口
type UserService interface {
	// GetUserByID 根据 ID 获取用户
	GetUserByID(id uint) (*entity.User, error)

	// GetUserProfile 获取用户信息
	GetUserProfile(userID uint) (*entity.User, error)

	// GetFirstAdmin 获取第一个管理员用户（博客主人）
	GetFirstAdmin() (*entity.User, error)

	// UpdateUser 更新用户信息
	UpdateUser(id uint, req *request.UpdateUserRequest) error
}
