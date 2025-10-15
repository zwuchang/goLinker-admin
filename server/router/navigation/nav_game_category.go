package navigation

import (
	"goLinker-admin/server/middleware"

	"github.com/gin-gonic/gin"
)

type NavGameCategoryRouter struct{}

// InitGameCategoryRouter 初始化游戏类别路由
func (s *NavGameCategoryRouter) InitGameCategoryRouter(Router *gin.RouterGroup) {
	gameCategoryRouter := Router.Group("navigation/gameCategory")
	gameCategoryNoLoggedRouter := Router.Group("navigation/gameCategory")
	gameCategoryRouter.Use(middleware.OperationRecord())
	{
		gameCategoryRouter.POST("createGameCategory", navGameCategoryApi.CreateGameCategory)   // 创建游戏类别
		gameCategoryRouter.PUT("updateGameCategory", navGameCategoryApi.UpdateGameCategory)    // 更新游戏类别
		gameCategoryRouter.DELETE("deleteGameCategory", navGameCategoryApi.DeleteGameCategory) // 删除游戏类别

	}
	{

		gameCategoryNoLoggedRouter.POST("getGameCategoryList", navGameCategoryApi.GetGameCategoryList)  // 获取游戏类别列表
		gameCategoryNoLoggedRouter.GET("getAllGameCategories", navGameCategoryApi.GetAllGameCategories) // 获取所有启用的游戏类别
	}
}
