package service

import (
	"blog/internal/model/entity"
	"blog/internal/repository"
)

// userService 用户服务实现
type userService struct {
	userRepo repository.UserRepository
}

// NewUserService 创建用户服务
func NewUserService(userRepo repository.UserRepository) UserService {
	return &userService{userRepo: userRepo}
}

// GetUserByID 根据 ID 获取用户
func (s *userService) GetUserByID(id uint) (*entity.User, error) {
	return s.userRepo.FindByID(id)
}

// GetUserProfile 获取用户信息
func (s *userService) GetUserProfile(userID uint) (*entity.User, error) {
	return s.userRepo.FindByID(userID)
}
