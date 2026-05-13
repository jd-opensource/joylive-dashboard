package biz

import (
	"context"
	"time"

	"github.com/jd-opensource/joylive-dashboard/internal/mods/resource/dal"
	"github.com/jd-opensource/joylive-dashboard/internal/mods/resource/schema"
	"github.com/jd-opensource/joylive-dashboard/pkg/errors"
	"github.com/jd-opensource/joylive-dashboard/pkg/util"
)

// Service group management
type ServiceGroup struct {
	Trans           *util.Trans
	ServiceGroupDAL *dal.ServiceGroup
}

var DefaultGroup = &schema.ServiceGroup{
	Name:          "默认分组",
	Code:          "default",
	CreatedAt:     time.Now(),
	UpdatedAt:     time.Now(),
	Description:   util.StringPtr("未指定分组的实例，皆为默认分组；默认分组策略即服务级策略。"),
}

// Query service groups from the data access object based on the provided parameters and options.
func (a *ServiceGroup) Query(ctx context.Context, params schema.ServiceGroupQueryParam) (*schema.ServiceGroupQueryResult, error) {
	params.Pagination = false

	result, err := a.ServiceGroupDAL.Query(ctx, params, schema.ServiceGroupQueryOptions{
		QueryOptions: util.QueryOptions{
			OrderFields: []util.OrderByParam{
				{Field: "created_at", Direction: util.DESC},
			},
		},
	})
	if err != nil {
		return nil, err
	}

	// Inject a virtual default group at the head of the list
	result.Data = append(schema.ServiceGroups{DefaultGroup}, result.Data...)

	return result, nil
}

// Get the specified service group from the data access object.
func (a *ServiceGroup) Get(ctx context.Context, id string) (*schema.ServiceGroup, error) {
	serviceGroup, err := a.ServiceGroupDAL.Get(ctx, id)
	if err != nil {
		return nil, err
	} else if serviceGroup == nil {
		return nil, errors.NotFound("", "Service group not found")
	}
	return serviceGroup, nil
}

// Create a new service group in the data access object.
func (a *ServiceGroup) Create(ctx context.Context, formItem *schema.ServiceGroupForm) (*schema.ServiceGroup, error) {
	// Check unique key (code + service_id) before creating.
	if exists, err := a.ServiceGroupDAL.ExistsByUniqueKey(ctx, formItem.Code, formItem.ServiceId); err != nil {
		return nil, err
	} else if exists {
		return nil, errors.BadRequest("", "Service group with the same code and service already exists")
	}

	serviceGroup := &schema.ServiceGroup{
		ID:        util.NewXID(),
		Deleted:   "0",
		CreatedAt: time.Now(),
	}

	if err := formItem.FillTo(serviceGroup); err != nil {
		return nil, err
	}

	err := a.Trans.Exec(ctx, func(ctx context.Context) error {
		return a.ServiceGroupDAL.Create(ctx, serviceGroup)
	})
	if err != nil {
		return nil, err
	}
	return serviceGroup, nil
}

// Update the specified service group in the data access object.
func (a *ServiceGroup) Update(ctx context.Context, id string, formItem *schema.ServiceGroupForm) error {
	serviceGroup, err := a.ServiceGroupDAL.Get(ctx, id)
	if err != nil {
		return err
	} else if serviceGroup == nil {
		return errors.NotFound("", "Service group not found")
	}

	// If unique key fields changed, ensure the new combination is not occupied.
	if serviceGroup.Code != formItem.Code || serviceGroup.ServiceId != formItem.ServiceId {
		if exists, err := a.ServiceGroupDAL.ExistsByUniqueKey(ctx, formItem.Code, formItem.ServiceId); err != nil {
			return err
		} else if exists {
			return errors.BadRequest("", "Service group with the same code and service already exists")
		}
	}

	if err := formItem.FillTo(serviceGroup); err != nil {
		return err
	}
	serviceGroup.UpdatedAt = time.Now()

	return a.Trans.Exec(ctx, func(ctx context.Context) error {
		return a.ServiceGroupDAL.Update(ctx, serviceGroup)
	})
}

// Delete the specified service group from the data access object.
func (a *ServiceGroup) Delete(ctx context.Context, id string) error {
	exists, err := a.ServiceGroupDAL.Exists(ctx, id)
	if err != nil {
		return err
	} else if !exists {
		return errors.NotFound("", "Service group not found")
	}

	return a.Trans.Exec(ctx, func(ctx context.Context) error {
		return a.ServiceGroupDAL.Delete(ctx, id)
	})
}
