package navigation

type ServiceGroup struct {
	ContactConfigService
	ContactMethodService
	NavMenuCheckService
	NavApiCheckService
}

var ServiceGroupApp = new(ServiceGroup)
