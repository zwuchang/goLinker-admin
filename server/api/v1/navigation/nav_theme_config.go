package navigation

import (
	"goLinker-admin/server/global"
	"goLinker-admin/server/model/common/response"
	navRequest "goLinker-admin/server/model/navigation/request"
	navResponse "goLinker-admin/server/model/navigation/response"
	navService "goLinker-admin/server/service/navigation"

	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
)

type NavThemeConfigApi struct{}

var themeConfigService = navService.ServiceGroupApp.NavThemeConfigService

// CreateThemeConfig 创建主题配置
// @Tags     ThemeConfig
// @Summary  创建主题配置
// @Security ApiKeyAuth
// @accept   application/json
// @Produce  application/json
// @Param    data body navRequest.CreateNavThemeConfigRequest true "主题配置信息"
// @Success  200  {object} response.Response{msg=string} "创建成功"
// @Router   /navigation/themeConfig/createThemeConfig [post]
func (a *NavThemeConfigApi) CreateThemeConfig(c *gin.Context) {
	var req navRequest.CreateNavThemeConfigRequest
	err := c.ShouldBindJSON(&req)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}

	err = themeConfigService.CreateThemeConfig(req)
	if err != nil {
		global.GVA_LOG.Error("创建失败!", zap.Error(err))
		response.FailWithMessage(err.Error(), c)
	} else {
		response.OkWithMessage("创建成功", c)
	}
}

// UpdateThemeConfig 更新主题配置
// @Tags     ThemeConfig
// @Summary  更新主题配置
// @Security ApiKeyAuth
// @accept   application/json
// @Produce  application/json
// @Param    data body navRequest.UpdateNavThemeConfigRequest true "主题配置信息"
// @Success  200  {object} response.Response{msg=string} "更新成功"
// @Router   /navigation/themeConfig/updateThemeConfig [put]
func (a *NavThemeConfigApi) UpdateThemeConfig(c *gin.Context) {
	var req navRequest.UpdateNavThemeConfigRequest
	err := c.ShouldBindJSON(&req)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}

	err = themeConfigService.UpdateThemeConfig(req)
	if err != nil {
		global.GVA_LOG.Error("更新失败!", zap.Error(err))
		response.FailWithMessage(err.Error(), c)
	} else {
		response.OkWithMessage("更新成功", c)
	}
}

// DeleteThemeConfig 删除主题配置
// @Tags     ThemeConfig
// @Summary  删除主题配置
// @Security ApiKeyAuth
// @accept   application/json
// @Produce  application/json
// @Param    data body navRequest.GetById true "主题配置ID"
// @Success  200  {object} response.Response{msg=string} "删除成功"
// @Router   /navigation/themeConfig/deleteThemeConfig [delete]
func (a *NavThemeConfigApi) DeleteThemeConfig(c *gin.Context) {
	var req navRequest.GetById
	err := c.ShouldBindJSON(&req)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}

	err = themeConfigService.DeleteThemeConfig(req.ID)
	if err != nil {
		global.GVA_LOG.Error("删除失败!", zap.Error(err))
		response.FailWithMessage(err.Error(), c)
	} else {
		response.OkWithMessage("删除成功", c)
	}
}

// GetThemeConfigList 获取主题配置列表
// @Tags     ThemeConfig
// @Summary  获取主题配置列表
// @Security ApiKeyAuth
// @accept   application/json
// @Produce  application/json
// @Param    data body navRequest.NavThemeConfigSearch true "分页获取主题配置列表"
// @Success  200  {object} response.Response{data=navResponse.NavThemeConfigListResponse,msg=string} "获取成功"
// @Router   /navigation/themeConfig/getThemeConfigList [post]
func (a *NavThemeConfigApi) GetThemeConfigList(c *gin.Context) {
	var pageInfo navRequest.NavThemeConfigSearch
	err := c.ShouldBindJSON(&pageInfo)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}

	list, total, err := themeConfigService.GetThemeConfigList(pageInfo)
	if err != nil {
		global.GVA_LOG.Error("获取失败!", zap.Error(err))
		response.FailWithMessage(err.Error(), c)
	} else {
		// 转换为响应格式
		var responseList []navResponse.NavThemeConfigResponse
		for _, item := range list {
			responseList = append(responseList, navResponse.NavThemeConfigResponse{
				ID:          item.ID,
				Name:        item.Name,
				Description: item.Description,
				ConfigJson:  item.ConfigJson,
				IsDefault:   item.IsDefault,
				CreatedAt:   item.CreatedAt.Format("2006-01-02 15:04:05"),
				UpdatedAt:   item.UpdatedAt.Format("2006-01-02 15:04:05"),
			})
		}

		responseData := navResponse.NavThemeConfigListResponse{
			List:     responseList,
			Total:    total,
			Page:     pageInfo.Page,
			PageSize: pageInfo.PageSize,
		}

		response.OkWithDetailed(responseData, "获取成功", c)
	}
}

// GetThemeConfigById 根据ID获取主题配置
// @Tags     ThemeConfig
// @Summary  根据ID获取主题配置
// @Security ApiKeyAuth
// @accept   application/json
// @Produce  application/json
// @Param    data body navRequest.GetById true "主题配置ID"
// @Success  200  {object} response.Response{data=navResponse.NavThemeConfigDetailResponse,msg=string} "获取成功"
// @Router   /navigation/themeConfig/getThemeConfigById [post]
func (a *NavThemeConfigApi) GetThemeConfigById(c *gin.Context) {
	var req navRequest.GetById
	err := c.ShouldBindJSON(&req)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}

	config, err := themeConfigService.GetThemeConfigById(req.ID)
	if err != nil {
		global.GVA_LOG.Error("获取失败!", zap.Error(err))
		response.FailWithMessage(err.Error(), c)
	} else {
		responseData := navResponse.NavThemeConfigDetailResponse{
			ID:          config.ID,
			Name:        config.Name,
			Description: config.Description,
			ConfigJson:  config.ConfigJson,
			IsDefault:   config.IsDefault,
			CreatedAt:   config.CreatedAt.Format("2006-01-02 15:04:05"),
			UpdatedAt:   config.UpdatedAt.Format("2006-01-02 15:04:05"),
		}

		response.OkWithDetailed(responseData, "获取成功", c)
	}
}

// SetDefaultTheme 设置默认主题
// @Tags     ThemeConfig
// @Summary  设置默认主题
// @Security ApiKeyAuth
// @accept   application/json
// @Produce  application/json
// @Param    data body navRequest.SetDefaultThemeRequest true "主题配置ID"
// @Success  200  {object} response.Response{msg=string} "设置成功"
// @Router   /navigation/themeConfig/setDefaultTheme [post]
func (a *NavThemeConfigApi) SetDefaultTheme(c *gin.Context) {
	var req navRequest.SetDefaultThemeRequest
	err := c.ShouldBindJSON(&req)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}

	err = themeConfigService.SetDefaultTheme(req)
	if err != nil {
		global.GVA_LOG.Error("设置失败!", zap.Error(err))
		response.FailWithMessage(err.Error(), c)
	} else {
		response.OkWithMessage("设置成功", c)
	}
}

// GetDefaultTheme 获取默认主题
// @Tags     ThemeConfig
// @Summary  获取默认主题
// @Security ApiKeyAuth
// @accept   application/json
// @Produce  application/json
// @Success  200  {object} response.Response{data=navResponse.NavThemeConfigDetailResponse,msg=string} "获取成功"
// @Router   /navigation/themeConfig/getDefaultTheme [get]
func (a *NavThemeConfigApi) GetDefaultTheme(c *gin.Context) {
	config, err := themeConfigService.GetDefaultTheme()
	if err != nil {
		global.GVA_LOG.Error("获取失败!", zap.Error(err))
		response.FailWithMessage(err.Error(), c)
	} else {
		responseData := navResponse.NavThemeConfigDetailResponse{
			ID:          config.ID,
			Name:        config.Name,
			Description: config.Description,
			ConfigJson:  config.ConfigJson,
			IsDefault:   config.IsDefault,
			CreatedAt:   config.CreatedAt.Format("2006-01-02 15:04:05"),
			UpdatedAt:   config.UpdatedAt.Format("2006-01-02 15:04:05"),
		}

		response.OkWithDetailed(responseData, "获取成功", c)
	}
}

// GetAllThemes 获取所有主题（用于前端主题切换）
// @Tags     ThemeConfig
// @Summary  获取所有主题
// @Security ApiKeyAuth
// @accept   application/json
// @Produce  application/json
// @Success  200  {object} response.Response{data=[]navResponse.PublicThemeConfigResponse,msg=string} "获取成功"
// @Router   /navigation/themeConfig/getAllThemes [get]
func (a *NavThemeConfigApi) GetAllThemes(c *gin.Context) {
	list, err := themeConfigService.GetAllThemes()
	if err != nil {
		global.GVA_LOG.Error("获取失败!", zap.Error(err))
		response.FailWithMessage(err.Error(), c)
	} else {
		// 转换为公开响应格式
		var responseList []navResponse.PublicThemeConfigResponse
		for _, item := range list {
			responseList = append(responseList, navResponse.PublicThemeConfigResponse{
				ID:         item.ID,
				Name:       item.Name,
				ConfigJson: item.ConfigJson,
				IsDefault:  item.IsDefault,
			})
		}

		response.OkWithDetailed(responseList, "获取成功", c)
	}
}
