package biz

import (
	"context"
	"time"

	"github.com/jd-opensource/joylive-dashboard/internal/mods/resource/dal"
	"github.com/jd-opensource/joylive-dashboard/internal/mods/resource/schema"
	"github.com/jd-opensource/joylive-dashboard/pkg/errors"
	"github.com/jd-opensource/joylive-dashboard/pkg/util"
)

// Application business logic layer
type Application struct {
	Trans         *util.Trans
	ApplicationDAL *dal.Application
}

// Query applications.
func (a *Application) Query(ctx context.Context, params schema.ApplicationQueryParam) (*schema.ApplicationQueryResult, error) {
	params.Pagination = true

	result, err := a.ApplicationDAL.Query(ctx, params, schema.ApplicationQueryOptions{
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

// Get the specified application.
func (a *Application) Get(ctx context.Context, id string) (*schema.Application, error) {
	app, err := a.ApplicationDAL.Get(ctx, id)
	if err != nil {
		return nil, err
	} else if app == nil {
		return nil, errors.NotFound("", "Application not found")
	}
	return app, nil
}

// Create a new application.
func (a *Application) Create(ctx context.Context, formItem *schema.ApplicationForm) (*schema.Application, error) {
	if exists, err := a.ApplicationDAL.ExistsName(ctx, formItem.Name); err != nil {
		return nil, err
	} else if exists {
		return nil, errors.BadRequest("", "Application name already exists")
	}

	app := &schema.Application{
		ID:        util.NewXID(),
		Creator:   util.FromUsername(ctx),
		CreatedAt: time.Now(),
	}
	if err := formItem.FillTo(app); err != nil {
		return nil, err
	}

	err := a.Trans.Exec(ctx, func(ctx context.Context) error {
		return a.ApplicationDAL.Create(ctx, app)
	})
	if err != nil {
		return nil, err
	}
	return app, nil
}

// Update the specified application.
func (a *Application) Update(ctx context.Context, id string, formItem *schema.ApplicationForm) error {
	app, err := a.ApplicationDAL.Get(ctx, id)
	if err != nil {
		return err
	} else if app == nil {
		return errors.NotFound("", "Application not found")
	} else if app.Name != formItem.Name {
		if exists, err := a.ApplicationDAL.ExistsName(ctx, formItem.Name); err != nil {
			return err
		} else if exists {
			return errors.BadRequest("", "Application name already exists")
		}
	}

	if err := formItem.FillTo(app); err != nil {
		return err
	}
	app.UpdatedAt = time.Now()

	return a.Trans.Exec(ctx, func(ctx context.Context) error {
		return a.ApplicationDAL.Update(ctx, app)
	})
}

// Delete the specified application.
func (a *Application) Delete(ctx context.Context, id string) error {
	exists, err := a.ApplicationDAL.Exists(ctx, id)
	if err != nil {
		return err
	} else if !exists {
		return errors.NotFound("", "Application not found")
	}

	return a.Trans.Exec(ctx, func(ctx context.Context) error {
		return a.ApplicationDAL.Delete(ctx, id)
	})
}
