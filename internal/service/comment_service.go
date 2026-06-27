package service

import (
	"blog/internal/model/dto/request"
	"blog/internal/model/dto/response"
	"blog/internal/model/entity"
	"blog/internal/repository"
	"blog/pkg/config"
	bizerrors "blog/pkg/errors"
	"blog/pkg/email"
	"blog/pkg/logger"
	"fmt"
	"strconv"
)

// commentService 评论服务实现
type commentService struct {
	commentRepo repository.CommentRepository
	articleRepo repository.ArticleRepository
	userRepo    repository.UserRepository
	emailSvc    email.EmailService
	config      *config.Config
}

// NewCommentService 创建评论服务
func NewCommentService(commentRepo repository.CommentRepository, articleRepo repository.ArticleRepository, userRepo repository.UserRepository, emailSvc email.EmailService, config *config.Config) CommentService {
	return &commentService{
		commentRepo: commentRepo,
		articleRepo: articleRepo,
		userRepo:    userRepo,
		emailSvc:    emailSvc,
		config:      config,
	}
}

// GetCommentsByArticle 获取文章评论列表（支持 slug 或数字 ID）
func (s *commentService) GetCommentsByArticle(articleParam string, req *request.CommentListRequest) (*response.PageResponse, error) {
	if req.Page <= 0 {
		req.Page = 1
	}
	if req.PageSize <= 0 {
		req.PageSize = 10
	}

	// 解析文章参数：先尝试数字 ID，再尝试 slug
	var articleID uint
	if id, err := strconv.ParseUint(articleParam, 10, 32); err == nil {
		articleID = uint(id)
	} else {
		article, err := s.articleRepo.FindBySlug(articleParam)
		if err != nil {
			return nil, fmt.Errorf("根据slug查询文章失败, %w", err)
		}
		if article == nil {
			return nil, bizerrors.New(bizerrors.CodeArticleNotFound, bizerrors.GetMessage(bizerrors.CodeArticleNotFound))
		}
		articleID = article.ID
	}

	list, total, err := s.commentRepo.ListByArticleID(articleID, (req.Page-1)*req.PageSize, req.PageSize)
	if err != nil {
		return nil, fmt.Errorf("获取文章评论列表失败, %w", err)
	}

	return response.NewPageResponse(list, total, req.Page, req.PageSize), nil
}

// CreateComment 创建评论
func (s *commentService) CreateComment(req *request.CreateCommentRequest) (uint, error) {
	article, err := s.articleRepo.FindByID(req.ArticleID)
	if err != nil {
		return 0, fmt.Errorf("查询文章失败, %w", err)
	}
	if article == nil {
		return 0, bizerrors.New(bizerrors.CodeArticleNotFound, bizerrors.GetMessage(bizerrors.CodeArticleNotFound))
	}

	comment := &entity.Comment{
		Content:   req.Content,
		Nickname:  req.Nickname,
		Email:     req.Email,
		Website:   req.Website,
		ArticleID: req.ArticleID,
		ParentID:  req.ParentID,
		Status:    "approved",
	}

	err = s.commentRepo.Create(comment)
	if err != nil {
		return 0, fmt.Errorf("创建评论失败, %w", err)
	}

	// 异步发送邮件通知
	go s.sendEmailNotifications(comment, article)

	return comment.ID, nil
}

// sendEmailNotifications 发送邮件通知
func (s *commentService) sendEmailNotifications(comment *entity.Comment, article *entity.Article) {
	adminEmail := s.config.Email.FromEmail
	logger.Infof("开始发送邮件通知，博主邮箱: %s", adminEmail)

	// 如果是回复评论，发送邮件给被回复的用户
	if comment.ParentID != nil && *comment.ParentID > 0 {
		parentComment, err := s.commentRepo.FindByID(*comment.ParentID)
		if err != nil || parentComment == nil {
			logger.Warnf("获取被回复评论失败, parentID: %d, err: %v", *comment.ParentID, err)
			return
		}

		if parentComment.Email != "" {
			err = s.emailSvc.SendReplyNotification(
				parentComment.Email,
				comment.Nickname,
				article.Title,
				article.Slug,
				comment.Content,
			)
			if err != nil {
				logger.Warnf("发送回复通知邮件失败, to: %s, err: %v", parentComment.Email, err)
			} else {
				logger.Infof("回复通知邮件发送成功, to: %s", parentComment.Email)
			}
		}
	} else {
		if adminEmail != "" {
			err := s.emailSvc.SendCommentNotification(
				adminEmail,
				comment.Nickname,
				article.Title,
				article.Slug,
				comment.Content,
			)
			if err != nil {
				logger.Warnf("发送评论通知邮件失败, to: %s, err: %v", adminEmail, err)
			} else {
				logger.Infof("评论通知邮件发送成功, to: %s", adminEmail)
			}
		} else {
			logger.Warn("博主邮箱为空，无法发送邮件通知")
		}
	}
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
		return nil, fmt.Errorf("获取后台评论列表失败, %w", err)
	}

	return response.NewPageResponse(list, total, req.Page, req.PageSize), nil
}

// UpdateCommentStatus 更新评论状态
func (s *commentService) UpdateCommentStatus(id uint, status string) error {
	comment, err := s.commentRepo.FindByID(id)
	if err != nil {
		return fmt.Errorf("查询评论失败, %w", err)
	}
	if comment == nil {
		return bizerrors.New(bizerrors.CodeCommentNotFound, bizerrors.GetMessage(bizerrors.CodeCommentNotFound))
	}

	comment.Status = status
	if err := s.commentRepo.Update(comment); err != nil {
		return fmt.Errorf("更新评论状态失败, %w", err)
	}
	return nil
}

// DeleteComment 删除评论
func (s *commentService) DeleteComment(id uint) error {
	comment, err := s.commentRepo.FindByID(id)
	if err != nil {
		return fmt.Errorf("查询评论失败, %w", err)
	}
	if comment == nil {
		return bizerrors.New(bizerrors.CodeCommentNotFound, bizerrors.GetMessage(bizerrors.CodeCommentNotFound))
	}

	if err := s.commentRepo.Delete(id); err != nil {
		return fmt.Errorf("删除评论失败, %w", err)
	}
	return nil
}

// BatchDeleteComments 批量删除评论
func (s *commentService) BatchDeleteComments(ids []uint) error {
	if err := s.commentRepo.BatchDelete(ids); err != nil {
		return fmt.Errorf("批量删除评论失败, %w", err)
	}
	return nil
}
