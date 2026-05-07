package schema

import (
	"time"

	"github.com/jd-opensource/joylive-dashboard/internal/config"
	"github.com/jd-opensource/joylive-dashboard/pkg/errors"
	"github.com/jd-opensource/joylive-dashboard/pkg/util"
	"gorm.io/gorm"
)

// Route policy detail management
type PolicyRouteDetail struct {
	ID           string          `json:"id" gorm:"size:20;primaryKey;<-:create;comment:Unique ID;"`                 // Unique ID
	RouteId      string          `json:"routeId" gorm:"size:20;not null;index:idx_routeid;comment:Route ID;"`       // Route ID
	RelationType string          `json:"relationType" gorm:"size:20;not null;comment:Relation type;"`               // Relation type
	Conditions   *string         `json:"conditions,omitempty" gorm:"type:json;comment:Match conditions (JSON);"`    // Match conditions (JSON)
	Destinations *string         `json:"destinations,omitempty" gorm:"type:json;comment:Destination rules (JSON);"` // Destination rules (JSON)
	Order        int             `json:"order" gorm:"not null;default:0;comment:Sort order;"`                       // Sort order
	Enabled      int             `json:"enabled" gorm:"not null;default:0;comment:Enabled;"`                        // Enabled
	Description  *string         `json:"description,omitempty" gorm:"size:255;comment:Details;"`                    // Details
	CreatedAt    time.Time       `json:"createdAt" gorm:"autoCreateTime;comment:Create timestamp;"`                 // Create timestamp
	UpdatedAt    time.Time       `json:"updatedAt,omitempty" gorm:"autoUpdateTime;comment:Update timestamp;"`       // Update timestamp
	Deleted      string          `json:"-" gorm:"index:idx_routeid;size:20;default:0;comment:Delete flag;"`         // Delete flag
	DeletedAt    *gorm.DeletedAt `json:"-" gorm:"comment:Delete timestamp;"`                                        // Delete timestamp
}

func (a PolicyRouteDetail) TableName() string {
	return config.C.FormatTableName("policy_route_detail")
}

// Defining the query parameters for the `PolicyRouteDetail` struct.
type PolicyRouteDetailQueryParam struct {
	util.PaginationParam
	RouteId string `form:"route_id"` // Route ID
}

// Defining the query options for the `PolicyRouteDetail` struct.
type PolicyRouteDetailQueryOptions struct {
	util.QueryOptions
}

// Defining the query result for the `PolicyRouteDetail` struct.
type PolicyRouteDetailQueryResult struct {
	Data       PolicyRouteDetails
	PageResult *util.PaginationResult
}

// Defining the slice of `PolicyRouteDetail` struct.
type PolicyRouteDetails []*PolicyRouteDetail

// Defining the data structure for creating a `PolicyRouteDetail` struct.
type PolicyRouteDetailForm struct {
	RouteId      string  `json:"routeId" binding:"required,max=20"`       // Route ID
	RelationType string  `json:"relationType" binding:"required,max=20"` // Relation type
	Conditions   *string `json:"conditions"`                              // Match conditions (JSON)
	Destinations *string `json:"destinations"`                            // Destination rules (JSON)
	Order        int     `json:"order"`                                   // Sort order
	Enabled      int     `json:"enabled"`                                 // Enabled
	Description  *string `json:"description"`                             // Details
}

// A validation function for the `PolicyRouteDetailForm` struct.
func (a *PolicyRouteDetailForm) Validate() error {
	if a.RouteId == "" {
		return errors.BadRequest("", "RouteId is required")
	}
	if a.RelationType == "" {
		return errors.BadRequest("", "RelationType is required")
	}
	return nil
}

// Convert `PolicyRouteDetailForm` to `PolicyRouteDetail` object.
func (a *PolicyRouteDetailForm) FillTo(policyRouteDetail *PolicyRouteDetail) error {
	policyRouteDetail.RouteId = a.RouteId
	policyRouteDetail.RelationType = a.RelationType
	policyRouteDetail.Conditions = a.Conditions
	policyRouteDetail.Destinations = a.Destinations
	policyRouteDetail.Order = a.Order
	policyRouteDetail.Enabled = a.Enabled
	policyRouteDetail.Description = a.Description
	return nil
}
