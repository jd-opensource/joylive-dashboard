package schema

import (
	"time"

	"github.com/jd-opensource/joylive-dashboard/internal/config"
	"github.com/jd-opensource/joylive-dashboard/pkg/errors"
	"github.com/jd-opensource/joylive-dashboard/pkg/util"
	"gorm.io/gorm"
)

// Auth policy management
type PolicyAuth struct {
	ID                  string          `json:"id" gorm:"size:20;primaryKey;<-:create;comment:Unique ID;"`                     // Unique ID
	SpaceCode           string          `json:"space_code" gorm:"size:255;not null;comment:Microservice space code;"`          // Microservice space code
	SourceApplicationID *string         `json:"source_application_id,omitempty" gorm:"size:20;comment:Source application ID;"` // Source application ID
	TargetServiceId     string          `json:"target_service_id" gorm:"size:20;not null;comment:Target service ID;"`          // Target service ID
	Type                string          `json:"type" gorm:"size:20;not null;comment:Auth type;"`                               // Auth type
	Params              *string         `json:"params,omitempty" gorm:"type:json;comment:Parameters (JSON);"`                  // Parameters (JSON)
	Version             int64           `json:"version" gorm:"not null;default:1;comment:Version;"`                            // Version
	Enabled             int             `json:"enabled" gorm:"not null;default:0;comment:Enabled;"`                            // Enabled
	Description         *string         `json:"description,omitempty" gorm:"size:255;comment:Details;"`                        // Details
	Creator             *string         `json:"creator,omitempty" gorm:"size:255;comment:Creator;"`                            // Creator
	Modifier            *string         `json:"modifier,omitempty" gorm:"size:255;comment:Modifier;"`                          // Modifier
	CreatedAt           time.Time       `json:"created_at" gorm:"autoCreateTime;comment:Create timestamp;"`                    // Create timestamp
	UpdatedAt           time.Time       `json:"updated_at,omitempty" gorm:"autoUpdateTime;comment:Update timestamp;"`          // Update timestamp
	Deleted             string          `json:"-" gorm:"size:20;default:0;comment:Delete flag;"`                               // Delete flag
	DeletedAt           *gorm.DeletedAt `json:"-" gorm:"comment:Delete timestamp;"`                                            // Delete timestamp
}

func (a PolicyAuth) TableName() string {
	return config.C.FormatTableName("policy_auth")
}

// ConvertTo Convert `PolicyAuth` to `PolicyAuthForm` object.
func (a PolicyAuth) ConvertTo(form *PolicyAuthForm) error {
	form.ID = a.ID
	form.SpaceCode = a.SpaceCode
	form.SourceApplicationID = a.SourceApplicationID
	form.TargetServiceId = a.TargetServiceId
	form.Type = a.Type
	form.Params = a.Params
	form.Version = a.Version
	form.Enabled = a.Enabled
	form.Description = a.Description
	form.Creator = a.Creator
	form.Modifier = a.Modifier
	form.CreatedAt = a.CreatedAt
	form.UpdatedAt = a.UpdatedAt
	return nil
}

// Defining the query parameters for the `PolicyAuth` struct.
type PolicyAuthQueryParam struct {
	util.PaginationParam
	SpaceCode       string `form:"space_code"`        // Microservice space code
	TargetServiceId string `form:"target_service_id"` // Target service ID
}

// Defining the query options for the `PolicyAuth` struct.
type PolicyAuthQueryOptions struct {
	util.QueryOptions
}

// Defining the query result for the `PolicyAuth` struct.
type PolicyAuthQueryResult struct {
	Data       PolicyAuths
	PageResult *util.PaginationResult
}

// Defining the slice of `PolicyAuth` struct.
type PolicyAuths []*PolicyAuth

// Defining the data structure for creating a `PolicyAuth` struct.
type PolicyAuthForm struct {
	ID                  string    `json:"id"`
	SpaceCode           string    `json:"space_code" binding:"required,max=255"`       // Microservice space code
	SourceApplicationID *string   `json:"source_application_id"`                       // Source application ID
	TargetServiceId     string    `json:"target_service_id" binding:"required,max=20"` // Target service ID
	Type                string    `json:"type" binding:"required,max=20"`              // Auth type
	Params              *string   `json:"params"`                                      // Parameters (JSON)
	Version             int64     `json:"version"`                                     // Version
	Enabled             int       `json:"enabled"`                                     // Enabled
	Description         *string   `json:"description"`                                 // Details
	Creator             *string   `json:"creator,omitempty"`                           // Creator
	Modifier            *string   `json:"modifier,omitempty"`                          // Modifier
	CreatedAt           time.Time `json:"created_at"`                                  // Create timestamp
	UpdatedAt           time.Time `json:"updated_at,omitempty"`                        // Update timestamp
}

// A validation function for the `PolicyAuthForm` struct.
func (a *PolicyAuthForm) Validate() error {
	if a.SpaceCode == "" {
		return errors.BadRequest("", "SpaceCode is required")
	}
	if a.TargetServiceId == "" {
		return errors.BadRequest("", "TargetServiceId is required")
	}
	if a.Type == "" {
		return errors.BadRequest("", "Type is required")
	}
	return nil
}

// Convert `PolicyAuthForm` to `PolicyAuth` object.
func (a *PolicyAuthForm) FillTo(policyAuth *PolicyAuth) error {
	policyAuth.SpaceCode = a.SpaceCode
	policyAuth.SourceApplicationID = a.SourceApplicationID
	policyAuth.TargetServiceId = a.TargetServiceId
	policyAuth.Type = a.Type
	policyAuth.Params = a.Params
	policyAuth.Enabled = a.Enabled
	policyAuth.Description = a.Description
	policyAuth.Version = time.Now().UnixMilli()
	return nil
}
