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

// GetCommentList 获取文章评论列表
func (s *commentService) GetCommentList(articleID uint, req *request.PageRequest) (*response.PageResponse, error) {
	list, total, err := s.commentRepo.ListByArticleID(articleID, req.GetOffset(), req.GetPageSize())
	if err != nil {
		return nil, err
	}
	return response.NewPageResponse(list, total, req.Page, req.GetPageSize()), nil
}

// CreateComment 创建评论
func (s *commentService) CreateComment(articleID uint, req *request.CreateCommentRequest) (uint, error) {
	// 检查文章是否存在
	article, err := s.articleRepo.FindByID(articleID)
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
		ArticleID: articleID,
		ParentID:  req.ParentID,
		Status:    "pending",
	}

	err = s.commentRepo.Create(comment)
	if err != nil {
		return 0, err
	}

	// 更新文章评论数
	_ = s.articleRepo.IncrementCommentCount(articleID)

	return comment.ID, nil
}

// GetAdminCommentList 获取评论列表（后台）
func (s *commentService) GetAdminCommentList(req *request.CommentListRequest) (*response.PageResponse, error) {
	list, total, err := s.commentRepo.ListAll(req.GetOffset(), req.GetPageSize(), req.Status, req.ArticleID)
	if err != nil {
		return nil, err
	}
	return response.NewPageResponse(list, total, req.Page, req.GetPageSize()), nil
}

// UpdateCommentStatus 审核评论
func (s *commentService) UpdateCommentStatus(id uint, req *request.UpdateCommentStatusRequest) error {
	comment, err := s.commentRepo.FindByID(id)
	if err != nil {
		return err
	}
	if comment == nil {
		return bizerrors.New(bizerrors.CodeCommentNotFound, "评论不存在")
	}

	comment.Status = req.Status
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
