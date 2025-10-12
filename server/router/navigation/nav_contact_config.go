package navigation

import (
	"goLinker-admin/server/middleware"

	"github.com/gin-gonic/gin"
)

type ContactConfigRouter struct{}

func (s *ContactConfigRouter) InitContactConfigRouter(Router *gin.RouterGroup) {
	navContactConfigRouter := Router.Group("navigation/contactConfig").Use(middleware.OperationRecord())
	navContactConfigRouterWithoutRecord := Router.Group("navigation/contactConfig")
	{
		navContactConfigRouter.PUT("update", navContactConfigApi.UpdateNavContactConfig) // 更新联系配置
	}
	{
		navContactConfigRouterWithoutRecord.GET("get", navContactConfigApi.GetNavContactConfig) // 获取联系配置
	}
}
