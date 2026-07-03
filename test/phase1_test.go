package test

import (
	"blog/pkg/bcrypt"
	"blog/pkg/config"
	"blog/pkg/jwt"
	"testing"
	"time"
)

// TestPasswordEncryption 测试密码加密功能
func TestPasswordEncryption(t *testing.T) {
	password := "TestPassword123!"
	
	// 测试加密
	hashedPassword, err := bcrypt.HashPassword(password)
	if err != nil {
		t.Fatalf("密码加密失败: %v", err)
	}
	
	t.Logf("原始密码: %s", password)
	t.Logf("加密后密码: %s", hashedPassword)
	
	// 验证密码长度
	if len(hashedPassword) < 60 {
		t.Errorf("加密后的密码长度不正确, 期望至少60个字符, 实际: %d", len(hashedPassword))
	}
	
	// 测试正确密码验证
	if !bcrypt.CheckPassword(password, hashedPassword) {
		t.Error("正确密码验证失败")
	}
	
	// 测试错误密码验证
	wrongPassword := "WrongPassword!"
	if bcrypt.CheckPassword(wrongPassword, hashedPassword) {
		t.Error("错误密码验证应该失败但通过了")
	}
	
	t.Log("✅ 密码加密功能测试通过")
}

// TestJWTExpiration 测试JWT过期时间
func TestJWTExpiration(t *testing.T) {
	cfg := config.JWTConfig{
		Secret:     "test-secret-key",
		ExpireHour: 24, // 24小时
	}
	
	jwtInstance := jwt.NewJWT(cfg)
	
	// 生成Token
	userID := uint(1)
	username := "testuser"
	tokenString, expiresAt, err := jwtInstance.GenerateToken(userID, username)
	
	if err != nil {
		t.Fatalf("生成Token失败: %v", err)
	}
	
	t.Logf("生成的Token: %s", tokenString)
	t.Logf("过期时间戳: %d", expiresAt)
	
	// 验证过期时间不为0
	if expiresAt == 0 {
		t.Error("❌ 过期时间戳不应该为0")
	}
	
	// 验证过期时间大约在24小时后
	now := time.Now().Unix()
	expectedExpiration := now + 24*60*60
	timeDiff := expiresAt - expectedExpiration
	
	if timeDiff < -10 || timeDiff > 10 {
		t.Errorf("过期时间不正确, 预期约: %d, 实际: %d, 差距: %d秒", expectedExpiration, expiresAt, timeDiff)
	}
	
	// 验证可以解析Token
	claims, err := jwtInstance.ParseToken(tokenString)
	if err != nil {
		t.Fatalf("解析Token失败: %v", err)
	}
	
	if claims.UserID != userID {
		t.Errorf("UserID不匹配, 期望: %d, 实际: %d", userID, claims.UserID)
	}
	
	if claims.Username != username {
		t.Errorf("Username不匹配, 期望: %s, 实际: %s", username, claims.Username)
	}
	
	t.Log("✅ JWT过期时间功能测试通过")
}

// TestBcryptConsistency 测试bcrypt一致性
func TestBcryptConsistency(t *testing.T) {
	password := "SamePassword"
	
	// 多次加密同一密码，结果应该不同（因为使用了salt）
	hash1, _ := bcrypt.HashPassword(password)
	hash2, _ := bcrypt.HashPassword(password)
	
	if hash1 == hash2 {
		t.Error("相同密码的两次加密结果不应该相同")
	}
	
	// 但两个hash都应该能验证原密码
	if !bcrypt.CheckPassword(password, hash1) {
		t.Error("hash1验证失败")
	}
	
	if !bcrypt.CheckPassword(password, hash2) {
		t.Error("hash2验证失败")
	}
	
	t.Log("✅ Bcrypt一致性测试通过")
}
