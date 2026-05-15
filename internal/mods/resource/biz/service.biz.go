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
	Trans                 *util.Trans
	ServiceDAL            *dal.Service
	ApplicationServiceBIZ *ApplicationService
	DataPermissionBIZ     *DataPermission
}

// Query services.
func (a *Service) Query(ctx context.Context, params schema.ServiceQueryParam) (*schema.ServiceQueryResult, error) {
	params.Pagination = true

	// Always carry current login username so DAL can filter by role of user's applications.
	params.Creator = util.FromUsername(ctx)

	if !util.FromIsRootUser(ctx) {
		params.UserID = util.FromUsername(ctx)
		params.Tenant = util.FromTenant(ctx)
	}

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

	// Populate the typed Authorized field for each service in the result.
	for _, svc := range result.Data {
		if extra, err := schema.ParseServiceExtra(svc.Extra); err == nil {
			svc.Authorized = extra.Authorized
		}
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

	if !util.FromIsRootUser(ctx) {
		ok, err := a.DataPermissionBIZ.HasReadPermission(ctx, schema.DataPermissionTypeService, id)
		if err != nil {
			return nil, err
		} else if !ok {
			return nil, errors.NotFound("", "Service not found")
		}
	}

	// Populate the typed Authorized field from extra so callers don't need to parse JSON.
	if extra, err := schema.ParseServiceExtra(svc.Extra); err == nil {
		svc.Authorized = extra.Authorized
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
		Tenant:    util.FromTenant(ctx),
		Creator:   util.FromUsername(ctx),
		Version:   1,
		CreatedAt: time.Now(),
	}
	if err := formItem.FillTo(svc); err != nil {
		return nil, err
	}

	err := a.Trans.Exec(ctx, func(ctx context.Context) error {
		if err := a.ServiceDAL.Create(ctx, svc); err != nil {
			return err
		}
		// Create application_service relationship with role=provider
		appSvcForm := &schema.ApplicationServiceForm{
			ApplicationId: formItem.ApplicationId,
			ServiceId:     svc.ID,
			Role:          "provider",
		}
		if _, err := a.ApplicationServiceBIZ.Create(ctx, appSvcForm); err != nil {
			return err
		}
		return a.DataPermissionBIZ.CreateByOwner(ctx, schema.DataPermissionTypeService, svc.ID, svc.Tenant)
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
	svc.Version++
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
		if err := a.ServiceDAL.Delete(ctx, id); err != nil {
			return err
		}
		return a.DataPermissionBIZ.DeleteByTypeAndDataId(ctx, schema.DataPermissionTypeService, id)
	})
}

// ToggleAuth toggles the authorized flag in the service's extra field.
func (a *Service) ToggleAuth(ctx context.Context, id string, authorized int) error {
	svc, err := a.ServiceDAL.Get(ctx, id)
	if err != nil {
		return err
	} else if svc == nil {
		return errors.NotFound("", "Service not found")
	}

	extra, err := schema.ParseServiceExtra(svc.Extra)
	if err != nil {
		return errors.BadRequest("", "Invalid extra JSON")
	}
	extra.Authorized = &authorized
	svc.Extra, err = extra.MarshalToJSON()
	if err != nil {
		return errors.BadRequest("", "Failed to marshal extra JSON")
	}
	svc.Version++
	svc.UpdatedAt = time.Now()

	return a.Trans.Exec(ctx, func(ctx context.Context) error {
		return a.ServiceDAL.Update(ctx, svc)
	})
}

// DeleteConsumer removes the consumer relationship for the specified service.
func (a *Service) DeleteConsumer(ctx context.Context, serviceId string) error {
	return a.Trans.Exec(ctx, func(ctx context.Context) error {
		return a.ApplicationServiceBIZ.ApplicationServiceDAL.DeleteByServiceIDAndRole(ctx, serviceId, "consumer")
	})
}
