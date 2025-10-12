package navigation

import (
	"github.com/gin-gonic/gin"
)

type NavGameCategoryRouter struct{}

// InitGameCategoryRouter 初始化游戏类别路由
func (s *NavGameCategoryRouter) InitGameCategoryRouter(Router *gin.RouterGroup) {
	gameCategoryRouter := Router.Group("navigation/gameCategory")
	{
		gameCategoryRouter.POST("createGameCategory", navGameCategoryApi.CreateGameCategory)    // 创建游戏类别
		gameCategoryRouter.PUT("updateGameCategory", navGameCategoryApi.UpdateGameCategory)     // 更新游戏类别
		gameCategoryRouter.DELETE("deleteGameCategory", navGameCategoryApi.DeleteGameCategory)  // 删除游戏类别
		gameCategoryRouter.POST("getGameCategoryList", navGameCategoryApi.GetGameCategoryList)  // 获取游戏类别列表
		gameCategoryRouter.GET("getAllGameCategories", navGameCategoryApi.GetAllGameCategories) // 获取所有启用的游戏类别
	}
}
