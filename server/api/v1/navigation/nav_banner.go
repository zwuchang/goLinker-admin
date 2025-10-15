package navigation

import (
	"goLinker-admin/server/global"
	"goLinker-admin/server/model/common/response"
	"goLinker-admin/server/model/navigation/request"
	"strconv"

	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
)

type NavBannerApi struct{}

// GetBannerList 获取Banner列表
// @Tags     BannerModule
// @Summary  获取Banner列表
// @Security ApiKeyAuth
// @accept   application/json
// @Produce  application/json
// @Param    data body request.NavBannerSearch true "Banner搜索参数"
// @Success  200  {object} response.Response{data=response.NavBannerListResponse} "获取成功"
// @Router   /navigation/banner/getBannerList [post]
func (a *NavBannerApi) GetBannerList(c *gin.Context) {
	var req request.NavBannerSearch
	err := c.ShouldBindJSON(&req)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}

	// 设置默认分页参数
	if req.Page == 0 {
		req.Page = 1
	}
	if req.PageSize == 0 {
		req.PageSize = 10
	}

	list, total, err := navBannerService.GetBannerList(req)
	if err != nil {
		global.GVA_LOG.Error("获取Banner列表失败!", zap.Error(err))
		response.FailWithMessage("获取Banner列表失败", c)
		return
	}

	response.OkWithDetailed(map[string]interface{}{
		"list":     list,
		"total":    total,
		"page":     req.Page,
		"pageSize": req.PageSize,
	}, "获取成功", c)
}

// GetBannerById 根据ID获取Banner
// @Tags     BannerModule
// @Summary  根据ID获取Banner
// @Security ApiKeyAuth
// @accept   application/json
// @Produce  application/json
// @Param    id path uint true "Banner ID"
// @Success  200  {object} response.Response{data=response.NavBannerResponse} "获取成功"
// @Router   /navigation/banner/getBannerById/{id} [get]
func (a *NavBannerApi) GetBannerById(c *gin.Context) {
	id, err := strconv.ParseUint(c.Param("id"), 10, 32)
	if err != nil || id == 0 {
		response.FailWithMessage("ID不能为空", c)
		return
	}

	banner, err := navBannerService.GetBannerById(uint(id))
	if err != nil {
		global.GVA_LOG.Error("获取Banner失败!", zap.Error(err))
		response.FailWithMessage(err.Error(), c)
		return
	}

	response.OkWithDetailed(map[string]interface{}{
		"banner": banner,
	}, "获取成功", c)
}

// CreateBanner 创建Banner
// @Tags     BannerModule
// @Summary  创建Banner
// @Security ApiKeyAuth
// @accept   application/json
// @Produce  application/json
// @Param    data body request.NavBannerCreateRequest true "Banner信息"
// @Success  200  {object} response.Response{msg=string} "创建成功"
// @Router   /navigation/banner/createBanner [post]
func (a *NavBannerApi) CreateBanner(c *gin.Context) {
	var req request.NavBannerCreateRequest
	err := c.ShouldBindJSON(&req)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}

	err = navBannerService.CreateBanner(req)
	if err != nil {
		global.GVA_LOG.Error("创建Banner失败!", zap.Error(err))
		response.FailWithMessage(err.Error(), c)
		return
	}

	response.OkWithMessage("创建成功", c)
}

// UpdateBanner 更新Banner
// @Tags     BannerModule
// @Summary  更新Banner
// @Security ApiKeyAuth
// @accept   application/json
// @Produce  application/json
// @Param    data body request.NavBannerUpdateRequest true "Banner信息"
// @Success  200  {object} response.Response{msg=string} "更新成功"
// @Router   /navigation/banner/updateBanner [post]
func (a *NavBannerApi) UpdateBanner(c *gin.Context) {
	var req request.NavBannerUpdateRequest
	err := c.ShouldBindJSON(&req)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}

	err = navBannerService.UpdateBanner(req)
	if err != nil {
		global.GVA_LOG.Error("更新Banner失败!", zap.Error(err))
		response.FailWithMessage(err.Error(), c)
		return
	}

	response.OkWithMessage("更新成功", c)
}

// DeleteBanner 删除Banner
// @Tags     BannerModule
// @Summary  删除Banner
// @Security ApiKeyAuth
// @accept   application/json
// @Produce  application/json
// @Param    data body request.DeleteBannerRequest true "删除Banner请求"
// @Success  200  {object} response.Response{msg=string} "删除成功"
// @Router   /navigation/banner/deleteBanner [post]
func (a *NavBannerApi) DeleteBanner(c *gin.Context) {
	var req request.DeleteBannerRequest
	err := c.ShouldBindJSON(&req)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}

	err = navBannerService.DeleteBanner(req.ID)
	if err != nil {
		global.GVA_LOG.Error("删除Banner失败!", zap.Error(err))
		response.FailWithMessage(err.Error(), c)
		return
	}

	response.OkWithMessage("删除成功", c)
}
