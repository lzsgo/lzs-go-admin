package request

import (
	"github.com/lzsgo/lzs-go-admin/server/model/{{.Package}}"
	"github.com/lzsgo/lzs-go-admin/server/model/common/request"
)

type {{.StructName}}Search struct{
    {{.Package}}.{{.StructName}}
    request.PageInfo
}
