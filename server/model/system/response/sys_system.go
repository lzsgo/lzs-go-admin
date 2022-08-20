package response

import "github.com/lzsgo/lzs-go-admin/server/config"

type SysConfigResponse struct {
	Config config.Server `json:"config"`
}
