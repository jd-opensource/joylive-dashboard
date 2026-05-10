package schema

import (
	"fmt"
	"time"

	"github.com/jd-opensource/joylive-dashboard/internal/config"
	"github.com/jd-opensource/joylive-dashboard/pkg/util"
	"gorm.io/gorm"
)

// Data permission type constants
const (
	DataPermissionTypeApplication = "application"
	DataPermissionTypeService     = "service"
)

var dataPermissionTypes = map[string]bool{
	DataPermissionTypeApplication: true,
	DataPermissionTypeService:     true,
}

// Data permission management
type DataPermission struct {
	ID         string          `json:"id" gorm:"size:20;primaryKey;<-:create;comment:Unique ID;"`                                                    // Unique ID
	Type       string          `json:"type" gorm:"size:50;not null;uniqueIndex:uniq_data_permission;index:idx_type;comment:Data type (table name);"` // Data type (table name)
	DataId     string          `json:"dataId" gorm:"size:20;not null;uniqueIndex:uniq_data_permission;index:idx_data_id;comment:Data ID;"`           // Data ID
	User       string          `json:"user" gorm:"size:50;not null;uniqueIndex:uniq_data_permission;index:idx_user;comment:User;"`                   // User
	Tenant     string          `json:"tenant" gorm:"size:50;not null;uniqueIndex:uniq_data_permission;comment:Tenant;"`                              // Tenant
	Role       string          `json:"role" gorm:"size:20;not null;uniqueIndex:uniq_data_permission;comment:Role code;"`                             // Role code
	Permission uint            `json:"permission" gorm:"not null;default:0;comment:Data permission bits - format(read,write,delete);"`               // Data permission bits - format(read,write,delete)
	Creator    string          `json:"creator" gorm:"size:255"`                                                                                      // Creator
	CreatedAt  time.Time       `json:"createdAt" gorm:"autoCreateTime;comment:Create timestamp;"`                                                    // Create timestamp
	UpdatedAt  time.Time       `json:"updatedAt,omitempty" gorm:"autoUpdateTime;comment:Update timestamp;"`                                          // Update timestamp
	Deleted    string          `json:"-" gorm:"uniqueIndex:uniq_data_permission;size:20;default:0;comment:Delete flag;"`                             // Delete flag
	DeletedAt  *gorm.DeletedAt `json:"-" gorm:"comment:Delete timestamp;"`                                                                           // Delete timestamp
}

func (a DataPermission) TableName() string {
	return config.C.FormatTableName("data_permission")
}

// Defining the query parameters for the `DataPermission` struct.
type DataPermissionQueryParam struct {
	util.PaginationParam
}

// Defining the query options for the `DataPermission` struct.
type DataPermissionQueryOptions struct {
	util.QueryOptions
}

// Defining the query result for the `DataPermission` struct.
type DataPermissionQueryResult struct {
	Data       DataPermissions
	PageResult *util.PaginationResult
}

// Defining the slice of `DataPermission` struct.
type DataPermissions []*DataPermission

// Defining the data structure for creating a `DataPermission` struct.
type DataPermissionForm struct {
	Type       string `json:"type" binding:"required,max=50"`   // Data type (table name)
	DataId     string `json:"dataId" binding:"required,max=20"` // Data ID
	User       string `json:"user" binding:"required,max=50"`   // User
	Tenant     string `json:"tenant" binding:"required,max=50"` // Tenant
	Role       string `json:"role" binding:"required,max=20"`   // Role code
	Permission uint   `json:"permission"`                       // Data permission bits - format(read,write,delete)
}

// A validation function for the `DataPermissionForm` struct.
func (a *DataPermissionForm) Validate() error {
	if !dataPermissionTypes[a.Type] {
		return fmt.Errorf("invalid data permission type: %s, must be one of: application, service", a.Type)
	}
	return nil
}

// Convert `DataPermissionForm` to `DataPermission` object.
func (a *DataPermissionForm) FillTo(dataPermission *DataPermission) error {
	dataPermission.Type = a.Type
	dataPermission.DataId = a.DataId
	dataPermission.User = a.User
	dataPermission.Tenant = a.Tenant
	dataPermission.Role = a.Role
	dataPermission.Permission = a.Permission
	return nil
}
