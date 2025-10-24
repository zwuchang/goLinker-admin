package navigation

import (
	"goLinker-admin/server/middleware"

	"github.com/gin-gonic/gin"
)

type NavPWAConfigRouter struct{}

// InitNavPWAConfigRouter 初始化PWA配置路由信息
func (s *NavPWAConfigRouter) InitNavPWAConfigRouter(Router *gin.RouterGroup) {
	navPWAConfigRouter := Router.Group("navigation/pwaConfig").Use(middleware.OperationRecord())
	navPWAConfigRouterWithoutRecord := Router.Group("navigation/pwaConfig")
	{
		navPWAConfigRouter.POST("createPWAConfig", navPWAConfigApi.CreatePWAConfig) // 新建PWA配置
		navPWAConfigRouter.POST("deletePWAConfig", navPWAConfigApi.DeletePWAConfig) // 删除PWA配置
		navPWAConfigRouter.POST("updatePWAConfig", navPWAConfigApi.UpdatePWAConfig) // 更新PWA配置
		navPWAConfigRouter.POST("clearPWACache", navPWAConfigApi.ClearPWACache)     // 清除PWA配置缓存
	}
	{
		navPWAConfigRouterWithoutRecord.POST("findPWAConfig", navPWAConfigApi.GetPWAConfigById)    // 根据ID获取PWA配置
		navPWAConfigRouterWithoutRecord.POST("getPWAConfigList", navPWAConfigApi.GetPWAConfigList) // 获取PWA配置列表
	}
}
