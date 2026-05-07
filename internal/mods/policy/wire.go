package policy

import (
	"github.com/google/wire"
	"github.com/jd-opensource/joylive-dashboard/internal/mods/policy/api"
	"github.com/jd-opensource/joylive-dashboard/internal/mods/policy/biz"
	"github.com/jd-opensource/joylive-dashboard/internal/mods/policy/dal"
)

var Set = wire.NewSet(
	wire.Struct(new(Policy), "*"),
	wire.Struct(new(dal.PolicyLoadbalance), "*"),
	wire.Struct(new(biz.PolicyLoadbalance), "*"),
	wire.Struct(new(api.PolicyLoadbalance), "*"),
	wire.Struct(new(dal.PolicyRoute), "*"),
	wire.Struct(new(biz.PolicyRoute), "*"),
	wire.Struct(new(api.PolicyRoute), "*"),
	wire.Struct(new(dal.PolicyRouteDetail), "*"),
	wire.Struct(new(biz.PolicyRouteDetail), "*"),
	wire.Struct(new(api.PolicyRouteDetail), "*"),
)
