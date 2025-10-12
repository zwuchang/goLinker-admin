package navigation

type ServiceGroup struct {
	ContactConfigService
	ContactMethodService
	NavMenuCheckService
	NavApiCheckService
	NavGameCategoryService
	NavGameService
	NavGameConfigService
}

var ServiceGroupApp = new(ServiceGroup)
