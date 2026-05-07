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
	DB                   *gorm.DB
	PolicyLoadbalanceAPI *api.PolicyLoadbalance
	PolicyRouteAPI       *api.PolicyRoute
	PolicyRouteDetailAPI *api.PolicyRouteDetail
}

func (a *Policy) AutoMigrate(ctx context.Context) error {
	return a.DB.AutoMigrate(new(schema.PolicyLoadbalance), new(schema.PolicyRoute), new(schema.PolicyRouteDetail))
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
	return nil
}

func (a *Policy) Release(ctx context.Context) error {
	return nil
}
