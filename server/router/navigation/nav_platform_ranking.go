package navigation

import (
	"goLinker-admin/server/middleware"

	"github.com/gin-gonic/gin"
)

type NavPlatformRankingRouter struct{}

// InitNavPlatformRankingRouter 初始化平台排行榜路由信息
func (s *NavPlatformRankingRouter) InitNavPlatformRankingRouter(Router *gin.RouterGroup) {
	navPlatformRankingRouter := Router.Group("navigation/platformRanking").Use(middleware.OperationRecord())
	navPlatformRankingRouterWithoutRecord := Router.Group("navigation/platformRanking")
	{
		navPlatformRankingRouter.POST("createPlatformRanking", navPlatformRankingApi.CreatePlatformRanking) // 新建平台排行榜
		navPlatformRankingRouter.POST("deletePlatformRanking", navPlatformRankingApi.DeletePlatformRanking) // 删除平台排行榜
		navPlatformRankingRouter.POST("updatePlatformRanking", navPlatformRankingApi.UpdatePlatformRanking) // 更新平台排行榜
	}
	{
		navPlatformRankingRouterWithoutRecord.POST("findPlatformRanking", navPlatformRankingApi.GetPlatformRankingById)    // 根据ID获取平台排行榜
		navPlatformRankingRouterWithoutRecord.POST("getPlatformRankingList", navPlatformRankingApi.GetPlatformRankingList) // 获取平台排行榜列表
	}
}
