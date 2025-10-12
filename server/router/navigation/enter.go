package navigation

import api "goLinker-admin/server/api/v1"

type RouterGroup struct {
	ContactConfigRouter
	ContactMethodRouter
	NavGameCategoryRouter
	NavGameRouter
}

var (
	navContactConfigApi = api.ApiGroupApp.NavigationApiGroup.ContactConfigApi
	navContactMethodApi = api.ApiGroupApp.NavigationApiGroup.ContactMethodApi
	navGameCategoryApi  = api.ApiGroupApp.NavigationApiGroup.NavGameCategoryApi
	navGameApi          = api.ApiGroupApp.NavigationApiGroup.NavGameApi
)
