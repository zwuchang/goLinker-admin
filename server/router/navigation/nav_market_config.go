package navigation

import (
	v1 "goLinker-admin/server/api/v1"
	"goLinker-admin/server/middleware"

	"github.com/gin-gonic/gin"
)

type NavMarketConfigRouter struct{}

// InitNavMarketConfigRouter 初始化 市场配置 路由信息
func (s *NavMarketConfigRouter) InitNavMarketConfigRouter(Router *gin.RouterGroup) {
	navMarketConfigRouter := Router.Group("navigation/marketConfig").Use(middleware.OperationRecord())
	navMarketConfigRouterWithoutRecord := Router.Group("navigation/marketConfig")
	var navMarketConfigApi = v1.ApiGroupApp.NavigationApiGroup.NavMarketConfigApi
	{
		navMarketConfigRouter.POST("createMarketConfig", navMarketConfigApi.CreateMarketConfig) // 创建市场配置
		navMarketConfigRouter.POST("updateMarketConfig", navMarketConfigApi.UpdateMarketConfig) // 更新市场配置
		navMarketConfigRouter.POST("deleteMarketConfig", navMarketConfigApi.DeleteMarketConfig) // 删除市场配置
	}
	{
		navMarketConfigRouterWithoutRecord.POST("getMarketConfigList", navMarketConfigApi.GetMarketConfigList)    // 获取市场配置列表
		navMarketConfigRouterWithoutRecord.GET("getMarketConfigById/:id", navMarketConfigApi.GetMarketConfigById) // 根据ID获取市场配置
	}
}


