package service

import (
	"blog/internal/model/dto/request"
	"blog/internal/model/entity"
	"blog/internal/repository"
	bizcrypt "blog/pkg/bcrypt"
	"blog/pkg/config"
	bizerrors "blog/pkg/errors"
	"blog/pkg/logger"
	"fmt"
)

// userService 用户服务实现
type userService struct {
	userRepo repository.UserRepository
	config   *config.Config
}

// NewUserService 创建用户服务
func NewUserService(userRepo repository.UserRepository, cfg *config.Config) UserService {
	return &userService{
		userRepo: userRepo,
		config:   cfg,
	}
}

// GetUserByID 根据 ID 获取用户
func (s *userService) GetUserByID(id uint) (*entity.User, error) {
	user, err := s.userRepo.FindByID(id)
	if err != nil {
		return nil, fmt.Errorf("查询用户失败, %w", err)
	}
	// 博主用户：找不到时自动基于配置创建一条，保证数据源完整
	if user == nil && s.isBloggerID(id) {
		user, err = s.ensureBloggerUser()
		if err != nil {
			return nil, err
		}
	}
	if user == nil {
		return nil, bizerrors.New(bizerrors.CodeUserNotFound, bizerrors.GetMessage(bizerrors.CodeUserNotFound))
	}
	return user, nil
}

// GetUserProfile 获取用户信息
func (s *userService) GetUserProfile(userID uint) (*entity.User, error) {
	user, err := s.userRepo.FindByID(userID)
	if err != nil {
		return nil, fmt.Errorf("查询用户信息失败, %w", err)
	}
	if user == nil && s.isBloggerID(userID) {
		user, err = s.ensureBloggerUser()
		if err != nil {
			return nil, err
		}
	}
	if user == nil {
		return nil, bizerrors.New(bizerrors.CodeUserNotFound, bizerrors.GetMessage(bizerrors.CodeUserNotFound))
	}
	return user, nil
}

// GetFirstAdmin 获取第一个管理员用户（博客主人）
func (s *userService) GetFirstAdmin() (*entity.User, error) {
	user, err := s.userRepo.FindFirstAdmin()
	if err != nil {
		return nil, fmt.Errorf("查询管理员用户失败, %w", err)
	}
	return user, nil
}

// UpdateUser 更新用户信息
func (s *userService) UpdateUser(id uint, req *request.UpdateUserRequest) error {
	user, err := s.userRepo.FindByID(id)
	if err != nil {
		return fmt.Errorf("查询用户失败, %w", err)
	}

	// 博主用户：user 表中无记录时自动基于配置创建一条，保证后续更新可用
	if user == nil && s.isBloggerID(id) {
		user, err = s.ensureBloggerUser()
		if err != nil {
			return err
		}
	}

	if user == nil {
		return bizerrors.New(bizerrors.CodeUserNotFound, bizerrors.GetMessage(bizerrors.CodeUserNotFound))
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

	if err := s.userRepo.Update(user); err != nil {
		return fmt.Errorf("更新用户信息失败, %w", err)
	}

	logger.Infof("更新用户信息成功, id: %d", id)
	return nil
}

// isBloggerID 判断 userID 是否为配置中的博主 ID
func (s *userService) isBloggerID(userID uint) bool {
	if s.config == nil {
		return false
	}
	return userID == s.config.Blogger.UserID
}

// ensureBloggerUser 基于 blogger 配置创建用户表记录并返回（幂等，已存在则直接返回）
func (s *userService) ensureBloggerUser() (*entity.User, error) {
	if s.config == nil {
		return nil, fmt.Errorf("配置未加载，无法创建博主用户")
	}
	b := s.config.Blogger
	existing, err := s.userRepo.FindByID(b.UserID)
	if err != nil {
		return nil, fmt.Errorf("查询博主用户失败: %w", err)
	}
	if existing != nil {
		return existing, nil
	}
	hashed, hashErr := bizcrypt.HashPassword(b.Password)
	if hashErr != nil {
		return nil, fmt.Errorf("创建博主用户失败(密码加密): %w", hashErr)
	}
	user := &entity.User{}
	user.ID = b.UserID
	user.Username = b.Username
	user.Password = hashed
	user.Nickname = b.Nickname
	user.Email = b.Email
	user.Avatar = b.Avatar
	user.Status = 1
	if createErr := s.userRepo.Create(user); createErr != nil {
		return nil, fmt.Errorf("创建博主用户失败: %w", createErr)
	}
	logger.Infof("自动创建博主用户记录, id: %d, username: %s", user.ID, user.Username)
	return user, nil
}
