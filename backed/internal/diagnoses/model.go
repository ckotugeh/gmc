package diagnoses

import (
	"time"

	"gorm.io/gorm"
)

// Diagnosis represents a diagnosis made by a doctor for a patient.
type Diagnosis struct {
	ID uint `gorm:"primaryKey" json:"id"`

	// Relationships
	AppointmentID uint `gorm:"not null;index" json:"appointment_id"`
	DoctorID      uint `gorm:"not null;index" json:"doctor_id"`
	PatientID     uint `gorm:"not null;index" json:"patient_id"`

	// Diagnosis information
	DiagnosisCode string `gorm:"size:20;index" json:"diagnosis_code"` // e.g. ICD-10 code
	Condition     string `gorm:"size:255;not null" json:"condition"`
	Description   string `gorm:"type:text" json:"description"`

	// Clinical details
	Severity string `gorm:"size:50" json:"severity"`                // Mild, Moderate, Severe, Critical
	Status   string `gorm:"size:50;default:'Active'" json:"status"` // Active, Resolved, Chronic

	// Follow-up
	Notes        string     `gorm:"type:text" json:"notes"`
	FollowUpDate *time.Time `json:"follow_up_date,omitempty"`

	CreatedAt time.Time      `json:"created_at"`
	UpdatedAt time.Time      `json:"updated_at"`
	DeletedAt gorm.DeletedAt `gorm:"index" json:"-"`
}
