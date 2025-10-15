package navigation

import api "goLinker-admin/server/api/v1"

type RouterGroup struct {
	ContactConfigRouter
	ContactMethodRouter
	NavGameCategoryRouter
	NavGameRouter
	NavGameConfigRouter
	NavBannerRouter
	PublicRouter
	NavAccessStatsRouter
}

var (
	navContactConfigApi = api.ApiGroupApp.NavigationApiGroup.ContactConfigApi
	navContactMethodApi = api.ApiGroupApp.NavigationApiGroup.ContactMethodApi
	navGameCategoryApi  = api.ApiGroupApp.NavigationApiGroup.NavGameCategoryApi
	navGameApi          = api.ApiGroupApp.NavigationApiGroup.NavGameApi
	navGameConfigApi    = api.ApiGroupApp.NavigationApiGroup.NavGameConfigApi
	navBannerApi        = api.ApiGroupApp.NavigationApiGroup.NavBannerApi
	publicApi           = api.ApiGroupApp.NavigationApiGroup.PublicApi
	navAccessStatsApi   = api.ApiGroupApp.NavigationApiGroup.NavAccessStatsApi
)
