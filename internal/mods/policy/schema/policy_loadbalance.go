package schema

import (
	"time"

	"github.com/jd-opensource/joylive-dashboard/internal/config"
	"github.com/jd-opensource/joylive-dashboard/pkg/errors"
	"github.com/jd-opensource/joylive-dashboard/pkg/util"
	"gorm.io/gorm"
)

// Loadbalance policy management
type PolicyLoadbalance struct {
	ID                  string          `json:"id" gorm:"size:20;primaryKey;<-:create;comment:Unique ID;"`                                                                                                          // Unique ID
	Name                string          `json:"name" gorm:"size:100;not null;uniqueIndex:uniq_policy_loadbalance_name;comment:Policy name;"`                                                                        // Policy name
	SpaceCode           *string         `json:"space_code,omitempty" gorm:"size:255;uniqueIndex:uniq_policy_loadbalance_name;uniqueIndex:uniq_policy_loadbalance_service;comment:Microservice space code;"`         // Microservice space code
	SourceApplicationID *string         `json:"source_application_id,omitempty" gorm:"size:20;uniqueIndex:uniq_policy_loadbalance_name;uniqueIndex:uniq_policy_loadbalance_service;comment:Source application ID;"` // Source application ID
	TargetServiceId     string          `json:"target_service_id" gorm:"size:20;not null;uniqueIndex:uniq_policy_loadbalance_name;uniqueIndex:uniq_policy_loadbalance_service;comment:Target service ID;"`          // Target service ID
	Group               string          `json:"group" gorm:"size:255;not null;default:default;comment:Group;"`                                                                                                      // Group
	Path                *string         `json:"path,omitempty" gorm:"size:255;comment:Path or interface;"`                                                                                                          // Path or interface
	Method              *string         `json:"method,omitempty" gorm:"size:255;comment:Method;"`                                                                                                                   // Method
	PolicyType          string          `json:"policy_type" gorm:"size:20;not null;comment:Loadbalance policy type;"`                                                                                               // Loadbalance policy type
	StickyType          *string         `json:"sticky_type,omitempty" gorm:"size:20;comment:Sticky type;"`                                                                                                          // Sticky type
	Version             int64           `json:"version" gorm:"not null;default:1;comment:Version;"`                                                                                                                 // Version
	Enabled             int             `json:"enabled" gorm:"not null;default:0;comment:Enabled;"`                                                                                                                 // Enabled
	Description         *string         `json:"description,omitempty" gorm:"size:255;comment:Details;"`                                                                                                             // Details
	Creator             *string         `json:"creator,omitempty" gorm:"size:255;comment:Creator;"`                                                                                                                 // Creator
	Modifier            *string         `json:"modifier,omitempty" gorm:"size:255;comment:Modifier;"`                                                                                                               // Modifier
	CreatedAt           time.Time       `json:"created_at" gorm:"autoCreateTime;comment:Create timestamp;"`                                                                                                         // Create timestamp
	UpdatedAt           time.Time       `json:"updated_at,omitempty" gorm:"autoUpdateTime;comment:Update timestamp;"`                                                                                               // Update timestamp
	Deleted             string          `json:"-" gorm:"uniqueIndex:uniq_policy_loadbalance_name;uniqueIndex:uniq_policy_loadbalance_service;size:20;default:0;comment:Delete flag;"`                               // Delete flag
	DeletedAt           *gorm.DeletedAt `json:"-" gorm:"comment:Delete timestamp;"`                                                                                                                                 // Delete timestamp
}

func (a PolicyLoadbalance) TableName() string {
	return config.C.FormatTableName("policy_loadbalance")
}

// Defining the query parameters for the `PolicyLoadbalance` struct.
type PolicyLoadbalanceQueryParam struct {
	util.PaginationParam
	SpaceCode       string `form:"space_code"`        // Microservice space code
	Name            string `form:"name"`              // Policy name (like)
	TargetServiceId string `form:"target_service_id"` // Target service ID
}

// Defining the query options for the `PolicyLoadbalance` struct.
type PolicyLoadbalanceQueryOptions struct {
	util.QueryOptions
}

// Defining the query result for the `PolicyLoadbalance` struct.
type PolicyLoadbalanceQueryResult struct {
	Data       PolicyLoadbalances
	PageResult *util.PaginationResult
}

// Defining the slice of `PolicyLoadbalance` struct.
type PolicyLoadbalances []*PolicyLoadbalance

// Defining the data structure for creating a `PolicyLoadbalance` struct.
type PolicyLoadbalanceForm struct {
	Name                string  `json:"name" binding:"required,max=100"`             // Policy name
	SpaceCode           *string `json:"space_code"`                                  // Microservice space code
	SourceApplicationID *string `json:"source_application_id"`                       // Source application ID
	TargetServiceId     string  `json:"target_service_id" binding:"required,max=20"` // Target service ID
	Group               string  `json:"group" binding:"required,max=255"`            // Group
	Path                *string `json:"path"`                                        // Path or interface
	Method              *string `json:"method"`                                      // Method
	PolicyType          string  `json:"policy_type" binding:"required,max=20"`       // Loadbalance policy type
	StickyType          *string `json:"sticky_type"`                                 // Sticky type
	Version             int64   `json:"version"`                                     // Version
	Enabled             int     `json:"enabled"`                                     // Enabled
	Description         *string `json:"description"`                                 // Details
}

// A validation function for the `PolicyLoadbalanceForm` struct.
func (a *PolicyLoadbalanceForm) Validate() error {
	if a.Name == "" {
		return errors.BadRequest("", "Name is required")
	}
	if a.SpaceCode == nil || *a.SpaceCode == "" {
		return errors.BadRequest("", "SpaceCode is required")
	}
	if a.TargetServiceId == "" {
		return errors.BadRequest("", "TargetServiceId is required")
	}
	if a.Group == "" {
		return errors.BadRequest("", "Group is required")
	}
	if a.PolicyType == "" {
		return errors.BadRequest("", "PolicyType is required")
	}
	return nil
}

// Convert `PolicyLoadbalanceForm` to `PolicyLoadbalance` object.
func (a *PolicyLoadbalanceForm) FillTo(policyLoadbalance *PolicyLoadbalance) error {
	policyLoadbalance.Name = a.Name
	policyLoadbalance.SpaceCode = a.SpaceCode
	policyLoadbalance.SourceApplicationID = a.SourceApplicationID
	policyLoadbalance.TargetServiceId = a.TargetServiceId
	policyLoadbalance.Group = a.Group
	policyLoadbalance.Path = a.Path
	policyLoadbalance.Method = a.Method
	policyLoadbalance.PolicyType = a.PolicyType
	policyLoadbalance.StickyType = a.StickyType
	policyLoadbalance.Version = a.Version
	policyLoadbalance.Enabled = a.Enabled
	policyLoadbalance.Description = a.Description
	return nil
}
