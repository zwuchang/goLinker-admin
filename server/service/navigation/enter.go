package navigation

type ServiceGroup struct {
	ContactConfigService
	ContactMethodService
	NavMenuCheckService
	NavApiCheckService
	NavGameCategoryService
	NavGameService
	NavGameConfigService
	NavBannerService
	NavAccessStatsService
	NavGameCategoryRelationService
	NavPlatformConfigService
	NavMarketConfigService
	NavPlatformRankingService
	NavScrollNoticeService
	NavPWAConfigService
}

var ServiceGroupApp = new(ServiceGroup)
