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

type NavPWAConfigApi struct{}

var navPWAConfigApiService = service.ServiceGroupApp.NavigationServiceGroup.NavPWAConfigService

// CreatePWAConfig 创建PWA配置
// @Tags     NavPWAConfig
// @Summary  创建PWA配置
// @Security ApiKeyAuth
// @accept   application/json
// @Produce  application/json
// @Param    data body navRequest.CreateNavPWAConfigRequest true "创建PWA配置"
// @Success  200  {object} response.Response{msg=string} "创建成功"
// @Router   /pwaConfig/createPWAConfig [post]
func (a *NavPWAConfigApi) CreatePWAConfig(c *gin.Context) {
	var config navRequest.CreateNavPWAConfigRequest
	err := c.ShouldBindJSON(&config)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	if err := navPWAConfigApiService.CreatePWAConfig(config); err != nil {
		global.GVA_LOG.Error("创建失败!", zap.Error(err))
		response.FailWithMessage("创建失败", c)
	} else {
		response.OkWithMessage("创建成功", c)
	}
}

// DeletePWAConfig 删除PWA配置
// @Tags     NavPWAConfig
// @Summary  删除PWA配置
// @Security ApiKeyAuth
// @accept   application/json
// @Produce  application/json
// @Param    data body request.GetById true "删除PWA配置"
// @Success  200  {object} response.Response{msg=string} "删除成功"
// @Router   /pwaConfig/deletePWAConfig [post]
func (a *NavPWAConfigApi) DeletePWAConfig(c *gin.Context) {
	var reqId request.GetById
	err := c.ShouldBindJSON(&reqId)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	if err := navPWAConfigApiService.DeletePWAConfig(uint(reqId.ID)); err != nil {
		global.GVA_LOG.Error("删除失败!", zap.Error(err))
		response.FailWithMessage("删除失败", c)
	} else {
		response.OkWithMessage("删除成功", c)
	}
}

// UpdatePWAConfig 更新PWA配置
// @Tags     NavPWAConfig
// @Summary  更新PWA配置
// @Security ApiKeyAuth
// @accept   application/json
// @Produce  application/json
// @Param    data body navRequest.UpdateNavPWAConfigRequest true "更新PWA配置"
// @Success  200  {object} response.Response{msg=string} "更新成功"
// @Router   /pwaConfig/updatePWAConfig [post]
func (a *NavPWAConfigApi) UpdatePWAConfig(c *gin.Context) {
	var config navRequest.UpdateNavPWAConfigRequest
	err := c.ShouldBindJSON(&config)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	if err := navPWAConfigApiService.UpdatePWAConfig(config); err != nil {
		global.GVA_LOG.Error("更新失败!", zap.Error(err))
		response.FailWithMessage("更新失败", c)
	} else {
		response.OkWithMessage("更新成功", c)
	}
}

// GetPWAConfigById 根据ID获取PWA配置
// @Tags     NavPWAConfig
// @Summary  根据ID获取PWA配置
// @Security ApiKeyAuth
// @accept   application/json
// @Produce  application/json
// @Param    data body request.GetById true "根据ID获取PWA配置"
// @Success  200  {object} response.Response{data=navResponse.NavPWAConfigResponse,msg=string} "获取成功"
// @Router   /pwaConfig/findPWAConfig [post]
func (a *NavPWAConfigApi) GetPWAConfigById(c *gin.Context) {
	var reqId request.GetById
	err := c.ShouldBindJSON(&reqId)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	if reConfig, err := navPWAConfigApiService.GetPWAConfigById(uint(reqId.ID)); err != nil {
		global.GVA_LOG.Error("获取失败!", zap.Error(err))
		response.FailWithMessage("获取失败", c)
	} else {
		response.OkWithDetailed(navResponse.NavPWAConfigResponse{Config: reConfig}, "获取成功", c)
	}
}

// GetPWAConfigList 获取PWA配置列表（单条记录）
// @Tags     NavPWAConfig
// @Summary  获取PWA配置列表
// @Security ApiKeyAuth
// @accept   application/json
// @Produce  application/json
// @Param    data body navRequest.NavPWAConfigSearch true "分页获取PWA配置列表"
// @Success  200  {object} response.Response{data=navResponse.NavPWAConfigListResponse,msg=string} "获取PWA配置列表成功"
// @Router   /pwaConfig/getPWAConfigList [post]
func (a *NavPWAConfigApi) GetPWAConfigList(c *gin.Context) {
	var pageInfo navRequest.NavPWAConfigSearch
	err := c.ShouldBindJSON(&pageInfo)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	if list, total, err := navPWAConfigApiService.GetPWAConfigList(pageInfo); err != nil {
		global.GVA_LOG.Error("获取失败!", zap.Error(err))
		response.FailWithMessage("获取失败", c)
	} else {
		response.OkWithDetailed(navResponse.NavPWAConfigListResponse{
			List:     list,
			Total:    total,
			Page:     pageInfo.Page,
			PageSize: pageInfo.PageSize,
		}, "获取成功", c)
	}
}

// ClearPWACache 清除PWA配置缓存
// @Tags     NavPWAConfig
// @Summary  清除PWA配置缓存
// @Security ApiKeyAuth
// @accept   application/json
// @Produce  application/json
// @Success  200  {object} response.Response{msg=string} "清除成功"
// @Failure  400  {object} response.Response{msg=string} "清除失败（如Redis未开启）"
// @Router   /pwaConfig/clearPWACache [post]
func (a *NavPWAConfigApi) ClearPWACache(c *gin.Context) {
	if err := navPWAConfigApiService.ClearPWACache(); err != nil {
		global.GVA_LOG.Error("清除缓存失败!", zap.Error(err))
		response.FailWithMessage(err.Error(), c)
	} else {
		response.OkWithMessage("清除缓存成功", c)
	}
}
