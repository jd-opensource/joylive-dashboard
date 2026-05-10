package resource

import (
	"context"

	"github.com/gin-gonic/gin"
	"github.com/jd-opensource/joylive-dashboard/internal/config"
	"github.com/jd-opensource/joylive-dashboard/internal/mods/resource/api"
	"github.com/jd-opensource/joylive-dashboard/internal/mods/resource/schema"
	"gorm.io/gorm"
)

type Resource struct {
	DB                    *gorm.DB
	ApplicationAPI        *api.Application
	ServiceAPI            *api.Service
	ApplicationServiceAPI *api.ApplicationService
	ServiceGroupAPI       *api.ServiceGroup
	ServiceAliasAPI       *api.ServiceAlias
	DataPermissionAPI     *api.DataPermission
}

func (a *Resource) AutoMigrate(ctx context.Context) error {
	return a.DB.AutoMigrate(
		new(schema.Application),
		new(schema.Service),
		new(schema.ApplicationService),
		new(schema.ServiceGroup),
		new(schema.ServiceAlias),
		new(schema.DataPermission),
	)
}

func (a *Resource) Init(ctx context.Context) error {
	if config.C.Storage.DB.AutoMigrate {
		if err := a.AutoMigrate(ctx); err != nil {
			return err
		}
	}
	return nil
}

func (a *Resource) RegisterV1Routers(ctx context.Context, v1 *gin.RouterGroup) error {
	app := v1.Group("applications")
	{
		app.GET("", a.ApplicationAPI.Query)
		app.GET(":id", a.ApplicationAPI.Get)
		app.POST("", a.ApplicationAPI.Create)
		app.PUT(":id", a.ApplicationAPI.Update)
		app.DELETE(":id", a.ApplicationAPI.Delete)
	}

	svc := v1.Group("services")
	{
		svc.GET("", a.ServiceAPI.Query)
		svc.GET(":id", a.ServiceAPI.Get)
		svc.POST("", a.ServiceAPI.Create)
		svc.PUT(":id", a.ServiceAPI.Update)
		svc.DELETE(":id", a.ServiceAPI.Delete)
	}
	applicationService := v1.Group("application-services")
	{
		applicationService.GET("", a.ApplicationServiceAPI.Query)
		applicationService.GET(":id", a.ApplicationServiceAPI.Get)
		applicationService.POST("", a.ApplicationServiceAPI.Create)
		applicationService.PUT(":id", a.ApplicationServiceAPI.Update)
		applicationService.DELETE(":id", a.ApplicationServiceAPI.Delete)
	}
	serviceGroup := v1.Group("service-groups")
	{
		serviceGroup.GET("", a.ServiceGroupAPI.Query)
		serviceGroup.GET(":id", a.ServiceGroupAPI.Get)
		serviceGroup.POST("", a.ServiceGroupAPI.Create)
		serviceGroup.PUT(":id", a.ServiceGroupAPI.Update)
		serviceGroup.DELETE(":id", a.ServiceGroupAPI.Delete)
	}
	serviceAlias := v1.Group("service-aliases")
	{
		serviceAlias.GET("", a.ServiceAliasAPI.Query)
		serviceAlias.GET(":id", a.ServiceAliasAPI.Get)
		serviceAlias.POST("", a.ServiceAliasAPI.Create)
		serviceAlias.PUT(":id", a.ServiceAliasAPI.Update)
		serviceAlias.DELETE(":id", a.ServiceAliasAPI.Delete)
	}
	dataPermission := v1.Group("data-permissions")
	{
		dataPermission.GET("", a.DataPermissionAPI.Query)
		dataPermission.GET(":id", a.DataPermissionAPI.Get)
		dataPermission.POST("", a.DataPermissionAPI.Create)
		dataPermission.PUT(":id", a.DataPermissionAPI.Update)
		dataPermission.DELETE(":id", a.DataPermissionAPI.Delete)
	}

	return nil
}

func (a *Resource) Release(ctx context.Context) error {
	return nil
}
