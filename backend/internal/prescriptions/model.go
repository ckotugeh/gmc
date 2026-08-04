package prescriptions

import (
	"time"

	"gorm.io/gorm"
)

// PrescriptionStatus represents the status of a prescription.
type PrescriptionStatus string

const (
	StatusActive    PrescriptionStatus = "active"
	StatusCompleted PrescriptionStatus = "completed"
	StatusCancelled PrescriptionStatus = "cancelled"
	StatusExpired   PrescriptionStatus = "expired"
)

// Prescription represents a doctor's prescription for a patient.
type Prescription struct {
	ID uint `gorm:"primaryKey" json:"id"`

	// Relationships
	DoctorID      uint `gorm:"not null;index" json:"doctor_id"`
	PatientID     uint `gorm:"not null;index" json:"patient_id"`
	AppointmentID uint `gorm:"not null;index" json:"appointment_id"`

	// Prescription details
	Diagnosis string             `gorm:"type:text;not null" json:"diagnosis"`
	Notes     string             `gorm:"type:text" json:"notes"`
	Status    PrescriptionStatus `gorm:"type:varchar(20);default:'active'" json:"status"`

	// Validity
	IssuedAt  time.Time  `json:"issued_at"`
	ExpiresAt *time.Time `json:"expires_at,omitempty"`

	// Medicines prescribed
	Items []PrescriptionItem `gorm:"foreignKey:PrescriptionID;constraint:OnDelete:CASCADE" json:"items"`

	CreatedAt time.Time      `json:"created_at"`
	UpdatedAt time.Time      `json:"updated_at"`
	DeletedAt gorm.DeletedAt `gorm:"index" json:"-"`
}

// PrescriptionItem represents a single medication in a prescription.
type PrescriptionItem struct {
	ID uint `gorm:"primaryKey" json:"id"`

	PrescriptionID uint `gorm:"not null;index" json:"prescription_id"`

	MedicationName string `gorm:"size:255;not null" json:"medication_name"`
	Dosage         string `gorm:"size:100;not null" json:"dosage"`    // e.g. 500mg
	Frequency      string `gorm:"size:100;not null" json:"frequency"` // e.g. Twice daily
	Duration       string `gorm:"size:100;not null" json:"duration"`  // e.g. 7 days
	Instructions   string `gorm:"type:text" json:"instructions"`      // e.g. Take after meals
	Quantity       int    `gorm:"not null" json:"quantity"`
	Refills        int    `gorm:"default:0" json:"refills"`

	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}
