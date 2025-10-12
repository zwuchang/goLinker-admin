package navigation

type ServiceGroup struct {
	ContactConfigService
	ContactMethodService
	NavMenuCheckService
	NavApiCheckService
	NavGameCategoryService
	NavGameService
}

var ServiceGroupApp = new(ServiceGroup)
