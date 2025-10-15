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
}

var ServiceGroupApp = new(ServiceGroup)
