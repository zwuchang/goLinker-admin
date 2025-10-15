package navigation

import (
	"goLinker-admin/server/global"
	commonRequest "goLinker-admin/server/model/common/request"
	"goLinker-admin/server/model/common/response"
	"goLinker-admin/server/model/navigation"
	navRequest "goLinker-admin/server/model/navigation/request"
	navResponse "goLinker-admin/server/model/navigation/response"
	"goLinker-admin/server/service"

	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
)

type PublicApi struct{}

// GetSiteInfo 获取网站信息（公开接口，无需认证）
// @Tags     PublicApi
// @Summary  获取网站信息
// @accept   application/json
// @Produce  application/json
// @Success  200  {object} response.Response{data=navResponse.PublicSiteInfoResponse} "获取成功"
// @Router   /public/site/info [post]
func (a *PublicApi) GetSiteInfo(c *gin.Context) {
	// 获取最新的游戏配置作为网站信息
	gameConfigs, _, err := service.ServiceGroupApp.NavigationServiceGroup.NavGameConfigService.GetGameConfigList(navRequest.NavGameConfigSearch{
		PageInfo: commonRequest.PageInfo{Page: 1, PageSize: 1},
		NavGameConfig: navigation.NavGameConfig{
			Status: 1, // 只获取启用的配置
		},
	})
	if err != nil || len(gameConfigs) == 0 {
		global.GVA_LOG.Error("获取网站信息失败!", zap.Error(err))
		response.FailWithMessage("获取网站信息失败", c)
		return
	}

	config := gameConfigs[0]
	siteInfo := navResponse.PublicSiteInfoResponse{
		WebsiteTitle: config.WebsiteTitle,
		WebsiteDesc:  config.WebsiteDesc,
		DownloadUrl:  config.DownloadUrl,
		AudioUrl:     config.AudioUrl,
		WebsiteIcon:  config.WebsiteIcon,
		WebsiteLogo:  config.WebsiteLogo,
		UpdateTime:   config.UpdatedAt.Format("2006-01-02 15:04:05"),
	}

	response.OkWithDetailed(siteInfo, "获取成功", c)
}

// GetBannerList 获取Banner列表（公开接口，无需认证）
// @Tags     PublicApi
// @Summary  获取Banner列表
// @accept   application/json
// @Produce  application/json
// @Param    data body navRequest.NavBannerSearch true "Banner搜索参数"
// @Success  200  {object} response.Response{data=map[string]interface{}} "获取成功"
// @Router   /public/banner/list [post]
func (a *PublicApi) GetBannerList(c *gin.Context) {
	var req navRequest.NavBannerSearch
	err := c.ShouldBindJSON(&req)
	if err != nil {
		// 如果JSON绑定失败，使用默认值
		req.Page = 1
		req.PageSize = 10
	}

	// 设置默认分页参数
	if req.Page <= 0 {
		req.Page = 1
	}
	if req.PageSize <= 0 {
		req.PageSize = 10
	}

	// 只获取可见的Banner
	visible := 1
	status := 1
	req.IsVisible = &visible
	req.Status = &status

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

// GetGameList 获取游戏列表（公开接口，无需认证）
// @Tags     PublicApi
// @Summary  获取游戏列表
// @accept   application/json
// @Produce  application/json
// @Param    data body navRequest.NavGameSearch true "游戏搜索参数"
// @Success  200  {object} response.Response{data=map[string]interface{}} "获取成功"
// @Router   /public/game/list [post]
func (a *PublicApi) GetGameList(c *gin.Context) {
	var req navRequest.NavGameSearch
	err := c.ShouldBindJSON(&req)
	if err != nil {
		// 如果JSON绑定失败，使用默认值
		req.Page = 1
		req.PageSize = 10
	}

	// 设置默认分页参数
	if req.Page <= 0 {
		req.Page = 1
	}
	if req.PageSize <= 0 {
		req.PageSize = 10
	}

	// 只获取可见的游戏
	req.IsVisible = 1
	req.Status = 1

	list, total, err := navGameService.GetGameList(req)
	if err != nil {
		global.GVA_LOG.Error("获取游戏列表失败!", zap.Error(err))
		response.FailWithMessage("获取游戏列表失败", c)
		return
	}

	response.OkWithDetailed(map[string]interface{}{
		"list":     list,
		"total":    total,
		"page":     req.Page,
		"pageSize": req.PageSize,
	}, "获取成功", c)
}

// GetGameCategoryList 获取游戏类别列表（公开接口，无需认证）
// @Tags     PublicApi
// @Summary  获取游戏类别列表
// @accept   application/json
// @Produce  application/json
// @Param    data body navRequest.NavGameCategorySearch true "游戏类别搜索参数"
// @Success  200  {object} response.Response{data=map[string]interface{}} "获取成功"
// @Router   /public/gameCategory/list [post]
func (a *PublicApi) GetGameCategoryList(c *gin.Context) {
	var req navRequest.NavGameCategorySearch
	err := c.ShouldBindJSON(&req)
	if err != nil {
		// 如果JSON绑定失败，使用默认值
		req.Page = 1
		req.PageSize = 10
	}

	if req.Page <= 0 {
		req.Page = 1
	}
	if req.PageSize <= 0 {
		req.PageSize = 10
	}

	// 只获取启用的游戏类别
	req.Status = 1

	list, total, err := navGameCategoryService.GetGameCategoryList(req)
	if err != nil {
		global.GVA_LOG.Error("获取游戏类别列表失败!", zap.Error(err))
		response.FailWithMessage("获取游戏类别列表失败", c)
		return
	}

	// 只返回ID和名称
	type CategoryInfo struct {
		ID   uint   `json:"id"`
		Name string `json:"name"`
	}

	var categoryList []CategoryInfo
	for _, category := range list {
		categoryList = append(categoryList, CategoryInfo{
			ID:   category.ID,
			Name: category.Name,
		})
	}

	response.OkWithDetailed(map[string]interface{}{
		"list":     categoryList,
		"total":    total,
		"page":     req.Page,
		"pageSize": req.PageSize,
	}, "获取成功", c)
}

// GetContactInfo 获取联系方式信息（公开接口，无需认证）
// @Tags     PublicApi
// @Summary  获取联系方式信息
// @accept   application/json
// @Produce  application/json
// @Success  200  {object} response.Response{data=navResponse.PublicContactResponse} "获取成功"
// @Router   /public/contact/info [post]
func (a *PublicApi) GetContactInfo(c *gin.Context) {
	// 获取启用的联系方式列表
	contactMethods, err := contactMethodService.GetEnabledContactMethods()
	if err != nil {
		global.GVA_LOG.Error("获取联系方式列表失败!", zap.Error(err))
		response.FailWithMessage("获取联系方式列表失败", c)
		return
	}

	// 获取联系配置
	contactConfig, err := contactConfigService.GetNavContactConfig()
	if err != nil {
		global.GVA_LOG.Error("获取联系配置失败!", zap.Error(err))
		response.FailWithMessage("获取联系配置失败", c)
		return
	}

	// 转换联系方式数据
	var contactMethodList []navResponse.ContactMethodInfo
	for _, method := range contactMethods {
		contactMethodList = append(contactMethodList, navResponse.ContactMethodInfo{
			ID:          method.ID,
			Image:       method.Image,
			JumpUrl:     method.JumpUrl,
			DisplayName: method.DisplayName,
			Sort:        method.Sort,
		})
	}

	// 构建响应数据
	contactInfo := navResponse.PublicContactResponse{
		ContactMethods: contactMethodList,
		ContactConfig: navResponse.ContactConfigInfo{
			BannerImage: contactConfig.BannerImage,
			Email:       contactConfig.Email,
		},
		UpdateTime: contactConfig.UpdatedAt.Format("2006-01-02 15:04:05"),
	}

	response.OkWithDetailed(contactInfo, "获取成功", c)
}
