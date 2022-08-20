package request

import (
	"github.com/lzsgo/lzs-go-admin/server/model/common/request"
	"github.com/lzsgo/lzs-go-admin/server/model/system"
)

type SysOperationRecordSearch struct {
	system.SysOperationRecord
	request.PageInfo
}
