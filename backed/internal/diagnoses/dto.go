package diagnoses

import "time"

// CreateDiagnosisRequest represents a request to create a diagnosis.
type CreateDiagnosisRequest struct {
	AppointmentID uint `json:"appointment_id" binding:"required"`
	DoctorID      uint `json:"doctor_id" binding:"required"`
	PatientID     uint `json:"patient_id" binding:"required"`

	DiagnosisCode string `json:"diagnosis_code"`
	Condition     string `json:"condition" binding:"required"`
	Description   string `json:"description"`

	Severity     string     `json:"severity"`
	Status       string     `json:"status"`
	Notes        string     `json:"notes"`
	FollowUpDate *time.Time `json:"follow_up_date,omitempty"`
}

// UpdateDiagnosisRequest represents a request to update a diagnosis.
type UpdateDiagnosisRequest struct {
	DiagnosisCode string `json:"diagnosis_code,omitempty"`
	Condition     string `json:"condition,omitempty"`
	Description   string `json:"description,omitempty"`

	Severity     string     `json:"severity,omitempty"`
	Status       string     `json:"status,omitempty"`
	Notes        string     `json:"notes,omitempty"`
	FollowUpDate *time.Time `json:"follow_up_date,omitempty"`
}

// DiagnosisResponse represents a diagnosis returned to the client.
type DiagnosisResponse struct {
	ID uint `json:"id"`

	AppointmentID uint `json:"appointment_id"`
	DoctorID      uint `json:"doctor_id"`
	PatientID     uint `json:"patient_id"`

	DiagnosisCode string `json:"diagnosis_code"`
	Condition     string `json:"condition"`
	Description   string `json:"description"`

	Severity string `json:"severity"`
	Status   string `json:"status"`
	Notes    string `json:"notes"`

	FollowUpDate *time.Time `json:"follow_up_date,omitempty"`

	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}

// DiagnosisFilterRequest represents query parameters for filtering diagnoses.
type DiagnosisFilterRequest struct {
	DoctorID      uint   `form:"doctor_id"`
	PatientID     uint   `form:"patient_id"`
	AppointmentID uint   `form:"appointment_id"`
	Status        string `form:"status"`
	Severity      string `form:"severity"`
	Condition     string `form:"condition"`
}
