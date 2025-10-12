package navigation

import "goLinker-admin/server/service/navigation"

type ApiGroup struct {
	ContactConfigApi
	ContactMethodApi
}

var (
	contactConfigService = navigation.ServiceGroupApp.ContactConfigService
	contactMethodService = navigation.ServiceGroupApp.ContactMethodService
)
