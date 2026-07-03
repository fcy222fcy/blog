package repository

import (
	"blog/internal/model/entity"

	"gorm.io/gorm"
)

// articleRepository 文章数据访问实现
type articleRepository struct {
	db *gorm.DB
}

// NewArticleRepository 创建文章数据访问
func NewArticleRepository(db *gorm.DB) ArticleRepository {
	return &articleRepository{db: db}
}

// FindByID 根据 ID 查找文章
func (r *articleRepository) FindByID(id uint) (*entity.Article, error) {
	var article entity.Article
	err := r.db.Preload("Category").Preload("Tags").First(&article, id).Error
	if err != nil {
		if err == gorm.ErrRecordNotFound {
			return nil, nil
		}
		return nil, err
	}
	return &article, nil
}

// FindBySlug 根据 slug 查找文章
func (r *articleRepository) FindBySlug(slug string) (*entity.Article, error) {
	var article entity.Article
	err := r.db.Preload("Category").Preload("Tags").Where("slug = ?", slug).First(&article).Error
	if err != nil {
		if err == gorm.ErrRecordNotFound {
			return nil, nil
		}
		return nil, err
	}
	return &article, nil
}

// Create 创建文章
func (r *articleRepository) Create(article *entity.Article) error {
	return r.db.Create(article).Error
}

// Update 更新文章
func (r *articleRepository) Update(article *entity.Article) error {
	return r.db.Save(article).Error
}

// Delete 删除文章（软删除）
func (r *articleRepository) Delete(id uint) error {
	return r.db.Delete(&entity.Article{}, id).Error
}

// ListPublished 已发布文章列表（前台）
func (r *articleRepository) ListPublished(offset, limit int, categoryId, tagId uint, keyword string) ([]*entity.Article, int64, error) {
	var articles []*entity.Article
	var total int64

	query := r.db.Model(&entity.Article{}).Where("status = ?", entity.ArticleStatusPublished)
	if keyword != "" {
		query = query.Where("title LIKE ?", "%"+keyword+"%")
	}
	if categoryId > 0 {
		query = query.Where("category_id = ?", categoryId)
	}
	if tagId > 0 {
		query = query.Joins("JOIN article_tags ON article_tags.article_id = articles.id").
			Where("article_tags.tag_id = ?", tagId)
	}

	err := query.Count(&total).Error
	if err != nil {
		return nil, 0, err
	}

	err = query.Preload("Category").Preload("Tags").
		Offset(offset).Limit(limit).
		Order("is_top DESC, created_at DESC").
		Find(&articles).Error
	if err != nil {
		return nil, 0, err
	}

	return articles, total, nil
}

// ListAll 所有文章列表（后台）
func (r *articleRepository) ListAll(offset, limit int, status, keyword string, categoryID uint) ([]*entity.Article, int64, error) {
	var articles []*entity.Article
	var total int64

	query := r.db.Model(&entity.Article{})
	if keyword != "" {
		query = query.Where("title LIKE ?", "%"+keyword+"%")
	}
	if status != "" {
		query = query.Where("status = ?", status)
	}
	if categoryID > 0 {
		query = query.Where("category_id = ?", categoryID)
	}

	err := query.Count(&total).Error
	if err != nil {
		return nil, 0, err
	}

	err = query.Preload("Category").Preload("Tags").
		Offset(offset).Limit(limit).
		Order("created_at DESC").
		Find(&articles).Error
	if err != nil {
		return nil, 0, err
	}

	return articles, total, nil
}

// IncrementViewCount 增加浏览量
func (r *articleRepository) IncrementViewCount(id uint) error {
	return r.db.Model(&entity.Article{}).Where("id = ?", id).
		UpdateColumn("view_count", gorm.Expr("view_count + 1")).Error
}

// BatchDelete 批量删除
func (r *articleRepository) BatchDelete(ids []uint) error {
	return r.db.Delete(&entity.Article{}, ids).Error
}

// Count 统计文章数量
func (r *articleRepository) Count(status string) (int64, error) {
	var total int64
	query := r.db.Model(&entity.Article{})
	if status != "" {
		query = query.Where("status = ?", status)
	}
	err := query.Count(&total).Error
	return total, err
}

// SumViewCount 统计总浏览量
func (r *articleRepository) SumViewCount() (int64, error) {
	var sum int64
	err := r.db.Model(&entity.Article{}).Select("COALESCE(SUM(view_count), 0)").Scan(&sum).Error
	return sum, err
}

// FindByCategoryID 根据分类ID查找文章
func (r *articleRepository) FindByCategoryID(categoryID uint, offset, limit int) ([]*entity.Article, int64, error) {
	var articles []*entity.Article
	var total int64

	query := r.db.Model(&entity.Article{}).Where("category_id = ? AND status = ?", categoryID, entity.ArticleStatusPublished)
	err := query.Count(&total).Error
	if err != nil {
		return nil, 0, err
	}

	err = query.Preload("Category").Preload("Tags").
		Offset(offset).Limit(limit).
		Order("created_at DESC").
		Find(&articles).Error
	return articles, total, err
}

// FindByTagID 根据标签ID查找文章
func (r *articleRepository) FindByTagID(tagID uint, offset, limit int) ([]*entity.Article, int64, error) {
	var articles []*entity.Article
	var total int64

	query := r.db.Model(&entity.Article{}).
		Joins("JOIN article_tags ON article_tags.article_id = articles.id").
		Where("article_tags.tag_id = ? AND articles.status = ?", tagID, entity.ArticleStatusPublished)
	err := query.Count(&total).Error
	if err != nil {
		return nil, 0, err
	}

	err = query.Preload("Category").Preload("Tags").
		Offset(offset).Limit(limit).
		Order("articles.created_at DESC").
		Find(&articles).Error
	return articles, total, err
}

// GetArchives 获取文章归档
func (r *articleRepository) GetArchives() ([]map[string][]*entity.Article, error) {
	var articles []*entity.Article
	err := r.db.Where("status = ?", entity.ArticleStatusPublished).
		Order("created_at DESC").
		Find(&articles).Error
	if err != nil {
		return nil, err
	}

	// 按年份分组
	yearMap := make(map[string][]*entity.Article)
	for _, article := range articles {
		year := article.CreatedAt.Format("2006")
		yearMap[year] = append(yearMap[year], article)
	}

	var result []map[string][]*entity.Article
	for year, arts := range yearMap {
		result = append(result, map[string][]*entity.Article{year: arts})
	}
	return result, nil
}

// GetRecent 获取最近文章
func (r *articleRepository) GetRecent(limit int) ([]entity.Article, error) {
	var articles []entity.Article
	err := r.db.
		Where("status = ?", entity.ArticleStatusPublished).
		Order("created_at DESC").
		Limit(limit).
		Preload("Category").
		Preload("Tags").
		Find(&articles).Error
	return articles, err
}

// Search 搜索文章（标题、内容、摘要模糊搜索）
func (r *articleRepository) Search(keyword string, offset, limit int) ([]*entity.Article, int64, error) {
	var articles []*entity.Article
	var total int64

	likePattern := "%" + keyword + "%"
	query := r.db.Model(&entity.Article{}).
		Where("status = ?", "published").
		Where("title LIKE ? OR content LIKE ? OR summary LIKE ?", likePattern, likePattern, likePattern)

	if err := query.Count(&total).Error; err != nil {
		return nil, 0, err
	}

	err := query.
		Preload("Category").Preload("Tags").
		Offset(offset).Limit(limit).
		Order("is_top DESC, created_at DESC").
		Find(&articles).Error
	if err != nil {
		return nil, 0, err
	}

	return articles, total, nil
}

// GetDB 获取数据库实例
func (r *articleRepository) GetDB() *gorm.DB {
	return r.db
}

// UpdateTags 更新文章标签关联
func (r *articleRepository) UpdateTags(article *entity.Article, tags []entity.Tag) error {
	return r.db.Model(article).Association("Tags").Replace(tags)
}
