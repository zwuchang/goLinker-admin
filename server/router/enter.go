package router

import (
	"goLinker-admin/server/router/example"
	"goLinker-admin/server/router/navigation"
	"goLinker-admin/server/router/system"
)

var RouterGroupApp = new(RouterGroup)

type RouterGroup struct {
	System     system.RouterGroup
	Example    example.RouterGroup
	Navigation navigation.RouterGroup
}
