package schema

import (
	"time"

	"github.com/jd-opensource/joylive-dashboard/internal/config"
	"github.com/jd-opensource/joylive-dashboard/pkg/encoding/json"
	"github.com/jd-opensource/joylive-dashboard/pkg/errors"
	"github.com/jd-opensource/joylive-dashboard/pkg/util"
	"gorm.io/gorm"
)

// Invocation policy management
type PolicyInvocation struct {
	ID                  string          `json:"id" gorm:"size:20;primaryKey;<-:create;comment:Unique ID;"`                                                           // Unique ID
	Name                string          `json:"name" gorm:"size:100;not null;uniqueIndex:uniq_policy_invocation_name;comment:Policy name;"`                          // Policy name
	SpaceCode           string          `json:"space_code" gorm:"size:255;not null;uniqueIndex:uniq_policy_invocation_name;comment:Microservice space code;"`         // Microservice space code
	SourceApplicationID *string         `json:"source_application_id,omitempty" gorm:"size:20;uniqueIndex:uniq_policy_invocation_name;comment:Source application ID;"` // Source application ID
	TargetServiceId     string          `json:"target_service_id" gorm:"size:20;not null;uniqueIndex:uniq_policy_invocation_name;comment:Target service ID;"`          // Target service ID
	Group               string          `json:"group" gorm:"size:255;not null;default:default;comment:Group;"`                                                       // Group
	Path                *string         `json:"path,omitempty" gorm:"size:255;comment:Path or interface;"`                                                           // Path or interface
	Method              *string         `json:"method,omitempty" gorm:"size:255;comment:Method;"`                                                                    // Method
	Type                string          `json:"type" gorm:"size:20;not null;comment:Invocation type (failfast | failover | failsafe);"`                              // Invocation type (failfast | failover | failsafe)
	RetryPolicy         *string         `json:"retry_policy,omitempty" gorm:"type:json;comment:Retry policy (JSON);"`                                                 // Retry policy (JSON)
	Version             int64           `json:"version" gorm:"not null;default:1;comment:Version;"`                                                                  // Version
	Enabled             int             `json:"enabled" gorm:"not null;default:0;comment:Enabled;"`                                                                  // Enabled
	Description         *string         `json:"description,omitempty" gorm:"size:255;comment:Details;"`                                                              // Details
	Creator             *string         `json:"creator,omitempty" gorm:"size:255;comment:Creator;"`                                                                  // Creator
	Modifier            *string         `json:"modifier,omitempty" gorm:"size:255;comment:Modifier;"`                                                                // Modifier
	CreatedAt           time.Time       `json:"created_at" gorm:"autoCreateTime;comment:Create timestamp;"`                                                           // Create timestamp
	UpdatedAt           time.Time       `json:"updated_at,omitempty" gorm:"autoUpdateTime;comment:Update timestamp;"`                                                 // Update timestamp
	Deleted             string          `json:"-" gorm:"uniqueIndex:uniq_policy_invocation_name;size:20;default:0;comment:Delete flag;"`                             // Delete flag
	DeletedAt           *gorm.DeletedAt `json:"-" gorm:"comment:Delete timestamp;"`                                                                                  // Delete timestamp
}

func (a PolicyInvocation) TableName() string {
	return config.C.FormatTableName("policy_invocation")
}

// ConvertTo Convert `PolicyInvocation` to `PolicyInvocationForm` object.
func (a PolicyInvocation) ConvertTo(form *PolicyInvocationForm) error {
	form.ID = a.ID
	form.Name = a.Name
	form.SpaceCode = a.SpaceCode
	form.SourceApplicationID = a.SourceApplicationID
	form.TargetServiceId = a.TargetServiceId
	if !util.IsEmptyOrBlank(a.Group) {
		form.Group = a.Group
	} else {
		form.Group = DefaultGroup
	}
	form.Path = a.Path
	form.Method = a.Method
	form.Type = a.Type
	if !util.IsNilOrEmpty(a.RetryPolicy) {
		rp := new(RetryPolicy)
		json.UnMarshalToObject(*a.RetryPolicy, rp)
		form.RetryPolicy = rp
	}
	form.Version = a.Version
	form.Enabled = a.Enabled
	form.Description = a.Description
	form.Creator = a.Creator
	form.Modifier = a.Modifier
	form.CreatedAt = a.CreatedAt
	form.UpdatedAt = a.UpdatedAt
	return nil
}

// Defining the query parameters for the `PolicyInvocation` struct.
type PolicyInvocationQueryParam struct {
	util.PaginationParam
	SpaceCode       string `form:"space_code"`        // Microservice space code
	Name            string `form:"name"`              // Policy name (like)
	TargetServiceId string `form:"target_service_id"` // Target service ID
}

// Defining the query options for the `PolicyInvocation` struct.
type PolicyInvocationQueryOptions struct {
	util.QueryOptions
}

// Defining the query result for the `PolicyInvocation` struct.
type PolicyInvocationQueryResult struct {
	Data       PolicyInvocations
	PageResult *util.PaginationResult
}

// Defining the slice of `PolicyInvocation` struct.
type PolicyInvocations []*PolicyInvocation

// Defining the data structure for creating a `PolicyInvocation` struct.
type PolicyInvocationForm struct {
	ID                  string         `json:"id"`
	Name                string         `json:"name" binding:"required,max=100"`           // Policy name
	SpaceCode           string         `json:"space_code" binding:"required,max=255"`      // Microservice space code
	SourceApplicationID *string        `json:"source_application_id"`                       // Source application ID
	TargetServiceId     string         `json:"target_service_id" binding:"required,max=20"` // Target service ID
	Group               string         `json:"group" binding:"required,max=255"`          // Group
	Path                *string        `json:"path"`                                      // Path or interface
	Method              *string        `json:"method"`                                    // Method
	Type                string         `json:"type" binding:"required,max=20"`            // Invocation type (failfast | failover | failsafe)
	RetryPolicy         *RetryPolicy   `json:"retry_policy"`                               // Retry policy
	Version             int64          `json:"version"`                                   // Version
	Enabled             int            `json:"enabled"`                                   // Enabled
	Description         *string        `json:"description"`                               // Details
	Creator             *string        `json:"creator,omitempty"`                           // Creator
	Modifier            *string        `json:"modifier,omitempty"`                          // Modifier
	CreatedAt           time.Time      `json:"created_at"`                                  // Create timestamp
	UpdatedAt           time.Time      `json:"updated_at,omitempty"`                        // Update timestamp
}

// A validation function for the `PolicyInvocationForm` struct.
func (a *PolicyInvocationForm) Validate() error {
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
	if a.Type == "" {
		return errors.BadRequest("", "Type is required")
	}
	return nil
}

// Convert `PolicyInvocationForm` to `PolicyInvocation` object.
func (a *PolicyInvocationForm) FillTo(policyInvocation *PolicyInvocation) error {
	policyInvocation.Name = a.Name
	policyInvocation.SpaceCode = a.SpaceCode
	policyInvocation.SourceApplicationID = a.SourceApplicationID
	policyInvocation.TargetServiceId = a.TargetServiceId
	if a.Group != "" {
		policyInvocation.Group = a.Group
	} else {
		policyInvocation.Group = DefaultGroup
	}
	policyInvocation.Path = a.Path
	policyInvocation.Method = a.Method
	policyInvocation.Type = a.Type
	if a.RetryPolicy != nil {
		a.RetryPolicy.Version = time.Now().UnixMilli()
	}
	policyInvocation.RetryPolicy = func() *string { return json.MarshalToString(a.RetryPolicy) }()
	policyInvocation.Enabled = a.Enabled
	policyInvocation.Description = a.Description
	policyInvocation.Version = time.Now().UnixMilli()
	return nil
}
