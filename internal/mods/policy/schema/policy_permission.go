package schema

import (
	"time"

	"github.com/jd-opensource/joylive-dashboard/internal/config"
	"github.com/jd-opensource/joylive-dashboard/pkg/errors"
	"github.com/jd-opensource/joylive-dashboard/pkg/util"
	"gorm.io/gorm"
)

// Permission policy management
type PolicyPermission struct {
	ID                  string          `json:"id" gorm:"size:20;primaryKey;<-:create;comment:Unique ID;"`                                                  // Unique ID
	Name                string          `json:"name" gorm:"size:255;not null;uniqueIndex:uniq_policy_permission_name;comment:Policy name;"`                 // Policy name
	SpaceCode           string          `json:"space_code" gorm:"size:255;not null;comment:Microservice space code;"`                                        // Microservice space code
	SourceApplicationID *string         `json:"source_application_id,omitempty" gorm:"size:20;comment:Source application ID;"`                                // Source application ID
	TargetServiceId     string          `json:"target_service_id" gorm:"size:20;not null;uniqueIndex:uniq_policy_permission_name;comment:Target service ID;"` // Target service ID
	Group               string          `json:"group" gorm:"size:255;not null;default:default;comment:Group;"`                                              // Group
	Path                *string         `json:"path,omitempty" gorm:"size:255;comment:Path or interface;"`                                                  // Path or interface
	Method              *string         `json:"method,omitempty" gorm:"size:255;comment:Method;"`                                                           // Method
	RelationType        string          `json:"relation_type" gorm:"size:20;not null;comment:Relation type;"`                                                // Relation type
	Conditions          *string         `json:"conditions,omitempty" gorm:"type:json;comment:Match conditions (JSON);"`                                     // Match conditions (JSON)
	Type                string          `json:"type" gorm:"size:20;not null;comment:Permission type;"`                                                      // Permission type
	Version             int64           `json:"version" gorm:"not null;default:1;comment:Version;"`                                                         // Version
	Enabled             int             `json:"enabled" gorm:"not null;default:0;comment:Enabled;"`                                                         // Enabled
	Description         *string         `json:"description,omitempty" gorm:"size:255;comment:Details;"`                                                     // Details
	Creator             *string         `json:"creator,omitempty" gorm:"size:255;comment:Creator;"`                                                         // Creator
	Modifier            *string         `json:"modifier,omitempty" gorm:"size:255;comment:Modifier;"`                                                       // Modifier
	CreatedAt           time.Time       `json:"created_at" gorm:"autoCreateTime;comment:Create timestamp;"`                                                  // Create timestamp
	UpdatedAt           time.Time       `json:"updated_at,omitempty" gorm:"autoUpdateTime;comment:Update timestamp;"`                                        // Update timestamp
	Deleted             string          `json:"-" gorm:"uniqueIndex:uniq_policy_permission_name;size:20;default:0;comment:Delete flag;"`                    // Delete flag
	DeletedAt           *gorm.DeletedAt `json:"-" gorm:"comment:Delete timestamp;"`                                                                         // Delete timestamp
}

func (a PolicyPermission) TableName() string {
	return config.C.FormatTableName("policy_permission")
}

// Defining the query parameters for the `PolicyPermission` struct.
type PolicyPermissionQueryParam struct {
	util.PaginationParam
	SpaceCode       string `form:"space_code"`        // Microservice space code
	Name            string `form:"name"`              // Policy name (like)
	TargetServiceId string `form:"target_service_id"` // Target service ID
}

// Defining the query options for the `PolicyPermission` struct.
type PolicyPermissionQueryOptions struct {
	util.QueryOptions
}

// Defining the query result for the `PolicyPermission` struct.
type PolicyPermissionQueryResult struct {
	Data       PolicyPermissions
	PageResult *util.PaginationResult
}

// Defining the slice of `PolicyPermission` struct.
type PolicyPermissions []*PolicyPermission

// Defining the data structure for creating a `PolicyPermission` struct.
type PolicyPermissionForm struct {
	Name                string  `json:"name" binding:"required,max=255"`           // Policy name
	SpaceCode           string  `json:"space_code" binding:"required,max=255"`      // Microservice space code
	SourceApplicationID *string `json:"source_application_id"`                       // Source application ID
	TargetServiceId     string  `json:"target_service_id" binding:"required,max=20"` // Target service ID
	Group               string  `json:"group" binding:"required,max=255"`          // Group
	Path                *string `json:"path"`                                      // Path or interface
	Method              *string `json:"method"`                                    // Method
	RelationType        string  `json:"relation_type" binding:"required,max=20"`    // Relation type
	Conditions          *string `json:"conditions"`                                // Match conditions (JSON)
	Type                string  `json:"type" binding:"required,max=20"`            // Permission type
	Version             int64   `json:"version"`                                   // Version
	Enabled             int     `json:"enabled"`                                   // Enabled
	Description         *string `json:"description"`                               // Details
}

// A validation function for the `PolicyPermissionForm` struct.
func (a *PolicyPermissionForm) Validate() error {
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
	if a.RelationType == "" {
		return errors.BadRequest("", "RelationType is required")
	}
	if a.Type == "" {
		return errors.BadRequest("", "Type is required")
	}
	return nil
}

// Convert `PolicyPermissionForm` to `PolicyPermission` object.
func (a *PolicyPermissionForm) FillTo(policyPermission *PolicyPermission) error {
	policyPermission.Name = a.Name
	policyPermission.SpaceCode = a.SpaceCode
	policyPermission.SourceApplicationID = a.SourceApplicationID
	policyPermission.TargetServiceId = a.TargetServiceId
	policyPermission.Group = a.Group
	policyPermission.Path = a.Path
	policyPermission.Method = a.Method
	policyPermission.RelationType = a.RelationType
	policyPermission.Conditions = a.Conditions
	policyPermission.Type = a.Type
	policyPermission.Version = a.Version
	policyPermission.Enabled = a.Enabled
	policyPermission.Description = a.Description
	return nil
}
