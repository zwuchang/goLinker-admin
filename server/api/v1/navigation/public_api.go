package navigation

import (
	"encoding/json"
	"goLinker-admin/server/global"
	commonRequest "goLinker-admin/server/model/common/request"
	"goLinker-admin/server/model/common/response"
	"goLinker-admin/server/model/navigation"
	navRequest "goLinker-admin/server/model/navigation/request"
	navResponse "goLinker-admin/server/model/navigation/response"
	"goLinker-admin/server/service"
	"math/rand/v2"
	"time"

	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
)

type PublicApi struct{}

// GetSiteInfo 获取网站信息（公开接口，无需认证）
// @Tags     PublicApi
// @Summary  Get website information
// @accept   application/json
// @Produce  application/json
// @Success  200  {object} response.Response{data=navResponse.PublicSiteInfoResponse} "Success"
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
		response.FailWithMessage("Failed to get website information", c)
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

	response.OkWithDetailed(siteInfo, "Success", c)
}

// GetBannerList 获取Banner列表（公开接口，无需认证）
// @Tags     PublicApi
// @Summary  Get banner list
// @accept   application/json
// @Produce  application/json
// @Param    data body navRequest.NavBannerSearch true "Banner search parameters"
// @Success  200  {object} response.Response{data=map[string]interface{}} "Success"
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
		response.FailWithMessage("Failed to get banner list", c)
		return
	}

	response.OkWithDetailed(map[string]interface{}{
		"list":     list,
		"total":    total,
		"page":     req.Page,
		"pageSize": req.PageSize,
	}, "Success", c)
}

// GetGameList 获取游戏列表（公开接口，无需认证）
// @Tags     PublicApi
// @Summary  Get game list
// @accept   application/json
// @Produce  application/json
// @Param    data body navRequest.NavGameSearch true "Game search parameters"
// @Success  200  {object} response.Response{data=map[string]interface{}} "Success"
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
		response.FailWithMessage("Failed to get game list", c)
		return
	}

	// 随机打乱数据
	rand.Shuffle(len(list), func(i, j int) {
		list[i], list[j] = list[j], list[i]
	})

	var itemList []navResponse.PublicGameItemResponse
	for _, game := range list {
		itemList = append(itemList, navResponse.PublicGameItemResponse{
			ID:        game.ID,
			Title:     game.Title,
			Image:     game.Icon,
			Views:     game.Views,
			CreatedAt: game.CreatedAt.Format("2006-01-02 15:04:05"),
		})
	}

	// 转换为公开游戏列表响应
	publicGameList := navResponse.PublicGameListResponse{
		List:     itemList,
		Total:    total,
		Page:     req.Page,
		PageSize: req.PageSize,
	}

	response.OkWithDetailed(publicGameList, "Success", c)
}

// GetGameCategoryList 获取游戏类别列表（公开接口，无需认证）
// @Tags     PublicApi
// @Summary  Get game category list
// @accept   application/json
// @Produce  application/json
// @Param    data body navRequest.NavGameCategorySearch true "Game category search parameters"
// @Success  200  {object} response.Response{data=map[string]interface{}} "Success"
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
		response.FailWithMessage("Failed to get game category list", c)
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
	}, "Success", c)
}

// GetContactInfo 获取联系方式信息（公开接口，无需认证）
// @Tags     PublicApi
// @Summary  Get contact information
// @accept   application/json
// @Produce  application/json
// @Success  200  {object} response.Response{data=navResponse.PublicContactResponse} "Success"
// @Router   /public/contact/info [post]
func (a *PublicApi) GetContactInfo(c *gin.Context) {
	// 获取启用的联系方式列表
	contactMethods, err := contactMethodService.GetEnabledContactMethods()
	if err != nil {
		global.GVA_LOG.Error("获取联系方式列表失败!", zap.Error(err))
		response.FailWithMessage("Failed to get contact information", c)
		return
	}

	// 获取联系配置
	contactConfig, err := contactConfigService.GetNavContactConfig()
	if err != nil {
		global.GVA_LOG.Error("获取联系配置失败!", zap.Error(err))
		response.FailWithMessage("Failed to get contact information", c)
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

	response.OkWithDetailed(contactInfo, "Success", c)
}

// GetGameArticle 根据游戏ID获取文章内容（公开接口，无需认证）
// @Tags     PublicApi
// @Summary  Get game article by game ID
// @accept   application/json
// @Produce  application/json
// @Param    data body navRequest.NavGameArticleRequest true "Game article request parameters"
// @Success  200  {object} response.Response{data=navResponse.NavGameArticleResponse} "Success"
// @Router   /public/game/article [post]
func (a *PublicApi) GetGameArticle(c *gin.Context) {
	var req navRequest.NavGameArticleRequest
	err := c.ShouldBindJSON(&req)
	if err != nil {
		response.FailWithMessage("Invalid request parameters", c)
		return
	}

	// 获取游戏信息
	game, err := navGameService.GetGameById(req.GameId)
	if err != nil {
		global.GVA_LOG.Error("获取游戏信息失败!", zap.Error(err))
		response.FailWithMessage("Game not found", c)
		return
	}

	// 检查游戏是否可见
	if game.IsVisible != 1 || game.Status != 1 {
		response.FailWithMessage("Game not available", c)
		return
	}

	// 构建响应数据
	articleInfo := navResponse.NavGameArticleResponse{
		GameId:     game.ID,
		Title:      game.Title,
		Article:    game.Article,
		UpdateTime: game.UpdatedAt.Format("2006-01-02 15:04:05"),
	}

	// 增加浏览次数
	navGameService.UpdateGameViews(game.ID)

	// 使用json.NewEncoder并关闭HTML转义
	c.Header("Content-Type", "application/json; charset=utf-8")
	c.Status(200)

	encoder := json.NewEncoder(c.Writer)
	encoder.SetEscapeHTML(false) // 关闭HTML转义

	responseData := response.Response{
		Code: 0,
		Data: articleInfo,
		Msg:  "Success",
	}

	encoder.Encode(responseData)
}

// GetAdsList 获取广告列表（置顶游戏）
// @Tags     PublicApi
// @Summary  Get ads list (sticky games)
// @accept   application/json
// @Produce  application/json
// @Param    data body navRequest.NavAdsSearch true "Ads search parameters"
// @Success  200  {object} response.Response{data=navResponse.NavAdsListResponse} "Success"
// @Router   /public/ads/index [post]
func (a *PublicApi) GetAdsList(c *gin.Context) {
	var req navRequest.NavAdsSearch
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

	// 固定取24条数据
	req.Page = 1
	req.PageSize = 24

	// 获取置顶游戏列表
	list, total, err := navGameService.GetStickyGames(req)
	if err != nil {
		global.GVA_LOG.Error("获取广告列表失败!", zap.Error(err))
		response.FailWithMessage("Failed to get ads list", c)
		return
	}

	// 随机打乱数据
	rand.Shuffle(len(list), func(i, j int) {
		list[i], list[j] = list[j], list[i]
	})

	// 转换为广告响应格式
	var adsList []navResponse.NavAdsItemResponse
	for _, game := range list {
		adsList = append(adsList, navResponse.NavAdsItemResponse{
			ID:          game.ID,
			AdName:      game.AdName,
			Image:       game.Icon,
			RedirectUrl: game.JumpUrl,
		})
	}

	// 构建响应数据
	adsResponse := navResponse.NavAdsListResponse{
		List:     adsList,
		Total:    total,
		Page:     req.Page,
		PageSize: req.PageSize,
	}

	response.OkWithDetailed(adsResponse, "Success", c)
}

// GetMenus 获取平台菜单列表（公开接口，无需认证）
// @Tags     PublicApi
// @Summary  Get platform menus list
// @accept   application/json
// @Produce  application/json
// @Success  200  {object} response.Response{data=[]navResponse.PublicPlatformMenuResponse} "Success"
// @Router   /public/menus [post]
func (a *PublicApi) GetMenus(c *gin.Context) {
	// 获取可见的平台配置列表
	platformConfigs, err := navPlatformConfigService.GetVisiblePlatformConfigs()
	if err != nil {
		global.GVA_LOG.Error("获取平台菜单列表失败!", zap.Error(err))
		response.FailWithMessage("Failed to get platform menus", c)
		return
	}

	// 转换为公开平台菜单响应格式
	var menuList []navResponse.PublicPlatformMenuResponse
	for _, config := range platformConfigs {
		menuList = append(menuList, navResponse.PublicPlatformMenuResponse{
			ID:           config.ID,
			PlatformName: config.PlatformName,
			PlatformIcon: config.PlatformIcon,
		})
	}

	response.OkWithDetailed(menuList, "Success", c)
}

// GetGame 获取对应游戏（公开接口，无需认证）
// @Tags     PublicApi
// @Summary  Get game by ID
// @accept   application/json
// @Produce  application/json
// @Param    data body navRequest.NavGameRequest true "Game request parameters"
// @Success  200  {object} response.Response{data=navResponse.NavGameResponse} "Success"
// @Router   /public/platform/getGame [post]
func (a *PublicApi) GetGame(c *gin.Context) {

	time.Sleep(1 * time.Second)
	response.OkWithDetailed(nil, "Success", c)

}

// GetMarket 根据类型获取市场配置列表（公开接口，无需认证）
// @Tags     PublicApi
// @Summary  Get market config list by type
// @accept   application/json
// @Produce  application/json
// @Param    data body navRequest.GetMarketByTypeRequest true "Market type request parameters"
// @Success  200  {object} response.Response{data=navResponse.PublicMarketListResponse} "Success"
// @Router   /public/game/getMarket [post]
func (a *PublicApi) GetMarket(c *gin.Context) {
	var req navRequest.GetMarketByTypeRequest
	err := c.ShouldBindJSON(&req)
	if err != nil {
		response.FailWithMessage("Invalid request parameters", c)
		return
	}

	// 获取指定类型的可见市场配置
	marketConfigs, err := navMarketConfigService.GetMarketConfigsByType(req.Type, 10)
	if err != nil {
		global.GVA_LOG.Error("获取市场配置列表失败!", zap.Error(err))
		response.FailWithMessage("Failed to get market config list", c)
		return
	}

	// 随机打乱数据
	rand.Shuffle(len(marketConfigs), func(i, j int) {
		marketConfigs[i], marketConfigs[j] = marketConfigs[j], marketConfigs[i]
	})

	// 转换为公开市场配置响应格式
	var marketList []navResponse.PublicMarketItemResponse
	for _, config := range marketConfigs {
		marketList = append(marketList, navResponse.PublicMarketItemResponse{
			ID:      config.ID,
			Name:    config.Name,
			Logo:    config.Logo,
			JumpUrl: config.JumpUrl,
		})
	}

	// 构建响应数据
	marketResponse := navResponse.PublicMarketListResponse{
		List:     marketList,
		Total:    int64(len(marketList)),
		Page:     1,
		PageSize: len(marketList),
	}

	response.OkWithDetailed(marketResponse, "Success", c)
}
