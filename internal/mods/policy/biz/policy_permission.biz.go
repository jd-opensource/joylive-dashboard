package biz

import (
	"context"
	"time"

	"github.com/jd-opensource/joylive-dashboard/internal/mods/policy/dal"
	"github.com/jd-opensource/joylive-dashboard/internal/mods/policy/schema"
	"github.com/jd-opensource/joylive-dashboard/pkg/errors"
	"github.com/jd-opensource/joylive-dashboard/pkg/util"
)

// Permission policy management
type PolicyPermission struct {
	Trans               *util.Trans
	PolicyPermissionDAL *dal.PolicyPermission
}

// Query policy permissions from the data access object based on the provided parameters and options.
func (a *PolicyPermission) Query(ctx context.Context, params schema.PolicyPermissionQueryParam) (*schema.PolicyPermissionQueryResult, error) {
	params.Pagination = false

	result, err := a.PolicyPermissionDAL.Query(ctx, params, schema.PolicyPermissionQueryOptions{
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

// Get the specified policy permission from the data access object.
func (a *PolicyPermission) Get(ctx context.Context, id string) (*schema.PolicyPermission, error) {
	policyPermission, err := a.PolicyPermissionDAL.Get(ctx, id)
	if err != nil {
		return nil, err
	} else if policyPermission == nil {
		return nil, errors.NotFound("", "Policy permission not found")
	}
	return policyPermission, nil
}

// Create a new policy permission in the data access object.
func (a *PolicyPermission) Create(ctx context.Context, formItem *schema.PolicyPermissionForm) (*schema.PolicyPermission, error) {
	// Check unique key (name + target_service_id) before creating.
	if exists, err := a.PolicyPermissionDAL.ExistsByUniqueKey(ctx, formItem.Name, formItem.TargetServiceId); err != nil {
		return nil, err
	} else if exists {
		return nil, errors.BadRequest("", "Policy permission with the same name and target service already exists")
	}

	policyPermission := &schema.PolicyPermission{
		ID:        util.NewXID(),
		Deleted:   "0",
		CreatedAt: time.Now(),
	}

	if err := formItem.FillTo(policyPermission); err != nil {
		return nil, err
	}

	err := a.Trans.Exec(ctx, func(ctx context.Context) error {
		return a.PolicyPermissionDAL.Create(ctx, policyPermission)
	})
	if err != nil {
		return nil, err
	}
	return policyPermission, nil
}

// Update the specified policy permission in the data access object.
func (a *PolicyPermission) Update(ctx context.Context, id string, formItem *schema.PolicyPermissionForm) error {
	policyPermission, err := a.PolicyPermissionDAL.Get(ctx, id)
	if err != nil {
		return err
	} else if policyPermission == nil {
		return errors.NotFound("", "Policy permission not found")
	}

	// If unique key fields changed, ensure the new combination is not occupied.
	if policyPermission.Name != formItem.Name || policyPermission.TargetServiceId != formItem.TargetServiceId {
		if exists, err := a.PolicyPermissionDAL.ExistsByUniqueKey(ctx, formItem.Name, formItem.TargetServiceId); err != nil {
			return err
		} else if exists {
			return errors.BadRequest("", "Policy permission with the same name and target service already exists")
		}
	}

	if err := formItem.FillTo(policyPermission); err != nil {
		return err
	}
	policyPermission.UpdatedAt = time.Now()

	return a.Trans.Exec(ctx, func(ctx context.Context) error {
		return a.PolicyPermissionDAL.Update(ctx, policyPermission)
	})
}

// Delete the specified policy permission from the data access object.
func (a *PolicyPermission) Delete(ctx context.Context, id string) error {
	exists, err := a.PolicyPermissionDAL.Exists(ctx, id)
	if err != nil {
		return err
	} else if !exists {
		return errors.NotFound("", "Policy permission not found")
	}

	return a.Trans.Exec(ctx, func(ctx context.Context) error {
		return a.PolicyPermissionDAL.Delete(ctx, id)
	})
}
