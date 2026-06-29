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
	link := fmt.Sprintf("wzx.glaty.cn/%s", articleSlug)

	body := fmt.Sprintf(`
<!DOCTYPE html>
<html>
<head>
    <meta charset="UTF-8">
    <style>
        body { font-family: -apple-system, BlinkMacSystemFont, 'Segoe UI', Roboto, sans-serif; margin: 0; padding: 0; background: #fff; }
        .container { max-width: 600px; margin: 0 auto; border: 1px solid #e0e0e0; }
        .header { padding: 20px; border-bottom: 2px solid #333; }
        .header h3 { margin: 0; font-size: 16px; color: #333; font-weight: 600; }
        .content { padding: 24px; }
        .row { display: flex; padding: 10px 0; border-bottom: 1px solid #f0f0f0; }
        .row:last-child { border-bottom: none; }
        .row .label { width: 80px; color: #999; font-size: 13px; }
        .row .value { color: #333; font-size: 14px; }
        .message { background: #fafafa; padding: 16px; margin: 16px 0; border-radius: 4px; color: #555; line-height: 1.6; font-size: 14px; }
        .btn { display: inline-block; background: #333; color: #fff; padding: 10px 24px; border-radius: 4px; text-decoration: none; font-size: 14px; margin-top: 16px; }
        .footer { padding: 16px; background: #fafafa; text-align: center; font-size: 12px; color: #999; border-top: 1px solid #eee; }
    </style>
</head>
<body>
    <div class="container">
        <div class="header">
            <h3>新评论通知</h3>
        </div>
        <div class="content">
            <div class="row">
                <div class="label">评论者</div>
                <div class="value">%s</div>
            </div>
            <div class="row">
                <div class="label">文章</div>
                <div class="value">%s</div>
            </div>
            <div class="message">%s</div>
            <a href="%s" class="btn">查看文章</a>
        </div>
        <div class="footer">此邮件由博客系统自动发送，请勿直接回复</div>
    </div>
</body>
</html>`, nickname, articleTitle, commentContent, link)

	return s.sendMail(to, subject, body)
}

// SendReplyNotification 发送回复通知邮件（通知被回复用户）
func (s *emailService) SendReplyNotification(to, nickname, articleTitle, articleSlug, replyContent string) error {
	subject := fmt.Sprintf("评论回复通知 - %s", articleTitle)
	link := fmt.Sprintf("wzx.glaty.cn/%s", articleSlug)

	body := fmt.Sprintf(`
<!DOCTYPE html>
<html>
<head>
    <meta charset="UTF-8">
    <style>
        body { font-family: -apple-system, BlinkMacSystemFont, 'Segoe UI', Roboto, sans-serif; margin: 0; padding: 0; background: #fff; }
        .container { max-width: 600px; margin: 0 auto; border: 1px solid #e0e0e0; }
        .header { padding: 20px; border-bottom: 2px solid #333; }
        .header h3 { margin: 0; font-size: 16px; color: #333; font-weight: 600; }
        .content { padding: 24px; }
        .row { display: flex; padding: 10px 0; border-bottom: 1px solid #f0f0f0; }
        .row:last-child { border-bottom: none; }
        .row .label { width: 80px; color: #999; font-size: 13px; }
        .row .value { color: #333; font-size: 14px; }
        .message { background: #fafafa; padding: 16px; margin: 16px 0; border-radius: 4px; color: #555; line-height: 1.6; font-size: 14px; }
        .btn { display: inline-block; background: #333; color: #fff; padding: 10px 24px; border-radius: 4px; text-decoration: none; font-size: 14px; margin-top: 16px; }
        .footer { padding: 16px; background: #fafafa; text-align: center; font-size: 12px; color: #999; border-top: 1px solid #eee; }
    </style>
</head>
<body>
    <div class="container">
        <div class="header">
            <h3>评论回复通知</h3>
        </div>
        <div class="content">
            <div class="row">
                <div class="label">回复者</div>
                <div class="value">%s</div>
            </div>
            <div class="row">
                <div class="label">文章</div>
                <div class="value">%s</div>
            </div>
            <div class="message">%s</div>
            <a href="%s" class="btn">查看回复</a>
        </div>
        <div class="footer">此邮件由博客系统自动发送，请勿直接回复</div>
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
