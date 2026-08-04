package admin

import (
	"time"

	"gorm.io/gorm"
)

// AdminAction represents an administrative action.
type AdminAction string

const (
	ActionCreate  AdminAction = "create"
	ActionUpdate  AdminAction = "update"
	ActionDelete  AdminAction = "delete"
	ActionSuspend AdminAction = "suspend"
	ActionApprove AdminAction = "approve"
	ActionReject  AdminAction = "reject"
	ActionBan     AdminAction = "ban"
)

// ResourceType represents the resource affected by an admin action.
type ResourceType string

const (
	ResourceUser        ResourceType = "user"
	ResourceHospital    ResourceType = "hospital"
	ResourceCommunity   ResourceType = "community"
	ResourcePost        ResourceType = "post"
	ResourceComment     ResourceType = "comment"
	ResourceAppointment ResourceType = "appointment"
)

// Admin represents an admin audit log entry.
type Admin struct {
	ID uint `gorm:"primaryKey" json:"id"`

	// Administrator performing the action
	AdminID uint `gorm:"not null;index" json:"admin_id"`

	// Resource being affected
	ResourceID uint         `gorm:"not null;index" json:"resource_id"`
	Resource   ResourceType `gorm:"type:varchar(30);not null;index" json:"resource"`

	// Action performed
	Action AdminAction `gorm:"type:varchar(30);not null;index" json:"action"`

	// Optional description
	Description string `gorm:"type:text" json:"description"`

	CreatedAt time.Time      `json:"created_at"`
	UpdatedAt time.Time      `json:"updated_at"`
	DeletedAt gorm.DeletedAt `gorm:"index" json:"-"`
}
