package biz

import (
	"context"
	"time"

	"github.com/jd-opensource/joylive-dashboard/internal/mods/policy/dal"
	"github.com/jd-opensource/joylive-dashboard/internal/mods/policy/schema"
	"github.com/jd-opensource/joylive-dashboard/pkg/errors"
	"github.com/jd-opensource/joylive-dashboard/pkg/util"
)

// Route policy management
type PolicyRoute struct {
	Trans          *util.Trans
	PolicyRouteDAL *dal.PolicyRoute
}

// Query policy routes from the data access object based on the provided parameters and options.
func (a *PolicyRoute) Query(ctx context.Context, params schema.PolicyRouteQueryParam) (*schema.PolicyRouteQueryResult, error) {
	params.Pagination = false

	result, err := a.PolicyRouteDAL.Query(ctx, params, schema.PolicyRouteQueryOptions{
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

// Get the specified policy route from the data access object.
func (a *PolicyRoute) Get(ctx context.Context, id string) (*schema.PolicyRouteForm, error) {
	policyRoute, err := a.PolicyRouteDAL.Get(ctx, id)
	if err != nil {
		return nil, err
	} else if policyRoute == nil {
		return nil, errors.NotFound("", "Policy route not found")
	}
	var policyRouteForm schema.PolicyRouteForm
	if err := policyRoute.ConvertTo(&policyRouteForm); err != nil {
		return nil, err
	}
	return &policyRouteForm, nil
}

// Create a new policy route in the data access object.
func (a *PolicyRoute) Create(ctx context.Context, formItem *schema.PolicyRouteForm) (*schema.PolicyRoute, error) {
	// Check unique key before creating.
	sourceAppID := ""
	if formItem.SourceApplicationID != nil {
		sourceAppID = *formItem.SourceApplicationID
	}
	if exists, err := a.PolicyRouteDAL.ExistsByUniqueKey(ctx, formItem.Name, formItem.SpaceCode, sourceAppID); err != nil {
		return nil, err
	} else if exists {
		return nil, errors.BadRequest("", "Policy route with the same name, space code and source application already exists")
	}

	policyRoute := &schema.PolicyRoute{
		ID:        util.NewXID(),
		Deleted:   "0",
		CreatedAt: time.Now(),
	}

	if err := formItem.FillTo(policyRoute); err != nil {
		return nil, err
	}

	err := a.Trans.Exec(ctx, func(ctx context.Context) error {
		return a.PolicyRouteDAL.Create(ctx, policyRoute)
	})
	if err != nil {
		return nil, err
	}
	return policyRoute, nil
}

// Update the specified policy route in the data access object.
func (a *PolicyRoute) Update(ctx context.Context, id string, formItem *schema.PolicyRouteForm) error {
	policyRoute, err := a.PolicyRouteDAL.Get(ctx, id)
	if err != nil {
		return err
	} else if policyRoute == nil {
		return errors.NotFound("", "Policy route not found")
	}

	// If unique key fields changed, ensure the new combination is not occupied.
	sourceAppID := ""
	if formItem.SourceApplicationID != nil {
		sourceAppID = *formItem.SourceApplicationID
	}
	existingSourceAppID := ""
	if policyRoute.SourceApplicationID != nil {
		existingSourceAppID = *policyRoute.SourceApplicationID
	}
	if policyRoute.Name != formItem.Name || policyRoute.SpaceCode != formItem.SpaceCode || existingSourceAppID != sourceAppID {
		if exists, err := a.PolicyRouteDAL.ExistsByUniqueKey(ctx, formItem.Name, formItem.SpaceCode, sourceAppID); err != nil {
			return err
		} else if exists {
			return errors.BadRequest("", "Policy route with the same name, space code and source application already exists")
		}
	}

	if err := formItem.FillTo(policyRoute); err != nil {
		return err
	}
	policyRoute.UpdatedAt = time.Now()

	return a.Trans.Exec(ctx, func(ctx context.Context) error {
		return a.PolicyRouteDAL.Update(ctx, policyRoute)
	})
}

// Delete the specified policy route from the data access object.
func (a *PolicyRoute) Delete(ctx context.Context, id string) error {
	exists, err := a.PolicyRouteDAL.Exists(ctx, id)
	if err != nil {
		return err
	} else if !exists {
		return errors.NotFound("", "Policy route not found")
	}

	return a.Trans.Exec(ctx, func(ctx context.Context) error {
		return a.PolicyRouteDAL.Delete(ctx, id)
	})
}
