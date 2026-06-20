package response

// SettingsResponse 系统设置响应
type SettingsResponse struct {
	SiteName        string `json:"site_name"`
	SiteDescription string `json:"site_description"`
	SiteURL         string `json:"site_url"`
	SiteKeywords    string `json:"site_keywords"`
	SEOTitle        string `json:"seo_title"`
	SEODescription  string `json:"seo_description"`
	SEOKeywords     string `json:"seo_keywords"`
	PageSize        int    `json:"page_size"`
	Favicon         string `json:"favicon"`
	Logo            string `json:"logo"`
}
