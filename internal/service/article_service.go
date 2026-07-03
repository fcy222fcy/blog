package service

import (
	"blog/internal/model/dto/request"
	"blog/internal/model/dto/response"
	"blog/internal/model/entity"
	"blog/internal/repository"
	bizerrors "blog/pkg/errors"
	"blog/pkg/logger"
	"fmt"
	"sort"
	"strings"
	"time"
)

// calculateReadingTime 计算阅读时间（分钟）
// 中文约 400 字/分钟，英文约 200 词/分钟
func calculateReadingTime(content string) int {
	if content == "" {
		return 1
	}
	// 去除 markdown 标记
	cleaned := content
	replacements := []string{"#", "*", "`", ">", "[", "]", "(", ")", "!", "-"}
	for _, r := range replacements {
		cleaned = strings.ReplaceAll(cleaned, r, "")
	}
	cleaned = strings.TrimSpace(cleaned)

	if cleaned == "" {
		return 1
	}

	// 计算中文字符数
	runeCount := len([]rune(cleaned))
	// 计算英文单词数
	wordCount := len(strings.Fields(cleaned))

	// 混合计算：中文字符按 400 字/分钟，英文单词按 200 词/分钟
	minutes := (runeCount / 400) + (wordCount / 200)
	if minutes < 1 {
		return 1
	}
	return minutes
}

// articleService 文章服务实现
type articleService struct {
	articleRepo  repository.ArticleRepository
	categoryRepo repository.CategoryRepository
	tagRepo      repository.TagRepository
	visitRepo    repository.VisitRepository
}

// NewArticleService 创建文章服务
func NewArticleService(
	articleRepo repository.ArticleRepository,
	categoryRepo repository.CategoryRepository,
	tagRepo repository.TagRepository,
	visitRepo repository.VisitRepository,
) ArticleService {
	return &articleService{
		articleRepo:  articleRepo,
		categoryRepo: categoryRepo,
		tagRepo:      tagRepo,
		visitRepo:    visitRepo,
	}
}

// GetArticleList 获取文章列表（前台）
func (s *articleService) GetArticleList(req *request.ArticleListRequest) (*response.PageResponse, error) {
	list, total, err := s.articleRepo.ListPublished(req.GetOffset(), req.GetPageSize(), req.Category, req.Tag, req.Keyword)
	if err != nil {
		return nil, fmt.Errorf("获取文章列表失败, %w", err)
	}
	return response.NewPageResponse(list, total, req.Page, req.GetPageSize()), nil
}

// GetArticleDetail 获取文章详情（前台）
func (s *articleService) GetArticleDetail(slug string, clientIP string) (*response.ArticleDetailResponse, error) {
	article, err := s.articleRepo.FindBySlug(slug)
	if err != nil {
		return nil, fmt.Errorf("获取文章详情失败, %w", err)
	}
	if article == nil {
		return nil, bizerrors.New(bizerrors.CodeArticleNotFound, bizerrors.GetMessage(bizerrors.CodeArticleNotFound))
	}

	// 防刷量：同一 IP 24 小时内只计数一次
	viewIncremented := false
	if clientIP != "" && s.visitRepo != nil {
		since := time.Now().Add(-24 * time.Hour)
		hasVisited, _ := s.visitRepo.HasVisited(article.ID, clientIP, since)
		if !hasVisited {
			_ = s.articleRepo.IncrementViewCount(article.ID)
			_ = s.visitRepo.Create(&entity.VisitLog{
				ArticleID: article.ID,
				IP:        clientIP,
			})
			viewIncremented = true
		}
	} else {
		// 无法判断 IP 时仍增加浏览量
		_ = s.articleRepo.IncrementViewCount(article.ID)
		viewIncremented = true
	}

	viewCount := article.ViewCount
	if viewIncremented {
		viewCount++
	}

	return &response.ArticleDetailResponse{
		ID:           article.ID,
		Title:        article.Title,
		Slug:         article.Slug,
		Content:      article.Content,
		Summary:      article.Summary,
		Cover:        article.Cover,
		ViewCount:    viewCount,
		CommentCount: article.CommentCount,
		Status:       article.Status,
		IsTop:        article.IsTop,
		ReadingTime:  article.ReadingTime,
		CreatedAt:    article.CreatedAt,
		UpdatedAt:    article.UpdatedAt,
	}, nil
}

// GetArticleArchives 获取文章归档
func (s *articleService) GetArticleArchives() ([]response.ArchiveResponse, error) {
	// 获取所有已发布文章，按创建时间降序排序
	articles, _, err := s.articleRepo.ListPublished(0, 1000, 0, 0, "")
	if err != nil {
		return nil, fmt.Errorf("获取文章归档失败, %w", err)
	}

	// 按年份分组
	yearMap := make(map[int][]response.ArchiveArticleResponse)
	for _, article := range articles {
		year := article.CreatedAt.Year()
		yearMap[year] = append(yearMap[year], response.ArchiveArticleResponse{
			ID:        article.ID,
			Title:     article.Title,
			Slug:      article.Slug,
			CreatedAt: article.CreatedAt,
		})
	}

	// 转换为响应格式，并按年份降序排序
	var result []response.ArchiveResponse
	for year, arts := range yearMap {
		// 确保每个年份内的文章按创建时间降序排序
		sort.Slice(arts, func(i, j int) bool {
			return arts[i].CreatedAt.After(arts[j].CreatedAt)
		})
		result = append(result, response.ArchiveResponse{
			Year:     year,
			Articles: arts,
		})
	}

	// 按年份降序排序
	sort.Slice(result, func(i, j int) bool {
		return result[i].Year > result[j].Year
	})

	return result, nil
}

// GetAdminArticleList 获取文章列表（后台）
func (s *articleService) GetAdminArticleList(req *request.ArticleListRequest) (*response.PageResponse, error) {
	list, total, err := s.articleRepo.ListAll(req.GetOffset(), req.GetPageSize(), req.Status, req.Keyword, req.Category)
	if err != nil {
		return nil, fmt.Errorf("获取后台文章列表失败, %w", err)
	}
	return response.NewPageResponse(list, total, req.Page, req.GetPageSize()), nil
}

// GetAdminArticleDetail 获取文章详情（后台）
func (s *articleService) GetAdminArticleDetail(id uint) (*response.AdminArticleDetailResponse, error) {
	article, err := s.articleRepo.FindByID(id)
	if err != nil {
		return nil, fmt.Errorf("获取后台文章详情失败, %w", err)
	}
	if article == nil {
		return nil, bizerrors.New(bizerrors.CodeArticleNotFound, bizerrors.GetMessage(bizerrors.CodeArticleNotFound))
	}

	// 从 Tags 中提取 tag IDs
	var tagIDs []uint
	for _, tag := range article.Tags {
		tagIDs = append(tagIDs, tag.ID)
	}

	return &response.AdminArticleDetailResponse{
		ID:         article.ID,
		Title:      article.Title,
		Slug:       article.Slug,
		Content:    article.Content,
		Summary:    article.Summary,
		Cover:      article.Cover,
		Status:     article.Status,
		IsTop:      article.IsTop,
		CategoryID: article.CategoryID,
		TagIDs:     tagIDs,
	}, nil
}

// CreateArticle 创建文章
func (s *articleService) CreateArticle(req *request.CreateArticleRequest) (uint, error) {
	// 查询标签
	var tags []entity.Tag
	if len(req.TagIDs) > 0 {
		tagList, err := s.tagRepo.FindByIDs(req.TagIDs)
		if err != nil {
			return 0, fmt.Errorf("查询标签失败, %w", err)
		}
		for _, t := range tagList {
			tags = append(tags, *t)
		}
	}

	article := &entity.Article{
		Title:       req.Title,
		Content:     req.Content,
		Summary:     req.Summary,
		Cover:       req.Cover,
		CategoryID:  req.CategoryID,
		Status:      req.Status,
		IsTop:       req.IsTop,
		Tags:        tags,
		ReadingTime: calculateReadingTime(req.Content),
	}

	if article.Status == "" {
		article.Status = entity.ArticleStatusDraft
	}

	err := s.articleRepo.Create(article)
	if err != nil {
		return 0, fmt.Errorf("创建文章失败, %w", err)
	}

	logger.Infof("文章创建成功, id: %d, title: %s, tags: %d", article.ID, article.Title, len(tags))

	return article.ID, nil
}

// UpdateArticle 更新文章
func (s *articleService) UpdateArticle(id uint, req *request.UpdateArticleRequest) error {
	article, err := s.articleRepo.FindByID(id)
	if err != nil {
		return fmt.Errorf("查询文章失败, %w", err)
	}
	if article == nil {
		return bizerrors.New(bizerrors.CodeArticleNotFound, bizerrors.GetMessage(bizerrors.CodeArticleNotFound))
	}

	if req.Title != "" {
		article.Title = req.Title
	}
	if req.Content != "" {
		article.Content = req.Content
		article.ReadingTime = calculateReadingTime(req.Content)
	}
	if req.Summary != "" {
		article.Summary = req.Summary
	}
	if req.Cover != "" {
		article.Cover = req.Cover
	}
	if req.CategoryID != 0 {
		article.CategoryID = req.CategoryID
	}
	if req.Status != "" {
		article.Status = req.Status
	}
	article.IsTop = req.IsTop

	// 更新标签关联
	if req.TagIDs != nil {
		var tags []entity.Tag
		if len(req.TagIDs) > 0 {
			tagList, err := s.tagRepo.FindByIDs(req.TagIDs)
			if err != nil {
				return fmt.Errorf("查询标签失败, %w", err)
			}
			for _, t := range tagList {
				tags = append(tags, *t)
			}
		}
		// 清除旧关联，建立新关联
		if err := s.articleRepo.UpdateTags(article, tags); err != nil {
			return fmt.Errorf("更新文章标签失败, %w", err)
		}
	}

	if err := s.articleRepo.Update(article); err != nil {
		return fmt.Errorf("更新文章失败, %w", err)
	}

	logger.Infof("文章更新成功, id: %d", id)
	return nil
}

// DeleteArticle 删除文章
func (s *articleService) DeleteArticle(id uint) error {
	article, err := s.articleRepo.FindByID(id)
	if err != nil {
		return fmt.Errorf("查询文章失败, %w", err)
	}
	if article == nil {
		return bizerrors.New(bizerrors.CodeArticleNotFound, bizerrors.GetMessage(bizerrors.CodeArticleNotFound))
	}

	if err := s.articleRepo.Delete(id); err != nil {
		return fmt.Errorf("删除文章失败, %w", err)
	}

	logger.Infof("文章删除成功, id: %d", id)
	return nil
}

// BatchDeleteArticles 批量删除文章
func (s *articleService) BatchDeleteArticles(ids []uint) error {
	if err := s.articleRepo.BatchDelete(ids); err != nil {
		return fmt.Errorf("批量删除文章失败, %w", err)
	}

	logger.Infof("批量删除文章成功, ids: %v", ids)
	return nil
}

// Search 搜索文章（标题、内容、摘要模糊搜索）
func (s *articleService) Search(keyword string, page, pageSize int) (*response.PageResponse, error) {
	if page <= 0 {
		page = 1
	}
	if pageSize <= 0 {
		pageSize = 10
	}
	if pageSize > 100 {
		pageSize = 100
	}

	offset := (page - 1) * pageSize
	articles, total, err := s.articleRepo.Search(keyword, offset, pageSize)
	if err != nil {
		return nil, fmt.Errorf("搜索文章失败, %w", err)
	}

	return response.NewPageResponse(articles, total, page, pageSize), nil
}
