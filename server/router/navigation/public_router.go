package navigation

import (
	"goLinker-admin/server/middleware"

	"github.com/gin-gonic/gin"
)

type PublicRouter struct{}

// InitPublicRouter 初始化公开接口路由
func (s *PublicRouter) InitPublicRouter(Router *gin.RouterGroup) {
	publicRouter := Router.Group("v1")
	publicRouter.Use(middleware.AccessStatsMiddleware()) // 使用访问统计中间件
	{
		// 网站信息相关接口（无需认证）
		publicRouter.POST("site/info", publicApi.GetSiteInfo)                // 获取网站信息
		publicRouter.POST("banner/list", publicApi.GetBannerList)            // 获取Banner列表
		publicRouter.POST("game/list", publicApi.GetGameList)                // 获取游戏列表
		publicRouter.POST("category/list", publicApi.GetGameCategoryList)    // 获取游戏类别列表
		publicRouter.POST("contact/info", publicApi.GetContactInfo)          // 获取联系方式信息
		publicRouter.POST("game/article", publicApi.GetGameArticle)          // 根据游戏ID获取文章内容
		publicRouter.POST("ads/index", publicApi.GetAdsList)                 // 获取广告列表（置顶游戏）
		publicRouter.POST("platform/menus", publicApi.GetMenus)              // 获取平台菜单列表
		publicRouter.POST("platform/getGame", publicApi.GetGame)             // 获取对应游戏
		publicRouter.POST("game/getMarket", publicApi.GetMarket)             // 获取对应市场
		publicRouter.POST("ranking/platforms", publicApi.GetPlatformRanking) // 获取平台排行榜
	}
}
