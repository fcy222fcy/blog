package service

import (
	"blog/internal/model/dto/request"
	"blog/internal/model/dto/response"
	"blog/internal/model/entity"
	"blog/internal/repository"
	bizerrors "blog/pkg/errors"
)

// articleService 文章服务实现
type articleService struct {
	articleRepo  repository.ArticleRepository
	categoryRepo repository.CategoryRepository
	tagRepo      repository.TagRepository
}

// NewArticleService 创建文章服务
func NewArticleService(
	articleRepo repository.ArticleRepository,
	categoryRepo repository.CategoryRepository,
	tagRepo repository.TagRepository,
) ArticleService {
	return &articleService{
		articleRepo:  articleRepo,
		categoryRepo: categoryRepo,
		tagRepo:      tagRepo,
	}
}

// GetArticleList 获取文章列表（前台）
func (s *articleService) GetArticleList(req *request.ArticleListRequest) (*response.PageResponse, error) {
	list, total, err := s.articleRepo.ListPublished(req.GetOffset(), req.GetPageSize(), req.Category, req.Tag, req.Keyword)
	if err != nil {
		return nil, err
	}
	return response.NewPageResponse(list, total, req.Page, req.GetPageSize()), nil
}

// GetArticleDetail 获取文章详情（前台）
func (s *articleService) GetArticleDetail(slug string) (*response.ArticleDetailResponse, error) {
	article, err := s.articleRepo.FindBySlug(slug)
	if err != nil {
		return nil, err
	}
	if article == nil {
		return nil, bizerrors.New(bizerrors.CodeArticleNotFound, "文章不存在")
	}

	// 增加浏览量
	_ = s.articleRepo.IncrementViewCount(article.ID)

	return &response.ArticleDetailResponse{
		ID:           article.ID,
		Title:        article.Title,
		Slug:         article.Slug,
		Content:      article.Content,
		Summary:      article.Summary,
		Cover:        article.Cover,
		ViewCount:    article.ViewCount + 1,
		LikeCount:    article.LikeCount,
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
	// 获取所有已发布文章
	articles, _, err := s.articleRepo.ListPublished(0, 1000, 0, 0, "")
	if err != nil {
		return nil, err
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

	// 转换为响应格式
	var result []response.ArchiveResponse
	for year, arts := range yearMap {
		result = append(result, response.ArchiveResponse{
			Year:     year,
			Articles: arts,
		})
	}

	return result, nil
}

// LikeArticle 文章点赞
func (s *articleService) LikeArticle(id uint) (int64, error) {
	likeCount, err := s.articleRepo.IncrementLikeCount(id)
	if err != nil {
		return 0, err
	}
	return likeCount, nil
}

// GetAdminArticleList 获取文章列表（后台）
func (s *articleService) GetAdminArticleList(req *request.ArticleListRequest) (*response.PageResponse, error) {
	list, total, err := s.articleRepo.ListAll(req.GetOffset(), req.GetPageSize(), req.Status, req.Keyword)
	if err != nil {
		return nil, err
	}
	return response.NewPageResponse(list, total, req.Page, req.GetPageSize()), nil
}

// GetAdminArticleDetail 获取文章详情（后台）
func (s *articleService) GetAdminArticleDetail(id uint) (*response.AdminArticleDetailResponse, error) {
	article, err := s.articleRepo.FindByID(id)
	if err != nil {
		return nil, err
	}
	if article == nil {
		return nil, bizerrors.New(bizerrors.CodeArticleNotFound, "文章不存在")
	}

	return &response.AdminArticleDetailResponse{
		ID:      article.ID,
		Title:   article.Title,
		Slug:    article.Slug,
		Content: article.Content,
		Summary: article.Summary,
		Cover:   article.Cover,
		Status:  article.Status,
		IsTop:   article.IsTop,
	}, nil
}

// CreateArticle 创建文章
func (s *articleService) CreateArticle(req *request.CreateArticleRequest) (uint, error) {
	article := &entity.Article{
		Title:      req.Title,
		Content:    req.Content,
		Summary:    req.Summary,
		Cover:      req.Cover,
		CategoryID: req.CategoryID,
		Status:     req.Status,
		IsTop:      req.IsTop,
	}

	if article.Status == "" {
		article.Status = "draft"
	}

	err := s.articleRepo.Create(article)
	if err != nil {
		return 0, err
	}

	// TODO: 处理标签关联

	return article.ID, nil
}

// UpdateArticle 更新文章
func (s *articleService) UpdateArticle(id uint, req *request.UpdateArticleRequest) error {
	article, err := s.articleRepo.FindByID(id)
	if err != nil {
		return err
	}
	if article == nil {
		return bizerrors.New(bizerrors.CodeArticleNotFound, "文章不存在")
	}

	if req.Title != "" {
		article.Title = req.Title
	}
	if req.Content != "" {
		article.Content = req.Content
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

	return s.articleRepo.Update(article)
}

// DeleteArticle 删除文章
func (s *articleService) DeleteArticle(id uint) error {
	article, err := s.articleRepo.FindByID(id)
	if err != nil {
		return err
	}
	if article == nil {
		return bizerrors.New(bizerrors.CodeArticleNotFound, "文章不存在")
	}
	return s.articleRepo.Delete(id)
}

// BatchDeleteArticles 批量删除文章
func (s *articleService) BatchDeleteArticles(ids []uint) error {
	return s.articleRepo.BatchDelete(ids)
}
