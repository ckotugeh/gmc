package medical_specialties

import (
	"time"

	"gorm.io/gorm"
)

// MedicalSpecialty represents a medical specialization.
type MedicalSpecialty struct {
	ID uint `gorm:"primaryKey" json:"id"`

	// Basic information
	Name        string `gorm:"size:100;uniqueIndex;not null" json:"name"`
	Code        string `gorm:"size:20;uniqueIndex;not null" json:"code"`
	Description string `gorm:"type:text" json:"description"`

	// Status
	IsActive bool `gorm:"default:true" json:"is_active"`

	CreatedAt time.Time      `json:"created_at"`
	UpdatedAt time.Time      `json:"updated_at"`
	DeletedAt gorm.DeletedAt `gorm:"index" json:"-"`
}
