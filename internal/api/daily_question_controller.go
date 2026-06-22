package api

import (
	"blog/internal/model/dto/request"
	"blog/internal/service"
	"blog/pkg/response"
	"strconv"

	"github.com/gin-gonic/gin"
)

// DailyQuestionController 每日一问控制器
type DailyQuestionController struct {
	dailyQSvc service.DailyQuestionService
}

// NewDailyQuestionController 创建每日一问控制器
func NewDailyQuestionController(dailyQSvc service.DailyQuestionService) *DailyQuestionController {
	return &DailyQuestionController{dailyQSvc: dailyQSvc}
}

// RegisterPublicRoutes 注册公开路由
func (c *DailyQuestionController) RegisterPublicRoutes(rg *gin.RouterGroup) {
	dailyQ := rg.Group("/daily-questions")
	{
		dailyQ.GET("/latest", c.GetLatestQuestion)
		dailyQ.GET("/date/:date", c.GetQuestionByDate)
		dailyQ.GET("/previous/:date", c.GetPreviousQuestion)
		dailyQ.GET("/next/:date", c.GetNextQuestion)
		dailyQ.POST("/:id/like", c.LikeQuestion)
	}
}

// RegisterAdminRoutes 注册后台路由
func (c *DailyQuestionController) RegisterAdminRoutes(rg *gin.RouterGroup) {
	admin := rg.Group("/daily-questions")
	{
		admin.GET("", c.GetAdminQuestionList)
		admin.POST("", c.CreateQuestion)
		admin.PUT("/:id", c.UpdateQuestion)
		admin.DELETE("/:id", c.DeleteQuestion)
		admin.PUT("/:id/status", c.UpdateQuestionStatus)
	}
}

// GetLatestQuestion 获取最新问题
func (c *DailyQuestionController) GetLatestQuestion(ctx *gin.Context) {
	result, err := c.dailyQSvc.GetLatestQuestion()
	if err != nil {
		response.Error(ctx, 500, "获取最新问题失败")
		return
	}

	response.Success(ctx, result)
}

// GetQuestionByDate 根据日期获取问题
func (c *DailyQuestionController) GetQuestionByDate(ctx *gin.Context) {
	date := ctx.Param("date")
	if date == "" {
		response.Error(ctx, 400, "日期不能为空")
		return
	}

	result, err := c.dailyQSvc.GetQuestionByDate(date)
	if err != nil {
		response.Error(ctx, 404, "该日期没有问题")
		return
	}

	response.Success(ctx, result)
}

// GetPreviousQuestion 获取前一天的问题
func (c *DailyQuestionController) GetPreviousQuestion(ctx *gin.Context) {
	date := ctx.Param("date")
	if date == "" {
		response.Error(ctx, 400, "日期不能为空")
		return
	}

	result, err := c.dailyQSvc.GetPreviousQuestion(date)
	if err != nil {
		response.Error(ctx, 404, "没有前一天的问题")
		return
	}

	response.Success(ctx, result)
}

// GetNextQuestion 获取后一天的问题
func (c *DailyQuestionController) GetNextQuestion(ctx *gin.Context) {
	date := ctx.Param("date")
	if date == "" {
		response.Error(ctx, 400, "日期不能为空")
		return
	}

	result, err := c.dailyQSvc.GetNextQuestion(date)
	if err != nil {
		response.Error(ctx, 404, "没有后一天的问题")
		return
	}

	response.Success(ctx, result)
}

// LikeQuestion 问题点赞
func (c *DailyQuestionController) LikeQuestion(ctx *gin.Context) {
	id, err := strconv.ParseUint(ctx.Param("id"), 10, 32)
	if err != nil {
		response.Error(ctx, 400, "问题ID无效")
		return
	}

	likeCount, err := c.dailyQSvc.LikeQuestion(uint(id))
	if err != nil {
		response.Error(ctx, 500, "点赞失败")
		return
	}

	response.Success(ctx, gin.H{"like_count": likeCount})
}

// GetAdminQuestionList 获取问题列表（后台）
func (c *DailyQuestionController) GetAdminQuestionList(ctx *gin.Context) {
	var req request.DailyQuestionListRequest
	if err := ctx.ShouldBindQuery(&req); err != nil {
		response.Error(ctx, 400, "参数错误")
		return
	}

	result, err := c.dailyQSvc.GetAdminQuestionList(&req)
	if err != nil {
		response.Error(ctx, 500, "获取问题列表失败")
		return
	}

	response.Success(ctx, result)
}

// CreateQuestion 创建问题
func (c *DailyQuestionController) CreateQuestion(ctx *gin.Context) {
	var req request.CreateDailyQuestionRequest
	if err := ctx.ShouldBindJSON(&req); err != nil {
		response.Error(ctx, 400, "参数错误: "+err.Error())
		return
	}

	id, err := c.dailyQSvc.CreateQuestion(&req)
	if err != nil {
		response.Error(ctx, 500, "创建问题失败: "+err.Error())
		return
	}

	response.Success(ctx, gin.H{"id": id})
}

// UpdateQuestion 更新问题
func (c *DailyQuestionController) UpdateQuestion(ctx *gin.Context) {
	id, err := strconv.ParseUint(ctx.Param("id"), 10, 32)
	if err != nil {
		response.Error(ctx, 400, "问题ID无效")
		return
	}

	var req request.UpdateDailyQuestionRequest
	if err := ctx.ShouldBindJSON(&req); err != nil {
		response.Error(ctx, 400, "参数错误: "+err.Error())
		return
	}

	err = c.dailyQSvc.UpdateQuestion(uint(id), &req)
	if err != nil {
		response.Error(ctx, 500, "更新问题失败: "+err.Error())
		return
	}

	response.Success(ctx, nil)
}

// DeleteQuestion 删除问题
func (c *DailyQuestionController) DeleteQuestion(ctx *gin.Context) {
	id, err := strconv.ParseUint(ctx.Param("id"), 10, 32)
	if err != nil {
		response.Error(ctx, 400, "问题ID无效")
		return
	}

	err = c.dailyQSvc.DeleteQuestion(uint(id))
	if err != nil {
		response.Error(ctx, 500, "删除问题失败")
		return
	}

	response.Success(ctx, nil)
}

// UpdateQuestionStatus 更新问题状态
func (c *DailyQuestionController) UpdateQuestionStatus(ctx *gin.Context) {
	id, err := strconv.ParseUint(ctx.Param("id"), 10, 32)
	if err != nil {
		response.Error(ctx, 400, "问题ID无效")
		return
	}

	var req struct {
		Status int `json:"status" binding:"required"`
	}
	if err := ctx.ShouldBindJSON(&req); err != nil {
		response.Error(ctx, 400, "参数错误")
		return
	}

	err = c.dailyQSvc.UpdateQuestionStatus(uint(id), req.Status)
	if err != nil {
		response.Error(ctx, 500, "更新问题状态失败")
		return
	}

	response.Success(ctx, nil)
}
