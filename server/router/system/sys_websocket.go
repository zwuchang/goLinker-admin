package system

import (
	"goLinker-admin/server/middleware"

	"github.com/gin-gonic/gin"
)

type WebSocketRouter struct{}

func (s *WebSocketRouter) InitWebSocketRouter(Router *gin.RouterGroup) {
	wsRouter := Router.Group("ws")
	{
		// 升级为 Websocket
		wsRouter.GET("/go", middleware.JWTAuth(), wsApi.Upgrade)
	}
}
