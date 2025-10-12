package service

import (
	"goLinker-admin/server/service/example"
	"goLinker-admin/server/service/navigation"
	"goLinker-admin/server/service/system"
)

var ServiceGroupApp = new(ServiceGroup)

type ServiceGroup struct {
	SystemServiceGroup     system.ServiceGroup
	ExampleServiceGroup    example.ServiceGroup
	NavigationServiceGroup navigation.ServiceGroup
}
