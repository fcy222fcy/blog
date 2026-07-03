package rss

import (
	"blog/internal/repository"
	"fmt"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
)

// Handler RSS 处理器
type Handler struct {
	articleRepo repository.ArticleRepository
}

// NewHandler 创建 RSS 处理器
func NewHandler(articleRepo repository.ArticleRepository) *Handler {
	return &Handler{articleRepo: articleRepo}
}

// GenerateRSS 生成 RSS 订阅
func (h *Handler) GenerateRSS(c *gin.Context) {
	articles, _, err := h.articleRepo.ListPublished(0, 20, 0, 0, "")
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"code":    500,
			"message": "生成 RSS 失败",
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

	rss := `<?xml version="1.0" encoding="UTF-8"?>
<rss version="2.0" xmlns:atom="http://www.w3.org/2005/Atom">
<channel>
  <title>我的博客</title>
  <link>` + siteURL + `</link>
  <description>分享技术、生活与思考</description>
  <language>zh-CN</language>
  <lastBuildDate>` + time.Now().Format(time.RFC1123Z) + `</lastBuildDate>
  <atom:link href="` + siteURL + `/api/v1/feed.xml" rel="self" type="application/rss+xml"/>`

	for _, article := range articles {
		rss += `
  <item>
    <title><![CDATA[` + article.Title + `]]></title>
    <link>` + siteURL + `/post/` + article.Slug + `</link>
    <guid isPermaLink="true">` + siteURL + `/post/` + article.Slug + `</guid>
    <description><![CDATA[` + article.Summary + `]]></description>
    <pubDate>` + article.CreatedAt.Format(time.RFC1123Z) + `</pubDate>
  </item>`
	}

	rss += `
</channel>
</rss>`

	c.Header("Content-Type", "application/rss+xml; charset=utf-8")
	c.String(http.StatusOK, rss)
}

// RegisterRoutes 注册路由
func RegisterRoutes(rg *gin.RouterGroup, handler *Handler) {
	feed := rg.Group("")
	{
		feed.GET("/feed.xml", handler.GenerateRSS)
		feed.GET("/rss", handler.GenerateRSS)
	}
}
