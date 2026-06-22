package api

import (
	"blog/internal/model/dto/request"
	"blog/internal/service"
	"blog/pkg/response"
	"strconv"

	"github.com/gin-gonic/gin"
)

// CategoryController 分类控制器
type CategoryController struct {
	categorySvc service.CategoryService
}

// NewCategoryController 创建分类控制器
func NewCategoryController(categorySvc service.CategoryService) *CategoryController {
	return &CategoryController{categorySvc: categorySvc}
}

// RegisterPublicRoutes 注册公开路由
func (c *CategoryController) RegisterPublicRoutes(rg *gin.RouterGroup) {
	categories := rg.Group("/categories")
	{
		categories.GET("", c.GetCategoryList)
		categories.GET("/:id", c.GetCategoryDetail)
	}
}

// RegisterAdminRoutes 注册后台路由
func (c *CategoryController) RegisterAdminRoutes(rg *gin.RouterGroup) {
	admin := rg.Group("/categories")
	{
		admin.POST("", c.CreateCategory)
		admin.PUT("/:id", c.UpdateCategory)
		admin.DELETE("/:id", c.DeleteCategory)
	}
}

// GetCategoryList 获取分类列表
func (c *CategoryController) GetCategoryList(ctx *gin.Context) {
	result, err := c.categorySvc.GetCategoryList()
	if err != nil {
		response.Error(ctx, 500, "获取分类列表失败")
		return
	}

	response.Success(ctx, result)
}

// GetCategoryDetail 获取分类详情
func (c *CategoryController) GetCategoryDetail(ctx *gin.Context) {
	id, err := strconv.ParseUint(ctx.Param("id"), 10, 32)
	if err != nil {
		response.Error(ctx, 400, "分类ID无效")
		return
	}

	result, err := c.categorySvc.GetCategoryByID(uint(id))
	if err != nil {
		response.Error(ctx, 404, "分类不存在")
		return
	}

	response.Success(ctx, result)
}

// CreateCategory 创建分类
func (c *CategoryController) CreateCategory(ctx *gin.Context) {
	var req request.CreateCategoryRequest
	if err := ctx.ShouldBindJSON(&req); err != nil {
		response.Error(ctx, 400, "参数错误: "+err.Error())
		return
	}

	id, err := c.categorySvc.CreateCategory(&req)
	if err != nil {
		response.Error(ctx, 500, "创建分类失败: "+err.Error())
		return
	}

	response.Success(ctx, gin.H{"id": id})
}

// UpdateCategory 更新分类
func (c *CategoryController) UpdateCategory(ctx *gin.Context) {
	id, err := strconv.ParseUint(ctx.Param("id"), 10, 32)
	if err != nil {
		response.Error(ctx, 400, "分类ID无效")
		return
	}

	var req request.UpdateCategoryRequest
	if err := ctx.ShouldBindJSON(&req); err != nil {
		response.Error(ctx, 400, "参数错误: "+err.Error())
		return
	}

	err = c.categorySvc.UpdateCategory(uint(id), &req)
	if err != nil {
		response.Error(ctx, 500, "更新分类失败: "+err.Error())
		return
	}

	response.Success(ctx, nil)
}

// DeleteCategory 删除分类
func (c *CategoryController) DeleteCategory(ctx *gin.Context) {
	id, err := strconv.ParseUint(ctx.Param("id"), 10, 32)
	if err != nil {
		response.Error(ctx, 400, "分类ID无效")
		return
	}

	err = c.categorySvc.DeleteCategory(uint(id))
	if err != nil {
		response.Error(ctx, 500, "删除分类失败")
		return
	}

	response.Success(ctx, nil)
}
