package medicalrecords

import (
	"time"

	"gorm.io/gorm"
)

// MedicalRecord represents a patient's medical record.
type MedicalRecord struct {
	ID uint `gorm:"primaryKey" json:"id"`

	// Relationships
	PatientID uint `gorm:"not null;index" json:"patient_id"`
	DoctorID  uint `gorm:"not null;index" json:"doctor_id"`

	// Medical information
	Diagnosis    string `gorm:"type:text;not null" json:"diagnosis"`
	Symptoms     string `gorm:"type:text" json:"symptoms"`
	Treatment    string `gorm:"type:text" json:"treatment"`
	Prescription string `gorm:"type:text" json:"prescription"`
	Notes        string `gorm:"type:text" json:"notes"`

	// Optional follow-up date
	FollowUpDate *time.Time `json:"follow_up_date,omitempty"`

	CreatedAt time.Time      `json:"created_at"`
	UpdatedAt time.Time      `json:"updated_at"`
	DeletedAt gorm.DeletedAt `gorm:"index" json:"-"`
}
