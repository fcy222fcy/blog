package e2e

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"os"
	"testing"
	"time"
)

// E2ETestConfig 端到端测试配置
type E2ETestConfig struct {
	BaseURL  string
	Username string
	Password string
	Token    string
	Client   *http.Client
}

// NewE2ETest 创建端到端测试实例
func NewE2ETest() *E2ETestConfig {
	baseURL := os.Getenv("E2E_BASE_URL")
	if baseURL == "" {
		baseURL = "http://localhost:8080/api/v1"
	}

	return &E2ETestConfig{
		BaseURL:  baseURL,
		Username: "admin",
		Password: "123456",
		Client: &http.Client{
			Timeout: 30 * time.Second,
		},
	}
}

// doRequest 发送 HTTP 请求
func (e *E2ETestConfig) doRequest(method, path string, body interface{}) (*http.Response, []byte, error) {
	url := e.BaseURL + path

	var reqBody io.Reader
	if body != nil {
		jsonData, err := json.Marshal(body)
		if err != nil {
			return nil, nil, fmt.Errorf("序列化请求体失败: %w", err)
		}
		reqBody = bytes.NewBuffer(jsonData)
	}

	req, err := http.NewRequest(method, url, reqBody)
	if err != nil {
		return nil, nil, fmt.Errorf("创建请求失败: %w", err)
	}

	req.Header.Set("Content-Type", "application/json")
	if e.Token != "" {
		req.Header.Set("Authorization", "Bearer "+e.Token)
	}

	resp, err := e.Client.Do(req)
	if err != nil {
		return nil, nil, fmt.Errorf("发送请求失败: %w", err)
	}
	defer resp.Body.Close()

	respBody, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, nil, fmt.Errorf("读取响应体失败: %w", err)
	}

	return resp, respBody, nil
}

// APIResponse 通用 API 响应
type APIResponse struct {
	Code    int             `json:"code"`
	Message string          `json:"message"`
	Data    json.RawMessage `json:"data"`
}

// ========== E2E 测试用例 ==========

// TestE2E_FullFlow 完整业务流程测试
// 运行方式: go test -v -run TestE2E_FullFlow ./test/e2e/
// 需要先启动服务: go run cmd/server/main.go
func TestE2E_FullFlow(t *testing.T) {
	if os.Getenv("E2E_TEST") != "true" {
		t.Skip("跳过 E2E 测试，设置 E2E_TEST=true 以运行")
	}

	e := NewE2ETest()

	// ========== 1. 认证流程 ==========
	t.Run("01_Login", func(t *testing.T) {
		resp, body, err := e.doRequest("POST", "/auth/login", map[string]string{
			"username": e.Username,
			"password": e.Password,
		})
		if err != nil {
			t.Fatalf("登录请求失败: %v", err)
		}
		if resp.StatusCode != http.StatusOK {
			t.Fatalf("登录失败, status=%d, body=%s", resp.StatusCode, string(body))
		}

		var apiResp APIResponse
		json.Unmarshal(body, &apiResp)
		if apiResp.Code != 0 {
			t.Fatalf("登录失败, code=%d, message=%s", apiResp.Code, apiResp.Message)
		}

		// 提取 Token
		var data map[string]interface{}
		json.Unmarshal(apiResp.Data, &data)
		if token, ok := data["token"].(string); ok {
			e.Token = token
			t.Logf("登录成功, Token: %s...", token[:20])
		} else {
			t.Fatal("无法从响应中提取 Token")
		}
	})

	// ========== 2. 获取用户信息 ==========
	t.Run("02_GetProfile", func(t *testing.T) {
		resp, body, err := e.doRequest("GET", "/auth/profile", nil)
		if err != nil {
			t.Fatalf("获取用户信息失败: %v", err)
		}
		if resp.StatusCode != http.StatusOK {
			t.Fatalf("获取用户信息失败, status=%d, body=%s", resp.StatusCode, string(body))
		}
		t.Logf("用户信息: %s", string(body))
	})

	// ========== 3. 分类 CRUD ==========
	var categoryID uint
	t.Run("03_CreateCategory", func(t *testing.T) {
		resp, body, err := e.doRequest("POST", "/admin/categories", map[string]interface{}{
			"name":        "E2E测试分类",
			"slug":        "e2e-test-category",
			"description": "端到端测试分类",
		})
		if err != nil {
			t.Fatalf("创建分类失败: %v", err)
		}
		if resp.StatusCode != http.StatusOK {
			t.Fatalf("创建分类失败, status=%d, body=%s", resp.StatusCode, string(body))
		}

		var apiResp APIResponse
		json.Unmarshal(body, &apiResp)
		if apiResp.Code != 0 {
			t.Fatalf("创建分类失败, code=%d, message=%s", apiResp.Code, apiResp.Message)
		}

		var data map[string]interface{}
		json.Unmarshal(apiResp.Data, &data)
		if id, ok := data["id"].(float64); ok {
			categoryID = uint(id)
			t.Logf("创建分类成功, ID=%d", categoryID)
		}
	})

	t.Run("04_GetCategories", func(t *testing.T) {
		resp, body, err := e.doRequest("GET", "/categories", nil)
		if err != nil {
			t.Fatalf("获取分类列表失败: %v", err)
		}
		if resp.StatusCode != http.StatusOK {
			t.Fatalf("获取分类列表失败, status=%d", resp.StatusCode)
		}
		t.Logf("分类列表: %s", string(body))
	})

	// ========== 4. 标签 CRUD ==========
	var tagID uint
	t.Run("05_CreateTag", func(t *testing.T) {
		resp, body, err := e.doRequest("POST", "/admin/tags", map[string]string{
			"name": "E2E测试标签",
			"slug": "e2e-test-tag",
		})
		if err != nil {
			t.Fatalf("创建标签失败: %v", err)
		}
		if resp.StatusCode != http.StatusOK {
			t.Fatalf("创建标签失败, status=%d, body=%s", resp.StatusCode, string(body))
		}

		var apiResp APIResponse
		json.Unmarshal(body, &apiResp)
		if apiResp.Code != 0 {
			t.Fatalf("创建标签失败, code=%d, message=%s", apiResp.Code, apiResp.Message)
		}

		var data map[string]interface{}
		json.Unmarshal(apiResp.Data, &data)
		if id, ok := data["id"].(float64); ok {
			tagID = uint(id)
			t.Logf("创建标签成功, ID=%d", tagID)
		}
	})

	// ========== 5. 文章 CRUD ==========
	var articleID uint
	t.Run("06_CreateArticle", func(t *testing.T) {
		resp, body, err := e.doRequest("POST", "/admin/articles", map[string]interface{}{
			"title":       "E2E测试文章",
			"content":     "这是端到端测试的文章内容，用于验证完整的业务流程。",
			"summary":     "E2E测试文章摘要",
			"category_id": categoryID,
			"tag_ids":     []uint{tagID},
			"status":      "published",
		})
		if err != nil {
			t.Fatalf("创建文章失败: %v", err)
		}
		if resp.StatusCode != http.StatusOK {
			t.Fatalf("创建文章失败, status=%d, body=%s", resp.StatusCode, string(body))
		}

		var apiResp APIResponse
		json.Unmarshal(body, &apiResp)
		if apiResp.Code != 0 {
			t.Fatalf("创建文章失败, code=%d, message=%s", apiResp.Code, apiResp.Message)
		}

		var data map[string]interface{}
		json.Unmarshal(apiResp.Data, &data)
		if id, ok := data["id"].(float64); ok {
			articleID = uint(id)
			t.Logf("创建文章成功, ID=%d", articleID)
		}
	})

	t.Run("07_GetArticles", func(t *testing.T) {
		resp, body, err := e.doRequest("GET", "/articles", nil)
		if err != nil {
			t.Fatalf("获取文章列表失败: %v", err)
		}
		if resp.StatusCode != http.StatusOK {
			t.Fatalf("获取文章列表失败, status=%d", resp.StatusCode)
		}
		t.Logf("文章列表: %s", string(body))
	})

	t.Run("08_GetArticleDetail", func(t *testing.T) {
		resp, body, err := e.doRequest("GET", "/articles/e2e-test-article", nil)
		if err != nil {
			t.Fatalf("获取文章详情失败: %v", err)
		}
		if resp.StatusCode != http.StatusOK {
			t.Fatalf("获取文章详情失败, status=%d, body=%s", resp.StatusCode, string(body))
		}
		t.Logf("文章详情: %s", string(body))
	})

	t.Run("09_UpdateArticle", func(t *testing.T) {
		resp, body, err := e.doRequest("PUT", fmt.Sprintf("/admin/articles/%d", articleID), map[string]interface{}{
			"title":   "E2E测试文章（已更新）",
			"content": "更新后的文章内容。",
		})
		if err != nil {
			t.Fatalf("更新文章失败: %v", err)
		}
		if resp.StatusCode != http.StatusOK {
			t.Fatalf("更新文章失败, status=%d, body=%s", resp.StatusCode, string(body))
		}
		t.Logf("更新文章成功")
	})

	// ========== 6. 评论流程 ==========
	var commentID uint
	t.Run("10_CreateComment", func(t *testing.T) {
		resp, body, err := e.doRequest("POST", "/comments", map[string]interface{}{
			"article_id": articleID,
			"content":    "这是一条E2E测试评论",
			"nickname":   "E2E测试用户",
			"email":      "e2e@test.com",
		})
		if err != nil {
			t.Fatalf("创建评论失败: %v", err)
		}
		if resp.StatusCode != http.StatusOK {
			t.Fatalf("创建评论失败, status=%d, body=%s", resp.StatusCode, string(body))
		}

		var apiResp APIResponse
		json.Unmarshal(body, &apiResp)
		if apiResp.Code != 0 {
			t.Fatalf("创建评论失败, code=%d, message=%s", apiResp.Code, apiResp.Message)
		}

		var data map[string]interface{}
		json.Unmarshal(apiResp.Data, &data)
		if id, ok := data["id"].(float64); ok {
			commentID = uint(id)
			t.Logf("创建评论成功, ID=%d", commentID)
		}
	})

	t.Run("11_GetComments", func(t *testing.T) {
		resp, body, err := e.doRequest("GET", fmt.Sprintf("/comments/article/%d", articleID), nil)
		if err != nil {
			t.Fatalf("获取评论列表失败: %v", err)
		}
		if resp.StatusCode != http.StatusOK {
			t.Fatalf("获取评论列表失败, status=%d", resp.StatusCode)
		}
		t.Logf("评论列表: %s", string(body))
	})

	t.Run("12_ApproveComment", func(t *testing.T) {
		resp, body, err := e.doRequest("PUT", fmt.Sprintf("/admin/comments/%d/status", commentID), map[string]string{
			"status": "approved",
		})
		if err != nil {
			t.Fatalf("审核评论失败: %v", err)
		}
		if resp.StatusCode != http.StatusOK {
			t.Fatalf("审核评论失败, status=%d, body=%s", resp.StatusCode, string(body))
		}
		t.Logf("审核评论成功")
	})

	// ========== 7. 友链流程 ==========
	var linkID uint
	t.Run("13_CreateLink", func(t *testing.T) {
		resp, body, err := e.doRequest("POST", "/admin/links", map[string]string{
			"name":        "E2E测试友链",
			"url":         "https://e2e-test.example.com",
			"description": "端到端测试友链",
		})
		if err != nil {
			t.Fatalf("创建友链失败: %v", err)
		}
		if resp.StatusCode != http.StatusOK {
			t.Fatalf("创建友链失败, status=%d, body=%s", resp.StatusCode, string(body))
		}

		var apiResp APIResponse
		json.Unmarshal(body, &apiResp)
		if apiResp.Code != 0 {
			t.Fatalf("创建友链失败, code=%d, message=%s", apiResp.Code, apiResp.Message)
		}

		var data map[string]interface{}
		json.Unmarshal(apiResp.Data, &data)
		if id, ok := data["id"].(float64); ok {
			linkID = uint(id)
			t.Logf("创建友链成功, ID=%d", linkID)
		}
	})

	t.Run("14_GetLinks", func(t *testing.T) {
		resp, _, err := e.doRequest("GET", "/links", nil)
		if err != nil {
			t.Fatalf("获取友链列表失败: %v", err)
		}
		if resp.StatusCode != http.StatusOK {
			t.Fatalf("获取友链列表失败, status=%d", resp.StatusCode)
		}
		t.Logf("获取友链列表成功")
	})

	// ========== 8. 每日一问流程 ==========
	var dailyQID uint
	t.Run("15_CreateDailyQuestion", func(t *testing.T) {
		resp, body, err := e.doRequest("POST", "/admin/daily-questions", map[string]interface{}{
			"question": "E2E测试：今天学了什么？",
			"answer":   "学了 Go 语言的端到端测试",
			"date":     "2024-12-01",
			"status":   1,
		})
		if err != nil {
			t.Fatalf("创建每日一问失败: %v", err)
		}
		if resp.StatusCode != http.StatusOK {
			t.Fatalf("创建每日一问失败, status=%d, body=%s", resp.StatusCode, string(body))
		}

		var apiResp APIResponse
		json.Unmarshal(body, &apiResp)
		if apiResp.Code != 0 {
			t.Fatalf("创建每日一问失败, code=%d, message=%s", apiResp.Code, apiResp.Message)
		}

		var data map[string]interface{}
		json.Unmarshal(apiResp.Data, &data)
		if id, ok := data["id"].(float64); ok {
			dailyQID = uint(id)
			t.Logf("创建每日一问成功, ID=%d", dailyQID)
		}
	})

	t.Run("16_GetLatestDailyQuestion", func(t *testing.T) {
		resp, _, err := e.doRequest("GET", "/daily-questions/latest", nil)
		if err != nil {
			t.Fatalf("获取最新每日一问失败: %v", err)
		}
		if resp.StatusCode != http.StatusOK {
			t.Fatalf("获取最新每日一问失败, status=%d", resp.StatusCode)
		}
		t.Logf("获取最新每日一问成功")
	})

	// ========== 9. 清理测试数据 ==========
	t.Run("17_Cleanup", func(t *testing.T) {
		// 删除评论
		if commentID > 0 {
			_, _, _ = e.doRequest("DELETE", fmt.Sprintf("/admin/comments/%d", commentID), nil)
		}
		// 删除文章
		if articleID > 0 {
			_, _, _ = e.doRequest("DELETE", fmt.Sprintf("/admin/articles/%d", articleID), nil)
		}
		// 删除每日一问
		if dailyQID > 0 {
			_, _, _ = e.doRequest("DELETE", fmt.Sprintf("/admin/daily-questions/%d", dailyQID), nil)
		}
		// 删除友链
		if linkID > 0 {
			_, _, _ = e.doRequest("DELETE", fmt.Sprintf("/admin/links/%d", linkID), nil)
		}
		// 删除标签
		if tagID > 0 {
			_, _, _ = e.doRequest("DELETE", fmt.Sprintf("/admin/tags/%d", tagID), nil)
		}
		// 删除分类
		if categoryID > 0 {
			_, _, _ = e.doRequest("DELETE", fmt.Sprintf("/admin/categories/%d", categoryID), nil)
		}
		t.Logf("清理测试数据完成")
	})
}

// TestE2E_AuthFlow 认证流程测试
func TestE2E_AuthFlow(t *testing.T) {
	if os.Getenv("E2E_TEST") != "true" {
		t.Skip("跳过 E2E 测试，设置 E2E_TEST=true 以运行")
	}

	e := NewE2ETest()

	// 测试登录失败
	t.Run("Login_Failure", func(t *testing.T) {
		resp, body, err := e.doRequest("POST", "/auth/login", map[string]string{
			"username": "admin",
			"password": "wrong_password",
		})
		if err != nil {
			t.Fatalf("请求失败: %v", err)
		}

		var apiResp APIResponse
		json.Unmarshal(body, &apiResp)
		if apiResp.Code == 0 {
			t.Error("期望登录失败但返回了成功")
		}
		t.Logf("登录失败响应: status=%d, code=%d, message=%s", resp.StatusCode, apiResp.Code, apiResp.Message)
	})

	// 测试登录成功
	t.Run("Login_Success", func(t *testing.T) {
		resp, body, err := e.doRequest("POST", "/auth/login", map[string]string{
			"username": e.Username,
			"password": e.Password,
		})
		if err != nil {
			t.Fatalf("请求失败: %v", err)
		}
		if resp.StatusCode != http.StatusOK {
			t.Fatalf("登录失败, status=%d", resp.StatusCode)
		}

		var apiResp APIResponse
		json.Unmarshal(body, &apiResp)
		if apiResp.Code != 0 {
			t.Fatalf("登录失败, code=%d, message=%s", apiResp.Code, apiResp.Message)
		}

		var data map[string]interface{}
		json.Unmarshal(apiResp.Data, &data)
		if token, ok := data["token"].(string); ok {
			e.Token = token
		}
		t.Logf("登录成功")
	})

	// 测试未授权访问
	t.Run("Unauthorized_Access", func(t *testing.T) {
		// 先清除 Token 测试未授权
		oldToken := e.Token
		e.Token = ""
		resp, _, _ := e.doRequest("GET", "/auth/profile", nil)
		if resp.StatusCode != http.StatusUnauthorized {
			t.Errorf("期望 401, 实际 %d", resp.StatusCode)
		}
		e.Token = oldToken
	})
}

// TestE2E_ArticleFlow 文章流程测试
func TestE2E_ArticleFlow(t *testing.T) {
	if os.Getenv("E2E_TEST") != "true" {
		t.Skip("跳过 E2E 测试，设置 E2E_TEST=true 以运行")
	}

	e := NewE2ETest()

	// 先登录
	resp, body, err := e.doRequest("POST", "/auth/login", map[string]string{
		"username": e.Username,
		"password": e.Password,
	})
	if err != nil || resp.StatusCode != http.StatusOK {
		t.Fatal("登录失败，无法继续测试")
	}
	var apiResp APIResponse
	json.Unmarshal(body, &apiResp)
	var data map[string]interface{}
	json.Unmarshal(apiResp.Data, &data)
	e.Token = data["token"].(string)

	// 测试获取文章列表
	t.Run("GetArticles", func(t *testing.T) {
		resp, _, err := e.doRequest("GET", "/articles?page=1&page_size=5", nil)
		if err != nil {
			t.Fatalf("请求失败: %v", err)
		}
		if resp.StatusCode != http.StatusOK {
			t.Errorf("期望 200, 实际 %d", resp.StatusCode)
		}
	})

	// 测试获取文章归档
	t.Run("GetArchives", func(t *testing.T) {
		resp, _, err := e.doRequest("GET", "/articles/archives", nil)
		if err != nil {
			t.Fatalf("请求失败: %v", err)
		}
		if resp.StatusCode != http.StatusOK {
			t.Errorf("期望 200, 实际 %d", resp.StatusCode)
		}
	})

	// 测试获取不存在的文章
	t.Run("GetNonExistentArticle", func(t *testing.T) {
		resp, body, err := e.doRequest("GET", "/articles/non-existent-slug-12345", nil)
		if err != nil {
			t.Fatalf("请求失败: %v", err)
		}
		json.Unmarshal(body, &apiResp)
		if apiResp.Code == 0 {
			t.Error("期望返回错误但返回了成功")
		}
		_ = resp
	})
}

// TestE2E_HealthCheck 健康检查
func TestE2E_HealthCheck(t *testing.T) {
	if os.Getenv("E2E_TEST") != "true" {
		t.Skip("跳过 E2E 测试，设置 E2E_TEST=true 以运行")
	}

	e := NewE2ETest()

	// 测试服务是否可达
	resp, _, err := e.doRequest("GET", "/articles", nil)
	if err != nil {
		t.Fatalf("服务不可达: %v", err)
	}
	if resp.StatusCode == http.StatusOK {
		t.Logf("服务健康检查通过, status=%d", resp.StatusCode)
	} else {
		t.Errorf("服务响应异常, status=%d", resp.StatusCode)
	}
}