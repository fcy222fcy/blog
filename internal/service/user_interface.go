package service

import (
	"blog/internal/model/entity"
)

// UserService 用户服务接口
type UserService interface {
	// GetUserByID 根据 ID 获取用户
	GetUserByID(id uint) (*entity.User, error)

	// GetUserByUsername 根据用户名获取用户
	GetUserProfile(userID uint) (*entity.User, error)
}
