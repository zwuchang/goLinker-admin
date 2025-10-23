package navigation

import (
	"goLinker-admin/server/middleware"

	"github.com/gin-gonic/gin"
)

type NavPlatformConfigRouter struct{}

// InitNavPlatformConfigRouter 初始化平台游戏配置路由
func (s *NavPlatformConfigRouter) InitNavPlatformConfigRouter(Router *gin.RouterGroup) {
	navPlatformConfigRouter := Router.Group("navigation/platformConfig").Use(middleware.OperationRecord())
	navPlatformConfigRouterWithoutRecord := Router.Group("navigation/platformConfig")
	{
		navPlatformConfigRouter.POST("createPlatformConfig", navPlatformConfigApi.CreatePlatformConfig) // 创建平台游戏配置
		navPlatformConfigRouter.POST("updatePlatformConfig", navPlatformConfigApi.UpdatePlatformConfig) // 更新平台游戏配置
		navPlatformConfigRouter.POST("deletePlatformConfig", navPlatformConfigApi.DeletePlatformConfig) // 删除平台游戏配置
	}
	{
		navPlatformConfigRouterWithoutRecord.POST("getPlatformConfigList", navPlatformConfigApi.GetPlatformConfigList)    // 获取平台游戏配置列表
		navPlatformConfigRouterWithoutRecord.GET("getPlatformConfigById/:id", navPlatformConfigApi.GetPlatformConfigById) // 根据ID获取平台游戏配置
	}
}
