package email

import (
	"blog/pkg/config"
	"fmt"
	"net/smtp"
	"strings"
)

// EmailService 邮件服务接口
type EmailService interface {
	SendCommentNotification(to, nickname, articleTitle, articleSlug, commentContent string) error
	SendReplyNotification(to, nickname, articleTitle, articleSlug, replyContent string) error
}

// emailService 邮件服务实现
type emailService struct {
	config config.EmailConfig
}

// NewEmailService 创建邮件服务
func NewEmailService(config config.EmailConfig) EmailService {
	return &emailService{config: config}
}

// SendCommentNotification 发送评论通知邮件（通知博主）
func (s *emailService) SendCommentNotification(to, nickname, articleTitle, articleSlug, commentContent string) error {
	subject := fmt.Sprintf("新评论通知 - %s", articleTitle)
	link := fmt.Sprintf("https://liuhouliang.com/post/%s", articleSlug)

	body := fmt.Sprintf(`
<!DOCTYPE html>
<html>
<head>
    <meta charset="UTF-8">
    <style>
        body { font-family: 'Segoe UI', Tahoma, Geneva, Verdana, sans-serif; line-height: 1.6; color: #333; background-color: #f5f5f5; margin: 0; padding: 20px; }
        .container { max-width: 600px; margin: 0 auto; background-color: #fff; border-radius: 8px; box-shadow: 0 2px 10px rgba(0,0,0,0.1); overflow: hidden; }
        .header { background: linear-gradient(135deg, #667eea 0%%, #764ba2 100%%); color: white; padding: 20px; text-align: center; }
        .header h1 { margin: 0; font-size: 20px; font-weight: 500; }
        .content { padding: 25px; }
        .info-box { background-color: #f8f9fa; border-left: 4px solid #667eea; padding: 15px; margin: 15px 0; border-radius: 0 4px 4px 0; }
        .comment-box { background-color: #fff; border: 1px solid #e9ecef; border-radius: 6px; padding: 15px; margin: 15px 0; }
        .comment-box p { margin: 0; color: #495057; }
        .btn { display: inline-block; background: linear-gradient(135deg, #667eea 0%%, #764ba2 100%%); color: white; text-decoration: none; padding: 12px 30px; border-radius: 25px; margin-top: 20px; font-weight: 500; }
        .btn:hover { opacity: 0.9; }
        .footer { background-color: #f8f9fa; padding: 15px; text-align: center; font-size: 12px; color: #6c757d; border-top: 1px solid #e9ecef; }
    </style>
</head>
<body>
    <div class="container">
        <div class="header">
            <h1>📝 新评论通知</h1>
        </div>
        <div class="content">
            <div class="info-box">
                <strong>%s</strong> 评论了您的文章《%s》
            </div>
            <div class="comment-box">
                <p>%s</p>
            </div>
            <div style="text-align: center;">
                <a href="%s" class="btn">查看文章</a>
            </div>
        </div>
        <div class="footer">
            此邮件由博客系统自动发送，请勿直接回复
        </div>
    </div>
</body>
</html>`, nickname, articleTitle, commentContent, link)

	return s.sendMail(to, subject, body)
}

// SendReplyNotification 发送回复通知邮件（通知被回复用户）
func (s *emailService) SendReplyNotification(to, nickname, articleTitle, articleSlug, replyContent string) error {
	subject := fmt.Sprintf("评论回复通知 - %s", articleTitle)
	link := fmt.Sprintf("https://liuhouliang.com/post/%s", articleSlug)

	body := fmt.Sprintf(`
<!DOCTYPE html>
<html>
<head>
    <meta charset="UTF-8">
    <style>
        body { font-family: 'Segoe UI', Tahoma, Geneva, Verdana, sans-serif; line-height: 1.6; color: #333; background-color: #f5f5f5; margin: 0; padding: 20px; }
        .container { max-width: 600px; margin: 0 auto; background-color: #fff; border-radius: 8px; box-shadow: 0 2px 10px rgba(0,0,0,0.1); overflow: hidden; }
        .header { background: linear-gradient(135deg, #11998e 0%%, #38ef7d 100%%); color: white; padding: 20px; text-align: center; }
        .header h1 { margin: 0; font-size: 20px; font-weight: 500; }
        .content { padding: 25px; }
        .info-box { background-color: #f8f9fa; border-left: 4px solid #11998e; padding: 15px; margin: 15px 0; border-radius: 0 4px 4px 0; }
        .comment-box { background-color: #fff; border: 1px solid #e9ecef; border-radius: 6px; padding: 15px; margin: 15px 0; }
        .comment-box p { margin: 0; color: #495057; }
        .btn { display: inline-block; background: linear-gradient(135deg, #11998e 0%%, #38ef7d 100%%); color: white; text-decoration: none; padding: 12px 30px; border-radius: 25px; margin-top: 20px; font-weight: 500; }
        .btn:hover { opacity: 0.9; }
        .footer { background-color: #f8f9fa; padding: 15px; text-align: center; font-size: 12px; color: #6c757d; border-top: 1px solid #e9ecef; }
    </style>
</head>
<body>
    <div class="container">
        <div class="header">
            <h1>💬 评论回复通知</h1>
        </div>
        <div class="content">
            <div class="info-box">
                <strong>%s</strong> 回复了您在《%s》下的评论
            </div>
            <div class="comment-box">
                <p>%s</p>
            </div>
            <div style="text-align: center;">
                <a href="%s" class="btn">查看回复</a>
            </div>
        </div>
        <div class="footer">
            此邮件由博客系统自动发送，请勿直接回复
        </div>
    </div>
</body>
</html>`, nickname, articleTitle, replyContent, link)

	return s.sendMail(to, subject, body)
}

// sendMail 发送邮件
func (s *emailService) sendMail(to, subject, body string) error {
	// 构建邮件头
	headers := make(map[string]string)
	headers["From"] = fmt.Sprintf("%s <%s>", s.config.From, s.config.FromEmail)
	headers["To"] = to
	headers["Subject"] = subject
	headers["MIME-Version"] = "1.0"
	headers["Content-Type"] = "text/html; charset=UTF-8"

	// 构建邮件内容
	var message strings.Builder
	for k, v := range headers {
		message.WriteString(fmt.Sprintf("%s: %s\r\n", k, v))
	}
	message.WriteString("\r\n")
	message.WriteString(body)

	// 认证
	auth := smtp.PlainAuth("", s.config.Username, s.config.Password, s.config.Host)

	// 发送邮件
	addr := fmt.Sprintf("%s:%d", s.config.Host, s.config.Port)
	return smtp.SendMail(addr, auth, s.config.FromEmail, []string{to}, []byte(message.String()))
}
