package navigation

import "goLinker-admin/server/service/navigation"

type ApiGroup struct {
	ContactConfigApi
	ContactMethodApi
	NavGameCategoryApi
	NavGameApi
	NavGameConfigApi
	NavBannerApi
	PublicApi
	NavAccessStatsApi
	NavPlatformConfigApi
	NavMarketConfigApi
}

var (
	contactConfigService     = navigation.ServiceGroupApp.ContactConfigService
	contactMethodService     = navigation.ServiceGroupApp.ContactMethodService
	navGameCategoryService   = navigation.ServiceGroupApp.NavGameCategoryService
	navGameService           = navigation.ServiceGroupApp.NavGameService
	navGameConfigService     = navigation.ServiceGroupApp.NavGameConfigService
	navBannerService         = navigation.ServiceGroupApp.NavBannerService
	navAccessStatsService    = navigation.ServiceGroupApp.NavAccessStatsService
	navPlatformConfigService = navigation.ServiceGroupApp.NavPlatformConfigService
	navMarketConfigService   = navigation.ServiceGroupApp.NavMarketConfigService
)
