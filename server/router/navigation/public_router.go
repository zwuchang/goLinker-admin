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
		publicRouter.POST("site/info", publicApi.GetSiteInfo)                 // 获取网站信息
		publicRouter.POST("banner/list", publicApi.GetBannerList)             // 获取Banner列表
		publicRouter.POST("game/list", publicApi.GetGameList)                 // 获取游戏列表
		publicRouter.POST("gameCategory/list", publicApi.GetGameCategoryList) // 获取游戏类别列表
	}
}
