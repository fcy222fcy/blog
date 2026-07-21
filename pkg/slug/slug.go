package slug

import (
	"regexp"
	"strings"
	"time"
)

var (
	// 非字母数字字符的正则
	nonAlphanumericRegex = regexp.MustCompile(`[^a-zA-Z0-9]+`)
	// 多个连字符的正则
	multipleHyphensRegex = regexp.MustCompile(`-{2,}`)
)

// Generate 根据名称生成 slug
// 如果名称为空或只包含非 ASCII 字符，则生成基于时间戳的 slug
func Generate(name string) string {
	if name == "" {
		return generateTimestampSlug()
	}

	// 转换为小写
	slug := strings.ToLower(name)

	// 尝试提取 ASCII 字符
	slug = extractASCII(slug)

	// 如果提取后为空，说明原名称全是非 ASCII 字符（如中文）
	if slug == "" {
		return generateTimestampSlug()
	}

	// 替换非字母数字字符为连字符
	slug = nonAlphanumericRegex.ReplaceAllString(slug, "-")

	// 移除首尾连字符
	slug = strings.Trim(slug, "-")

	// 合并多个连字符
	slug = multipleHyphensRegex.ReplaceAllString(slug, "-")

	if slug == "" {
		return generateTimestampSlug()
	}

	return slug
}

// extractASCII 从字符串中提取 ASCII 字符
func extractASCII(s string) string {
	var result []byte
	for _, r := range s {
		if r >= 32 && r <= 126 {
			result = append(result, byte(r))
		}
	}
	return string(result)
}

// generateTimestampSlug 生成基于时间戳的 slug
func generateTimestampSlug() string {
	return time.Now().Format("20060102150405")
}
