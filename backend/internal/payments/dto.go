package payments

import "time"

// CreatePaymentRequest represents a request to create a payment.
type CreatePaymentRequest struct {
	AppointmentID uint          `json:"appointment_id" binding:"required"`
	PatientID     uint          `json:"patient_id" binding:"required"`
	DoctorID      uint          `json:"doctor_id" binding:"required"`
	HospitalID    *uint         `json:"hospital_id,omitempty"`
	Amount        float64       `json:"amount" binding:"required,gt=0"`
	Currency      Currency      `json:"currency"`
	Method        PaymentMethod `json:"method" binding:"required"`
	Description   string        `json:"description"`
}

// UpdatePaymentRequest represents a request to update a payment.
type UpdatePaymentRequest struct {
	Status               PaymentStatus `json:"status"`
	Method               PaymentMethod `json:"method"`
	TransactionReference string        `json:"transaction_reference"`
	Description          string        `json:"description"`
}

// PaymentResponse represents a payment returned to clients.
type PaymentResponse struct {
	ID                   uint          `json:"id"`
	AppointmentID        uint          `json:"appointment_id"`
	PatientID            uint          `json:"patient_id"`
	DoctorID             uint          `json:"doctor_id"`
	HospitalID           *uint         `json:"hospital_id,omitempty"`
	Amount               float64       `json:"amount"`
	Currency             Currency      `json:"currency"`
	Method               PaymentMethod `json:"method"`
	Status               PaymentStatus `json:"status"`
	TransactionReference string        `json:"transaction_reference"`
	Description          string        `json:"description"`
	PaidAt               *time.Time    `json:"paid_at,omitempty"`
	CreatedAt            time.Time     `json:"created_at"`
	UpdatedAt            time.Time     `json:"updated_at"`
}

// PaymentSummaryResponse represents payment statistics.
type PaymentSummaryResponse struct {
	TotalPayments  int64   `json:"total_payments"`
	TotalRevenue   float64 `json:"total_revenue"`
	PendingAmount  float64 `json:"pending_amount"`
	PaidAmount     float64 `json:"paid_amount"`
	RefundedAmount float64 `json:"refunded_amount"`
}

// PaymentFilterRequest represents payment query filters.
type PaymentFilterRequest struct {
	PatientID     uint          `form:"patient_id"`
	DoctorID      uint          `form:"doctor_id"`
	HospitalID    uint          `form:"hospital_id"`
	AppointmentID uint          `form:"appointment_id"`
	Status        PaymentStatus `form:"status"`
	Method        PaymentMethod `form:"method"`
}
