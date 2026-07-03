package integration

import (
	"encoding/json"
	"net/http"
	"testing"

	"blog/internal/model/entity"
	"blog/pkg/response"
)

func TestCategory_GetList_Empty(t *testing.T) {
	ts := NewTestServer(t)

	w := ts.DoRequest("GET", "/api/v1/categories")
	if w.Code != http.StatusOK {
		t.Errorf("期望状态码 200, 实际 %d", w.Code)
	}
}

func TestCategory_GetList_WithData(t *testing.T) {
	ts := NewTestServer(t)

	ts.DB.Create(&entity.Category{Name: "技术", Slug: "tech"})
	ts.DB.Create(&entity.Category{Name: "生活", Slug: "life"})

	w := ts.DoRequest("GET", "/api/v1/categories")
	if w.Code != http.StatusOK {
		t.Errorf("期望状态码 200, 实际 %d", w.Code)
	}
}

func TestCategory_GetDetail_Success(t *testing.T) {
	ts := NewTestServer(t)

	c := entity.Category{Name: "技术", Slug: "tech"}
	ts.DB.Create(&c)

	w := ts.DoRequest("GET", "/api/v1/categories/"+uintToStr(c.ID))
	if w.Code != http.StatusOK {
		t.Errorf("期望状态码 200, 实际 %d, body=%s", w.Code, w.Body.String())
	}
}

func TestCategory_GetDetail_NotFound(t *testing.T) {
	ts := NewTestServer(t)

	w := ts.DoRequest("GET", "/api/v1/categories/99999")
	// 现在返回 404 状态码
	if w.Code != http.StatusNotFound {
		t.Errorf("期望状态码 404, 实际 %d", w.Code)
	}

	var resp response.Response
	json.Unmarshal(w.Body.Bytes(), &resp)
	if resp.Code == 0 {
		t.Error("期望 code != 0（分类不存在）")
	}
}

func TestCategory_Create_Unauthorized(t *testing.T) {
	ts := NewTestServer(t)

	body := `{"name":"新分类","slug":"new-category"}`
	w := ts.DoRequestWithBody("POST", "/api/v1/admin/categories", body)

	if w.Code != http.StatusUnauthorized {
		t.Errorf("期望状态码 401, 实际 %d", w.Code)
	}
}

func TestCategory_Create_Success(t *testing.T) {
	ts := NewTestServer(t)
	ts.Login(t, "admin", "123456")

	body := `{"name":"新分类","slug":"new-category","description":"描述"}`
	w := ts.DoAuthRequestWithBody("POST", "/api/v1/admin/categories", body)

	if w.Code != http.StatusOK {
		t.Errorf("期望状态码 200, 实际 %d, body=%s", w.Code, w.Body.String())
	}
}

func TestCategory_Create_DuplicateName(t *testing.T) {
	ts := NewTestServer(t)
	ts.Login(t, "admin", "123456")

	ts.DB.Create(&entity.Category{Name: "技术", Slug: "tech"})

	body := `{"name":"技术","slug":"tech-duplicate"}`
	w := ts.DoAuthRequestWithBody("POST", "/api/v1/admin/categories", body)

	var resp response.Response
	json.Unmarshal(w.Body.Bytes(), &resp)
	if resp.Code == 0 {
		t.Error("期望 code != 0（分类名已存在）")
	}
}

func TestCategory_Update_Success(t *testing.T) {
	ts := NewTestServer(t)
	ts.Login(t, "admin", "123456")

	c := entity.Category{Name: "原始", Slug: "original"}
	ts.DB.Create(&c)

	body := `{"name":"更新后","slug":"updated"}`
	w := ts.DoAuthRequestWithBody("PUT", "/api/v1/admin/categories/"+uintToStr(c.ID), body)

	if w.Code != http.StatusOK {
		t.Errorf("期望状态码 200, 实际 %d, body=%s", w.Code, w.Body.String())
	}
}

func TestCategory_Delete_Success(t *testing.T) {
	ts := NewTestServer(t)
	ts.Login(t, "admin", "123456")

	c := entity.Category{Name: "待删除", Slug: "to-delete"}
	ts.DB.Create(&c)

	w := ts.DoAuthRequest("DELETE", "/api/v1/admin/categories/"+uintToStr(c.ID))
	if w.Code != http.StatusOK {
		t.Errorf("期望状态码 200, 实际 %d, body=%s", w.Code, w.Body.String())
	}
}

func TestCategory_Delete_NotFound(t *testing.T) {
	ts := NewTestServer(t)
	ts.Login(t, "admin", "123456")

	w := ts.DoAuthRequest("DELETE", "/api/v1/admin/categories/99999")
	var resp response.Response
	json.Unmarshal(w.Body.Bytes(), &resp)
	if resp.Code == 0 {
		t.Error("期望 code != 0（分类不存在）")
	}
}

// ========== Tag Tests ==========

func TestTag_GetList_Empty(t *testing.T) {
	ts := NewTestServer(t)

	w := ts.DoRequest("GET", "/api/v1/tags")
	if w.Code != http.StatusOK {
		t.Errorf("期望状态码 200, 实际 %d", w.Code)
	}
}

func TestTag_GetList_WithData(t *testing.T) {
	ts := NewTestServer(t)

	ts.DB.Create(&entity.Tag{Name: "Go", Slug: "go"})
	ts.DB.Create(&entity.Tag{Name: "Gin", Slug: "gin"})

	w := ts.DoRequest("GET", "/api/v1/tags")
	if w.Code != http.StatusOK {
		t.Errorf("期望状态码 200, 实际 %d", w.Code)
	}
}

func TestTag_Create_Unauthorized(t *testing.T) {
	ts := NewTestServer(t)

	body := `{"name":"新标签","slug":"new-tag"}`
	w := ts.DoRequestWithBody("POST", "/api/v1/admin/tags", body)

	if w.Code != http.StatusUnauthorized {
		t.Errorf("期望状态码 401, 实际 %d", w.Code)
	}
}

func TestTag_Create_Success(t *testing.T) {
	ts := NewTestServer(t)
	ts.Login(t, "admin", "123456")

	body := `{"name":"新标签","slug":"new-tag"}`
	w := ts.DoAuthRequestWithBody("POST", "/api/v1/admin/tags", body)

	if w.Code != http.StatusOK {
		t.Errorf("期望状态码 200, 实际 %d, body=%s", w.Code, w.Body.String())
	}
}

func TestTag_Create_DuplicateName(t *testing.T) {
	ts := NewTestServer(t)
	ts.Login(t, "admin", "123456")

	ts.DB.Create(&entity.Tag{Name: "Go", Slug: "go"})

	body := `{"name":"Go","slug":"go-duplicate"}`
	w := ts.DoAuthRequestWithBody("POST", "/api/v1/admin/tags", body)

	var resp response.Response
	json.Unmarshal(w.Body.Bytes(), &resp)
	if resp.Code == 0 {
		t.Error("期望 code != 0（标签名已存在）")
	}
}

func TestTag_Update_Success(t *testing.T) {
	ts := NewTestServer(t)
	ts.Login(t, "admin", "123456")

	tag := entity.Tag{Name: "原始", Slug: "original"}
	ts.DB.Create(&tag)

	body := `{"name":"更新后","slug":"updated"}`
	w := ts.DoAuthRequestWithBody("PUT", "/api/v1/admin/tags/"+uintToStr(tag.ID), body)

	if w.Code != http.StatusOK {
		t.Errorf("期望状态码 200, 实际 %d, body=%s", w.Code, w.Body.String())
	}
}

func TestTag_Delete_Success(t *testing.T) {
	ts := NewTestServer(t)
	ts.Login(t, "admin", "123456")

	tag := entity.Tag{Name: "待删除", Slug: "to-delete"}
	ts.DB.Create(&tag)

	w := ts.DoAuthRequest("DELETE", "/api/v1/admin/tags/"+uintToStr(tag.ID))
	if w.Code != http.StatusOK {
		t.Errorf("期望状态码 200, 实际 %d, body=%s", w.Code, w.Body.String())
	}
}

func TestTag_Delete_NotFound(t *testing.T) {
	ts := NewTestServer(t)
	ts.Login(t, "admin", "123456")

	w := ts.DoAuthRequest("DELETE", "/api/v1/admin/tags/99999")
	var resp response.Response
	json.Unmarshal(w.Body.Bytes(), &resp)
	if resp.Code == 0 {
		t.Error("期望 code != 0（标签不存在）")
	}
}