package policy

import (
	"context"

	"github.com/jd-opensource/joylive-dashboard/internal/config"

	"github.com/gin-gonic/gin"
	"github.com/jd-opensource/joylive-dashboard/internal/mods/policy/api"
	"github.com/jd-opensource/joylive-dashboard/internal/mods/policy/schema"
	"gorm.io/gorm"
)

type Policy struct {
	DB                    *gorm.DB
	PolicyLoadbalanceAPI  *api.PolicyLoadbalance
	PolicyRouteAPI        *api.PolicyRoute
	PolicyRouteDetailAPI  *api.PolicyRouteDetail
	PolicyLimitAPI        *api.PolicyLimit
	PolicyCircuitBreakAPI *api.PolicyCircuitBreak
	PolicyPermissionAPI   *api.PolicyPermission
	PolicyAuthAPI         *api.PolicyAuth
	PolicyFaultAPI        *api.PolicyFault
	PolicyInvocationAPI   *api.PolicyInvocation
}

func (a *Policy) AutoMigrate(ctx context.Context) error {
	return a.DB.AutoMigrate(new(schema.PolicyLoadbalance), new(schema.PolicyRoute), new(schema.PolicyRouteDetail), new(schema.PolicyLimit), new(schema.PolicyCircuitBreak), new(schema.PolicyPermission), new(schema.PolicyAuth), new(schema.PolicyFault), new(schema.PolicyInvocation))
}

func (a *Policy) Init(ctx context.Context) error {
	if config.C.Storage.DB.AutoMigrate {
		if err := a.AutoMigrate(ctx); err != nil {
			return err
		}
	}
	return nil
}

func (a *Policy) RegisterV1Routers(ctx context.Context, v1 *gin.RouterGroup) error {
	v1 = v1.Group("policy")
	policyLoadbalance := v1.Group("policy-loadbalances")
	{
		policyLoadbalance.GET("", a.PolicyLoadbalanceAPI.Query)
		policyLoadbalance.GET(":id", a.PolicyLoadbalanceAPI.Get)
		policyLoadbalance.POST("", a.PolicyLoadbalanceAPI.Create)
		policyLoadbalance.PUT(":id", a.PolicyLoadbalanceAPI.Update)
		policyLoadbalance.DELETE(":id", a.PolicyLoadbalanceAPI.Delete)
	}
	policyRoute := v1.Group("policy-routes")
	{
		policyRoute.GET("", a.PolicyRouteAPI.Query)
		policyRoute.GET(":id", a.PolicyRouteAPI.Get)
		policyRoute.POST("", a.PolicyRouteAPI.Create)
		policyRoute.PUT(":id", a.PolicyRouteAPI.Update)
		policyRoute.DELETE(":id", a.PolicyRouteAPI.Delete)
	}
	policyRouteDetail := v1.Group("policy-route-details")
	{
		policyRouteDetail.GET("", a.PolicyRouteDetailAPI.Query)
		policyRouteDetail.GET(":id", a.PolicyRouteDetailAPI.Get)
		policyRouteDetail.POST("", a.PolicyRouteDetailAPI.Create)
		policyRouteDetail.PUT(":id", a.PolicyRouteDetailAPI.Update)
		policyRouteDetail.DELETE(":id", a.PolicyRouteDetailAPI.Delete)
	}
	policyLimit := v1.Group("policy-limits")
	{
		policyLimit.GET("", a.PolicyLimitAPI.Query)
		policyLimit.GET(":id", a.PolicyLimitAPI.Get)
		policyLimit.POST("", a.PolicyLimitAPI.Create)
		policyLimit.PUT(":id", a.PolicyLimitAPI.Update)
		policyLimit.DELETE(":id", a.PolicyLimitAPI.Delete)
	}
	policyCircuitBreak := v1.Group("policy-circuit-breaks")
	{
		policyCircuitBreak.GET("", a.PolicyCircuitBreakAPI.Query)
		policyCircuitBreak.GET(":id", a.PolicyCircuitBreakAPI.Get)
		policyCircuitBreak.POST("", a.PolicyCircuitBreakAPI.Create)
		policyCircuitBreak.PUT(":id", a.PolicyCircuitBreakAPI.Update)
		policyCircuitBreak.DELETE(":id", a.PolicyCircuitBreakAPI.Delete)
	}
	policyPermission := v1.Group("policy-permissions")
	{
		policyPermission.GET("", a.PolicyPermissionAPI.Query)
		policyPermission.GET(":id", a.PolicyPermissionAPI.Get)
		policyPermission.POST("", a.PolicyPermissionAPI.Create)
		policyPermission.PUT(":id", a.PolicyPermissionAPI.Update)
		policyPermission.DELETE(":id", a.PolicyPermissionAPI.Delete)
	}
	policyAuth := v1.Group("policy-auths")
	{
		policyAuth.GET("", a.PolicyAuthAPI.Query)
		policyAuth.GET(":id", a.PolicyAuthAPI.Get)
		policyAuth.POST("", a.PolicyAuthAPI.Create)
		policyAuth.PUT(":id", a.PolicyAuthAPI.Update)
		policyAuth.DELETE(":id", a.PolicyAuthAPI.Delete)
	}
	policyFault := v1.Group("policy-faults")
	{
		policyFault.GET("", a.PolicyFaultAPI.Query)
		policyFault.GET(":id", a.PolicyFaultAPI.Get)
		policyFault.POST("", a.PolicyFaultAPI.Create)
		policyFault.PUT(":id", a.PolicyFaultAPI.Update)
		policyFault.DELETE(":id", a.PolicyFaultAPI.Delete)
	}
	policyInvocation := v1.Group("policy-invocations")
	{
		policyInvocation.GET("", a.PolicyInvocationAPI.Query)
		policyInvocation.GET(":id", a.PolicyInvocationAPI.Get)
		policyInvocation.POST("", a.PolicyInvocationAPI.Create)
		policyInvocation.PUT(":id", a.PolicyInvocationAPI.Update)
		policyInvocation.DELETE(":id", a.PolicyInvocationAPI.Delete)
	}
	return nil
}

func (a *Policy) Release(ctx context.Context) error {
	return nil
}
