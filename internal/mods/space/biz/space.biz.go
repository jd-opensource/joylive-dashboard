package biz

import (
	"context"
	"time"

	"github.com/jd-opensource/joylive-dashboard/internal/mods/space/dal"
	"github.com/jd-opensource/joylive-dashboard/internal/mods/space/schema"
	"github.com/jd-opensource/joylive-dashboard/pkg/errors"
	"github.com/jd-opensource/joylive-dashboard/pkg/util"
)

// Space business logic layer
type Space struct {
	Trans    *util.Trans
	SpaceDAL *dal.Space
}

// Query spaces from the data access object based on the provided parameters.
func (a *Space) Query(ctx context.Context, params schema.SpaceQueryParam) (*schema.SpaceQueryResult, error) {
	params.Pagination = true

	result, err := a.SpaceDAL.Query(ctx, params, schema.SpaceQueryOptions{
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

// Get the specified space from the data access object.
func (a *Space) Get(ctx context.Context, id string) (*schema.Space, error) {
	space, err := a.SpaceDAL.Get(ctx, id)
	if err != nil {
		return nil, err
	} else if space == nil {
		return nil, errors.NotFound("", "Space not found")
	}
	return space, nil
}

// Create a new space in the data access object.
func (a *Space) Create(ctx context.Context, formItem *schema.SpaceForm) (*schema.Space, error) {
	if exists, err := a.SpaceDAL.ExistsCode(ctx, formItem.Code); err != nil {
		return nil, err
	} else if exists {
		return nil, errors.BadRequest("", "Space code already exists")
	}

	space := &schema.Space{
		ID:        util.NewXID(),
		Creator:   util.FromUsername(ctx),
		CreatedAt: time.Now(),
	}
	if err := formItem.FillTo(space); err != nil {
		return nil, err
	}

	err := a.Trans.Exec(ctx, func(ctx context.Context) error {
		return a.SpaceDAL.Create(ctx, space)
	})
	if err != nil {
		return nil, err
	}

	return space, nil
}

// Update the specified space in the data access object.
func (a *Space) Update(ctx context.Context, id string, formItem *schema.SpaceForm) error {
	space, err := a.SpaceDAL.Get(ctx, id)
	if err != nil {
		return err
	} else if space == nil {
		return errors.NotFound("", "Space not found")
	} else if space.Code != formItem.Code {
		if exists, err := a.SpaceDAL.ExistsCode(ctx, formItem.Code); err != nil {
			return err
		} else if exists {
			return errors.BadRequest("", "Space code already exists")
		}
	}

	if err := formItem.FillTo(space); err != nil {
		return err
	}
	space.UpdatedAt = time.Now()

	return a.Trans.Exec(ctx, func(ctx context.Context) error {
		return a.SpaceDAL.Update(ctx, space)
	})
}

// Delete the specified space from the data access object.
func (a *Space) Delete(ctx context.Context, id string) error {
	exists, err := a.SpaceDAL.Exists(ctx, id)
	if err != nil {
		return err
	} else if !exists {
		return errors.NotFound("", "Space not found")
	}

	return a.Trans.Exec(ctx, func(ctx context.Context) error {
		return a.SpaceDAL.Delete(ctx, id)
	})
}
