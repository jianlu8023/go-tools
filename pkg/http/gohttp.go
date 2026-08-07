package http

import (
	"context"
	"github.com/tjfoc/gmsm/gmtls"
	"net"
	"net/http"
	"time"
)

// NewGMGoHttpClient 使用国密 TLS 配置创建一个 Go 标准库 HTTP 客户端。
// 若 config 为 nil 则基于默认配置创建标准 TLS 的 HTTP 客户端（非 nil，可直接使用），
// 与 NewClientWithGMTls 等上层 API 的 nil 校验逻辑协同，避免 nil 客户端传递。
func NewGMGoHttpClient(config *gmtls.Config) *http.Client {
	dialer := &net.Dialer{
		Timeout:   30 * time.Second,
		KeepAlive: 60 * time.Second,
	}
	if config == nil {
		// 无国密配置时退化为标准 HTTP/HTTPS 客户端
		return &http.Client{
			Transport: &http.Transport{
				TLSHandshakeTimeout: 30 * time.Second,
				IdleConnTimeout:     30 * time.Second,
			},
			Timeout: 30 * time.Second,
		}
	}
	return NewGMGoHttpClientWithDialer(config, dialer)
}

// NewGMGoHttpClientWithDialer 使用国密 TLS 配置和自定义 Dialer 创建 HTTP 客户端。
// 若 config 或 dialer 为 nil 则返回 nil，调用方需自行处理（或改用 NewGMGoHttpClient）。
func NewGMGoHttpClientWithDialer(config *gmtls.Config, dialer *net.Dialer) *http.Client {
	if config == nil || dialer == nil {
		return nil
	}
	return &http.Client{
		Transport: &http.Transport{
			DialTLSContext: func(ctx context.Context, network, addr string) (net.Conn, error) {
				return gmtls.DialWithDialer(dialer, network, addr, config)
			},
			TLSHandshakeTimeout: 30 * time.Second,
			IdleConnTimeout:     30 * time.Second,
		},
	}
}
