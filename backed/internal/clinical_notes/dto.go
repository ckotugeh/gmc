package clinical_notes

// CreateClinicalNoteRequest represents a request to create a clinical note.
type CreateClinicalNoteRequest struct {
	AppointmentID uint `json:"appointment_id" binding:"required"`
	DoctorID      uint `json:"doctor_id" binding:"required"`
	PatientID     uint `json:"patient_id" binding:"required"`
	DiagnosisID   uint `json:"diagnosis_id"`

	Subject string `json:"subject" binding:"required"`
	Note    string `json:"note" binding:"required"`

	Assessment string `json:"assessment"`
	Plan       string `json:"plan"`

	IsConfidential *bool `json:"is_confidential,omitempty"`
}

// UpdateClinicalNoteRequest represents a request to update a clinical note.
type UpdateClinicalNoteRequest struct {
	Subject string `json:"subject,omitempty"`
	Note    string `json:"note,omitempty"`

	Assessment string `json:"assessment,omitempty"`
	Plan       string `json:"plan,omitempty"`

	IsConfidential *bool `json:"is_confidential,omitempty"`
}

// ClinicalNoteResponse represents a clinical note returned to the client.
type ClinicalNoteResponse struct {
	ID uint `json:"id"`

	AppointmentID uint `json:"appointment_id"`
	DoctorID      uint `json:"doctor_id"`
	PatientID     uint `json:"patient_id"`
	DiagnosisID   uint `json:"diagnosis_id"`

	Subject string `json:"subject"`
	Note    string `json:"note"`

	Assessment string `json:"assessment"`
	Plan       string `json:"plan"`

	IsConfidential bool `json:"is_confidential"`
}

// ClinicalNoteFilterRequest represents query parameters for filtering clinical notes.
type ClinicalNoteFilterRequest struct {
	AppointmentID uint `form:"appointment_id"`
	DoctorID      uint `form:"doctor_id"`
	PatientID     uint `form:"patient_id"`
	DiagnosisID   uint `form:"diagnosis_id"`

	IsConfidential *bool `form:"is_confidential"`
}
