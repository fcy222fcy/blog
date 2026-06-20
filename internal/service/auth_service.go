package service

import (
	"blog/internal/model/dto/request"
	"blog/internal/model/dto/response"
	"blog/internal/repository"
	bizerrors "blog/pkg/errors"
	"blog/pkg/jwt"
)

// authService 认证服务实现
type authService struct {
	userRepo repository.UserRepository
	jwt      *jwt.JWT
}

// NewAuthService 创建认证服务
func NewAuthService(userRepo repository.UserRepository, jwt *jwt.JWT) AuthService {
	return &authService{
		userRepo: userRepo,
		jwt:      jwt,
	}
}

// Login 用户登录
func (s *authService) Login(req *request.LoginRequest) (*response.LoginResponse, error) {
	// 查找用户
	user, err := s.userRepo.FindByUsername(req.Username)
	if err != nil {
		return nil, err
	}
	if user == nil {
		return nil, bizerrors.New(bizerrors.CodeUserNotFound, "用户不存在")
	}

	// TODO: 验证密码（需要 bcrypt）
	// if !bcrypt.CheckPassword(req.Password, user.Password) {
	//     return nil, bizerrors.New(bizerrors.CodePasswordIncorrect, "密码错误")
	// }

	// 生成 Token
	token, err := s.jwt.GenerateToken(user.ID, user.Username)
	if err != nil {
		return nil, err
	}

	return &response.LoginResponse{
		Token:     token,
		ExpiresAt: 0, // TODO: 计算过期时间
	}, nil
}

// GetProfile 获取用户信息
func (s *authService) GetProfile(userID uint) (*response.UserProfileResponse, error) {
	user, err := s.userRepo.FindByID(userID)
	if err != nil {
		return nil, err
	}
	if user == nil {
		return nil, bizerrors.New(bizerrors.CodeUserNotFound, "用户不存在")
	}

	return &response.UserProfileResponse{
		ID:       user.ID,
		Username: user.Username,
		Nickname: user.Nickname,
		Email:    user.Email,
		Avatar:   user.Avatar,
		Bio:      user.Bio,
	}, nil
}

// ChangePassword 修改密码
func (s *authService) ChangePassword(userID uint, req *request.ChangePasswordRequest) error {
	user, err := s.userRepo.FindByID(userID)
	if err != nil {
		return err
	}
	if user == nil {
		return bizerrors.New(bizerrors.CodeUserNotFound, "用户不存在")
	}

	// TODO: 验证旧密码
	// if !bcrypt.CheckPassword(req.OldPassword, user.Password) {
	//     return bizerrors.New(bizerrors.CodePasswordIncorrect, "旧密码错误")
	// }

	// TODO: 加密新密码
	// user.Password = bcrypt.HashPassword(req.NewPassword)
	// return s.userRepo.Update(user)

	return nil
}
