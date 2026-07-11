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

// SendCommentNotification 发送评论通知邮件（通知博主）— Style2 极简线条风
func (s *emailService) SendCommentNotification(to, nickname, articleTitle, articleSlug, commentContent string) error {
	subject := fmt.Sprintf("新评论通知 - %s", articleTitle)
	link := fmt.Sprintf("https://wzx.glaty.cn/%s", articleSlug)

	body := fmt.Sprintf(`
<!DOCTYPE html>
<html>
<head>
    <meta charset="UTF-8">
    <style>
        body {
            font-family: -apple-system, BlinkMacSystemFont, "Segoe UI", "Droid Sans", "Helvetica Neue", "HarmonyOS Sans SC", sans-serif;
            margin: 0;
            padding: 24px 0;
            background: #f8f7f2;
            line-height: 1.6;
        }
        .wrap {
            max-width: 600px;
            margin: 0 auto;
            padding: 0 20px;
        }
        .card {
            background: #fdfdfb;
            border-radius: 10px;
            overflow: hidden;
            box-shadow:
                0px 4px 8px rgba(0,0,0,0.04),
                0px 0px 2px rgba(0,0,0,0.06);
        }
        .eb-header { padding: 28px 32px 0; }
        .eb-brand {
            font-size: 11px;
            color: #1B365D;
            letter-spacing: 3px;
            text-transform: uppercase;
            font-weight: 600;
            margin-bottom: 24px;
            opacity: 0.85;
        }
        .eb-title {
            font-size: 20px;
            color: #1B365D;
            font-weight: 700;
            margin: 0 0 6px 0;
        }
        .eb-sub {
            font-size: 13px;
            color: #8c8c8c;
            margin: 0 0 24px 0;
        }
        .eb-sub strong { color: #1B365D; font-weight: 600; }
        .eb-divider {
            height: 1px;
            background: rgba(0,0,0,0.08);
        }
        .eb-content { padding: 24px 32px 4px; }
        .info-row {
            display: flex;
            padding: 10px 0;
            font-size: 14px;
        }
        .info-row .k {
            width: 72px;
            color: #8c8c8c;
            flex-shrink: 0;
        }
        .info-row .v {
            color: #333333;
            flex: 1;
            font-weight: 500;
        }
        .eb-quote-wrap { padding: 20px 32px; }
        .eb-quote {
            padding: 16px 0;
            border-top: 1px solid rgba(0,0,0,0.06);
            border-bottom: 1px solid rgba(0,0,0,0.06);
            color: #5d5d5d;
            font-size: 14px;
            line-height: 1.8;
        }
        .eb-quote-label {
            font-size: 11px;
            color: #1B365D;
            letter-spacing: 1px;
            margin-bottom: 10px;
            font-weight: 600;
        }
        .eb-action { padding: 8px 32px 28px; }
        .eb-btn {
            display: inline-flex;
            align-items: center;
            gap: 8px;
            color: #1B365D;
            text-decoration: none;
            font-size: 14px;
            font-weight: 600;
            border-bottom: 1px solid #1B365D;
            padding-bottom: 2px;
        }
        .eb-foot {
            padding: 20px 32px 24px;
            font-size: 12px;
            color: #8c8c8c;
            text-align: center;
            background: #faf9f5;
            border-top: 1px solid rgba(0,0,0,0.04);
        }
    </style>
</head>
<body>
    <div class="wrap">
        <div class="card">
            <div class="eb-header">
                <div class="eb-brand">— F c y &nbsp; B L O G —</div>
                <h1 class="eb-title">您收到了一条新评论</h1>
                <p class="eb-sub"><strong>%s</strong> 在您的博客文章下发表了评论</p>
            </div>
            <div class="eb-divider"></div>
            <div class="eb-content">
                <div class="info-row"><span class="k">评论者</span><span class="v">%s</span></div>
                <div class="info-row"><span class="k">文&nbsp;&nbsp;章</span><span class="v">%s</span></div>
            </div>
            <div class="eb-quote-wrap">
                <div class="eb-quote">
                    <div class="eb-quote-label">评 论 内 容</div>
                    %s
                </div>
            </div>
            <div class="eb-action">
                <a href="%s" class="eb-btn">查看评论并回复 →</a>
            </div>
            <div class="eb-foot">此邮件由 Fcy's Blog 自动发送 · 日常落灰的个人博客 · 请勿直接回复</div>
        </div>
    </div>
</body>
</html>`, nickname, nickname, articleTitle, commentContent, link)

	return s.sendMail(to, subject, body)
}

// SendReplyNotification 发送回复通知邮件（通知被回复用户）— Style2 极简线条风
func (s *emailService) SendReplyNotification(to, nickname, articleTitle, articleSlug, replyContent string) error {
	subject := fmt.Sprintf("评论回复通知 - %s", articleTitle)
	link := fmt.Sprintf("https://wzx.glaty.cn/%s", articleSlug)

	body := fmt.Sprintf(`
<!DOCTYPE html>
<html>
<head>
    <meta charset="UTF-8">
    <style>
        body {
            font-family: -apple-system, BlinkMacSystemFont, "Segoe UI", "Droid Sans", "Helvetica Neue", "HarmonyOS Sans SC", sans-serif;
            margin: 0;
            padding: 24px 0;
            background: #f8f7f2;
            line-height: 1.6;
        }
        .wrap {
            max-width: 600px;
            margin: 0 auto;
            padding: 0 20px;
        }
        .card {
            background: #fdfdfb;
            border-radius: 10px;
            overflow: hidden;
            box-shadow:
                0px 4px 8px rgba(0,0,0,0.04),
                0px 0px 2px rgba(0,0,0,0.06);
        }
        .eb-header { padding: 28px 32px 0; }
        .eb-brand {
            font-size: 11px;
            color: #1B365D;
            letter-spacing: 3px;
            text-transform: uppercase;
            font-weight: 600;
            margin-bottom: 24px;
            opacity: 0.85;
        }
        .eb-title {
            font-size: 20px;
            color: #1B365D;
            font-weight: 700;
            margin: 0 0 6px 0;
        }
        .eb-sub {
            font-size: 13px;
            color: #8c8c8c;
            margin: 0 0 24px 0;
        }
        .eb-sub strong { color: #1B365D; font-weight: 600; }
        .eb-divider {
            height: 1px;
            background: rgba(0,0,0,0.08);
        }
        .eb-content { padding: 24px 32px 4px; }
        .info-row {
            display: flex;
            padding: 10px 0;
            font-size: 14px;
        }
        .info-row .k {
            width: 72px;
            color: #8c8c8c;
            flex-shrink: 0;
        }
        .info-row .v {
            color: #333333;
            flex: 1;
            font-weight: 500;
        }
        .eb-quote-wrap { padding: 20px 32px; }
        .eb-quote {
            padding: 16px 0;
            border-top: 1px solid rgba(0,0,0,0.06);
            border-bottom: 1px solid rgba(0,0,0,0.06);
            color: #5d5d5d;
            font-size: 14px;
            line-height: 1.8;
        }
        .eb-quote-label {
            font-size: 11px;
            color: #1B365D;
            letter-spacing: 1px;
            margin-bottom: 10px;
            font-weight: 600;
        }
        .eb-action { padding: 8px 32px 28px; }
        .eb-btn {
            display: inline-flex;
            align-items: center;
            gap: 8px;
            color: #1B365D;
            text-decoration: none;
            font-size: 14px;
            font-weight: 600;
            border-bottom: 1px solid #1B365D;
            padding-bottom: 2px;
        }
        .eb-foot {
            padding: 20px 32px 24px;
            font-size: 12px;
            color: #8c8c8c;
            text-align: center;
            background: #faf9f5;
            border-top: 1px solid rgba(0,0,0,0.04);
        }
    </style>
</head>
<body>
    <div class="wrap">
        <div class="card">
            <div class="eb-header">
                <div class="eb-brand">— F c y &nbsp; B L O G —</div>
                <h1 class="eb-title">您的评论收到了回复</h1>
                <p class="eb-sub"><strong>%s</strong> 在文章下回复了您的评论</p>
            </div>
            <div class="eb-divider"></div>
            <div class="eb-content">
                <div class="info-row"><span class="k">回复者</span><span class="v">%s</span></div>
                <div class="info-row"><span class="k">文&nbsp;&nbsp;章</span><span class="v">%s</span></div>
            </div>
            <div class="eb-quote-wrap">
                <div class="eb-quote">
                    <div class="eb-quote-label">回 复 内 容</div>
                    %s
                </div>
            </div>
            <div class="eb-action">
                <a href="%s" class="eb-btn">查看回复详情 →</a>
            </div>
            <div class="eb-foot">此邮件由 Fcy's Blog 自动发送 · 日常落灰的个人博客 · 请勿直接回复</div>
        </div>
    </div>
</body>
</html>`, nickname, nickname, articleTitle, replyContent, link)

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
