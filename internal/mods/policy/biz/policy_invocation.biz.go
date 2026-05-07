package biz

import (
	"context"
	"time"

	"github.com/jd-opensource/joylive-dashboard/internal/mods/policy/dal"
	"github.com/jd-opensource/joylive-dashboard/internal/mods/policy/schema"
	"github.com/jd-opensource/joylive-dashboard/pkg/errors"
	"github.com/jd-opensource/joylive-dashboard/pkg/util"
)

// Invocation policy management
type PolicyInvocation struct {
	Trans               *util.Trans
	PolicyInvocationDAL *dal.PolicyInvocation
}

// Query policy invocations from the data access object based on the provided parameters and options.
func (a *PolicyInvocation) Query(ctx context.Context, params schema.PolicyInvocationQueryParam) (*schema.PolicyInvocationQueryResult, error) {
	params.Pagination = false

	result, err := a.PolicyInvocationDAL.Query(ctx, params, schema.PolicyInvocationQueryOptions{
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

// Get the specified policy invocation from the data access object.
func (a *PolicyInvocation) Get(ctx context.Context, id string) (*schema.PolicyInvocation, error) {
	policyInvocation, err := a.PolicyInvocationDAL.Get(ctx, id)
	if err != nil {
		return nil, err
	} else if policyInvocation == nil {
		return nil, errors.NotFound("", "Policy invocation not found")
	}
	return policyInvocation, nil
}

// Create a new policy invocation in the data access object.
func (a *PolicyInvocation) Create(ctx context.Context, formItem *schema.PolicyInvocationForm) (*schema.PolicyInvocation, error) {
	// Check unique key before creating.
	sourceAppID := ""
	if formItem.SourceApplicationID != nil {
		sourceAppID = *formItem.SourceApplicationID
	}
	if exists, err := a.PolicyInvocationDAL.ExistsByUniqueKey(ctx, formItem.Name, formItem.SpaceCode, sourceAppID, formItem.TargetServiceId); err != nil {
		return nil, err
	} else if exists {
		return nil, errors.BadRequest("", "Policy invocation with the same name, space code, source application and target service already exists")
	}

	policyInvocation := &schema.PolicyInvocation{
		ID:        util.NewXID(),
		Deleted:   "0",
		CreatedAt: time.Now(),
	}

	if err := formItem.FillTo(policyInvocation); err != nil {
		return nil, err
	}

	err := a.Trans.Exec(ctx, func(ctx context.Context) error {
		return a.PolicyInvocationDAL.Create(ctx, policyInvocation)
	})
	if err != nil {
		return nil, err
	}
	return policyInvocation, nil
}

// Update the specified policy invocation in the data access object.
func (a *PolicyInvocation) Update(ctx context.Context, id string, formItem *schema.PolicyInvocationForm) error {
	policyInvocation, err := a.PolicyInvocationDAL.Get(ctx, id)
	if err != nil {
		return err
	} else if policyInvocation == nil {
		return errors.NotFound("", "Policy invocation not found")
	}

	// If unique key fields changed, ensure the new combination is not occupied.
	sourceAppID := ""
	if formItem.SourceApplicationID != nil {
		sourceAppID = *formItem.SourceApplicationID
	}
	existingSourceAppID := ""
	if policyInvocation.SourceApplicationID != nil {
		existingSourceAppID = *policyInvocation.SourceApplicationID
	}
	if policyInvocation.Name != formItem.Name || policyInvocation.SpaceCode != formItem.SpaceCode || existingSourceAppID != sourceAppID || policyInvocation.TargetServiceId != formItem.TargetServiceId {
		if exists, err := a.PolicyInvocationDAL.ExistsByUniqueKey(ctx, formItem.Name, formItem.SpaceCode, sourceAppID, formItem.TargetServiceId); err != nil {
			return err
		} else if exists {
			return errors.BadRequest("", "Policy invocation with the same name, space code, source application and target service already exists")
		}
	}

	if err := formItem.FillTo(policyInvocation); err != nil {
		return err
	}
	policyInvocation.UpdatedAt = time.Now()

	return a.Trans.Exec(ctx, func(ctx context.Context) error {
		return a.PolicyInvocationDAL.Update(ctx, policyInvocation)
	})
}

// Delete the specified policy invocation from the data access object.
func (a *PolicyInvocation) Delete(ctx context.Context, id string) error {
	exists, err := a.PolicyInvocationDAL.Exists(ctx, id)
	if err != nil {
		return err
	} else if !exists {
		return errors.NotFound("", "Policy invocation not found")
	}

	return a.Trans.Exec(ctx, func(ctx context.Context) error {
		return a.PolicyInvocationDAL.Delete(ctx, id)
	})
}
