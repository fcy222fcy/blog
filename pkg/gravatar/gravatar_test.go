package gravatar

import (
	"strings"
	"testing"
)

func TestGetQQAvatarURL(t *testing.T) {
	tests := []struct {
		name     string
		qq       string
		size     int
		expected string
	}{
		{
			name:     "size 80 maps to 100 (nearest supported >= 80)",
			qq:       "123456789",
			size:     80,
			expected: "https://q1.qlogo.cn/g?b=qq&nk=123456789&s=100",
		},
		{
			name:     "empty QQ number returns empty",
			qq:       "",
			size:     80,
			expected: "",
		},
		{
			name:     "size <= 0 defaults to 100",
			qq:       "987654321",
			size:     0,
			expected: "https://q1.qlogo.cn/g?b=qq&nk=987654321&s=100",
		},
		{
			name:     "QQ number with spaces trimmed",
			qq:       "  111222333  ",
			size:     100,
			expected: "https://q1.qlogo.cn/g?b=qq&nk=111222333&s=100",
		},
		{
			name:     "size 40 stays 40 (exactly supported)",
			qq:       "222333444",
			size:     40,
			expected: "https://q1.qlogo.cn/g?b=qq&nk=222333444&s=40",
		},
		{
			name:     "size 41 maps to 100",
			qq:       "333444555",
			size:     41,
			expected: "https://q1.qlogo.cn/g?b=qq&nk=333444555&s=100",
		},
		{
			name:     "size 100 stays 100",
			qq:       "444555666",
			size:     100,
			expected: "https://q1.qlogo.cn/g?b=qq&nk=444555666&s=100",
		},
		{
			name:     "size 101 maps to 140",
			qq:       "555666777",
			size:     101,
			expected: "https://q1.qlogo.cn/g?b=qq&nk=555666777&s=140",
		},
		{
			name:     "size 250 maps to 640 (exceeds 240)",
			qq:       "666777888",
			size:     250,
			expected: "https://q1.qlogo.cn/g?b=qq&nk=666777888&s=640",
		},
		{
			name:     "size 1000 maps to 640 (max)",
			qq:       "777888999",
			size:     1000,
			expected: "https://q1.qlogo.cn/g?b=qq&nk=777888999&s=640",
		},
		{
			name:     "negative size defaults to 100",
			qq:       "888999000",
			size:     -5,
			expected: "https://q1.qlogo.cn/g?b=qq&nk=888999000&s=100",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := GetQQAvatarURL(tt.qq, tt.size)
			if result != tt.expected {
				t.Errorf("GetQQAvatarURL(%q, %d) = %q, want %q", tt.qq, tt.size, result, tt.expected)
			}
		})
	}
}

func TestIsQQEmail(t *testing.T) {
	tests := []struct {
		name       string
		email      string
		wantQQ     string
		wantIsQQ   bool
	}{
		{
			name:     "normal QQ email",
			email:    "123456789@qq.com",
			wantQQ:   "123456789",
			wantIsQQ: true,
		},
		{
			name:     "VIP QQ email",
			email:    "987654321@vip.qq.com",
			wantQQ:   "987654321",
			wantIsQQ: true,
		},
		{
			name:     "QQ email with uppercase",
			email:    "111222333@QQ.COM",
			wantQQ:   "111222333",
			wantIsQQ: true,
		},
		{
			name:     "VIP QQ email with uppercase",
			email:    "444555666@VIP.QQ.COM",
			wantQQ:   "444555666",
			wantIsQQ: true,
		},
		{
			name:     "non-QQ email (gmail)",
			email:    "user@gmail.com",
			wantQQ:   "",
			wantIsQQ: false,
		},
		{
			name:     "non-QQ email (163)",
			email:    "test@163.com",
			wantQQ:   "",
			wantIsQQ: false,
		},
		{
			name:     "fake QQ email with letters",
			email:    "abc123@qq.com",
			wantQQ:   "",
			wantIsQQ: false,
		},
		{
			name:     "empty email",
			email:    "",
			wantQQ:   "",
			wantIsQQ: false,
		},
		{
			name:     "email with leading/trailing spaces",
			email:    "  555666777@qq.com  ",
			wantQQ:   "555666777",
			wantIsQQ: true,
		},
		{
			name:     "qqq.com (not qq.com)",
			email:    "123456@qqq.com",
			wantQQ:   "",
			wantIsQQ: false,
		},
		{
			name:     "subdomain other than vip",
			email:    "123456@mail.qq.com",
			wantQQ:   "",
			wantIsQQ: false,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			gotQQ, gotIsQQ := isQQEmail(tt.email)
			if gotQQ != tt.wantQQ || gotIsQQ != tt.wantIsQQ {
				t.Errorf("isQQEmail(%q) = (%q, %v), want (%q, %v)",
					tt.email, gotQQ, gotIsQQ, tt.wantQQ, tt.wantIsQQ)
			}
		})
	}
}

func TestGetAvatarURLByEmail(t *testing.T) {
	tests := []struct {
		name       string
		email      string
		size       int
		checkQQ    bool
		checkGrav  bool
		checkEmpty bool
	}{
		{
			name:    "QQ email returns QQ avatar URL",
			email:   "123456789@qq.com",
			size:    80,
			checkQQ: true,
		},
		{
			name:    "VIP QQ email returns QQ avatar URL",
			email:   "987654321@vip.qq.com",
			size:    100,
			checkQQ: true,
		},
		{
			name:     "non-QQ email returns Gravatar URL",
			email:    "user@gmail.com",
			size:     80,
			checkGrav: true,
		},
		{
			name:       "empty email returns empty",
			email:      "",
			size:       80,
			checkEmpty: true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := GetAvatarURLByEmail(tt.email, tt.size)
			if tt.checkQQ {
				if !strings.Contains(result, "qlogo.cn") {
					t.Errorf("Expected QQ avatar URL (qlogo.cn), got: %s", result)
				}
				if !strings.Contains(result, "b=qq") {
					t.Errorf("Expected QQ avatar URL to contain b=qq, got: %s", result)
				}
			}
			if tt.checkGrav {
				if !strings.Contains(result, "gravatar.com") {
					t.Errorf("Expected Gravatar URL, got: %s", result)
				}
			}
			if tt.checkEmpty {
				if result != "" {
					t.Errorf("Expected empty string, got: %s", result)
				}
			}
		})
	}
}

func TestGetAvatarURL(t *testing.T) {
	url := GetAvatarURL("test@example.com", nil)
	if !strings.Contains(url, "gravatar.com/avatar/") {
		t.Errorf("Expected Gravatar URL, got: %s", url)
	}
	if !strings.Contains(url, "s=80") {
		t.Errorf("Expected default size 80, got: %s", url)
	}
	if !strings.Contains(url, "d=identicon") {
		t.Errorf("Expected default identicon, got: %s", url)
	}

	customOpts := &Options{
		Size:    200,
		Default: DefaultRetro,
		Rating:  RatingPG,
	}
	url2 := GetAvatarURL("test@example.com", customOpts)
	if !strings.Contains(url2, "s=200") {
		t.Errorf("Expected size 200, got: %s", url2)
	}
	if !strings.Contains(url2, "d=retro") {
		t.Errorf("Expected d=retro, got: %s", url2)
	}
	if !strings.Contains(url2, "r=pg") {
		t.Errorf("Expected r=pg, got: %s", url2)
	}
}

func TestGetEmailHash(t *testing.T) {
	hash := GetEmailHash("MyEmailAddress@example.com")
	expected := "0bc83cb571cd1c50ba6f3e8a78ef1346"
	if hash != expected {
		t.Errorf("GetEmailHash() = %s, want %s", hash, expected)
	}

	hash2 := GetEmailHash("  MyEmailAddress@EXAMPLE.COM  ")
	if hash2 != expected {
		t.Errorf("GetEmailHash() with spaces/uppercase = %s, want %s", hash2, expected)
	}
}
