package main

import (
	"fmt"
	"os"
	"regexp"
	"text/tabwriter"

	"blog/internal/model/entity"
	"blog/pkg/config"
	"blog/pkg/database"

	"gorm.io/gorm"
)

const articleTagJoinTable = "article_tags"

// 分类/标签的"测试数据"启发式判断
var testPatterns = []*regexp.Regexp{
	regexp.MustCompile(`(?i)test`),
	regexp.MustCompile(`测试`),
	regexp.MustCompile(`临时`),
	regexp.MustCompile(`(?i)temp`),
	regexp.MustCompile(`(?i)demo`),
	regexp.MustCompile(`示例`),
	regexp.MustCompile(`(?i)example`),
	regexp.MustCompile(`未命名|未分类`),
	regexp.MustCompile(`(?i)untitled`),
	regexp.MustCompile(`(?i)^\s*[a-z0-9]{1,3}\s*$`), // 过短纯字母/数字如 "go" / "ai" 单个字（正常标签一般中文或英文组合 3+ chars）
}

func isProbablyTest(name, slug string) bool {
	s := name + "|" + slug
	for _, p := range testPatterns {
		if p.MatchString(s) {
			return true
		}
	}
	return false
}

func printTable(title string, headers []string, rows [][]string) {
	fmt.Printf("\n====== %s ======\n", title)
	if len(rows) == 0 {
		fmt.Println("  (空)")
		return
	}
	tw := tabwriter.NewWriter(os.Stdout, 0, 0, 2, ' ', 0)
	fmt.Fprintln(tw, joinTab(headers))
	for _, r := range rows {
		fmt.Fprintln(tw, joinTab(r))
	}
	tw.Flush()
}

func joinTab(ss []string) string {
	out := ""
	for i, s := range ss {
		if i > 0 {
			out += "\t"
		}
		out += s
	}
	return out
}

func main() {
	cfg, err := config.Load()
	if err != nil {
		fmt.Fprintf(os.Stderr, "[ERROR] 读取配置失败: %v\n", err)
		os.Exit(1)
	}
	db, err := database.NewDatabase(cfg.MySQL)
	if err != nil {
		fmt.Fprintf(os.Stderr, "[ERROR] 连接数据库失败: %v\n", err)
		os.Exit(1)
	}
	defer func() { _ = db.Close() }()
	g := db.DB

	// ===== 1. 先 dump 原始分类 =====
	var categories []entity.Category
	if err := g.Order("id ASC").Find(&categories).Error; err != nil {
		fmt.Fprintf(os.Stderr, "[ERROR] 查分类失败: %v\n", err)
		os.Exit(1)
	}
	catRows := make([][]string, 0, len(categories))
	catToDel := map[uint]bool{}
	for _, c := range categories {
		mark := ""
		if isProbablyTest(c.Name, c.Slug) {
			mark = " [DEL?]"
			catToDel[c.ID] = true
		}
		catRows = append(catRows, []string{
			fmt.Sprintf("%d", c.ID),
			c.Slug,
			c.Name,
			c.CreatedAt.Format("2006-01-02 15:04") + mark,
		})
	}
	printTable("当前所有分类（标记 [DEL?] 的是疑似测试数据）", []string{"ID", "Slug", "Name", "CreatedAt"}, catRows)

	// ===== 2. dump 原始标签 =====
	var tags []entity.Tag
	if err := g.Order("id ASC").Find(&tags).Error; err != nil {
		fmt.Fprintf(os.Stderr, "[ERROR] 查标签失败: %v\n", err)
		os.Exit(1)
	}
	tagRows := make([][]string, 0, len(tags))
	tagToDel := map[uint]bool{}
	for _, t := range tags {
		mark := ""
		if isProbablyTest(t.Name, t.Slug) {
			mark = " [DEL?]"
			tagToDel[t.ID] = true
		}
		tagRows = append(tagRows, []string{
			fmt.Sprintf("%d", t.ID),
			t.Slug,
			t.Name,
			t.CreatedAt.Format("2006-01-02 15:04") + mark,
		})
	}
	printTable("当前所有标签（标记 [DEL?] 的是疑似测试数据）", []string{"ID", "Slug", "Name", "CreatedAt"}, tagRows)

	// ===== 3. 事务：删除标记的测试项 =====
	delCatIDs := make([]uint, 0, len(catToDel))
	for id := range catToDel {
		delCatIDs = append(delCatIDs, id)
	}
	delTagIDs := make([]uint, 0, len(tagToDel))
	for id := range tagToDel {
		delTagIDs = append(delTagIDs, id)
	}

	fmt.Printf("\n------ 准备执行删除 ------\n")
	fmt.Printf("  拟删除分类 %d 个（ID: %v）\n", len(delCatIDs), delCatIDs)
	fmt.Printf("  拟删除标签 %d 个（ID: %v）\n", len(delTagIDs), delTagIDs)
	if len(delCatIDs) == 0 && len(delTagIDs) == 0 {
		fmt.Println("✅ 没有需要删除的测试数据，直接退出。")
		return
	}

	err = g.Transaction(func(tx *gorm.DB) error {
		// 3.1 清理文章分类外键（被删分类的文章 → category_id=0）
		if len(delCatIDs) > 0 {
			res := tx.Model(&entity.Article{}).
				Where("category_id IN ?", delCatIDs).
				Update("category_id", 0)
			if res.Error != nil {
				return fmt.Errorf("重置文章 category_id 失败: %w", res.Error)
			}
			fmt.Printf("  · 已把 %d 篇文章的 category_id 重置为 0（原分类被删除）\n", res.RowsAffected)
		}

		// 3.2 清理 article_tag 关联（被删标签）
		if len(delTagIDs) > 0 {
			res := tx.Table(articleTagJoinTable).Where("tag_id IN ?", delTagIDs).Delete(nil)
			if res.Error != nil {
				return fmt.Errorf("清理 %s 关联失败: %w", articleTagJoinTable, res.Error)
			}
			fmt.Printf("  · 已删除 %d 条 %s 关联（对应被删标签）\n", res.RowsAffected, articleTagJoinTable)
		}

		// 3.3 删分类
		if len(delCatIDs) > 0 {
			res := tx.Where("id IN ?", delCatIDs).Delete(&entity.Category{})
			if res.Error != nil {
				return fmt.Errorf("删除分类失败: %w", res.Error)
			}
			fmt.Printf("  · 已删除 %d 条 categories\n", res.RowsAffected)
		}

		// 3.4 删标签
		if len(delTagIDs) > 0 {
			res := tx.Where("id IN ?", delTagIDs).Delete(&entity.Tag{})
			if res.Error != nil {
				return fmt.Errorf("删除标签失败: %w", res.Error)
			}
			fmt.Printf("  · 已删除 %d 条 tags\n", res.RowsAffected)
		}
		return nil
	})
	if err != nil {
		fmt.Fprintf(os.Stderr, "[ERROR] 事务失败，已回滚: %v\n", err)
		os.Exit(1)
	}

	// ===== 4. 删除后的校验 dump =====
	fmt.Println("\n====== 删除后：分类与标签 最终校验 ======")
	if err := g.Order("id ASC").Find(&categories).Error; err != nil {
		fmt.Fprintf(os.Stderr, "[ERROR] 重查分类失败: %v\n", err)
		os.Exit(1)
	}
	catRows = catRows[:0]
	for _, c := range categories {
		catRows = append(catRows, []string{
			fmt.Sprintf("%d", c.ID), c.Slug, c.Name, c.CreatedAt.Format("2006-01-02 15:04"),
		})
	}
	printTable("删除后 - categories", []string{"ID", "Slug", "Name", "CreatedAt"}, catRows)

	if err := g.Order("id ASC").Find(&tags).Error; err != nil {
		fmt.Fprintf(os.Stderr, "[ERROR] 重查标签失败: %v\n", err)
		os.Exit(1)
	}
	tagRows = tagRows[:0]
	for _, t := range tags {
		tagRows = append(tagRows, []string{
			fmt.Sprintf("%d", t.ID), t.Slug, t.Name, t.CreatedAt.Format("2006-01-02 15:04"),
		})
	}
	printTable("删除后 - tags", []string{"ID", "Slug", "Name", "CreatedAt"}, tagRows)

	// 统计文章关联一致性
	type stat struct {
		TotalArticles    int64
		ArticlesWithCat  int64
		ArticlesWithTag  int64
		ArticleTagLinks  int64
		OrphanTagLinks   int64 // 指向不存在 tag_id 的 article_tag
		OrphanCatLinks   int64 // 指向不存在 category_id 的 articles
	}
	var s stat
	g.Model(&entity.Article{}).Count(&s.TotalArticles)
	g.Model(&entity.Article{}).Where("category_id > 0").Count(&s.ArticlesWithCat)
	g.Table(articleTagJoinTable).Count(&s.ArticleTagLinks)
	g.Raw(fmt.Sprintf(`SELECT COUNT(DISTINCT article_id) FROM %s`, articleTagJoinTable)).Scan(&s.ArticlesWithTag)
	g.Raw(fmt.Sprintf(`SELECT COUNT(*) FROM %s at LEFT JOIN tags t ON t.id = at.tag_id WHERE t.id IS NULL`, articleTagJoinTable)).Scan(&s.OrphanTagLinks)
	g.Raw(`SELECT COUNT(*) FROM articles a LEFT JOIN categories c ON c.id = a.category_id WHERE a.category_id > 0 AND c.id IS NULL`).Scan(&s.OrphanCatLinks)

	fmt.Printf("\n------ 关联一致性统计 ------\n")
	fmt.Printf("  文章总数            : %d\n", s.TotalArticles)
	fmt.Printf("  有分类的文章        : %d\n", s.ArticlesWithCat)
	fmt.Printf("  至少有 1 个标签的文章: %d\n", s.ArticlesWithTag)
	fmt.Printf("  %s 总记录  : %d\n", articleTagJoinTable, s.ArticleTagLinks)
	fmt.Printf("  孤儿 %s    : %d (应为 0)\n", articleTagJoinTable, s.OrphanTagLinks)
	fmt.Printf("  孤儿 category_id    : %d (应为 0)\n", s.OrphanCatLinks)

	fmt.Println("\n✅ 清理完成！建议：重启后端，硬刷新前端查看详情页。")
}
