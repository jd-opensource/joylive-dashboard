package dal

import (
	"context"

	"github.com/jd-opensource/joylive-dashboard/internal/mods/resource/schema"
	"github.com/jd-opensource/joylive-dashboard/pkg/errors"
	"github.com/jd-opensource/joylive-dashboard/pkg/util"
	"gorm.io/gorm"
)

// Get service group storage instance (only active records)
func GetServiceGroupDB(ctx context.Context, defDB *gorm.DB) *gorm.DB {
	return util.GetDB(ctx, defDB).Model(new(schema.ServiceGroup)).Where("deleted = '0'")
}

// Service group management
type ServiceGroup struct {
	DB *gorm.DB
}

// Query service groups from the database based on the provided parameters and options.
func (a *ServiceGroup) Query(ctx context.Context, params schema.ServiceGroupQueryParam, opts ...schema.ServiceGroupQueryOptions) (*schema.ServiceGroupQueryResult, error) {
	var opt schema.ServiceGroupQueryOptions
	if len(opts) > 0 {
		opt = opts[0]
	}

	db := GetServiceGroupDB(ctx, a.DB)
	if v := params.ServiceId; v != "" {
		db = db.Where("service_id = ?", v)
	}
	if v := params.Code; v != "" {
		db = db.Where("code = ?", v)
	}
	if v := params.LikeName; v != "" {
		db = db.Where("name LIKE ?", "%"+v+"%")
	}

	var list schema.ServiceGroups
	pageResult, err := util.WrapPageQuery(ctx, db, params.PaginationParam, opt.QueryOptions, &list)
	if err != nil {
		return nil, errors.WithStack(err)
	}

	queryResult := &schema.ServiceGroupQueryResult{
		PageResult: pageResult,
		Data:       list,
	}
	return queryResult, nil
}

// Get the specified service group from the database.
func (a *ServiceGroup) Get(ctx context.Context, id string, opts ...schema.ServiceGroupQueryOptions) (*schema.ServiceGroup, error) {
	var opt schema.ServiceGroupQueryOptions
	if len(opts) > 0 {
		opt = opts[0]
	}

	item := new(schema.ServiceGroup)
	ok, err := util.FindOne(ctx, GetServiceGroupDB(ctx, a.DB).Where("id=?", id), opt.QueryOptions, item)
	if err != nil {
		return nil, errors.WithStack(err)
	} else if !ok {
		return nil, nil
	}
	return item, nil
}

// Exists checks if the specified service group exists in the database.
func (a *ServiceGroup) Exists(ctx context.Context, id string) (bool, error) {
	ok, err := util.Exists(ctx, GetServiceGroupDB(ctx, a.DB).Where("id=?", id))
	return ok, errors.WithStack(err)
}

// ExistsByUniqueKey checks whether a service group with the given unique key (code + service_id) already exists.
func (a *ServiceGroup) ExistsByUniqueKey(ctx context.Context, code, serviceId string) (bool, error) {
	ok, err := util.Exists(ctx, GetServiceGroupDB(ctx, a.DB).
		Where("code = ? AND service_id = ?", code, serviceId))
	return ok, errors.WithStack(err)
}

// Create a new service group.
func (a *ServiceGroup) Create(ctx context.Context, item *schema.ServiceGroup) error {
	result := GetServiceGroupDB(ctx, a.DB).Create(item)
	return errors.WithStack(result.Error)
}

// Update the specified service group in the database.
func (a *ServiceGroup) Update(ctx context.Context, item *schema.ServiceGroup) error {
	result := GetServiceGroupDB(ctx, a.DB).Where("id=?", item.ID).Select("*").Omit("created_at").Updates(item)
	return errors.WithStack(result.Error)
}

// Delete the specified service group from the database using logical deletion.
func (a *ServiceGroup) Delete(ctx context.Context, id string) error {
	return errors.WithStack(util.SoftDelete(ctx, GetServiceGroupDB(ctx, a.DB), id))
}
