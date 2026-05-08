package schema

import (
	"time"

	"github.com/jd-opensource/joylive-dashboard/internal/config"
	"github.com/jd-opensource/joylive-dashboard/pkg/util"
)

// Service management for microservices
type Service struct {
	ID               string     `json:"id" gorm:"size:20;primarykey;"`                        // Unique ID
	Name             string     `json:"name" gorm:"size:255;uniqueIndex:uniq_svc_name"`       // Name
	SpaceCode        string     `json:"space_code" gorm:"size:255;uniqueIndex:uniq_svc_name"` // Space code
	RegistrationType string     `json:"registration_type" gorm:"size:20"`                     // Registration type: HTTP, RPC_APP, RPC_INTERFACE
	Source           string     `json:"source" gorm:"size:255"`                               // Data source: Local, JSF, JDAP
	Tenant           string     `json:"tenant" gorm:"size:255"`                               // Tenant
	Creator          string     `json:"creator" gorm:"size:255"`                              // Creator
	Extra            *string    `json:"extra,omitempty" gorm:"type:json"`                     // Extra info
	Version          int64      `json:"version" gorm:"not null;default:1"`                    // Version
	Description      string     `json:"description" gorm:"size:255"`                          // Description
	CreatedAt        time.Time  `json:"created_at" gorm:"index;"`                             // Create time
	UpdatedAt        time.Time  `json:"updated_at" gorm:"index;"`                             // Update time
	Deleted          string     `json:"-" gorm:"size:20;uniqueIndex:uniq_svc_name;default:0"` // Logical delete flag
	DeletedAt        *time.Time `json:"-"`                                                    // Delete time
}

func (a *Service) TableName() string {
	return config.C.FormatTableName("service")
}

// ServiceQueryParam defines the query parameters for Service.
type ServiceQueryParam struct {
	util.PaginationParam
	LikeName  string `form:"name"`       // Name (like)
	SpaceCode string `form:"space_code"` // Space code
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
	Name             string `json:"name" binding:"required,max=255"`             // Name
	SpaceCode        string `json:"space_code" binding:"required,max=255"`       // Space code
	ApplicationId    string `json:"application_id" binding:"required,max=20"`    // Application ID
	RegistrationType string `json:"registration_type" binding:"required,max=20"` // Registration type
	Source           string `json:"source" binding:"required,max=255"`           // Data source
	Extra            *string `json:"extra"`                                       // Extra info
	Description      string `json:"description"`                                 // Description
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
