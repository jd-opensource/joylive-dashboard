package schema

import (
	"time"

	"github.com/jd-opensource/joylive-dashboard/internal/config"
	"github.com/jd-opensource/joylive-dashboard/pkg/util"
)

// Application management
type Application struct {
	ID          string    `json:"id" gorm:"size:20;primarykey;"`                  // Unique ID
	Name        string    `json:"name" gorm:"size:255;uniqueIndex:uniq_app_name"` // Application name
	Alias       string    `json:"alias" gorm:"size:255"`                          // Application Chinese name
	Language    string    `json:"language" gorm:"size:20"`                         // Language: Java, Python, Golang
	Enhance     string    `json:"enhance" gorm:"size:20"`                          // Enhance method: Agent, Sidecar
	Source      string    `json:"source" gorm:"size:255"`                          // Data source: Local, JSF, JDAP
	Tenant      string    `json:"tenant" gorm:"size:255"`                          // Tenant
	Creator     string    `json:"creator" gorm:"size:255"`                         // Creator
	Description string    `json:"description" gorm:"size:255"`                     // Description
	CreatedAt   time.Time `json:"created_at" gorm:"index;"`                        // Create time
	UpdatedAt   time.Time `json:"updated_at" gorm:"index;"`                        // Update time
	Deleted     string    `json:"-" gorm:"size:20;default:0"`                      // Logical delete flag
	DeletedAt   *time.Time `json:"-"`                                              // Delete time
}

func (a *Application) TableName() string {
	return config.C.FormatTableName("application")
}

// ApplicationQueryParam defines the query parameters for Application.
type ApplicationQueryParam struct {
	util.PaginationParam
	LikeName string `form:"name"` // Name (like)
}

// ApplicationQueryOptions defines the query options for Application.
type ApplicationQueryOptions struct {
	util.QueryOptions
}

// ApplicationQueryResult defines the query result for Application.
type ApplicationQueryResult struct {
	Data       Applications
	PageResult *util.PaginationResult
}

// Applications defines a slice of Application.
type Applications []*Application

// ApplicationForm defines the form for creating/updating an Application.
type ApplicationForm struct {
	Name        string `json:"name" binding:"required,max=255"`        // Application name
	Alias       string `json:"alias"`                                  // Application Chinese name
	Language    string `json:"language"`                               // Language: Java, Python, Golang
	Enhance     string `json:"enhance" binding:"required,max=20"`      // Enhance method: Agent, Sidecar
	Source      string `json:"source" binding:"required,max=255"`      // Data source: Local, JSF, JDAP
	Description string `json:"description"`                            // Description
}

func (a *ApplicationForm) Validate() error {
	return nil
}

func (a *ApplicationForm) FillTo(app *Application) error {
	app.Name = a.Name
	app.Alias = a.Alias
	app.Language = a.Language
	app.Enhance = a.Enhance
	app.Source = a.Source
	app.Description = a.Description
	return nil
}
