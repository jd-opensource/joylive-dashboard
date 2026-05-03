package biz

import (
	"context"
	"time"

	"github.com/jd-opensource/joylive-dashboard/internal/mods/resource/dal"
	"github.com/jd-opensource/joylive-dashboard/internal/mods/resource/schema"
	"github.com/jd-opensource/joylive-dashboard/pkg/errors"
	"github.com/jd-opensource/joylive-dashboard/pkg/util"
)

// Service business logic layer
type Service struct {
	Trans      *util.Trans
	ServiceDAL *dal.Service
}

// Query services.
func (a *Service) Query(ctx context.Context, params schema.ServiceQueryParam) (*schema.ServiceQueryResult, error) {
	params.Pagination = true

	result, err := a.ServiceDAL.Query(ctx, params, schema.ServiceQueryOptions{
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

// Get the specified service.
func (a *Service) Get(ctx context.Context, id string) (*schema.Service, error) {
	svc, err := a.ServiceDAL.Get(ctx, id)
	if err != nil {
		return nil, err
	} else if svc == nil {
		return nil, errors.NotFound("", "Service not found")
	}
	return svc, nil
}

// Create a new service.
func (a *Service) Create(ctx context.Context, formItem *schema.ServiceForm) (*schema.Service, error) {
	if exists, err := a.ServiceDAL.ExistsName(ctx, formItem.Name, formItem.SpaceCode); err != nil {
		return nil, err
	} else if exists {
		return nil, errors.BadRequest("", "Service name already exists in this space")
	}

	svc := &schema.Service{
		ID:        util.NewXID(),
		Creator:   util.FromUsername(ctx),
		Version:   1,
		CreatedAt: time.Now(),
	}
	if err := formItem.FillTo(svc); err != nil {
		return nil, err
	}

	err := a.Trans.Exec(ctx, func(ctx context.Context) error {
		return a.ServiceDAL.Create(ctx, svc)
	})
	if err != nil {
		return nil, err
	}
	return svc, nil
}

// Update the specified service.
func (a *Service) Update(ctx context.Context, id string, formItem *schema.ServiceForm) error {
	svc, err := a.ServiceDAL.Get(ctx, id)
	if err != nil {
		return err
	} else if svc == nil {
		return errors.NotFound("", "Service not found")
	} else if svc.Name != formItem.Name || svc.SpaceCode != formItem.SpaceCode {
		if exists, err := a.ServiceDAL.ExistsName(ctx, formItem.Name, formItem.SpaceCode); err != nil {
			return err
		} else if exists {
			return errors.BadRequest("", "Service name already exists in this space")
		}
	}

	if err := formItem.FillTo(svc); err != nil {
		return err
	}
	svc.UpdatedAt = time.Now()

	return a.Trans.Exec(ctx, func(ctx context.Context) error {
		return a.ServiceDAL.Update(ctx, svc)
	})
}

// Delete the specified service.
func (a *Service) Delete(ctx context.Context, id string) error {
	exists, err := a.ServiceDAL.Exists(ctx, id)
	if err != nil {
		return err
	} else if !exists {
		return errors.NotFound("", "Service not found")
	}

	return a.Trans.Exec(ctx, func(ctx context.Context) error {
		return a.ServiceDAL.Delete(ctx, id)
	})
}
