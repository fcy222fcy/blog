package response

// LoginResponse 登录响应
type LoginResponse struct {
	Token     string `json:"token"`
	ExpiresAt int64  `json:"expires_at"`
}

// UserProfileResponse 用户信息响应
type UserProfileResponse struct {
	ID          uint          `json:"id"`
	Username    string        `json:"username"`
	Nickname    string        `json:"nickname"`
	Email       string        `json:"email"`
	Avatar      string        `json:"avatar"`
	Bio         string        `json:"bio"`
	SocialLinks []SocialLink  `json:"social_links"`
}

// SocialLink 社交链接
type SocialLink struct {
	Name string `json:"name"`
	URL  string `json:"url"`
}
