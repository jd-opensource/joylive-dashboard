package schema

import (
	"time"

	"github.com/jd-opensource/joylive-dashboard/internal/config"
	"github.com/jd-opensource/joylive-dashboard/pkg/errors"
	"github.com/jd-opensource/joylive-dashboard/pkg/util"
	"gorm.io/gorm"
)

// Application and service relationship management
type ApplicationService struct {
	ID            string          `json:"id" gorm:"size:20;primaryKey;<-:create;comment:Unique ID;"`                               // Unique ID
	ApplicationId string          `json:"application_id" gorm:"size:20;uniqueIndex:idx_name;comment:ApplicationId;"`               // ApplicationId
	ServiceId     string          `json:"service_id" gorm:"size:20;uniqueIndex:idx_name;comment:ServiceId;"`                       // ServiceId
	Role          string          `json:"role" gorm:"size:20;uniqueIndex:idx_name;comment:App role such as provider or consumer;"` // App role such as provider or consumer
	Status        string          `json:"status" gorm:"size:20;comment:消费状态: approved, rejected, pending;"`                        // 消费状态: approved, rejected, pending
	Description   *string         `json:"description" gorm:"size:255;comment:Details;"`                                            // Details
	CreatedAt     time.Time       `json:"created_at" gorm:"autoCreateTime;comment:Create timestamp;"`                              // Create timestamp
	UpdatedAt     time.Time       `json:"updated_at,omitempty" gorm:"autoUpdateTime;comment:Update timestamp;"`                    // Update timestamp
	Deleted       string          `json:"-" gorm:"uniqueIndex:idx_name;comment:Delete flag;size:20;default:0"`                     // Delete flag
	DeletedAt     *gorm.DeletedAt `json:"-" gorm:"comment:Delete timestamp;"`                                                      // Delete timestamp
}

func (a ApplicationService) TableName() string {
	return config.C.FormatTableName("application_service")
}

// Defining the query parameters for the `ApplicationService` struct.
type ApplicationServiceQueryParam struct {
	util.PaginationParam
}

// Defining the query options for the `ApplicationService` struct.
type ApplicationServiceQueryOptions struct {
	util.QueryOptions
}

// Defining the query result for the `ApplicationService` struct.
type ApplicationServiceQueryResult struct {
	Data       ApplicationServices
	PageResult *util.PaginationResult
}

// Defining the slice of `ApplicationService` struct.
type ApplicationServices []*ApplicationService

// Defining the data structure for creating a `ApplicationService` struct.
type ApplicationServiceForm struct {
	ApplicationId string  `json:"application_id" binding:"required,max=20"` // ApplicationId
	ServiceId     string  `json:"service_id" binding:"required,max=20"`     // ServiceId
	Role          string  `json:"role" binding:"required,max=20"`           // App role such as provider or consumer
	Status        string  `json:"status" binding:"max=20"`                  // Status: approved, rejected, pending
	Description   *string `json:"description"`                              // Details
}

// A validation function for the `ApplicationServiceForm` struct.
func (a *ApplicationServiceForm) Validate() error {
	if a.Role != "" && a.Role != "provider" && a.Role != "consumer" {
		return errors.BadRequest("", "Role must be provider or consumer")
	}
	if a.Status != "" && a.Status != "approved" && a.Status != "rejected" && a.Status != "pending" {
		return errors.BadRequest("", "Status must be approved, rejected, or pending")
	}
	return nil
}

// Convert `ApplicationServiceForm` to `ApplicationService` object.
func (a *ApplicationServiceForm) FillTo(applicationService *ApplicationService) error {
	applicationService.ApplicationId = a.ApplicationId
	applicationService.ServiceId = a.ServiceId
	applicationService.Role = a.Role
	if a.Status != "" {
		applicationService.Status = a.Status
	}
	applicationService.Description = a.Description
	return nil
}
