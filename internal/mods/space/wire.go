package space

import (
	"github.com/jd-opensource/joylive-dashboard/internal/mods/space/api"
	"github.com/jd-opensource/joylive-dashboard/internal/mods/space/biz"
	"github.com/jd-opensource/joylive-dashboard/internal/mods/space/dal"
	"github.com/google/wire"
)

// Collection of wire providers
var Set = wire.NewSet(
	wire.Struct(new(Space), "*"),
	wire.Struct(new(dal.Space), "*"),
	wire.Struct(new(biz.Space), "*"),
	wire.Struct(new(api.Space), "*"),
)
