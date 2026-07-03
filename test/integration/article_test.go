package integration

import (
	"encoding/json"
	"net/http"
	"testing"

	"blog/internal/model/entity"
	"blog/pkg/response"
)

// 辅助函数：创建测试分类
func createTestCategory(ts *TestServer, t *testing.T, name, slug string) uint {
	t.Helper()
	c := entity.Category{Name: name, Slug: slug}
	ts.DB.Create(&c)
	return c.ID
}

// 辅助函数：创建测试标签
func createTestTag(ts *TestServer, t *testing.T, name, slug string) uint {
	t.Helper()
	tag := entity.Tag{Name: name, Slug: slug}
	ts.DB.Create(&tag)
	return tag.ID
}

func TestArticle_GetList_Empty(t *testing.T) {
	ts := NewTestServer(t)

	w := ts.DoRequest("GET", "/api/v1/articles?page=1&page_size=10")
	if w.Code != http.StatusOK {
		t.Errorf("期望状态码 200, 实际 %d", w.Code)
	}
}

func TestArticle_GetList_WithData(t *testing.T) {
	ts := NewTestServer(t)

	// 创建测试数据
	catID := createTestCategory(ts, t, "技术", "tech")
	ts.DB.Create(&entity.Article{
		Title:      "测试文章1",
		Slug:       "test-article-1",
		Content:    "这是测试内容",
		CategoryID: catID,
		Status:     "published",
	})
	ts.DB.Create(&entity.Article{
		Title:      "测试文章2",
		Slug:       "test-article-2",
		Content:    "这是草稿内容",
		CategoryID: catID,
		Status:     "draft",
	})

	w := ts.DoRequest("GET", "/api/v1/articles?page=1&page_size=10")
	if w.Code != http.StatusOK {
		t.Errorf("期望状态码 200, 实际 %d", w.Code)
	}

	var resp response.Response
	json.Unmarshal(w.Body.Bytes(), &resp)
	if resp.Code != 0 {
		t.Errorf("期望 code=0, 实际 code=%d", resp.Code)
	}
}

func TestArticle_GetDetail_Success(t *testing.T) {
	ts := NewTestServer(t)

	catID := createTestCategory(ts, t, "技术", "tech")
	ts.DB.Create(&entity.Article{
		Title:      "测试文章",
		Slug:       "test-article",
		Content:    "这是测试内容",
		CategoryID: catID,
		Status:     "published",
	})

	w := ts.DoRequest("GET", "/api/v1/articles/test-article")
	if w.Code != http.StatusOK {
		t.Errorf("期望状态码 200, 实际 %d, body=%s", w.Code, w.Body.String())
	}
}

func TestArticle_GetDetail_NotFound(t *testing.T) {
	ts := NewTestServer(t)

	w := ts.DoRequest("GET", "/api/v1/articles/non-existent-slug")
	// 现在返回 404 状态码
	if w.Code != http.StatusNotFound {
		t.Errorf("期望状态码 404, 实际 %d", w.Code)
	}

	var resp response.Response
	json.Unmarshal(w.Body.Bytes(), &resp)
	if resp.Code == 0 {
		t.Error("期望 code != 0（文章不存在）")
	}
}

func TestArticle_GetArchives(t *testing.T) {
	ts := NewTestServer(t)

	catID := createTestCategory(ts, t, "技术", "tech")
	ts.DB.Create(&entity.Article{
		Title:      "归档文章",
		Slug:       "archive-article",
		Content:    "内容",
		CategoryID: catID,
		Status:     "published",
	})

	w := ts.DoRequest("GET", "/api/v1/articles/archives")
	if w.Code != http.StatusOK {
		t.Errorf("期望状态码 200, 实际 %d", w.Code)
	}
}

func TestArticle_Create_Unauthorized(t *testing.T) {
	ts := NewTestServer(t)

	body := `{"title":"新文章","content":"内容","category_id":1}`
	w := ts.DoRequestWithBody("POST", "/api/v1/admin/articles", body)

	if w.Code != http.StatusUnauthorized {
		t.Errorf("期望状态码 401, 实际 %d", w.Code)
	}
}

func TestArticle_Create_Success(t *testing.T) {
	ts := NewTestServer(t)
	ts.Login(t, "admin", "123456")
	catID := createTestCategory(ts, t, "技术", "tech")

	body := `{"title":"新文章","content":"这是文章内容","category_id":` + uintToStr(catID) + `,"status":"published"}`
	w := ts.DoAuthRequestWithBody("POST", "/api/v1/admin/articles", body)

	if w.Code != http.StatusOK {
		t.Errorf("期望状态码 200, 实际 %d, body=%s", w.Code, w.Body.String())
	}

	var resp response.Response
	json.Unmarshal(w.Body.Bytes(), &resp)
	if resp.Code != 0 {
		t.Errorf("期望 code=0, 实际 code=%d, message=%s", resp.Code, resp.Message)
	}
}

func TestArticle_Update_Success(t *testing.T) {
	ts := NewTestServer(t)
	ts.Login(t, "admin", "123456")
	catID := createTestCategory(ts, t, "技术", "tech")

	// 创建文章
	a := entity.Article{
		Title:      "原始标题",
		Slug:       "original-slug",
		Content:    "原始内容",
		CategoryID: catID,
		Status:     "published",
	}
	ts.DB.Create(&a)

	body := `{"title":"更新后的标题","content":"更新后的内容","category_id":` + uintToStr(catID) + `}`
	w := ts.DoAuthRequestWithBody("PUT", "/api/v1/admin/articles/"+uintToStr(a.ID), body)

	if w.Code != http.StatusOK {
		t.Errorf("期望状态码 200, 实际 %d, body=%s", w.Code, w.Body.String())
	}
}

func TestArticle_Delete_Success(t *testing.T) {
	ts := NewTestServer(t)
	ts.Login(t, "admin", "123456")
	catID := createTestCategory(ts, t, "技术", "tech")

	a := entity.Article{
		Title:      "待删除文章",
		Slug:       "to-delete",
		Content:    "内容",
		CategoryID: catID,
		Status:     "published",
	}
	ts.DB.Create(&a)

	w := ts.DoAuthRequest("DELETE", "/api/v1/admin/articles/"+uintToStr(a.ID))
	if w.Code != http.StatusOK {
		t.Errorf("期望状态码 200, 实际 %d, body=%s", w.Code, w.Body.String())
	}
}

func TestArticle_Delete_NotFound(t *testing.T) {
	ts := NewTestServer(t)
	ts.Login(t, "admin", "123456")

	w := ts.DoAuthRequest("DELETE", "/api/v1/admin/articles/99999")
	// 现在返回 404 状态码
	if w.Code != http.StatusNotFound {
		t.Errorf("期望状态码 404, 实际 %d", w.Code)
	}

	var resp response.Response
	json.Unmarshal(w.Body.Bytes(), &resp)
	if resp.Code == 0 {
		t.Error("期望 code != 0（文章不存在）")
	}
}

func TestArticle_BatchDelete(t *testing.T) {
	ts := NewTestServer(t)
	ts.Login(t, "admin", "123456")
	catID := createTestCategory(ts, t, "技术", "tech")

	// 创建多篇文章
	a1 := entity.Article{Title: "文章1", Slug: "article-1", Content: "内容1", CategoryID: catID, Status: "published"}
	a2 := entity.Article{Title: "文章2", Slug: "article-2", Content: "内容2", CategoryID: catID, Status: "published"}
	ts.DB.Create(&a1)
	ts.DB.Create(&a2)

	body := `{"ids":[` + uintToStr(a1.ID) + `,` + uintToStr(a2.ID) + `]}`
	w := ts.DoAuthRequestWithBody("POST", "/api/v1/admin/articles/batch-delete", body)

	if w.Code != http.StatusOK {
		t.Errorf("期望状态码 200, 实际 %d, body=%s", w.Code, w.Body.String())
	}
}

func TestArticle_GetAdminList(t *testing.T) {
	ts := NewTestServer(t)
	ts.Login(t, "admin", "123456")
	catID := createTestCategory(ts, t, "技术", "tech")

	ts.DB.Create(&entity.Article{Title: "管理文章", Slug: "admin-article", Content: "内容", CategoryID: catID, Status: "published"})

	w := ts.DoAuthRequest("GET", "/api/v1/admin/articles?page=1&page_size=10")
	if w.Code != http.StatusOK {
		t.Errorf("期望状态码 200, 实际 %d", w.Code)
	}
}

func TestArticle_GetAdminDetail(t *testing.T) {
	ts := NewTestServer(t)
	ts.Login(t, "admin", "123456")
	catID := createTestCategory(ts, t, "技术", "tech")

	a := entity.Article{Title: "详情文章", Slug: "detail-article", Content: "内容", CategoryID: catID, Status: "published"}
	ts.DB.Create(&a)

	w := ts.DoAuthRequest("GET", "/api/v1/admin/articles/"+uintToStr(a.ID))
	if w.Code != http.StatusOK {
		t.Errorf("期望状态码 200, 实际 %d, body=%s", w.Code, w.Body.String())
	}
}

// uintToStr 将 uint 转换为字符串
func uintToStr(n uint) string {
	if n == 0 {
		return "0"
	}
	digits := ""
	for tmp := n; tmp > 0; tmp /= 10 {
		digits = string(rune('0'+tmp%10)) + digits
	}
	return digits
}
