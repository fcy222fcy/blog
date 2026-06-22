package service

import (
	"blog/internal/model/dto/request"
	"blog/internal/model/entity"
	"blog/internal/repository"
	bizerrors "blog/pkg/errors"
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

// UpdateUser 更新用户信息
func (s *userService) UpdateUser(id uint, req *request.UpdateUserRequest) error {
	user, err := s.userRepo.FindByID(id)
	if err != nil {
		return err
	}
	if user == nil {
		return bizerrors.New(bizerrors.CodeUserNotFound, "用户不存在")
	}

	if req.Nickname != "" {
		user.Nickname = req.Nickname
	}
	if req.Email != "" {
		user.Email = req.Email
	}
	if req.Avatar != "" {
		user.Avatar = req.Avatar
	}
	if req.Bio != "" {
		user.Bio = req.Bio
	}

	return s.userRepo.Update(user)
}
