package schema

import (
	"time"

	"github.com/jd-opensource/joylive-dashboard/internal/config"
	"github.com/jd-opensource/joylive-dashboard/pkg/errors"
	"github.com/jd-opensource/joylive-dashboard/pkg/util"
	"gorm.io/gorm"
)

// Service group management
type ServiceGroup struct {
	ID            string          `json:"id" gorm:"size:20;primaryKey;<-:create;comment:Unique ID;"`                                                       // Unique ID
	Name          string          `json:"name" gorm:"size:255;not null;comment:Group name;"`                                                               // Group name
	Code          string          `json:"code" gorm:"size:255;not null;uniqueIndex:uniq_service_group_code;comment:Group code;"`                           // Group code
	IsolationCode *string         `json:"isolationCode,omitempty" gorm:"size:255;comment:Isolation code;"`                                                 // Isolation code
	ServiceId     string          `json:"serviceId" gorm:"size:20;not null;uniqueIndex:uniq_service_group_code;comment:Service ID this group belongs to;"` // Service ID this group belongs to
	Labels        *string         `json:"labels,omitempty" gorm:"size:1024;comment:Labels;"`                                                               // Labels
	Description   *string         `json:"description,omitempty" gorm:"size:255;comment:Details;"`                                                          // Details
	CreatedAt     time.Time       `json:"createdAt" gorm:"autoCreateTime;comment:Create timestamp;"`                                                       // Create timestamp
	UpdatedAt     time.Time       `json:"updatedAt,omitempty" gorm:"autoUpdateTime;comment:Update timestamp;"`                                             // Update timestamp
	Deleted       string          `json:"-" gorm:"size:20;uniqueIndex:uniq_service_group_code;default:0;comment:Delete flag;"`                             // Delete flag
	DeletedAt     *gorm.DeletedAt `json:"-" gorm:"comment:Delete timestamp;"`                                                                              // Delete timestamp
}

func (a ServiceGroup) TableName() string {
	return config.C.FormatTableName("service_group")
}

// ServiceGroupQueryParam defines the query parameters for Service group.
type ServiceGroupQueryParam struct {
	util.PaginationParam
	ServiceId string `form:"service_id"` // Service ID
	Code      string `form:"code"`       // Group code
	LikeName  string `form:"name"`       // Group name (like)
}

// ServiceGroupQueryOptions defines the query options for Service group.
type ServiceGroupQueryOptions struct {
	util.QueryOptions
}

// ServiceGroupQueryResult defines the query result for Service group.
type ServiceGroupQueryResult struct {
	Data       ServiceGroups
	PageResult *util.PaginationResult
}

// ServiceGroups defines a slice of Service group.
type ServiceGroups []*ServiceGroup

// ServiceGroupForm defines the data structure for creating/updating a Service group.
type ServiceGroupForm struct {
	Name          string  `json:"name" binding:"required,max=255"`            // Group name
	Code          string  `json:"code" binding:"required,max=255"`            // Group code
	IsolationCode *string `json:"isolation_code" binding:"omitempty,max=255"` // Isolation code
	ServiceId     string  `json:"service_id" binding:"required,max=20"`       // Service ID this group belongs to
	Labels        *string `json:"labels" binding:"omitempty,max=1024"`        // Labels
	Description   *string `json:"description" binding:"omitempty,max=255"`    // Details
}

// A validation function for the `ServiceGroupForm` struct.
func (a *ServiceGroupForm) Validate() error {
	if a.Name == "" {
		return errors.BadRequest("", "Name is required")
	}
	if a.Code == "" {
		return errors.BadRequest("", "Code is required")
	}
	if a.ServiceId == "" {
		return errors.BadRequest("", "ServiceId is required")
	}
	return nil
}

// Convert `ServiceGroupForm` to `ServiceGroup` object.
func (a *ServiceGroupForm) FillTo(serviceGroup *ServiceGroup) error {
	serviceGroup.Name = a.Name
	serviceGroup.Code = a.Code
	serviceGroup.IsolationCode = a.IsolationCode
	serviceGroup.ServiceId = a.ServiceId
	serviceGroup.Labels = a.Labels
	serviceGroup.Description = a.Description
	return nil
}
