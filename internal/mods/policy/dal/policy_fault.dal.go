package dal

import (
	"context"

	"github.com/jd-opensource/joylive-dashboard/internal/mods/policy/schema"
	"github.com/jd-opensource/joylive-dashboard/pkg/errors"
	"github.com/jd-opensource/joylive-dashboard/pkg/util"
	"gorm.io/gorm"
)

// Get policy fault storage instance (only active records)
func GetPolicyFaultDB(ctx context.Context, defDB *gorm.DB) *gorm.DB {
	return util.GetDB(ctx, defDB).Model(new(schema.PolicyFault)).Where("deleted = '0'")
}

// Fault injection policy management
type PolicyFault struct {
	DB *gorm.DB
}

// Query policy faults from the database based on the provided parameters and options.
func (a *PolicyFault) Query(ctx context.Context, params schema.PolicyFaultQueryParam, opts ...schema.PolicyFaultQueryOptions) (*schema.PolicyFaultQueryResult, error) {
	var opt schema.PolicyFaultQueryOptions
	if len(opts) > 0 {
		opt = opts[0]
	}

	db := GetPolicyFaultDB(ctx, a.DB)
	if v := params.SpaceCode; v != "" {
		db = db.Where("space_code = ?", v)
	}
	if v := params.Name; v != "" {
		db = db.Where("name LIKE ?", "%"+v+"%")
	}
	if v := params.TargetServiceId; v != "" {
		db = db.Where("target_service_id = ?", v)
	}

	var list schema.PolicyFaults
	pageResult, err := util.WrapPageQuery(ctx, db, params.PaginationParam, opt.QueryOptions, &list)
	if err != nil {
		return nil, errors.WithStack(err)
	}

	queryResult := &schema.PolicyFaultQueryResult{
		PageResult: pageResult,
		Data:       list,
	}
	return queryResult, nil
}

// Get the specified policy fault from the database.
func (a *PolicyFault) Get(ctx context.Context, id string, opts ...schema.PolicyFaultQueryOptions) (*schema.PolicyFault, error) {
	var opt schema.PolicyFaultQueryOptions
	if len(opts) > 0 {
		opt = opts[0]
	}

	item := new(schema.PolicyFault)
	ok, err := util.FindOne(ctx, GetPolicyFaultDB(ctx, a.DB).Where("id=?", id), opt.QueryOptions, item)
	if err != nil {
		return nil, errors.WithStack(err)
	} else if !ok {
		return nil, nil
	}
	return item, nil
}

// Exists checks if the specified policy fault exists in the database.
func (a *PolicyFault) Exists(ctx context.Context, id string) (bool, error) {
	ok, err := util.Exists(ctx, GetPolicyFaultDB(ctx, a.DB).Where("id=?", id))
	return ok, errors.WithStack(err)
}

// Create a new policy fault.
func (a *PolicyFault) Create(ctx context.Context, item *schema.PolicyFault) error {
	result := GetPolicyFaultDB(ctx, a.DB).Create(item)
	return errors.WithStack(result.Error)
}

// Update the specified policy fault in the database.
func (a *PolicyFault) Update(ctx context.Context, item *schema.PolicyFault) error {
	result := GetPolicyFaultDB(ctx, a.DB).Where("id=?", item.ID).Select("*").Omit("created_at").Updates(item)
	return errors.WithStack(result.Error)
}

// ExistsByUniqueKey checks whether a policy fault with the given unique key already exists.
func (a *PolicyFault) ExistsByUniqueKey(ctx context.Context, name, spaceCode, sourceApplicationID, targetServiceId string) (bool, error) {
	ok, err := util.Exists(ctx, GetPolicyFaultDB(ctx, a.DB).
		Where("name = ? AND space_code = ? AND source_application_id = ? AND target_service_id = ?", name, spaceCode, sourceApplicationID, targetServiceId))
	return ok, errors.WithStack(err)
}

// Delete the specified policy fault from the database using logical deletion.
func (a *PolicyFault) Delete(ctx context.Context, id string) error {
	return errors.WithStack(util.SoftDelete(ctx, GetPolicyFaultDB(ctx, a.DB), id))
}
