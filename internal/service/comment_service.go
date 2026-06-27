package service

import (
	"blog/internal/model/dto/request"
	"blog/internal/model/dto/response"
	"blog/internal/model/entity"
	"blog/internal/repository"
	bizerrors "blog/pkg/errors"
	"blog/pkg/email"
	"blog/pkg/logger"
)

// commentService 评论服务实现
type commentService struct {
	commentRepo repository.CommentRepository
	articleRepo repository.ArticleRepository
	userRepo    repository.UserRepository
	emailSvc    email.EmailService
}

// NewCommentService 创建评论服务
func NewCommentService(commentRepo repository.CommentRepository, articleRepo repository.ArticleRepository, userRepo repository.UserRepository, emailSvc email.EmailService) CommentService {
	return &commentService{
		commentRepo: commentRepo,
		articleRepo: articleRepo,
		userRepo:    userRepo,
		emailSvc:    emailSvc,
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

	// 异步发送邮件通知
	go s.sendEmailNotifications(comment, article)

	return comment.ID, nil
}

// sendEmailNotifications 发送邮件通知
func (s *commentService) sendEmailNotifications(comment *entity.Comment, article *entity.Article) {
	// 获取博主信息（假设博主是第一个用户）
	adminUsers, _, err := s.userRepo.List(0, 1)
	if err != nil || len(adminUsers) == 0 {
		logger.Error("获取博主信息失败: %v", err)
		return
	}
	admin := adminUsers[0]

	// 如果是回复评论，发送邮件给被回复的用户
	if comment.ParentID > 0 {
		// 获取被回复的评论
		parentComment, err := s.commentRepo.FindByID(comment.ParentID)
		if err != nil || parentComment == nil {
			logger.Error("获取被回复评论失败: %v", err)
			return
		}

		// 如果被回复用户有邮箱，发送回复通知
		if parentComment.Email != "" {
			err = s.emailSvc.SendReplyNotification(
				parentComment.Email,
				comment.Nickname,
				article.Title,
				article.Slug,
				comment.Content,
			)
			if err != nil {
				logger.Error("发送回复通知邮件失败: %v", err)
			}
		}
	} else {
		// 如果是评论文章，发送邮件给博主
		if admin.Email != "" {
			err = s.emailSvc.SendCommentNotification(
				admin.Email,
				comment.Nickname,
				article.Title,
				article.Slug,
				comment.Content,
			)
			if err != nil {
				logger.Error("发送评论通知邮件失败: %v", err)
			}
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
