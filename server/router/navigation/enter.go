package navigation

import api "goLinker-admin/server/api/v1"

type RouterGroup struct {
	ContactConfigRouter
	ContactMethodRouter
}

var (
	navContactConfigApi = api.ApiGroupApp.NavigationApiGroup.ContactConfigApi
	navContactMethodApi = api.ApiGroupApp.NavigationApiGroup.ContactMethodApi
)
