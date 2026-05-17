package biz

import (
	"context"
	"time"

	"github.com/jd-opensource/joylive-dashboard/internal/mods/policy/dal"
	"github.com/jd-opensource/joylive-dashboard/internal/mods/policy/schema"
	"github.com/jd-opensource/joylive-dashboard/pkg/errors"
	"github.com/jd-opensource/joylive-dashboard/pkg/util"
)

// Fault injection policy management
type PolicyFault struct {
	Trans          *util.Trans
	PolicyFaultDAL *dal.PolicyFault
}

// Query policy faults from the data access object based on the provided parameters and options.
func (a *PolicyFault) Query(ctx context.Context, params schema.PolicyFaultQueryParam) (*schema.PolicyFaultQueryResult, error) {
	params.Pagination = false

	result, err := a.PolicyFaultDAL.Query(ctx, params, schema.PolicyFaultQueryOptions{
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

// Get the specified policy fault from the data access object.
func (a *PolicyFault) Get(ctx context.Context, id string) (*schema.PolicyFaultForm, error) {
	policyFault, err := a.PolicyFaultDAL.Get(ctx, id)
	if err != nil {
		return nil, err
	} else if policyFault == nil {
		return nil, errors.NotFound("", "Policy fault not found")
	}
	var form schema.PolicyFaultForm
	if err := policyFault.ConvertTo(&form); err != nil {
		return nil, err
	}
	return &form, nil
}

// Create a new policy fault in the data access object.
func (a *PolicyFault) Create(ctx context.Context, formItem *schema.PolicyFaultForm) (*schema.PolicyFault, error) {
	// Validate: method requires path
	if formItem.Method != nil && *formItem.Method != "" && (formItem.Path == nil || *formItem.Path == "") {
		return nil, errors.BadRequest("", "Service path is required when method is specified")
	}

	// Check unique key before creating.
	sourceAppID := ""
	if formItem.SourceApplicationID != nil {
		sourceAppID = *formItem.SourceApplicationID
	}
	if exists, err := a.PolicyFaultDAL.ExistsByUniqueKey(ctx, formItem.Name, formItem.SpaceCode, sourceAppID, formItem.TargetServiceId); err != nil {
		return nil, err
	} else if exists {
		return nil, errors.BadRequest("", "Policy fault with the same name, space code, source application and target service already exists")
	}

	creator := util.FromUsername(ctx)
	policyFault := &schema.PolicyFault{
		ID:        util.NewXID(),
		Deleted:   "0",
		Creator:   &creator,
		CreatedAt: time.Now(),
	}

	if err := formItem.FillTo(policyFault); err != nil {
		return nil, err
	}

	err := a.Trans.Exec(ctx, func(ctx context.Context) error {
		return a.PolicyFaultDAL.Create(ctx, policyFault)
	})
	if err != nil {
		return nil, err
	}
	return policyFault, nil
}

// Update the specified policy fault in the data access object.
func (a *PolicyFault) Update(ctx context.Context, id string, formItem *schema.PolicyFaultForm) error {
	policyFault, err := a.PolicyFaultDAL.Get(ctx, id)
	if err != nil {
		return err
	} else if policyFault == nil {
		return errors.NotFound("", "Policy fault not found")
	}

	// Validate: method requires path
	if formItem.Method != nil && *formItem.Method != "" && (formItem.Path == nil || *formItem.Path == "") {
		return errors.BadRequest("", "Service path is required when method is specified")
	}

	// If unique key fields changed, ensure the new combination is not occupied.
	sourceAppID := ""
	if formItem.SourceApplicationID != nil {
		sourceAppID = *formItem.SourceApplicationID
	}
	existingSourceAppID := ""
	if policyFault.SourceApplicationID != nil {
		existingSourceAppID = *policyFault.SourceApplicationID
	}
	if policyFault.Name != formItem.Name || policyFault.SpaceCode != formItem.SpaceCode || existingSourceAppID != sourceAppID || policyFault.TargetServiceId != formItem.TargetServiceId {
		if exists, err := a.PolicyFaultDAL.ExistsByUniqueKey(ctx, formItem.Name, formItem.SpaceCode, sourceAppID, formItem.TargetServiceId); err != nil {
			return err
		} else if exists {
			return errors.BadRequest("", "Policy fault with the same name, space code, source application and target service already exists")
		}
	}

	if err := formItem.FillTo(policyFault); err != nil {
		return err
	}
	modifier := util.FromUsername(ctx)
	policyFault.Modifier = &modifier
	policyFault.UpdatedAt = time.Now()

	return a.Trans.Exec(ctx, func(ctx context.Context) error {
		return a.PolicyFaultDAL.Update(ctx, policyFault)
	})
}

// Delete the specified policy fault from the data access object.
func (a *PolicyFault) Delete(ctx context.Context, id string) error {
	exists, err := a.PolicyFaultDAL.Exists(ctx, id)
	if err != nil {
		return err
	} else if !exists {
		return errors.NotFound("", "Policy fault not found")
	}

	return a.Trans.Exec(ctx, func(ctx context.Context) error {
		return a.PolicyFaultDAL.Delete(ctx, id)
	})
}
