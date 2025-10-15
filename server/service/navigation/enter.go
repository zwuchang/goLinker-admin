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
}

var ServiceGroupApp = new(ServiceGroup)
