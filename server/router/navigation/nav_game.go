package navigation

import (
	"github.com/gin-gonic/gin"
)

type NavGameRouter struct{}

// InitGameRouter 初始化游戏路由
func (s *NavGameRouter) InitGameRouter(Router *gin.RouterGroup) {
	gameRouter := Router.Group("navigation/game")
	{
		gameRouter.POST("createGame", navGameApi.CreateGame)           // 创建游戏
		gameRouter.PUT("updateGame", navGameApi.UpdateGame)            // 更新游戏
		gameRouter.DELETE("deleteGame", navGameApi.DeleteGame)         // 删除游戏
		gameRouter.POST("getGameList", navGameApi.GetGameList)         // 获取游戏列表
		gameRouter.POST("updateGameViews", navGameApi.UpdateGameViews) // 更新游戏浏览次数
	}
}
