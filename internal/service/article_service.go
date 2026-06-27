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
		return nil, fmt.Errorf("获取文章列表失败, %w", err)
	}
	return response.NewPageResponse(list, total, req.Page, req.GetPageSize()), nil
}

// GetArticleDetail 获取文章详情（前台）
func (s *articleService) GetArticleDetail(slug string) (*response.ArticleDetailResponse, error) {
	article, err := s.articleRepo.FindBySlug(slug)
	if err != nil {
		return nil, fmt.Errorf("获取文章详情失败, %w", err)
	}
	if article == nil {
		return nil, bizerrors.New(bizerrors.CodeArticleNotFound, bizerrors.GetMessage(bizerrors.CodeArticleNotFound))
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

// LikeArticle 文章点赞
func (s *articleService) LikeArticle(id uint) (int64, error) {
	likeCount, err := s.articleRepo.IncrementLikeCount(id)
	if err != nil {
		return 0, fmt.Errorf("文章点赞失败, %w", err)
	}
	return likeCount, nil
}

// GetAdminArticleList 获取文章列表（后台）
func (s *articleService) GetAdminArticleList(req *request.ArticleListRequest) (*response.PageResponse, error) {
	list, total, err := s.articleRepo.ListAll(req.GetOffset(), req.GetPageSize(), req.Status, req.Keyword)
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
		return 0, fmt.Errorf("创建文章失败, %w", err)
	}

	logger.Infof("文章创建成功, id: %d, title: %s", article.ID, article.Title)

	// TODO: 处理标签关联

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
