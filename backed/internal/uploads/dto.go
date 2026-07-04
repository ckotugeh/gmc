package uploads

// UploadResponse represents a file returned to the client.
type UploadResponse struct {
	ID              uint  `json:"id"`
	UserID          uint  `json:"user_id"`
	AppointmentID   *uint `json:"appointment_id,omitempty"`
	MedicalRecordID *uint `json:"medical_record_id,omitempty"`
	HospitalID      *uint `json:"hospital_id,omitempty"`

	FileName     string `json:"file_name"`
	OriginalName string `json:"original_name"`
	FileType     string `json:"file_type"`
	FileSize     int64  `json:"file_size"`
	FilePath     string `json:"file_path"`

	Description string `json:"description"`
	IsPublic    bool   `json:"is_public"`

	CreatedAt string `json:"created_at"`
	UpdatedAt string `json:"updated_at"`
}

// UpdateUploadRequest represents the payload for updating upload metadata.
type UpdateUploadRequest struct {
	Description string `json:"description"`
	IsPublic    *bool  `json:"is_public"`
}
