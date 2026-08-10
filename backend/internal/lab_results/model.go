package lab_results

import "time"

// LabResult represents the outcome of a laboratory test.
type LabResult struct {
	ID uint `gorm:"primaryKey" json:"id"`

	LabRequestID uint `gorm:"not null;index" json:"lab_request_id"`

	PatientID uint `gorm:"not null;index" json:"patient_id"`
	DoctorID  uint `gorm:"not null;index" json:"doctor_id"`

	// Test information
	TestName string `gorm:"size:255;not null" json:"test_name"`

	Result string `gorm:"type:text;not null" json:"result"`

	ReferenceRange string `gorm:"size:255" json:"reference_range"`

	Units string `gorm:"size:50" json:"units"`

	Interpretation string `gorm:"type:text" json:"interpretation"` // Normal, Abnormal, Critical...

	Status string `gorm:"size:20;default:'Completed'" json:"status"` // Completed, Verified, Amended

	Remarks string `gorm:"type:text" json:"remarks,omitempty"`

	PerformedAt time.Time  `gorm:"not null" json:"performed_at"`
	VerifiedAt  *time.Time `json:"verified_at,omitempty"`

	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}
