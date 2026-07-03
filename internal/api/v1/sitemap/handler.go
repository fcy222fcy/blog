package sitemap

import (
	"blog/internal/repository"
	"fmt"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
)

// Handler Sitemap 处理器
type Handler struct {
	articleRepo repository.ArticleRepository
}

// NewHandler 创建 Sitemap 处理器
func NewHandler(articleRepo repository.ArticleRepository) *Handler {
	return &Handler{articleRepo: articleRepo}
}

// GenerateSitemap 生成 sitemap.xml
func (h *Handler) GenerateSitemap(c *gin.Context) {
	articles, _, err := h.articleRepo.ListPublished(0, 1000, 0, 0, "")
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"code":    500,
			"message": "生成站点地图失败",
		})
		return
	}

	siteURL := "https://myblog.com"
	if host := c.Request.Host; host != "" {
		scheme := "https"
		if c.Request.TLS == nil {
			scheme = "http"
		}
		siteURL = fmt.Sprintf("%s://%s", scheme, host)
	}

	sitemap := `<?xml version="1.0" encoding="UTF-8"?>
<urlset xmlns="http://www.sitemaps.org/schemas/sitemap/0.9">`

	// 首页
	sitemap += fmt.Sprintf(`
  <url>
    <loc>%s/</loc>
    <lastmod>%s</lastmod>
    <changefreq>daily</changefreq>
    <priority>1.0</priority>
  </url>`, siteURL, time.Now().Format("2006-01-02"))

	// 文章页面
	for _, article := range articles {
		lastmod := article.UpdatedAt.Format("2006-01-02")
		if article.UpdatedAt.IsZero() {
			lastmod = article.CreatedAt.Format("2006-01-02")
		}
		sitemap += fmt.Sprintf(`
  <url>
    <loc>%s/post/%s</loc>
    <lastmod>%s</lastmod>
    <changefreq>weekly</changefreq>
    <priority>0.8</priority>
  </url>`, siteURL, article.Slug, lastmod)
	}

	sitemap += `
</urlset>`

	c.Header("Content-Type", "application/xml; charset=utf-8")
	c.String(http.StatusOK, sitemap)
}

// RegisterRoutes 注册路由
func RegisterRoutes(rg *gin.RouterGroup, handler *Handler) {
	sm := rg.Group("")
	{
		sm.GET("/sitemap.xml", handler.GenerateSitemap)
	}
}
