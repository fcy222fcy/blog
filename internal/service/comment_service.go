package service

import (
	"blog/internal/model/dto/request"
	"blog/internal/model/dto/response"
	"blog/internal/model/entity"
	"blog/internal/repository"
	bizerrors "blog/pkg/errors"
)

// commentService 评论服务实现
type commentService struct {
	commentRepo repository.CommentRepository
	articleRepo repository.ArticleRepository
}

// NewCommentService 创建评论服务
func NewCommentService(commentRepo repository.CommentRepository, articleRepo repository.ArticleRepository) CommentService {
	return &commentService{
		commentRepo: commentRepo,
		articleRepo: articleRepo,
	}
}

// GetCommentsByArticle 获取文章评论列表
func (s *commentService) GetCommentsByArticle(articleID uint, req *request.CommentListRequest) (*response.PageResponse, error) {
	if req.Page <= 0 {
		req.Page = 1
	}
	if req.PageSize <= 0 {
		req.PageSize = 10
	}

	list, total, err := s.commentRepo.ListByArticleID(articleID, (req.Page-1)*req.PageSize, req.PageSize)
	if err != nil {
		return nil, err
	}

	return response.NewPageResponse(list, total, req.Page, req.PageSize), nil
}

// CreateComment 创建评论
func (s *commentService) CreateComment(req *request.CreateCommentRequest) (uint, error) {
	article, err := s.articleRepo.FindByID(req.ArticleID)
	if err != nil {
		return 0, err
	}
	if article == nil {
		return 0, bizerrors.New(bizerrors.CodeArticleNotFound, "文章不存在")
	}

	comment := &entity.Comment{
		Content:   req.Content,
		Nickname:  req.Nickname,
		Email:     req.Email,
		Website:   req.Website,
		ArticleID: req.ArticleID,
		ParentID:  req.ParentID,
		Status:    "pending",
	}

	err = s.commentRepo.Create(comment)
	if err != nil {
		return 0, err
	}

	return comment.ID, nil
}

// GetAdminCommentList 获取评论列表（后台）
func (s *commentService) GetAdminCommentList(req *request.CommentListRequest) (*response.PageResponse, error) {
	if req.Page <= 0 {
		req.Page = 1
	}
	if req.PageSize <= 0 {
		req.PageSize = 10
	}

	list, total, err := s.commentRepo.AdminList((req.Page-1)*req.PageSize, req.PageSize, req.Status)
	if err != nil {
		return nil, err
	}

	return response.NewPageResponse(list, total, req.Page, req.PageSize), nil
}

// UpdateCommentStatus 更新评论状态
func (s *commentService) UpdateCommentStatus(id uint, status string) error {
	comment, err := s.commentRepo.FindByID(id)
	if err != nil {
		return err
	}
	if comment == nil {
		return bizerrors.New(bizerrors.CodeCommentNotFound, "评论不存在")
	}

	comment.Status = status
	return s.commentRepo.Update(comment)
}

// DeleteComment 删除评论
func (s *commentService) DeleteComment(id uint) error {
	comment, err := s.commentRepo.FindByID(id)
	if err != nil {
		return err
	}
	if comment == nil {
		return bizerrors.New(bizerrors.CodeCommentNotFound, "评论不存在")
	}

	return s.commentRepo.Delete(id)
}

// BatchDeleteComments 批量删除评论
func (s *commentService) BatchDeleteComments(ids []uint) error {
	return s.commentRepo.BatchDelete(ids)
}
