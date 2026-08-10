package vitals

import "time"

// Vital represents a patient's vital signs recorded during a consultation.
type Vital struct {
	ID uint `gorm:"primaryKey" json:"id"`

	PatientID     uint `gorm:"not null;index" json:"patient_id"`
	DoctorID      uint `gorm:"not null;index" json:"doctor_id"`
	AppointmentID uint `gorm:"index" json:"appointment_id"`

	// Vital signs
	Temperature      float64 `gorm:"type:decimal(4,1)" json:"temperature"` // °C
	HeartRate        int     `json:"heart_rate"`                           // bpm
	RespiratoryRate  int     `json:"respiratory_rate"`                     // breaths/min
	SystolicBP       int     `json:"systolic_bp"`                          // mmHg
	DiastolicBP      int     `json:"diastolic_bp"`                         // mmHg
	OxygenSaturation int     `json:"oxygen_saturation"`                    // %
	Weight           float64 `gorm:"type:decimal(5,2)" json:"weight"`      // kg
	Height           float64 `gorm:"type:decimal(5,2)" json:"height"`      // cm
	BMI              float64 `gorm:"type:decimal(5,2)" json:"bmi"`         // kg/m²

	Notes string `gorm:"type:text" json:"notes,omitempty"`

	RecordedAt time.Time `gorm:"not null" json:"recorded_at"`

	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}
