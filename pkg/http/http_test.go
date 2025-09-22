package http

import (
	"fmt"
	"io"
	"net/http"
	"net/http/httptest"
	"os"
	"path/filepath"
	"testing"
	"time"

	"github.com/go-resty/resty/v2"
	"github.com/stretchr/testify/assert"
)

// 测试客户端创建和基本配置
func TestClientCreation(t *testing.T) {
	// 测试基本客户端创建
	client := NewClient()
	assert.NotNil(t, client, "客户端创建失败")

	// 测试自定义超时客户端
	clientWithTimeout := NewClientWithTimeout(5000)
	assert.NotNil(t, clientWithTimeout, "自定义超时客户端创建失败")

	// 测试链式调用设置
	client = client.SetHeader("X-Test", "value").
		SetTimeout(10000).
		SetRetry(3, 100, 1000)
	assert.NotNil(t, client, "链式调用失败")

	// 测试获取原始客户端
	rawClient := client.GetRawClient()
	assert.NotNil(t, rawClient, "获取原始客户端失败")
}

// 测试HTTP方法
func TestHTTPMethods(t *testing.T) {
	// 创建测试服务器
	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")

		switch r.Method {
		case "GET":
			query := r.URL.Query()
			name := query.Get("name")
			if name != "" {
				w.Write([]byte(fmt.Sprintf(`{"message":"Hello, %s"}`, name)))
			} else {
				w.Write([]byte(`{"message":"Hello, World"}`))
			}
		case "POST":
			w.WriteHeader(http.StatusCreated)
			w.Write([]byte(`{"status":"created"}`))
		case "PUT":
			w.Write([]byte(`{"status":"updated"}`))
		case "DELETE":
			w.Write([]byte(`{"status":"deleted"}`))
		case "HEAD":
			w.Header().Set("X-Custom", "value")
		case "PATCH":
			w.Write([]byte(`{"status":"patched"}`))
		default:
			w.WriteHeader(http.StatusMethodNotAllowed)
		}
	}))
	defer server.Close()

	client := NewClient()

	// 测试GET方法
	body, statusCode, err := client.GET(server.URL, map[string]interface{}{"name": "Test"})
	assert.NoError(t, err, "GET请求失败")
	assert.Equal(t, http.StatusOK, statusCode, "GET响应状态码错误")
	assert.Contains(t, string(body), "Hello, Test", "GET响应内容错误")

	// 测试POST方法
	body, statusCode, err = client.POST(server.URL, map[string]interface{}{"key": "value"})
	assert.NoError(t, err, "POST请求失败")
	assert.Equal(t, http.StatusCreated, statusCode, "POST响应状态码错误")
	assert.Contains(t, string(body), "created", "POST响应内容错误")

	// 测试PUT方法
	body, statusCode, err = client.PUT(server.URL, map[string]interface{}{"key": "value"})
	assert.NoError(t, err, "PUT请求失败")
	assert.Equal(t, http.StatusOK, statusCode, "PUT响应状态码错误")
	assert.Contains(t, string(body), "updated", "PUT响应内容错误")

	// 测试DELETE方法
	body, statusCode, err = client.DELETE(server.URL)
	assert.NoError(t, err, "DELETE请求失败")
	assert.Equal(t, http.StatusOK, statusCode, "DELETE响应状态码错误")
	assert.Contains(t, string(body), "deleted", "DELETE响应内容错误")

	// 测试HEAD方法
	header, statusCode, err := client.HEAD(server.URL)
	assert.NoError(t, err, "HEAD请求失败")
	assert.Equal(t, http.StatusOK, statusCode, "HEAD响应状态码错误")
	assert.Equal(t, "value", header.Get("X-Custom"), "HEAD响应头错误")

	// 测试PATCH方法
	body, statusCode, err = client.PATCH(server.URL, map[string]interface{}{"key": "value"})
	assert.NoError(t, err, "PATCH请求失败")
	assert.Equal(t, http.StatusOK, statusCode, "PATCH响应状态码错误")
	assert.Contains(t, string(body), "patched", "PATCH响应内容错误")

	// 测试SendRequest方法
	body, statusCode, err = client.SendRequest("GET", server.URL, map[string]string{"X-Custom-Header": "test"}, nil)
	assert.NoError(t, err, "SendRequest请求失败")
	assert.Equal(t, http.StatusOK, statusCode, "SendRequest响应状态码错误")
}

// 测试JSON解析功能
func TestJSONParsing(t *testing.T) {
	// 创建测试服务器
	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")
		w.Write([]byte(`{"name":"John","age":30,"city":"New York"}`))
	}))
	defer server.Close()

	client := NewClient()

	// 定义响应结构体
	type User struct {
		Name string `json:"name"`
		Age  int    `json:"age"`
		City string `json:"city"`
	}

	// 测试GetJSON
	var user User
	statusCode, err := client.GetJSON(server.URL, nil, &user)
	assert.NoError(t, err, "GetJSON失败")
	assert.Equal(t, http.StatusOK, statusCode, "GetJSON响应状态码错误")
	assert.Equal(t, "John", user.Name, "GetJSON解析错误")
	assert.Equal(t, 30, user.Age, "GetJSON解析错误")
	assert.Equal(t, "New York", user.City, "GetJSON解析错误")

	// 测试PostJSON
	var user2 User
	statusCode, err = client.PostJSON(server.URL, map[string]interface{}{"test": "data"}, &user2)
	assert.NoError(t, err, "PostJSON失败")
	assert.Equal(t, http.StatusOK, statusCode, "PostJSON响应状态码错误")
	assert.Equal(t, "John", user2.Name, "PostJSON解析错误")
}

// 测试文件下载功能
func TestDownloadFile(t *testing.T) {
	// 创建测试文件内容
	testContent := "This is a test file content."

	// 创建测试服务器
	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "text/plain")
		w.Write([]byte(testContent))
	}))
	defer server.Close()

	// 创建临时文件路径
	tempDir := t.TempDir()
	filePath := filepath.Join(tempDir, "test.txt")

	// 测试DownloadFile
	size, err := DownloadFile(server.URL, filePath)
	assert.NoError(t, err, "文件下载失败")
	assert.Equal(t, int64(len(testContent)), size, "下载文件大小错误")

	// 验证文件内容
	content, err := os.ReadFile(filePath)
	assert.NoError(t, err, "读取下载文件失败")
	assert.Equal(t, testContent, string(content), "下载文件内容错误")
}

// 测试文件上传功能
func TestUploadFile(t *testing.T) {
	// 创建临时文件
	tempDir := t.TempDir()
	filePath := filepath.Join(tempDir, "upload.txt")
	testContent := "This is a test upload file content."
	os.WriteFile(filePath, []byte(testContent), 0644)

	// 创建测试服务器
	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		r.ParseMultipartForm(10 << 20) // 10MB
		file, handler, err := r.FormFile("file")
		assert.NoError(t, err, "服务器接收文件失败")
		defer file.Close()

		// 验证文件名
		assert.Equal(t, "upload.txt", handler.Filename, "上传文件名错误")

		// 读取文件内容
		content, err := io.ReadAll(file)
		assert.NoError(t, err, "服务器读取文件内容失败")
		assert.Equal(t, testContent, string(content), "上传文件内容错误")

		// 验证其他表单参数
		assert.Equal(t, "value1", r.FormValue("param1"), "表单参数错误")

		w.Write([]byte(`{"status":"uploaded"}`))
	}))
	defer server.Close()

	// 测试UploadFile
	body, statusCode, err := UploadFile(
		server.URL,
		"file",
		filePath,
		map[string]string{"param1": "value1"},
	)
	assert.NoError(t, err, "文件上传失败")
	assert.Equal(t, http.StatusOK, statusCode, "文件上传响应状态码错误")
	assert.Contains(t, string(body), "uploaded", "文件上传响应内容错误")
}

// 测试认证功能
func TestAuthentication(t *testing.T) {
	// 创建需要认证的测试服务器
	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")

		// 测试基本认证
		user, pass, ok := r.BasicAuth()
		if ok && user == "username" && pass == "password" {
			w.Write([]byte(`{"authenticated": true, "type": "basic"}`))
			return
		}

		// 测试Bearer认证
		auth := r.Header.Get("Authorization")
		if auth == "Bearer token123" {
			w.Write([]byte(`{"authenticated": true, "type": "bearer"}`))
			return
		}

		w.WriteHeader(http.StatusUnauthorized)
		w.Write([]byte(`{"authenticated": false}`))
	}))
	defer server.Close()

	// 测试基本认证
	client := NewClient().SetBasicAuth("username", "password")
	body, statusCode, err := client.GET(server.URL, nil)
	assert.NoError(t, err, "基本认证请求失败")
	assert.Equal(t, http.StatusOK, statusCode, "基本认证响应状态码错误")
	assert.Contains(t, string(body), "basic", "基本认证响应内容错误")

	// 测试Bearer认证
	client = NewClient().SetBearerAuth("token123")
	body, statusCode, err = client.GET(server.URL, nil)
	assert.NoError(t, err, "Bearer认证请求失败")
	assert.Equal(t, http.StatusOK, statusCode, "Bearer认证响应状态码错误")
	assert.Contains(t, string(body), "bearer", "Bearer认证响应内容错误")
}

// 测试快捷函数
func TestShortcutFunctions(t *testing.T) {
	// 创建测试服务器
	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")
		w.Write([]byte(`{"message":"Hello from shortcut"}`))
	}))
	defer server.Close()

	// 测试GET快捷函数
	body, statusCode, err := GET(server.URL, nil)
	assert.NoError(t, err, "GET快捷函数失败")
	assert.Equal(t, http.StatusOK, statusCode, "GET快捷函数响应状态码错误")
	assert.Contains(t, string(body), "Hello from shortcut", "GET快捷函数响应内容错误")

	// 测试POST快捷函数
	body, statusCode, err = POST(server.URL, nil)
	assert.NoError(t, err, "POST快捷函数失败")
	assert.Equal(t, http.StatusOK, statusCode, "POST快捷函数响应状态码错误")
	assert.Contains(t, string(body), "Hello from shortcut", "POST快捷函数响应内容错误")

	// 测试PUT快捷函数
	body, statusCode, err = PUT(server.URL, nil)
	assert.NoError(t, err, "PUT快捷函数失败")
	assert.Equal(t, http.StatusOK, statusCode, "PUT快捷函数响应状态码错误")
	assert.Contains(t, string(body), "Hello from shortcut", "PUT快捷函数响应内容错误")

	// 测试DELETE快捷函数
	body, statusCode, err = DELETE(server.URL)
	assert.NoError(t, err, "DELETE快捷函数失败")
	assert.Equal(t, http.StatusOK, statusCode, "DELETE快捷函数响应状态码错误")
	assert.Contains(t, string(body), "Hello from shortcut", "DELETE快捷函数响应内容错误")

	// 测试GetJSON快捷函数
	type Response struct {
		Message string `json:"message"`
	}
	var resp Response
	statusCode, err = GetJSON(server.URL, nil, &resp)
	assert.NoError(t, err, "GetJSON快捷函数失败")
	assert.Equal(t, http.StatusOK, statusCode, "GetJSON快捷函数响应状态码错误")
	assert.Equal(t, "Hello from shortcut", resp.Message, "GetJSON快捷函数解析错误")

	// 测试PostJSON快捷函数
	var resp2 Response
	statusCode, err = PostJSON(server.URL, nil, &resp2)
	assert.NoError(t, err, "PostJSON快捷函数失败")
	assert.Equal(t, http.StatusOK, statusCode, "PostJSON快捷函数响应状态码错误")
}

// 测试错误处理
func TestErrorHandling(t *testing.T) {
	// 测试连接超时
	client := NewClient().SetTimeout(10)
	_, statusCode, err := client.GET("http://192.0.2.1:8080", nil) // 无效IP地址
	assert.Error(t, err, "连接超时应该返回错误")
	assert.Equal(t, 0, statusCode, "连接超时状态码应该为0")

	// 测试无效URL
	_, statusCode, err = client.GET("http://invalid-url-format", nil)
	assert.Error(t, err, "无效URL应该返回错误")
	assert.Equal(t, 0, statusCode, "无效URL状态码应该为0")
}

// 测试超时设置
func TestTimeout(t *testing.T) {
	// 创建一个会超时的测试服务器
	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		time.Sleep(200 * time.Millisecond) // 睡眠200ms
		w.Write([]byte(`{"message":"delayed response"}`))
	}))
	defer server.Close()

	// 设置50ms超时，应该超时失败
	// 在测试超时时禁用重试机制，以便更准确地测试超时功能
	client := NewClient().SetTimeout(50).SetRetry(0, 0, 0)
	startTime := time.Now()
	_, _, err := client.GET(server.URL, nil)
	duration := time.Since(startTime)

	assert.Error(t, err, "超时设置应该生效")
	// 禁用重试后，断言值可以设置为更接近超时时间的值
	assert.Less(t, duration.Seconds(), 0.1, "请求应该在设置的超时时间内失败")

	// 设置500ms超时，应该成功
	client = NewClient().SetTimeout(500)
	_, statusCode, err := client.GET(server.URL, nil)
	assert.NoError(t, err, "足够的超时时间应该成功")
	assert.Equal(t, http.StatusOK, statusCode, "状态码应该为200")
}

// 测试路径参数功能
func TestPathParams(t *testing.T) {
	// 创建测试服务器
	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")
		path := r.URL.Path
		w.Write([]byte(fmt.Sprintf(`{"path":"%s"}`, path)))
	}))
	defer server.Close()

	// 测试单个路径参数设置
	client := NewClient()
	client.SetPathParam("id", "123")
	client.SetPathParam("name", "test")
	url := server.URL + "/users/{id}/profiles/{name}"
	body, statusCode, err := client.GET(url, nil)
	assert.NoError(t, err, "带路径参数的GET请求失败")
	assert.Equal(t, http.StatusOK, statusCode, "带路径参数的GET响应状态码错误")
	assert.Contains(t, string(body), "/users/123/profiles/test", "路径参数替换错误")

	// 测试多个路径参数设置
	client = NewClient()
	client.SetPathParams(map[string]string{
		"resource":  "articles",
		"articleId": "456",
	})
	url = server.URL + "/api/{resource}/{articleId}"
	body, statusCode, err = client.GET(url, nil)
	assert.NoError(t, err, "带多个路径参数的GET请求失败")
	assert.Equal(t, http.StatusOK, statusCode, "带多个路径参数的GET响应状态码错误")
	assert.Contains(t, string(body), "/api/articles/456", "多个路径参数替换错误")
}

// 测试请求跟踪功能
func TestRequestTracing(t *testing.T) {
	// 创建测试服务器
	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")
		w.Write([]byte(`{"message":"trace test"}`))
	}))
	defer server.Close()

	// 启用请求跟踪
	client := NewClient()
	client.SetEnableTrace(true)
	respBody, statusCode, err := client.GET(server.URL, nil)
	assert.NoError(t, err, "带跟踪的GET请求失败")
	assert.Equal(t, http.StatusOK, statusCode, "带跟踪的GET响应状态码错误")
	assert.Contains(t, string(respBody), "trace test", "带跟踪的GET响应内容错误")

	// 由于我们的GET方法没有直接返回resty.Response对象，
	// 这里无法直接测试GetTraceInfo方法
	// 但我们可以通过SendRequest方法来测试
	rawClient := client.GetRawClient()
	resp, err := rawClient.R().
		EnableTrace().
		Get(server.URL)
	assert.NoError(t, err, "获取原始客户端跟踪信息失败")
	traceInfo := resp.Request.TraceInfo()
	assert.NotZero(t, traceInfo.TotalTime, "跟踪信息总时间不应为零")
}

// 测试中间件功能
func TestMiddlewares(t *testing.T) {
	// 创建测试服务器
	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")
		customHeader := r.Header.Get("X-Custom-Middleware")
		w.Write([]byte(fmt.Sprintf(`{"customHeader":"%s"}`, customHeader)))
	}))
	defer server.Close()

	// 添加请求中间件
	client := NewClient()
	middlewareApplied := false
	client.AddMiddleware(func(c *resty.Client, r *resty.Request) error {
		r.SetHeader("X-Custom-Middleware", "applied")
		middlewareApplied = true
		return nil
	})

	body, statusCode, err := client.GET(server.URL, nil)
	assert.NoError(t, err, "带中间件的GET请求失败")
	assert.Equal(t, http.StatusOK, statusCode, "带中间件的GET响应状态码错误")
	assert.Contains(t, string(body), "applied", "中间件未正确应用")
	assert.True(t, middlewareApplied, "中间件函数未被调用")
}

// 测试自定义重试条件
func TestRetryCondition(t *testing.T) {
	// 创建一个会在第一次请求返回500错误，第二次返回200的测试服务器
	count := 0
	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		count++
		w.Header().Set("Content-Type", "application/json")
		if count == 1 {
			w.WriteHeader(http.StatusInternalServerError)
			w.Write([]byte(`{"error":"retry needed"}`))
		} else {
			w.Write([]byte(fmt.Sprintf(`{"success":true, "count":%d}`, count)))
		}
	}))
	defer server.Close()

	// 设置重试条件
	client := NewClient()
	client.SetRetry(3, 100, 500)
	client.SetRetryCondition(func(r *resty.Response, err error) bool {
		// 只有当状态码为500时才重试
		return r != nil && r.StatusCode() == http.StatusInternalServerError
	})

	// 发送请求，应该重试一次
	body, statusCode, err := client.GET(server.URL, nil)
	assert.NoError(t, err, "带自定义重试条件的GET请求失败")
	assert.Equal(t, http.StatusOK, statusCode, "带自定义重试条件的GET响应状态码错误")
	assert.Contains(t, string(body), `"count":2`, "请求未按预期重试")
}

// 测试请求头设置功能
func TestHeadersSetting(t *testing.T) {
	// 创建测试服务器
	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")
		// 检查请求头
		customHeader := r.Header.Get("X-Custom-Header")
		customHeader2 := r.Header.Get("X-Custom-Header-2")
		w.Write([]byte(fmt.Sprintf(`{"header1":"%s", "header2":"%s"}`, customHeader, customHeader2)))
	}))
	defer server.Close()

	// 测试单个请求头设置
	client := NewClient()
	client.SetHeader("X-Custom-Header", "value1")
	body, statusCode, err := client.GET(server.URL, nil)
	assert.NoError(t, err, "带单个请求头的GET请求失败")
	assert.Equal(t, http.StatusOK, statusCode, "带单个请求头的GET响应状态码错误")
	assert.Contains(t, string(body), `"header1":"value1"`, "单个请求头设置错误")

	// 测试多个请求头设置
	client = NewClient()
	client.SetHeaders(map[string]string{
		"X-Custom-Header":   "value1",
		"X-Custom-Header-2": "value2",
	})
	body, statusCode, err = client.GET(server.URL, nil)
	assert.NoError(t, err, "带多个请求头的GET请求失败")
	assert.Equal(t, http.StatusOK, statusCode, "带多个请求头的GET响应状态码错误")
	assert.Contains(t, string(body), `"header1":"value1"`, "第一个请求头设置错误")
	assert.Contains(t, string(body), `"header2":"value2"`, "第二个请求头设置错误")
}
