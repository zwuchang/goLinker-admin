package navigation

import (
	"goLinker-admin/server/global"
	"goLinker-admin/server/model/common/response"
	navRequest "goLinker-admin/server/model/navigation/request"
	navResponse "goLinker-admin/server/model/navigation/response"
	"strconv"

	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
)

type NavPlatformConfigApi struct{}

// GetPlatformConfigList 获取平台游戏配置列表
// @Tags     PlatformConfig
// @Summary  获取平台游戏配置列表
// @Security ApiKeyAuth
// @accept   application/json
// @Produce  application/json
// @Param    data body navRequest.NavPlatformConfigSearch true "分页获取平台游戏配置列表"
// @Success  200  {object} response.Response{data=response.PageResult{list=[]navResponse.NavPlatformConfigResponse,total=int64}} "获取成功"
// @Router   /navigation/platformConfig/getPlatformConfigList [post]
func (a *NavPlatformConfigApi) GetPlatformConfigList(c *gin.Context) {
	var pageInfo navRequest.NavPlatformConfigSearch
	err := c.ShouldBindJSON(&pageInfo)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}

	list, total, err := navPlatformConfigService.GetPlatformConfigList(pageInfo)
	if err != nil {
		global.GVA_LOG.Error("获取失败!", zap.Error(err))
		response.FailWithMessage(err.Error(), c)
	} else {
		// 转换为响应格式
		var responseList []navResponse.NavPlatformConfigResponse
		for _, item := range list {
			responseList = append(responseList, navResponse.NavPlatformConfigResponse{
				ID:           item.ID,
				PlatformName: item.PlatformName,
				PlatformIcon: item.PlatformIcon,
				PlatformApi:  item.PlatformApi,
				Sort:         item.Sort,
				Status:       item.Status,
				Description:  item.Description,
				IsVisible:    item.IsVisible,
				CreatedAt:    item.CreatedAt.Format("2006-01-02 15:04:05"),
				UpdatedAt:    item.UpdatedAt.Format("2006-01-02 15:04:05"),
			})
		}
		response.OkWithDetailed(response.PageResult{
			List:     responseList,
			Total:    total,
			Page:     pageInfo.Page,
			PageSize: pageInfo.PageSize,
		}, "获取成功", c)
	}
}

// CreatePlatformConfig 创建平台游戏配置
// @Tags     PlatformConfig
// @Summary  创建平台游戏配置
// @Security ApiKeyAuth
// @accept   application/json
// @Produce  application/json
// @Param    data body navRequest.CreateNavPlatformConfigRequest true "创建平台游戏配置"
// @Success  200  {object} response.Response{msg=string} "创建成功"
// @Router   /navigation/platformConfig/createPlatformConfig [post]
func (a *NavPlatformConfigApi) CreatePlatformConfig(c *gin.Context) {
	var req navRequest.CreateNavPlatformConfigRequest
	err := c.ShouldBindJSON(&req)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}

	err = navPlatformConfigService.CreatePlatformConfig(req)
	if err != nil {
		global.GVA_LOG.Error("创建失败!", zap.Error(err))
		response.FailWithMessage(err.Error(), c)
	} else {
		response.OkWithMessage("创建成功", c)
	}
}

// UpdatePlatformConfig 更新平台游戏配置
// @Tags     PlatformConfig
// @Summary  更新平台游戏配置
// @Security ApiKeyAuth
// @accept   application/json
// @Produce  application/json
// @Param    data body navRequest.UpdateNavPlatformConfigRequest true "更新平台游戏配置"
// @Success  200  {object} response.Response{msg=string} "更新成功"
// @Router   /navigation/platformConfig/updatePlatformConfig [post]
func (a *NavPlatformConfigApi) UpdatePlatformConfig(c *gin.Context) {
	var req navRequest.UpdateNavPlatformConfigRequest
	err := c.ShouldBindJSON(&req)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}

	err = navPlatformConfigService.UpdatePlatformConfig(req)
	if err != nil {
		global.GVA_LOG.Error("更新失败!", zap.Error(err))
		response.FailWithMessage(err.Error(), c)
	} else {
		response.OkWithMessage("更新成功", c)
	}
}

// GetPlatformConfigById 根据ID获取平台游戏配置
// @Tags     PlatformConfig
// @Summary  根据ID获取平台游戏配置
// @Security ApiKeyAuth
// @accept   application/json
// @Produce  application/json
// @Param    id path int true "平台游戏配置ID"
// @Success  200  {object} response.Response{data=navResponse.NavPlatformConfigResponse} "获取成功"
// @Router   /navigation/platformConfig/getPlatformConfigById/{id} [get]
func (a *NavPlatformConfigApi) GetPlatformConfigById(c *gin.Context) {
	idStr := c.Param("id")
	id, err := strconv.ParseUint(idStr, 10, 32)
	if err != nil {
		response.FailWithMessage("参数错误", c)
		return
	}

	platformConfig, err := navPlatformConfigService.GetPlatformConfigById(uint(id))
	if err != nil {
		global.GVA_LOG.Error("获取失败!", zap.Error(err))
		response.FailWithMessage(err.Error(), c)
	} else {
		responseData := navResponse.NavPlatformConfigResponse{
			ID:           platformConfig.ID,
			PlatformName: platformConfig.PlatformName,
			PlatformIcon: platformConfig.PlatformIcon,
			PlatformApi:  platformConfig.PlatformApi,
			Sort:         platformConfig.Sort,
			Status:       platformConfig.Status,
			Description:  platformConfig.Description,
			IsVisible:    platformConfig.IsVisible,
			CreatedAt:    platformConfig.CreatedAt.Format("2006-01-02 15:04:05"),
			UpdatedAt:    platformConfig.UpdatedAt.Format("2006-01-02 15:04:05"),
		}
		response.OkWithDetailed(responseData, "获取成功", c)
	}
}

// DeletePlatformConfig 删除平台游戏配置
// @Tags     PlatformConfig
// @Summary  删除平台游戏配置
// @Security ApiKeyAuth
// @accept   application/json
// @Produce  application/json
// @Param    data body navRequest.DeleteNavPlatformConfigRequest true "删除平台游戏配置"
// @Success  200  {object} response.Response{msg=string} "删除成功"
// @Router   /navigation/platformConfig/deletePlatformConfig [post]
func (a *NavPlatformConfigApi) DeletePlatformConfig(c *gin.Context) {
	var req navRequest.DeleteNavPlatformConfigRequest
	err := c.ShouldBindJSON(&req)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}

	err = navPlatformConfigService.DeletePlatformConfig(req.ID)
	if err != nil {
		global.GVA_LOG.Error("删除失败!", zap.Error(err))
		response.FailWithMessage(err.Error(), c)
	} else {
		response.OkWithMessage("删除成功", c)
	}
}
