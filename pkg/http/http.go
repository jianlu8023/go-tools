package http

import (
	"crypto/tls"
	"fmt"
	"io"
	"net/http"
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

// GET 发送GET请求
// @param url: 请求URL
// @param params: URL查询参数
// @return []byte: 响应体
// @return int: 响应状态码
// @return error: 错误信息
func (c *Client) GET(url string, params map[string]interface{}) ([]byte, int, error) {
	r := c.client.R()
	// 设置路径参数
	if len(c.PathParams) > 0 {
		r.SetPathParams(c.PathParams)
	}
	// 设置查询参数
	if params != nil {
		for k, v := range params {
			r.SetQueryParam(k, fmt.Sprintf("%v", v))
		}
	}
	// 启用跟踪
	if c.EnableTrace {
		r.EnableTrace()
	}
	resp, err := r.Get(url)
	if err != nil {
		return nil, 0, fmt.Errorf("GET请求失败: %v", err)
	}
	return resp.Body(), resp.StatusCode(), nil
}

// POST 发送POST请求
// @param url: 请求URL
// @param body: 请求体
// @return []byte: 响应体
// @return int: 响应状态码
// @return error: 错误信息
func (c *Client) POST(url string, body interface{}) ([]byte, int, error) {
	r := c.client.R()
	// 设置路径参数
	if len(c.PathParams) > 0 {
		r.SetPathParams(c.PathParams)
	}
	// 设置请求体
	r.SetBody(body)
	// 启用跟踪
	if c.EnableTrace {
		r.EnableTrace()
	}
	resp, err := r.Post(url)
	if err != nil {
		return nil, 0, fmt.Errorf("POST请求失败: %v", err)
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
	r := c.client.R()
	// 设置路径参数
	if len(c.PathParams) > 0 {
		r.SetPathParams(c.PathParams)
	}
	// 设置请求体
	r.SetBody(body)
	// 启用跟踪
	if c.EnableTrace {
		r.EnableTrace()
	}
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
	r := c.client.R()
	// 设置路径参数
	if len(c.PathParams) > 0 {
		r.SetPathParams(c.PathParams)
	}
	// 启用跟踪
	if c.EnableTrace {
		r.EnableTrace()
	}
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
	r := c.client.R()
	// 设置路径参数
	if len(c.PathParams) > 0 {
		r.SetPathParams(c.PathParams)
	}
	// 启用跟踪
	if c.EnableTrace {
		r.EnableTrace()
	}
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
	r := c.client.R()
	// 设置路径参数
	if len(c.PathParams) > 0 {
		r.SetPathParams(c.PathParams)
	}
	// 设置请求体
	r.SetBody(body)
	// 启用跟踪
	if c.EnableTrace {
		r.EnableTrace()
	}
	resp, err := r.Patch(url)
	if err != nil {
		return nil, 0, fmt.Errorf("PATCH请求失败: %v", err)
	}
	return resp.Body(), resp.StatusCode(), nil
}

// DownloadFile 下载文件
// @param url: 文件URL
// @param filePath: 保存路径
// @return int64: 下载的字节数
// @return error: 错误信息
func (c *Client) DownloadFile(url, filePath string) (int64, error) {
	r := c.client.R()
	// 设置路径参数
	if len(c.PathParams) > 0 {
		r.SetPathParams(c.PathParams)
	}
	// 设置输出文件
	r.SetOutput(filePath)
	// 启用跟踪
	if c.EnableTrace {
		r.EnableTrace()
	}
	resp, err := r.Get(url)
	if err != nil {
		return 0, fmt.Errorf("文件下载失败: %v", err)
	}
	if resp.StatusCode() != http.StatusOK {
		return 0, fmt.Errorf("下载文件失败，状态码: %d", resp.StatusCode())
	}
	return resp.Size(), nil
}

// UploadFile 上传文件
// @param url: 上传URL
// @param fieldName: 文件字段名
// @param filePath: 要上传的文件路径
// @param params: 其他表单参数
// @return []byte: 响应体
// @return int: 响应状态码
// @return error: 错误信息
func (c *Client) UploadFile(url, fieldName, filePath string, params map[string]string) ([]byte, int, error) {
	r := c.client.R()
	// 设置路径参数
	if len(c.PathParams) > 0 {
		r.SetPathParams(c.PathParams)
	}
	// 设置表单参数
	if params != nil {
		r.SetFormData(params)
	}
	// 设置文件
	r.SetFile(fieldName, filePath)
	// 启用跟踪
	if c.EnableTrace {
		r.EnableTrace()
	}
	resp, err := r.Post(url)
	if err != nil {
		return nil, 0, fmt.Errorf("文件上传失败: %v", err)
	}
	return resp.Body(), resp.StatusCode(), nil
}

// GetJSON 发送GET请求并解析JSON响应
// @param url: 请求URL
// @param params: URL查询参数
// @param result: 解析JSON的目标结构体指针
// @return int: 响应状态码
// @return error: 错误信息
func (c *Client) GetJSON(url string, params map[string]interface{}, result interface{}) (int, error) {
	r := c.client.R()
	// 设置路径参数
	if len(c.PathParams) > 0 {
		r.SetPathParams(c.PathParams)
	}
	// 设置查询参数
	if params != nil {
		for k, v := range params {
			r.SetQueryParam(k, fmt.Sprintf("%v", v))
		}
	}
	// 设置结果结构体
	r.SetResult(result)
	// 启用跟踪
	if c.EnableTrace {
		r.EnableTrace()
	}
	resp, err := r.Get(url)
	if err != nil {
		return 0, fmt.Errorf("GET请求失败: %v", err)
	}
	return resp.StatusCode(), nil
}

// PostJSON 发送POST请求并解析JSON响应
// @param url: 请求URL
// @param body: 请求体
// @param result: 解析JSON的目标结构体指针
// @return int: 响应状态码
// @return error: 错误信息
func (c *Client) PostJSON(url string, body interface{}, result interface{}) (int, error) {
	r := c.client.R()
	// 设置路径参数
	if len(c.PathParams) > 0 {
		r.SetPathParams(c.PathParams)
	}
	// 设置请求体
	r.SetBody(body)
	// 设置结果结构体
	r.SetResult(result)
	// 启用跟踪
	if c.EnableTrace {
		r.EnableTrace()
	}
	resp, err := r.Post(url)
	if err != nil {
		return 0, fmt.Errorf("POST请求失败: %v", err)
	}
	return resp.StatusCode(), nil
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

// SendRequest 发送自定义请求
// @param method: HTTP方法
// @param url: 请求URL
// @param headers: 请求头
// @param body: 请求体
// @return []byte: 响应体
// @return int: 响应状态码
// @return error: 错误信息
func (c *Client) SendRequest(method, url string, headers map[string]string, body interface{}) ([]byte, int, error) {
	r := c.client.R()
	// 设置路径参数
	if len(c.PathParams) > 0 {
		r.SetPathParams(c.PathParams)
	}
	if headers != nil {
		r.SetHeaders(headers)
	}
	if body != nil {
		r.SetBody(body)
	}
	// 启用跟踪
	if c.EnableTrace {
		r.EnableTrace()
	}

	resp, err := r.Execute(method, url)
	if err != nil {
		return nil, 0, fmt.Errorf("请求失败: %v", err)
	}
	return resp.Body(), resp.StatusCode(), nil
}

// GetRawClient 获取底层的resty客户端
// @return *resty.Client: 底层的resty客户端
func (c *Client) GetRawClient() *resty.Client {
	return c.client
}

// SetRedirectPolicy 设置重定向策略
// @param redirectPolicy: 重定向策略
// @return *Client: 当前客户端实例，支持链式调用
func (c *Client) SetRedirectPolicy(redirectPolicy resty.RedirectPolicy) *Client {
	c.client.SetRedirectPolicy(redirectPolicy)
	return c
}

// GetTraceInfo 获取请求跟踪信息
// @param resp: 请求响应
// @return resty.TraceInfo: 跟踪信息
func (c *Client) GetTraceInfo(resp *resty.Response) resty.TraceInfo {
	return resp.Request.TraceInfo()
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

// StreamCallback 定义流式响应的回调函数类型
// @param data: 接收到的数据片段
// @param err: 错误信息，如果有的话
// @return bool: 是否继续接收数据，true表示继续，false表示停止
type StreamCallback func(data []byte, err error) bool

// Stream 发送流式请求并处理响应
// @param method: HTTP方法
// @param url: 请求URL
// @param headers: 请求头
// @param body: 请求体
// @param callback: 处理流式响应的回调函数
// @return error: 错误信息
func (c *Client) Stream(method, url string, headers map[string]string, body interface{}, callback StreamCallback) error {
	// 创建请求
	r := c.client.R()

	// 设置路径参数
	if len(c.PathParams) > 0 {
		r.SetPathParams(c.PathParams)
	}

	// 设置请求头
	if headers != nil {
		r.SetHeaders(headers)
	}

	// 设置请求体
	if body != nil {
		r.SetBody(body)
	}

	// 启用跟踪
	if c.EnableTrace {
		r.EnableTrace()
	}

	// 不自动解析响应体，而是直接获取原始响应
	r.SetDoNotParseResponse(true)

	// 执行请求
	resp, err := r.Execute(method, url)
	if err != nil {
		callback(nil, fmt.Errorf("请求失败: %v", err))
		return err
	}

	// 检查响应状态码
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
	// 设置查询参数
	if params != nil {
		for k, v := range params {
			c.client.SetQueryParam(k, fmt.Sprintf("%v", v))
		}
	}

	return c.Stream("GET", url, nil, nil, callback)
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

// GetJSON 发送GET请求并解析JSON响应(使用默认客户端)
// @param url: 请求URL
// @param params: URL查询参数
// @param result: 解析JSON的目标结构体指针
// @return int: 响应状态码
// @return error: 错误信息
func GetJSON(url string, params map[string]interface{}, result interface{}) (int, error) {
	return defaultClient.GetJSON(url, params, result)
}

// PostJSON 发送POST请求并解析JSON响应(使用默认客户端)
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
