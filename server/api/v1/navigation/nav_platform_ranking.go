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

type NavPlatformRankingApi struct{}

var navPlatformRankingApiService = service.ServiceGroupApp.NavigationServiceGroup.NavPlatformRankingService

// GetPlatformRankingList 获取平台排行榜列表
// @Tags     NavPlatformRanking
// @Summary  获取平台排行榜列表
// @Security ApiKeyAuth
// @accept   application/json
// @Produce  application/json
// @Param    data body navRequest.NavPlatformRankingSearch true "分页获取平台排行榜列表"
// @Success  200  {object} response.Response{data=navResponse.NavPlatformRankingListResponse,msg=string} "获取平台排行榜列表成功"
// @Router   /platformRanking/getPlatformRankingList [post]
func (a *NavPlatformRankingApi) GetPlatformRankingList(c *gin.Context) {
	var pageInfo navRequest.NavPlatformRankingSearch
	err := c.ShouldBindJSON(&pageInfo)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	if list, total, err := navPlatformRankingApiService.GetPlatformRankingList(pageInfo); err != nil {
		global.GVA_LOG.Error("获取失败!", zap.Error(err))
		response.FailWithMessage("获取失败", c)
	} else {
		response.OkWithDetailed(navResponse.NavPlatformRankingListResponse{
			List:     list,
			Total:    total,
			Page:     pageInfo.Page,
			PageSize: pageInfo.PageSize,
		}, "获取成功", c)
	}
}

// CreatePlatformRanking 创建平台排行榜
// @Tags     NavPlatformRanking
// @Summary  创建平台排行榜
// @Security ApiKeyAuth
// @accept   application/json
// @Produce  application/json
// @Param    data body navRequest.CreateNavPlatformRankingRequest true "创建平台排行榜"
// @Success  200  {object} response.Response{msg=string} "创建成功"
// @Router   /platformRanking/createPlatformRanking [post]
func (a *NavPlatformRankingApi) CreatePlatformRanking(c *gin.Context) {
	var platformRanking navRequest.CreateNavPlatformRankingRequest
	err := c.ShouldBindJSON(&platformRanking)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	if err := navPlatformRankingApiService.CreatePlatformRanking(platformRanking); err != nil {
		global.GVA_LOG.Error("创建失败!", zap.Error(err))
		response.FailWithMessage("创建失败", c)
	} else {
		response.OkWithMessage("创建成功", c)
	}
}

// UpdatePlatformRanking 更新平台排行榜
// @Tags     NavPlatformRanking
// @Summary  更新平台排行榜
// @Security ApiKeyAuth
// @accept   application/json
// @Produce  application/json
// @Param    data body navRequest.UpdateNavPlatformRankingRequest true "更新平台排行榜"
// @Success  200  {object} response.Response{msg=string} "更新成功"
// @Router   /platformRanking/updatePlatformRanking [post]
func (a *NavPlatformRankingApi) UpdatePlatformRanking(c *gin.Context) {
	var platformRanking navRequest.UpdateNavPlatformRankingRequest
	err := c.ShouldBindJSON(&platformRanking)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	if err := navPlatformRankingApiService.UpdatePlatformRanking(platformRanking); err != nil {
		global.GVA_LOG.Error("更新失败!", zap.Error(err))
		response.FailWithMessage("更新失败", c)
	} else {
		response.OkWithMessage("更新成功", c)
	}
}

// DeletePlatformRanking 删除平台排行榜
// @Tags     NavPlatformRanking
// @Summary  删除平台排行榜
// @Security ApiKeyAuth
// @accept   application/json
// @Produce  application/json
// @Param    data body request.GetById true "删除平台排行榜"
// @Success  200  {object} response.Response{msg=string} "删除成功"
// @Router   /platformRanking/deletePlatformRanking [post]
func (a *NavPlatformRankingApi) DeletePlatformRanking(c *gin.Context) {
	var platformRanking request.GetById
	err := c.ShouldBindJSON(&platformRanking)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	if err := navPlatformRankingApiService.DeletePlatformRanking(uint(platformRanking.ID)); err != nil {
		global.GVA_LOG.Error("删除失败!", zap.Error(err))
		response.FailWithMessage("删除失败", c)
	} else {
		response.OkWithMessage("删除成功", c)
	}
}

// GetPlatformRankingById 根据ID获取平台排行榜
// @Tags     NavPlatformRanking
// @Summary  根据ID获取平台排行榜
// @Security ApiKeyAuth
// @accept   application/json
// @Produce  application/json
// @Param    data body request.GetById true "根据ID获取平台排行榜"
// @Success  200  {object} response.Response{data=navResponse.NavPlatformRankingResponse,msg=string} "获取成功"
// @Router   /platformRanking/getPlatformRankingById [post]
func (a *NavPlatformRankingApi) GetPlatformRankingById(c *gin.Context) {
	var platformRanking request.GetById
	err := c.ShouldBindJSON(&platformRanking)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	if rePlatformRanking, err := navPlatformRankingApiService.GetPlatformRankingById(uint(platformRanking.ID)); err != nil {
		global.GVA_LOG.Error("获取失败!", zap.Error(err))
		response.FailWithMessage("获取失败", c)
	} else {
		response.OkWithDetailed(navResponse.NavPlatformRankingResponse{
			Ranking: rePlatformRanking,
		}, "获取成功", c)
	}
}
