package navigation

import (
	"goLinker-admin/server/middleware"

	"github.com/gin-gonic/gin"
)

type NavGameConfigRouter struct{}

// InitGameConfigRouter 初始化游戏配置路由
func (s *NavGameConfigRouter) InitGameConfigRouter(Router *gin.RouterGroup) {
	gameConfigRouter := Router.Group("navigation/gameConfig")
	gameConfigNoLoggedRouter := Router.Group("navigation/gameConfig")
	gameConfigRouter.Use(middleware.OperationRecord())
	{
		gameConfigRouter.POST("createGameConfig", navGameConfigApi.CreateGameConfig)   // 创建游戏配置
		gameConfigRouter.PUT("updateGameConfig", navGameConfigApi.UpdateGameConfig)    // 更新游戏配置
		gameConfigRouter.DELETE("deleteGameConfig", navGameConfigApi.DeleteGameConfig) // 删除游戏配置
	}
	{
		gameConfigNoLoggedRouter.GET("getGameConfig", navGameConfigApi.GetGameConfig)          // 获取游戏配置
		gameConfigNoLoggedRouter.POST("getGameConfigList", navGameConfigApi.GetGameConfigList) // 获取游戏配置列表
	}
}
