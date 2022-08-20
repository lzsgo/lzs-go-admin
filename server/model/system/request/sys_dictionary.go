package request

import (
	"github.com/lzsgo/lzs-go-admin/server/model/common/request"
	"github.com/lzsgo/lzs-go-admin/server/model/system"
)

type SysDictionarySearch struct {
	system.SysDictionary
	request.PageInfo
}
