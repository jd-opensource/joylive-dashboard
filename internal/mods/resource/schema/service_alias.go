package schema

import (
	"time"

	"github.com/jd-opensource/joylive-dashboard/internal/config"
	"github.com/jd-opensource/joylive-dashboard/pkg/errors"
	"github.com/jd-opensource/joylive-dashboard/pkg/util"
	"gorm.io/gorm"
)

// Service alias management
type ServiceAlias struct {
	ID          string          `json:"id" gorm:"size:20;primaryKey;<-:create;comment:Unique ID;"`                                               // Unique ID
	SpaceCode   string          `json:"space_code" gorm:"size:255;not null;uniqueIndex:uniq_service_alias_code;comment:Microservice space code;"` // Microservice space code
	Alias       string          `json:"alias" gorm:"size:255;not null;uniqueIndex:uniq_service_alias_code;comment:Service alias;"`               // Service alias
	ServiceId   string          `json:"service_id" gorm:"size:20;not null;comment:Service ID this alias belongs to;"`                             // Service ID this alias belongs to
	Description *string         `json:"description,omitempty" gorm:"size:255;comment:Details;"`                                                  // Details
	CreatedAt   time.Time       `json:"created_at" gorm:"autoCreateTime;comment:Create timestamp;"`                                               // Create timestamp
	UpdatedAt   time.Time       `json:"updated_at,omitempty" gorm:"autoUpdateTime;comment:Update timestamp;"`                                     // Update timestamp
	Deleted     string          `json:"-" gorm:"uniqueIndex:uniq_service_alias_code;size:20;default:0;comment:Delete flag;"`                     // Delete flag
	DeletedAt   *gorm.DeletedAt `json:"-" gorm:"comment:Delete timestamp;"`                                                                      // Delete timestamp
}

func (a ServiceAlias) TableName() string {
	return config.C.FormatTableName("service_alias")
}

// Defining the query parameters for the `ServiceAlias` struct.
type ServiceAliasQueryParam struct {
	util.PaginationParam
	SpaceCode string `form:"space_code"` // Microservice space code
	Alias     string `form:"alias"`      // Service alias
	ServiceId string `form:"service_id"` // Service ID
}

// Defining the query options for the `ServiceAlias` struct.
type ServiceAliasQueryOptions struct {
	util.QueryOptions
}

// Defining the query result for the `ServiceAlias` struct.
type ServiceAliasQueryResult struct {
	Data       ServiceAliases
	PageResult *util.PaginationResult
}

// Defining the slice of `ServiceAlias` struct.
type ServiceAliases []*ServiceAlias

// Defining the data structure for creating a `ServiceAlias` struct.
type ServiceAliasForm struct {
	SpaceCode   string  `json:"space_code" binding:"required,max=255"` // Microservice space code
	Alias       string  `json:"alias" binding:"required,max=255"`     // Service alias
	ServiceId   string  `json:"service_id" binding:"required,max=20"`  // Service ID this alias belongs to
	Description *string `json:"description"`                          // Details
}

// A validation function for the `ServiceAliasForm` struct.
func (a *ServiceAliasForm) Validate() error {
	if a.SpaceCode == "" {
		return errors.BadRequest("", "SpaceCode is required")
	}
	if a.Alias == "" {
		return errors.BadRequest("", "Alias is required")
	}
	if a.ServiceId == "" {
		return errors.BadRequest("", "ServiceId is required")
	}
	return nil
}

// Convert `ServiceAliasForm` to `ServiceAlias` object.
func (a *ServiceAliasForm) FillTo(serviceAlias *ServiceAlias) error {
	serviceAlias.SpaceCode = a.SpaceCode
	serviceAlias.Alias = a.Alias
	serviceAlias.ServiceId = a.ServiceId
	serviceAlias.Description = a.Description
	return nil
}
