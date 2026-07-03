package unit

import (
	"testing"

	"blog/internal/model/entity"
)

func TestUser_TableName(t *testing.T) {
	u := entity.User{}
	if u.TableName() != "users" {
		t.Errorf("TableName() = %v, want users", u.TableName())
	}
}

func TestArticle_TableName(t *testing.T) {
	a := entity.Article{}
	if a.TableName() != "articles" {
		t.Errorf("TableName() = %v, want articles", a.TableName())
	}
}

func TestCategory_TableName(t *testing.T) {
	c := entity.Category{}
	if c.TableName() != "categories" {
		t.Errorf("TableName() = %v, want categories", c.TableName())
	}
}

func TestTag_TableName(t *testing.T) {
	tag := entity.Tag{}
	if tag.TableName() != "tags" {
		t.Errorf("TableName() = %v, want tags", tag.TableName())
	}
}

func TestComment_TableName(t *testing.T) {
	c := entity.Comment{}
	if c.TableName() != "comments" {
		t.Errorf("TableName() = %v, want comments", c.TableName())
	}
}

func TestLink_TableName(t *testing.T) {
	l := entity.Link{}
	if l.TableName() != "links" {
		t.Errorf("TableName() = %v, want links", l.TableName())
	}
}

func TestDailyQuestion_TableName(t *testing.T) {
	dq := entity.DailyQuestion{}
	if dq.TableName() != "daily_questions" {
		t.Errorf("TableName() = %v, want daily_questions", dq.TableName())
	}
}

func TestArticle_DefaultValues(t *testing.T) {
	a := entity.Article{
		Title:   "Test Title",
		Slug:    "test-title",
		Content: "Test Content",
	}

	if a.Status != "" {
		// 默认值应为空字符串，GORM 会在数据库层面设置默认值
	}
	if a.ViewCount != 0 {
		t.Errorf("ViewCount default = %v, want 0", a.ViewCount)
	}
	if a.CommentCount != 0 {
		t.Errorf("CommentCount default = %v, want 0", a.CommentCount)
	}
	if a.IsTop {
		t.Errorf("IsTop default = true, want false")
	}
	if a.ReadingTime != 0 {
		t.Errorf("ReadingTime default = %v, want 0", a.ReadingTime)
	}
	_ = a // 使用变量
}

func TestBaseEntity_Fields(t *testing.T) {
	// 验证 BaseEntity 包含必要的字段
	b := entity.BaseEntity{}
	if b.ID != 0 {
		t.Errorf("ID default = %v, want 0", b.ID)
	}
}