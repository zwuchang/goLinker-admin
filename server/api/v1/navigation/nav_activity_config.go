package navigation

import (
	"goLinker-admin/server/global"
	"goLinker-admin/server/model/common/request"
	"goLinker-admin/server/model/common/response"
	navRequest "goLinker-admin/server/model/navigation/request"
	navResponse "goLinker-admin/server/model/navigation/response"
	"goLinker-admin/server/service"

	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
)

type NavActivityConfigApi struct{}

var navActivityConfigApiService = service.ServiceGroupApp.NavigationServiceGroup.NavActivityConfigService

// GetActivityConfigList 获取活动配置列表
// @Tags     NavActivityConfig
// @Summary  获取活动配置列表
// @Security ApiKeyAuth
// @accept   application/json
// @Produce  application/json
// @Param    data body navRequest.NavActivityConfigSearch true "分页获取活动配置列表"
// @Success  200  {object} response.Response{data=navResponse.NavActivityConfigListResponse,msg=string} "获取活动配置列表成功"
// @Router   /activityConfig/getActivityConfigList [post]
func (a *NavActivityConfigApi) GetActivityConfigList(c *gin.Context) {
	var pageInfo navRequest.NavActivityConfigSearch
	err := c.ShouldBindJSON(&pageInfo)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	if list, total, err := navActivityConfigApiService.GetActivityConfigList(pageInfo); err != nil {
		global.GVA_LOG.Error("获取失败!", zap.Error(err))
		response.FailWithMessage("获取失败", c)
	} else {
		// 转换为响应格式
		var responseList []navResponse.NavActivityConfigResponse
		for _, item := range list {
			responseList = append(responseList, navResponse.NavActivityConfigResponse{
				ID:           item.ID,
				Title:        item.Title,
				Image:        item.Image,
				JumpUrl:      item.JumpUrl,
				CategoryName: item.CategoryName,
				CategoryIcon: item.CategoryIcon,
				Content:      item.Content,
				Status:       item.Status,
				IsVisible:    item.IsVisible,
				Sort:         item.Sort,
				CreatedAt:    item.CreatedAt.Format("2006-01-02 15:04:05"),
				UpdatedAt:    item.UpdatedAt.Format("2006-01-02 15:04:05"),
			})
		}
		response.OkWithDetailed(navResponse.NavActivityConfigListResponse{
			List:     responseList,
			Total:    total,
			Page:     pageInfo.Page,
			PageSize: pageInfo.PageSize,
		}, "获取成功", c)
	}
}

// CreateActivityConfig 创建活动配置
// @Tags     NavActivityConfig
// @Summary  创建活动配置
// @Security ApiKeyAuth
// @accept   application/json
// @Produce  application/json
// @Param    data body navRequest.CreateNavActivityConfigRequest true "创建活动配置"
// @Success  200  {object} response.Response{msg=string} "创建成功"
// @Router   /activityConfig/createActivityConfig [post]
func (a *NavActivityConfigApi) CreateActivityConfig(c *gin.Context) {
	var req navRequest.CreateNavActivityConfigRequest
	err := c.ShouldBindJSON(&req)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	if err := navActivityConfigApiService.CreateActivityConfig(req); err != nil {
		global.GVA_LOG.Error("创建失败!", zap.Error(err))
		response.FailWithMessage("创建失败", c)
	} else {
		response.OkWithMessage("创建成功", c)
	}
}

// UpdateActivityConfig 更新活动配置
// @Tags     NavActivityConfig
// @Summary  更新活动配置
// @Security ApiKeyAuth
// @accept   application/json
// @Produce  application/json
// @Param    data body navRequest.UpdateNavActivityConfigRequest true "更新活动配置"
// @Success  200  {object} response.Response{msg=string} "更新成功"
// @Router   /activityConfig/updateActivityConfig [post]
func (a *NavActivityConfigApi) UpdateActivityConfig(c *gin.Context) {
	var req navRequest.UpdateNavActivityConfigRequest
	err := c.ShouldBindJSON(&req)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	if err := navActivityConfigApiService.UpdateActivityConfig(req); err != nil {
		global.GVA_LOG.Error("更新失败!", zap.Error(err))
		response.FailWithMessage("更新失败", c)
	} else {
		response.OkWithMessage("更新成功", c)
	}
}

// DeleteActivityConfig 删除活动配置
// @Tags     NavActivityConfig
// @Summary  删除活动配置
// @Security ApiKeyAuth
// @accept   application/json
// @Produce  application/json
// @Param    data body request.GetById true "删除活动配置"
// @Success  200  {object} response.Response{msg=string} "删除成功"
// @Router   /activityConfig/deleteActivityConfig [post]
func (a *NavActivityConfigApi) DeleteActivityConfig(c *gin.Context) {
	var req request.GetById
	err := c.ShouldBindJSON(&req)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	if err := navActivityConfigApiService.DeleteActivityConfig(uint(req.ID)); err != nil {
		global.GVA_LOG.Error("删除失败!", zap.Error(err))
		response.FailWithMessage("删除失败", c)
	} else {
		response.OkWithMessage("删除成功", c)
	}
}

// GetActivityConfigById 根据ID获取活动配置
// @Tags     NavActivityConfig
// @Summary  根据ID获取活动配置
// @Security ApiKeyAuth
// @accept   application/json
// @Produce  application/json
// @Param    data body request.GetById true "根据ID获取活动配置"
// @Success  200  {object} response.Response{data=navResponse.NavActivityConfigDetailResponse,msg=string} "获取成功"
// @Router   /activityConfig/getActivityConfigById [post]
func (a *NavActivityConfigApi) GetActivityConfigById(c *gin.Context) {
	var req request.GetById
	err := c.ShouldBindJSON(&req)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	if reActivityConfig, err := navActivityConfigApiService.GetActivityConfigById(uint(req.ID)); err != nil {
		global.GVA_LOG.Error("获取失败!", zap.Error(err))
		response.FailWithMessage("获取失败", c)
	} else {
		response.OkWithDetailed(navResponse.NavActivityConfigDetailResponse{
			Config: reActivityConfig,
		}, "获取成功", c)
	}
}
