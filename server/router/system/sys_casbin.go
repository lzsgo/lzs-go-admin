package system

import (
	"github.com/gin-gonic/gin"
	v1 "github.com/lzsgo/lzs-go-admin/server/api/v1"
	"github.com/lzsgo/lzs-go-admin/server/middleware"
)

type CasbinRouter struct{}

func (s *CasbinRouter) InitCasbinRouter(Router *gin.RouterGroup) {
	casbinRouter := Router.Group("casbin").Use(middleware.OperationRecord())
	casbinRouterWithoutRecord := Router.Group("casbin")
	casbinApi := v1.ApiGroupApp.SystemApiGroup.CasbinApi
	{
		casbinRouter.POST("updateCasbin", casbinApi.UpdateCasbin)
	}
	{
		casbinRouterWithoutRecord.POST("getPolicyPathByAuthorityId", casbinApi.GetPolicyPathByAuthorityId)
	}
}
