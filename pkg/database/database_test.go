package database

import (
	"os"
	"strings"
	"testing"
)

func TestSplitSQLStatementsPreservesQuotedStrings(t *testing.T) {
	f, err := os.CreateTemp(t.TempDir(), "seed-*.sql")
	if err != nil {
		t.Fatal(err)
	}
	defer f.Close()

	const sql = "INSERT INTO categories (name, description) VALUES ('搭建网站', '网站搭建、博客部署相关');\nINSERT INTO tags (name) VALUES ('Go; Gin');"
	if _, err := f.WriteString(sql); err != nil {
		t.Fatal(err)
	}
	if _, err := f.Seek(0, 0); err != nil {
		t.Fatal(err)
	}

	var statements []string
	for _, statement := range splitSQLStatements(f) {
		if strings.TrimSpace(statement) != "" {
			statements = append(statements, statement)
		}
	}
	if len(statements) != 2 {
		t.Fatalf("statement count = %d, want 2", len(statements))
	}
	if !strings.Contains(statements[0], "'搭建网站'") {
		t.Fatalf("first statement lost quoted string: %q", statements[0])
	}
	if !strings.Contains(statements[1], "'Go; Gin'") {
		t.Fatalf("second statement split quoted semicolon: %q", statements[1])
	}
}
