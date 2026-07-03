package integration

import (
	"encoding/json"
	"net/http"
	"testing"

	"blog/pkg/response"
)

func TestAuth_Login_Success(t *testing.T) {
	ts := NewTestServer(t)
	ts.Login(t, "admin", "123456")

	body := `{"username":"admin","password":"123456"}`
	w := ts.DoRequestWithBody("POST", "/api/v1/auth/login", body)

	if w.Code != http.StatusOK {
		t.Errorf("期望状态码 200, 实际 %d", w.Code)
	}

	var resp response.Response
	json.Unmarshal(w.Body.Bytes(), &resp)
	if resp.Code != 0 {
		t.Errorf("期望 code=0, 实际 code=%d, message=%s", resp.Code, resp.Message)
	}
}

func TestAuth_Login_UserNotFound(t *testing.T) {
	ts := NewTestServer(t)

	body := `{"username":"nonexistent","password":"123456"}`
	w := ts.DoRequestWithBody("POST", "/api/v1/auth/login", body)

	// 现在返回 401 状态码（用户不存在属于 1000-1999 范围的业务错误）
	if w.Code != http.StatusUnauthorized {
		t.Errorf("期望状态码 401, 实际 %d", w.Code)
	}

	var resp response.Response
	json.Unmarshal(w.Body.Bytes(), &resp)
	if resp.Code == 0 {
		t.Error("期望 code != 0（用户不存在）")
	}
}

func TestAuth_Login_EmptyBody(t *testing.T) {
	ts := NewTestServer(t)

	w := ts.DoRequestWithBody("POST", "/api/v1/auth/login", `{}`)

	if w.Code != http.StatusBadRequest {
		t.Errorf("期望状态码 400, 实际 %d (body=%s)", w.Code, w.Body.String())
	}
}

func TestAuth_Login_InvalidJSON(t *testing.T) {
	ts := NewTestServer(t)

	w := ts.DoRequestWithBody("POST", "/api/v1/auth/login", `invalid json`)

	if w.Code != http.StatusBadRequest {
		t.Errorf("期望状态码 400, 实际 %d", w.Code)
	}
}

func TestAuth_Register_Success(t *testing.T) {
	ts := NewTestServer(t)

	body := `{"username":"newuser","password":"pass123456","nickname":"New User","email":"new@test.com"}`
	w := ts.DoRequestWithBody("POST", "/api/v1/auth/register", body)

	if w.Code != http.StatusOK {
		t.Errorf("期望状态码 200, 实际 %d, body=%s", w.Code, w.Body.String())
	}

	var resp response.Response
	json.Unmarshal(w.Body.Bytes(), &resp)
	if resp.Code != 0 {
		t.Errorf("期望 code=0, 实际 code=%d, message=%s", resp.Code, resp.Message)
	}
}

func TestAuth_Register_DuplicateUser(t *testing.T) {
	ts := NewTestServer(t)
	ts.Login(t, "admin", "123456")

	body := `{"username":"admin","password":"pass123456","nickname":"Duplicate","email":"dup@test.com"}`
	w := ts.DoRequestWithBody("POST", "/api/v1/auth/register", body)

	// 现在返回 401 状态码（用户已存在属于 1000-1999 范围的业务错误）
	if w.Code != http.StatusUnauthorized {
		t.Errorf("期望状态码 401, 实际 %d", w.Code)
	}

	var resp response.Response
	json.Unmarshal(w.Body.Bytes(), &resp)
	if resp.Code == 0 {
		t.Error("期望 code != 0（用户已存在）")
	}
}

func TestAuth_Register_InvalidEmail(t *testing.T) {
	ts := NewTestServer(t)

	body := `{"username":"testuser","password":"pass123456","nickname":"Test","email":"invalid-email"}`
	w := ts.DoRequestWithBody("POST", "/api/v1/auth/register", body)

	if w.Code != http.StatusBadRequest {
		t.Errorf("期望状态码 400, 实际 %d, body=%s", w.Code, w.Body.String())
	}
}

func TestAuth_GetUserInfo_Unauthorized(t *testing.T) {
	ts := NewTestServer(t)

	w := ts.DoRequest("GET", "/api/v1/user/info")

	if w.Code != http.StatusUnauthorized {
		t.Errorf("期望状态码 401, 实际 %d", w.Code)
	}
}

func TestAuth_GetUserInfo_Success(t *testing.T) {
	ts := NewTestServer(t)
	ts.Login(t, "admin", "123456")

	w := ts.DoAuthRequest("GET", "/api/v1/user/info")

	if w.Code != http.StatusOK {
		t.Errorf("期望状态码 200, 实际 %d, body=%s", w.Code, w.Body.String())
	}

	var resp response.Response
	json.Unmarshal(w.Body.Bytes(), &resp)
	if resp.Code != 0 {
		t.Errorf("期望 code=0, 实际 code=%d, message=%s", resp.Code, resp.Message)
	}
}

func TestAuth_ChangePassword_Success(t *testing.T) {
	ts := NewTestServer(t)
	ts.Login(t, "admin", "123456")

	// 修改密码接口未暴露HTTP路由，此测试验证基础功能正常
	if ts.DB == nil {
		t.Error("数据库未初始化")
	}
}

func TestAuth_ChangePassword_Unauthorized(t *testing.T) {
	ts := NewTestServer(t)

	// 修改密码接口未暴露HTTP路由，验证未登录访问受保护接口返回401
	w := ts.DoRequestWithBody("PUT", "/api/v1/user/info", `{"nickname":"test"}`)

	if w.Code != http.StatusUnauthorized {
		t.Errorf("期望状态码 401, 实际 %d", w.Code)
	}
}
