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
	link := fmt.Sprintf("https://wzx.glaty.cn/%s", articleSlug)

	// 获取评论者首字母作为头像
	avatarInitial := string([]rune(nickname)[0])
	if len([]rune(nickname)) > 1 {
		avatarInitial = string([]rune(nickname)[:1])
	}

	body := fmt.Sprintf(`
<!DOCTYPE html>
<html>
<head>
    <meta charset="UTF-8">
    <style>
        body {
            font-family: -apple-system, BlinkMacSystemFont, "Segoe UI", "Droid Sans", "Helvetica Neue", "HarmonyOS Sans SC", sans-serif;
            margin: 0;
            padding: 0;
            background: #f8f7f2;
            line-height: 1.6;
        }
        .container {
            max-width: 600px;
            margin: 0 auto;
            background: #fdfdfb;
            border-radius: 10px;
            overflow: hidden;
            box-shadow: 0 4px 6px rgba(0, 0, 0, 0.1);
        }
        .header {
            padding: 24px;
            background: linear-gradient(135deg, #1B365D 0%%, #202A44 100%%);
            text-align: center;
        }
        .header h1 {
            margin: 0;
            font-size: 20px;
            color: #fff;
            font-weight: 700;
        }
        .content {
            padding: 32px;
        }
        .avatar-section {
            text-align: center;
            margin-bottom: 24px;
        }
        .avatar-circle {
            width: 80px;
            height: 80px;
            border-radius: 50%%;
            background: linear-gradient(135deg, rgba(27, 54, 93, 0.1), rgba(27, 54, 93, 0.2));
            display: inline-flex;
            align-items: center;
            justify-content: center;
            margin: 0 auto 16px;
            box-shadow: 0 4px 8px rgba(27, 54, 93, 0.3);
        }
        .avatar-circle .initials {
            font-size: 32px;
            font-weight: 700;
            color: #1B365D;
        }
        .info-card {
            background: rgba(27, 54, 93, 0.05);
            border-radius: 10px;
            padding: 20px;
            margin-bottom: 24px;
        }
        .info-row {
            display: flex;
            padding: 12px 0;
            border-bottom: 1px solid rgba(27, 54, 93, 0.1);
        }
        .info-row:last-child {
            border-bottom: none;
        }
        .info-label {
            width: 80px;
            color: #8c8c8c;
            font-size: 14px;
            font-weight: 500;
        }
        .info-value {
            color: #333333;
            font-size: 14px;
            flex: 1;
        }
        .comment-box {
            background: #fdfdfb;
            padding: 20px;
            margin: 20px 0;
            border-radius: 8px;
            border-left: 4px solid #1B365D;
            color: #5d5d5d;
            line-height: 1.7;
            font-size: 15px;
            box-shadow: 0 2px 4px rgba(0, 0, 0, 0.05);
        }
        .btn-section {
            text-align: center;
            margin-top: 28px;
        }
        .btn {
            display: inline-block;
            background: #1B365D;
            color: #fff;
            padding: 12px 32px;
            border-radius: 25px;
            text-decoration: none;
            font-size: 15px;
            font-weight: 600;
            box-shadow: 0 4px 8px rgba(27, 54, 93, 0.3);
            transition: transform 0.2s ease, box-shadow 0.2s ease;
        }
        .btn:hover {
            transform: translateY(-2px);
            box-shadow: 0 6px 12px rgba(27, 54, 93, 0.4);
        }
        .footer {
            padding: 20px;
            background: rgba(27, 54, 93, 0.05);
            text-align: center;
            font-size: 13px;
            color: #8c8c8c;
        }
        .icon {
            display: inline-block;
            margin-right: 6px;
            vertical-align: middle;
        }
    </style>
</head>
<body>
    <div class="container">
        <div class="header">
            <h1>新评论通知</h1>
        </div>
        <div class="content">
            <div class="avatar-section">
                <div class="avatar-circle">
                    <span class="initials">%s</span>
                </div>
                <div style="color: #333333; font-weight: 600; font-size: 16px;">%s</div>
                <div style="color: #8c8c8c; font-size: 13px;">评论了您的文章</div>
            </div>

            <div class="info-card">
                <div class="info-row">
                    <div class="info-label">文章</div>
                    <div class="info-value">%s</div>
                </div>
            </div>

            <div class="comment-box">
                <strong>评论内容：</strong><br>
                %s
            </div>

            <div class="btn-section">
                <a href="%s" class="btn">查看文章</a>
            </div>
        </div>
        <div class="footer">
            此邮件由博客系统自动发送，请勿直接回复
        </div>
    </div>
</body>
</html>`, avatarInitial, nickname, articleTitle, commentContent, link)

	return s.sendMail(to, subject, body)
}

// SendReplyNotification 发送回复通知邮件（通知被回复用户）
func (s *emailService) SendReplyNotification(to, nickname, articleTitle, articleSlug, replyContent string) error {
	subject := fmt.Sprintf("评论回复通知 - %s", articleTitle)
	link := fmt.Sprintf("https://wzx.glaty.cn/%s", articleSlug)

	// 获取回复者首字母作为头像
	avatarInitial := string([]rune(nickname)[0])
	if len([]rune(nickname)) > 1 {
		avatarInitial = string([]rune(nickname)[:1])
	}

	body := fmt.Sprintf(`
<!DOCTYPE html>
<html>
<head>
    <meta charset="UTF-8">
    <style>
        body {
            font-family: -apple-system, BlinkMacSystemFont, "Segoe UI", "Droid Sans", "Helvetica Neue", "HarmonyOS Sans SC", sans-serif;
            margin: 0;
            padding: 0;
            background: #f8f7f2;
            line-height: 1.6;
        }
        .container {
            max-width: 600px;
            margin: 0 auto;
            background: #fdfdfb;
            border-radius: 10px;
            overflow: hidden;
            box-shadow: 0 4px 6px rgba(0, 0, 0, 0.1);
        }
        .header {
            padding: 24px;
            background: linear-gradient(135deg, #1B365D 0%%, #202A44 100%%);
            text-align: center;
        }
        .header h1 {
            margin: 0;
            font-size: 20px;
            color: #fff;
            font-weight: 700;
        }
        .content {
            padding: 32px;
        }
        .avatar-section {
            text-align: center;
            margin-bottom: 24px;
        }
        .avatar-circle {
            width: 80px;
            height: 80px;
            border-radius: 50%%;
            background: linear-gradient(135deg, rgba(27, 54, 93, 0.1), rgba(27, 54, 93, 0.2));
            display: inline-flex;
            align-items: center;
            justify-content: center;
            margin: 0 auto 16px;
            box-shadow: 0 4px 8px rgba(27, 54, 93, 0.3);
        }
        .avatar-circle .initials {
            font-size: 32px;
            font-weight: 700;
            color: #1B365D;
        }
        .info-card {
            background: rgba(27, 54, 93, 0.05);
            border-radius: 10px;
            padding: 20px;
            margin-bottom: 24px;
        }
        .info-row {
            display: flex;
            padding: 12px 0;
            border-bottom: 1px solid rgba(27, 54, 93, 0.1);
        }
        .info-row:last-child {
            border-bottom: none;
        }
        .info-label {
            width: 80px;
            color: #8c8c8c;
            font-size: 14px;
            font-weight: 500;
        }
        .info-value {
            color: #333333;
            font-size: 14px;
            flex: 1;
        }
        .reply-box {
            background: #fdfdfb;
            padding: 20px;
            margin: 20px 0;
            border-radius: 8px;
            border-left: 4px solid #1B365D;
            color: #5d5d5d;
            line-height: 1.7;
            font-size: 15px;
            box-shadow: 0 2px 4px rgba(0, 0, 0, 0.05);
        }
        .btn-section {
            text-align: center;
            margin-top: 28px;
        }
        .btn {
            display: inline-block;
            background: #1B365D;
            color: #fff;
            padding: 12px 32px;
            border-radius: 25px;
            text-decoration: none;
            font-size: 15px;
            font-weight: 600;
            box-shadow: 0 4px 8px rgba(27, 54, 93, 0.3);
            transition: transform 0.2s ease, box-shadow 0.2s ease;
        }
        .btn:hover {
            transform: translateY(-2px);
            box-shadow: 0 6px 12px rgba(27, 54, 93, 0.4);
        }
        .footer {
            padding: 20px;
            background: rgba(27, 54, 93, 0.05);
            text-align: center;
            font-size: 13px;
            color: #8c8c8c;
        }
        .reply-icon {
            display: inline-block;
            width: 24px;
            height: 24px;
            background: #1B365D;
            border-radius: 50%%;
            text-align: center;
            line-height: 24px;
            color: white;
            font-size: 12px;
            margin-right: 8px;
            vertical-align: middle;
        }
    </style>
</head>
<body>
    <div class="container">
        <div class="header">
            <h1>评论回复通知</h1>
        </div>
        <div class="content">
            <div class="avatar-section">
                <div class="avatar-circle">
                    <span class="initials">%s</span>
                </div>
                <div style="color: #333333; font-weight: 600; font-size: 16px;">
                    <span class="reply-icon">↩</span>%s
                </div>
                <div style="color: #8c8c8c; font-size: 13px;">回复了您的评论</div>
            </div>

            <div class="info-card">
                <div class="info-row">
                    <div class="info-label">文章</div>
                    <div class="info-value">%s</div>
                </div>
            </div>

            <div class="reply-box">
                <strong>回复内容：</strong><br>
                %s
            </div>

            <div class="btn-section">
                <a href="%s" class="btn">查看回复</a>
            </div>
        </div>
        <div class="footer">
            此邮件由博客系统自动发送，请勿直接回复
        </div>
    </div>
</body>
</html>`, avatarInitial, nickname, articleTitle, replyContent, link)

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
