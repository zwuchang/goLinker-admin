package initialize

import (
	_ "goLinker-admin/server/source/example"
	_ "goLinker-admin/server/source/system"
)

func init() {
	// do nothing,only import source package so that inits can be registered
}
