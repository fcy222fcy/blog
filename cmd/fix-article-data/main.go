package main

import (
	"fmt"
	"os"
	"strings"
	"text/tabwriter"

	"blog/internal/model/entity"
	"blog/pkg/config"
	"blog/pkg/database"

	"gorm.io/gorm"
)

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

	// ===== 1. 打印所有分类 =====
	fmt.Println("\n====== 当前所有分类 ======")
	var categories []entity.Category
	if err := g.Order("id ASC").Find(&categories).Error; err != nil {
		fmt.Fprintf(os.Stderr, "[ERROR] 查分类失败: %v\n", err)
		os.Exit(1)
	}
	if len(categories) == 0 {
		fmt.Println("  (空，没有任何分类)")
	} else {
		tw := tabwriter.NewWriter(os.Stdout, 0, 0, 2, ' ', 0)
		fmt.Fprintln(tw, "ID\tSlug\tName\tCreatedAt")
		for _, c := range categories {
			fmt.Fprintf(tw, "%d\t%s\t%s\t%s\n", c.ID, c.Slug, c.Name, c.CreatedAt.Format("2006-01-02 15:04"))
		}
		tw.Flush()
	}

	// ===== 2. 打印所有文章 =====
	fmt.Println("\n====== 当前所有文章 ======")
	type articleRow struct {
		ID         uint
		Title      string
		CategoryID uint
		Cover      string
		Status     string
	}
	var rows []articleRow
	if err := g.Model(&entity.Article{}).Select("id, title, category_id, cover, status").Order("id ASC").Scan(&rows).Error; err != nil {
		fmt.Fprintf(os.Stderr, "[ERROR] 查文章失败: %v\n", err)
		os.Exit(1)
	}
	if len(rows) == 0 {
		fmt.Println("  (空，没有任何文章)")
	} else {
		tw := tabwriter.NewWriter(os.Stdout, 0, 0, 2, ' ', 0)
		fmt.Fprintln(tw, "ID\tCategoryID\tStatus\tCover?\tTitle")
		noCat := 0
		noCover := 0
		for _, r := range rows {
			hasCover := "NO"
			if strings.TrimSpace(r.Cover) != "" {
				hasCover = "YES"
			} else {
				noCover++
			}
			if r.CategoryID == 0 {
				noCat++
			}
			fmt.Fprintf(tw, "%d\t%d\t%s\t%s\t%s\n", r.ID, r.CategoryID, r.Status, hasCover, r.Title)
		}
		tw.Flush()
		fmt.Printf("\n[统计] 共 %d 篇文章，未设置分类: %d，未设置封面: %d\n", len(rows), noCat, noCover)
	}

	// ===== 3. 确定默认分类 ID =====
	var defaultCatID uint
	if len(categories) == 0 {
		fmt.Println("\n[操作] 创建默认分类「未分类」...")
		def := entity.Category{Name: "未分类", Slug: "uncategorized", Description: "系统自动创建的默认分类"}
		if err := g.Create(&def).Error; err != nil {
			fmt.Fprintf(os.Stderr, "[ERROR] 创建默认分类失败: %v\n", err)
			os.Exit(1)
		}
		defaultCatID = def.ID
		fmt.Printf("  ✓ 创建成功，ID=%d\n", defaultCatID)
	} else {
		defaultCatID = categories[0].ID
		fmt.Printf("\n[信息] 选第一个分类作为默认分类：ID=%d Name=%q\n", defaultCatID, categories[0].Name)
	}

	// ===== 4. 修复所有 category_id=0 / NULL 的文章 =====
	fmt.Println("\n[操作] 给 category_id=0 / NULL 的文章设置默认分类...")
	res := g.Model(&entity.Article{}).
		Where("category_id IS NULL OR category_id = ?", 0).
		Update("category_id", defaultCatID)
	if res.Error != nil {
		fmt.Fprintf(os.Stderr, "[ERROR] 更新 category_id 失败: %v\n", res.Error)
		os.Exit(1)
	}
	fmt.Printf("  ✓ 更新了 %d 篇文章的分类\n", res.RowsAffected)

	// ===== 5. 再打印一遍，确认分类修复生效 =====
	fmt.Println("\n====== 修复后文章分类校验 ======")
	rows = rows[:0]
	if err := g.Model(&entity.Article{}).Select("id, title, category_id").Order("id ASC").Scan(&rows).Error; err != nil {
		fmt.Fprintf(os.Stderr, "[ERROR] 二次查询失败: %v\n", err)
		os.Exit(1)
	}
	tw := tabwriter.NewWriter(os.Stdout, 0, 0, 2, ' ', 0)
	fmt.Fprintln(tw, "ID\tCategoryID\tTitle")
	for _, r := range rows {
		fmt.Fprintf(tw, "%d\t%d\t%s\n", r.ID, r.CategoryID, r.Title)
	}
	tw.Flush()

	// =======================================================
	// ===== 第二部分：给文章批量补 TAGS（加 # 前缀样式用）=====
	// =======================================================

	// ===== 6. 打印所有标签 =====
	fmt.Println("\n====== 当前所有标签 ======")
	var tags []entity.Tag
	if err := g.Order("id ASC").Find(&tags).Error; err != nil {
		fmt.Fprintf(os.Stderr, "[ERROR] 查标签失败: %v\n", err)
		os.Exit(1)
	}
	if len(tags) == 0 {
		fmt.Println("  (空，还没创建任何标签)")
	} else {
		tw = tabwriter.NewWriter(os.Stdout, 0, 0, 2, ' ', 0)
		fmt.Fprintln(tw, "ID\tSlug\tName")
		for _, t := range tags {
			fmt.Fprintf(tw, "%d\t%s\t%s\n", t.ID, t.Slug, t.Name)
		}
		tw.Flush()
	}

	// ===== 7. 打印文章标签关联现状 =====
	fmt.Println("\n====== 文章标签关联现状 ======")
	type atRow struct {
		ArticleID uint
		TagCount  int64
	}
	var atRows []atRow
	if err := g.Model(&entity.ArticleTag{}).
		Select("article_id, COUNT(tag_id) as tag_count").
		Group("article_id").Scan(&atRows).Error; err != nil {
		fmt.Fprintf(os.Stderr, "[ERROR] 查 article_tags 失败: %v\n", err)
		os.Exit(1)
	}
	cntMap := map[uint]int64{}
	for _, r := range atRows {
		cntMap[r.ArticleID] = r.TagCount
	}
	// 打所有文章 + 它的 tag 数
	tw = tabwriter.NewWriter(os.Stdout, 0, 0, 2, ' ', 0)
	fmt.Fprintln(tw, "ID\tTags\tTitle")
	noTagArticles := 0
	var allArticleIDs []uint
	for _, r := range rows {
		allArticleIDs = append(allArticleIDs, r.ID)
		c := cntMap[r.ID]
		if c == 0 {
			noTagArticles++
		}
		fmt.Fprintf(tw, "%d\t%d\t%s\n", r.ID, c, r.Title)
	}
	tw.Flush()
	fmt.Printf("\n[统计] 共 %d 篇文章，完全没有标签的: %d 篇\n", len(rows), noTagArticles)

	// ===== 8. 定义「按分类→推荐标签」映射，不存在就创建 =====
	fmt.Println("\n[操作] 准备基础标签（按分类分组）...")
	// 分类ID → 要挂的推荐标签名（按 slug 建唯一索引）
	categoryRecommend := map[uint][]string{
		1: {"博客搭建", "个人博客", "VPS", "Cloudflare", "SEO优化", "开源主题", "Gravatar"},
		2: {"Go语言", "Gin框架", "Vue3", "JavaScript", "MySQL", "Redis", "Docker",
			"Kubernetes", "Git", "软件架构", "性能优化", "AI编程助手"},
		3: {"周末随笔", "读书笔记", "旅行记录", "生活记录", "美食", "心情随笔",
			"程序员生活", "健康生活", "年度总结", "摄影"},
		// 测试分类默认挂通用标签
	}
	universalTags := []string{"技术分享", "原创", "干货"}

	// 把所有需要的标签先确保创建，拿到 name->ID 映射
	nameToID := map[string]uint{}
	// 先把已存在的标签 name -> ID 填进去
	for _, t := range tags {
		nameToID[t.Name] = t.ID
	}
	// 把所有推荐标签名 collect 出来去重
	needTagNames := map[string]struct{}{}
	for _, names := range categoryRecommend {
		for _, n := range names {
			needTagNames[n] = struct{}{}
		}
	}
	for _, n := range universalTags {
		needTagNames[n] = struct{}{}
	}
	createdTags := 0
	for name := range needTagNames {
		if _, ok := nameToID[name]; ok {
			continue
		}
		slug := strings.ToLower(strings.ReplaceAll(name, " ", "-"))
		newTag := entity.Tag{Name: name, Slug: slug}
		// FirstOrCreate 防 unique 冲突
		if err := g.Where(entity.Tag{Name: name}).FirstOrCreate(&newTag).Error; err != nil {
			fmt.Fprintf(os.Stderr, "[WARN] 创建标签 name=%q 失败: %v (跳过)\n", name, err)
			continue
		}
		nameToID[name] = newTag.ID
		createdTags++
	}
	fmt.Printf("  ✓ 新创建了 %d 个标签，当前标签库共 %d 个\n", createdTags, len(nameToID))

	// ===== 9. 给没标签的文章，按它的 category_id 智能挂 2-3 个 =====
	fmt.Println("\n[操作] 给完全没标签的文章批量绑定标签...")
	articlesForBind := []entity.Article{}
	if err := g.Select("id, category_id").Find(&articlesForBind, allArticleIDs).Error; err != nil {
		fmt.Fprintf(os.Stderr, "[ERROR] 查询文章用于绑定标签失败: %v\n", err)
		os.Exit(1)
	}
	// 为了防止重复插入，先读已有的 (article_id, tag_id) 对
	type atPair struct {
		ArticleID uint
		TagID     uint
	}
	var existingPairs []atPair
	if err := g.Model(&entity.ArticleTag{}).Find(&existingPairs).Error; err != nil {
		fmt.Fprintf(os.Stderr, "[ERROR] 查现有关联失败: %v\n", err)
		os.Exit(1)
	}
	existing := map[[2]uint]struct{}{}
	for _, p := range existingPairs {
		existing[[2]uint{p.ArticleID, p.TagID}] = struct{}{}
	}

	// 简单确定性选标签：用文章ID做伪随机，保证同一篇每次跑选的都一样
	pick := func(pool []string, id uint, n int) []string {
		if len(pool) == 0 {
			return nil
		}
		// 复制并做简单洗牌（按 id 固定偏移）
		out := make([]string, len(pool))
		copy(out, pool)
		off := int(id) % len(out)
		out = append(out[off:], out[:off]...)
		if n > len(out) {
			n = len(out)
		}
		return out[:n]
	}

	var toBind []entity.ArticleTag
	boundCount := 0
	for _, a := range articlesForBind {
		if cntMap[a.ID] > 0 {
			continue // 已经有标签的跳过
		}
		var pool []string
		if rec, ok := categoryRecommend[a.CategoryID]; ok {
			pool = append(pool, rec...)
		} else {
			pool = append(pool, universalTags...)
		}
		// 每篇文章选 3 个；如果分类池没匹配上，就拿 universal 里的 3 个补
		names := pick(pool, a.ID, 3)
		if len(names) < 3 {
			extra := pick(universalTags, a.ID+1, 3-len(names))
			names = append(names, extra...)
		}
		for _, nm := range names {
			tid, ok := nameToID[nm]
			if !ok {
				continue
			}
			key := [2]uint{a.ID, tid}
			if _, dup := existing[key]; dup {
				continue
			}
			existing[key] = struct{}{}
			toBind = append(toBind, entity.ArticleTag{ArticleID: a.ID, TagID: tid})
			boundCount++
		}
	}
	// 批量 INSERT
	if len(toBind) > 0 {
		// 分批次防止 SQL 太长
		batch := 200
		for i := 0; i < len(toBind); i += batch {
			end := i + batch
			if end > len(toBind) {
				end = len(toBind)
			}
			if err := g.Create(toBind[i:end]).Error; err != nil {
				fmt.Fprintf(os.Stderr, "[ERROR] 批量写入标签关联失败 (i=%d): %v\n", i, err)
				os.Exit(1)
			}
		}
		fmt.Printf("  ✓ 给 %d 篇文章新增了 %d 条标签关联\n", noTagArticles, boundCount)
	} else {
		fmt.Println("  (没标签需要新增，跳过)")
	}

	// ===== 10. 最终校验：打印每篇文章的标签名 =====
	fmt.Println("\n====== 最终校验：每篇文章的标签（# 号显示预览）======")
	// 把 tag_id → name 做个反查
	idToName := map[uint]string{}
	for name, id := range nameToID {
		idToName[id] = name
	}
	// 再全量查一次关联
	type atFull struct {
		ArticleID uint
		TagID     uint
	}
	var full []atFull
	if err := g.Model(&entity.ArticleTag{}).Order("article_id ASC, tag_id ASC").Find(&full).Error; err != nil {
		fmt.Fprintf(os.Stderr, "[ERROR] 查最终关联失败: %v\n", err)
		os.Exit(1)
	}
	artToTags := map[uint][]string{}
	for _, r := range full {
		artToTags[r.ArticleID] = append(artToTags[r.ArticleID], idToName[r.TagID])
	}
	tw = tabwriter.NewWriter(os.Stdout, 0, 0, 2, ' ', 0)
	fmt.Fprintln(tw, "ArticleID\t#标签（空格分隔预览）")
	for _, a := range articlesForBind {
		ts := artToTags[a.ID]
		preview := ""
		for i, t := range ts {
			if i > 0 {
				preview += "  "
			}
			preview += "#" + t
		}
		if preview == "" {
			preview = "(无)"
		}
		fmt.Fprintf(tw, "%d\t%s\n", a.ID, preview)
	}
	tw.Flush()

	// 屏蔽 unused 警告
	_ = gorm.ErrRecordNotFound

	fmt.Println("\n✅ 完成！现在重启后端 + Ctrl+Shift+R 刷新详情页，分类胶囊就显示了。")
}
