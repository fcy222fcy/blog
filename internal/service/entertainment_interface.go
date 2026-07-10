package service

import (
	"blog/internal/model/dto/request"
	"blog/internal/model/dto/response"
)

type EntertainmentService interface {
	GetPublicList(typeStr string, year *int) (map[string]interface{}, error)
	GetAdminList(req *request.EntertainmentListRequest) (*response.PageResponse, error)
	Create(req *request.CreateEntertainmentRequest) (uint, error)
	Update(id uint, req *request.UpdateEntertainmentRequest) error
	Delete(id uint) error
	GetByID(id uint) (*response.EntertainmentResponse, error)
}
