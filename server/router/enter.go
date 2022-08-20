package router

import (
	"github.com/lzsgo/lzs-go-admin/server/router/example"
	"github.com/lzsgo/lzs-go-admin/server/router/system"
)

type RouterGroup struct {
	System  system.RouterGroup
	Example example.RouterGroup
}

var RouterGroupApp = new(RouterGroup)
