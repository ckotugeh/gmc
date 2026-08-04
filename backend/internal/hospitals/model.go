package hospitals

import (
	"time"

	"gorm.io/gorm"
)

// Hospital represents a healthcare institution.
type Hospital struct {
	ID uint `gorm:"primaryKey" json:"id"`

	// Basic information
	Name        string `gorm:"size:255;not null" json:"name"`
	Description string `gorm:"type:text" json:"description"`

	// Contact information
	Email   string `gorm:"size:255;uniqueIndex" json:"email"`
	Phone   string `gorm:"size:50" json:"phone"`
	Website string `gorm:"size:255" json:"website"`

	// Address
	Address string `gorm:"size:255" json:"address"`
	City    string `gorm:"size:100" json:"city"`
	State   string `gorm:"size:100" json:"state"`
	Country string `gorm:"size:100" json:"country"`
	ZipCode string `gorm:"size:20" json:"zip_code"`

	// Additional details
	LicenseNumber string `gorm:"size:100;uniqueIndex" json:"license_number"`
	IsActive      bool   `gorm:"default:true" json:"is_active"`

	CreatedAt time.Time      `json:"created_at"`
	UpdatedAt time.Time      `json:"updated_at"`
	DeletedAt gorm.DeletedAt `gorm:"index" json:"-"`
}
