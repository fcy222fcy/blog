package service

import (
	"blog/internal/model/dto/request"
	"blog/internal/model/dto/response"
	"blog/internal/model/entity"
	"blog/internal/repository"
	"blog/pkg/config"
	"blog/pkg/email"
	bizerrors "blog/pkg/errors"
	"blog/pkg/gravatar"
	"blog/pkg/logger"
	"blog/pkg/ua"
	"fmt"
	"strconv"
	"strings"
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

// isBlogger 判断 userID 是否为博主（硬编码配置，不查用户表）
func (s *commentService) isBlogger(userID *uint) bool {
	if s.config == nil || userID == nil || *userID == 0 {
		return false
	}
	return *userID == s.config.Blogger.UserID
}

// convertToCommentResponse 将 entity.Comment 转换为 response.CommentResponse
// 博主身份判断：UserID 等于配置中博主虚拟 ID 即标记为博主
func (s *commentService) convertToCommentResponse(comment *entity.Comment) response.CommentResponse {
	resp := response.CommentResponse{
		ID:        comment.ID,
		Content:   comment.Content,
		Nickname:  comment.Nickname,
		Email:     comment.Email,
		Website:   comment.Website,
		Avatar:    comment.Avatar,
		Status:    comment.Status,
		LikeCount: comment.LikeCount,
		ParentID:  comment.ParentID,
		CreatedAt: comment.CreatedAt,
		IsAdmin:   s.isBlogger(comment.UserID),
	}

	if resp.Avatar == "" && resp.Email != "" {
		resp.Avatar = gravatar.GetAvatarURLByEmail(resp.Email, 80)
	}

	if comment.ReplyToNickname != "" {
		resp.ReplyTo = comment.ReplyToNickname
	}

	// 客户端信息优先直接取已经存好的列（前端精确检测/后端兜底），
	// 历史空值数据才回退到解析 User-Agent
	hasStoredInfo := (comment.OS != "" && comment.OS != "未知") || (comment.Browser != "" && comment.Browser != "未知")
	if hasStoredInfo {
		if comment.OS != "未知" {
			resp.OS = comment.OS
			resp.OSVersion = comment.OSVersion
		}
		if comment.Browser != "未知" {
			resp.Browser = comment.Browser
			resp.BrowserVersion = comment.BrowserVersion
		}
	} else if strings.TrimSpace(comment.UserAgent) != "" {
		uaInfo := ua.Parse(comment.UserAgent)
		if uaInfo.OS != "未知" {
			resp.OS = uaInfo.OS
			resp.OSVersion = uaInfo.OSVersion
		}
		if uaInfo.Browser != "未知" {
			resp.Browser = uaInfo.Browser
			resp.BrowserVersion = uaInfo.BrowserVersion
		}
	}

	return resp
}

// buildCommentTree 构建评论响应树并注入博主标识
func (s *commentService) buildCommentResponseList(rootComments []*entity.Comment) []response.CommentResponse {
	result := make([]response.CommentResponse, 0, len(rootComments))
	for _, root := range rootComments {
		rootResp := s.convertToCommentResponse(root)
		if len(root.Replies) > 0 {
			repliesResp := make([]response.CommentResponse, 0, len(root.Replies))
			for i := range root.Replies {
				reply := &root.Replies[i]
				replyResp := s.convertToCommentResponse(reply)
				repliesResp = append(repliesResp, replyResp)
			}
			rootResp.Replies = repliesResp
		}
		result = append(result, rootResp)
	}
	return result
}

// GetCommentsByArticle 获取文章评论列表（支持 slug 或数字 ID）
func (s *commentService) GetCommentsByArticle(articleParam string, req *request.CommentListRequest) (*response.PageResponse, error) {
	if req.Page <= 0 {
		req.Page = 1
	}
	if req.PageSize <= 0 {
		req.PageSize = 10
	}

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

	sortBy := strings.ToLower(strings.TrimSpace(req.SortBy))
	if sortBy != "asc" && sortBy != "desc" && sortBy != "hot" {
		sortBy = "desc"
	}

	list, total, err := s.commentRepo.ListByArticleID(articleID, (req.Page-1)*req.PageSize, req.PageSize, sortBy)
	if err != nil {
		return nil, fmt.Errorf("获取文章评论列表失败, %w", err)
	}

	respList := s.buildCommentResponseList(list)

	return response.NewPageResponse(respList, total, req.Page, req.PageSize), nil
}

// CreateComment 创建评论
// userID: 当前登录用户 ID，0 表示访客
// 博主身份判定条件：userID > 0（已通过登录按钮登录）
func (s *commentService) CreateComment(req *request.CreateCommentRequest, userID uint, ip, userAgent string) (uint, error) {
	if strings.TrimSpace(req.Content) == "" {
		return 0, bizerrors.New(bizerrors.CodeInvalidParams, "评论内容不能为空")
	}

	if strings.TrimSpace(req.Nickname) == "" {
		return 0, bizerrors.New(bizerrors.CodeInvalidParams, "昵称不能为空")
	}

	var article *entity.Article
	var err error
	if req.ArticleSlug != "" {
		article, err = s.articleRepo.FindBySlug(req.ArticleSlug)
	} else if req.ArticleID > 0 {
		article, err = s.articleRepo.FindByID(req.ArticleID)
	} else {
		return 0, bizerrors.New(bizerrors.CodeArticleNotFound, bizerrors.GetMessage(bizerrors.CodeArticleNotFound))
	}
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
		ArticleID: article.ID,
		ParentID:  req.ParentID,
		Status:    "approved",
		IP:        ip,
		UserAgent: userAgent,
	}

	// 客户端信息：优先使用前端通过 JS 精确检测的结果（能区分 Win11/Win10），
	// 空值才用后端 UA 解析兜底
	uaInfo := ua.Parse(userAgent)
	switch {
	case strings.TrimSpace(req.OS) != "" && req.OS != "未知":
		comment.OS = strings.TrimSpace(req.OS)
		comment.OSVersion = strings.TrimSpace(req.OSVersion)
	default:
		if uaInfo.OS != "未知" {
			comment.OS = uaInfo.OS
			comment.OSVersion = uaInfo.OSVersion
		}
	}
	switch {
	case strings.TrimSpace(req.Browser) != "" && req.Browser != "未知":
		comment.Browser = strings.TrimSpace(req.Browser)
		comment.BrowserVersion = strings.TrimSpace(req.BrowserVersion)
	default:
		if uaInfo.Browser != "未知" {
			comment.Browser = uaInfo.Browser
			comment.BrowserVersion = uaInfo.BrowserVersion
		}
	}

	if userID > 0 {
		// 博主账号：直接用配置中的信息（不查用户表）
		if s.config != nil && userID == s.config.Blogger.UserID {
			b := s.config.Blogger
			comment.UserID = &userID
			if comment.Nickname == "" {
				comment.Nickname = b.Nickname
			}
			if comment.Email == "" {
				comment.Email = b.Email
			}
			if comment.Avatar == "" {
				comment.Avatar = b.Avatar
			}
		} else {
			// 其他登录用户：查用户表
			user, err := s.userRepo.FindByID(userID)
			if err == nil && user != nil {
				comment.UserID = &userID
				if comment.Nickname == "" {
					comment.Nickname = user.Nickname
				}
				if comment.Email == "" {
					comment.Email = user.Email
				}
				if comment.Avatar == "" {
					comment.Avatar = user.Avatar
				}
			}
		}
	}

	if comment.Avatar == "" && comment.Email != "" {
		comment.Avatar = gravatar.GetAvatarURLByEmail(comment.Email, 80)
	}

	err = s.commentRepo.Create(comment)
	if err != nil {
		return 0, fmt.Errorf("创建评论失败, %w", err)
	}

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

// LikeComment 点赞评论（防重复）
func (s *commentService) LikeComment(commentID uint, visitorIP string) error {
	// 检查评论是否存在
	comment, err := s.commentRepo.FindByID(commentID)
	if err != nil {
		return fmt.Errorf("查询评论失败, %w", err)
	}
	if comment == nil {
		return bizerrors.New(bizerrors.CodeCommentNotFound, bizerrors.GetMessage(bizerrors.CodeCommentNotFound))
	}

	// 检查是否已点赞（防重复）
	hasLiked, err := s.commentRepo.HasLiked(commentID, visitorIP)
	if err != nil {
		return fmt.Errorf("查询点赞记录失败, %w", err)
	}
	if hasLiked {
		return bizerrors.New(bizerrors.CodeCommentAlreadyLiked, bizerrors.GetMessage(bizerrors.CodeCommentAlreadyLiked))
	}

	// 增加点赞数
	if err := s.commentRepo.IncrementLikeCount(commentID); err != nil {
		return fmt.Errorf("增加点赞数失败, %w", err)
	}

	// 记录点赞日志
	likeLog := &entity.CommentLikeLog{
		CommentID: commentID,
		VisitorIP: visitorIP,
	}
	if err := s.commentRepo.CreateLikeLog(likeLog); err != nil {
		return fmt.Errorf("记录点赞日志失败, %w", err)
	}

	return nil
}
