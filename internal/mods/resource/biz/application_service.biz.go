package biz

import (
	"context"
	"time"

	"github.com/jd-opensource/joylive-dashboard/internal/mods/resource/dal"
	"github.com/jd-opensource/joylive-dashboard/internal/mods/resource/schema"
	"github.com/jd-opensource/joylive-dashboard/pkg/errors"
	"github.com/jd-opensource/joylive-dashboard/pkg/util"
)

// Application and service relationship management
type ApplicationService struct {
	Trans                 *util.Trans
	ApplicationServiceDAL *dal.ApplicationService
}

// Query application services from the data access object based on the provided parameters and options.
func (a *ApplicationService) Query(ctx context.Context, params schema.ApplicationServiceQueryParam) (*schema.ApplicationServiceQueryResult, error) {
	params.Pagination = false

	result, err := a.ApplicationServiceDAL.Query(ctx, params, schema.ApplicationServiceQueryOptions{
		QueryOptions: util.QueryOptions{
			OrderFields: []util.OrderByParam{
				{Field: "created_at", Direction: util.DESC},
			},
		},
	})
	if err != nil {
		return nil, err
	}
	return result, nil
}

// Get the specified application service from the data access object.
func (a *ApplicationService) Get(ctx context.Context, id string) (*schema.ApplicationService, error) {
	applicationService, err := a.ApplicationServiceDAL.Get(ctx, id)
	if err != nil {
		return nil, err
	} else if applicationService == nil {
		return nil, errors.NotFound("", "Application service not found")
	}
	return applicationService, nil
}

// Create a new application service in the data access object.
func (a *ApplicationService) Create(ctx context.Context, formItem *schema.ApplicationServiceForm) (*schema.ApplicationService, error) {
	// Check if the combination already exists
	if exists, err := a.ApplicationServiceDAL.ExistsByUniqueKey(ctx, formItem.ApplicationId, formItem.ServiceId, formItem.Role); err != nil {
		return nil, err
	} else if exists {
		return nil, errors.BadRequest("", "Application service relationship already exists")
	}

	applicationService := &schema.ApplicationService{
		ID:        util.NewXID(),
		Status:    "pending",
		Deleted:   "0",
		CreatedAt: time.Now(),
	}

	if err := formItem.FillTo(applicationService); err != nil {
		return nil, err
	}

	err := a.Trans.Exec(ctx, func(ctx context.Context) error {
		return a.ApplicationServiceDAL.Create(ctx, applicationService)
	})
	if err != nil {
		return nil, err
	}
	return applicationService, nil
}

// Update the specified application service in the data access object.
func (a *ApplicationService) Update(ctx context.Context, id string, formItem *schema.ApplicationServiceForm) error {
	applicationService, err := a.ApplicationServiceDAL.Get(ctx, id)
	if err != nil {
		return err
	} else if applicationService == nil {
		return errors.NotFound("", "Application service not found")
	}

	if err := formItem.FillTo(applicationService); err != nil {
		return err
	}
	applicationService.UpdatedAt = time.Now()

	return a.Trans.Exec(ctx, func(ctx context.Context) error {
		if err := a.ApplicationServiceDAL.Update(ctx, applicationService); err != nil {
			return err
		}
		return nil
	})
}

// Delete the specified application service from the data access object.
func (a *ApplicationService) Delete(ctx context.Context, id string) error {
	exists, err := a.ApplicationServiceDAL.Exists(ctx, id)
	if err != nil {
		return err
	} else if !exists {
		return errors.NotFound("", "Application service not found")
	}

	return a.Trans.Exec(ctx, func(ctx context.Context) error {
		if err := a.ApplicationServiceDAL.Delete(ctx, id); err != nil {
			return err
		}
		return nil
	})
}
