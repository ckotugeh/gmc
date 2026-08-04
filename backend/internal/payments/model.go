package payments

import (
	"time"

	"gorm.io/gorm"
)

// PaymentStatus represents the status of a payment.
type PaymentStatus string

const (
	StatusPending   PaymentStatus = "pending"
	StatusPaid      PaymentStatus = "paid"
	StatusFailed    PaymentStatus = "failed"
	StatusRefunded  PaymentStatus = "refunded"
	StatusCancelled PaymentStatus = "cancelled"
)

// PaymentMethod represents the payment method used.
type PaymentMethod string

const (
	MethodCard         PaymentMethod = "card"
	MethodMobileMoney  PaymentMethod = "mobile_money"
	MethodBankTransfer PaymentMethod = "bank_transfer"
	MethodCash         PaymentMethod = "cash"
)

// Currency represents supported currencies.
type Currency string

const (
	CurrencyKES Currency = "KES"
	CurrencyUSD Currency = "USD"
	CurrencyEUR Currency = "EUR"
)

// Payment represents a payment transaction.
type Payment struct {
	ID uint `gorm:"primaryKey" json:"id"`

	// Appointment this payment belongs to.
	AppointmentID uint `gorm:"not null;index" json:"appointment_id"`

	// Patient making the payment.
	PatientID uint `gorm:"not null;index" json:"patient_id"`

	// Doctor receiving the payment.
	DoctorID uint `gorm:"not null;index" json:"doctor_id"`

	// Optional hospital receiving payment.
	HospitalID *uint `gorm:"index" json:"hospital_id,omitempty"`

	// Financial details.
	Amount   float64  `gorm:"type:decimal(12,2);not null" json:"amount"`
	Currency Currency `gorm:"type:varchar(10);default:'KES'" json:"currency"`

	Method PaymentMethod `gorm:"type:varchar(30);not null" json:"method"`
	Status PaymentStatus `gorm:"type:varchar(30);default:'pending'" json:"status"`

	// External transaction reference (Stripe, M-Pesa, etc.).
	TransactionReference string `gorm:"size:255;uniqueIndex" json:"transaction_reference"`

	// Optional notes.
	Description string `gorm:"type:text" json:"description"`

	// Payment timestamps.
	PaidAt *time.Time `json:"paid_at,omitempty"`

	CreatedAt time.Time      `json:"created_at"`
	UpdatedAt time.Time      `json:"updated_at"`
	DeletedAt gorm.DeletedAt `gorm:"index" json:"-"`
}
