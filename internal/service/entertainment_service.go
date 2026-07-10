package service

import (
	"blog/internal/model/dto/request"
	"blog/internal/model/dto/response"
	"blog/internal/model/entity"
	"blog/internal/repository"
	bizerrors "blog/pkg/errors"
	"blog/pkg/logger"
	"fmt"
	"sort"
)

type entertainmentService struct {
	entertainmentRepo repository.EntertainmentRepository
}

func NewEntertainmentService(entertainmentRepo repository.EntertainmentRepository) EntertainmentService {
	return &entertainmentService{entertainmentRepo: entertainmentRepo}
}

func (s *entertainmentService) GetPublicList(typeStr string, year *int) (map[string]interface{}, error) {
	years, err := s.entertainmentRepo.ListYears()
	if err != nil {
		return nil, fmt.Errorf("获取年份列表失败, %w", err)
	}

	list, err := s.entertainmentRepo.ListPublic(typeStr, year)
	if err != nil {
		return nil, fmt.Errorf("获取娱乐列表失败, %w", err)
	}

	resultList := make([]response.EntertainmentResponse, 0, len(list))
	for _, item := range list {
		resultList = append(resultList, *s.toResponse(item))
	}

	groups := make(map[string][]response.EntertainmentResponse)
	for _, item := range resultList {
		key := item.Type
		if _, ok := groups[key]; !ok {
			groups[key] = make([]response.EntertainmentResponse, 0)
		}
		groups[key] = append(groups[key], item)
	}

	sort.Ints(years)
	reversedYears := make([]int, 0, len(years))
	for i := len(years) - 1; i >= 0; i-- {
		reversedYears = append(reversedYears, years[i])
	}

	return map[string]interface{}{
		"years": reversedYears,
		"list":  resultList,
		"groups": groups,
	}, nil
}

func (s *entertainmentService) GetAdminList(req *request.EntertainmentListRequest) (*response.PageResponse, error) {
	if req.Page <= 0 {
		req.Page = 1
	}
	if req.PageSize <= 0 {
		req.PageSize = 20
	}

	list, total, err := s.entertainmentRepo.List(
		(req.Page-1)*req.PageSize,
		req.PageSize,
		req.Type,
		req.Status,
		req.Year,
		req.Keyword,
	)
	if err != nil {
		return nil, fmt.Errorf("获取后台娱乐列表失败, %w", err)
	}

	var result []response.EntertainmentResponse
	for _, item := range list {
		result = append(result, *s.toResponse(item))
	}

	return response.NewPageResponse(result, total, req.Page, req.PageSize), nil
}

func (s *entertainmentService) Create(req *request.CreateEntertainmentRequest) (uint, error) {
	item := &entity.Entertainment{
		Title:          req.Title,
		TitleEn:        req.TitleEn,
		Type:           req.Type,
		Year:           req.Year,
		Cover:          req.Cover,
		Rating:         req.Rating,
		RatingExternal: req.RatingExternal,
		Platform:       req.Platform,
		Playtime:       req.Playtime,
		Comment:        req.Comment,
		Status:         req.Status,
		Link:           req.Link,
	}

	if err := s.entertainmentRepo.Create(item); err != nil {
		return 0, fmt.Errorf("创建娱乐条目失败, %w", err)
	}

	logger.Infof("创建娱乐条目成功, id=%d, title=%s", item.ID, item.Title)
	return item.ID, nil
}

func (s *entertainmentService) Update(id uint, req *request.UpdateEntertainmentRequest) error {
	item, err := s.entertainmentRepo.FindByID(id)
	if err != nil {
		return fmt.Errorf("查询娱乐条目失败, %w", err)
	}
	if item == nil {
		return bizerrors.New(bizerrors.CodeEntertainmentNotFound, bizerrors.GetMessage(bizerrors.CodeEntertainmentNotFound))
	}

	if req.Title != "" {
		item.Title = req.Title
	}
	if req.TitleEn != "" {
		item.TitleEn = req.TitleEn
	}
	if req.Type != "" {
		item.Type = req.Type
	}
	if req.Year != 0 {
		item.Year = req.Year
	}
	if req.Cover != "" {
		item.Cover = req.Cover
	}
	if req.Rating != 0 {
		item.Rating = req.Rating
	}
	if req.RatingExternal != 0 {
		item.RatingExternal = req.RatingExternal
	}
	if req.Platform != "" {
		item.Platform = req.Platform
	}
	if req.Playtime != "" {
		item.Playtime = req.Playtime
	}
	if req.Comment != "" {
		item.Comment = req.Comment
	}
	if req.Status != "" {
		item.Status = req.Status
	}
	if req.Link != "" {
		item.Link = req.Link
	}

	if err := s.entertainmentRepo.Update(item); err != nil {
		return fmt.Errorf("更新娱乐条目失败, %w", err)
	}

	logger.Infof("更新娱乐条目成功, id=%d", id)
	return nil
}

func (s *entertainmentService) Delete(id uint) error {
	item, err := s.entertainmentRepo.FindByID(id)
	if err != nil {
		return fmt.Errorf("查询娱乐条目失败, %w", err)
	}
	if item == nil {
		return bizerrors.New(bizerrors.CodeEntertainmentNotFound, bizerrors.GetMessage(bizerrors.CodeEntertainmentNotFound))
	}

	if err := s.entertainmentRepo.Delete(id); err != nil {
		return fmt.Errorf("删除娱乐条目失败, %w", err)
	}

	logger.Infof("删除娱乐条目成功, id=%d", id)
	return nil
}

func (s *entertainmentService) GetByID(id uint) (*response.EntertainmentResponse, error) {
	item, err := s.entertainmentRepo.FindByID(id)
	if err != nil {
		return nil, fmt.Errorf("查询娱乐条目失败, %w", err)
	}
	if item == nil {
		return nil, bizerrors.New(bizerrors.CodeEntertainmentNotFound, bizerrors.GetMessage(bizerrors.CodeEntertainmentNotFound))
	}
	return s.toResponse(item), nil
}

func (s *entertainmentService) toResponse(item *entity.Entertainment) *response.EntertainmentResponse {
	return &response.EntertainmentResponse{
		ID:             item.ID,
		Title:          item.Title,
		TitleEn:        item.TitleEn,
		Type:           item.Type,
		Year:           item.Year,
		Cover:          item.Cover,
		Rating:         item.Rating,
		RatingExternal: item.RatingExternal,
		Platform:       item.Platform,
		Playtime:       item.Playtime,
		Comment:        item.Comment,
		Status:         item.Status,
		Link:           item.Link,
		CreatedAt:      item.CreatedAt,
		UpdatedAt:      item.UpdatedAt,
	}
}
