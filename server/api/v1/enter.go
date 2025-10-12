package v1

import (
	"goLinker-admin/server/api/v1/example"
	"goLinker-admin/server/api/v1/navigation"
	"goLinker-admin/server/api/v1/system"
)

var ApiGroupApp = new(ApiGroup)

type ApiGroup struct {
	SystemApiGroup     system.ApiGroup
	ExampleApiGroup    example.ApiGroup
	NavigationApiGroup navigation.ApiGroup
}
