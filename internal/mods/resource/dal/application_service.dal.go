package dal

import (
	"context"

	"github.com/jd-opensource/joylive-dashboard/internal/mods/resource/schema"
	"github.com/jd-opensource/joylive-dashboard/pkg/errors"
	"github.com/jd-opensource/joylive-dashboard/pkg/util"
	"gorm.io/gorm"
)

// Get application service storage instance (only active records)
func GetApplicationServiceDB(ctx context.Context, defDB *gorm.DB) *gorm.DB {
	return util.GetDB(ctx, defDB).Model(new(schema.ApplicationService)).Where("deleted = '0'")
}

// Application and service relationship management
type ApplicationService struct {
	DB *gorm.DB
}

// Query application services from the database based on the provided parameters and options.
func (a *ApplicationService) Query(ctx context.Context, params schema.ApplicationServiceQueryParam, opts ...schema.ApplicationServiceQueryOptions) (*schema.ApplicationServiceQueryResult, error) {
	var opt schema.ApplicationServiceQueryOptions
	if len(opts) > 0 {
		opt = opts[0]
	}

	db := GetApplicationServiceDB(ctx, a.DB)

	var list schema.ApplicationServices
	pageResult, err := util.WrapPageQuery(ctx, db, params.PaginationParam, opt.QueryOptions, &list)
	if err != nil {
		return nil, errors.WithStack(err)
	}

	queryResult := &schema.ApplicationServiceQueryResult{
		PageResult: pageResult,
		Data:       list,
	}
	return queryResult, nil
}

// Get the specified application service from the database.
func (a *ApplicationService) Get(ctx context.Context, id string, opts ...schema.ApplicationServiceQueryOptions) (*schema.ApplicationService, error) {
	var opt schema.ApplicationServiceQueryOptions
	if len(opts) > 0 {
		opt = opts[0]
	}

	item := new(schema.ApplicationService)
	ok, err := util.FindOne(ctx, GetApplicationServiceDB(ctx, a.DB).Where("id=?", id), opt.QueryOptions, item)
	if err != nil {
		return nil, errors.WithStack(err)
	} else if !ok {
		return nil, nil
	}
	return item, nil
}

// Exists checks if the specified application service exists in the database.
func (a *ApplicationService) Exists(ctx context.Context, id string) (bool, error) {
	ok, err := util.Exists(ctx, GetApplicationServiceDB(ctx, a.DB).Where("id=?", id))
	return ok, errors.WithStack(err)
}

// ExistsByUniqueKey checks if an application service with the given unique key exists.
func (a *ApplicationService) ExistsByUniqueKey(ctx context.Context, applicationId, serviceId, role string) (bool, error) {
	ok, err := util.Exists(ctx, GetApplicationServiceDB(ctx, a.DB).
		Where("application_id=? AND service_id=? AND role=?", applicationId, serviceId, role))
	return ok, errors.WithStack(err)
}

// Create a new application service.
func (a *ApplicationService) Create(ctx context.Context, item *schema.ApplicationService) error {
	result := GetApplicationServiceDB(ctx, a.DB).Create(item)
	return errors.WithStack(result.Error)
}

// Update the specified application service in the database.
func (a *ApplicationService) Update(ctx context.Context, item *schema.ApplicationService) error {
	result := GetApplicationServiceDB(ctx, a.DB).Where("id=?", item.ID).Select("*").Omit("created_at").Updates(item)
	return errors.WithStack(result.Error)
}

// Delete the specified application service from the database.
func (a *ApplicationService) Delete(ctx context.Context, id string) error {
	return errors.WithStack(util.SoftDelete(ctx, GetApplicationServiceDB(ctx, a.DB), id))
}
