package request

import (
	"goLinker-admin/server/model/common/request"
	"goLinker-admin/server/model/system"
)

type SysDictionaryDetailSearch struct {
	system.SysDictionaryDetail
	request.PageInfo
}
