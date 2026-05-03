package mods

import (
	"context"

	"github.com/jd-opensource/joylive-dashboard/internal/mods/rbac"
	"github.com/jd-opensource/joylive-dashboard/internal/mods/resource"
	"github.com/jd-opensource/joylive-dashboard/internal/mods/space"
	"github.com/gin-gonic/gin"
	"github.com/google/wire"
)

const (
	apiPrefix = "/api/"
)

// Collection of wire providers
var Set = wire.NewSet(
	wire.Struct(new(Mods), "*"),
	rbac.Set,
	resource.Set,
	space.Set,
)

type Mods struct {
	RBAC     *rbac.RBAC
	Resource *resource.Resource
	Space    *space.Space
}

func (a *Mods) Init(ctx context.Context) error {
	if err := a.RBAC.Init(ctx); err != nil {
		return err
	}
	if err := a.Resource.Init(ctx); err != nil {
		return err
	}
	if err := a.Space.Init(ctx); err != nil {
		return err
	}

	return nil
}

func (a *Mods) RouterPrefixes() []string {
	return []string{
		apiPrefix,
	}
}

func (a *Mods) RegisterRouters(ctx context.Context, e *gin.Engine) error {
	gAPI := e.Group(apiPrefix)
	v1 := gAPI.Group("v1")

	if err := a.RBAC.RegisterV1Routers(ctx, v1); err != nil {
		return err
	}
	if err := a.Resource.RegisterV1Routers(ctx, v1); err != nil {
		return err
	}
	if err := a.Space.RegisterV1Routers(ctx, v1); err != nil {
		return err
	}

	return nil
}

func (a *Mods) Release(ctx context.Context) error {
	if err := a.RBAC.Release(ctx); err != nil {
		return err
	}
	if err := a.Resource.Release(ctx); err != nil {
		return err
	}
	if err := a.Space.Release(ctx); err != nil {
		return err
	}

	return nil
}
