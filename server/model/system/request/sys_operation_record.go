package request

import (
	"goLinker-admin/server/model/common/request"
	"goLinker-admin/server/model/system"
)

type SysOperationRecordSearch struct {
	system.SysOperationRecord
	request.PageInfo
}
