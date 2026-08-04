package uploads

import (
	"time"

	"gorm.io/gorm"
)

// Upload represents a file uploaded to the platform.
type Upload struct {
	ID uint `gorm:"primaryKey" json:"id"`

	// Owner
	UserID uint `gorm:"not null;index" json:"user_id"`

	// Optional associations
	AppointmentID   *uint `gorm:"index" json:"appointment_id,omitempty"`
	MedicalRecordID *uint `gorm:"index" json:"medical_record_id,omitempty"`
	HospitalID      *uint `gorm:"index" json:"hospital_id,omitempty"`

	// File information
	FileName     string `gorm:"size:255;not null" json:"file_name"`
	OriginalName string `gorm:"size:255;not null" json:"original_name"`
	FileType     string `gorm:"size:100;not null" json:"file_type"`
	FileSize     int64  `gorm:"not null" json:"file_size"`
	FilePath     string `gorm:"size:500;not null" json:"file_path"`

	// Metadata
	Description string `gorm:"type:text" json:"description"`
	IsPublic    bool   `gorm:"default:false" json:"is_public"`

	CreatedAt time.Time      `json:"created_at"`
	UpdatedAt time.Time      `json:"updated_at"`
	DeletedAt gorm.DeletedAt `gorm:"index" json:"-"`
}
