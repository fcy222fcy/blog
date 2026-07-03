package gravatar

import (
	"crypto/md5"
	"encoding/hex"
	"fmt"
	"strings"
)

// DefaultType Gravatar 默认头像类型
type DefaultType string

const (
	// 404 - 不返回图片，返回404
	Default404 DefaultType = "404"
	// mp - 神秘人（默认）
	DefaultMP DefaultType = "mp"
	// identicon - 几何图案
	DefaultIdenticon DefaultType = "identicon"
	// monsterid - 怪物
	DefaultMonsterid DefaultType = "monsterid"
	// wavatar - 生成的面孔
	DefaultWavatar DefaultType = "wavatar"
	// retro - 8位复古
	DefaultRetro DefaultType = "retro"
	// robohash - 生成的机器人
	DefaultRobohash DefaultType = "robohash"
	// blank - 透明PNG
	DefaultBlank DefaultType = "blank"
)

// Rating Gravatar 图片评级
type Rating string

const (
	RatingG  Rating = "g"  // 适合所有网站
	RatingPG Rating = "pg" // 可能不适合13岁以下
	RatingR  Rating = "r"  // 可能不适合17岁以下
	RatingX  Rating = "x"  // 仅适合成人
)

// Options Gravatar 选项
type Options struct {
	Size    int         // 图片大小 (1-2048)，默认 80
	Default DefaultType // 默认图片类型
	Rating  Rating      // 图片评级
	ForceDefault bool   // 强制显示默认头像（用于测试）
}

// DefaultOptions 默认选项
var DefaultOptions = Options{
	Size:    80,
	Default: DefaultIdenticon, // 使用几何图案作为默认头像
	Rating:  RatingG,
	ForceDefault: false,
}

// GetAvatarURL 获取 Gravatar 头像 URL
func GetAvatarURL(email string, opts *Options) string {
	if opts == nil {
		opts = &DefaultOptions
	}

	// 标准化邮箱：转小写并去除空格
	email = strings.TrimSpace(strings.ToLower(email))

	// 计算 MD5 哈希
	hash := md5.Sum([]byte(email))
	emailHash := hex.EncodeToString(hash[:])

	// 构建 URL
	url := fmt.Sprintf("https://www.gravatar.com/avatar/%s", emailHash)

	// 添加参数
	params := make([]string, 0)
	
	if opts.Size > 0 {
		if opts.Size > 2048 {
			opts.Size = 2048
		}
		params = append(params, fmt.Sprintf("s=%d", opts.Size))
	}
	
	if opts.Default != "" {
		params = append(params, fmt.Sprintf("d=%s", opts.Default))
	}
	
	if opts.Rating != "" {
		params = append(params, fmt.Sprintf("r=%s", opts.Rating))
	}
	
	if opts.ForceDefault {
		params = append(params, "f=y")
	}

	if len(params) > 0 {
		url += "?" + strings.Join(params, "&")
	}

	return url
}

// GetAvatarURLSimple 获取头像 URL (简化版本，使用默认选项)
func GetAvatarURLSimple(email string) string {
	return GetAvatarURL(email, &DefaultOptions)
}

// GetAvatarURLWithSize 获取指定尺寸的头像 URL
func GetAvatarURLWithSize(email string, size int) string {
	opts := DefaultOptions
	opts.Size = size
	return GetAvatarURL(email, &opts)
}

// GetEmailHash 获取邮箱的 MD5 哈希值
func GetEmailHash(email string) string {
	email = strings.TrimSpace(strings.ToLower(email))
	hash := md5.Sum([]byte(email))
	return hex.EncodeToString(hash[:])
}
