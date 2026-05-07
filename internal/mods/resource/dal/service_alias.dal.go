package dal

import (
	"context"

	"github.com/jd-opensource/joylive-dashboard/internal/mods/resource/schema"
	"github.com/jd-opensource/joylive-dashboard/pkg/errors"
	"github.com/jd-opensource/joylive-dashboard/pkg/util"
	"gorm.io/gorm"
)

// Get service alias storage instance (only active records)
func GetServiceAliasDB(ctx context.Context, defDB *gorm.DB) *gorm.DB {
	return util.GetDB(ctx, defDB).Model(new(schema.ServiceAlias)).Where("deleted = '0'")
}

// Service alias management
type ServiceAlias struct {
	DB *gorm.DB
}

// Query service aliases from the database based on the provided parameters and options.
func (a *ServiceAlias) Query(ctx context.Context, params schema.ServiceAliasQueryParam, opts ...schema.ServiceAliasQueryOptions) (*schema.ServiceAliasQueryResult, error) {
	var opt schema.ServiceAliasQueryOptions
	if len(opts) > 0 {
		opt = opts[0]
	}

	db := GetServiceAliasDB(ctx, a.DB)
	if v := params.SpaceCode; v != "" {
		db = db.Where("space_code = ?", v)
	}
	if v := params.Alias; v != "" {
		db = db.Where("alias = ?", v)
	}
	if v := params.ServiceId; v != "" {
		db = db.Where("service_id = ?", v)
	}

	var list schema.ServiceAliases
	pageResult, err := util.WrapPageQuery(ctx, db, params.PaginationParam, opt.QueryOptions, &list)
	if err != nil {
		return nil, errors.WithStack(err)
	}

	queryResult := &schema.ServiceAliasQueryResult{
		PageResult: pageResult,
		Data:       list,
	}
	return queryResult, nil
}

// Get the specified service alias from the database.
func (a *ServiceAlias) Get(ctx context.Context, id string, opts ...schema.ServiceAliasQueryOptions) (*schema.ServiceAlias, error) {
	var opt schema.ServiceAliasQueryOptions
	if len(opts) > 0 {
		opt = opts[0]
	}

	item := new(schema.ServiceAlias)
	ok, err := util.FindOne(ctx, GetServiceAliasDB(ctx, a.DB).Where("id=?", id), opt.QueryOptions, item)
	if err != nil {
		return nil, errors.WithStack(err)
	} else if !ok {
		return nil, nil
	}
	return item, nil
}

// Exists checks if the specified service alias exists in the database.
func (a *ServiceAlias) Exists(ctx context.Context, id string) (bool, error) {
	ok, err := util.Exists(ctx, GetServiceAliasDB(ctx, a.DB).Where("id=?", id))
	return ok, errors.WithStack(err)
}

// ExistsByUniqueKey checks whether a service alias with the given unique key (space_code + alias) already exists.
func (a *ServiceAlias) ExistsByUniqueKey(ctx context.Context, spaceCode, alias string) (bool, error) {
	ok, err := util.Exists(ctx, GetServiceAliasDB(ctx, a.DB).
		Where("space_code = ? AND alias = ?", spaceCode, alias))
	return ok, errors.WithStack(err)
}

// Create a new service alias.
func (a *ServiceAlias) Create(ctx context.Context, item *schema.ServiceAlias) error {
	result := GetServiceAliasDB(ctx, a.DB).Create(item)
	return errors.WithStack(result.Error)
}

// Update the specified service alias in the database.
func (a *ServiceAlias) Update(ctx context.Context, item *schema.ServiceAlias) error {
	result := GetServiceAliasDB(ctx, a.DB).Where("id=?", item.ID).Select("*").Omit("created_at").Updates(item)
	return errors.WithStack(result.Error)
}

// Delete the specified service alias from the database using logical deletion.
func (a *ServiceAlias) Delete(ctx context.Context, id string) error {
	return errors.WithStack(util.SoftDelete(ctx, GetServiceAliasDB(ctx, a.DB), id))
}
