package service

import (
	"github.com/lzsgo/lzs-go-admin/server/service/example"
	"github.com/lzsgo/lzs-go-admin/server/service/system"
)

type ServiceGroup struct {
	SystemServiceGroup  system.ServiceGroup
	ExampleServiceGroup example.ServiceGroup
}

var ServiceGroupApp = new(ServiceGroup)
