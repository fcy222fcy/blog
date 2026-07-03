package service

import (
	"blog/internal/model/dto/request"
	"blog/internal/model/dto/response"
	"blog/internal/model/entity"
	"blog/internal/repository"
	bizcrypt "blog/pkg/bcrypt"
	bizerrors "blog/pkg/errors"
	"blog/pkg/jwt"
	"blog/pkg/logger"
	"fmt"
	"strings"

	"go.uber.org/zap"
)

// validateInput 验证输入是否包含危险字符
func validateInput(input string) bool {
	dangerous := []string{"'", "--", "#", ";", "/*", "*/", "UNION", "SELECT", "DROP", "DELETE", "INSERT", "UPDATE", "OR ", "AND "}
	upper := strings.ToUpper(input)
	for _, d := range dangerous {
		if strings.Contains(upper, d) {
			return false
		}
	}
	return true
}

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
	logger.Infof("用户登录, username: %s", req.Username)

	// 验证输入是否包含危险字符
	if !validateInput(req.Username) || !validateInput(req.Password) {
		return nil, bizerrors.New(bizerrors.CodeInvalidParams, "输入包含非法字符")
	}

	// 查找用户
	user, err := s.userRepo.FindByUsername(req.Username)
	if err != nil {
		logger.Error("查询用户失败", zap.Error(err))
		return nil, fmt.Errorf("查询用户失败, %w", err)
	}
	if user == nil {
		logger.Warn("用户不存在", zap.String("username", req.Username))
		return nil, bizerrors.New(bizerrors.CodeUserNotFound, bizerrors.GetMessage(bizerrors.CodeUserNotFound))
	}

	// 验证密码
	if !bizcrypt.CheckPassword(req.Password, user.Password) {
		return nil, bizerrors.New(bizerrors.CodePasswordIncorrect, bizerrors.GetMessage(bizerrors.CodePasswordIncorrect))
	}

	// 生成 Token
	token, expiresAt, err := s.jwt.GenerateToken(user.ID, user.Username)
	if err != nil {
		logger.Error("生成Token失败", zap.Error(err))
		return nil, fmt.Errorf("生成Token失败, %w", err)
	}

	logger.Infof("用户登录成功, username: %s", req.Username)
	return &response.LoginResponse{
		Token:     token,
		ExpiresAt: expiresAt,
		User: response.UserInfo{
			ID:       user.ID,
			Username: user.Username,
			Nickname: user.Nickname,
			Avatar:   user.Avatar,
			Email:    user.Email,
		},
	}, nil
}

// GetProfile 获取用户信息
func (s *authService) GetProfile(userID uint) (*response.UserProfileResponse, error) {
	logger.Infof("获取用户信息, userID: %d", userID)

	user, err := s.userRepo.FindByID(userID)
	if err != nil {
		logger.Error("查询用户失败", zap.Error(err))
		return nil, fmt.Errorf("查询用户失败, %w", err)
	}
	if user == nil {
		logger.Warn("用户不存在", zap.Uint("userID", userID))
		return nil, bizerrors.New(bizerrors.CodeUserNotFound, bizerrors.GetMessage(bizerrors.CodeUserNotFound))
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
	logger.Infof("修改密码, userID: %d", userID)

	user, err := s.userRepo.FindByID(userID)
	if err != nil {
		logger.Error("查询用户失败", zap.Error(err))
		return fmt.Errorf("查询用户失败, %w", err)
	}
	if user == nil {
		logger.Warn("用户不存在", zap.Uint("userID", userID))
		return bizerrors.New(bizerrors.CodeUserNotFound, bizerrors.GetMessage(bizerrors.CodeUserNotFound))
	}

	// 验证旧密码
	if !bizcrypt.CheckPassword(req.OldPassword, user.Password) {
		return bizerrors.New(bizerrors.CodePasswordIncorrect, bizerrors.GetMessage(bizerrors.CodePasswordIncorrect))
	}

	// 加密新密码
	hashedPassword, err := bizcrypt.HashPassword(req.NewPassword)
	if err != nil {
		return bizerrors.New(bizerrors.CodeInternalServer, "密码加密失败")
	}

	user.Password = hashedPassword
	if err := s.userRepo.Update(user); err != nil {
		logger.Error("更新密码失败", zap.Error(err))
		return fmt.Errorf("更新密码失败, %w", err)
	}

	logger.Infof("修改密码成功, userID: %d", userID)
	return nil
}

// Register 用户注册
func (s *authService) Register(req *request.RegisterRequest) error {
	logger.Infof("用户注册, username: %s", req.Username)

	// 检查用户名是否已存在
	existing, err := s.userRepo.FindByUsername(req.Username)
	if err != nil {
		logger.Error("查询用户失败", zap.Error(err))
		return fmt.Errorf("查询用户失败, %w", err)
	}
	if existing != nil {
		logger.Warn("用户名已存在", zap.String("username", req.Username))
		return bizerrors.New(bizerrors.CodeUserAlreadyExists, bizerrors.GetMessage(bizerrors.CodeUserAlreadyExists))
	}

	// 加密密码
	hashedPassword, err := bizcrypt.HashPassword(req.Password)
	if err != nil {
		return bizerrors.New(bizerrors.CodeInternalServer, "密码加密失败")
	}

	user := &entity.User{
		Username: req.Username,
		Password: hashedPassword,
		Nickname: req.Nickname,
		Email:    req.Email,
		Status:   1,
	}

	if err := s.userRepo.Create(user); err != nil {
		logger.Error("创建用户失败", zap.Error(err))
		return fmt.Errorf("创建用户失败, %w", err)
	}

	logger.Infof("用户注册成功, username: %s", req.Username)
	return nil
}
