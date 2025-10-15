package navigation

import (
	"goLinker-admin/server/middleware"

	"github.com/gin-gonic/gin"
)

type NavAccessStatsRouter struct{}

// InitAccessStatsRouter 初始化访问统计路由
func (s *NavAccessStatsRouter) InitAccessStatsRouter(Router *gin.RouterGroup) {
	accessStatsRouter := Router.Group("navigation/accessStats")
	accessStatsNoLoggedRouter := Router.Group("navigation/accessStats")
	accessStatsRouter.Use(middleware.OperationRecord())
	{
		accessStatsRouter.POST("getAccessStatsList", navAccessStatsApi.GetAccessStatsList)       // 获取访问统计列表
		accessStatsRouter.POST("getAccessStatsSummary", navAccessStatsApi.GetAccessStatsSummary) // 获取访问统计汇总
	}
	{
		accessStatsNoLoggedRouter.POST("cleanOldStats", navAccessStatsApi.CleanOldStats) // 清理旧数据

	}
}
