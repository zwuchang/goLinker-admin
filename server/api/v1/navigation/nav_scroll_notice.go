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

type NavScrollNoticeApi struct{}

var navScrollNoticeApiService = service.ServiceGroupApp.NavigationServiceGroup.NavScrollNoticeService

// GetScrollNoticeList 获取滚动通知列表
// @Tags     NavScrollNotice
// @Summary  获取滚动通知列表
// @Security ApiKeyAuth
// @accept   application/json
// @Produce  application/json
// @Param    data body navRequest.NavScrollNoticeSearch true "分页获取滚动通知列表"
// @Success  200  {object} response.Response{data=navResponse.NavScrollNoticeListResponse,msg=string} "获取滚动通知列表成功"
// @Router   /scrollNotice/getScrollNoticeList [post]
func (a *NavScrollNoticeApi) GetScrollNoticeList(c *gin.Context) {
	var pageInfo navRequest.NavScrollNoticeSearch
	err := c.ShouldBindJSON(&pageInfo)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	if list, total, err := navScrollNoticeApiService.GetScrollNoticeList(pageInfo); err != nil {
		global.GVA_LOG.Error("获取失败!", zap.Error(err))
		response.FailWithMessage("获取失败", c)
	} else {
		response.OkWithDetailed(navResponse.NavScrollNoticeListResponse{
			List:     list,
			Total:    total,
			Page:     pageInfo.Page,
			PageSize: pageInfo.PageSize,
		}, "获取成功", c)
	}
}

// CreateScrollNotice 创建滚动通知
// @Tags     NavScrollNotice
// @Summary  创建滚动通知
// @Security ApiKeyAuth
// @accept   application/json
// @Produce  application/json
// @Param    data body navRequest.CreateNavScrollNoticeRequest true "创建滚动通知"
// @Success  200  {object} response.Response{msg=string} "创建成功"
// @Router   /scrollNotice/createScrollNotice [post]
func (a *NavScrollNoticeApi) CreateScrollNotice(c *gin.Context) {
	var scrollNotice navRequest.CreateNavScrollNoticeRequest
	err := c.ShouldBindJSON(&scrollNotice)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	if err := navScrollNoticeApiService.CreateScrollNotice(scrollNotice); err != nil {
		global.GVA_LOG.Error("创建失败!", zap.Error(err))
		response.FailWithMessage("创建失败", c)
	} else {
		response.OkWithMessage("创建成功", c)
	}
}

// UpdateScrollNotice 更新滚动通知
// @Tags     NavScrollNotice
// @Summary  更新滚动通知
// @Security ApiKeyAuth
// @accept   application/json
// @Produce  application/json
// @Param    data body navRequest.UpdateNavScrollNoticeRequest true "更新滚动通知"
// @Success  200  {object} response.Response{msg=string} "更新成功"
// @Router   /scrollNotice/updateScrollNotice [post]
func (a *NavScrollNoticeApi) UpdateScrollNotice(c *gin.Context) {
	var scrollNotice navRequest.UpdateNavScrollNoticeRequest
	err := c.ShouldBindJSON(&scrollNotice)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	if err := navScrollNoticeApiService.UpdateScrollNotice(scrollNotice); err != nil {
		global.GVA_LOG.Error("更新失败!", zap.Error(err))
		response.FailWithMessage("更新失败", c)
	} else {
		response.OkWithMessage("更新成功", c)
	}
}

// DeleteScrollNotice 删除滚动通知
// @Tags     NavScrollNotice
// @Summary  删除滚动通知
// @Security ApiKeyAuth
// @accept   application/json
// @Produce  application/json
// @Param    data body request.GetById true "删除滚动通知"
// @Success  200  {object} response.Response{msg=string} "删除成功"
// @Router   /scrollNotice/deleteScrollNotice [post]
func (a *NavScrollNoticeApi) DeleteScrollNotice(c *gin.Context) {
	var scrollNotice request.GetById
	err := c.ShouldBindJSON(&scrollNotice)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	if err := navScrollNoticeApiService.DeleteScrollNotice(uint(scrollNotice.ID)); err != nil {
		global.GVA_LOG.Error("删除失败!", zap.Error(err))
		response.FailWithMessage("删除失败", c)
	} else {
		response.OkWithMessage("删除成功", c)
	}
}

// GetScrollNoticeById 根据ID获取滚动通知
// @Tags     NavScrollNotice
// @Summary  根据ID获取滚动通知
// @Security ApiKeyAuth
// @accept   application/json
// @Produce  application/json
// @Param    data body request.GetById true "根据ID获取滚动通知"
// @Success  200  {object} response.Response{data=navResponse.NavScrollNoticeResponse,msg=string} "获取成功"
// @Router   /scrollNotice/getScrollNoticeById [post]
func (a *NavScrollNoticeApi) GetScrollNoticeById(c *gin.Context) {
	var scrollNotice request.GetById
	err := c.ShouldBindJSON(&scrollNotice)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	if reScrollNotice, err := navScrollNoticeApiService.GetScrollNoticeById(uint(scrollNotice.ID)); err != nil {
		global.GVA_LOG.Error("获取失败!", zap.Error(err))
		response.FailWithMessage("获取失败", c)
	} else {
		response.OkWithDetailed(navResponse.NavScrollNoticeResponse{
			Notice: reScrollNotice,
		}, "获取成功", c)
	}
}
