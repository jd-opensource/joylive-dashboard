package schema

import (
	"encoding/json"
	"time"

	"github.com/jd-opensource/joylive-dashboard/internal/config"
	"github.com/jd-opensource/joylive-dashboard/internal/mods/space/schema"
	"github.com/jd-opensource/joylive-dashboard/pkg/util"
	"gorm.io/gorm"
)

// ServiceExtra is a typed view over the Service.Extra JSON field.
// Known keys are exposed as typed fields; unrecognised keys are round-tripped
// transparently so that data written by other clients is never silently dropped.
type ServiceExtra struct {
	Authorized *int `json:"authorized,omitempty"` // 1 = auth enabled, 0 = disabled

	// unknown holds every key not matched by a named field above, enabling
	// safe round-trip of data we don't yet know about.
	unknown map[string]json.RawMessage
}

// UnmarshalJSON implements json.Unmarshaler.
// It deserialises all keys into the unknown map first, then promotes
// recognised keys into their typed struct fields.
func (e *ServiceExtra) UnmarshalJSON(data []byte) error {
	raw := make(map[string]json.RawMessage)
	if err := json.Unmarshal(data, &raw); err != nil {
		return err
	}

	if v, ok := raw["authorized"]; ok {
		var authorized int
		if err := json.Unmarshal(v, &authorized); err != nil {
			return err
		}
		e.Authorized = &authorized
		delete(raw, "authorized")
	}

	e.unknown = raw
	return nil
}

// MarshalJSON implements json.Marshaler.
// It merges the unknown keys with the known typed fields, giving typed fields
// precedence so an in-place mutation always wins over a stale unknown value.
func (e ServiceExtra) MarshalJSON() ([]byte, error) {
	out := make(map[string]json.RawMessage, len(e.unknown)+1)
	for k, v := range e.unknown {
		out[k] = v
	}
	if e.Authorized != nil {
		b, err := json.Marshal(*e.Authorized)
		if err != nil {
			return nil, err
		}
		out["authorized"] = b
	}
	return json.Marshal(out)
}

// ParseServiceExtra parses the raw JSON pointer stored in Service.Extra.
// Returns a zero-value ServiceExtra (all fields nil/unset) when raw is nil or empty.
func ParseServiceExtra(raw *string) (ServiceExtra, error) {
	if raw == nil || *raw == "" {
		return ServiceExtra{}, nil
	}
	var e ServiceExtra
	if err := json.Unmarshal([]byte(*raw), &e); err != nil {
		return ServiceExtra{}, err
	}
	return e, nil
}

// MarshalToJSON serialises ServiceExtra back to a JSON string pointer.
// Returns nil when the result would be an empty object, keeping the DB field clean.
func (e ServiceExtra) MarshalToJSON() (*string, error) {
	b, err := json.Marshal(e)
	if err != nil {
		return nil, err
	}
	s := string(b)
	if s == "{}" || s == "null" {
		return nil, nil
	}
	return &s, nil
}

// Service management for microservices
type Service struct {
	ID                       string          `json:"id" gorm:"size:20;primarykey;"`                               // Unique ID
	Name                     string          `json:"name" gorm:"size:255;uniqueIndex:uniq_svc_name"`              // Name
	SpaceCode                string          `json:"space_code" gorm:"size:255;uniqueIndex:uniq_svc_name"`        // Space code
	Space                    *schema.Space   `json:"space,omitempty" gorm:"foreignKey:SpaceCode;references:Code"` // Space
	RegistrationType         string          `json:"registration_type" gorm:"size:20"`                            // Registration type: HTTP, RPC_APP, RPC_INTERFACE
	Source                   string          `json:"source" gorm:"size:255"`                                      // Data source: Local, JSF, JDAP
	Tenant                   string          `json:"tenant" gorm:"size:255"`                                      // Tenant
	Creator                  string          `json:"creator" gorm:"size:255"`                                     // Creator
	Extra                    *string         `json:"extra,omitempty" gorm:"type:json"`                            // Extra info
	Authorized               *int            `json:"authorized,omitempty" gorm:"-"`                               // Parsed from Extra.authorized; not stored directly
	Version                  int64           `json:"version" gorm:"not null;default:1"`                           // Version
	Description              string          `json:"description" gorm:"size:255"`                                 // Description
	CreatedAt                time.Time       `json:"created_at" gorm:"index;"`                                    // Create time
	UpdatedAt                time.Time       `json:"updated_at" gorm:"index;"`                                    // Update time
	Deleted                  string          `json:"-" gorm:"size:20;uniqueIndex:uniq_svc_name;default:0"`        // Logical delete flag
	DeletedAt                *gorm.DeletedAt `json:"-" gorm:"comment:Delete time;"`                               // Delete time
	ApplicationServiceStatus string          `json:"application_service_status,omitempty" gorm:"->"`              // Application service status (virtual field, read-only)
	ApplicationName          string          `json:"application_name,omitempty" gorm:"->"`                        // Application name (virtual field, read-only)
	ApplicationId            string          `json:"application_id,omitempty" gorm:"->"`                          // Application ID (virtual field, read-only)
}

func (a *Service) TableName() string {
	return config.C.FormatTableName("service")
}

// ServiceQueryParam defines the query parameters for Service.
type ServiceQueryParam struct {
	util.PaginationParam
	LikeName  string `form:"name"`       // Name (like)
	SpaceCode string `form:"space_code"` // Space code
	Role      string `form:"role"`       // Filter by role of current user's applications: provider, consumer
	Creator   string `form:"-"`          // Current login username (set programmatically; used together with Role)
	UserID    string `form:"-"`          // User ID (set programmatically for permission filter)
	Tenant    string `form:"-"`          // Tenant (set programmatically for permission filter)
}

// ServiceQueryOptions defines the query options for Service.
type ServiceQueryOptions struct {
	util.QueryOptions
}

// ServiceQueryResult defines the query result for Service.
type ServiceQueryResult struct {
	Data       Services
	PageResult *util.PaginationResult
}

// Services defines a slice of Service.
type Services []*Service

// ServiceForm defines the form for creating/updating a Service.
type ServiceForm struct {
	Name             string  `json:"name" binding:"required,max=255"`             // Name
	SpaceCode        string  `json:"space_code" binding:"required,max=255"`       // Space code
	ApplicationId    string  `json:"application_id" binding:"required,max=20"`    // Application ID
	RegistrationType string  `json:"registration_type" binding:"required,max=20"` // Registration type
	Source           string  `json:"source"`                                      // Data source
	Extra            *string `json:"extra"`                                       // Extra info
	Description      string  `json:"description"`                                 // Description
}

func (a *ServiceForm) Validate() error {
	return nil
}

func (a *ServiceForm) FillTo(svc *Service) error {
	svc.Name = a.Name
	svc.SpaceCode = a.SpaceCode
	svc.RegistrationType = a.RegistrationType
	svc.Source = a.Source
	svc.Extra = toNilIfEmpty(a.Extra)
	svc.Description = a.Description
	return nil
}

func toNilIfEmpty(s *string) *string {
	if s != nil && *s == "" {
		return nil
	}
	return s
}
