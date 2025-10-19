package navigation

import (
	"goLinker-admin/server/global"
	"goLinker-admin/server/model/common/response"
	navRequest "goLinker-admin/server/model/navigation/request"
	navResponse "goLinker-admin/server/model/navigation/response"
	"goLinker-admin/server/service"
	"strconv"

	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
)

type NavMarketConfigApi struct{}

// GetMarketConfigList 获取市场配置列表
// @Tags     MarketConfig
// @Summary  获取市场配置列表
// @Security ApiKeyAuth
// @accept   application/json
// @Produce  application/json
// @Param    data body navRequest.NavMarketConfigSearch true "市场配置搜索条件"
// @Success  200  {object} response.Response{data=navResponse.NavMarketConfigListResponse} "获取成功"
// @Router   /navigation/marketConfig/getMarketConfigList [post]
func (a *NavMarketConfigApi) GetMarketConfigList(c *gin.Context) {
	var pageInfo navRequest.NavMarketConfigSearch
	err := c.ShouldBindJSON(&pageInfo)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}

	list, total, err := service.ServiceGroupApp.NavigationServiceGroup.NavMarketConfigService.GetMarketConfigList(pageInfo)
	if err != nil {
		global.GVA_LOG.Error("获取失败!", zap.Error(err))
		response.FailWithMessage(err.Error(), c)
	} else {
		// 转换为响应格式
		var resList []navResponse.NavMarketConfigResponse
		for _, item := range list {
			resList = append(resList, navResponse.NavMarketConfigResponse{
				ID:          item.ID,
				Name:        item.Name,
				Logo:        item.Logo,
				JumpUrl:     item.JumpUrl,
				Type:        item.Type,
				Status:      item.Status,
				IsVisible:   item.IsVisible,
				Sort:        item.Sort,
				Description: item.Description,
				CreatedAt:   item.CreatedAt.Format("2006-01-02 15:04:05"),
				UpdatedAt:   item.UpdatedAt.Format("2006-01-02 15:04:05"),
			})
		}

		response.OkWithDetailed(navResponse.NavMarketConfigListResponse{
			List:     resList,
			Total:    total,
			Page:     pageInfo.Page,
			PageSize: pageInfo.PageSize,
		}, "获取成功", c)
	}
}

// CreateMarketConfig 创建市场配置
// @Tags     MarketConfig
// @Summary  创建市场配置
// @Security ApiKeyAuth
// @accept   application/json
// @Produce  application/json
// @Param    data body navRequest.CreateMarketConfig true "市场配置信息"
// @Success  200  {object} response.Response{msg=string} "创建成功"
// @Router   /navigation/marketConfig/createMarketConfig [post]
func (a *NavMarketConfigApi) CreateMarketConfig(c *gin.Context) {
	var req navRequest.CreateMarketConfig
	err := c.ShouldBindJSON(&req)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}

	err = service.ServiceGroupApp.NavigationServiceGroup.NavMarketConfigService.CreateMarketConfig(req)
	if err != nil {
		global.GVA_LOG.Error("创建失败!", zap.Error(err))
		response.FailWithMessage(err.Error(), c)
	} else {
		response.OkWithMessage("创建成功", c)
	}
}

// UpdateMarketConfig 更新市场配置
// @Tags     MarketConfig
// @Summary  更新市场配置
// @Security ApiKeyAuth
// @accept   application/json
// @Produce  application/json
// @Param    data body navRequest.UpdateMarketConfig true "市场配置信息"
// @Success  200  {object} response.Response{msg=string} "更新成功"
// @Router   /navigation/marketConfig/updateMarketConfig [post]
func (a *NavMarketConfigApi) UpdateMarketConfig(c *gin.Context) {
	var req navRequest.UpdateMarketConfig
	err := c.ShouldBindJSON(&req)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}

	err = service.ServiceGroupApp.NavigationServiceGroup.NavMarketConfigService.UpdateMarketConfig(req)
	if err != nil {
		global.GVA_LOG.Error("更新失败!", zap.Error(err))
		response.FailWithMessage(err.Error(), c)
	} else {
		response.OkWithMessage("更新成功", c)
	}
}

// GetMarketConfigById 根据ID获取市场配置
// @Tags     MarketConfig
// @Summary  根据ID获取市场配置
// @Security ApiKeyAuth
// @accept   application/json
// @Produce  application/json
// @Param    id path uint true "市场配置ID"
// @Success  200  {object} response.Response{data=navResponse.NavMarketConfigResponse} "获取成功"
// @Router   /navigation/marketConfig/getMarketConfigById/{id} [get]
func (a *NavMarketConfigApi) GetMarketConfigById(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		response.FailWithMessage("参数错误", c)
		return
	}

	marketConfig, err := service.ServiceGroupApp.NavigationServiceGroup.NavMarketConfigService.GetMarketConfigById(uint(id))
	if err != nil {
		global.GVA_LOG.Error("获取失败!", zap.Error(err))
		response.FailWithMessage(err.Error(), c)
	} else {
		responseData := navResponse.NavMarketConfigResponse{
			ID:          marketConfig.ID,
			Name:        marketConfig.Name,
			Logo:        marketConfig.Logo,
			JumpUrl:     marketConfig.JumpUrl,
			Type:        marketConfig.Type,
			Status:      marketConfig.Status,
			IsVisible:   marketConfig.IsVisible,
			Sort:        marketConfig.Sort,
			Description: marketConfig.Description,
			CreatedAt:   marketConfig.CreatedAt.Format("2006-01-02 15:04:05"),
			UpdatedAt:   marketConfig.UpdatedAt.Format("2006-01-02 15:04:05"),
		}
		response.OkWithDetailed(responseData, "获取成功", c)
	}
}

// DeleteMarketConfig 删除市场配置
// @Tags     MarketConfig
// @Summary  删除市场配置
// @Security ApiKeyAuth
// @accept   application/json
// @Produce  application/json
// @Param    data body navRequest.DeleteMarketConfig true "市场配置ID"
// @Success  200  {object} response.Response{msg=string} "删除成功"
// @Router   /navigation/marketConfig/deleteMarketConfig [post]
func (a *NavMarketConfigApi) DeleteMarketConfig(c *gin.Context) {
	var req navRequest.DeleteMarketConfig
	err := c.ShouldBindJSON(&req)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}

	err = service.ServiceGroupApp.NavigationServiceGroup.NavMarketConfigService.DeleteMarketConfig(req.ID)
	if err != nil {
		global.GVA_LOG.Error("删除失败!", zap.Error(err))
		response.FailWithMessage(err.Error(), c)
	} else {
		response.OkWithMessage("删除成功", c)
	}
}
