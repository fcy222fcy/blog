package integration

import (
	"encoding/json"
	"net/http"
	"testing"

	"blog/internal/model/entity"
	"blog/pkg/response"
)

// ========== Comment Tests ==========

func TestComment_Create_Success(t *testing.T) {
	ts := NewTestServer(t)

	// 创建文章
	catID := createTestCategory(ts, t, "技术", "tech")
	a := entity.Article{Title: "评论测试文章", Slug: "comment-test", Content: "内容", CategoryID: catID, Status: "published"}
	ts.DB.Create(&a)

	body := `{"article_id":` + uintToStr(a.ID) + `,"content":"这是一条测试评论","nickname":"测试用户","email":"test@example.com"}`
	w := ts.DoRequestWithBody("POST", "/api/v1/comments", body)

	if w.Code != http.StatusOK {
		t.Errorf("期望状态码 200, 实际 %d, body=%s", w.Code, w.Body.String())
	}

	var resp response.Response
	json.Unmarshal(w.Body.Bytes(), &resp)
	if resp.Code != 0 {
		t.Errorf("期望 code=0, 实际 code=%d, message=%s", resp.Code, resp.Message)
	}
}

func TestComment_Create_Reply(t *testing.T) {
	ts := NewTestServer(t)

	catID := createTestCategory(ts, t, "技术", "tech")
	a := entity.Article{Title: "评论测试文章", Slug: "comment-test-2", Content: "内容", CategoryID: catID, Status: "published"}
	ts.DB.Create(&a)

	parent := entity.Comment{Content: "父评论", Nickname: "父用户", ArticleID: a.ID, Status: "approved"}
	ts.DB.Create(&parent)

	parentID := parent.ID
	body := `{"article_id":` + uintToStr(a.ID) + `,"content":"这是一条回复","nickname":"回复者","parent_id":` + uintToStr(parentID) + `}`
	w := ts.DoRequestWithBody("POST", "/api/v1/comments", body)

	if w.Code != http.StatusOK {
		t.Errorf("期望状态码 200, 实际 %d, body=%s", w.Code, w.Body.String())
	}
}

func TestComment_Create_EmptyContent(t *testing.T) {
	ts := NewTestServer(t)

	body := `{"article_id":1,"content":"","nickname":"用户"}`
	w := ts.DoRequestWithBody("POST", "/api/v1/comments", body)

	if w.Code != http.StatusBadRequest {
		t.Errorf("期望状态码 400, 实际 %d, body=%s", w.Code, w.Body.String())
	}
}

func TestComment_GetByArticle_Success(t *testing.T) {
	ts := NewTestServer(t)

	catID := createTestCategory(ts, t, "技术", "tech")
	a := entity.Article{Title: "评论文章", Slug: "comment-article", Content: "内容", CategoryID: catID, Status: "published"}
	ts.DB.Create(&a)

	ts.DB.Create(&entity.Comment{Content: "评论1", Nickname: "用户1", ArticleID: a.ID, Status: "approved"})
	ts.DB.Create(&entity.Comment{Content: "评论2", Nickname: "用户2", ArticleID: a.ID, Status: "approved"})

	// 使用 slug 查询
	w := ts.DoRequest("GET", "/api/v1/comments/article/comment-article?page=1&page_size=10")
	if w.Code != http.StatusOK {
		t.Errorf("期望状态码 200, 实际 %d, body=%s", w.Code, w.Body.String())
	}
}

func TestComment_GetAdminList_Unauthorized(t *testing.T) {
	ts := NewTestServer(t)

	w := ts.DoRequest("GET", "/api/v1/admin/comments?page=1&page_size=10")
	if w.Code != http.StatusUnauthorized {
		t.Errorf("期望状态码 401, 实际 %d", w.Code)
	}
}

func TestComment_GetAdminList_Success(t *testing.T) {
	ts := NewTestServer(t)
	ts.Login(t, "admin", "123456")

	catID := createTestCategory(ts, t, "技术", "tech")
	a := entity.Article{Title: "评论文章", Slug: "comment-admin", Content: "内容", CategoryID: catID, Status: "published"}
	ts.DB.Create(&a)
	ts.DB.Create(&entity.Comment{Content: "评论", Nickname: "用户", ArticleID: a.ID, Status: "pending"})

	w := ts.DoAuthRequest("GET", "/api/v1/admin/comments?page=1&page_size=10")
	if w.Code != http.StatusOK {
		t.Errorf("期望状态码 200, 实际 %d", w.Code)
	}
}

func TestComment_UpdateStatus_Success(t *testing.T) {
	ts := NewTestServer(t)
	ts.Login(t, "admin", "123456")

	catID := createTestCategory(ts, t, "技术", "tech")
	a := entity.Article{Title: "评论文章", Slug: "comment-status", Content: "内容", CategoryID: catID, Status: "published"}
	ts.DB.Create(&a)
	c := entity.Comment{Content: "待审核评论", Nickname: "用户", ArticleID: a.ID, Status: "pending"}
	ts.DB.Create(&c)

	body := `{"status":"approved"}`
	w := ts.DoAuthRequestWithBody("PUT", "/api/v1/admin/comments/"+uintToStr(c.ID)+"/status", body)

	if w.Code != http.StatusOK {
		t.Errorf("期望状态码 200, 实际 %d, body=%s", w.Code, w.Body.String())
	}
}

func TestComment_Delete_Success(t *testing.T) {
	ts := NewTestServer(t)
	ts.Login(t, "admin", "123456")

	catID := createTestCategory(ts, t, "技术", "tech")
	a := entity.Article{Title: "评论文章", Slug: "comment-delete", Content: "内容", CategoryID: catID, Status: "published"}
	ts.DB.Create(&a)
	c := entity.Comment{Content: "待删除评论", Nickname: "用户", ArticleID: a.ID, Status: "approved"}
	ts.DB.Create(&c)

	w := ts.DoAuthRequest("DELETE", "/api/v1/admin/comments/"+uintToStr(c.ID))
	if w.Code != http.StatusOK {
		t.Errorf("期望状态码 200, 实际 %d, body=%s", w.Code, w.Body.String())
	}
}

// ========== Daily Question Tests ==========

func TestDailyQuestion_GetLatest_Empty(t *testing.T) {
	ts := NewTestServer(t)

	w := ts.DoRequest("GET", "/api/v1/daily-questions/latest")
	// 现在返回 404 状态码（没有每日一问属于 3000-3999 范围的资源错误）
	if w.Code != http.StatusNotFound {
		t.Errorf("期望状态码 404, 实际 %d", w.Code)
	}
}

func TestDailyQuestion_GetLatest_WithData(t *testing.T) {
	ts := NewTestServer(t)

	ts.DB.Create(&entity.DailyQuestion{
		Question: "今天学了什么？",
		Answer:   "学了 Go 语言",
		Date:     "2024-01-01",
		Status:   1,
	})

	w := ts.DoRequest("GET", "/api/v1/daily-questions/latest")
	if w.Code != http.StatusOK {
		t.Errorf("期望状态码 200, 实际 %d, body=%s", w.Code, w.Body.String())
	}
}

func TestDailyQuestion_GetByDate(t *testing.T) {
	ts := NewTestServer(t)

	ts.DB.Create(&entity.DailyQuestion{
		Question: "今天学了什么？",
		Answer:   "学了 Go 语言",
		Date:     "2024-01-01",
		Status:   1,
	})

	w := ts.DoRequest("GET", "/api/v1/daily-questions/date/2024-01-01")
	if w.Code != http.StatusOK {
		t.Errorf("期望状态码 200, 实际 %d, body=%s", w.Code, w.Body.String())
	}
}

func TestDailyQuestion_Like(t *testing.T) {
	ts := NewTestServer(t)

	dq := entity.DailyQuestion{
		Question: "问题",
		Answer:   "答案",
		Date:     "2024-01-01",
		Status:   1,
	}
	ts.DB.Create(&dq)

	w := ts.DoRequestWithBody("POST", "/api/v1/daily-questions/"+uintToStr(dq.ID)+"/like", "")
	if w.Code != http.StatusOK {
		t.Errorf("期望状态码 200, 实际 %d, body=%s", w.Code, w.Body.String())
	}
}

func TestDailyQuestion_Create_Unauthorized(t *testing.T) {
	ts := NewTestServer(t)

	body := `{"question":"新问题","answer":"新答案","date":"2024-01-01"}`
	w := ts.DoRequestWithBody("POST", "/api/v1/admin/daily-questions", body)

	if w.Code != http.StatusUnauthorized {
		t.Errorf("期望状态码 401, 实际 %d", w.Code)
	}
}

func TestDailyQuestion_Create_Success(t *testing.T) {
	ts := NewTestServer(t)
	ts.Login(t, "admin", "123456")

	body := `{"question":"新问题","answer":"新答案","date":"2024-06-01","status":1}`
	w := ts.DoAuthRequestWithBody("POST", "/api/v1/admin/daily-questions", body)

	if w.Code != http.StatusOK {
		t.Errorf("期望状态码 200, 实际 %d, body=%s", w.Code, w.Body.String())
	}
}

func TestDailyQuestion_Update_Success(t *testing.T) {
	ts := NewTestServer(t)
	ts.Login(t, "admin", "123456")

	dq := entity.DailyQuestion{Question: "原始问题", Answer: "原始答案", Date: "2024-01-01", Status: 1}
	ts.DB.Create(&dq)

	body := `{"question":"更新后的问题","answer":"更新后的答案"}`
	w := ts.DoAuthRequestWithBody("PUT", "/api/v1/admin/daily-questions/"+uintToStr(dq.ID), body)

	if w.Code != http.StatusOK {
		t.Errorf("期望状态码 200, 实际 %d, body=%s", w.Code, w.Body.String())
	}
}

func TestDailyQuestion_Delete_Success(t *testing.T) {
	ts := NewTestServer(t)
	ts.Login(t, "admin", "123456")

	dq := entity.DailyQuestion{Question: "待删除问题", Answer: "答案", Date: "2024-01-01", Status: 1}
	ts.DB.Create(&dq)

	w := ts.DoAuthRequest("DELETE", "/api/v1/admin/daily-questions/"+uintToStr(dq.ID))
	if w.Code != http.StatusOK {
		t.Errorf("期望状态码 200, 实际 %d, body=%s", w.Code, w.Body.String())
	}
}

func TestDailyQuestion_Navigation(t *testing.T) {
	ts := NewTestServer(t)

	ts.DB.Create(&entity.DailyQuestion{Question: "第一天", Answer: "答案1", Date: "2024-01-01", Status: 1})
	ts.DB.Create(&entity.DailyQuestion{Question: "第二天", Answer: "答案2", Date: "2024-01-02", Status: 1})
	ts.DB.Create(&entity.DailyQuestion{Question: "第三天", Answer: "答案3", Date: "2024-01-03", Status: 1})

	// 获取前一天
	w := ts.DoRequest("GET", "/api/v1/daily-questions/previous/2024-01-02")
	if w.Code != http.StatusOK {
		t.Errorf("前一天 - 期望状态码 200, 实际 %d, body=%s", w.Code, w.Body.String())
	}

	// 获取后一天
	w = ts.DoRequest("GET", "/api/v1/daily-questions/next/2024-01-02")
	if w.Code != http.StatusOK {
		t.Errorf("后一天 - 期望状态码 200, 实际 %d, body=%s", w.Code, w.Body.String())
	}
}

func TestDailyQuestion_GetAdminList(t *testing.T) {
	ts := NewTestServer(t)
	ts.Login(t, "admin", "123456")

	ts.DB.Create(&entity.DailyQuestion{Question: "管理问题", Answer: "答案", Date: "2024-01-01", Status: 1})

	w := ts.DoAuthRequest("GET", "/api/v1/admin/daily-questions?page=1&page_size=10")
	if w.Code != http.StatusOK {
		t.Errorf("期望状态码 200, 实际 %d", w.Code)
	}
}

// ========== User Tests ==========

func TestUser_GetInfo_Unauthorized(t *testing.T) {
	ts := NewTestServer(t)

	w := ts.DoRequest("GET", "/api/v1/user/info")
	if w.Code != http.StatusUnauthorized {
		t.Errorf("期望状态码 401, 实际 %d", w.Code)
	}
}

func TestUser_GetInfo_Success(t *testing.T) {
	ts := NewTestServer(t)
	ts.Login(t, "admin", "123456")

	w := ts.DoAuthRequest("GET", "/api/v1/user/info")
	if w.Code != http.StatusOK {
		t.Errorf("期望状态码 200, 实际 %d, body=%s", w.Code, w.Body.String())
	}
}

func TestUser_UpdateInfo_Success(t *testing.T) {
	ts := NewTestServer(t)
	ts.Login(t, "admin", "123456")

	body := `{"nickname":"管理员","email":"admin@blog.com","avatar":"https://example.com/avatar.png","bio":"博客作者"}`
	w := ts.DoAuthRequestWithBody("PUT", "/api/v1/user/info", body)

	if w.Code != http.StatusOK {
		t.Errorf("期望状态码 200, 实际 %d, body=%s", w.Code, w.Body.String())
	}
}

func TestUser_UpdateInfo_Unauthorized(t *testing.T) {
	ts := NewTestServer(t)

	body := `{"nickname":"hacker"}`
	w := ts.DoRequestWithBody("PUT", "/api/v1/user/info", body)

	if w.Code != http.StatusUnauthorized {
		t.Errorf("期望状态码 401, 实际 %d", w.Code)
	}
}
