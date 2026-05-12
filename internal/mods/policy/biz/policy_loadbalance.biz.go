package biz

import (
	"context"
	"time"

	"github.com/jd-opensource/joylive-dashboard/internal/mods/policy/dal"
	"github.com/jd-opensource/joylive-dashboard/internal/mods/policy/schema"
	"github.com/jd-opensource/joylive-dashboard/pkg/errors"
	"github.com/jd-opensource/joylive-dashboard/pkg/util"
)

// Loadbalance policy management
type PolicyLoadbalance struct {
	Trans                *util.Trans
	PolicyLoadbalanceDAL *dal.PolicyLoadbalance
}

// Query policy loadbalances from the data access object based on the provided parameters and options.
func (a *PolicyLoadbalance) Query(ctx context.Context, params schema.PolicyLoadbalanceQueryParam) (*schema.PolicyLoadbalanceQueryResult, error) {
	params.Pagination = false

	result, err := a.PolicyLoadbalanceDAL.Query(ctx, params, schema.PolicyLoadbalanceQueryOptions{
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

// Get the specified policy loadbalance from the data access object.
func (a *PolicyLoadbalance) Get(ctx context.Context, id string) (*schema.PolicyLoadbalance, error) {
	policyLoadbalance, err := a.PolicyLoadbalanceDAL.Get(ctx, id)
	if err != nil {
		return nil, err
	} else if policyLoadbalance == nil {
		return nil, errors.NotFound("", "Policy loadbalance not found")
	}
	return policyLoadbalance, nil
}

// Create a new policy loadbalance in the data access object.
func (a *PolicyLoadbalance) Create(ctx context.Context, formItem *schema.PolicyLoadbalanceForm) (*schema.PolicyLoadbalance, error) {
	// Check unique key before creating.
	spaceCode := ""
	if formItem.SpaceCode != nil {
		spaceCode = *formItem.SpaceCode
	}
	sourceAppID := ""
	if formItem.SourceApplicationID != nil {
		sourceAppID = *formItem.SourceApplicationID
	}
	if exists, err := a.PolicyLoadbalanceDAL.ExistsByUniqueKey(ctx, spaceCode, sourceAppID, formItem.TargetServiceId); err != nil {
		return nil, err
	} else if exists {
		return nil, errors.BadRequest("", "Policy loadbalance with the same space code, source application and target service already exists")
	}

	creator := util.FromUsername(ctx)
	policyLoadbalance := &schema.PolicyLoadbalance{
		ID:        util.NewXID(),
		Deleted:   "0",
		Creator:   &creator,
		CreatedAt: time.Now(),
	}

	if err := formItem.FillTo(policyLoadbalance); err != nil {
		return nil, err
	}

	err := a.Trans.Exec(ctx, func(ctx context.Context) error {
		return a.PolicyLoadbalanceDAL.Create(ctx, policyLoadbalance)
	})
	if err != nil {
		return nil, err
	}
	return policyLoadbalance, nil
}

// Update the specified policy loadbalance in the data access object.
func (a *PolicyLoadbalance) Update(ctx context.Context, id string, formItem *schema.PolicyLoadbalanceForm) error {
	policyLoadbalance, err := a.PolicyLoadbalanceDAL.Get(ctx, id)
	if err != nil {
		return err
	} else if policyLoadbalance == nil {
		return errors.NotFound("", "Policy loadbalance not found")
	}

	// If unique key fields changed, ensure the new combination is not occupied.
	spaceCode := ""
	if formItem.SpaceCode != nil {
		spaceCode = *formItem.SpaceCode
	}
	sourceAppID := ""
	if formItem.SourceApplicationID != nil {
		sourceAppID = *formItem.SourceApplicationID
	}
	existingSpaceCode := ""
	if policyLoadbalance.SpaceCode != nil {
		existingSpaceCode = *policyLoadbalance.SpaceCode
	}
	existingSourceAppID := ""
	if policyLoadbalance.SourceApplicationID != nil {
		existingSourceAppID = *policyLoadbalance.SourceApplicationID
	}
	if existingSpaceCode != spaceCode || existingSourceAppID != sourceAppID || policyLoadbalance.TargetServiceId != formItem.TargetServiceId {
		if exists, err := a.PolicyLoadbalanceDAL.ExistsByUniqueKey(ctx, spaceCode, sourceAppID, formItem.TargetServiceId); err != nil {
			return err
		} else if exists {
			return errors.BadRequest("", "Policy loadbalance with the same space code, source application and target service already exists")
		}
	}

	if err := formItem.FillTo(policyLoadbalance); err != nil {
		return err
	}
	modifier := util.FromUsername(ctx)
	policyLoadbalance.Modifier = &modifier
	policyLoadbalance.Version++
	policyLoadbalance.UpdatedAt = time.Now()

	return a.Trans.Exec(ctx, func(ctx context.Context) error {
		return a.PolicyLoadbalanceDAL.Update(ctx, policyLoadbalance)
	})
}

// Delete the specified policy loadbalance from the data access object.
func (a *PolicyLoadbalance) Delete(ctx context.Context, id string) error {
	exists, err := a.PolicyLoadbalanceDAL.Exists(ctx, id)
	if err != nil {
		return err
	} else if !exists {
		return errors.NotFound("", "Policy loadbalance not found")
	}

	return a.Trans.Exec(ctx, func(ctx context.Context) error {
		return a.PolicyLoadbalanceDAL.Delete(ctx, id)
	})
}
