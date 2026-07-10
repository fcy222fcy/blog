package config

import (
	"fmt"
	"os"

	"github.com/spf13/viper"
)

// Config 应用配置
type Config struct {
	Server  ServerConfig  `mapstructure:"server"`
	MySQL   MySQLConfig   `mapstructure:"mysql"`
	Redis   RedisConfig   `mapstructure:"redis"`
	JWT     JWTConfig     `mapstructure:"jwt"`
	Log     LogConfig     `mapstructure:"log"`
	Email   EmailConfig   `mapstructure:"email"`
	Blogger BloggerConfig `mapstructure:"blogger"`
}

// ServerConfig 服务器配置
type ServerConfig struct {
	Port int `mapstructure:"port"`
}

// MySQLConfig 数据库配置
type MySQLConfig struct {
	Host     string `mapstructure:"host"`
	Port     int    `mapstructure:"port"`
	Username string `mapstructure:"username"`
	Password string `mapstructure:"password"`
	DBName   string `mapstructure:"dbname"`
}

// DSN 获取数据库连接字符串
func (m *MySQLConfig) DSN() string {
	return fmt.Sprintf("%s:%s@tcp(%s:%d)/%s?charset=utf8mb4&parseTime=True&loc=Local",
		m.Username, m.Password, m.Host, m.Port, m.DBName)
}

// RedisConfig Redis 配置
type RedisConfig struct {
	Host     string `mapstructure:"host"`
	Port     int    `mapstructure:"port"`
	Password string `mapstructure:"password"`
	DB       int    `mapstructure:"db"`
}

// JWTConfig JWT 配置
type JWTConfig struct {
	Secret     string `mapstructure:"secret"`
	ExpireHour int    `mapstructure:"expire_hour"`
}

// LogConfig 日志配置
type LogConfig struct {
	Level      string `mapstructure:"level"`
	Filename   string `mapstructure:"filename"`
	MaxSize    int    `mapstructure:"max_size"`
	MaxBackups int    `mapstructure:"max_backups"`
	MaxAge     int    `mapstructure:"max_age"`
}

// EmailConfig 邮件配置
type EmailConfig struct {
	Host      string `mapstructure:"host"`
	Port      int    `mapstructure:"port"`
	Username  string `mapstructure:"username"`
	Password  string `mapstructure:"password"`
	From      string `mapstructure:"from"`
	FromEmail string `mapstructure:"from_email"`
}

// BloggerConfig 博主账号配置（写死配置，不走用户表）
type BloggerConfig struct {
	UserID   uint   `mapstructure:"user_id"`  // 博主虚拟 UserID，固定值用于判断博主身份
	Username string `mapstructure:"username"` // 登录用户名
	Password string `mapstructure:"password"` // 登录密码（明文，注意生产环境应加密或环境变量覆盖）
	Nickname string `mapstructure:"nickname"` // 显示昵称
	Email    string `mapstructure:"email"`    // 邮箱（用于头像/通知，与 from_email 保持一致）
	Avatar   string `mapstructure:"avatar"`   // 头像URL（可选，空则用邮箱Gravatar）
}

// Load 加载配置
func Load() (*Config, error) {
	viper.SetConfigName("config")
	viper.SetConfigType("yaml")
	viper.AddConfigPath("configs")
	viper.AddConfigPath(".")

	// 读取环境变量
	viper.AutomaticEnv()

	// 读取配置文件
	if err := viper.ReadInConfig(); err != nil {
		return nil, fmt.Errorf("读取配置文件失败: %w", err)
	}

	var config Config
	if err := viper.Unmarshal(&config); err != nil {
		return nil, fmt.Errorf("解析配置文件失败: %w", err)
	}

	// 环境变量覆盖敏感配置
	if password := os.Getenv("MYSQL_PASSWORD"); password != "" {
		config.MySQL.Password = password
	}
	if password := os.Getenv("REDIS_PASSWORD"); password != "" {
		config.Redis.Password = password
	}
	if secret := os.Getenv("JWT_SECRET"); secret != "" {
		config.JWT.Secret = secret
	}

	return &config, nil
}
