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
	NavPlatformRankingApi
	NavScrollNoticeApi
	NavPWAConfigApi
	NavThemeConfigApi
	NavActivityConfigApi
}

var (
	contactConfigService      = navigation.ServiceGroupApp.ContactConfigService
	contactMethodService      = navigation.ServiceGroupApp.ContactMethodService
	navGameCategoryService    = navigation.ServiceGroupApp.NavGameCategoryService
	navGameService            = navigation.ServiceGroupApp.NavGameService
	navGameConfigService      = navigation.ServiceGroupApp.NavGameConfigService
	navBannerService          = navigation.ServiceGroupApp.NavBannerService
	navAccessStatsService     = navigation.ServiceGroupApp.NavAccessStatsService
	navPlatformConfigService  = navigation.ServiceGroupApp.NavPlatformConfigService
	navMarketConfigService    = navigation.ServiceGroupApp.NavMarketConfigService
	navPlatformRankingService = navigation.ServiceGroupApp.NavPlatformRankingService
	navScrollNoticeService    = navigation.ServiceGroupApp.NavScrollNoticeService
	navPWAConfigService       = navigation.ServiceGroupApp.NavPWAConfigService
	navThemeConfigService     = navigation.ServiceGroupApp.NavThemeConfigService
	navActivityConfigService  = navigation.ServiceGroupApp.NavActivityConfigService
)
