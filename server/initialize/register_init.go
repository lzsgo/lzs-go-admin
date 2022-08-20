package initialize

import (
	_ "github.com/lzsgo/lzs-go-admin/server/source/example"
	_ "github.com/lzsgo/lzs-go-admin/server/source/system"
)

func init() {
	// do nothing,only import source package so that inits can be registered
}
