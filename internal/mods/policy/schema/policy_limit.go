package schema

import (
	"time"

	"github.com/jd-opensource/joylive-dashboard/internal/config"
	"github.com/jd-opensource/joylive-dashboard/pkg/encoding/json"
	"github.com/jd-opensource/joylive-dashboard/pkg/errors"
	"github.com/jd-opensource/joylive-dashboard/pkg/util"
	"gorm.io/gorm"
)

// Limit policy management
type PolicyLimit struct {
	ID                  string          `json:"id" gorm:"size:20;primaryKey;<-:create;comment:Unique ID;"`                                               // Unique ID
	Name                string          `json:"name" gorm:"size:100;not null;uniqueIndex:uniq_policy_limit_name;comment:Policy name;"`                   // Policy name
	SpaceCode           string          `json:"space_code" gorm:"size:255;not null;comment:Microservice space code;"`                                    // Microservice space code
	SourceApplicationID *string         `json:"source_application_id,omitempty" gorm:"size:20;comment:Source application ID;"`                           // Source application ID
	TargetServiceId     string          `json:"target_service_id" gorm:"size:20;not null;uniqueIndex:uniq_policy_limit_name;comment:Target service ID;"` // Target service ID
	Group               string          `json:"group" gorm:"size:255;not null;default:default;comment:Group;"`                                           // Group
	Path                *string         `json:"path,omitempty" gorm:"size:255;comment:Path or interface;"`                                               // Path or interface
	Method              *string         `json:"method,omitempty" gorm:"size:255;comment:Method;"`                                                        // Method
	RelationType        string          `json:"relation_type" gorm:"size:20;not null;comment:Relation type;"`                                            // Relation type
	Conditions          *string         `json:"conditions,omitempty" gorm:"type:json;comment:Match conditions (JSON);"`                                  // Match conditions (JSON)
	Type                string          `json:"type" gorm:"size:20;not null;comment:Limit policy type (Rate, Concurrency, Load);"`                       // Limit policy type (Rate, Concurrency, Load)
	RealizeType         string          `json:"realize_type" gorm:"size:20;not null;comment:Realize type;"`                                              // Realize type
	SlidingWindows      *string         `json:"sliding_windows,omitempty" gorm:"type:json;comment:Sliding windows (JSON);"`                              // Sliding windows (JSON)
	MaxConcurrency      *int            `json:"max_concurrency,omitempty" gorm:"default:0;comment:Max concurrency;"`                                     // Max concurrency
	MaxWaitMs           int64           `json:"max_wait_ms" gorm:"not null;default:0;comment:Max wait time (ms);"`                                       // Max wait time (ms)
	CpuUsage            *int            `json:"cpu_usage,omitempty" gorm:"default:0;comment:Max CPU usage;"`                                             // Max CPU usage
	LoadUsage           *int            `json:"load_usage,omitempty" gorm:"default:0;comment:Max system load;"`                                          // Max system load
	ActionParameters    *string         `json:"action_parameters,omitempty" gorm:"type:json;comment:Action parameters (JSON);"`                          // Action parameters (JSON)
	Version             int64           `json:"version" gorm:"not null;default:1;comment:Version;"`                                                      // Version
	Enabled             int             `json:"enabled" gorm:"not null;default:0;comment:Enabled;"`                                                      // Enabled
	Description         *string         `json:"description,omitempty" gorm:"size:255;comment:Details;"`                                                  // Details
	Creator             *string         `json:"creator,omitempty" gorm:"size:255;comment:Creator;"`                                                      // Creator
	Modifier            *string         `json:"modifier,omitempty" gorm:"size:255;comment:Modifier;"`                                                    // Modifier
	CreatedAt           time.Time       `json:"created_at" gorm:"autoCreateTime;comment:Create timestamp;"`                                              // Create timestamp
	UpdatedAt           time.Time       `json:"updated_at,omitempty" gorm:"autoUpdateTime;comment:Update timestamp;"`                                    // Update timestamp
	Deleted             string          `json:"-" gorm:"uniqueIndex:uniq_policy_limit_name;size:20;default:0;comment:Delete flag;"`                      // Delete flag
	DeletedAt           *gorm.DeletedAt `json:"-" gorm:"comment:Delete timestamp;"`                                                                      // Delete timestamp
}

func (a PolicyLimit) TableName() string {
	return config.C.FormatTableName("policy_limit")
}

// ConvertTo Convert `PolicyLimit` to `PolicyLimitForm` object.
func (a PolicyLimit) ConvertTo(limit *PolicyLimitForm) error {
	limit.ID = a.ID
	limit.Name = a.Name
	limit.SpaceCode = a.SpaceCode
	limit.SourceApplicationID = a.SourceApplicationID
	limit.TargetServiceId = a.TargetServiceId
	// limit.TargetService = a.TargetService
	if !util.IsEmptyOrBlank(a.Group) {
		limit.Group = a.Group
	} else {
		limit.Group = DefaultGroup
	}
	limit.Path = a.Path
	limit.Method = a.Method
	limit.RelationType = a.RelationType
	if !util.IsNilOrEmpty(a.Conditions) {
		conditions := make([]TagCondition, 0)
		json.UnMarshalToObject(*a.Conditions, &conditions)
		limit.Conditions = &conditions
	}
	limit.Type = a.Type
	limit.RealizeType = a.RealizeType
	if !util.IsNilOrEmpty(a.SlidingWindows) {
		sw := make([]SlidingWindow, 0)
		json.UnMarshalToObject(*a.SlidingWindows, &sw)
		limit.SlidingWindows = &sw
	}
	limit.MaxConcurrency = a.MaxConcurrency
	limit.MaxWaitMs = a.MaxWaitMs
	limit.CpuUsage = a.CpuUsage
	limit.LoadUsage = a.LoadUsage
	if !util.IsNilOrEmpty(a.ActionParameters) {
		limit.ActionParameters, _ = json.UnMarshalToMap(*a.ActionParameters)
	}
	limit.Version = a.Version
	limit.Enabled = a.Enabled
	limit.Description = a.Description
	limit.Creator = a.Creator
	limit.Modifier = a.Modifier
	limit.CreatedAt = a.CreatedAt
	limit.UpdatedAt = a.UpdatedAt
	return nil
}

// Defining the query parameters for the `PolicyLimit` struct.
type PolicyLimitQueryParam struct {
	util.PaginationParam
	SpaceCode       string `form:"space_code"`        // Microservice space code
	Name            string `form:"name"`              // Policy name (like)
	TargetServiceId string `form:"target_service_id"` // Target service ID
}

// Defining the query options for the `PolicyLimit` struct.
type PolicyLimitQueryOptions struct {
	util.QueryOptions
}

// Defining the query result for the `PolicyLimit` struct.
type PolicyLimitQueryResult struct {
	Data       PolicyLimits
	PageResult *util.PaginationResult
}

// Defining the slice of `PolicyLimit` struct.
type PolicyLimits []*PolicyLimit

// Defining the data structure for creating a `PolicyLimit` struct.
type PolicyLimitForm struct {
	ID                  string            `json:"id"`
	Name                string            `json:"name" binding:"required,max=100"`             // Policy name
	SpaceCode           string            `json:"space_code" binding:"required,max=255"`       // Microservice space code
	SourceApplicationID *string           `json:"source_application_id"`                       // Source application ID
	TargetServiceId     string            `json:"target_service_id" binding:"required,max=20"` // Target service ID
	Group               string            `json:"group" binding:"required,max=255"`            // Group
	Path                *string           `json:"path"`                                        // Path or interface
	Method              *string           `json:"method"`                                      // Method
	RelationType        string            `json:"relation_type" binding:"required,max=20"`     // Relation type
	Conditions          *[]TagCondition   `json:"conditions"`                                  // Match conditions (JSON)
	Type                string            `json:"type" binding:"required,max=20"`              // Limit policy type (Rate, Concurrency, Load)
	RealizeType         string            `json:"realize_type" binding:"required,max=20"`      // Realize type
	SlidingWindows      *[]SlidingWindow  `json:"sliding_windows"`                             // Sliding windows (JSON)
	MaxConcurrency      *int              `json:"max_concurrency"`                             // Max concurrency
	MaxWaitMs           int64             `json:"max_wait_ms"`                                 // Max wait time (ms)
	CpuUsage            *int              `json:"cpu_usage"`                                   // Max CPU usage
	LoadUsage           *int              `json:"load_usage"`                                  // Max system load
	ActionParameters    map[string]string `json:"action_parameters"`                           // Action parameters (JSON)
	Version             int64             `json:"version"`                                     // Version
	Enabled             int               `json:"enabled"`                                     // Enabled
	Description         *string           `json:"description"`                                 // Details
	Creator             *string           `json:"creator,omitempty"`                           // Creator
	Modifier            *string           `json:"modifier,omitempty"`                          // Modifier
	CreatedAt           time.Time         `json:"created_at"`                                  // Create timestamp
	UpdatedAt           time.Time         `json:"updated_at,omitempty"`                        // Update timestamp
}

// A validation function for the `PolicyLimitForm` struct.
func (a *PolicyLimitForm) Validate() error {
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
	if a.RealizeType == "" {
		return errors.BadRequest("", "RealizeType is required")
	}
	return nil
}

// Convert `PolicyLimitForm` to `PolicyLimit` object.
func (a *PolicyLimitForm) FillTo(policyLimit *PolicyLimit) error {
	policyLimit.Name = a.Name
	policyLimit.SpaceCode = a.SpaceCode
	policyLimit.SourceApplicationID = a.SourceApplicationID
	policyLimit.TargetServiceId = a.TargetServiceId
	//policyLimit.SourceService = a.SourceService
	// policyLimit.TargetService = a.TargetService
	if a.Group != "" {
		policyLimit.Group = a.Group
	} else {
		policyLimit.Group = DefaultGroup
	}
	policyLimit.Path = a.Path
	if !util.IsNilOrEmpty(policyLimit.Path) {
		policyLimit.Method = a.Method
	}
	policyLimit.RelationType = a.RelationType
	policyLimit.Conditions = func() *string { return json.MarshalToString(a.Conditions) }()
	if len(a.Type) != 0 {
		policyLimit.Type = a.Type
	} else {
		policyLimit.Type = DefaultLimitType
	}
	if len(a.RealizeType) != 0 {
		policyLimit.RealizeType = a.RealizeType
	} else {
		policyLimit.RealizeType = DefaultAgentLimitRealizeType
	}
	policyLimit.SlidingWindows = func() *string { return json.MarshalToString(a.SlidingWindows) }()
	policyLimit.MaxConcurrency = a.MaxConcurrency
	policyLimit.MaxWaitMs = a.MaxWaitMs
	policyLimit.CpuUsage = a.CpuUsage
	policyLimit.LoadUsage = a.LoadUsage
	policyLimit.ActionParameters = func() *string { return json.MarshalToString(a.ActionParameters) }()
	policyLimit.Enabled = a.Enabled
	policyLimit.Description = a.Description
	policyLimit.Version = time.Now().UnixMilli()
	return nil
}
