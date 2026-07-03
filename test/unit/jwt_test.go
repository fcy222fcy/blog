package unit

import (
	"testing"

	"blog/pkg/config"
	"blog/pkg/jwt"
)

func TestJWT_GenerateToken(t *testing.T) {
	cfg := config.JWTConfig{
		Secret:     "test-secret-key",
		ExpireHour: 24,
	}
	j := jwt.NewJWT(cfg)

	token, expiresAt, err := j.GenerateToken(1, "admin")
	if err != nil {
		t.Fatalf("GenerateToken() error = %v", err)
	}
	if token == "" {
		t.Fatal("GenerateToken() returned empty token")
	}
	if expiresAt == 0 {
		t.Fatal("GenerateToken() returned zero ExpiresAt")
	}
}

func TestJWT_ParseToken(t *testing.T) {
	cfg := config.JWTConfig{
		Secret:     "test-secret-key",
		ExpireHour: 24,
	}
	j := jwt.NewJWT(cfg)

	token, _, err := j.GenerateToken(1, "admin")
	if err != nil {
		t.Fatalf("GenerateToken() error = %v", err)
	}

	claims, err := j.ParseToken(token)
	if err != nil {
		t.Fatalf("ParseToken() error = %v", err)
	}
	if claims.UserID != 1 {
		t.Errorf("ParseToken() UserID = %v, want 1", claims.UserID)
	}
	if claims.Username != "admin" {
		t.Errorf("ParseToken() Username = %v, want admin", claims.Username)
	}
}

func TestJWT_ParseToken_InvalidToken(t *testing.T) {
	cfg := config.JWTConfig{
		Secret:     "test-secret-key",
		ExpireHour: 24,
	}
	j := jwt.NewJWT(cfg)

	_, err := j.ParseToken("invalid-token")
	if err == nil {
		t.Fatal("ParseToken() expected error for invalid token")
	}
}

func TestJWT_ParseToken_WrongSecret(t *testing.T) {
	cfg1 := config.JWTConfig{
		Secret:     "secret-1",
		ExpireHour: 24,
	}
	cfg2 := config.JWTConfig{
		Secret:     "secret-2",
		ExpireHour: 24,
	}

	j1 := jwt.NewJWT(cfg1)
	j2 := jwt.NewJWT(cfg2)

	token, _, _ := j1.GenerateToken(1, "admin")
	_, err := j2.ParseToken(token)
	if err == nil {
		t.Fatal("ParseToken() expected error with wrong secret")
	}
}

func TestJWT_GenerateToken_DifferentUsers(t *testing.T) {
	cfg := config.JWTConfig{
		Secret:     "test-secret-key",
		ExpireHour: 24,
	}
	j := jwt.NewJWT(cfg)

	token1, _, _ := j.GenerateToken(1, "admin")
	token2, _, _ := j.GenerateToken(2, "editor")

	if token1 == token2 {
		t.Fatal("Different users should have different tokens")
	}

	claims1, _ := j.ParseToken(token1)
	claims2, _ := j.ParseToken(token2)

	if claims1.UserID == claims2.UserID {
		t.Fatal("Different users should have different UserIDs")
	}
}

func TestJWT_Claims_Fields(t *testing.T) {
	cfg := config.JWTConfig{
		Secret:     "test-secret-key",
		ExpireHour: 24,
	}
	j := jwt.NewJWT(cfg)

	token, _, _ := j.GenerateToken(42, "testuser")
	claims, err := j.ParseToken(token)
	if err != nil {
		t.Fatalf("ParseToken() error = %v", err)
	}

	if claims.UserID != 42 {
		t.Errorf("UserID = %v, want 42", claims.UserID)
	}
	if claims.Username != "testuser" {
		t.Errorf("Username = %v, want testuser", claims.Username)
	}
	if claims.Issuer != "blog" {
		t.Errorf("Issuer = %v, want blog", claims.Issuer)
	}
	if claims.ExpiresAt == nil {
		t.Fatal("ExpiresAt should not be nil")
	}
	if claims.IssuedAt == nil {
		t.Fatal("IssuedAt should not be nil")
	}
}