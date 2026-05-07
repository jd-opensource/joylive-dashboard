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
	wire.Struct(new(dal.PolicyLimit), "*"),
	wire.Struct(new(biz.PolicyLimit), "*"),
	wire.Struct(new(api.PolicyLimit), "*"),
	wire.Struct(new(dal.PolicyCircuitBreak), "*"),
	wire.Struct(new(biz.PolicyCircuitBreak), "*"),
	wire.Struct(new(api.PolicyCircuitBreak), "*"),
	wire.Struct(new(dal.PolicyPermission), "*"),
	wire.Struct(new(biz.PolicyPermission), "*"),
	wire.Struct(new(api.PolicyPermission), "*"),
	wire.Struct(new(dal.PolicyAuth), "*"),
	wire.Struct(new(biz.PolicyAuth), "*"),
	wire.Struct(new(api.PolicyAuth), "*"),
	wire.Struct(new(dal.PolicyFault), "*"),
	wire.Struct(new(biz.PolicyFault), "*"),
	wire.Struct(new(api.PolicyFault), "*"),
	wire.Struct(new(dal.PolicyInvocation), "*"),
	wire.Struct(new(biz.PolicyInvocation), "*"),
	wire.Struct(new(api.PolicyInvocation), "*"),
)
