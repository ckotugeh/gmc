package clinical_notes

import (
	"time"

	"gorm.io/gorm"
)

// ClinicalNote represents notes recorded by a doctor during or after a consultation.
type ClinicalNote struct {
	ID uint `gorm:"primaryKey" json:"id"`

	// Relationships
	AppointmentID uint `gorm:"not null;index" json:"appointment_id"`
	DoctorID      uint `gorm:"not null;index" json:"doctor_id"`
	PatientID     uint `gorm:"not null;index" json:"patient_id"`
	DiagnosisID   uint `gorm:"index" json:"diagnosis_id"`

	// Clinical note
	Subject string `gorm:"size:255;not null" json:"subject"`
	Note    string `gorm:"type:text;not null" json:"note"`

	// Consultation details
	Assessment string `gorm:"type:text" json:"assessment"`
	Plan       string `gorm:"type:text" json:"plan"`

	// Visibility
	IsConfidential bool `gorm:"default:true" json:"is_confidential"`

	CreatedAt time.Time      `json:"created_at"`
	UpdatedAt time.Time      `json:"updated_at"`
	DeletedAt gorm.DeletedAt `gorm:"index" json:"-"`
}
