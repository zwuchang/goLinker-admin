package navigation

import (
	"goLinker-admin/server/middleware"

	"github.com/gin-gonic/gin"
)

type NavActivityConfigRouter struct{}

// InitNavActivityConfigRouter 初始化活动配置路由
func (s *NavActivityConfigRouter) InitNavActivityConfigRouter(Router *gin.RouterGroup) {
	navActivityConfigRouter := Router.Group("/navigation/activityConfig").Use(middleware.OperationRecord())
	navActivityConfigRouterWithoutRecord := Router.Group("/navigation/activityConfig")
	{
		navActivityConfigRouter.POST("createActivityConfig", navActivityConfigApi.CreateActivityConfig) // 创建活动配置
		navActivityConfigRouter.POST("updateActivityConfig", navActivityConfigApi.UpdateActivityConfig) // 更新活动配置
		navActivityConfigRouter.POST("deleteActivityConfig", navActivityConfigApi.DeleteActivityConfig) // 删除活动配置

	}
	{
		navActivityConfigRouterWithoutRecord.POST("getActivityConfigList", navActivityConfigApi.GetActivityConfigList) // 获取活动配置列表
		navActivityConfigRouterWithoutRecord.POST("getActivityConfigById", navActivityConfigApi.GetActivityConfigById) // 根据ID获取活动配置
	}
}
