package navigation

import "goLinker-admin/server/service/navigation"

type ApiGroup struct {
	ContactConfigApi
	ContactMethodApi
	NavGameCategoryApi
	NavGameApi
	NavGameConfigApi
}

var (
	contactConfigService   = navigation.ServiceGroupApp.ContactConfigService
	contactMethodService   = navigation.ServiceGroupApp.ContactMethodService
	navGameCategoryService = navigation.ServiceGroupApp.NavGameCategoryService
	navGameService         = navigation.ServiceGroupApp.NavGameService
	navGameConfigService   = navigation.ServiceGroupApp.NavGameConfigService
)
