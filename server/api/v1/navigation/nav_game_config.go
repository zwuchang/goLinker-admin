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

type NavGameConfigApi struct{}

// CreateGameConfig 创建游戏配置
// @Tags     NavGameConfig
// @Summary  创建游戏配置
// @Security ApiKeyAuth
// @accept   application/json
// @Produce  application/json
// @Param    data body navRequest.CreateNavGameConfigRequest true "游戏配置信息"
// @Success  200  {object} response.Response{msg=string} "创建成功"
// @Router   /navigation/gameConfig/createGameConfig [post]
func (api *NavGameConfigApi) CreateGameConfig(c *gin.Context) {
	var gameConfig navRequest.CreateNavGameConfigRequest
	err := c.ShouldBindJSON(&gameConfig)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}

	// 参数验证
	if gameConfig.DownloadUrl == "" {
		response.FailWithMessage("下载地址不能为空", c)
		return
	}
	if gameConfig.WebsiteTitle == "" {
		response.FailWithMessage("网站标题不能为空", c)
		return
	}

	err = navService.ServiceGroupApp.NavGameConfigService.CreateGameConfig(gameConfig)
	if err != nil {
		global.GVA_LOG.Error("创建失败!", zap.Error(err))
		response.FailWithMessage(err.Error(), c)
	} else {
		response.OkWithMessage("创建成功", c)
	}
}

// UpdateGameConfig 更新游戏配置
// @Tags     NavGameConfig
// @Summary  更新游戏配置
// @Security ApiKeyAuth
// @accept   application/json
// @Produce  application/json
// @Param    data body navRequest.UpdateNavGameConfigRequest true "游戏配置信息"
// @Success  200  {object} response.Response{msg=string} "更新成功"
// @Router   /navigation/gameConfig/updateGameConfig [put]
func (api *NavGameConfigApi) UpdateGameConfig(c *gin.Context) {
	var gameConfig navRequest.UpdateNavGameConfigRequest
	err := c.ShouldBindJSON(&gameConfig)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}

	// 参数验证
	if gameConfig.ID == 0 {
		response.FailWithMessage("ID不能为空", c)
		return
	}
	if gameConfig.DownloadUrl == "" {
		response.FailWithMessage("下载地址不能为空", c)
		return
	}
	if gameConfig.WebsiteTitle == "" {
		response.FailWithMessage("网站标题不能为空", c)
		return
	}

	err = navService.ServiceGroupApp.NavGameConfigService.UpdateGameConfig(gameConfig)
	if err != nil {
		global.GVA_LOG.Error("更新失败!", zap.Error(err))
		response.FailWithMessage(err.Error(), c)
	} else {
		response.OkWithMessage("更新成功", c)
	}
}

// GetGameConfig 获取游戏配置
// @Tags     NavGameConfig
// @Summary  获取游戏配置
// @Security ApiKeyAuth
// @accept   application/json
// @Produce  application/json
// @Success  200  {object} response.Response{data=navResponse.NavGameConfigResponse,msg=string} "获取成功"
// @Router   /navigation/gameConfig/getGameConfig [get]
func (api *NavGameConfigApi) GetGameConfig(c *gin.Context) {
	gameConfig, err := navService.ServiceGroupApp.NavGameConfigService.GetGameConfig()
	if err != nil {
		global.GVA_LOG.Error("获取失败!", zap.Error(err))
		response.FailWithMessage(err.Error(), c)
	} else {
		responseData := navResponse.NavGameConfigResponse{
			ID:           gameConfig.ID,
			DownloadUrl:  gameConfig.DownloadUrl,
			AudioUrl:     gameConfig.AudioUrl,
			WebsiteTitle: gameConfig.WebsiteTitle,
			WebsiteDesc:  gameConfig.WebsiteDesc,
			WebsiteIcon:  gameConfig.WebsiteIcon,
			WebsiteLogo:  gameConfig.WebsiteLogo,
			MarketLogo:   gameConfig.MarketLogo,
			Status:       gameConfig.Status,
			CreatedAt:    gameConfig.CreatedAt.Format("2006-01-02 15:04:05"),
			UpdatedAt:    gameConfig.UpdatedAt.Format("2006-01-02 15:04:05"),
		}
		response.OkWithDetailed(responseData, "获取成功", c)
	}
}

// GetGameConfigList 获取游戏配置列表
// @Tags     NavGameConfig
// @Summary  获取游戏配置列表
// @Security ApiKeyAuth
// @accept   application/json
// @Produce  application/json
// @Param    data body navRequest.NavGameConfigSearch true "分页获取游戏配置列表"
// @Success  200  {object} response.Response{data=response.PageResult,msg=string} "获取成功"
// @Router   /navigation/gameConfig/getGameConfigList [post]
func (api *NavGameConfigApi) GetGameConfigList(c *gin.Context) {
	var pageInfo navRequest.NavGameConfigSearch
	err := c.ShouldBindJSON(&pageInfo)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}

	list, total, err := navService.ServiceGroupApp.NavGameConfigService.GetGameConfigList(pageInfo)
	if err != nil {
		global.GVA_LOG.Error("获取失败!", zap.Error(err))
		response.FailWithMessage(err.Error(), c)
	} else {
		response.OkWithDetailed(response.PageResult{
			List:     list,
			Total:    total,
			Page:     pageInfo.Page,
			PageSize: pageInfo.PageSize,
		}, "获取成功", c)
	}
}

// DeleteGameConfig 删除游戏配置
// @Tags     NavGameConfig
// @Summary  删除游戏配置
// @Security ApiKeyAuth
// @accept   application/json
// @Produce  application/json
// @Param    data body navRequest.GetById true "ID"
// @Success  200  {object} response.Response{msg=string} "删除成功"
// @Router   /navigation/gameConfig/deleteGameConfig [delete]
func (api *NavGameConfigApi) DeleteGameConfig(c *gin.Context) {
	var req navRequest.GetById
	err := c.ShouldBindJSON(&req)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}

	err = navGameConfigService.DeleteGameConfig(req.ID)
	if err != nil {
		global.GVA_LOG.Error("删除失败!", zap.Error(err))
		response.FailWithMessage(err.Error(), c)
	} else {
		response.OkWithMessage("删除成功", c)
	}
}
