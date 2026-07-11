package media

import (
	"blog/pkg/response"
	"fmt"
	"net/http"
	"os"
	"path/filepath"
	"strconv"
	"strings"
	"time"

	"github.com/gin-gonic/gin"
)

const (
	BaseUploadDir = "./uploads"
)

var allowedCategories = map[string]bool{
	"avatar":      true,
	"article":     true,
	"daily":       true,
	"entertainment": true,
	"link":        true,
	"common":      true,
}

func normalizeCategory(cat string) string {
	cat = strings.TrimSpace(strings.ToLower(cat))
	if !allowedCategories[cat] {
		return "common"
	}
	return cat
}

// Controller 媒体控制器
type Controller struct{}

// NewController 创建媒体控制器
func NewController() *Controller {
	return &Controller{}
}

// Upload 上传文件
// category: avatar(主页头像/基础配置) | article(文章封面/插图) | daily(每日一问) | entertainment(娱乐) | link(友链) | common(默认)
func (c *Controller) Upload(ctx *gin.Context) {
	file, err := ctx.FormFile("file")
	if err != nil {
		response.BadRequest(ctx, "请选择文件")
		return
	}

	category := normalizeCategory(ctx.PostForm("category"))

	// 验证文件大小 (5MB)
	if file.Size > 5*1024*1024 {
		response.BadRequest(ctx, "文件大小不能超过 5MB")
		return
	}

	// 验证文件扩展名
	allowedExtensions := map[string]bool{
		".jpg": true, ".jpeg": true, ".png": true,
		".gif": true, ".webp": true, ".pdf": true,
	}
	ext := strings.ToLower(filepath.Ext(file.Filename))
	if !allowedExtensions[ext] {
		response.BadRequest(ctx, "不支持的文件类型")
		return
	}

	// 验证 Content-Type（防止扩展名伪造）
	if fileHeader, err := file.Open(); err == nil {
		buf := make([]byte, 512)
		n, _ := fileHeader.Read(buf)
		contentType := http.DetectContentType(buf[:n])
		fileHeader.Close()
		// 对于图片文件，额外校验 Content-Type
		if ext != ".pdf" && !strings.HasPrefix(contentType, "image/") {
			response.BadRequest(ctx, "文件类型与内容不匹配")
			return
		}
	}

	// 生成分类目录
	uploadDir := filepath.Join(BaseUploadDir, category)
	if err := os.MkdirAll(uploadDir, 0755); err != nil {
		response.ServerError(ctx, "创建上传目录失败")
		return
	}

	// 生成唯一文件名
	filename := fmt.Sprintf("%d%s", time.Now().UnixNano(), ext)
	savePath := filepath.Join(uploadDir, filename)

	if err := ctx.SaveUploadedFile(file, savePath); err != nil {
		response.ServerError(ctx, "保存文件失败")
		return
	}

	// 返回文件URL（包含分类子目录，前端直接可访问）
	url := fmt.Sprintf("/uploads/%s/%s", category, filename)
	response.Success(ctx, gin.H{
		"url":      url,
		"category": category,
		"filename": file.Filename,
		"size":     file.Size,
	})
}

// List 获取媒体列表（支持按分类过滤，自动递归扫描分类子目录+兼容根目录旧文件）
func (c *Controller) List(ctx *gin.Context) {
	// 可选分类过滤
	filterCat := strings.TrimSpace(strings.ToLower(ctx.Query("category")))
	if filterCat != "" && !allowedCategories[filterCat] {
		filterCat = ""
	}

	// 获取分页参数
	page, _ := strconv.Atoi(ctx.DefaultQuery("page", "1"))
	pageSize, _ := strconv.Atoi(ctx.DefaultQuery("page_size", "20"))
	keyword := ctx.Query("keyword")

	if page < 1 {
		page = 1
	}
	if pageSize < 1 || pageSize > 50 {
		pageSize = 20
	}

	// 扫描上传目录（递归），构造所有文件列表
	var mediaList []gin.H
	baseDir := BaseUploadDir

	// 先检查根目录是否存在
	if _, err := os.Stat(baseDir); os.IsNotExist(err) {
		response.Success(ctx, gin.H{
			"list":  []interface{}{},
			"total": 0,
		})
		return
	}

	// 收集文件：{相对路径, 分类, 文件名, FileInfo}
	type fileEntry struct {
		relPath  string // 相对 uploads 的路径，如 "avatar/123.jpg" 或 "123.jpg"（旧文件）
		category string
		name     string
		info     os.FileInfo
	}
	var entries []fileEntry

	_ = filepath.Walk(baseDir, func(path string, info os.FileInfo, walkErr error) error {
		if walkErr != nil || info == nil || info.IsDir() {
			return nil
		}

		// 求相对于 baseDir 的路径，使用 filepath.ToSlash 统一分隔符
		rel, err := filepath.Rel(baseDir, path)
		if err != nil {
			return nil
		}
		rel = filepath.ToSlash(rel)

		// 判断分类：有目录前缀的取第一段，根目录文件归为 common（兼容旧数据）
		var cat, fname string
		if idx := strings.Index(rel, "/"); idx >= 0 {
			cat = strings.ToLower(rel[:idx])
			fname = rel[idx+1:]
			if !allowedCategories[cat] {
				// 不合法分类目录下的文件跳过
				return nil
			}
		} else {
			cat = "common"
			fname = rel
		}

		// 分类过滤
		if filterCat != "" && cat != filterCat {
			return nil
		}

		entries = append(entries, fileEntry{
			relPath:  rel,
			category: cat,
			name:     fname,
			info:     info,
		})
		return nil
	})

	// 构建响应
	for _, e := range entries {
		// 关键词过滤
		if keyword != "" && !strings.Contains(strings.ToLower(e.name), strings.ToLower(keyword)) {
			continue
		}

		mediaList = append(mediaList, gin.H{
			"id":       len(mediaList) + 1,
			"name":     e.name,
			"category": e.category,
			"url":      "/uploads/" + e.relPath,
			"size":     e.info.Size(),
			"type":     getMimeType(e.name),
			"modified": e.info.ModTime().Format(time.RFC3339),
		})
	}

	// 分页
	total := len(mediaList)
	start := (page - 1) * pageSize
	end := start + pageSize
	if start > total {
		start = total
	}
	if end > total {
		end = total
	}

	response.Success(ctx, gin.H{
		"list":      mediaList[start:end],
		"total":     total,
		"page":      page,
		"page_size": pageSize,
		"category":  filterCat,
	})
}

// Delete 删除媒体文件
// 路由：DELETE /api/v1/media/:filename
// 规则：
//   - 新格式：通过 query 传 category（如 ?category=avatar），filename 为纯文件名
//     目标文件 = uploads/<category>/<filename>
//   - 兼容格式：不传 category，filename 允许 "分类/文件名" 或 "文件名"
//     - 含 "/"：自动解析第一段为分类
//     - 不含 "/"：视为旧根目录文件，分类 = common
func (c *Controller) Delete(ctx *gin.Context) {
	filenameParam := strings.TrimSpace(ctx.Param("filename"))
	if filenameParam == "" {
		response.BadRequest(ctx, "文件名不能为空")
		return
	}

	category := normalizeCategory(ctx.Query("category"))

	// 若 query 未传合法分类，则从 filenameParam 中尝试解析 "分类/文件名"
	var fname, relPath string
	if ctx.Query("category") == "" && strings.Contains(filenameParam, "/") {
		// 兼容：filename = "avatar/xxx.jpg"，路径中自带分类
		parts := strings.SplitN(filenameParam, "/", 2)
		parsedCat := normalizeCategory(parts[0])
		fname = parts[1]
		relPath = filepath.ToSlash(filepath.Join(parsedCat, fname))
	} else if ctx.Query("category") != "" {
		// 新格式：category=xxx + filename=纯文件名
		fname = filenameParam
		relPath = filepath.ToSlash(filepath.Join(category, fname))
	} else {
		// 旧数据：根目录裸文件，分类 = common
		fname = filenameParam
		relPath = fname
	}

	// 安全检查：禁止 ".."、绝对路径、反斜杠（经过 join 后仍需校验）
	if strings.Contains(fname, "..") || strings.Contains(relPath, "..") ||
		filepath.IsAbs(relPath) || strings.Contains(relPath, "\\") {
		response.BadRequest(ctx, "非法文件名或分类")
		return
	}

	filePath := filepath.Join(BaseUploadDir, relPath)

	// 检查文件是否存在
	if _, err := os.Stat(filePath); os.IsNotExist(err) {
		response.NotFound(ctx, "文件不存在")
		return
	}

	// 删除文件
	if err := os.Remove(filePath); err != nil {
		response.ServerError(ctx, "删除文件失败")
		return
	}

	response.Success(ctx, nil)
}

// getMimeType 获取文件MIME类型
func getMimeType(filename string) string {
	ext := strings.ToLower(filepath.Ext(filename))
	mimeTypes := map[string]string{
		".jpg":  "image/jpeg",
		".jpeg": "image/jpeg",
		".png":  "image/png",
		".gif":  "image/gif",
		".webp": "image/webp",
		".pdf":  "application/pdf",
	}
	if mime, ok := mimeTypes[ext]; ok {
		return mime
	}
	return "application/octet-stream"
}
