package dal

import (
	"context"

	"github.com/jd-opensource/joylive-dashboard/internal/mods/policy/schema"
	"github.com/jd-opensource/joylive-dashboard/pkg/errors"
	"github.com/jd-opensource/joylive-dashboard/pkg/util"
	"gorm.io/gorm"
)

// Get policy auth storage instance (only active records)
func GetPolicyAuthDB(ctx context.Context, defDB *gorm.DB) *gorm.DB {
	return util.GetDB(ctx, defDB).Model(new(schema.PolicyAuth)).Where("deleted = '0'")
}

// Auth policy management
type PolicyAuth struct {
	DB *gorm.DB
}

// Query policy auths from the database based on the provided parameters and options.
func (a *PolicyAuth) Query(ctx context.Context, params schema.PolicyAuthQueryParam, opts ...schema.PolicyAuthQueryOptions) (*schema.PolicyAuthQueryResult, error) {
	var opt schema.PolicyAuthQueryOptions
	if len(opts) > 0 {
		opt = opts[0]
	}

	db := GetPolicyAuthDB(ctx, a.DB)
	if v := params.SpaceCode; v != "" {
		db = db.Where("space_code = ?", v)
	}
	if v := params.TargetServiceId; v != "" {
		db = db.Where("target_service_id = ?", v)
	}

	var list schema.PolicyAuths
	pageResult, err := util.WrapPageQuery(ctx, db, params.PaginationParam, opt.QueryOptions, &list)
	if err != nil {
		return nil, errors.WithStack(err)
	}

	queryResult := &schema.PolicyAuthQueryResult{
		PageResult: pageResult,
		Data:       list,
	}
	return queryResult, nil
}

// Get the specified policy auth from the database.
func (a *PolicyAuth) Get(ctx context.Context, id string, opts ...schema.PolicyAuthQueryOptions) (*schema.PolicyAuth, error) {
	var opt schema.PolicyAuthQueryOptions
	if len(opts) > 0 {
		opt = opts[0]
	}

	item := new(schema.PolicyAuth)
	ok, err := util.FindOne(ctx, GetPolicyAuthDB(ctx, a.DB).Where("id=?", id), opt.QueryOptions, item)
	if err != nil {
		return nil, errors.WithStack(err)
	} else if !ok {
		return nil, nil
	}
	return item, nil
}

// Exists checks if the specified policy auth exists in the database.
func (a *PolicyAuth) Exists(ctx context.Context, id string) (bool, error) {
	ok, err := util.Exists(ctx, GetPolicyAuthDB(ctx, a.DB).Where("id=?", id))
	return ok, errors.WithStack(err)
}

// Create a new policy auth.
func (a *PolicyAuth) Create(ctx context.Context, item *schema.PolicyAuth) error {
	result := GetPolicyAuthDB(ctx, a.DB).Create(item)
	return errors.WithStack(result.Error)
}

// Update the specified policy auth in the database.
func (a *PolicyAuth) Update(ctx context.Context, item *schema.PolicyAuth) error {
	result := GetPolicyAuthDB(ctx, a.DB).Where("id=?", item.ID).Select("*").Omit("created_at").Updates(item)
	return errors.WithStack(result.Error)
}

// Delete the specified policy auth from the database using logical deletion.
func (a *PolicyAuth) Delete(ctx context.Context, id string) error {
	return errors.WithStack(util.SoftDelete(ctx, GetPolicyAuthDB(ctx, a.DB), id))
}
