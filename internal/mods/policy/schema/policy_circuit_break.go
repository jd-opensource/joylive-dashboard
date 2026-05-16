package schema

import (
	"time"

	"github.com/jd-opensource/joylive-dashboard/internal/config"
	"github.com/jd-opensource/joylive-dashboard/pkg/encoding/json"
	"github.com/jd-opensource/joylive-dashboard/pkg/errors"
	"github.com/jd-opensource/joylive-dashboard/pkg/util"
	"gorm.io/gorm"
)

// Circuit break policy management
type PolicyCircuitBreak struct {
	ID                          string          `json:"id" gorm:"size:20;primaryKey;<-:create;comment:Unique ID;"`                                                              // Unique ID
	Name                        string          `json:"name" gorm:"size:100;not null;uniqueIndex:uniq_policy_circuit_break_name;comment:Policy name;"`                          // Policy name
	SpaceCode                   string          `json:"space_code" gorm:"size:255;not null;uniqueIndex:uniq_policy_circuit_break_name;comment:Microservice space code;"`         // Microservice space code
	SourceApplicationID         *string         `json:"source_application_id,omitempty" gorm:"size:20;uniqueIndex:uniq_policy_circuit_break_name;comment:Source application ID;"` // Source application ID
	TargetServiceId             string          `json:"target_service_id" gorm:"size:20;not null;uniqueIndex:uniq_policy_circuit_break_name;comment:Target service ID;"`          // Target service ID
	Group                       string          `json:"group" gorm:"size:255;not null;default:default;comment:Group;"`                                                          // Group
	Path                        *string         `json:"path,omitempty" gorm:"size:255;comment:Path or interface;"`                                                              // Path or interface
	Method                      *string         `json:"method,omitempty" gorm:"size:255;comment:Method;"`                                                                       // Method
	Level                       string          `json:"level" gorm:"size:20;not null;comment:Policy level;"`                                                                    // Policy level
	SlidingWindowType           string          `json:"sliding_window_type" gorm:"size:20;not null;comment:Sliding window type;"`                                                 // Sliding window type
	SlidingWindowSize           int             `json:"sliding_window_size" gorm:"not null;default:1;comment:Sliding window size;"`                                               // Sliding window size
	MinCallsThreshold           int             `json:"min_calls_threshold" gorm:"not null;default:1;comment:Min calls threshold;"`                                               // Min calls threshold
	CodePolicy                  *string         `json:"code_policy,omitempty" gorm:"type:json;comment:Code policy (JSON);"`                                                      // Code policy (JSON)
	ErrorCodes                  *string         `json:"error_codes,omitempty" gorm:"type:json;comment:Error codes (JSON);"`                                                      // Error codes (JSON)
	MessagePolicy               *string         `json:"message_policy,omitempty" gorm:"type:json;comment:Message policy (JSON);"`                                                // Message policy (JSON)
	ErrorMessages               *string         `json:"error_messages,omitempty" gorm:"type:json;comment:Error messages (JSON);"`                                                // Error messages (JSON)
	Exceptions                  *string         `json:"exceptions,omitempty" gorm:"type:json;comment:Exceptions (JSON);"`                                                       // Exceptions (JSON)
	FailureRateThreshold        int             `json:"failure_rate_threshold" gorm:"not null;default:1;comment:Failure rate threshold;"`                                         // Failure rate threshold
	SlowCallRateThreshold       int             `json:"slow_call_rate_threshold" gorm:"not null;default:1;comment:Slow call rate threshold;"`                                      // Slow call rate threshold
	SlowCallDurationThreshold   int             `json:"slow_call_duration_threshold" gorm:"not null;default:10000;comment:Slow call duration threshold;"`                          // Slow call duration threshold
	WaitDurationInOpenState     int             `json:"wait_duration_in_open_state" gorm:"not null;default:10000;comment:Wait duration in open state;"`                             // Wait duration in open state
	OutlierMaxPercent           *int            `json:"outlier_max_percent,omitempty" gorm:"comment:Outlier max percent;"`                                                        // Outlier max percent
	AllowedCallsInHalfOpenState int             `json:"allowed_calls_in_half_open_state" gorm:"not null;default:5;comment:Allowed calls in half open state;"`                        // Allowed calls in half open state
	ForceOpen                   int             `json:"force_open" gorm:"not null;default:0;comment:Force open;"`                                                                // Force open
	RecoveryEnabled             int             `json:"recovery_enabled" gorm:"not null;default:0;comment:Recovery enabled;"`                                                    // Recovery enabled
	RecoveryDuration            int             `json:"recovery_duration" gorm:"not null;default:0;comment:Recovery duration (ms);"`                                             // Recovery duration (ms)
	RecoveryPhase               int             `json:"recovery_phase" gorm:"not null;default:0;comment:Recovery phase;"`                                                        // Recovery phase
	DegradeConfig               *string         `json:"degrade_config,omitempty" gorm:"type:json;comment:Degrade config (JSON);"`                                                // Degrade config (JSON)
	RealizeType                 string          `json:"realize_type" gorm:"size:20;not null;default:Resilience4j;comment:Realize type;"`                                         // Realize type
	Version                     int64           `json:"version" gorm:"not null;default:1;comment:Version;"`                                                                     // Version
	Enabled                     int             `json:"enabled" gorm:"not null;default:0;comment:Enabled;"`                                                                     // Enabled
	Description                 *string         `json:"description,omitempty" gorm:"size:255;comment:Details;"`                                                                 // Details
	Creator                     *string         `json:"creator,omitempty" gorm:"size:255;comment:Creator;"`                                                                     // Creator
	Modifier                    *string         `json:"modifier,omitempty" gorm:"size:255;comment:Modifier;"`                                                                   // Modifier
	CreatedAt                   time.Time       `json:"created_at" gorm:"autoCreateTime;comment:Create timestamp;"`                                                              // Create timestamp
	UpdatedAt                   time.Time       `json:"updated_at,omitempty" gorm:"autoUpdateTime;comment:Update timestamp;"`                                                    // Update timestamp
	Deleted                     string          `json:"-" gorm:"uniqueIndex:uniq_policy_circuit_break_name;size:20;default:0;comment:Delete flag;"`                             // Delete flag
	DeletedAt                   *gorm.DeletedAt `json:"-" gorm:"comment:Delete timestamp;"`                                                                                     // Delete timestamp
}

func (a PolicyCircuitBreak) TableName() string {
	return config.C.FormatTableName("policy_circuit_break")
}

// ConvertTo Convert `PolicyCircuitBreak` to `PolicyCircuitBreakForm` object.
func (a PolicyCircuitBreak) ConvertTo(form *PolicyCircuitBreakForm) error {
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
	form.Level = a.Level
	form.SlidingWindowType = a.SlidingWindowType
	form.SlidingWindowSize = a.SlidingWindowSize
	form.MinCallsThreshold = a.MinCallsThreshold
	if !util.IsNilOrEmpty(a.CodePolicy) {
		cp := new(ErrorParserPolicy)
		json.UnMarshalToObject(*a.CodePolicy, cp)
		form.CodePolicy = cp
	}
	if !util.IsNilOrEmpty(a.ErrorCodes) {
		ec := make([]string, 0)
		json.UnMarshalToObject(*a.ErrorCodes, &ec)
		form.ErrorCodes = ec
	}
	if !util.IsNilOrEmpty(a.MessagePolicy) {
		mp := new(ErrorParserPolicy)
		json.UnMarshalToObject(*a.MessagePolicy, mp)
		form.MessagePolicy = mp
	}
	if !util.IsNilOrEmpty(a.ErrorMessages) {
		em := make([]string, 0)
		json.UnMarshalToObject(*a.ErrorMessages, &em)
		form.ErrorMessages = em
	}
	if !util.IsNilOrEmpty(a.Exceptions) {
		ex := make([]string, 0)
		json.UnMarshalToObject(*a.Exceptions, &ex)
		form.Exceptions = ex
	}
	form.FailureRateThreshold = a.FailureRateThreshold
	form.SlowCallRateThreshold = a.SlowCallRateThreshold
	form.SlowCallDurationThreshold = a.SlowCallDurationThreshold
	form.WaitDurationInOpenState = a.WaitDurationInOpenState
	form.OutlierMaxPercent = a.OutlierMaxPercent
	form.AllowedCallsInHalfOpenState = a.AllowedCallsInHalfOpenState
	form.ForceOpen = a.ForceOpen
	form.RecoveryEnabled = a.RecoveryEnabled
	form.RecoveryDuration = a.RecoveryDuration
	form.RecoveryPhase = a.RecoveryPhase
	if !util.IsNilOrEmpty(a.DegradeConfig) {
		dc := new(DegradeConfig)
		json.UnMarshalToObject(*a.DegradeConfig, dc)
		form.DegradeConfig = dc
	}
	form.RealizeType = a.RealizeType
	form.Version = a.Version
	form.Enabled = a.Enabled
	form.Description = a.Description
	form.Creator = a.Creator
	form.Modifier = a.Modifier
	form.CreatedAt = a.CreatedAt
	form.UpdatedAt = a.UpdatedAt
	return nil
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
	ID                          string              `json:"id"`
	Name                        string              `json:"name" binding:"required,max=100"`             // Policy name
	SpaceCode                   string              `json:"space_code" binding:"required,max=255"`        // Microservice space code
	SourceApplicationID         *string             `json:"source_application_id"`                         // Source application ID
	TargetServiceId             string              `json:"target_service_id" binding:"required,max=20"`   // Target service ID
	Group                       string              `json:"group" binding:"required,max=255"`            // Group
	Path                        *string             `json:"path"`                                        // Path or interface
	Method                      *string             `json:"method"`                                      // Method
	Level                       string              `json:"level" binding:"required,max=20"`             // Policy level
	SlidingWindowType           string              `json:"sliding_window_type" binding:"required,max=20"` // Sliding window type
	SlidingWindowSize           int                 `json:"sliding_window_size"`                           // Sliding window size
	MinCallsThreshold           int                 `json:"min_calls_threshold"`                           // Min calls threshold
	CodePolicy                  *ErrorParserPolicy  `json:"code_policy"`                                  // Code policy
	ErrorCodes                  []string            `json:"error_codes"`                                  // Error codes
	MessagePolicy               *ErrorParserPolicy  `json:"message_policy"`                               // Message policy
	ErrorMessages               []string            `json:"error_messages"`                               // Error messages
	Exceptions                  []string            `json:"exceptions"`                                  // Exceptions
	FailureRateThreshold        int                 `json:"failure_rate_threshold"`                        // Failure rate threshold
	SlowCallRateThreshold       int                 `json:"slow_call_rate_threshold"`                       // Slow call rate threshold
	SlowCallDurationThreshold   int                 `json:"slow_call_duration_threshold"`                   // Slow call duration threshold
	WaitDurationInOpenState     int                 `json:"wait_duration_in_open_state"`                     // Wait duration in open state
	OutlierMaxPercent           *int                `json:"outlier_max_percent"`                           // Outlier max percent
	AllowedCallsInHalfOpenState int                 `json:"allowed_calls_in_half_open_state"`                 // Allowed calls in half open state
	ForceOpen                   int                 `json:"force_open"`                                   // Force open
	RecoveryEnabled             int                 `json:"recovery_enabled"`                             // Recovery enabled
	RecoveryDuration            int                 `json:"recovery_duration"`                            // Recovery duration (ms)
	RecoveryPhase               int                 `json:"recovery_phase"`                               // Recovery phase
	DegradeConfig               *DegradeConfig      `json:"degrade_config"`                               // Degrade config
	RealizeType                 string              `json:"realize_type" binding:"required,max=20"`       // Realize type
	Version                     int64               `json:"version"`                                     // Version
	Enabled                     int                 `json:"enabled"`                                     // Enabled
	Description                 *string             `json:"description"`                                 // Details
	Creator                     *string             `json:"creator,omitempty"`                           // Creator
	Modifier                    *string             `json:"modifier,omitempty"`                          // Modifier
	CreatedAt                   time.Time           `json:"created_at"`                                  // Create timestamp
	UpdatedAt                   time.Time           `json:"updated_at,omitempty"`                        // Update timestamp
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
	if a.Group != "" {
		policyCircuitBreak.Group = a.Group
	} else {
		policyCircuitBreak.Group = DefaultGroup
	}
	policyCircuitBreak.Path = a.Path
	policyCircuitBreak.Method = a.Method
	policyCircuitBreak.Level = a.Level
	policyCircuitBreak.SlidingWindowType = a.SlidingWindowType
	policyCircuitBreak.SlidingWindowSize = a.SlidingWindowSize
	policyCircuitBreak.MinCallsThreshold = a.MinCallsThreshold
	policyCircuitBreak.CodePolicy = func() *string { return json.MarshalToString(a.CodePolicy) }()
	policyCircuitBreak.ErrorCodes = func() *string { return json.MarshalToString(a.ErrorCodes) }()
	policyCircuitBreak.MessagePolicy = func() *string { return json.MarshalToString(a.MessagePolicy) }()
	policyCircuitBreak.ErrorMessages = func() *string { return json.MarshalToString(a.ErrorMessages) }()
	policyCircuitBreak.Exceptions = func() *string { return json.MarshalToString(a.Exceptions) }()
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
	policyCircuitBreak.DegradeConfig = func() *string { return json.MarshalToString(a.DegradeConfig) }()
	policyCircuitBreak.RealizeType = a.RealizeType
	policyCircuitBreak.Enabled = a.Enabled
	policyCircuitBreak.Description = a.Description
	policyCircuitBreak.Version = time.Now().UnixMilli()
	return nil
}
