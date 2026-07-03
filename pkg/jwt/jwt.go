package jwt

import (
	"errors"
	"time"

	"blog/pkg/config"

	"github.com/golang-jwt/jwt/v5"
)

// Claims JWT Claims
type Claims struct {
	UserID   uint   `json:"user_id"`
	Username string `json:"username"`
	jwt.RegisteredClaims
}

// JWT JWT 工具
type JWT struct {
	config config.JWTConfig
}

// NewJWT 创建 JWT 实例
func NewJWT(cfg config.JWTConfig) *JWT {
	return &JWT{config: cfg}
}

// GenerateToken 生成 Token，返回 token 字符串、过期时间戳（Unix 秒）和错误
func (j *JWT) GenerateToken(userID uint, username string) (string, int64, error) {
	now := time.Now()
	expiresAt := now.Add(time.Duration(j.config.ExpireHour) * time.Hour)

	claims := Claims{
		UserID:   userID,
		Username: username,
		RegisteredClaims: jwt.RegisteredClaims{
			ExpiresAt: jwt.NewNumericDate(expiresAt),
			IssuedAt:  jwt.NewNumericDate(now),
			Issuer:    "blog",
		},
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	tokenString, err := token.SignedString([]byte(j.config.Secret))
	return tokenString, expiresAt.Unix(), err
}

// ParseToken 解析 Token
func (j *JWT) ParseToken(tokenString string) (*Claims, error) {
	token, err := jwt.ParseWithClaims(tokenString, &Claims{}, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, errors.New("unexpected signing method")
		}
		return []byte(j.config.Secret), nil
	})

	if err != nil {
		return nil, err
	}

	if claims, ok := token.Claims.(*Claims); ok && token.Valid {
		return claims, nil
	}

	return nil, errors.New("invalid token")
}
