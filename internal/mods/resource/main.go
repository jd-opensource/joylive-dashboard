package resource

import (
	"context"

	"github.com/jd-opensource/joylive-dashboard/internal/config"
	"github.com/jd-opensource/joylive-dashboard/internal/mods/resource/api"
	"github.com/jd-opensource/joylive-dashboard/internal/mods/resource/schema"
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

type Resource struct {
	DB            *gorm.DB
	ApplicationAPI *api.Application
	ServiceAPI     *api.Service
}

func (a *Resource) AutoMigrate(ctx context.Context) error {
	return a.DB.AutoMigrate(
		new(schema.Application),
		new(schema.Service),
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

	return nil
}

func (a *Resource) Release(ctx context.Context) error {
	return nil
}
