package request

import (
	"goLinker-admin/server/model/common/request"
	"goLinker-admin/server/model/navigation"
)

type NavContactConfigSearch struct {
	navigation.NavContactConfig
	request.PageInfo
}

type NavContactMethodSearch struct {
	navigation.NavContactMethod
	request.PageInfo
}


