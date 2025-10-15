package navigation

import (
	v1 "goLinker-admin/server/api/v1"
	"goLinker-admin/server/middleware"

	"github.com/gin-gonic/gin"
)

type NavBannerRouter struct{}

// InitBannerRouter 初始化Banner路由
func (s *NavBannerRouter) InitBannerRouter(Router *gin.RouterGroup) {
	bannerRouter := Router.Group("navigation/banner")
	bannerRouter.Use(middleware.OperationRecord())
	{
		bannerRouter.POST("getBannerList", v1.ApiGroupApp.NavigationApiGroup.NavBannerApi.GetBannerList)    // 获取Banner列表
		bannerRouter.GET("getBannerById/:id", v1.ApiGroupApp.NavigationApiGroup.NavBannerApi.GetBannerById) // 根据ID获取Banner
		bannerRouter.POST("createBanner", v1.ApiGroupApp.NavigationApiGroup.NavBannerApi.CreateBanner)      // 创建Banner
		bannerRouter.POST("updateBanner", v1.ApiGroupApp.NavigationApiGroup.NavBannerApi.UpdateBanner)      // 更新Banner
		bannerRouter.POST("deleteBanner", v1.ApiGroupApp.NavigationApiGroup.NavBannerApi.DeleteBanner)      // 删除Banner
	}
}
