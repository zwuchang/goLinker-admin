package navigation

import (
	api "goLinker-admin/server/api/v1"

	"github.com/gin-gonic/gin"
)

type NavGameCategoryRouter struct{}

// InitGameCategoryRouter 初始化游戏类别路由
func (s *NavGameCategoryRouter) InitGameCategoryRouter(Router *gin.RouterGroup) {
	gameCategoryRouter := Router.Group("navigation/gameCategory")
	{
		gameCategoryRouter.POST("createGameCategory", api.ApiGroupApp.NavigationApiGroup.NavGameCategoryApi.CreateGameCategory)    // 创建游戏类别
		gameCategoryRouter.PUT("updateGameCategory", api.ApiGroupApp.NavigationApiGroup.NavGameCategoryApi.UpdateGameCategory)     // 更新游戏类别
		gameCategoryRouter.DELETE("deleteGameCategory", api.ApiGroupApp.NavigationApiGroup.NavGameCategoryApi.DeleteGameCategory)  // 删除游戏类别
		gameCategoryRouter.POST("getGameCategoryList", api.ApiGroupApp.NavigationApiGroup.NavGameCategoryApi.GetGameCategoryList)  // 获取游戏类别列表
		gameCategoryRouter.GET("getAllGameCategories", api.ApiGroupApp.NavigationApiGroup.NavGameCategoryApi.GetAllGameCategories) // 获取所有启用的游戏类别
	}
}
