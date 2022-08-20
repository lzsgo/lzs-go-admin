package request

import (
	"github.com/lzsgo/lzs-go-admin/server/model/common/request"
	"github.com/lzsgo/lzs-go-admin/server/model/system"
)

type SysDictionaryDetailSearch struct {
	system.SysDictionaryDetail
	request.PageInfo
}
