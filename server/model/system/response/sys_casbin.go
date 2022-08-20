package response

import (
	"github.com/lzsgo/lzs-go-admin/server/model/system/request"
)

type PolicyPathResponse struct {
	Paths []request.CasbinInfo `json:"paths"`
}
