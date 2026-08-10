package vitals

import "time"

// CreateVitalRequest represents the payload for creating a vital record.
type CreateVitalRequest struct {
	PatientID     uint `json:"patient_id" binding:"required"`
	DoctorID      uint `json:"doctor_id" binding:"required"`
	AppointmentID uint `json:"appointment_id"`

	Temperature      float64 `json:"temperature"`
	HeartRate        int     `json:"heart_rate"`
	RespiratoryRate  int     `json:"respiratory_rate"`
	SystolicBP       int     `json:"systolic_bp"`
	DiastolicBP      int     `json:"diastolic_bp"`
	OxygenSaturation int     `json:"oxygen_saturation"`
	Weight           float64 `json:"weight"`
	Height           float64 `json:"height"`
	BMI              float64 `json:"bmi"`

	Notes string `json:"notes"`

	RecordedAt time.Time `json:"recorded_at"`
}

// UpdateVitalRequest represents the payload for updating a vital record.
type UpdateVitalRequest struct {
	Temperature      *float64   `json:"temperature,omitempty"`
	HeartRate        *int       `json:"heart_rate,omitempty"`
	RespiratoryRate  *int       `json:"respiratory_rate,omitempty"`
	SystolicBP       *int       `json:"systolic_bp,omitempty"`
	DiastolicBP      *int       `json:"diastolic_bp,omitempty"`
	OxygenSaturation *int       `json:"oxygen_saturation,omitempty"`
	Weight           *float64   `json:"weight,omitempty"`
	Height           *float64   `json:"height,omitempty"`
	BMI              *float64   `json:"bmi,omitempty"`
	Notes            *string    `json:"notes,omitempty"`
	RecordedAt       *time.Time `json:"recorded_at,omitempty"`
}

// VitalResponse represents a vital record returned by the API.
type VitalResponse struct {
	ID uint `json:"id"`

	PatientID     uint `json:"patient_id"`
	DoctorID      uint `json:"doctor_id"`
	AppointmentID uint `json:"appointment_id"`

	Temperature      float64 `json:"temperature"`
	HeartRate        int     `json:"heart_rate"`
	RespiratoryRate  int     `json:"respiratory_rate"`
	SystolicBP       int     `json:"systolic_bp"`
	DiastolicBP      int     `json:"diastolic_bp"`
	OxygenSaturation int     `json:"oxygen_saturation"`
	Weight           float64 `json:"weight"`
	Height           float64 `json:"height"`
	BMI              float64 `json:"bmi"`

	Notes string `json:"notes"`

	RecordedAt time.Time `json:"recorded_at"`
	CreatedAt  time.Time `json:"created_at"`
	UpdatedAt  time.Time `json:"updated_at"`
}
