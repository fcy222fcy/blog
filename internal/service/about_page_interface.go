package service

import "blog/internal/model/dto/request"

// AboutPageService 关于页面服务接口
type AboutPageService interface {
	// GetAboutPage 获取关于页面
	GetAboutPage() (*request.AboutPageResponse, error)

	// UpdateAboutPage 更新关于页面
	UpdateAboutPage(req *request.UpdateAboutPageRequest) error
}
