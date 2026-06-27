package email

import (
	"blog/pkg/config"
	"crypto/tls"
	"encoding/base64"
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

// encodeHeader RFC 2047 编码中文标题
func encodeHeader(header string) string {
	return fmt.Sprintf("=?UTF-8?B?%s?=", base64.StdEncoding.EncodeToString([]byte(header)))
}

// sendMail 发送邮件
func (s *emailService) sendMail(to, subject, body string) error {
	// 构建邮件头
	var message strings.Builder
	message.WriteString(fmt.Sprintf("From: %s <%s>\r\n", encodeHeader(s.config.From), s.config.FromEmail))
	message.WriteString(fmt.Sprintf("To: %s\r\n", to))
	message.WriteString(fmt.Sprintf("Subject: %s\r\n", encodeHeader(subject)))
	message.WriteString("MIME-Version: 1.0\r\n")
	message.WriteString("Content-Type: text/html; charset=UTF-8\r\n")
	message.WriteString("\r\n")
	message.WriteString(body)

	// 认证
	auth := smtp.PlainAuth("", s.config.Username, s.config.Password, s.config.Host)

	// 使用 TLS 连接
	addr := fmt.Sprintf("%s:%d", s.config.Host, s.config.Port)
	conn, err := tls.Dial("tcp", addr, &tls.Config{
		InsecureSkipVerify: true,
		ServerName:         s.config.Host,
	})
	if err != nil {
		return fmt.Errorf("TLS连接失败: %w", err)
	}
	defer conn.Close()

	client, err := smtp.NewClient(conn, s.config.Host)
	if err != nil {
		return fmt.Errorf("创建SMTP客户端失败: %w", err)
	}
	defer client.Close()

	// 认证
	if err = client.Auth(auth); err != nil {
		return fmt.Errorf("SMTP认证失败: %w", err)
	}

	// 设置发件人
	if err = client.Mail(s.config.FromEmail); err != nil {
		return fmt.Errorf("设置发件人失败: %w", err)
	}

	// 设置收件人
	if err = client.Rcpt(to); err != nil {
		return fmt.Errorf("设置收件人失败: %w", err)
	}

	// 发送邮件
	w, err := client.Data()
	if err != nil {
		return fmt.Errorf("获取数据通道失败: %w", err)
	}

	if _, err = w.Write([]byte(message.String())); err != nil {
		return fmt.Errorf("写入邮件内容失败: %w", err)
	}

	if err = w.Close(); err != nil {
		return fmt.Errorf("关闭数据通道失败: %w", err)
	}

	return nil
}
