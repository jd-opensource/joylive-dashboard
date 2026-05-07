package biz

import (
	"context"
	"time"

	"github.com/jd-opensource/joylive-dashboard/internal/mods/policy/dal"
	"github.com/jd-opensource/joylive-dashboard/internal/mods/policy/schema"
	"github.com/jd-opensource/joylive-dashboard/pkg/errors"
	"github.com/jd-opensource/joylive-dashboard/pkg/util"
)

// Limit policy management
type PolicyLimit struct {
	Trans          *util.Trans
	PolicyLimitDAL *dal.PolicyLimit
}

// Query policy limits from the data access object based on the provided parameters and options.
func (a *PolicyLimit) Query(ctx context.Context, params schema.PolicyLimitQueryParam) (*schema.PolicyLimitQueryResult, error) {
	params.Pagination = false

	result, err := a.PolicyLimitDAL.Query(ctx, params, schema.PolicyLimitQueryOptions{
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

// Get the specified policy limit from the data access object.
func (a *PolicyLimit) Get(ctx context.Context, id string) (*schema.PolicyLimit, error) {
	policyLimit, err := a.PolicyLimitDAL.Get(ctx, id)
	if err != nil {
		return nil, err
	} else if policyLimit == nil {
		return nil, errors.NotFound("", "Policy limit not found")
	}
	return policyLimit, nil
}

// Create a new policy limit in the data access object.
func (a *PolicyLimit) Create(ctx context.Context, formItem *schema.PolicyLimitForm) (*schema.PolicyLimit, error) {
	// Check unique key (name + target_service_id) before creating.
	if exists, err := a.PolicyLimitDAL.ExistsByUniqueKey(ctx, formItem.Name, formItem.TargetServiceId); err != nil {
		return nil, err
	} else if exists {
		return nil, errors.BadRequest("", "Policy limit with the same name and target service already exists")
	}

	policyLimit := &schema.PolicyLimit{
		ID:        util.NewXID(),
		Deleted:   "0",
		CreatedAt: time.Now(),
	}

	if err := formItem.FillTo(policyLimit); err != nil {
		return nil, err
	}

	err := a.Trans.Exec(ctx, func(ctx context.Context) error {
		return a.PolicyLimitDAL.Create(ctx, policyLimit)
	})
	if err != nil {
		return nil, err
	}
	return policyLimit, nil
}

// Update the specified policy limit in the data access object.
func (a *PolicyLimit) Update(ctx context.Context, id string, formItem *schema.PolicyLimitForm) error {
	policyLimit, err := a.PolicyLimitDAL.Get(ctx, id)
	if err != nil {
		return err
	} else if policyLimit == nil {
		return errors.NotFound("", "Policy limit not found")
	}

	// If unique key fields changed, ensure the new combination is not occupied.
	if policyLimit.Name != formItem.Name || policyLimit.TargetServiceId != formItem.TargetServiceId {
		if exists, err := a.PolicyLimitDAL.ExistsByUniqueKey(ctx, formItem.Name, formItem.TargetServiceId); err != nil {
			return err
		} else if exists {
			return errors.BadRequest("", "Policy limit with the same name and target service already exists")
		}
	}

	if err := formItem.FillTo(policyLimit); err != nil {
		return err
	}
	policyLimit.UpdatedAt = time.Now()

	return a.Trans.Exec(ctx, func(ctx context.Context) error {
		return a.PolicyLimitDAL.Update(ctx, policyLimit)
	})
}

// Delete the specified policy limit from the data access object.
func (a *PolicyLimit) Delete(ctx context.Context, id string) error {
	exists, err := a.PolicyLimitDAL.Exists(ctx, id)
	if err != nil {
		return err
	} else if !exists {
		return errors.NotFound("", "Policy limit not found")
	}

	return a.Trans.Exec(ctx, func(ctx context.Context) error {
		return a.PolicyLimitDAL.Delete(ctx, id)
	})
}
