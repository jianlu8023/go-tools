package http

import (
	"context"
	"crypto/tls"
	"fmt"
	"github.com/tjfoc/gmsm/gmtls"
	"io"
	"net/http"
	"net/url"
	"strconv"
	"time"

	jsoniter "github.com/json-iterator/go"

	"github.com/go-resty/resty/v2"
)

// Client 是一个HTTP客户端封装，基于resty实现
type Client struct {
	client         *resty.Client
	Middlewares    []resty.RequestMiddleware
	RetryCondition func(*resty.Response, error) bool
	EnableTrace    bool
	PathParams     map[string]string
}

// NewClient 创建一个新的HTTP客户端
// @return *Client: 新创建的HTTP客户端
func NewClient() *Client {
	client := resty.New()
	client.SetTimeout(30 * time.Second)
	client.SetRetryCount(3)
	client.SetRetryWaitTime(500 * time.Millisecond)
	client.SetRetryMaxWaitTime(2 * time.Second)
	client.SetHeader("Content-Type", "application/json")
	client.SetJSONMarshaler(jsoniter.Marshal)
	client.SetJSONUnmarshaler(jsoniter.Unmarshal)

	return &Client{
		client:      client,
		Middlewares: make([]resty.RequestMiddleware, 0),
		PathParams:  make(map[string]string),
	}
}

// NewClientWithTimeout 创建一个自定义超时的HTTP客户端
// @param timeout: 超时时间(毫秒)
// @return *Client: 新创建的HTTP客户端
func NewClientWithTimeout(timeout int) *Client {
	client := NewClient()
	client.client.SetTimeout(time.Duration(timeout) * time.Millisecond)
	return client
}

// NewClientWithTLS 创建一个启用TLS自定义配置的HTTP客户端
// @param config: TLS配置
// @return *Client: 新创建的HTTP客户端
func NewClientWithTLS(config *tls.Config) *Client {
	client := NewClient()
	client.client.SetTLSClientConfig(config)
	return client
}

func NewClientWithGMTls(config *gmtls.Config) *Client {
	return NewClientWithGoHttpClient(NewGMGoHttpClient(config))
}

// NewClientWithGoHttpClient 基于自定义 Go 标准库 HTTP 客户端创建封装的 Client。
// 若 gohttpClient 为 nil，则使用默认配置（等价于 NewClient），避免 nil 传递给 resty。
func NewClientWithGoHttpClient(gohttpClient *http.Client) *Client {
	var client *resty.Client
	if gohttpClient == nil {
		client = resty.New()
	} else {
		client = resty.NewWithClient(gohttpClient)
	}
	client.SetTimeout(30 * time.Second)
	client.SetRetryCount(3)
	client.SetRetryWaitTime(500 * time.Millisecond)
	client.SetRetryMaxWaitTime(2 * time.Second)
	client.SetHeader("Content-Type", "application/json")
	client.SetJSONMarshaler(jsoniter.Marshal)
	client.SetJSONUnmarshaler(jsoniter.Unmarshal)

	return &Client{
		client:      client,
		Middlewares: make([]resty.RequestMiddleware, 0),
		PathParams:  make(map[string]string),
	}
}

// SetHeader 设置请求头
// @param key: 请求头键
// @param value: 请求头值
// @return *Client: 当前客户端实例，支持链式调用
func (c *Client) SetHeader(key, value string) *Client {
	c.client.SetHeader(key, value)
	return c
}

// SetHeaders 设置多个请求头
// @param headers: 请求头映射
// @return *Client: 当前客户端实例，支持链式调用
func (c *Client) SetHeaders(headers map[string]string) *Client {
	for k, v := range headers {
		c.client.SetHeader(k, v)
	}
	return c
}

// SetTimeout 设置超时时间
// @param timeout: 超时时间(毫秒)
// @return *Client: 当前客户端实例，支持链式调用
func (c *Client) SetTimeout(timeout int) *Client {
	c.client.SetTimeout(time.Duration(timeout) * time.Millisecond)
	return c
}

// SetRetry 设置重试策略
// @param count: 重试次数
// @param waitTime: 重试等待时间(毫秒)
// @param maxWaitTime: 最大重试等待时间(毫秒)
// @return *Client: 当前客户端实例，支持链式调用
func (c *Client) SetRetry(count, waitTime, maxWaitTime int) *Client {
	c.client.SetRetryCount(count)
	c.client.SetRetryWaitTime(time.Duration(waitTime) * time.Millisecond)
	c.client.SetRetryMaxWaitTime(time.Duration(maxWaitTime) * time.Millisecond)
	return c
}

// SetRetryCondition 设置重试条件函数
// @param condition: 重试条件函数，返回true表示需要重试
// @return *Client: 当前客户端实例，支持链式调用
func (c *Client) SetRetryCondition(condition func(*resty.Response, error) bool) *Client {
	c.RetryCondition = condition
	if condition != nil {
		c.client.AddRetryCondition(condition)
	}
	return c
}

// AddMiddleware 添加请求中间件
// @param middleware: 请求中间件函数
// @return *Client: 当前客户端实例，支持链式调用
func (c *Client) AddMiddleware(middleware resty.RequestMiddleware) *Client {
	c.Middlewares = append(c.Middlewares, middleware)
	c.client.OnBeforeRequest(middleware)
	return c
}

// SetEnableTrace 启用请求跟踪
// @param enable: 是否启用
// @return *Client: 当前客户端实例，支持链式调用
func (c *Client) SetEnableTrace(enable bool) *Client {
	c.EnableTrace = enable
	if enable {
		c.client.EnableTrace()
	}
	return c
}

// SetPathParam 设置URL路径参数
// @param key: 参数名
// @param value: 参数值
// @return *Client: 当前客户端实例，支持链式调用
func (c *Client) SetPathParam(key, value string) *Client {
	c.PathParams[key] = value
	return c
}

// SetPathParams 设置多个URL路径参数
// @param params: 参数映射
// @return *Client: 当前客户端实例，支持链式调用
func (c *Client) SetPathParams(params map[string]string) *Client {
	for k, v := range params {
		c.PathParams[k] = v
	}
	return c
}

// SetCookies 设置Cookie
// @param cookies: Cookie列表
// @return *Client: 当前客户端实例，支持链式调用
func (c *Client) SetCookies(cookies []*http.Cookie) *Client {
	c.client.SetCookies(cookies)
	return c
}

// SetUserAgent 设置User-Agent
// @param userAgent: User-Agent字符串
// @return *Client: 当前客户端实例，支持链式调用
func (c *Client) SetUserAgent(userAgent string) *Client {
	c.client.SetHeader("User-Agent", userAgent)
	return c
}

// SetDigestAuth 设置Digest认证
// @param username: 用户名
// @param password: 密码
// @return *Client: 当前客户端实例，支持链式调用
func (c *Client) SetDigestAuth(username, password string) *Client {
	c.client.SetDigestAuth(username, password)
	return c
}

// SetContentType 设置请求Content-Type
// @param contentType: 内容类型
// @return *Client: 当前客户端实例，支持链式调用
func (c *Client) SetContentType(contentType string) *Client {
	c.client.SetHeader("Content-Type", contentType)
	return c
}

// SetContentLength 是否设置Content-Length请求头
// @param flag: 是否启用
// @return *Client: 当前客户端实例，支持链式调用
func (c *Client) SetContentLength(flag bool) *Client {
	c.client.SetContentLength(flag)
	return c
}

// SetLogger 设置日志记录器
// @param logger: 日志接口
// @return *Client: 当前客户端实例，支持链式调用
func (c *Client) SetLogger(logger resty.Logger) *Client {
	c.client.SetLogger(logger)
	return c
}

// SetBasicAuth 设置基本认证
// @param username: 用户名
// @param password: 密码
// @return *Client: 当前客户端实例，支持链式调用
func (c *Client) SetBasicAuth(username, password string) *Client {
	c.client.SetBasicAuth(username, password)
	return c
}

// SetBearerAuth 设置Bearer令牌认证
// @param token: Bearer令牌
// @return *Client: 当前客户端实例，支持链式调用
func (c *Client) SetBearerAuth(token string) *Client {
	c.client.SetAuthToken(token)
	return c
}

// SetProxy 设置HTTP代理
// @param proxyURL: 代理URL
// @return *Client: 当前客户端实例，支持链式调用
func (c *Client) SetProxy(proxyURL string) *Client {
	c.client.SetProxy(proxyURL)
	return c
}

// SetTLSClientConfig 设置TLS配置
// @param config: TLS配置
// @return *Client: 当前客户端实例，支持链式调用
func (c *Client) SetTLSClientConfig(config *tls.Config) *Client {
	c.client.SetTLSClientConfig(config)
	return c
}

// SetRedirectPolicy 设置重定向策略
// @param redirectPolicy: 重定向策略
// @return *Client: 当前客户端实例，支持链式调用
func (c *Client) SetRedirectPolicy(redirectPolicy resty.RedirectPolicy) *Client {
	c.client.SetRedirectPolicy(redirectPolicy)
	return c
}

// GetRawClient 获取底层的resty客户端
// @return *resty.Client: 底层的resty客户端
func (c *Client) GetRawClient() *resty.Client {
	return c.client
}

// GetTraceInfo 获取请求跟踪信息
// @param resp: 请求响应
// @return resty.TraceInfo: 跟踪信息
func (c *Client) GetTraceInfo(resp *resty.Response) resty.TraceInfo {
	return resp.Request.TraceInfo()
}

// convertParam 将查询参数值转换为URL编码的字符串
func convertParam(v interface{}) string {
	switch val := v.(type) {
	case string:
		return val
	case int:
		return strconv.Itoa(val)
	case int64:
		return strconv.FormatInt(val, 10)
	case float64:
		return strconv.FormatFloat(val, 'f', -1, 64)
	case bool:
		return strconv.FormatBool(val)
	case time.Time:
		return val.Format(time.RFC3339)
	case fmt.Stringer:
		return val.String()
	default:
		return fmt.Sprintf("%v", v)
	}
}

// setQueryParams 将查询参数设置到请求中
func setQueryParams(r *resty.Request, params map[string]interface{}) {
	if params == nil {
		return
	}
	v := url.Values{}
	for k, param := range params {
		v.Set(k, convertParam(param))
	}
	r.SetQueryParamsFromValues(v)
}

// buildRequest 构建带公共配置的请求
func (c *Client) buildRequest() *resty.Request {
	r := c.client.R()
	if len(c.PathParams) > 0 {
		r.SetPathParams(c.PathParams)
	}
	if c.EnableTrace {
		r.EnableTrace()
	}
	return r
}

// buildRequestWithContext 构建带上下文的请求，复用 buildRequest 的公共配置。
// 所有 WithContext 变体应统一调用此方法，避免公共配置（PathParams、EnableTrace）变更时遗漏同步。
func (c *Client) buildRequestWithContext(ctx context.Context) *resty.Request {
	return c.buildRequest().SetContext(ctx)
}

// GET 发送GET请求
// @param url: 请求URL
// @param params: URL查询参数
// @return []byte: 响应体
// @return int: 响应状态码
// @return error: 错误信息
func (c *Client) GET(url string, params map[string]interface{}) ([]byte, int, error) {
	r := c.buildRequest()
	setQueryParams(r, params)

	resp, err := r.Get(url)
	if err != nil {
		return nil, 0, fmt.Errorf("GET请求失败: %v", err)
	}
	return resp.Body(), resp.StatusCode(), nil
}

// GETContext 发送带上下文的GET请求，支持超时和取消
// @param ctx: 上下文，可用于取消请求或设置请求级超时
// @param url: 请求URL
// @param params: URL查询参数
// @return []byte: 响应体
// @return int: 响应状态码
// @return error: 错误信息
func (c *Client) GETContext(ctx context.Context, url string, params map[string]interface{}) ([]byte, int, error) {
	r := c.buildRequestWithContext(ctx)
	setQueryParams(r, params)

	resp, err := r.Get(url)
	if err != nil {
		return nil, 0, fmt.Errorf("GET请求失败: %v", err)
	}
	return resp.Body(), resp.StatusCode(), nil
}

// GETWithResult 发送GET请求并自动解析JSON响应到结构体
// @param url: 请求URL
// @param params: URL查询参数
// @param result: 解析JSON的目标结构体指针
// @return *resty.Response: 原始响应
// @return error: 错误信息
func (c *Client) GETWithResult(url string, params map[string]interface{}, result interface{}) (*resty.Response, error) {
	r := c.buildRequest().SetResult(result)
	setQueryParams(r, params)
	return r.Get(url)
}

// GetJSON 发送GET请求并自动解析JSON响应到结构体（返回状态码）
// @param url: 请求URL
// @param params: URL查询参数
// @param result: 解析JSON的目标结构体指针
// @return int: 响应状态码
// @return error: 错误信息
func (c *Client) GetJSON(url string, params map[string]interface{}, result interface{}) (int, error) {
	resp, err := c.GETWithResult(url, params, result)
	if err != nil {
		return 0, err
	}
	return resp.StatusCode(), nil
}

// PostJSON 发送POST请求并自动解析JSON请求体和响应（包含响应JSON反序列化）。
// 与 POSTJSON 的区别：PostJSON 接受 result 参数用于反序列化响应体，POSTJSON 只返回原始响应字节。
// @param url: 请求URL
// @param body: 请求体结构
// @param result: 解析JSON的目标结构体指针
// @return int: 响应状态码
// @return error: 错误信息
func (c *Client) PostJSON(url string, body interface{}, result interface{}) (int, error) {
	resp, err := c.POSTJSONWithResult(url, body, result)
	if err != nil {
		return 0, err
	}
	return resp.StatusCode(), nil
}

// POST 发送POST请求(通用)
// @param url: 请求URL
// @param body: 请求体
// @return []byte: 响应体
// @return int: 响应状态码
// @return error: 错误信息
func (c *Client) POST(url string, body interface{}) ([]byte, int, error) {
	r := c.buildRequest()
	r.SetBody(body)

	resp, err := r.Post(url)
	if err != nil {
		return nil, 0, fmt.Errorf("POST请求失败: %v", err)
	}
	return resp.Body(), resp.StatusCode(), nil
}

// POSTJSON 发送POST请求并序列化JSON请求体（仅返回原始响应字节，不反序列化响应）。
// 与 PostJSON 的区别：POSTJSON 无 result 参数，只返回原始响应体字节；PostJSON 会自动反序列化响应体。
// @param url: 请求URL
// @param body: 请求体结构
// @return []byte: 响应体
// @return int: 响应状态码
// @return error: 错误信息
func (c *Client) POSTJSON(url string, body interface{}) ([]byte, int, error) {
	r := c.buildRequest()
	r.SetBody(body)
	r.SetHeader("Content-Type", "application/json")

	resp, err := r.Post(url)
	if err != nil {
		return nil, 0, fmt.Errorf("POST请求失败: %v", err)
	}
	return resp.Body(), resp.StatusCode(), nil
}

// POSTJSONWithResult 发送POST请求并序列化JSON请求体，自动解析JSON响应
// @param url: 请求URL
// @param body: 请求体结构
// @param result: 解析JSON的目标结构体指针
// @return *resty.Response: 原始响应
// @return error: 错误信息
func (c *Client) POSTJSONWithResult(url string, body interface{}, result interface{}) (*resty.Response, error) {
	r := c.buildRequest().SetBody(body).SetResult(result).SetHeader("Content-Type", "application/json")
	return r.Post(url)
}

// DELETEJSON 发送DELETE请求并携带JSON请求体
// @param url: 请求URL
// @param body: 请求体结构
// @return []byte: 响应体
// @return int: 响应状态码
// @return error: 错误信息
func (c *Client) DELETEJSON(url string, body interface{}) ([]byte, int, error) {
	r := c.buildRequest()
	r.SetBody(body)
	r.SetHeader("Content-Type", "application/json")

	resp, err := r.Delete(url)
	if err != nil {
		return nil, 0, fmt.Errorf("DELETE请求失败: %v", err)
	}
	return resp.Body(), resp.StatusCode(), nil
}

// PUT 发送PUT请求
// @param url: 请求URL
// @param body: 请求体
// @return []byte: 响应体
// @return int: 响应状态码
// @return error: 错误信息
func (c *Client) PUT(url string, body interface{}) ([]byte, int, error) {
	r := c.buildRequest()
	r.SetBody(body)

	resp, err := r.Put(url)
	if err != nil {
		return nil, 0, fmt.Errorf("PUT请求失败: %v", err)
	}
	return resp.Body(), resp.StatusCode(), nil
}

// DELETE 发送DELETE请求
// @param url: 请求URL
// @return []byte: 响应体
// @return int: 响应状态码
// @return error: 错误信息
func (c *Client) DELETE(url string) ([]byte, int, error) {
	r := c.buildRequest()

	resp, err := r.Delete(url)
	if err != nil {
		return nil, 0, fmt.Errorf("DELETE请求失败: %v", err)
	}
	return resp.Body(), resp.StatusCode(), nil
}

// HEAD 发送HEAD请求
// @param url: 请求URL
// @return http.Header: 响应头
// @return int: 响应状态码
// @return error: 错误信息
func (c *Client) HEAD(url string) (http.Header, int, error) {
	r := c.buildRequest()

	resp, err := r.Head(url)
	if err != nil {
		return nil, 0, fmt.Errorf("HEAD请求失败: %v", err)
	}
	return resp.Header(), resp.StatusCode(), nil
}

// PATCH 发送PATCH请求
// @param url: 请求URL
// @param body: 请求体
// @return []byte: 响应体
// @return int: 响应状态码
// @return error: 错误信息
func (c *Client) PATCH(url string, body interface{}) ([]byte, int, error) {
	r := c.buildRequest()
	r.SetBody(body)

	resp, err := r.Patch(url)
	if err != nil {
		return nil, 0, fmt.Errorf("PATCH请求失败: %v", err)
	}
	return resp.Body(), resp.StatusCode(), nil
}

// UploadFile 上传单个文件
// @param url: 上传URL
// @param fieldName: 文件字段名
// @param filePath: 要上传的文件路径
// @param params: 其他表单参数
// @return []byte: 响应体
// @return int: 响应状态码
// @return error: 错误信息
func (c *Client) UploadFile(url, fieldName, filePath string, params map[string]string) ([]byte, int, error) {
	r := c.buildRequest()
	if params != nil {
		r.SetFormData(params)
	}
	// 设置文件
	r.SetFile(fieldName, filePath)

	resp, err := r.Post(url)
	if err != nil {
		return nil, 0, fmt.Errorf("文件上传失败: %v", err)
	}
	return resp.Body(), resp.StatusCode(), nil
}

// UploadMultiFile 上传多个文件
// @param url: 上传URL
// @param files: 文件映射 (fieldName -> filePath)
// @param params: 其他表单参数
// @return []byte: 响应体
// @return int: 响应状态码
// @return error: 错误信息
func (c *Client) UploadMultiFile(url string, files map[string]string, params map[string]string) ([]byte, int, error) {
	r := c.buildRequest()
	if params != nil {
		r.SetFormData(params)
	}
	for fieldName, filePath := range files {
		r.SetFile(fieldName, filePath)
	}

	resp, err := r.Post(url)
	if err != nil {
		return nil, 0, fmt.Errorf("文件上传失败: %v", err)
	}
	return resp.Body(), resp.StatusCode(), nil
}

// POSTJSON 发送JSON POST请求(使用默认客户端)
// @param url: 请求URL
// @param body: 请求体
// @return []byte: 响应体
// @return int: 响应状态码
// @return error: 错误信息
func POSTJSON(url string, body interface{}) ([]byte, int, error) {
	return defaultClient.POSTJSON(url, body)
}

// SendRequest 发送自定义请求
// @param method: HTTP方法
// @param url: 请求URL
// @param headers: 请求头
// @param body: 请求体
// @return []byte: 响应体
// @return int: 响应状态码
// @return error: 错误信息
func (c *Client) SendRequest(method, url string, headers map[string]string, body interface{}) ([]byte, int, error) {
	r := c.buildRequest()
	if headers != nil {
		r.SetHeaders(headers)
	}
	if body != nil {
		r.SetBody(body)
	}

	resp, err := r.Execute(method, url)
	if err != nil {
		return nil, 0, fmt.Errorf("请求失败: %v", err)
	}
	return resp.Body(), resp.StatusCode(), nil
}

// Stream 发送流式请求并处理响应
// @param method: HTTP方法
// @param url: 请求URL
// @param headers: 请求头
// @param body: 请求体
// @param callback: 处理流式响应的回调函数
// @return error: 错误信息
func (c *Client) Stream(method, url string, headers map[string]string, body interface{}, callback StreamCallback) error {
	r := c.buildRequest()

	if headers != nil {
		r.SetHeaders(headers)
	}

	if body != nil {
		r.SetBody(body)
	}

	r.SetDoNotParseResponse(true)

	resp, err := r.Execute(method, url)
	if err != nil {
		callback(nil, fmt.Errorf("请求失败: %v", err))
		return err
	}

	return streamResponse(resp, callback)
}

// StreamCallback 定义流式响应的回调函数类型
// @param data: 接收到的数据片段
// @param err: 错误信息，如果有的话
// @return bool: 是否继续接收数据，true表示继续，false表示停止
type StreamCallback func(data []byte, err error) bool

// streamResponse 通用的流式响应处理逻辑
func streamResponse(resp *resty.Response, callback StreamCallback) error {
	if resp.StatusCode() < 200 || resp.StatusCode() >= 300 {
		// 读取错误响应体（如果有的话）
		if resp.RawResponse != nil && resp.RawResponse.Body != nil {
			defer resp.RawResponse.Body.Close()
			body, _ := io.ReadAll(resp.RawResponse.Body)
			callback(nil, fmt.Errorf("请求失败，状态码: %d, 响应: %s", resp.StatusCode(), string(body)))
		} else {
			callback(nil, fmt.Errorf("请求失败，状态码: %d", resp.StatusCode()))
		}
		return fmt.Errorf("请求失败，状态码: %d", resp.StatusCode())
	}

	// 获取原始响应体
	if resp.RawResponse == nil || resp.RawResponse.Body == nil {
		err := fmt.Errorf("无法获取响应体")
		callback(nil, err)
		return err
	}

	reader := resp.RawResponse.Body
	defer reader.Close()

	// 处理流式响应
	buf := make([]byte, 1024)
	for {
		n, err := reader.Read(buf)
		if n > 0 {
			// 将读取到的数据传递给回调函数
			if !callback(buf[:n], nil) {
				break // 用户选择停止接收数据
			}
		}

		if err != nil {
			if err == io.EOF {
				// 流结束，正常退出
				callback(nil, nil) // 发送一个结束信号
				break
			} else {
				// 其他错误
				callback(nil, err)
				return err
			}
		}
	}

	return nil
}

// StreamGET 发送GET流式请求
// @param url: 请求URL
// @param params: URL查询参数
// @param callback: 处理流式响应的回调函数
// @return error: 错误信息
func (c *Client) StreamGET(url string, params map[string]interface{}, callback StreamCallback) error {
	r := c.buildRequest()
	setQueryParams(r, params)

	r.SetDoNotParseResponse(true)

	resp, err := r.Get(url)
	if err != nil {
		callback(nil, fmt.Errorf("请求失败: %v", err))
		return err
	}

	return streamResponse(resp, callback)
}

// StreamPOST 发送POST流式请求
// @param url: 请求URL
// @param body: 请求体
// @param callback: 处理流式响应的回调函数
// @return error: 错误信息
func (c *Client) StreamPOST(url string, body interface{}, callback StreamCallback) error {
	return c.Stream("POST", url, nil, body, callback)
}

// StreamWithRequest 发送自定义流式请求
// @param method: HTTP方法
// @param url: 请求URL
// @param headers: 请求头
// @param body: 请求体
// @param callback: 处理流式响应的回调函数
// @return error: 错误信息
func (c *Client) StreamWithRequest(method, url string, headers map[string]string, body interface{}, callback StreamCallback) error {
	return c.Stream(method, url, headers, body, callback)
}

// 以下是快捷函数，直接使用默认客户端

var defaultClient = NewClient()

// GET 发送GET请求(使用默认客户端)
// @param url: 请求URL
// @param params: URL查询参数
// @return []byte: 响应体
// @return int: 响应状态码
// @return error: 错误信息
func GET(url string, params map[string]interface{}) ([]byte, int, error) {
	return defaultClient.GET(url, params)
}

// POST 发送POST请求(使用默认客户端)
// @param url: 请求URL
// @param body: 请求体
// @return []byte: 响应体
// @return int: 响应状态码
// @return error: 错误信息
func POST(url string, body interface{}) ([]byte, int, error) {
	return defaultClient.POST(url, body)
}

// PUT 发送PUT请求(使用默认客户端)
// @param url: 请求URL
// @param body: 请求体
// @return []byte: 响应体
// @return int: 响应状态码
// @return error: 错误信息
func PUT(url string, body interface{}) ([]byte, int, error) {
	return defaultClient.PUT(url, body)
}

// DELETE 发送DELETE请求(使用默认客户端)
// @param url: 请求URL
// @return []byte: 响应体
// @return int: 响应状态码
// @return error: 错误信息
func DELETE(url string) ([]byte, int, error) {
	return defaultClient.DELETE(url)
}

// GetJSON 发送GET请求并自动解析JSON响应到结构体(使用默认客户端)
// @param url: 请求URL
// @param params: URL查询参数
// @param result: 解析JSON的目标结构体指针
// @return int: 响应状态码
// @return error: 错误信息
func GetJSON(url string, params map[string]interface{}, result interface{}) (int, error) {
	return defaultClient.GetJSON(url, params, result)
}

// PostJSON 发送POST请求并自动解析JSON请求体和响应(使用默认客户端)
// @param url: 请求URL
// @param body: 请求体
// @param result: 解析JSON的目标结构体指针
// @return int: 响应状态码
// @return error: 错误信息
func PostJSON(url string, body interface{}, result interface{}) (int, error) {
	return defaultClient.PostJSON(url, body, result)
}

// DownloadFile 下载文件(使用默认客户端)
// @param url: 文件URL
// @param filePath: 保存路径
// @return int64: 下载的字节数
// @return error: 错误信息
func DownloadFile(url, filePath string) (int64, error) {
	return defaultClient.DownloadFile(url, filePath)
}

// UploadFile 上传文件(使用默认客户端)
// @param url: 上传URL
// @param fieldName: 文件字段名
// @param filePath: 要上传的文件路径
// @param params: 其他表单参数
// @return []byte: 响应体
// @return int: 响应状态码
// @return error: 错误信息
func UploadFile(url, fieldName, filePath string, params map[string]string) ([]byte, int, error) {
	return defaultClient.UploadFile(url, fieldName, filePath, params)
}

// StreamGET 发送GET流式请求(使用默认客户端)
// @param url: 请求URL
// @param params: URL查询参数
// @param callback: 处理流式响应的回调函数
// @return error: 错误信息
func StreamGET(url string, params map[string]interface{}, callback StreamCallback) error {
	return defaultClient.StreamGET(url, params, callback)
}

// StreamPOST 发送POST流式请求(使用默认客户端)
// @param url: 请求URL
// @param body: 请求体
// @param callback: 处理流式响应的回调函数
// @return error: 错误信息
func StreamPOST(url string, body interface{}, callback StreamCallback) error {
	return defaultClient.StreamPOST(url, body, callback)
}

// StreamWithRequest 发送自定义流式请求(使用默认客户端)
// @param method: HTTP方法
// @param url: 请求URL
// @param headers: 请求头
// @param body: 请求体
// @param callback: 处理流式响应的回调函数
// @return error: 错误信息
func StreamWithRequest(method, url string, headers map[string]string, body interface{}, callback StreamCallback) error {
	return defaultClient.StreamWithRequest(method, url, headers, body, callback)
}

// UploadMultiFile 上传多文件(使用默认客户端)
// @param url: 上传URL
// @param files: 文件映射
// @param params: 其他表单参数
// @return []byte: 响应体
// @return int: 响应状态码
// @return error: 错误信息
func UploadMultiFile(url string, files map[string]string, params map[string]string) ([]byte, int, error) {
	return defaultClient.UploadMultiFile(url, files, params)
}

// DELETEJSON 发送DELETE请求并携带JSON请求体(使用默认客户端)
// @param url: 请求URL
// @param body: 请求体
// @return []byte: 响应体
// @return int: 响应状态码
// @return error: 错误信息
func DELETEJSON(url string, body interface{}) ([]byte, int, error) {
	return defaultClient.DELETEJSON(url, body)
}

// GETContext 发送带上下文的GET请求(使用默认客户端)
// @param ctx: 上下文
// @param url: 请求URL
// @param params: URL查询参数
// @return []byte: 响应体
// @return int: 响应状态码
// @return error: 错误信息
func GETContext(ctx context.Context, url string, params map[string]interface{}) ([]byte, int, error) {
	return defaultClient.GETContext(ctx, url, params)
}

// StreamWithContext 发送带上下文的流式请求
// @param ctx: 上下文，可用于取消请求或设置请求级超时
// @param method: HTTP方法
// @param url: 请求URL
// @param headers: 请求头
// @param body: 请求体
// @param callback: 处理流式响应的回调函数
// @return error: 错误信息
func (c *Client) StreamWithContext(ctx context.Context, method, url string, headers map[string]string, body interface{}, callback StreamCallback) error {
	r := c.buildRequestWithContext(ctx)

	if headers != nil {
		r.SetHeaders(headers)
	}

	if body != nil {
		r.SetBody(body)
	}

	r.SetDoNotParseResponse(true)

	resp, err := r.Execute(method, url)
	if err != nil {
		callback(nil, fmt.Errorf("请求失败: %v", err))
		return err
	}

	return streamResponse(resp, callback)
}

// SendRequestWithContext 发送带上下文的自定义请求
// @param ctx: 上下文，可用于取消请求或设置请求级超时
// @param method: HTTP方法
// @param url: 请求URL
// @param headers: 请求头
// @param body: 请求体
// @return []byte: 响应体
// @return int: 响应状态码
// @return error: 错误信息
func (c *Client) SendRequestWithContext(ctx context.Context, method, url string, headers map[string]string, body interface{}) ([]byte, int, error) {
	r := c.buildRequestWithContext(ctx)

	if headers != nil {
		r.SetHeaders(headers)
	}

	if body != nil {
		r.SetBody(body)
	}

	resp, err := r.Execute(method, url)
	if err != nil {
		return nil, 0, fmt.Errorf("请求失败: %v", err)
	}
	return resp.Body(), resp.StatusCode(), nil
}

// DownloadFile 下载文件
// @param url: 文件URL
// @param filePath: 保存路径
// @return int64: 下载的字节数
// @return error: 错误信息
func (c *Client) DownloadFile(url, filePath string) (int64, error) {
	r := c.buildRequest()
	r.SetOutput(filePath)

	resp, err := r.Get(url)
	if err != nil {
		return 0, fmt.Errorf("文件下载失败: %v", err)
	}
	if resp.StatusCode() != http.StatusOK {
		return 0, fmt.Errorf("下载文件失败，状态码: %d", resp.StatusCode())
	}
	return resp.Size(), nil
}
