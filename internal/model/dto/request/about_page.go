package request

// UpdateAboutPageRequest 更新关于页面请求
type UpdateAboutPageRequest struct {
	Title       string `json:"title"`
	Subtitle    string `json:"subtitle"`
	Bio         string `json:"bio"`
	Skills      string `json:"skills"`
	AboutMe     string `json:"about_me"`
	AboutSite   string `json:"about_site"`
	Projects    string `json:"projects"`
	ContactInfo string `json:"contact_info"`
}

// AboutPageResponse 关于页面响应
type AboutPageResponse struct {
	ID          uint   `json:"id"`
	Title       string `json:"title"`
	Subtitle    string `json:"subtitle"`
	Bio         string `json:"bio"`
	Skills      string `json:"skills"`
	AboutMe     string `json:"about_me"`
	AboutSite   string `json:"about_site"`
	Projects    string `json:"projects"`
	ContactInfo string `json:"contact_info"`
}
