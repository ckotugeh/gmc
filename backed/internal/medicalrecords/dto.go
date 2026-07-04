package medicalrecords

import "time"

// CreateMedicalRecordRequest represents the payload for creating a medical record.
type CreateMedicalRecordRequest struct {
	PatientID    uint       `json:"patient_id" binding:"required"`
	DoctorID     uint       `json:"doctor_id" binding:"required"`
	Diagnosis    string     `json:"diagnosis" binding:"required"`
	Symptoms     string     `json:"symptoms"`
	Treatment    string     `json:"treatment"`
	Prescription string     `json:"prescription"`
	Notes        string     `json:"notes"`
	FollowUpDate *time.Time `json:"follow_up_date"`
}

// UpdateMedicalRecordRequest represents the payload for updating a medical record.
type UpdateMedicalRecordRequest struct {
	Diagnosis    string     `json:"diagnosis"`
	Symptoms     string     `json:"symptoms"`
	Treatment    string     `json:"treatment"`
	Prescription string     `json:"prescription"`
	Notes        string     `json:"notes"`
	FollowUpDate *time.Time `json:"follow_up_date"`
}

// MedicalRecordResponse represents the response returned to clients.
type MedicalRecordResponse struct {
	ID           uint       `json:"id"`
	PatientID    uint       `json:"patient_id"`
	DoctorID     uint       `json:"doctor_id"`
	Diagnosis    string     `json:"diagnosis"`
	Symptoms     string     `json:"symptoms"`
	Treatment    string     `json:"treatment"`
	Prescription string     `json:"prescription"`
	Notes        string     `json:"notes"`
	FollowUpDate *time.Time `json:"follow_up_date,omitempty"`
	CreatedAt    time.Time  `json:"created_at"`
	UpdatedAt    time.Time  `json:"updated_at"`
}
