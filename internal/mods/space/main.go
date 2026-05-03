package space

import (
	"context"

	"github.com/jd-opensource/joylive-dashboard/internal/config"
	"github.com/jd-opensource/joylive-dashboard/internal/mods/space/api"
	"github.com/jd-opensource/joylive-dashboard/internal/mods/space/schema"
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

type Space struct {
	DB      *gorm.DB
	SpaceAPI *api.Space
}

func (a *Space) AutoMigrate(ctx context.Context) error {
	return a.DB.AutoMigrate(
		new(schema.Space),
	)
}

func (a *Space) Init(ctx context.Context) error {
	if config.C.Storage.DB.AutoMigrate {
		if err := a.AutoMigrate(ctx); err != nil {
			return err
		}
	}
	return nil
}

func (a *Space) RegisterV1Routers(ctx context.Context, v1 *gin.RouterGroup) error {
	space := v1.Group("spaces")
	{
		space.GET("", a.SpaceAPI.Query)
		space.GET(":id", a.SpaceAPI.Get)
		space.POST("", a.SpaceAPI.Create)
		space.PUT(":id", a.SpaceAPI.Update)
		space.DELETE(":id", a.SpaceAPI.Delete)
	}
	return nil
}

func (a *Space) Release(ctx context.Context) error {
	return nil
}
