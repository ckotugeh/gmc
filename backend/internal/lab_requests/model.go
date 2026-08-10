package lab_requests

import "time"

// LabRequest represents a laboratory test request made by a doctor.
type LabRequest struct {
	ID uint `gorm:"primaryKey" json:"id"`

	PatientID     uint `gorm:"not null;index" json:"patient_id"`
	DoctorID      uint `gorm:"not null;index" json:"doctor_id"`
	AppointmentID uint `gorm:"index" json:"appointment_id"`

	// Request details
	TestName string `gorm:"size:255;not null" json:"test_name"`        // e.g. CBC, Blood Sugar, Lipid Profile
	Category string `gorm:"size:100" json:"category"`                  // Hematology, Chemistry, Microbiology...
	Priority string `gorm:"size:20;default:'Routine'" json:"priority"` // Routine, Urgent, STAT

	ClinicalNotes string `gorm:"type:text" json:"clinical_notes,omitempty"`
	Reason        string `gorm:"type:text" json:"reason,omitempty"`

	Status string `gorm:"size:20;default:'Pending'" json:"status"` // Pending, Sample Collected, Processing, Completed, Cancelled

	RequestedAt time.Time `gorm:"not null" json:"requested_at"`

	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}
