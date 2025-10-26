package navigation

import (
	"goLinker-admin/server/middleware"

	"github.com/gin-gonic/gin"
)

type NavThemeConfigRouter struct{}

// InitNavThemeConfigRouter 初始化主题配置路由
func (s *NavThemeConfigRouter) InitNavThemeConfigRouter(Router *gin.RouterGroup) {
	navThemeConfigNoLoggedRouter := Router.Group("navigation/themeConfig")
	navThemeConfigRouter := Router.Group("navigation/themeConfig").Use(middleware.OperationRecord())
	{
		navThemeConfigRouter.POST("createThemeConfig", navThemeConfigApi.CreateThemeConfig) // 创建主题配置
		navThemeConfigRouter.POST("updateThemeConfig", navThemeConfigApi.UpdateThemeConfig) // 更新主题配置
		navThemeConfigRouter.POST("deleteThemeConfig", navThemeConfigApi.DeleteThemeConfig) // 删除主题配置
		navThemeConfigRouter.POST("setDefaultTheme", navThemeConfigApi.SetDefaultTheme)     // 设置默认主题
	}
	{
		navThemeConfigNoLoggedRouter.POST("getThemeConfigById", navThemeConfigApi.GetThemeConfigById) // 根据ID获取主题配置
		navThemeConfigNoLoggedRouter.POST("getThemeConfigList", navThemeConfigApi.GetThemeConfigList) // 获取主题配置列表
		navThemeConfigNoLoggedRouter.POST("getDefaultTheme", navThemeConfigApi.GetDefaultTheme)       // 获取默认主题
		navThemeConfigNoLoggedRouter.POST("getAllThemes", navThemeConfigApi.GetAllThemes)             // 获取所有主题
	}
}
