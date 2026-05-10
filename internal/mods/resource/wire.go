package resource

import (
	"github.com/google/wire"
	"github.com/jd-opensource/joylive-dashboard/internal/mods/resource/api"
	"github.com/jd-opensource/joylive-dashboard/internal/mods/resource/biz"
	"github.com/jd-opensource/joylive-dashboard/internal/mods/resource/dal"
)

// Collection of wire providers
var Set = wire.NewSet(
	wire.Struct(new(Resource), "*"),
	wire.Struct(new(dal.Application), "*"),
	wire.Struct(new(biz.Application), "*"),
	wire.Struct(new(api.Application), "*"),
	wire.Struct(new(dal.Service), "*"),
	wire.Struct(new(biz.Service), "*"),
	wire.Struct(new(api.Service), "*"),
	wire.Struct(new(dal.ApplicationService), "*"),
	wire.Struct(new(biz.ApplicationService), "*"),
	wire.Struct(new(api.ApplicationService), "*"),
	wire.Struct(new(dal.ServiceGroup), "*"),
	wire.Struct(new(biz.ServiceGroup), "*"),
	wire.Struct(new(api.ServiceGroup), "*"),
	wire.Struct(new(dal.ServiceAlias), "*"),
	wire.Struct(new(biz.ServiceAlias), "*"),
	wire.Struct(new(api.ServiceAlias), "*"),
	wire.Struct(new(dal.DataPermission), "*"),
	wire.Struct(new(biz.DataPermission), "*"),
	wire.Struct(new(api.DataPermission), "*"),
)
