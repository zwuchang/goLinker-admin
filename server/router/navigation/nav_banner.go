package navigation

import (
	"goLinker-admin/server/middleware"

	"github.com/gin-gonic/gin"
)

type NavBannerRouter struct{}

// InitBannerRouter 初始化Banner路由
func (s *NavBannerRouter) InitBannerRouter(Router *gin.RouterGroup) {
	bannerNoLoggedRouter := Router.Group("navigation/banner")
	bannerRouter := Router.Group("navigation/banner")
	bannerRouter.Use(middleware.OperationRecord())
	{
		bannerRouter.POST("createBanner", navBannerApi.CreateBanner) // 创建Banner
		bannerRouter.POST("updateBanner", navBannerApi.UpdateBanner) // 更新Banner
		bannerRouter.POST("deleteBanner", navBannerApi.DeleteBanner) // 删除Banner
	}
	{
		bannerNoLoggedRouter.POST("getBannerList", navBannerApi.GetBannerList)    // 获取Banner列表
		bannerNoLoggedRouter.GET("getBannerById/:id", navBannerApi.GetBannerById) // 根据ID获取Banner
	}

}
