package dal

import (
	"context"

	"github.com/jd-opensource/joylive-dashboard/internal/config"
	"github.com/jd-opensource/joylive-dashboard/internal/mods/resource/schema"
	"github.com/jd-opensource/joylive-dashboard/pkg/errors"
	"github.com/jd-opensource/joylive-dashboard/pkg/util"
	"gorm.io/gorm"
)

// GetServiceDB returns the database instance for Service (only active records).
func GetServiceDB(ctx context.Context, defDB *gorm.DB) *gorm.DB {
	tableName := config.C.FormatTableName("service")
	return util.GetDB(ctx, defDB).Model(new(schema.Service)).Preload("Space").Where(tableName + ".deleted = '0'")
}

// Service data access layer
type Service struct {
	DB *gorm.DB
}

// Query services from the database.
func (a *Service) Query(ctx context.Context, params schema.ServiceQueryParam, opts ...schema.ServiceQueryOptions) (*schema.ServiceQueryResult, error) {
	var opt schema.ServiceQueryOptions
	if len(opts) > 0 {
		opt = opts[0]
	}

	tableName := config.C.FormatTableName("service")

	// Build role subquery if needed
	var roleSubQuery *gorm.DB
	if role := params.Role; role != "" {
		if params.UserID != "" {
			appPermQuery := GetDataPermissionDB(ctx, a.DB).
				Where("type = ? AND user = ? AND tenant = ? AND permission & 1 = 1", schema.DataPermissionTypeApplication, params.UserID, params.Tenant).
				Select("data_id")
			roleSubQuery = GetApplicationServiceDB(ctx, a.DB).
				Where("role = ?", role).
				Where("application_id IN (?)", appPermQuery).
				Select("service_id")
		} else {
			roleSubQuery = GetApplicationServiceDB(ctx, a.DB).
				Where("role = ?", role).
				Select("service_id")
		}
	}

	// For consumer role, use a fresh query with JOINs (no Preload to avoid conflict)
	if params.Role == "consumer" {
		appSvcTable := config.C.FormatTableName("application_service")
		appTable := config.C.FormatTableName("application")
		db := util.GetDB(ctx, a.DB).Model(new(schema.Service)).
			Where(tableName+".deleted = '0'").
			Select(
				tableName+".*, "+
					appSvcTable+".status as application_service_status, "+
					appTable+".name as application_name",
			).Joins(
			"LEFT JOIN "+appSvcTable+" ON "+appSvcTable+".service_id = "+tableName+".id AND "+appSvcTable+".role = ? AND "+appSvcTable+".deleted = '0'",
			params.Role,
		).Joins(
			"LEFT JOIN " + appTable + " ON " + appTable + ".id = " + appSvcTable + ".application_id AND " + appTable + ".deleted = '0'",
		)
		if v := params.LikeName; len(v) > 0 {
			db = db.Where(tableName+".name LIKE ?", "%"+v+"%")
		}
		if v := params.SpaceCode; len(v) > 0 {
			db = db.Where(tableName+".space_code = ?", v)
		}
		if roleSubQuery != nil {
			db = db.Where(tableName+".id IN (?)", roleSubQuery)
		}
		var list schema.Services
		pageResult, err := util.WrapPageQuery(ctx, db, params.PaginationParam, opt.QueryOptions, &list)
		if err != nil {
			return nil, errors.WithStack(err)
		}
		return &schema.ServiceQueryResult{
			PageResult: pageResult,
			Data:       list,
		}, nil
	}

	// Non-consumer queries: use GetServiceDB with Preload
	db := GetServiceDB(ctx, a.DB)
	if v := params.LikeName; len(v) > 0 {
		db = db.Where(tableName+".name LIKE ?", "%"+v+"%")
	}
	if v := params.SpaceCode; len(v) > 0 {
		db = db.Where(tableName+".space_code = ?", v)
	}
	if roleSubQuery != nil {
		db = db.Where(tableName+".id IN (?)", roleSubQuery)
	} else if v := params.UserID; len(v) > 0 {
		permQuery := GetDataPermissionDB(ctx, a.DB).
			Where("type = ? AND user = ? AND tenant = ? AND permission & 1 = 1", schema.DataPermissionTypeService, v, params.Tenant).
			Select("data_id")
		db = db.Where(tableName+".id IN (?)", permQuery)
	}

	var list schema.Services
	pageResult, err := util.WrapPageQuery(ctx, db, params.PaginationParam, opt.QueryOptions, &list)
	if err != nil {
		return nil, errors.WithStack(err)
	}
	return &schema.ServiceQueryResult{
		PageResult: pageResult,
		Data:       list,
	}, nil
}

// Get the specified service from the database.
func (a *Service) Get(ctx context.Context, id string, opts ...schema.ServiceQueryOptions) (*schema.Service, error) {
	var opt schema.ServiceQueryOptions
	if len(opts) > 0 {
		opt = opts[0]
	}

	item := new(schema.Service)
	ok, err := util.FindOne(ctx, GetServiceDB(ctx, a.DB).Where("id=?", id), opt.QueryOptions, item)
	if err != nil {
		return nil, errors.WithStack(err)
	} else if !ok {
		return nil, nil
	}
	return item, nil
}

// Exists checks if the specified service exists.
func (a *Service) Exists(ctx context.Context, id string) (bool, error) {
	ok, err := util.Exists(ctx, GetServiceDB(ctx, a.DB).Where("id=?", id))
	return ok, errors.WithStack(err)
}

// ExistsName checks if a service with the given name and space_code exists.
func (a *Service) ExistsName(ctx context.Context, name, spaceCode string) (bool, error) {
	ok, err := util.Exists(ctx, GetServiceDB(ctx, a.DB).Where("name=? AND space_code=?", name, spaceCode))
	return ok, errors.WithStack(err)
}

// Create a new service.
func (a *Service) Create(ctx context.Context, item *schema.Service) error {
	result := GetServiceDB(ctx, a.DB).Create(item)
	return errors.WithStack(result.Error)
}

// Update the specified service.
func (a *Service) Update(ctx context.Context, item *schema.Service) error {
	result := GetServiceDB(ctx, a.DB).Where("id=?", item.ID).Select("*").Omit("created_at").Updates(item)
	return errors.WithStack(result.Error)
}

// Delete the specified service logically.
func (a *Service) Delete(ctx context.Context, id string) error {
	return errors.WithStack(util.SoftDelete(ctx, GetServiceDB(ctx, a.DB), id))
}
