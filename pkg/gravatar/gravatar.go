package gravatar

import (
	"crypto/md5"
	"encoding/hex"
	"fmt"
	"regexp"
	"strings"
)

// qqEmailRegexp 匹配 QQ 邮箱（支持 @qq.com 和 @vip.qq.com）
var qqEmailRegexp = regexp.MustCompile(`^(\d+)@(?:vip\.)?qq\.com$`)

// qqSupportedSizes qlogo.cn 官方实测支持的尺寸列表（升序）
// 经实测不支持 50/60/64/80/120 等尺寸（会返回 400）
var qqSupportedSizes = []int{40, 100, 140, 240, 640}

// normalizeQQSize 将任意 size 映射到 qlogo.cn 支持的最近的不小于该值的尺寸
// 若超过最大支持值则返回最大值
func normalizeQQSize(size int) int {
	if size <= 0 {
		return 100
	}
	for _, s := range qqSupportedSizes {
		if s >= size {
			return s
		}
	}
	return qqSupportedSizes[len(qqSupportedSizes)-1]
}

// GetQQAvatarURL 根据 QQ 号获取 QQ 官方头像
// size: 任意整数尺寸，会自动映射到 qlogo.cn 支持的最近尺寸（40/100/140/240/640）
func GetQQAvatarURL(qqNumber string, size int) string {
	qqNumber = strings.TrimSpace(qqNumber)
	if qqNumber == "" {
		return ""
	}
	normalizedSize := normalizeQQSize(size)
	return fmt.Sprintf("https://q1.qlogo.cn/g?b=qq&nk=%s&s=%d", qqNumber, normalizedSize)
}

// isQQEmail 判断是否为 QQ 邮箱，返回 QQ 号和 true 或空串和 false
func isQQEmail(email string) (string, bool) {
	email = strings.TrimSpace(strings.ToLower(email))
	matches := qqEmailRegexp.FindStringSubmatch(email)
	if len(matches) >= 2 {
		return matches[1], true
	}
	return "", false
}

// GetAvatarURLByEmail 智能获取头像：
// 1. QQ 邮箱 → 使用 QQ 官方头像（qlogo.cn）
// 2. 其他邮箱 → 使用 Gravatar
func GetAvatarURLByEmail(email string, size int) string {
	email = strings.TrimSpace(strings.ToLower(email))
	if email == "" {
		return ""
	}
	if qq, ok := isQQEmail(email); ok {
		if url := GetQQAvatarURL(qq, size); url != "" {
			return url
		}
	}
	return GetAvatarURLWithSize(email, size)
}

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
	Size         int         // 图片大小 (1-2048)，默认 80
	Default      DefaultType // 默认图片类型
	Rating       Rating      // 图片评级
	ForceDefault bool        // 强制显示默认头像（用于测试）
}

// DefaultOptions 默认选项
var DefaultOptions = Options{
	Size:         80,
	Default:      DefaultIdenticon, // 使用几何图案作为默认头像
	Rating:       RatingG,
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
