package system

import (
	"github.com/gin-gonic/gin"
	v1 "github.com/lzsgo/lzs-go-admin/server/api/v1"
)

type InitRouter struct{}

func (s *InitRouter) InitInitRouter(Router *gin.RouterGroup) {
	initRouter := Router.Group("init")
	dbApi := v1.ApiGroupApp.SystemApiGroup.DBApi
	{
		initRouter.POST("initdb", dbApi.InitDB)   // 初始化数据库
		initRouter.POST("checkdb", dbApi.CheckDB) // 检查数据库
		initRouter.POST("resetdb", dbApi.ResetDB) // 重置数据库
	}
}
