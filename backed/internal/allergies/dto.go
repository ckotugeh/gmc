package allergies

import "time"

// CreateAllergyRequest represents the payload for creating an allergy record.
type CreateAllergyRequest struct {
	PatientID uint `json:"patient_id" binding:"required"`
	DoctorID  uint `json:"doctor_id" binding:"required"`

	Allergen string `json:"allergen" binding:"required"`
	Type     string `json:"type" binding:"required"`
	Severity string `json:"severity" binding:"required"`
	Reaction string `json:"reaction"`
	Status   string `json:"status"`
	Notes    string `json:"notes"`

	RecordedAt time.Time `json:"recorded_at"`
}

// UpdateAllergyRequest represents the payload for updating an allergy record.
type UpdateAllergyRequest struct {
	Allergen   *string    `json:"allergen,omitempty"`
	Type       *string    `json:"type,omitempty"`
	Severity   *string    `json:"severity,omitempty"`
	Reaction   *string    `json:"reaction,omitempty"`
	Status     *string    `json:"status,omitempty"`
	Notes      *string    `json:"notes,omitempty"`
	RecordedAt *time.Time `json:"recorded_at,omitempty"`
}

// AllergyResponse represents an allergy returned by the API.
type AllergyResponse struct {
	ID uint `json:"id"`

	PatientID uint `json:"patient_id"`
	DoctorID  uint `json:"doctor_id"`

	Allergen string `json:"allergen"`
	Type     string `json:"type"`
	Severity string `json:"severity"`
	Reaction string `json:"reaction"`
	Status   string `json:"status"`
	Notes    string `json:"notes"`

	RecordedAt time.Time `json:"recorded_at"`
	CreatedAt  time.Time `json:"created_at"`
	UpdatedAt  time.Time `json:"updated_at"`
}
