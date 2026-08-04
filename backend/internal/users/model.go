package users

import (
	"time"

	"gorm.io/gorm"
)

// User represents a registered platform user.
type User struct {
	ID        uint           `gorm:"primaryKey" json:"id"`
	CreatedAt time.Time      `json:"created_at"`
	UpdatedAt time.Time      `json:"updated_at"`
	DeletedAt gorm.DeletedAt `gorm:"index" json:"-"`

	FirstName string `gorm:"size:100;not null" json:"first_name"`
	LastName  string `gorm:"size:100;not null" json:"last_name"`

	Email    string `gorm:"size:255;uniqueIndex;not null" json:"email"`
	Password string `gorm:"not null" json:"-"`

	Role  string `gorm:"size:20;default:'patient'" json:"role"`
	Phone string `gorm:"size:20" json:"phone"`

	IsVerified bool `gorm:"default:false" json:"is_verified"`
	IsActive   bool `gorm:"default:true" json:"is_active"`

	LastLoginAt *time.Time `json:"last_login_at,omitempty"`
}
