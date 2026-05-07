package schema

import (
	"time"

	"github.com/jd-opensource/joylive-dashboard/internal/config"
	"github.com/jd-opensource/joylive-dashboard/pkg/errors"
	"github.com/jd-opensource/joylive-dashboard/pkg/util"
	"gorm.io/gorm"
)

// Circuit break policy management
type PolicyCircuitBreak struct {
	ID                          string          `json:"id" gorm:"size:20;primaryKey;<-:create;comment:Unique ID;"`                                         // Unique ID
	Name                        string          `json:"name" gorm:"size:100;not null;uniqueIndex:uniq_name;comment:Policy name;"`                          // Policy name
	SpaceCode                   string          `json:"spaceCode" gorm:"size:255;not null;uniqueIndex:uniq_name;comment:Microservice space code;"`         // Microservice space code
	SourceApplicationID         *string         `json:"sourceApplicationId,omitempty" gorm:"size:20;uniqueIndex:uniq_name;comment:Source application ID;"` // Source application ID
	TargetServiceId             string          `json:"targetServiceId" gorm:"size:20;not null;uniqueIndex:uniq_name;comment:Target service ID;"`          // Target service ID
	Group                       string          `json:"group" gorm:"size:255;not null;default:default;comment:Group;"`                                     // Group
	Path                        *string         `json:"path,omitempty" gorm:"size:255;comment:Path or interface;"`                                         // Path or interface
	Method                      *string         `json:"method,omitempty" gorm:"size:255;comment:Method;"`                                                  // Method
	Level                       string          `json:"level" gorm:"size:20;not null;comment:Policy level;"`                                               // Policy level
	SlidingWindowType           string          `json:"slidingWindowType" gorm:"size:20;not null;comment:Sliding window type;"`                            // Sliding window type
	SlidingWindowSize           int             `json:"slidingWindowSize" gorm:"not null;default:1;comment:Sliding window size;"`                          // Sliding window size
	MinCallsThreshold           int             `json:"minCallsThreshold" gorm:"not null;default:1;comment:Min calls threshold;"`                          // Min calls threshold
	CodePolicy                  *string         `json:"codePolicy,omitempty" gorm:"type:json;comment:Code policy (JSON);"`                                 // Code policy (JSON)
	ErrorCodes                  *string         `json:"errorCodes,omitempty" gorm:"type:json;comment:Error codes (JSON);"`                                 // Error codes (JSON)
	MessagePolicy               *string         `json:"messagePolicy,omitempty" gorm:"type:json;comment:Message policy (JSON);"`                           // Message policy (JSON)
	ErrorMessages               *string         `json:"errorMessages,omitempty" gorm:"type:json;comment:Error messages (JSON);"`                           // Error messages (JSON)
	Exceptions                  *string         `json:"exceptions,omitempty" gorm:"type:json;comment:Exceptions (JSON);"`                                  // Exceptions (JSON)
	FailureRateThreshold        int             `json:"failureRateThreshold" gorm:"not null;default:1;comment:Failure rate threshold;"`                    // Failure rate threshold
	SlowCallRateThreshold       int             `json:"slowCallRateThreshold" gorm:"not null;default:1;comment:Slow call rate threshold;"`                 // Slow call rate threshold
	SlowCallDurationThreshold   int             `json:"slowCallDurationThreshold" gorm:"not null;default:10000;comment:Slow call duration threshold;"`     // Slow call duration threshold
	WaitDurationInOpenState     int             `json:"waitDurationInOpenState" gorm:"not null;default:10000;comment:Wait duration in open state;"`        // Wait duration in open state
	OutlierMaxPercent           *int            `json:"outlierMaxPercent,omitempty" gorm:"comment:Outlier max percent;"`                                   // Outlier max percent
	AllowedCallsInHalfOpenState int             `json:"allowedCallsInHalfOpenState" gorm:"not null;default:5;comment:Allowed calls in half open state;"`   // Allowed calls in half open state
	ForceOpen                   int             `json:"forceOpen" gorm:"not null;default:0;comment:Force open;"`                                           // Force open
	RecoveryEnabled             int             `json:"recoveryEnabled" gorm:"not null;default:0;comment:Recovery enabled;"`                               // Recovery enabled
	RecoveryDuration            int             `json:"recoveryDuration" gorm:"not null;default:0;comment:Recovery duration (ms);"`                        // Recovery duration (ms)
	RecoveryPhase               int             `json:"recoveryPhase" gorm:"not null;default:0;comment:Recovery phase;"`                                   // Recovery phase
	DegradeConfig               *string         `json:"degradeConfig,omitempty" gorm:"type:json;comment:Degrade config (JSON);"`                           // Degrade config (JSON)
	RealizeType                 string          `json:"realizeType" gorm:"size:20;not null;default:Resilience4j;comment:Realize type;"`                    // Realize type
	Version                     int64           `json:"version" gorm:"not null;default:1;comment:Version;"`                                                // Version
	Enabled                     int             `json:"enabled" gorm:"not null;default:0;comment:Enabled;"`                                                // Enabled
	Description                 *string         `json:"description,omitempty" gorm:"size:255;comment:Details;"`                                            // Details
	Creator                     *string         `json:"creator,omitempty" gorm:"size:255;comment:Creator;"`                                                // Creator
	Modifier                    *string         `json:"modifier,omitempty" gorm:"size:255;comment:Modifier;"`                                              // Modifier
	CreatedAt                   time.Time       `json:"createdAt" gorm:"autoCreateTime;comment:Create timestamp;"`                                         // Create timestamp
	UpdatedAt                   time.Time       `json:"updatedAt,omitempty" gorm:"autoUpdateTime;comment:Update timestamp;"`                               // Update timestamp
	Deleted                     string          `json:"-" gorm:"uniqueIndex:uniq_name;size:20;default:0;comment:Delete flag;"`                             // Delete flag
	DeletedAt                   *gorm.DeletedAt `json:"-" gorm:"comment:Delete timestamp;"`                                                                // Delete timestamp
}

func (a PolicyCircuitBreak) TableName() string {
	return config.C.FormatTableName("policy_circuit_break")
}

// Defining the query parameters for the `PolicyCircuitBreak` struct.
type PolicyCircuitBreakQueryParam struct {
	util.PaginationParam
	SpaceCode       string `form:"space_code"`        // Microservice space code
	Name            string `form:"name"`              // Policy name (like)
	TargetServiceId string `form:"target_service_id"` // Target service ID
}

// Defining the query options for the `PolicyCircuitBreak` struct.
type PolicyCircuitBreakQueryOptions struct {
	util.QueryOptions
}

// Defining the query result for the `PolicyCircuitBreak` struct.
type PolicyCircuitBreakQueryResult struct {
	Data       PolicyCircuitBreaks
	PageResult *util.PaginationResult
}

// Defining the slice of `PolicyCircuitBreak` struct.
type PolicyCircuitBreaks []*PolicyCircuitBreak

// Defining the data structure for creating a `PolicyCircuitBreak` struct.
type PolicyCircuitBreakForm struct {
	Name                        string  `json:"name" binding:"required,max=100"`               // Policy name
	SpaceCode                   string  `json:"spaceCode" binding:"required,max=255"`          // Microservice space code
	SourceApplicationID         *string `json:"sourceApplicationId"`                           // Source application ID
	TargetServiceId             string  `json:"targetServiceId" binding:"required,max=20"`     // Target service ID
	Group                       string  `json:"group" binding:"required,max=255"`              // Group
	Path                        *string `json:"path"`                                          // Path or interface
	Method                      *string `json:"method"`                                        // Method
	Level                       string  `json:"level" binding:"required,max=20"`               // Policy level
	SlidingWindowType           string  `json:"slidingWindowType" binding:"required,max=20"`   // Sliding window type
	SlidingWindowSize           int     `json:"slidingWindowSize"`                             // Sliding window size
	MinCallsThreshold           int     `json:"minCallsThreshold"`                             // Min calls threshold
	CodePolicy                  *string `json:"codePolicy"`                                    // Code policy (JSON)
	ErrorCodes                  *string `json:"errorCodes"`                                    // Error codes (JSON)
	MessagePolicy               *string `json:"messagePolicy"`                                 // Message policy (JSON)
	ErrorMessages               *string `json:"errorMessages"`                                 // Error messages (JSON)
	Exceptions                  *string `json:"exceptions"`                                    // Exceptions (JSON)
	FailureRateThreshold        int     `json:"failureRateThreshold"`                          // Failure rate threshold
	SlowCallRateThreshold       int     `json:"slowCallRateThreshold"`                         // Slow call rate threshold
	SlowCallDurationThreshold   int     `json:"slowCallDurationThreshold"`                     // Slow call duration threshold
	WaitDurationInOpenState     int     `json:"waitDurationInOpenState"`                       // Wait duration in open state
	OutlierMaxPercent           *int    `json:"outlierMaxPercent"`                             // Outlier max percent
	AllowedCallsInHalfOpenState int     `json:"allowedCallsInHalfOpenState"`                   // Allowed calls in half open state
	ForceOpen                   int     `json:"forceOpen"`                                     // Force open
	RecoveryEnabled             int     `json:"recoveryEnabled"`                               // Recovery enabled
	RecoveryDuration            int     `json:"recoveryDuration"`                              // Recovery duration (ms)
	RecoveryPhase               int     `json:"recoveryPhase"`                                 // Recovery phase
	DegradeConfig               *string `json:"degradeConfig"`                                 // Degrade config (JSON)
	RealizeType                 string  `json:"realizeType" binding:"required,max=20"`         // Realize type
	Version                     int64   `json:"version"`                                       // Version
	Enabled                     int     `json:"enabled"`                                       // Enabled
	Description                 *string `json:"description"`                                   // Details
}

// A validation function for the `PolicyCircuitBreakForm` struct.
func (a *PolicyCircuitBreakForm) Validate() error {
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
	if a.Level == "" {
		return errors.BadRequest("", "Level is required")
	}
	if a.SlidingWindowType == "" {
		return errors.BadRequest("", "SlidingWindowType is required")
	}
	if a.RealizeType == "" {
		return errors.BadRequest("", "RealizeType is required")
	}
	return nil
}

// Convert `PolicyCircuitBreakForm` to `PolicyCircuitBreak` object.
func (a *PolicyCircuitBreakForm) FillTo(policyCircuitBreak *PolicyCircuitBreak) error {
	policyCircuitBreak.Name = a.Name
	policyCircuitBreak.SpaceCode = a.SpaceCode
	policyCircuitBreak.SourceApplicationID = a.SourceApplicationID
	policyCircuitBreak.TargetServiceId = a.TargetServiceId
	policyCircuitBreak.Group = a.Group
	policyCircuitBreak.Path = a.Path
	policyCircuitBreak.Method = a.Method
	policyCircuitBreak.Level = a.Level
	policyCircuitBreak.SlidingWindowType = a.SlidingWindowType
	policyCircuitBreak.SlidingWindowSize = a.SlidingWindowSize
	policyCircuitBreak.MinCallsThreshold = a.MinCallsThreshold
	policyCircuitBreak.CodePolicy = a.CodePolicy
	policyCircuitBreak.ErrorCodes = a.ErrorCodes
	policyCircuitBreak.MessagePolicy = a.MessagePolicy
	policyCircuitBreak.ErrorMessages = a.ErrorMessages
	policyCircuitBreak.Exceptions = a.Exceptions
	policyCircuitBreak.FailureRateThreshold = a.FailureRateThreshold
	policyCircuitBreak.SlowCallRateThreshold = a.SlowCallRateThreshold
	policyCircuitBreak.SlowCallDurationThreshold = a.SlowCallDurationThreshold
	policyCircuitBreak.WaitDurationInOpenState = a.WaitDurationInOpenState
	policyCircuitBreak.OutlierMaxPercent = a.OutlierMaxPercent
	policyCircuitBreak.AllowedCallsInHalfOpenState = a.AllowedCallsInHalfOpenState
	policyCircuitBreak.ForceOpen = a.ForceOpen
	policyCircuitBreak.RecoveryEnabled = a.RecoveryEnabled
	policyCircuitBreak.RecoveryDuration = a.RecoveryDuration
	policyCircuitBreak.RecoveryPhase = a.RecoveryPhase
	policyCircuitBreak.DegradeConfig = a.DegradeConfig
	policyCircuitBreak.RealizeType = a.RealizeType
	policyCircuitBreak.Version = a.Version
	policyCircuitBreak.Enabled = a.Enabled
	policyCircuitBreak.Description = a.Description
	return nil
}
