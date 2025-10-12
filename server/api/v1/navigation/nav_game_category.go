package navigation

import (
	"goLinker-admin/server/global"
	"goLinker-admin/server/model/common/response"
	"goLinker-admin/server/model/navigation"
	navRequest "goLinker-admin/server/model/navigation/request"

	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
)

type NavGameCategoryApi struct{}

// CreateGameCategory 创建游戏类别
// @Tags     GameCategory
// @Summary  创建游戏类别
// @Security ApiKeyAuth
// @accept   application/json
// @Produce  application/json
// @Param    data body navigation.NavGameCategory true "游戏类别信息"
// @Success  200  {object} response.Response{msg=string} "创建成功"
// @Router   /navigation/gameCategory/createGameCategory [post]
func (a *NavGameCategoryApi) CreateGameCategory(c *gin.Context) {
	var category navigation.NavGameCategory
	err := c.ShouldBindJSON(&category)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}

	// 参数验证
	if category.Name == "" {
		response.FailWithMessage("类别名称不能为空", c)
		return
	}

	err = navGameCategoryService.CreateGameCategory(category)
	if err != nil {
		global.GVA_LOG.Error("创建失败!", zap.Error(err))
		response.FailWithMessage(err.Error(), c)
	} else {
		response.OkWithMessage("创建成功", c)
	}
}

// UpdateGameCategory 更新游戏类别
// @Tags     GameCategory
// @Summary  更新游戏类别
// @Security ApiKeyAuth
// @accept   application/json
// @Produce  application/json
// @Param    data body navigation.NavGameCategory true "游戏类别信息"
// @Success  200  {object} response.Response{msg=string} "更新成功"
// @Router   /navigation/gameCategory/updateGameCategory [put]
func (a *NavGameCategoryApi) UpdateGameCategory(c *gin.Context) {
	var category navigation.NavGameCategory
	err := c.ShouldBindJSON(&category)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}

	// 参数验证
	if category.ID == 0 {
		response.FailWithMessage("ID不能为空", c)
		return
	}
	if category.Name == "" {
		response.FailWithMessage("类别名称不能为空", c)
		return
	}

	err = navGameCategoryService.UpdateGameCategory(category)
	if err != nil {
		global.GVA_LOG.Error("更新失败!", zap.Error(err))
		response.FailWithMessage(err.Error(), c)
	} else {
		response.OkWithMessage("更新成功", c)
	}
}

// DeleteGameCategory 删除游戏类别
// @Tags     GameCategory
// @Summary  删除游戏类别
// @Security ApiKeyAuth
// @accept   application/json
// @Produce  application/json
// @Param    data body navRequest.GetById true "游戏类别ID"
// @Success  200  {object} response.Response{msg=string} "删除成功"
// @Router   /navigation/gameCategory/deleteGameCategory [delete]
func (a *NavGameCategoryApi) DeleteGameCategory(c *gin.Context) {
	var req navRequest.GetById
	err := c.ShouldBindJSON(&req)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}

	err = navGameCategoryService.DeleteGameCategory(req.ID)
	if err != nil {
		global.GVA_LOG.Error("删除失败!", zap.Error(err))
		response.FailWithMessage(err.Error(), c)
	} else {
		response.OkWithMessage("删除成功", c)
	}
}

// GetGameCategoryList 获取游戏类别列表
// @Tags     GameCategory
// @Summary  获取游戏类别列表
// @Security ApiKeyAuth
// @accept   application/json
// @Produce  application/json
// @Param    data body navRequest.NavGameCategorySearch true "分页获取游戏类别列表"
// @Success  200  {object} response.Response{data=response.PageResult{list=[]navigation.NavGameCategory,total=int64}} "获取成功"
// @Router   /navigation/gameCategory/getGameCategoryList [post]
func (a *NavGameCategoryApi) GetGameCategoryList(c *gin.Context) {
	var pageInfo navRequest.NavGameCategorySearch
	err := c.ShouldBindJSON(&pageInfo)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}

	list, total, err := navGameCategoryService.GetGameCategoryList(pageInfo)
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

// GetAllGameCategories 获取所有启用的游戏类别
// @Tags     GameCategory
// @Summary  获取所有启用的游戏类别
// @Security ApiKeyAuth
// @accept   application/json
// @Produce  application/json
// @Success  200  {object} response.Response{data=[]navigation.NavGameCategory} "获取成功"
// @Router   /navigation/gameCategory/getAllGameCategories [get]
func (a *NavGameCategoryApi) GetAllGameCategories(c *gin.Context) {
	list, err := navGameCategoryService.GetAllGameCategories()
	if err != nil {
		global.GVA_LOG.Error("获取失败!", zap.Error(err))
		response.FailWithMessage(err.Error(), c)
	} else {
		response.OkWithDetailed(list, "获取成功", c)
	}
}
