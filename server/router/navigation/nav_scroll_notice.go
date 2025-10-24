package navigation

import (
	"goLinker-admin/server/middleware"

	"github.com/gin-gonic/gin"
)

type NavScrollNoticeRouter struct{}

// InitNavScrollNoticeRouter 初始化滚动通知路由信息
func (s *NavScrollNoticeRouter) InitNavScrollNoticeRouter(Router *gin.RouterGroup) {
	navScrollNoticeRouter := Router.Group("navigation/notice").Use(middleware.OperationRecord())
	navScrollNoticeRouterWithoutRecord := Router.Group("navigation/notice")
	{
		navScrollNoticeRouter.POST("createNotice", navScrollNoticeApi.CreateScrollNotice) // 新建滚动通知
		navScrollNoticeRouter.POST("deleteNotice", navScrollNoticeApi.DeleteScrollNotice) // 删除滚动通知
		navScrollNoticeRouter.POST("updateNotice", navScrollNoticeApi.UpdateScrollNotice) // 更新滚动通知
	}
	{
		navScrollNoticeRouterWithoutRecord.POST("findNotice", navScrollNoticeApi.GetScrollNoticeById)    // 根据ID获取滚动通知
		navScrollNoticeRouterWithoutRecord.POST("getNoticeList", navScrollNoticeApi.GetScrollNoticeList) // 获取滚动通知列表
	}
}
