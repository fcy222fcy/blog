package unit

import (
	"testing"

	"blog/internal/model/dto/request"
)

func TestPageRequest_GetOffset(t *testing.T) {
	tests := []struct {
		name     string
		page     int
		pageSize int
		want     int
	}{
		{"第1页", 1, 10, 0},
		{"第2页", 2, 10, 10},
		{"第3页", 3, 20, 40},
		{"零值", 0, 10, 0},
		{"负数", -1, 10, 0},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			p := &request.PageRequest{Page: tt.page, PageSize: tt.pageSize}
			if got := p.GetOffset(); got != tt.want {
				t.Errorf("GetOffset() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestPageRequest_GetPageSize(t *testing.T) {
	tests := []struct {
		name     string
		pageSize int
		want     int
	}{
		{"正常值", 10, 10},
		{"零值使用默认", 0, 10},
		{"超过最大值", 200, 100},
		{"边界值100", 100, 100},
		{"边界值101", 101, 100},
		{"负数使用默认", -1, 10},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			p := &request.PageRequest{Page: 1, PageSize: tt.pageSize}
			if got := p.GetPageSize(); got != tt.want {
				t.Errorf("GetPageSize() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestLoginRequest_Validation(t *testing.T) {
	// 验证 LoginRequest 结构体字段
	req := request.LoginRequest{
		Username: "admin",
		Password: "123456",
	}
	if req.Username != "admin" {
		t.Errorf("Username = %v, want admin", req.Username)
	}
	if req.Password != "123456" {
		t.Errorf("Password = %v, want 123456", req.Password)
	}
}

func TestRegisterRequest_Validation(t *testing.T) {
	req := request.RegisterRequest{
		Username: "newuser",
		Password: "password123",
		Nickname: "New User",
		Email:    "newuser@example.com",
	}
	if req.Username != "newuser" {
		t.Errorf("Username = %v, want newuser", req.Username)
	}
	if req.Email != "newuser@example.com" {
		t.Errorf("Email = %v, want newuser@example.com", req.Email)
	}
}

func TestCreateArticleRequest_Fields(t *testing.T) {
	req := request.CreateArticleRequest{
		Title:      "测试文章",
		Content:    "这是测试内容",
		Summary:    "摘要",
		Cover:      "https://example.com/cover.jpg",
		CategoryID: 1,
		TagIDs:     []uint{1, 2, 3},
		Status:     "published",
		IsTop:      true,
	}
	if req.Title != "测试文章" {
		t.Errorf("Title = %v", req.Title)
	}
	if req.CategoryID != 1 {
		t.Errorf("CategoryID = %v, want 1", req.CategoryID)
	}
	if len(req.TagIDs) != 3 {
		t.Errorf("len(TagIDs) = %v, want 3", len(req.TagIDs))
	}
	if req.Status != "published" {
		t.Errorf("Status = %v, want published", req.Status)
	}
	if !req.IsTop {
		t.Errorf("IsTop = false, want true")
	}
}

func TestUpdateArticleRequest_Fields(t *testing.T) {
	req := request.UpdateArticleRequest{
		Title:      "更新后的标题",
		Content:    "更新后的内容",
		CategoryID: 2,
		Status:     "draft",
		IsTop:      false,
	}
	if req.Title != "更新后的标题" {
		t.Errorf("Title = %v", req.Title)
	}
	if req.Status != "draft" {
		t.Errorf("Status = %v, want draft", req.Status)
	}
}

func TestCreateCommentRequest_Fields(t *testing.T) {
	req := request.CreateCommentRequest{
		ArticleID: 1,
		Content:   "这是一条评论",
		Nickname:  "评论者",
		Email:     "commenter@example.com",
		Website:   "https://example.com",
	}
	if req.ArticleID != 1 {
		t.Errorf("ArticleID = %v, want 1", req.ArticleID)
	}
	if req.Content != "这是一条评论" {
		t.Errorf("Content = %v", req.Content)
	}
	if req.Nickname != "评论者" {
		t.Errorf("Nickname = %v", req.Nickname)
	}
}

func TestCreateCommentRequest_WithParentID(t *testing.T) {
	parentID := uint(5)
	req := request.CreateCommentRequest{
		ArticleID: 1,
		Content:   "回复评论",
		Nickname:  "回复者",
		ParentID:  &parentID,
	}
	if req.ParentID == nil || *req.ParentID != 5 {
		t.Errorf("ParentID = %v, want 5", req.ParentID)
	}
}

func TestCreateCategoryRequest_Fields(t *testing.T) {
	req := request.CreateCategoryRequest{
		Name:        "技术",
		Slug:        "tech",
		Description: "技术相关文章",
		Icon:        "🔧",
		SortOrder:   1,
	}
	if req.Name != "技术" {
		t.Errorf("Name = %v", req.Name)
	}
	if req.Slug != "tech" {
		t.Errorf("Slug = %v, want tech", req.Slug)
	}
}

func TestCreateTagRequest_Fields(t *testing.T) {
	req := request.CreateTagRequest{
		Name: "Go",
		Slug: "go",
	}
	if req.Name != "Go" {
		t.Errorf("Name = %v", req.Name)
	}
	if req.Slug != "go" {
		t.Errorf("Slug = %v, want go", req.Slug)
	}
}

func TestCreateLinkRequest_Fields(t *testing.T) {
	req := request.CreateLinkRequest{
		Name:        "示例网站",
		URL:         "https://example.com",
		Description: "一个示例网站",
		Avatar:      "https://example.com/avatar.jpg",
		Logo:        "🔗",
		SortOrder:   1,
		Status:      "approved",
	}
	if req.Name != "示例网站" {
		t.Errorf("Name = %v", req.Name)
	}
	if req.URL != "https://example.com" {
		t.Errorf("URL = %v", req.URL)
	}
	if req.Status != "approved" {
		t.Errorf("Status = %v, want approved", req.Status)
	}
}

func TestCreateDailyQuestionRequest_Fields(t *testing.T) {
	req := request.CreateDailyQuestionRequest{
		Question: "今天学了什么？",
		Answer:   "学了 Go 语言",
		Date:     "2024-01-01",
		Status:   1,
	}
	if req.Question != "今天学了什么？" {
		t.Errorf("Question = %v", req.Question)
	}
	if req.Date != "2024-01-01" {
		t.Errorf("Date = %v", req.Date)
	}
	if req.Status != 1 {
		t.Errorf("Status = %v, want 1", req.Status)
	}
}

func TestChangePasswordRequest_Fields(t *testing.T) {
	req := request.ChangePasswordRequest{
		OldPassword: "old123456",
		NewPassword: "new123456",
	}
	if req.OldPassword != "old123456" {
		t.Errorf("OldPassword = %v", req.OldPassword)
	}
	if req.NewPassword != "new123456" {
		t.Errorf("NewPassword = %v", req.NewPassword)
	}
}

func TestArticleListRequest_Defaults(t *testing.T) {
	req := request.ArticleListRequest{
		PageRequest: request.PageRequest{Page: 1, PageSize: 10},
		Category:    0,
		Tag:         0,
		Keyword:     "",
		Status:      "",
	}
	if req.Page != 1 {
		t.Errorf("Page = %v, want 1", req.Page)
	}
	// 未设置 Category 时应为 0（表示全部）
	if req.Category != 0 {
		t.Errorf("Category = %v, want 0", req.Category)
	}
}

func TestCommentListRequest_Defaults(t *testing.T) {
	req := request.CommentListRequest{
		PageRequest: request.PageRequest{Page: 1, PageSize: 10},
		Status:      "",
		ArticleID:   0,
	}
	if req.Page != 1 {
		t.Errorf("Page = %v, want 1", req.Page)
	}
}

func TestDailyQuestionListRequest_Defaults(t *testing.T) {
	req := request.DailyQuestionListRequest{
		PageRequest: request.PageRequest{Page: 1, PageSize: 10},
		Keyword:     "",
		Date:        "",
	}
	if req.Page != 1 {
		t.Errorf("Page = %v, want 1", req.Page)
	}
}