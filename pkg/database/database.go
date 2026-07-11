package database

import (
	"fmt"
	"time"

	"blog/internal/model/entity"
	"blog/pkg/config"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
)

// Database 数据库
type Database struct {
	DB *gorm.DB
}

// NewDatabase 创建数据库连接
func NewDatabase(cfg config.MySQLConfig) (*Database, error) {
	db, err := gorm.Open(mysql.Open(cfg.DSN()), &gorm.Config{
		Logger: logger.Default.LogMode(logger.Info),
	})
	if err != nil {
		return nil, fmt.Errorf("连接数据库失败: %w", err)
	}

	// 获取底层 *sql.DB
	sqlDB, err := db.DB()
	if err != nil {
		return nil, fmt.Errorf("获取数据库连接失败: %w", err)
	}

	// 连接池配置
	sqlDB.SetMaxIdleConns(10)
	sqlDB.SetMaxOpenConns(100)
	sqlDB.SetConnMaxLifetime(time.Hour)

	return &Database{DB: db}, nil
}

// AutoMigrate 自动迁移
func (d *Database) AutoMigrate() error {
	return d.DB.AutoMigrate(
		&entity.User{},
		&entity.Article{},
		&entity.Category{},
		&entity.Tag{},
		&entity.Comment{},
		&entity.CommentLikeLog{},
		&entity.DailyQuestion{},
		&entity.Media{},
		&entity.AboutPage{},
		&entity.VisitLog{},
		&entity.AuditLog{},
	)
}

// Close 关闭数据库连接
func (d *Database) Close() error {
	sqlDB, err := d.DB.DB()
	if err != nil {
		return err
	}
	return sqlDB.Close()
}
