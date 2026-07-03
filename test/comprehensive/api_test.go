package comprehensive

import (
	"encoding/json"
	"fmt"
	"net/http"
	"os"
	"strings"
	"testing"
	"time"
)

// ========== API测试配置 ==========
type APITestConfig struct {
	BaseURL string
	Token   string
	Client  *http.Client
}

func NewAPITest() *APITestConfig {
	baseURL := os.Getenv("API_BASE_URL")
	if baseURL == "" {
		baseURL = "http://localhost:8888/api/v1"
	}
	return &APITestConfig{
		BaseURL: baseURL,
		Client: &http.Client{
			Timeout: 30 * time.Second,
		},
	}
}

// doRequest 发送HTTP请求
func (api *APITestConfig) doRequest(method, path string, body interface{}) (*http.Response, []byte, error) {
	url := api.BaseURL + path

	var reqBody *strings.Reader
	if body != nil {
		jsonData, _ := json.Marshal(body)
		reqBody = strings.NewReader(string(jsonData))
	} else {
		reqBody = strings.NewReader("")
	}

	req, err := http.NewRequest(method, url, reqBody)
	if err != nil {
		return nil, nil, err
	}
	req.Header.Set("Content-Type", "application/json")
	if api.Token != "" {
		req.Header.Set("Authorization", "Bearer "+api.Token)
	}

	resp, err := api.Client.Do(req)
	if err != nil {
		return nil, nil, err
	}
	defer resp.Body.Close()

	respBody := make([]byte, 0)
	buf := make([]byte, 1024)
	for {
		n, readErr := resp.Body.Read(buf)
		respBody = append(respBody, buf[:n]...)
		if readErr != nil {
			break
		}
	}

	return resp, respBody, nil
}

// login 登录获取token
func (api *APITestConfig) login(t *testing.T) {
	t.Helper()
	loginResp, loginBody, _ := api.doRequest("POST", "/auth/login", map[string]string{
		"username": "admin",
		"password": "123456",
	})
	if loginResp.StatusCode == http.StatusOK {
		var loginResult map[string]interface{}
		json.Unmarshal(loginBody, &loginResult)
		if data, ok := loginResult["data"].(map[string]interface{}); ok {
			if token, ok := data["token"].(string); ok {
				api.Token = token
			}
		}
	}
}

// ========== 认证API测试 ==========

// TestAPI_Auth 认证接口测试
func TestAPI_Auth(t *testing.T) {
	if os.Getenv("E2E_TEST") != "true" {
		t.Skip("跳过API测试，设置 E2E_TEST=true 以运行")
	}
	api := NewAPITest()

	t.Run("登录接口", func(t *testing.T) {
		t.Run("正常登录", func(t *testing.T) {
			resp, body, err := api.doRequest("POST", "/auth/login", map[string]string{
				"username": "admin",
				"password": "123456",
			})
			if err != nil {
				t.Fatalf("登录请求失败: %v", err)
			}
			if resp.StatusCode != http.StatusOK {
				t.Errorf("登录状态码异常: %d", resp.StatusCode)
			}
			var result map[string]interface{}
			json.Unmarshal(body, &result)
			if code, ok := result["code"].(float64); ok && code != 0 {
				t.Errorf("登录业务码异常: %v", result["code"])
			}
		})

		t.Run("密码错误", func(t *testing.T) {
			resp, body, err := api.doRequest("POST", "/auth/login", map[string]string{
				"username": "admin",
				"password": "wrong_password",
			})
			if err != nil {
				t.Fatalf("登录请求失败: %v", err)
			}
			var result map[string]interface{}
			json.Unmarshal(body, &result)
			if resp.StatusCode == http.StatusOK {
				if code, ok := result["code"].(float64); ok && code == 0 {
					t.Error("密码错误时应返回错误码")
				}
			}
		})

		t.Run("用户不存在", func(t *testing.T) {
			resp, body, err := api.doRequest("POST", "/auth/login", map[string]string{
				"username": "nonexistent_user",
				"password": "123456",
			})
			if err != nil {
				t.Fatalf("登录请求失败: %v", err)
			}
			var result map[string]interface{}
			json.Unmarshal(body, &result)
			if resp.StatusCode == http.StatusOK {
				if code, ok := result["code"].(float64); ok && code == 0 {
					t.Error("用户不存在时应返回错误码")
				}
			}
		})

		t.Run("空用户名", func(t *testing.T) {
			resp, body, err := api.doRequest("POST", "/auth/login", map[string]string{
				"username": "",
				"password": "123456",
			})
			if err != nil {
				t.Fatalf("登录请求失败: %v", err)
			}
			var result map[string]interface{}
			json.Unmarshal(body, &result)
			if resp.StatusCode == http.StatusOK {
				if code, ok := result["code"].(float64); ok && code == 0 {
					t.Error("空用户名时应返回错误码")
				}
			}
		})

		t.Run("空密码", func(t *testing.T) {
			resp, body, err := api.doRequest("POST", "/auth/login", map[string]string{
				"username": "admin",
				"password": "",
			})
			if err != nil {
				t.Fatalf("登录请求失败: %v", err)
			}
			var result map[string]interface{}
			json.Unmarshal(body, &result)
			if resp.StatusCode == http.StatusOK {
				if code, ok := result["code"].(float64); ok && code == 0 {
					t.Error("空密码时应返回错误码")
				}
			}
		})

		t.Run("无效JSON格式", func(t *testing.T) {
			resp, _, err := api.doRequest("POST", "/auth/login", "invalid json")
			if err != nil {
				t.Fatalf("登录请求失败: %v", err)
			}
			if resp.StatusCode == http.StatusOK {
				t.Error("无效JSON时应返回错误状态码")
			}
		})
	})

	t.Run("注册接口", func(t *testing.T) {
		t.Run("正常注册", func(t *testing.T) {
			resp, body, err := api.doRequest("POST", "/auth/register", map[string]string{
				"username": fmt.Sprintf("testuser_%d", time.Now().UnixNano()),
				"password": "testpassword123",
				"email":    fmt.Sprintf("test_%d@example.com", time.Now().UnixNano()),
				"nickname": "Test User",
			})
			if err != nil {
				t.Fatalf("注册请求失败: %v", err)
			}
			if resp.StatusCode != http.StatusOK {
				t.Errorf("注册状态码异常: %d", resp.StatusCode)
			}
			var result map[string]interface{}
			json.Unmarshal(body, &result)
			if resp.StatusCode == http.StatusOK {
				if code, ok := result["code"].(float64); ok && code != 0 {
					t.Errorf("注册业务码异常: %v", result["code"])
				}
			}
		})

		t.Run("重复用户名", func(t *testing.T) {
			api.doRequest("POST", "/auth/register", map[string]string{
				"username": "duplicate_user",
				"password": "testpassword123",
				"email":    "duplicate@example.com",
				"nickname": "Duplicate User",
			})
			resp, body, err := api.doRequest("POST", "/auth/register", map[string]string{
				"username": "duplicate_user",
				"password": "testpassword123",
				"email":    "duplicate2@example.com",
				"nickname": "Duplicate User 2",
			})
			if err != nil {
				t.Fatalf("注册请求失败: %v", err)
			}
			var result map[string]interface{}
			json.Unmarshal(body, &result)
			if resp.StatusCode == http.StatusOK {
				if code, ok := result["code"].(float64); ok && code == 0 {
					t.Error("重复用户名时应返回错误码")
				}
			}
		})

		t.Run("无效邮箱格式", func(t *testing.T) {
			resp, body, err := api.doRequest("POST", "/auth/register", map[string]string{
				"username": fmt.Sprintf("testuser_%d", time.Now().UnixNano()),
				"password": "testpassword123",
				"email":    "invalid_email",
				"nickname": "Test User",
			})
			if err != nil {
				t.Fatalf("注册请求失败: %v", err)
			}
			var result map[string]interface{}
			json.Unmarshal(body, &result)
			if resp.StatusCode == http.StatusOK {
				if code, ok := result["code"].(float64); ok && code == 0 {
					t.Error("无效邮箱时应返回错误码")
				}
			}
		})

		t.Run("密码过短", func(t *testing.T) {
			resp, body, err := api.doRequest("POST", "/auth/register", map[string]string{
				"username": fmt.Sprintf("testuser_%d", time.Now().UnixNano()),
				"password": "123",
				"email":    fmt.Sprintf("test_%d@example.com", time.Now().UnixNano()),
				"nickname": "Test User",
			})
			if err != nil {
				t.Fatalf("注册请求失败: %v", err)
			}
			var result map[string]interface{}
			json.Unmarshal(body, &result)
			if resp.StatusCode == http.StatusOK {
				if code, ok := result["code"].(float64); ok && code == 0 {
					t.Error("密码过短时应返回错误码")
				}
			}
		})
	})

	t.Run("获取用户信息接口", func(t *testing.T) {
		t.Run("未登录访问", func(t *testing.T) {
			oldToken := api.Token
			api.Token = ""
			resp, _, err := api.doRequest("GET", "/user/info", nil)
			api.Token = oldToken
			if err != nil {
				t.Fatalf("请求失败: %v", err)
			}
			if resp.StatusCode != http.StatusUnauthorized {
				t.Errorf("未登录应返回401，实际: %d", resp.StatusCode)
			}
		})

		t.Run("已登录访问", func(t *testing.T) {
			api.login(t)
			resp, body, err := api.doRequest("GET", "/user/info", nil)
			if err != nil {
				t.Fatalf("请求失败: %v", err)
			}
			if resp.StatusCode != http.StatusOK {
				t.Errorf("获取用户信息状态码异常: %d", resp.StatusCode)
			}
			var result map[string]interface{}
			json.Unmarshal(body, &result)
			if resp.StatusCode == http.StatusOK {
				if code, ok := result["code"].(float64); ok && code != 0 {
					t.Errorf("获取用户信息业务码异常: %v", result["code"])
				}
			}
		})
	})
}

// ========== 文章API测试 ==========

// TestAPI_Articles 文章接口测试
func TestAPI_Articles(t *testing.T) {
	if os.Getenv("E2E_TEST") != "true" {
		t.Skip("跳过API测试，设置 E2E_TEST=true 以运行")
	}
	api := NewAPITest()
	api.login(t)

	t.Run("获取文章列表", func(t *testing.T) {
		t.Run("正常获取", func(t *testing.T) {
			resp, body, err := api.doRequest("GET", "/articles?page=1&page_size=10", nil)
			if err != nil {
				t.Fatalf("请求失败: %v", err)
			}
			if resp.StatusCode != http.StatusOK {
				t.Errorf("获取文章列表状态码异常: %d", resp.StatusCode)
			}
			var result map[string]interface{}
			json.Unmarshal(body, &result)
			if resp.StatusCode == http.StatusOK {
				if code, ok := result["code"].(float64); ok && code != 0 {
					t.Errorf("获取文章列表业务码异常: %v", result["code"])
				}
			}
		})

		t.Run("分页参数测试", func(t *testing.T) {
			resp, _, err := api.doRequest("GET", "/articles?page=0&page_size=10", nil)
			if err != nil {
				t.Fatalf("请求失败: %v", err)
			}
			if resp.StatusCode != http.StatusOK {
				t.Errorf("无效页码应返回200，实际: %d", resp.StatusCode)
			}

			resp, _, err = api.doRequest("GET", "/articles?page=99999&page_size=10", nil)
			if err != nil {
				t.Fatalf("请求失败: %v", err)
			}
			if resp.StatusCode != http.StatusOK {
				t.Errorf("超大页码应返回200，实际: %d", resp.StatusCode)
			}

			resp, _, err = api.doRequest("GET", "/articles?page=1&page_size=0", nil)
			if err != nil {
				t.Fatalf("请求失败: %v", err)
			}
			if resp.StatusCode != http.StatusOK {
				t.Errorf("无效每页数量应返回200，实际: %d", resp.StatusCode)
			}
			_ = resp
		})

		t.Run("分类筛选", func(t *testing.T) {
			resp, _, err := api.doRequest("GET", "/articles?category_id=1", nil)
			if err != nil {
				t.Fatalf("请求失败: %v", err)
			}
			if resp.StatusCode != http.StatusOK {
				t.Errorf("分类筛选状态码异常: %d", resp.StatusCode)
			}
		})

		t.Run("标签筛选", func(t *testing.T) {
			resp, _, err := api.doRequest("GET", "/articles?tag_id=1", nil)
			if err != nil {
				t.Fatalf("请求失败: %v", err)
			}
			if resp.StatusCode != http.StatusOK {
				t.Errorf("标签筛选状态码异常: %d", resp.StatusCode)
			}
		})
	})

	t.Run("获取文章详情", func(t *testing.T) {
		t.Run("正常获取", func(t *testing.T) {
			resp, body, err := api.doRequest("GET", "/articles/gin-framework-quickstart", nil)
			if err != nil {
				t.Fatalf("请求失败: %v", err)
			}
			if resp.StatusCode != http.StatusOK {
				t.Errorf("获取文章详情状态码异常: %d", resp.StatusCode)
			}
			var result map[string]interface{}
			json.Unmarshal(body, &result)
			if resp.StatusCode == http.StatusOK {
				if code, ok := result["code"].(float64); ok && code != 0 {
					t.Errorf("获取文章详情业务码异常: %v", result["code"])
				}
			}
		})

		t.Run("文章不存在", func(t *testing.T) {
			resp, body, err := api.doRequest("GET", "/articles/nonexistent-article", nil)
			if err != nil {
				t.Fatalf("请求失败: %v", err)
			}
			var result map[string]interface{}
			json.Unmarshal(body, &result)
			if resp.StatusCode == http.StatusOK {
				if code, ok := result["code"].(float64); ok && code == 0 {
					t.Error("文章不存在时应返回错误码")
				}
			}
			_ = resp
		})
	})

	t.Run("创建文章", func(t *testing.T) {
		t.Run("正常创建", func(t *testing.T) {
			resp, body, err := api.doRequest("POST", "/admin/articles", map[string]interface{}{
				"title":       fmt.Sprintf("测试文章_%d", time.Now().UnixNano()),
				"content":     "这是测试文章的内容",
				"summary":     "这是测试文章的摘要",
				"category_id": 1,
				"tag_ids":     []int{1},
				"status":      "draft",
			})
			if err != nil {
				t.Fatalf("请求失败: %v", err)
			}
			if resp.StatusCode != http.StatusOK {
				t.Errorf("创建文章状态码异常: %d", resp.StatusCode)
			}
			var result map[string]interface{}
			json.Unmarshal(body, &result)
			if resp.StatusCode == http.StatusOK {
				if code, ok := result["code"].(float64); ok && code != 0 {
					t.Errorf("创建文章业务码异常: %v", result["code"])
				}
			}
		})

		t.Run("标题为空", func(t *testing.T) {
			resp, body, err := api.doRequest("POST", "/admin/articles", map[string]interface{}{
				"title":       "",
				"content":     "这是测试文章的内容",
				"category_id": 1,
				"status":      "draft",
			})
			if err != nil {
				t.Fatalf("请求失败: %v", err)
			}
			var result map[string]interface{}
			json.Unmarshal(body, &result)
			if resp.StatusCode == http.StatusOK {
				if code, ok := result["code"].(float64); ok && code == 0 {
					t.Error("标题为空时应返回错误码")
				}
			}
			_ = resp
		})

		t.Run("未登录创建", func(t *testing.T) {
			oldToken := api.Token
			api.Token = ""
			resp, _, err := api.doRequest("POST", "/admin/articles", map[string]interface{}{
				"title":       "测试文章",
				"content":     "内容",
				"category_id": 1,
				"status":      "draft",
			})
			api.Token = oldToken
			if err != nil {
				t.Fatalf("请求失败: %v", err)
			}
			if resp.StatusCode != http.StatusUnauthorized {
				t.Errorf("未登录创建文章应返回401，实际: %d", resp.StatusCode)
			}
		})
	})

	t.Run("更新文章", func(t *testing.T) {
		createResp, createBody, _ := api.doRequest("POST", "/admin/articles", map[string]interface{}{
			"title":       fmt.Sprintf("待更新文章_%d", time.Now().UnixNano()),
			"content":     "原始内容",
			"category_id": 1,
			"status":      "draft",
		})
		var createResult map[string]interface{}
		json.Unmarshal(createBody, &createResult)
		var articleID float64
		if data, ok := createResult["data"].(map[string]interface{}); ok {
			if id, ok := data["id"].(float64); ok {
				articleID = id
			}
		}
		_ = createResp

		if articleID > 0 {
			t.Run("正常更新", func(t *testing.T) {
				resp, body, err := api.doRequest("PUT", fmt.Sprintf("/admin/articles/%d", int(articleID)), map[string]interface{}{
					"title":   "更新后的标题",
					"content": "更新后的内容",
				})
				if err != nil {
					t.Fatalf("请求失败: %v", err)
				}
				if resp.StatusCode != http.StatusOK {
					t.Errorf("更新文章状态码异常: %d", resp.StatusCode)
				}
				var result map[string]interface{}
				json.Unmarshal(body, &result)
				if resp.StatusCode == http.StatusOK {
					if code, ok := result["code"].(float64); ok && code != 0 {
						t.Errorf("更新文章业务码异常: %v", result["code"])
					}
				}
			})
		}
	})

	t.Run("删除文章", func(t *testing.T) {
		createResp, createBody, _ := api.doRequest("POST", "/admin/articles", map[string]interface{}{
			"title":       fmt.Sprintf("待删除文章_%d", time.Now().UnixNano()),
			"content":     "内容",
			"category_id": 1,
			"status":      "draft",
		})
		var createResult map[string]interface{}
		json.Unmarshal(createBody, &createResult)
		var articleID float64
		if data, ok := createResult["data"].(map[string]interface{}); ok {
			if id, ok := data["id"].(float64); ok {
				articleID = id
			}
		}
		_ = createResp

		if articleID > 0 {
			t.Run("正常删除", func(t *testing.T) {
				resp, body, err := api.doRequest("DELETE", fmt.Sprintf("/admin/articles/%d", int(articleID)), nil)
				if err != nil {
					t.Fatalf("请求失败: %v", err)
				}
				if resp.StatusCode != http.StatusOK {
					t.Errorf("删除文章状态码异常: %d", resp.StatusCode)
				}
				var result map[string]interface{}
				json.Unmarshal(body, &result)
				if resp.StatusCode == http.StatusOK {
					if code, ok := result["code"].(float64); ok && code != 0 {
						t.Errorf("删除文章业务码异常: %v", result["code"])
					}
				}
			})

			t.Run("删除不存在的文章", func(t *testing.T) {
				resp, _, err := api.doRequest("DELETE", "/admin/articles/99999", nil)
				if err != nil {
					t.Fatalf("请求失败: %v", err)
				}
				t.Logf("删除不存在的文章返回状态码: %d", resp.StatusCode)
			})
		}
	})
}

// ========== 分类API测试 ==========

// TestAPI_Categories 分类接口测试
func TestAPI_Categories(t *testing.T) {
	if os.Getenv("E2E_TEST") != "true" {
		t.Skip("跳过API测试，设置 E2E_TEST=true 以运行")
	}
	api := NewAPITest()
	api.login(t)

	t.Run("获取分类列表", func(t *testing.T) {
		resp, body, err := api.doRequest("GET", "/categories", nil)
		if err != nil {
			t.Fatalf("请求失败: %v", err)
		}
		if resp.StatusCode != http.StatusOK {
			t.Errorf("获取分类列表状态码异常: %d", resp.StatusCode)
		}
		var result map[string]interface{}
		json.Unmarshal(body, &result)
		if resp.StatusCode == http.StatusOK {
			if code, ok := result["code"].(float64); ok && code != 0 {
				t.Errorf("获取分类列表业务码异常: %v", result["code"])
			}
		}
	})

	t.Run("创建分类", func(t *testing.T) {
		t.Run("正常创建", func(t *testing.T) {
			resp, body, err := api.doRequest("POST", "/admin/categories", map[string]interface{}{
				"name":        fmt.Sprintf("测试分类_%d", time.Now().UnixNano()),
				"slug":        fmt.Sprintf("test-category-%d", time.Now().UnixNano()),
				"description": "这是测试分类",
			})
			if err != nil {
				t.Fatalf("请求失败: %v", err)
			}
			if resp.StatusCode != http.StatusOK {
				t.Errorf("创建分类状态码异常: %d", resp.StatusCode)
			}
			var result map[string]interface{}
			json.Unmarshal(body, &result)
			if resp.StatusCode == http.StatusOK {
				if code, ok := result["code"].(float64); ok && code != 0 {
					t.Errorf("创建分类业务码异常: %v", result["code"])
				}
			}
		})

		t.Run("分类名为空", func(t *testing.T) {
			resp, body, err := api.doRequest("POST", "/admin/categories", map[string]interface{}{
				"name":        "",
				"description": "这是测试分类",
			})
			if err != nil {
				t.Fatalf("请求失败: %v", err)
			}
			var result map[string]interface{}
			json.Unmarshal(body, &result)
			if resp.StatusCode == http.StatusOK {
				if code, ok := result["code"].(float64); ok && code == 0 {
					t.Error("分类名为空时应返回错误码")
				}
			}
			_ = resp
		})
	})

	t.Run("删除分类", func(t *testing.T) {
		createResp, createBody, _ := api.doRequest("POST", "/admin/categories", map[string]interface{}{
			"name": fmt.Sprintf("待删除分类_%d", time.Now().UnixNano()),
		})
		var createResult map[string]interface{}
		json.Unmarshal(createBody, &createResult)
		var categoryID float64
		if data, ok := createResult["data"].(map[string]interface{}); ok {
			if id, ok := data["id"].(float64); ok {
				categoryID = id
			}
		}
		_ = createResp

		if categoryID > 0 {
			t.Run("正常删除", func(t *testing.T) {
				resp, body, err := api.doRequest("DELETE", fmt.Sprintf("/admin/categories/%d", int(categoryID)), nil)
				if err != nil {
					t.Fatalf("请求失败: %v", err)
				}
				if resp.StatusCode != http.StatusOK {
					t.Errorf("删除分类状态码异常: %d", resp.StatusCode)
				}
				var result map[string]interface{}
				json.Unmarshal(body, &result)
				if resp.StatusCode == http.StatusOK {
					if code, ok := result["code"].(float64); ok && code != 0 {
						t.Errorf("删除分类业务码异常: %v", result["code"])
					}
				}
			})
		}
	})
}

// ========== 标签API测试 ==========

// TestAPI_Tags 标签接口测试
func TestAPI_Tags(t *testing.T) {
	if os.Getenv("E2E_TEST") != "true" {
		t.Skip("跳过API测试，设置 E2E_TEST=true 以运行")
	}
	api := NewAPITest()
	api.login(t)

	t.Run("获取标签列表", func(t *testing.T) {
		resp, body, err := api.doRequest("GET", "/tags", nil)
		if err != nil {
			t.Fatalf("请求失败: %v", err)
		}
		if resp.StatusCode != http.StatusOK {
			t.Errorf("获取标签列表状态码异常: %d", resp.StatusCode)
		}
		var result map[string]interface{}
		json.Unmarshal(body, &result)
		if resp.StatusCode == http.StatusOK {
			if code, ok := result["code"].(float64); ok && code != 0 {
				t.Errorf("获取标签列表业务码异常: %v", result["code"])
			}
		}
	})

	t.Run("创建标签", func(t *testing.T) {
		t.Run("正常创建", func(t *testing.T) {
			resp, body, err := api.doRequest("POST", "/admin/tags", map[string]interface{}{
				"name": fmt.Sprintf("测试标签_%d", time.Now().UnixNano()),
				"slug": fmt.Sprintf("test-tag-%d", time.Now().UnixNano()),
			})
			if err != nil {
				t.Fatalf("请求失败: %v", err)
			}
			if resp.StatusCode != http.StatusOK {
				t.Errorf("创建标签状态码异常: %d", resp.StatusCode)
			}
			var result map[string]interface{}
			json.Unmarshal(body, &result)
			if resp.StatusCode == http.StatusOK {
				if code, ok := result["code"].(float64); ok && code != 0 {
					t.Errorf("创建标签业务码异常: %v", result["code"])
				}
			}
		})

		t.Run("标签名为空", func(t *testing.T) {
			resp, body, err := api.doRequest("POST", "/admin/tags", map[string]interface{}{
				"name": "",
			})
			if err != nil {
				t.Fatalf("请求失败: %v", err)
			}
			var result map[string]interface{}
			json.Unmarshal(body, &result)
			if resp.StatusCode == http.StatusOK {
				if code, ok := result["code"].(float64); ok && code == 0 {
					t.Error("标签名为空时应返回错误码")
				}
			}
			_ = resp
		})
	})

	t.Run("删除标签", func(t *testing.T) {
		createResp, createBody, _ := api.doRequest("POST", "/admin/tags", map[string]interface{}{
			"name": fmt.Sprintf("待删除标签_%d", time.Now().UnixNano()),
		})
		var createResult map[string]interface{}
		json.Unmarshal(createBody, &createResult)
		var tagID float64
		if data, ok := createResult["data"].(map[string]interface{}); ok {
			if id, ok := data["id"].(float64); ok {
				tagID = id
			}
		}
		_ = createResp

		if tagID > 0 {
			t.Run("正常删除", func(t *testing.T) {
				resp, body, err := api.doRequest("DELETE", fmt.Sprintf("/admin/tags/%d", int(tagID)), nil)
				if err != nil {
					t.Fatalf("请求失败: %v", err)
				}
				if resp.StatusCode != http.StatusOK {
					t.Errorf("删除标签状态码异常: %d", resp.StatusCode)
				}
				var result map[string]interface{}
				json.Unmarshal(body, &result)
				if resp.StatusCode == http.StatusOK {
					if code, ok := result["code"].(float64); ok && code != 0 {
						t.Errorf("删除标签业务码异常: %v", result["code"])
					}
				}
			})
		}
	})
}

// ========== 评论API测试 ==========

// TestAPI_Comments 评论接口测试
func TestAPI_Comments(t *testing.T) {
	if os.Getenv("E2E_TEST") != "true" {
		t.Skip("跳过API测试，设置 E2E_TEST=true 以运行")
	}
	api := NewAPITest()

	t.Run("获取文章评论", func(t *testing.T) {
		resp, body, err := api.doRequest("GET", "/comments/article/1", nil)
		if err != nil {
			t.Fatalf("请求失败: %v", err)
		}
		if resp.StatusCode != http.StatusOK {
			t.Errorf("获取文章评论状态码异常: %d", resp.StatusCode)
		}
		var result map[string]interface{}
		json.Unmarshal(body, &result)
		if resp.StatusCode == http.StatusOK {
			if code, ok := result["code"].(float64); ok && code != 0 {
				t.Errorf("获取文章评论业务码异常: %v", result["code"])
			}
		}
	})

	t.Run("创建评论", func(t *testing.T) {
		t.Run("正常创建", func(t *testing.T) {
			resp, body, err := api.doRequest("POST", "/comments", map[string]interface{}{
				"article_id": 1,
				"content":    "这是一条测试评论",
				"nickname":   "测试用户",
				"email":      "test@example.com",
			})
			if err != nil {
				t.Fatalf("请求失败: %v", err)
			}
			if resp.StatusCode != http.StatusOK {
				t.Errorf("创建评论状态码异常: %d", resp.StatusCode)
			}
			var result map[string]interface{}
			json.Unmarshal(body, &result)
			if resp.StatusCode == http.StatusOK {
				if code, ok := result["code"].(float64); ok && code != 0 {
					t.Errorf("创建评论业务码异常: %v", result["code"])
				}
			}
		})

		t.Run("评论内容为空", func(t *testing.T) {
			resp, body, err := api.doRequest("POST", "/comments", map[string]interface{}{
				"article_id": 1,
				"content":    "",
				"nickname":   "测试用户",
				"email":      "test@example.com",
			})
			if err != nil {
				t.Fatalf("请求失败: %v", err)
			}
			var result map[string]interface{}
			json.Unmarshal(body, &result)
			if resp.StatusCode == http.StatusOK {
				if code, ok := result["code"].(float64); ok && code == 0 {
					t.Error("评论内容为空时应返回错误码")
				}
			}
			_ = resp
		})

		t.Run("文章不存在", func(t *testing.T) {
			resp, body, err := api.doRequest("POST", "/comments", map[string]interface{}{
				"article_id": 99999,
				"content":    "这是一条测试评论",
				"nickname":   "测试用户",
				"email":      "test@example.com",
			})
			if err != nil {
				t.Fatalf("请求失败: %v", err)
			}
			var result map[string]interface{}
			json.Unmarshal(body, &result)
			if resp.StatusCode == http.StatusOK {
				if code, ok := result["code"].(float64); ok && code == 0 {
					t.Error("文章不存在时应返回错误码")
				}
			}
			_ = resp
		})
	})
}

// ========== 友链API测试 ==========

// TestAPI_Links 友链接口测试
func TestAPI_Links(t *testing.T) {
	if os.Getenv("E2E_TEST") != "true" {
		t.Skip("跳过API测试，设置 E2E_TEST=true 以运行")
	}
	api := NewAPITest()
	api.login(t)

	t.Run("获取友链列表", func(t *testing.T) {
		resp, body, err := api.doRequest("GET", "/links", nil)
		if err != nil {
			t.Fatalf("请求失败: %v", err)
		}
		if resp.StatusCode != http.StatusOK {
			t.Errorf("获取友链列表状态码异常: %d", resp.StatusCode)
		}
		var result map[string]interface{}
		json.Unmarshal(body, &result)
		if resp.StatusCode == http.StatusOK {
			if code, ok := result["code"].(float64); ok && code != 0 {
				t.Errorf("获取友链列表业务码异常: %v", result["code"])
			}
		}
	})

	t.Run("创建友链", func(t *testing.T) {
		t.Run("正常创建", func(t *testing.T) {
			resp, body, err := api.doRequest("POST", "/admin/links", map[string]interface{}{
				"name":        fmt.Sprintf("测试友链_%d", time.Now().UnixNano()),
				"url":         "https://test-example.com",
				"description": "这是测试友链",
			})
			if err != nil {
				t.Fatalf("请求失败: %v", err)
			}
			if resp.StatusCode != http.StatusOK {
				t.Errorf("创建友链状态码异常: %d", resp.StatusCode)
			}
			var result map[string]interface{}
			json.Unmarshal(body, &result)
			if resp.StatusCode == http.StatusOK {
				if code, ok := result["code"].(float64); ok && code != 0 {
					t.Errorf("创建友链业务码异常: %v", result["code"])
				}
			}
		})

		t.Run("名称为空", func(t *testing.T) {
			resp, body, err := api.doRequest("POST", "/admin/links", map[string]interface{}{
				"name": "",
				"url":  "https://test-example.com",
			})
			if err != nil {
				t.Fatalf("请求失败: %v", err)
			}
			var result map[string]interface{}
			json.Unmarshal(body, &result)
			if resp.StatusCode == http.StatusOK {
				if code, ok := result["code"].(float64); ok && code == 0 {
					t.Error("名称为空时应返回错误码")
				}
			}
			_ = resp
		})
	})

	t.Run("删除友链", func(t *testing.T) {
		createResp, createBody, _ := api.doRequest("POST", "/admin/links", map[string]interface{}{
			"name": fmt.Sprintf("待删除友链_%d", time.Now().UnixNano()),
			"url":  fmt.Sprintf("https://delete-test-%d.com", time.Now().UnixNano()),
		})
		var createResult map[string]interface{}
		json.Unmarshal(createBody, &createResult)
		var linkID float64
		if data, ok := createResult["data"].(map[string]interface{}); ok {
			if id, ok := data["id"].(float64); ok {
				linkID = id
			}
		}
		_ = createResp

		if linkID > 0 {
			t.Run("正常删除", func(t *testing.T) {
				resp, body, err := api.doRequest("DELETE", fmt.Sprintf("/admin/links/%d", int(linkID)), nil)
				if err != nil {
					t.Fatalf("请求失败: %v", err)
				}
				if resp.StatusCode != http.StatusOK {
					t.Errorf("删除友链状态码异常: %d", resp.StatusCode)
				}
				var result map[string]interface{}
				json.Unmarshal(body, &result)
				if resp.StatusCode == http.StatusOK {
					if code, ok := result["code"].(float64); ok && code != 0 {
						t.Errorf("删除友链业务码异常: %v", result["code"])
					}
				}
			})
		}
	})
}

// ========== 每日一问API测试 ==========

// TestAPI_DailyQuestions 每日一问接口测试
func TestAPI_DailyQuestions(t *testing.T) {
	if os.Getenv("E2E_TEST") != "true" {
		t.Skip("跳过API测试，设置 E2E_TEST=true 以运行")
	}
	api := NewAPITest()
	api.login(t)

	t.Run("获取最新每日一问", func(t *testing.T) {
		resp, body, err := api.doRequest("GET", "/daily-questions/latest", nil)
		if err != nil {
			t.Fatalf("请求失败: %v", err)
		}
		if resp.StatusCode != http.StatusOK {
			t.Errorf("获取最新每日一问状态码异常: %d", resp.StatusCode)
		}
		var result map[string]interface{}
		json.Unmarshal(body, &result)
		if resp.StatusCode == http.StatusOK {
			if code, ok := result["code"].(float64); ok && code != 0 {
				t.Errorf("获取最新每日一问业务码异常: %v", result["code"])
			}
		}
	})

	t.Run("创建每日一问", func(t *testing.T) {
		t.Run("正常创建", func(t *testing.T) {
			resp, body, err := api.doRequest("POST", "/admin/daily-questions", map[string]interface{}{
				"question": fmt.Sprintf("测试问题_%d", time.Now().UnixNano()),
				"answer":   "测试答案",
				"date":     time.Now().Format("2006-01-02"),
				"status":   1,
			})
			if err != nil {
				t.Fatalf("请求失败: %v", err)
			}
			if resp.StatusCode != http.StatusOK {
				t.Errorf("创建每日一问状态码异常: %d", resp.StatusCode)
			}
			var result map[string]interface{}
			json.Unmarshal(body, &result)
			if resp.StatusCode == http.StatusOK {
				if code, ok := result["code"].(float64); ok && code != 0 {
					t.Errorf("创建每日一问业务码异常: %v", result["code"])
				}
			}
		})

		t.Run("问题为空", func(t *testing.T) {
			resp, body, err := api.doRequest("POST", "/admin/daily-questions", map[string]interface{}{
				"question": "",
				"answer":   "测试答案",
				"date":     time.Now().Format("2006-01-02"),
			})
			if err != nil {
				t.Fatalf("请求失败: %v", err)
			}
			var result map[string]interface{}
			json.Unmarshal(body, &result)
			if resp.StatusCode == http.StatusOK {
				if code, ok := result["code"].(float64); ok && code == 0 {
					t.Error("问题为空时应返回错误码")
				}
			}
			_ = resp
		})
	})

	t.Run("点赞每日一问", func(t *testing.T) {
		_, body, err := api.doRequest("GET", "/daily-questions/latest", nil)
		if err != nil {
			t.Fatalf("请求失败: %v", err)
		}
		var result map[string]interface{}
		json.Unmarshal(body, &result)
		if data, ok := result["data"].(map[string]interface{}); ok {
			if id, ok := data["id"].(float64); ok {
				likeResp, _, err := api.doRequest("POST", fmt.Sprintf("/daily-questions/%d/like", int(id)), nil)
				if err != nil {
					t.Fatalf("点赞请求失败: %v", err)
				}
				if likeResp.StatusCode != http.StatusOK {
					t.Errorf("点赞状态码异常: %d", likeResp.StatusCode)
				}
			}
		}
	})
}
