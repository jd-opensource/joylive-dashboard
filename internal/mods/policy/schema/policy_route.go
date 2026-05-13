package schema

import (
	"time"

	"github.com/jd-opensource/joylive-dashboard/internal/config"
	"github.com/jd-opensource/joylive-dashboard/pkg/errors"
	"github.com/jd-opensource/joylive-dashboard/pkg/util"
	"gorm.io/gorm"
)

// Route policy management
type PolicyRoute struct {
	ID                  string          `json:"id" gorm:"size:20;primaryKey;<-:create;comment:Unique ID;"`                                                      // Unique ID
	Name                string          `json:"name" gorm:"size:100;not null;uniqueIndex:uniq_policy_route_name;comment:Policy name;"`                          // Policy name
	SpaceCode           string          `json:"space_code" gorm:"size:255;not null;uniqueIndex:uniq_policy_route_name;comment:Microservice space code;"`         // Microservice space code
	SourceApplicationID *string         `json:"source_application_id,omitempty" gorm:"size:20;uniqueIndex:uniq_policy_route_name;comment:Source application ID;"` // Source application ID
	TargetServiceId     string          `json:"target_service_id" gorm:"size:20;not null;comment:Target service ID;"`                                             // Target service ID
	Group               string          `json:"group" gorm:"size:255;not null;default:default;comment:Group;"`                                                  // Group
	Path                *string         `json:"path,omitempty" gorm:"size:255;comment:Path or interface;"`                                                      // Path or interface
	Method              *string         `json:"method,omitempty" gorm:"size:255;comment:Method;"`                                                               // Method
	Order               int             `json:"order" gorm:"not null;default:0;comment:Sort order;"`                                                            // Sort order
	Version             int64           `json:"version" gorm:"not null;default:1;comment:Version;"`                                                             // Version
	Enabled             int             `json:"enabled" gorm:"not null;default:0;comment:Enabled;"`                                                             // Enabled
	Description         *string         `json:"description,omitempty" gorm:"size:255;comment:Details;"`                                                         // Details
	Creator             *string         `json:"creator,omitempty" gorm:"size:255;comment:Creator;"`                                                             // Creator
	Modifier            *string         `json:"modifier,omitempty" gorm:"size:255;comment:Modifier;"`                                                           // Modifier
	CreatedAt           time.Time       `json:"created_at" gorm:"autoCreateTime;comment:Create timestamp;"`                                                      // Create timestamp
	UpdatedAt           time.Time       `json:"updated_at,omitempty" gorm:"autoUpdateTime;comment:Update timestamp;"`                                            // Update timestamp
	Deleted             string          `json:"-" gorm:"uniqueIndex:uniq_policy_route_name;size:20;default:0;comment:Delete flag;"`                             // Delete flag
	DeletedAt           *gorm.DeletedAt `json:"-" gorm:"comment:Delete timestamp;"`                                                                             // Delete timestamp
}

func (a PolicyRoute) TableName() string {
	return config.C.FormatTableName("policy_route")
}

// Defining the query parameters for the `PolicyRoute` struct.
type PolicyRouteQueryParam struct {
	util.PaginationParam
	SpaceCode       string `form:"space_code"`        // Microservice space code
	Name            string `form:"name"`              // Policy name (like)
	TargetServiceId string `form:"target_service_id"` // Target service ID
}

// Defining the query options for the `PolicyRoute` struct.
type PolicyRouteQueryOptions struct {
	util.QueryOptions
}

// Defining the query result for the `PolicyRoute` struct.
type PolicyRouteQueryResult struct {
	Data       PolicyRoutes
	PageResult *util.PaginationResult
}

// Defining the slice of `PolicyRoute` struct.
type PolicyRoutes []*PolicyRoute

// Defining the data structure for creating a `PolicyRoute` struct.
type PolicyRouteForm struct {
	Name                string  `json:"name" binding:"required,max=100"`           // Policy name
	SpaceCode           string  `json:"space_code" binding:"required,max=255"`      // Microservice space code
	SourceApplicationID *string `json:"source_application_id"`                       // Source application ID
	TargetServiceId     string  `json:"target_service_id" binding:"required,max=20"` // Target service ID
	Group               string  `json:"group" binding:"required,max=255"`          // Group
	Path                *string `json:"path"`                                      // Path or interface
	Method              *string `json:"method"`                                    // Method
	Order               int     `json:"order"`                                     // Sort order
	Version             int64   `json:"version"`                                   // Version
	Enabled             int     `json:"enabled"`                                   // Enabled
	Description         *string `json:"description"`                               // Details
}

// A validation function for the `PolicyRouteForm` struct.
func (a *PolicyRouteForm) Validate() error {
	if a.Name == "" {
		return errors.BadRequest("", "Name is required")
	}
	if a.SpaceCode == "" {
		return errors.BadRequest("", "SpaceCode is required")
	}
	if a.TargetServiceId == "" {
		return errors.BadRequest("", "TargetServiceId is required")
	}
	if a.Group == "" {
		return errors.BadRequest("", "Group is required")
	}
	return nil
}

// Convert `PolicyRouteForm` to `PolicyRoute` object.
func (a *PolicyRouteForm) FillTo(policyRoute *PolicyRoute) error {
	policyRoute.Name = a.Name
	policyRoute.SpaceCode = a.SpaceCode
	policyRoute.SourceApplicationID = a.SourceApplicationID
	policyRoute.TargetServiceId = a.TargetServiceId
	policyRoute.Group = a.Group
	policyRoute.Path = a.Path
	policyRoute.Method = a.Method
	policyRoute.Order = a.Order
	policyRoute.Version = a.Version
	policyRoute.Enabled = a.Enabled
	policyRoute.Description = a.Description
	return nil
}
