package biz

import (
	"context"
	"time"

	"github.com/jd-opensource/joylive-dashboard/internal/mods/resource/dal"
	"github.com/jd-opensource/joylive-dashboard/internal/mods/resource/schema"
	"github.com/jd-opensource/joylive-dashboard/pkg/errors"
	"github.com/jd-opensource/joylive-dashboard/pkg/util"
)

// Service alias management
type ServiceAlias struct {
	Trans           *util.Trans
	ServiceAliasDAL *dal.ServiceAlias
}

// Query service aliases from the data access object based on the provided parameters and options.
func (a *ServiceAlias) Query(ctx context.Context, params schema.ServiceAliasQueryParam) (*schema.ServiceAliasQueryResult, error) {
	params.Pagination = false

	result, err := a.ServiceAliasDAL.Query(ctx, params, schema.ServiceAliasQueryOptions{
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

// Get the specified service alias from the data access object.
func (a *ServiceAlias) Get(ctx context.Context, id string) (*schema.ServiceAlias, error) {
	serviceAlias, err := a.ServiceAliasDAL.Get(ctx, id)
	if err != nil {
		return nil, err
	} else if serviceAlias == nil {
		return nil, errors.NotFound("", "Service alias not found")
	}
	return serviceAlias, nil
}

// Create a new service alias in the data access object.
func (a *ServiceAlias) Create(ctx context.Context, formItem *schema.ServiceAliasForm) (*schema.ServiceAlias, error) {
	// Check unique key (space_code + alias) before creating.
	if exists, err := a.ServiceAliasDAL.ExistsByUniqueKey(ctx, formItem.SpaceCode, formItem.Alias); err != nil {
		return nil, err
	} else if exists {
		return nil, errors.BadRequest("", "Service alias with the same space code and alias already exists")
	}

	serviceAlias := &schema.ServiceAlias{
		ID:        util.NewXID(),
		Deleted:   "0",
		CreatedAt: time.Now(),
	}

	if err := formItem.FillTo(serviceAlias); err != nil {
		return nil, err
	}

	err := a.Trans.Exec(ctx, func(ctx context.Context) error {
		return a.ServiceAliasDAL.Create(ctx, serviceAlias)
	})
	if err != nil {
		return nil, err
	}
	return serviceAlias, nil
}

// Update the specified service alias in the data access object.
func (a *ServiceAlias) Update(ctx context.Context, id string, formItem *schema.ServiceAliasForm) error {
	serviceAlias, err := a.ServiceAliasDAL.Get(ctx, id)
	if err != nil {
		return err
	} else if serviceAlias == nil {
		return errors.NotFound("", "Service alias not found")
	}

	// If unique key fields changed, ensure the new combination is not occupied.
	if serviceAlias.SpaceCode != formItem.SpaceCode || serviceAlias.Alias != formItem.Alias {
		if exists, err := a.ServiceAliasDAL.ExistsByUniqueKey(ctx, formItem.SpaceCode, formItem.Alias); err != nil {
			return err
		} else if exists {
			return errors.BadRequest("", "Service alias with the same space code and alias already exists")
		}
	}

	if err := formItem.FillTo(serviceAlias); err != nil {
		return err
	}
	serviceAlias.UpdatedAt = time.Now()

	return a.Trans.Exec(ctx, func(ctx context.Context) error {
		return a.ServiceAliasDAL.Update(ctx, serviceAlias)
	})
}

// Delete the specified service alias from the data access object.
func (a *ServiceAlias) Delete(ctx context.Context, id string) error {
	exists, err := a.ServiceAliasDAL.Exists(ctx, id)
	if err != nil {
		return err
	} else if !exists {
		return errors.NotFound("", "Service alias not found")
	}

	return a.Trans.Exec(ctx, func(ctx context.Context) error {
		return a.ServiceAliasDAL.Delete(ctx, id)
	})
}
