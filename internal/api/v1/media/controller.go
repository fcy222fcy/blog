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

// Controller 媒体控制器
type Controller struct{}

// NewController 创建媒体控制器
func NewController() *Controller {
	return &Controller{}
}

// Upload 上传文件
func (c *Controller) Upload(ctx *gin.Context) {
	file, err := ctx.FormFile("file")
	if err != nil {
		response.BadRequest(ctx, "请选择文件")
		return
	}

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

	// 生成上传目录
	uploadDir := "./uploads"
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

	// 返回文件URL
	url := fmt.Sprintf("/uploads/%s", filename)
	response.Success(ctx, gin.H{
		"url":      url,
		"filename": file.Filename,
		"size":     file.Size,
	})
}

// List 获取媒体列表
func (c *Controller) List(ctx *gin.Context) {
	uploadDir := "./uploads"

	// 读取目录
	entries, err := os.ReadDir(uploadDir)
	if err != nil {
		// 目录不存在返回空列表
		response.Success(ctx, gin.H{
			"list":  []interface{}{},
			"total": 0,
		})
		return
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

	// 构建媒体列表
	var mediaList []gin.H
	for _, entry := range entries {
		if entry.IsDir() {
			continue
		}

		name := entry.Name()
		info, _ := entry.Info()
		if info == nil {
			continue
		}

		// 关键词过滤
		if keyword != "" && !strings.Contains(strings.ToLower(name), strings.ToLower(keyword)) {
			continue
		}

		mediaList = append(mediaList, gin.H{
			"id":       len(mediaList) + 1,
			"name":     name,
			"url":      fmt.Sprintf("/uploads/%s", name),
			"size":     info.Size(),
			"type":     getMimeType(name),
			"modified": info.ModTime().Format(time.RFC3339),
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
	})
}

// Delete 删除媒体文件
func (c *Controller) Delete(ctx *gin.Context) {
	filename := ctx.Param("filename")
	if filename == "" {
		response.BadRequest(ctx, "文件名不能为空")
		return
	}

	// 安全检查：防止路径遍历
	if strings.Contains(filename, "..") || strings.Contains(filename, "/") || strings.Contains(filename, "\\") {
		response.BadRequest(ctx, "非法文件名")
		return
	}

	filePath := filepath.Join("./uploads", filename)

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
