package system

import (
	"goLinker-admin/server/middleware"

	"github.com/gin-gonic/gin"
)

type CronRouter struct{}

func (s *CronRouter) InitCronRouter(Router *gin.RouterGroup) {
	cronRouter := Router.Group("cron").Use(middleware.OperationRecord())
	{
		cronRouter.POST("update", cronApi.UpdateCronSpec) // 更新定时任务规则
		cronRouter.POST("stop", cronApi.StopCron)         // 暂停定时任务
		cronRouter.POST("start", cronApi.StartCron)       // 启动定时任务
	}
	{
		Router.Group("cron").POST("list", cronApi.GetCronStatus) // 获取定时任务状态
	}
}
