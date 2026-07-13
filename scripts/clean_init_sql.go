// Command clean_init_sql normalizes text literals in the article seed rows.
package main

import (
	"fmt"
	"os"
	"regexp"
	"strings"
)

var articlePrefix = regexp.MustCompile(`^\('([^']*)', '([^']*)',\r?\n`)

func main() {
	if len(os.Args) != 3 {
		fmt.Fprintln(os.Stderr, "usage: go run scripts/clean_init_sql.go <input.sql> <output.sql>")
		os.Exit(2)
	}

	source, err := os.ReadFile(os.Args[1])
	if err != nil {
		panic(err)
	}

	content := string(source)
	start := strings.Index(content, "INSERT INTO articles")
	endMarker := "-- 文章标签关联"
	endOffset := strings.Index(content[start:], endMarker)
	if start < 0 || endOffset < 0 {
		panic("articles seed section not found")
	}
	end := start + endOffset
	section := content[start:end]
	section = strings.Replace(section,
		"(title, slug, content, summary, category_id, view_count, like_count, comment_count, status, is_top, reading_time, created_at, updated_at)",
		"(title, slug, content, summary, category_id, view_count, comment_count, status, is_top, reading_time, created_at, updated_at)",
		1,
	)

	rowCount := 0
	cleaned := regexp.MustCompile(`(?ms)^\('.*?NOW\(\), NOW\(\)\)[,;]$`).ReplaceAllStringFunc(section, func(row string) string {
		rowCount++
		return cleanArticleRow(row)
	})
	if rowCount == 0 {
		panic("no article seed rows cleaned")
	}

	result := removeDuplicateArticleTags(moveUsersBeforeComments(content[:start] + cleaned + content[end:]))
	if err := os.WriteFile(os.Args[2], []byte(result), 0644); err != nil {
		panic(err)
	}
	fmt.Printf("cleaned %d article rows\n", rowCount)
}

func escape(value string) string {
	return strings.ReplaceAll(value, "'", "''")
}

func removeObsoleteLikeCount(tail string) string {
	parts := strings.SplitN(tail, ", ", 4)
	if len(parts) != 4 {
		panic("unable to parse article counters")
	}
	return parts[0] + ", " + parts[1] + ", " + parts[3]
}

func moveUsersBeforeComments(sql string) string {
	userStart := strings.Index(sql, "\n-- 用户（密码")
	commentStart := strings.Index(sql, "\n-- 评论数据")
	if userStart < 0 || commentStart < 0 || userStart < commentStart {
		panic("unable to reorder user seed data")
	}
	users := sql[userStart:]
	withoutUsers := sql[:userStart]
	return withoutUsers[:commentStart] + users + "\n" + withoutUsers[commentStart:]
}

func removeDuplicateArticleTags(sql string) string {
	start := strings.Index(sql, "\n-- 文章标签关联\nINSERT INTO article_tags")
	if start < 0 {
		panic("first article_tags seed section not found")
	}
	endOffset := strings.Index(sql[start:], "\n-- 每日一问")
	if endOffset < 0 {
		panic("daily question seed section not found")
	}
	return sql[:start] + sql[start+endOffset:]
}

func cleanArticleRow(row string) string {
	prefix := articlePrefix.FindStringSubmatchIndex(row)
	if prefix == nil {
		panic("unable to parse article title and slug")
	}

	values := row[prefix[1]:]
	separator := "',\r\n"
	if !strings.Contains(values, separator) {
		separator = "',\n"
	}
	firstEnd := strings.Index(values, separator)
	if firstEnd < 1 {
		panic("unable to locate article content boundary")
	}
	secondStart := firstEnd + len(separator)
	secondEnd := strings.Index(values[secondStart:], separator)
	if secondEnd < 1 {
		panic("unable to locate article summary boundary")
	}
	secondEnd += secondStart

	content := values[1:firstEnd]
	summary := values[secondStart+1 : secondEnd]
	tail := values[secondEnd+len(separator):]
	tail = removeObsoleteLikeCount(tail)
	return fmt.Sprintf("('%s', '%s',\n'%s',\n'%s',\n%s",
		escape(row[prefix[2]:prefix[3]]),
		escape(row[prefix[4]:prefix[5]]),
		escape(content),
		escape(summary),
		tail,
	)
}
