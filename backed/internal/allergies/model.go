package allergies

import "time"

// Allergy represents a patient's allergy record.
type Allergy struct {
	ID uint `gorm:"primaryKey" json:"id"`

	PatientID uint `gorm:"not null;index" json:"patient_id"`
	DoctorID  uint `gorm:"not null;index" json:"doctor_id"`

	// Allergy information
	Allergen string `gorm:"size:255;not null" json:"allergen"` // e.g. Penicillin, Peanuts

	Type string `gorm:"size:50;not null" json:"type"` // Drug, Food, Environmental, Latex, Insect, Other

	Severity string `gorm:"size:20;not null" json:"severity"` // Mild, Moderate, Severe

	Reaction string `gorm:"type:text" json:"reaction"` // Rash, Anaphylaxis, Swelling...

	Status string `gorm:"size:20;default:'Active'" json:"status"` // Active, Resolved

	Notes string `gorm:"type:text" json:"notes,omitempty"`

	RecordedAt time.Time `gorm:"not null" json:"recorded_at"`

	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}
