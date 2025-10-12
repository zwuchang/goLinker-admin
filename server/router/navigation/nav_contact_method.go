package navigation

import (
	"goLinker-admin/server/middleware"

	"github.com/gin-gonic/gin"
)

type ContactMethodRouter struct{}

func (s *ContactMethodRouter) InitContactMethodRouter(Router *gin.RouterGroup) {
	navContactMethodRouter := Router.Group("navigation/contactMethod").Use(middleware.OperationRecord())
	navContactMethodRouterWithoutRecord := Router.Group("navigation/contactMethod")
	{
		navContactMethodRouter.POST("create", navContactMethodApi.CreateNavContactMethod)   // 创建联系方式
		navContactMethodRouter.PUT("update", navContactMethodApi.UpdateNavContactMethod)    // 更新联系方式
		navContactMethodRouter.DELETE("delete", navContactMethodApi.DeleteNavContactMethod) // 删除联系方式
	}
	{
		navContactMethodRouterWithoutRecord.GET("get", navContactMethodApi.GetNavContactMethod)      // 获取单个联系方式
		navContactMethodRouterWithoutRecord.GET("list", navContactMethodApi.GetNavContactMethodList) // 获取联系方式列表
	}
}
