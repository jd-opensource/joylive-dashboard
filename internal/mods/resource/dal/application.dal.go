package dal

import (
	"context"

	"github.com/jd-opensource/joylive-dashboard/internal/mods/resource/schema"
	"github.com/jd-opensource/joylive-dashboard/pkg/errors"
	"github.com/jd-opensource/joylive-dashboard/pkg/util"
	"gorm.io/gorm"
)

// GetApplicationDB returns the database instance for Application (only active records).
func GetApplicationDB(ctx context.Context, defDB *gorm.DB) *gorm.DB {
	return util.GetDB(ctx, defDB).Model(new(schema.Application)).Where("deleted = '0'")
}

// Application data access layer
type Application struct {
	DB *gorm.DB
}

// Query applications from the database.
func (a *Application) Query(ctx context.Context, params schema.ApplicationQueryParam, opts ...schema.ApplicationQueryOptions) (*schema.ApplicationQueryResult, error) {
	var opt schema.ApplicationQueryOptions
	if len(opts) > 0 {
		opt = opts[0]
	}

	db := GetApplicationDB(ctx, a.DB)
	if v := params.LikeName; len(v) > 0 {
		db = db.Where("name LIKE ?", "%"+v+"%")
	}
	if v := params.UserID; len(v) > 0 {
		permQuery := GetDataPermissionDB(ctx, a.DB).
			Where("type = ? AND user = ? AND tenant = ? AND permission & 1 = 1", schema.DataPermissionTypeApplication, v, params.Tenant).
			Select("data_id")
		db = db.Where("id IN (?)", permQuery)
	}

	var list schema.Applications
	pageResult, err := util.WrapPageQuery(ctx, db, params.PaginationParam, opt.QueryOptions, &list)
	if err != nil {
		return nil, errors.WithStack(err)
	}

	return &schema.ApplicationQueryResult{
		PageResult: pageResult,
		Data:       list,
	}, nil
}

// Get the specified application from the database.
func (a *Application) Get(ctx context.Context, id string, opts ...schema.ApplicationQueryOptions) (*schema.Application, error) {
	var opt schema.ApplicationQueryOptions
	if len(opts) > 0 {
		opt = opts[0]
	}

	item := new(schema.Application)
	ok, err := util.FindOne(ctx, GetApplicationDB(ctx, a.DB).Where("id=?", id), opt.QueryOptions, item)
	if err != nil {
		return nil, errors.WithStack(err)
	} else if !ok {
		return nil, nil
	}
	return item, nil
}

// Exists checks if the specified application exists.
func (a *Application) Exists(ctx context.Context, id string) (bool, error) {
	ok, err := util.Exists(ctx, GetApplicationDB(ctx, a.DB).Where("id=?", id))
	return ok, errors.WithStack(err)
}

// ExistsName checks if an application with the given name exists.
func (a *Application) ExistsName(ctx context.Context, name string) (bool, error) {
	ok, err := util.Exists(ctx, GetApplicationDB(ctx, a.DB).Where("name=?", name))
	return ok, errors.WithStack(err)
}

// Create a new application.
func (a *Application) Create(ctx context.Context, item *schema.Application) error {
	result := GetApplicationDB(ctx, a.DB).Create(item)
	return errors.WithStack(result.Error)
}

// Update the specified application.
func (a *Application) Update(ctx context.Context, item *schema.Application) error {
	result := GetApplicationDB(ctx, a.DB).Where("id=?", item.ID).Select("*").Omit("created_at").Updates(item)
	return errors.WithStack(result.Error)
}

// Delete the specified application logically.
func (a *Application) Delete(ctx context.Context, id string) error {
	return errors.WithStack(util.SoftDelete(ctx, GetApplicationDB(ctx, a.DB), id))
}
