package schema

import (
	"time"

	"github.com/jd-opensource/joylive-dashboard/internal/config"
	"github.com/jd-opensource/joylive-dashboard/pkg/errors"
	"github.com/jd-opensource/joylive-dashboard/pkg/util"
	"gorm.io/gorm"
)

// Fault injection policy management
type PolicyFault struct {
	ID                  string          `json:"id" gorm:"size:20;primaryKey;<-:create;comment:Unique ID;"`                                                      // Unique ID
	Name                string          `json:"name" gorm:"size:100;not null;uniqueIndex:uniq_policy_fault_name;comment:Policy name;"`                          // Policy name
	SpaceCode           string          `json:"spaceCode" gorm:"size:255;not null;uniqueIndex:uniq_policy_fault_name;comment:Microservice space code;"`         // Microservice space code
	SourceApplicationID *string         `json:"sourceApplicationId,omitempty" gorm:"size:20;uniqueIndex:uniq_policy_fault_name;comment:Source application ID;"` // Source application ID
	TargetServiceId     string          `json:"targetServiceId" gorm:"size:20;not null;uniqueIndex:uniq_policy_fault_name;comment:Target service ID;"`          // Target service ID
	Group               string          `json:"group" gorm:"size:255;not null;default:default;comment:Group;"`                                                  // Group
	Path                *string         `json:"path,omitempty" gorm:"size:255;comment:Path or interface;"`                                                      // Path or interface
	Method              *string         `json:"method,omitempty" gorm:"size:255;comment:Method;"`                                                               // Method
	RelationType        string          `json:"relationType" gorm:"size:20;not null;comment:Relation type;"`                                                    // Relation type
	Conditions          *string         `json:"conditions,omitempty" gorm:"type:json;comment:Match conditions (JSON);"`                                         // Match conditions (JSON)
	Type                string          `json:"type" gorm:"size:20;not null;comment:Fault injection type (error | delay);"`                                     // Fault injection type (error | delay)
	Scope               string          `json:"scope" gorm:"size:20;not null;comment:Injection scope (inbound | outbound);"`                                    // Injection scope (inbound | outbound)
	ErrorCode           *int            `json:"errorCode,omitempty" gorm:"default:0;comment:Error code;"`                                                       // Error code
	ErrorMsg            *string         `json:"errorMsg,omitempty" gorm:"size:1024;comment:Error message;"`                                                     // Error message
	DelayTimeMs         *int            `json:"delayTimeMs,omitempty" gorm:"default:0;comment:Delay time (ms);"`                                                // Delay time (ms)
	Percent             int             `json:"percent" gorm:"not null;default:0;comment:Fault injection percent;"`                                             // Fault injection percent
	Version             int64           `json:"version" gorm:"not null;default:1;comment:Version;"`                                                             // Version
	Enabled             int             `json:"enabled" gorm:"not null;default:0;comment:Enabled;"`                                                             // Enabled
	Description         *string         `json:"description,omitempty" gorm:"size:255;comment:Details;"`                                                         // Details
	Creator             *string         `json:"creator,omitempty" gorm:"size:255;comment:Creator;"`                                                             // Creator
	Modifier            *string         `json:"modifier,omitempty" gorm:"size:255;comment:Modifier;"`                                                           // Modifier
	CreatedAt           time.Time       `json:"createdAt" gorm:"autoCreateTime;comment:Create timestamp;"`                                                      // Create timestamp
	UpdatedAt           time.Time       `json:"updatedAt,omitempty" gorm:"autoUpdateTime;comment:Update timestamp;"`                                            // Update timestamp
	Deleted             string          `json:"-" gorm:"uniqueIndex:uniq_policy_fault_name;size:20;default:0;comment:Delete flag;"`                             // Delete flag
	DeletedAt           *gorm.DeletedAt `json:"-" gorm:"comment:Delete timestamp;"`                                                                             // Delete timestamp
}

func (a PolicyFault) TableName() string {
	return config.C.FormatTableName("policy_fault")
}

// Defining the query parameters for the `PolicyFault` struct.
type PolicyFaultQueryParam struct {
	util.PaginationParam
	SpaceCode       string `form:"space_code"`        // Microservice space code
	Name            string `form:"name"`              // Policy name (like)
	TargetServiceId string `form:"target_service_id"` // Target service ID
}

// Defining the query options for the `PolicyFault` struct.
type PolicyFaultQueryOptions struct {
	util.QueryOptions
}

// Defining the query result for the `PolicyFault` struct.
type PolicyFaultQueryResult struct {
	Data       PolicyFaults
	PageResult *util.PaginationResult
}

// Defining the slice of `PolicyFault` struct.
type PolicyFaults []*PolicyFault

// Defining the data structure for creating a `PolicyFault` struct.
type PolicyFaultForm struct {
	Name                string  `json:"name" binding:"required,max=100"`           // Policy name
	SpaceCode           string  `json:"spaceCode" binding:"required,max=255"`      // Microservice space code
	SourceApplicationID *string `json:"sourceApplicationId"`                       // Source application ID
	TargetServiceId     string  `json:"targetServiceId" binding:"required,max=20"` // Target service ID
	Group               string  `json:"group" binding:"required,max=255"`          // Group
	Path                *string `json:"path"`                                      // Path or interface
	Method              *string `json:"method"`                                    // Method
	RelationType        string  `json:"relationType" binding:"required,max=20"`    // Relation type
	Conditions          *string `json:"conditions"`                                // Match conditions (JSON)
	Type                string  `json:"type" binding:"required,max=20"`            // Fault injection type (error | delay)
	Scope               string  `json:"scope" binding:"required,max=20"`           // Injection scope (inbound | outbound)
	ErrorCode           *int    `json:"errorCode"`                                 // Error code
	ErrorMsg            *string `json:"errorMsg"`                                  // Error message
	DelayTimeMs         *int    `json:"delayTimeMs"`                               // Delay time (ms)
	Percent             int     `json:"percent"`                                   // Fault injection percent
	Version             int64   `json:"version"`                                   // Version
	Enabled             int     `json:"enabled"`                                   // Enabled
	Description         *string `json:"description"`                               // Details
}

// A validation function for the `PolicyFaultForm` struct.
func (a *PolicyFaultForm) Validate() error {
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
	if a.Scope == "" {
		return errors.BadRequest("", "Scope is required")
	}
	return nil
}

// Convert `PolicyFaultForm` to `PolicyFault` object.
func (a *PolicyFaultForm) FillTo(policyFault *PolicyFault) error {
	policyFault.Name = a.Name
	policyFault.SpaceCode = a.SpaceCode
	policyFault.SourceApplicationID = a.SourceApplicationID
	policyFault.TargetServiceId = a.TargetServiceId
	policyFault.Group = a.Group
	policyFault.Path = a.Path
	policyFault.Method = a.Method
	policyFault.RelationType = a.RelationType
	policyFault.Conditions = a.Conditions
	policyFault.Type = a.Type
	policyFault.Scope = a.Scope
	policyFault.ErrorCode = a.ErrorCode
	policyFault.ErrorMsg = a.ErrorMsg
	policyFault.DelayTimeMs = a.DelayTimeMs
	policyFault.Percent = a.Percent
	policyFault.Version = a.Version
	policyFault.Enabled = a.Enabled
	policyFault.Description = a.Description
	return nil
}
