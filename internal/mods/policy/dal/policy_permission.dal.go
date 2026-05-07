package dal

import (
	"context"

	"github.com/jd-opensource/joylive-dashboard/internal/mods/policy/schema"
	"github.com/jd-opensource/joylive-dashboard/pkg/errors"
	"github.com/jd-opensource/joylive-dashboard/pkg/util"
	"gorm.io/gorm"
)

// Get policy permission storage instance (only active records)
func GetPolicyPermissionDB(ctx context.Context, defDB *gorm.DB) *gorm.DB {
	return util.GetDB(ctx, defDB).Model(new(schema.PolicyPermission)).Where("deleted = '0'")
}

// Permission policy management
type PolicyPermission struct {
	DB *gorm.DB
}

// Query policy permissions from the database based on the provided parameters and options.
func (a *PolicyPermission) Query(ctx context.Context, params schema.PolicyPermissionQueryParam, opts ...schema.PolicyPermissionQueryOptions) (*schema.PolicyPermissionQueryResult, error) {
	var opt schema.PolicyPermissionQueryOptions
	if len(opts) > 0 {
		opt = opts[0]
	}

	db := GetPolicyPermissionDB(ctx, a.DB)
	if v := params.SpaceCode; v != "" {
		db = db.Where("space_code = ?", v)
	}
	if v := params.Name; v != "" {
		db = db.Where("name LIKE ?", "%"+v+"%")
	}
	if v := params.TargetServiceId; v != "" {
		db = db.Where("target_service_id = ?", v)
	}

	var list schema.PolicyPermissions
	pageResult, err := util.WrapPageQuery(ctx, db, params.PaginationParam, opt.QueryOptions, &list)
	if err != nil {
		return nil, errors.WithStack(err)
	}

	queryResult := &schema.PolicyPermissionQueryResult{
		PageResult: pageResult,
		Data:       list,
	}
	return queryResult, nil
}

// Get the specified policy permission from the database.
func (a *PolicyPermission) Get(ctx context.Context, id string, opts ...schema.PolicyPermissionQueryOptions) (*schema.PolicyPermission, error) {
	var opt schema.PolicyPermissionQueryOptions
	if len(opts) > 0 {
		opt = opts[0]
	}

	item := new(schema.PolicyPermission)
	ok, err := util.FindOne(ctx, GetPolicyPermissionDB(ctx, a.DB).Where("id=?", id), opt.QueryOptions, item)
	if err != nil {
		return nil, errors.WithStack(err)
	} else if !ok {
		return nil, nil
	}
	return item, nil
}

// Exists checks if the specified policy permission exists in the database.
func (a *PolicyPermission) Exists(ctx context.Context, id string) (bool, error) {
	ok, err := util.Exists(ctx, GetPolicyPermissionDB(ctx, a.DB).Where("id=?", id))
	return ok, errors.WithStack(err)
}

// ExistsByUniqueKey checks whether a policy permission with the given unique key already exists.
func (a *PolicyPermission) ExistsByUniqueKey(ctx context.Context, name, targetServiceId string) (bool, error) {
	ok, err := util.Exists(ctx, GetPolicyPermissionDB(ctx, a.DB).
		Where("name = ? AND target_service_id = ?", name, targetServiceId))
	return ok, errors.WithStack(err)
}

// Create a new policy permission.
func (a *PolicyPermission) Create(ctx context.Context, item *schema.PolicyPermission) error {
	result := GetPolicyPermissionDB(ctx, a.DB).Create(item)
	return errors.WithStack(result.Error)
}

// Update the specified policy permission in the database.
func (a *PolicyPermission) Update(ctx context.Context, item *schema.PolicyPermission) error {
	result := GetPolicyPermissionDB(ctx, a.DB).Where("id=?", item.ID).Select("*").Omit("created_at").Updates(item)
	return errors.WithStack(result.Error)
}

// Delete the specified policy permission from the database using logical deletion.
func (a *PolicyPermission) Delete(ctx context.Context, id string) error {
	return errors.WithStack(util.SoftDelete(ctx, GetPolicyPermissionDB(ctx, a.DB), id))
}
