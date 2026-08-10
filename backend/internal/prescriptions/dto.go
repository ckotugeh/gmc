package prescriptions

import "time"

// PrescriptionItemRequest represents a medication included in a prescription.
type PrescriptionItemRequest struct {
	MedicationName string `json:"medication_name" binding:"required"`
	Dosage         string `json:"dosage" binding:"required"`
	Frequency      string `json:"frequency" binding:"required"`
	Duration       string `json:"duration" binding:"required"`
	Instructions   string `json:"instructions"`
	Quantity       int    `json:"quantity" binding:"required,gt=0"`
	Refills        int    `json:"refills"`
}

// CreatePrescriptionRequest represents a request to create a prescription.
type CreatePrescriptionRequest struct {
	DoctorID      uint                      `json:"doctor_id" binding:"required"`
	PatientID     uint                      `json:"patient_id" binding:"required"`
	AppointmentID uint                      `json:"appointment_id" binding:"required"`
	Diagnosis     string                    `json:"diagnosis" binding:"required"`
	Notes         string                    `json:"notes"`
	ExpiresAt     *time.Time                `json:"expires_at,omitempty"`
	Items         []PrescriptionItemRequest `json:"items" binding:"required,min=1"`
}

// UpdatePrescriptionRequest represents a request to update a prescription.
type UpdatePrescriptionRequest struct {
	Diagnosis string                    `json:"diagnosis"`
	Notes     string                    `json:"notes"`
	Status    PrescriptionStatus        `json:"status"`
	ExpiresAt *time.Time                `json:"expires_at,omitempty"`
	Items     []PrescriptionItemRequest `json:"items"`
}

// PrescriptionItemResponse represents a medication returned to the client.
type PrescriptionItemResponse struct {
	ID             uint   `json:"id"`
	MedicationName string `json:"medication_name"`
	Dosage         string `json:"dosage"`
	Frequency      string `json:"frequency"`
	Duration       string `json:"duration"`
	Instructions   string `json:"instructions"`
	Quantity       int    `json:"quantity"`
	Refills        int    `json:"refills"`
}

// PrescriptionResponse represents a prescription returned to the client.
type PrescriptionResponse struct {
	ID            uint                       `json:"id"`
	DoctorID      uint                       `json:"doctor_id"`
	PatientID     uint                       `json:"patient_id"`
	AppointmentID uint                       `json:"appointment_id"`
	Diagnosis     string                     `json:"diagnosis"`
	Notes         string                     `json:"notes"`
	Status        PrescriptionStatus         `json:"status"`
	IssuedAt      time.Time                  `json:"issued_at"`
	ExpiresAt     *time.Time                 `json:"expires_at,omitempty"`
	Items         []PrescriptionItemResponse `json:"items"`
	CreatedAt     time.Time                  `json:"created_at"`
	UpdatedAt     time.Time                  `json:"updated_at"`
}

// PrescriptionFilterRequest represents query filters.
type PrescriptionFilterRequest struct {
	DoctorID      uint               `form:"doctor_id"`
	PatientID     uint               `form:"patient_id"`
	AppointmentID uint               `form:"appointment_id"`
	Status        PrescriptionStatus `form:"status"`
}
