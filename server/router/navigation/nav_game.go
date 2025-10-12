package navigation

import (
	api "goLinker-admin/server/api/v1"

	"github.com/gin-gonic/gin"
)

type NavGameRouter struct{}

// InitGameRouter 初始化游戏路由
func (s *NavGameRouter) InitGameRouter(Router *gin.RouterGroup) {
	gameRouter := Router.Group("navigation/game")
	{
		gameRouter.POST("createGame", api.ApiGroupApp.NavigationApiGroup.NavGameApi.CreateGame)           // 创建游戏
		gameRouter.PUT("updateGame", api.ApiGroupApp.NavigationApiGroup.NavGameApi.UpdateGame)            // 更新游戏
		gameRouter.DELETE("deleteGame", api.ApiGroupApp.NavigationApiGroup.NavGameApi.DeleteGame)         // 删除游戏
		gameRouter.POST("getGameList", api.ApiGroupApp.NavigationApiGroup.NavGameApi.GetGameList)         // 获取游戏列表
		gameRouter.POST("updateGameViews", api.ApiGroupApp.NavigationApiGroup.NavGameApi.UpdateGameViews) // 更新游戏浏览次数
	}
}
