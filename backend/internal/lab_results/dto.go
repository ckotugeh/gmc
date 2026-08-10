package lab_results

import "time"

// CreateLabResultRequest represents the payload for creating a lab result.
type CreateLabResultRequest struct {
	LabRequestID uint `json:"lab_request_id" binding:"required"`

	PatientID uint `json:"patient_id" binding:"required"`
	DoctorID  uint `json:"doctor_id" binding:"required"`

	TestName string `json:"test_name" binding:"required"`

	Result string `json:"result" binding:"required"`

	ReferenceRange string `json:"reference_range"`
	Units          string `json:"units"`

	Interpretation string `json:"interpretation"`

	Status string `json:"status"`

	Remarks string `json:"remarks"`

	PerformedAt time.Time  `json:"performed_at"`
	VerifiedAt  *time.Time `json:"verified_at,omitempty"`
}

// UpdateLabResultRequest represents the payload for updating a lab result.
type UpdateLabResultRequest struct {
	TestName       *string    `json:"test_name,omitempty"`
	Result         *string    `json:"result,omitempty"`
	ReferenceRange *string    `json:"reference_range,omitempty"`
	Units          *string    `json:"units,omitempty"`
	Interpretation *string    `json:"interpretation,omitempty"`
	Status         *string    `json:"status,omitempty"`
	Remarks        *string    `json:"remarks,omitempty"`
	PerformedAt    *time.Time `json:"performed_at,omitempty"`
	VerifiedAt     *time.Time `json:"verified_at,omitempty"`
}

// LabResultResponse represents a lab result returned by the API.
type LabResultResponse struct {
	ID uint `json:"id"`

	LabRequestID uint `json:"lab_request_id"`

	PatientID uint `json:"patient_id"`
	DoctorID  uint `json:"doctor_id"`

	TestName string `json:"test_name"`

	Result string `json:"result"`

	ReferenceRange string `json:"reference_range"`
	Units          string `json:"units"`

	Interpretation string `json:"interpretation"`

	Status string `json:"status"`

	Remarks string `json:"remarks"`

	PerformedAt time.Time  `json:"performed_at"`
	VerifiedAt  *time.Time `json:"verified_at,omitempty"`

	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}
