package lab_requests

import "time"

// CreateLabRequestRequest represents the payload for creating a lab request.
type CreateLabRequestRequest struct {
	PatientID     uint `json:"patient_id" binding:"required"`
	DoctorID      uint `json:"doctor_id" binding:"required"`
	AppointmentID uint `json:"appointment_id"`

	TestName string `json:"test_name" binding:"required"`
	Category string `json:"category"`
	Priority string `json:"priority"`

	ClinicalNotes string `json:"clinical_notes"`
	Reason        string `json:"reason"`

	Status string `json:"status"`

	RequestedAt time.Time `json:"requested_at"`
}

// UpdateLabRequestRequest represents the payload for updating a lab request.
type UpdateLabRequestRequest struct {
	TestName      *string    `json:"test_name,omitempty"`
	Category      *string    `json:"category,omitempty"`
	Priority      *string    `json:"priority,omitempty"`
	ClinicalNotes *string    `json:"clinical_notes,omitempty"`
	Reason        *string    `json:"reason,omitempty"`
	Status        *string    `json:"status,omitempty"`
	RequestedAt   *time.Time `json:"requested_at,omitempty"`
}

// LabRequestResponse represents a lab request returned by the API.
type LabRequestResponse struct {
	ID uint `json:"id"`

	PatientID     uint `json:"patient_id"`
	DoctorID      uint `json:"doctor_id"`
	AppointmentID uint `json:"appointment_id"`

	TestName string `json:"test_name"`
	Category string `json:"category"`
	Priority string `json:"priority"`

	ClinicalNotes string `json:"clinical_notes"`
	Reason        string `json:"reason"`

	Status string `json:"status"`

	RequestedAt time.Time `json:"requested_at"`
	CreatedAt   time.Time `json:"created_at"`
	UpdatedAt   time.Time `json:"updated_at"`
}
