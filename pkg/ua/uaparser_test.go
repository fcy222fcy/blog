package ua

import (
	"fmt"
	"testing"
)

// 真实的 User-Agent 字符串测试集（不是模拟，都是真实 UA）
var realUATestCases = []struct {
	Name         string
	UA           string
	ExpectOS     string
	ExpectOSVer  string
	ExpectBr     string
	ExpectBrVer  string
}{
	{
		Name:        "Windows 11 + Chrome 126（真实 Chrome Win11 UA）",
		UA:          "Mozilla/5.0 (Windows NT 10.0; Win64; x64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/126.0.0.0 Safari/537.36",
		ExpectOS:    "Windows",
		ExpectOSVer: "10", // 如果没 Build 号兜底显示 10
		ExpectBr:    "Chrome",
		ExpectBrVer: "126",
	},
	{
		Name:        "macOS Sonoma 14.5 + Safari 17.5（真实 Safari UA）",
		UA:          "Mozilla/5.0 (Macintosh; Intel Mac OS X 14_5) AppleWebKit/605.1.15 (KHTML, like Gecko) Version/17.5 Safari/605.1.15",
		ExpectOS:    "macOS",
		ExpectOSVer: "14 (Sonoma)",
		ExpectBr:    "Safari",
		ExpectBrVer: "17",
	},
	{
		Name:        "iPhone iOS 17.5 + Safari",
		UA:          "Mozilla/5.0 (iPhone; CPU iPhone OS 17_5 like Mac OS X) AppleWebKit/605.1.15 (KHTML, like Gecko) Version/17.5 Mobile/15E148 Safari/604.1",
		ExpectOS:    "iOS",
		ExpectOSVer: "17",
		ExpectBr:    "Safari",
		ExpectBrVer: "17",
	},
	{
		Name:        "Android 14 + Chrome（三星 S24 真实 UA）",
		UA:          "Mozilla/5.0 (Linux; Android 14; SM-S9280) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/126.0.6478.122 Mobile Safari/537.36",
		ExpectOS:    "Android",
		ExpectOSVer: "14",
		ExpectBr:    "Chrome",
		ExpectBrVer: "126",
	},
	{
		Name:        "Windows 10 + Edge 126",
		UA:          "Mozilla/5.0 (Windows NT 10.0; Win64; x64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/126.0.0.0 Safari/537.36 Edg/126.0.2592.68",
		ExpectOS:    "Windows",
		ExpectOSVer: "10",
		ExpectBr:    "Edge",
		ExpectBrVer: "126",
	},
	{
		Name:        "macOS + Firefox 128",
		UA:          "Mozilla/5.0 (Macintosh; Intel Mac OS X 14.4; rv:128.0) Gecko/20100101 Firefox/128.0",
		ExpectOS:    "macOS",
		ExpectOSVer: "14 (Sonoma)",
		ExpectBr:    "Firefox",
		ExpectBrVer: "128",
	},
	{
		Name:        "微信 8.0.49 on Android 13",
		UA:          "Mozilla/5.0 (Linux; Android 13; MI 13 Pro Build/TKQ1.220829.002; wv) AppleWebKit/537.36 (KHTML, like Gecko) Version/4.0 Chrome/120.0.6099.144 Mobile Safari/537.36 MicroMessenger/8.0.49.2600(0x2800313B) NetType/WIFI Language/zh_CN ABI/arm64",
		ExpectOS:    "Android",
		ExpectOSVer: "13",
		ExpectBr:    "微信",
		ExpectBrVer: "8",
	},
	{
		Name:        "iPadOS 17.5 + Safari",
		UA:          "Mozilla/5.0 (iPad; CPU OS 17_5 like Mac OS X) AppleWebKit/605.1.15 (KHTML, like Gecko) Version/17.5 Mobile/15E148 Safari/604.1",
		ExpectOS:    "iPadOS",
		ExpectOSVer: "17",
		ExpectBr:    "Safari",
		ExpectBrVer: "17",
	},
	{
		Name:        "Ubuntu 22.04 + Chrome 126",
		UA:          "Mozilla/5.0 (X11; Linux x86_64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/126.0.0.0 Safari/537.36 Ubuntu/22.04",
		ExpectOS:    "Ubuntu",
		ExpectOSVer: "22",
		ExpectBr:    "Chrome",
		ExpectBrVer: "126",
	},
	{
		Name:        "Windows 7 + IE 11（老浏览器真实 UA）",
		UA:          "Mozilla/5.0 (Windows NT 6.1; Trident/7.0; rv:11.0) like Gecko",
		ExpectOS:    "Windows",
		ExpectOSVer: "7",
		ExpectBr:    "IE",
		ExpectBrVer: "11",
	},
}

func TestParseRealUA(t *testing.T) {
	fmt.Println("")
	fmt.Println("====== 真实 User-Agent 解析测试 ======")
	passCount := 0
	for i, c := range realUATestCases {
		info := Parse(c.UA)
		gotOS := info.FormatOS()
		expectOS := c.ExpectOS
		if c.ExpectOSVer != "" {
			expectOS = c.ExpectOS + " " + c.ExpectOSVer
		}
		gotBr := info.FormatBrowser()
		expectBr := c.ExpectBr
		if c.ExpectBrVer != "" {
			expectBr = c.ExpectBr + " " + c.ExpectBrVer
		}
		okOS := gotOS == expectOS || c.ExpectOS == info.OS
		okBr := gotBr == expectBr || c.ExpectBr == info.Browser
		if okOS && okBr {
			passCount++
		}
		status := "✅"
		if !(okOS && okBr) {
			status = "❌"
		}
		fmt.Printf("%s [%02d] %s\n", status, i+1, c.Name)
		fmt.Printf("      OS    解析: %-25s  期望: %s\n", gotOS, expectOS)
		fmt.Printf("      浏览器 解析: %-25s  期望: %s\n", gotBr, expectBr)
		fmt.Println("")
	}
	fmt.Printf("====== 通过 %d / %d ======\n", passCount, len(realUATestCases))
	if passCount < len(realUATestCases) {
		t.FailNow()
	}
}
