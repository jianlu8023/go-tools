package smtp

import (
	"crypto/tls"
	"encoding/base64"
	"fmt"
	"net/smtp"
	"strings"
	"time"
)

// Config SMTP配置
type Config struct {
	Server             string        // SMTP服务器地址
	Port               int           // SMTP端口
	From               string        // 发件人邮箱
	Account            string        // 登录账号
	Password           string        // 登录密码
	ForceAuthLogin     bool          // 强制使用LOGIN认证
	SSLEnabled         bool          // 使用SSL连接（465端口直连TLS）
	StartTLSEnabled    bool          // 使用STARTTLS升级（587端口）
	InsecureSkipVerify bool          // 跳过TLS证书校验
	Timeout            time.Duration // 发送超时，0表示不限制
}

// Send 发送HTML邮件，receivers 支持以 ; 分隔的多个收件人
// @param cfg SMTP配置
// @param subject 邮件主题
// @param receivers 收件人（; 分隔）
// @param content HTML正文
// @return error 发送过程中的错误
func Send(cfg Config, subject string, receivers string, content string) error {
	from := cfg.From
	if from == "" {
		from = cfg.Account
	}
	if cfg.Server == "" && cfg.Account == "" {
		return fmt.Errorf("smtp server not configured")
	}
	id, err := generateMessageID(from)
	if err != nil {
		return err
	}

	encodedSubject := fmt.Sprintf("=?UTF-8?B?%s?=", base64.StdEncoding.EncodeToString([]byte(subject)))
	mail := []byte(fmt.Sprintf("To: %s\r\n"+
		"From: %s\r\n"+
		"Subject: %s\r\n"+
		"Date: %s\r\n"+
		"Message-ID: %s\r\n"+
		"Content-Type: text/html; charset=UTF-8\r\n\r\n%s\r\n",
		receivers, from, encodedSubject, time.Now().Format(time.RFC1123Z), id, content))

	addr := fmt.Sprintf("%s:%d", cfg.Server, cfg.Port)
	client, err := newClient(cfg, addr)
	if err != nil {
		return err
	}
	defer client.Close()

	if shouldAuthenticate(cfg) {
		if err = client.Auth(smtp.PlainAuth("", cfg.Account, cfg.Password, cfg.Server)); err != nil {
			return err
		}
	}
	if err = client.Mail(from); err != nil {
		return err
	}
	for _, receiver := range strings.Split(receivers, ";") {
		receiver = strings.TrimSpace(receiver)
		if receiver == "" {
			continue
		}
		if err = client.Rcpt(receiver); err != nil {
			return err
		}
	}
	w, err := client.Data()
	if err != nil {
		return err
	}
	if _, err = w.Write(mail); err != nil {
		return err
	}
	if err = w.Close(); err != nil {
		return err
	}
	return client.Quit()
}

// generateMessageID 生成RFC 5322格式的Message-ID
func generateMessageID(from string) (string, error) {
	at := strings.IndexByte(from, '@')
	if at < 0 || at >= len(from)-1 {
		return "", fmt.Errorf("invalid smtp account: %s", from)
	}
	domain := from[at+1:]
	return fmt.Sprintf("<%d.%d@%s>", time.Now().UnixNano(), time.Now().UnixNano()%100000000, domain), nil
}

// shouldAuthenticate 账号密码均设置时才认证
func shouldAuthenticate(cfg Config) bool {
	return cfg.Account != "" && cfg.Password != ""
}

// newClient 建立SMTP连接（SSL直连或明文+STARTTLS）
func newClient(cfg Config, addr string) (*smtp.Client, error) {
	tlsConfig := &tls.Config{
		ServerName:         cfg.Server,
		InsecureSkipVerify: cfg.InsecureSkipVerify,
	}
	if cfg.SSLEnabled || (cfg.Port == 465 && !cfg.StartTLSEnabled) {
		conn, err := tls.Dial("tcp", addr, tlsConfig)
		if err != nil {
			return nil, err
		}
		client, err := smtp.NewClient(conn, cfg.Server)
		if err != nil {
			conn.Close()
			return nil, err
		}
		return client, nil
	}

	client, err := smtp.Dial(addr)
	if err != nil {
		return nil, err
	}
	if cfg.StartTLSEnabled {
		startTLSSupported, _ := client.Extension("STARTTLS")
		if !startTLSSupported {
			client.Close()
			return nil, fmt.Errorf("smtp server does not support STARTTLS")
		}
		if err := client.StartTLS(tlsConfig); err != nil {
			client.Close()
			return nil, err
		}
	}
	return client, nil
}
