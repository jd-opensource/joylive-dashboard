package biz

import (
	"context"
	"time"

	"github.com/jd-opensource/joylive-dashboard/internal/mods/policy/dal"
	"github.com/jd-opensource/joylive-dashboard/internal/mods/policy/schema"
	"github.com/jd-opensource/joylive-dashboard/pkg/errors"
	"github.com/jd-opensource/joylive-dashboard/pkg/util"
)

// Auth policy management
type PolicyAuth struct {
	Trans         *util.Trans
	PolicyAuthDAL *dal.PolicyAuth
}

// Query policy auths from the data access object based on the provided parameters and options.
func (a *PolicyAuth) Query(ctx context.Context, params schema.PolicyAuthQueryParam) (*schema.PolicyAuthQueryResult, error) {
	params.Pagination = false

	result, err := a.PolicyAuthDAL.Query(ctx, params, schema.PolicyAuthQueryOptions{
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

// Get the specified policy auth from the data access object.
func (a *PolicyAuth) Get(ctx context.Context, id string) (*schema.PolicyAuth, error) {
	policyAuth, err := a.PolicyAuthDAL.Get(ctx, id)
	if err != nil {
		return nil, err
	} else if policyAuth == nil {
		return nil, errors.NotFound("", "Policy auth not found")
	}
	return policyAuth, nil
}

// Create a new policy auth in the data access object.
func (a *PolicyAuth) Create(ctx context.Context, formItem *schema.PolicyAuthForm) (*schema.PolicyAuth, error) {
	policyAuth := &schema.PolicyAuth{
		ID:        util.NewXID(),
		Deleted:   "0",
		CreatedAt: time.Now(),
	}

	if err := formItem.FillTo(policyAuth); err != nil {
		return nil, err
	}

	err := a.Trans.Exec(ctx, func(ctx context.Context) error {
		return a.PolicyAuthDAL.Create(ctx, policyAuth)
	})
	if err != nil {
		return nil, err
	}
	return policyAuth, nil
}

// Update the specified policy auth in the data access object.
func (a *PolicyAuth) Update(ctx context.Context, id string, formItem *schema.PolicyAuthForm) error {
	policyAuth, err := a.PolicyAuthDAL.Get(ctx, id)
	if err != nil {
		return err
	} else if policyAuth == nil {
		return errors.NotFound("", "Policy auth not found")
	}

	if err := formItem.FillTo(policyAuth); err != nil {
		return err
	}
	policyAuth.UpdatedAt = time.Now()

	return a.Trans.Exec(ctx, func(ctx context.Context) error {
		return a.PolicyAuthDAL.Update(ctx, policyAuth)
	})
}

// Delete the specified policy auth from the data access object.
func (a *PolicyAuth) Delete(ctx context.Context, id string) error {
	exists, err := a.PolicyAuthDAL.Exists(ctx, id)
	if err != nil {
		return err
	} else if !exists {
		return errors.NotFound("", "Policy auth not found")
	}

	return a.Trans.Exec(ctx, func(ctx context.Context) error {
		return a.PolicyAuthDAL.Delete(ctx, id)
	})
}
