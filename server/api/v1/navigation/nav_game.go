package navigation

import (
	"encoding/json"
	"goLinker-admin/server/global"
	"goLinker-admin/server/model/common/response"
	"goLinker-admin/server/model/navigation"
	navRequest "goLinker-admin/server/model/navigation/request"

	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
)

type NavGameApi struct{}

// CreateGame 创建游戏
// @Tags     Game
// @Summary  创建游戏
// @Security ApiKeyAuth
// @accept   application/json
// @Produce  application/json
// @Param    data body navigation.NavGame true "游戏信息"
// @Success  200  {object} response.Response{msg=string} "创建成功"
// @Router   /navigation/game/createGame [post]
func (a *NavGameApi) CreateGame(c *gin.Context) {
	var game navigation.NavGame
	err := c.ShouldBindJSON(&game)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}

	// 参数验证
	if game.Title == "" {
		response.FailWithMessage("游戏标题不能为空", c)
		return
	}
	if game.CategoryIds == "" {
		response.FailWithMessage("游戏类别不能为空", c)
		return
	}

	err = navGameService.CreateGame(game)
	if err != nil {
		global.GVA_LOG.Error("创建失败!", zap.Error(err))
		response.FailWithMessage(err.Error(), c)
	} else {
		response.OkWithMessage("创建成功", c)
	}
}

// UpdateGame 更新游戏
// @Tags     Game
// @Summary  更新游戏
// @Security ApiKeyAuth
// @accept   application/json
// @Produce  application/json
// @Param    data body navRequest.NavGameUpdateRequest true "游戏更新信息"
// @Success  200  {object} response.Response{msg=string} "更新成功"
// @Router   /navigation/game/updateGame [put]
func (a *NavGameApi) UpdateGame(c *gin.Context) {
	var req navRequest.NavGameUpdateRequest
	err := c.ShouldBindJSON(&req)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}

	err = navGameService.UpdateGame(req)
	if err != nil {
		global.GVA_LOG.Error("更新失败!", zap.Error(err))
		response.FailWithMessage(err.Error(), c)
	} else {
		response.OkWithMessage("更新成功", c)
	}
}

// DeleteGame 删除游戏
// @Tags     Game
// @Summary  删除游戏
// @Security ApiKeyAuth
// @accept   application/json
// @Produce  application/json
// @Param    data body navRequest.GetById true "游戏ID"
// @Success  200  {object} response.Response{msg=string} "删除成功"
// @Router   /navigation/game/deleteGame [delete]
func (a *NavGameApi) DeleteGame(c *gin.Context) {
	var req navRequest.GetById
	err := c.ShouldBindJSON(&req)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}

	err = navGameService.DeleteGame(req.ID)
	if err != nil {
		global.GVA_LOG.Error("删除失败!", zap.Error(err))
		response.FailWithMessage(err.Error(), c)
	} else {
		response.OkWithMessage("删除成功", c)
	}
}

// GetGameList 获取游戏列表
// @Tags     Game
// @Summary  获取游戏列表
// @Security ApiKeyAuth
// @accept   application/json
// @Produce  application/json
// @Param    data body navRequest.NavGameSearch true "分页获取游戏列表"
// @Success  200  {object} response.Response{data=response.PageResult{list=[]navigation.NavGame,total=int64}} "获取成功"
// @Router   /navigation/game/getGameList [post]
func (a *NavGameApi) GetGameList(c *gin.Context) {
	var pageInfo navRequest.NavGameSearch
	err := c.ShouldBindJSON(&pageInfo)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}

	list, total, err := navGameService.GetGameList(pageInfo)
	if err != nil {
		global.GVA_LOG.Error("获取失败!", zap.Error(err))
		response.FailWithMessage(err.Error(), c)
	} else {
		// 使用json.NewEncoder并关闭HTML转义
		c.Header("Content-Type", "application/json; charset=utf-8")
		c.Status(200)

		encoder := json.NewEncoder(c.Writer)
		encoder.SetEscapeHTML(false) // 关闭HTML转义

		responseData := response.Response{
			Code: 0,
			Data: response.PageResult{
				List:     list,
				Total:    total,
				Page:     pageInfo.Page,
				PageSize: pageInfo.PageSize,
			},
			Msg: "获取成功",
		}

		encoder.Encode(responseData)
	}
}

// UpdateGameViews 更新游戏浏览次数
// @Tags     Game
// @Summary  更新游戏浏览次数
// @Security ApiKeyAuth
// @accept   application/json
// @Produce  application/json
// @Param    data body navRequest.GetById true "游戏ID"
// @Success  200  {object} response.Response{msg=string} "更新成功"
// @Router   /navigation/game/updateGameViews [post]
func (a *NavGameApi) UpdateGameViews(c *gin.Context) {
	var req navRequest.GetById
	err := c.ShouldBindJSON(&req)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}

	err = navGameService.UpdateGameViews(req.ID)
	if err != nil {
		global.GVA_LOG.Error("更新失败!", zap.Error(err))
		response.FailWithMessage(err.Error(), c)
	} else {
		response.OkWithMessage("更新成功", c)
	}
}
