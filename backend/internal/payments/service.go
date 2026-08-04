package payments

import (
	"errors"
	"strings"
	"time"
)

var (
	ErrPaymentNotFound = errors.New("payment not found")
	ErrInvalidPayment  = errors.New("invalid payment")
)

// Service defines the payment business logic.
type Service interface {
	Create(req CreatePaymentRequest) (*Payment, error)
	GetByID(id uint) (*Payment, error)
	GetAll() ([]Payment, error)

	GetByAppointmentID(appointmentID uint) ([]Payment, error)
	GetByPatientID(patientID uint) ([]Payment, error)
	GetByDoctorID(doctorID uint) ([]Payment, error)
	GetByHospitalID(hospitalID uint) ([]Payment, error)

	Update(id uint, req UpdatePaymentRequest) (*Payment, error)
	Delete(id uint) error

	GetSummary() (*PaymentSummaryResponse, error)
}

type service struct {
	repo Repository
}

// NewService creates a new payment service.
func NewService(repo Repository) Service {
	return &service{
		repo: repo,
	}
}

// Create creates a new payment.
func (s *service) Create(req CreatePaymentRequest) (*Payment, error) {
	if req.AppointmentID == 0 ||
		req.PatientID == 0 ||
		req.DoctorID == 0 ||
		req.Amount <= 0 {
		return nil, ErrInvalidPayment
	}

	if req.Currency == "" {
		req.Currency = CurrencyKES
	}

	payment := &Payment{
		AppointmentID: req.AppointmentID,
		PatientID:     req.PatientID,
		DoctorID:      req.DoctorID,
		HospitalID:    req.HospitalID,
		Amount:        req.Amount,
		Currency:      req.Currency,
		Method:        req.Method,
		Status:        StatusPending,
		Description:   strings.TrimSpace(req.Description),
	}

	if err := s.repo.Create(payment); err != nil {
		return nil, err
	}

	return payment, nil
}

// GetByID retrieves a payment by ID.
func (s *service) GetByID(id uint) (*Payment, error) {
	payment, err := s.repo.GetByID(id)
	if err != nil {
		return nil, ErrPaymentNotFound
	}

	return payment, nil
}

// GetAll retrieves all payments.
func (s *service) GetAll() ([]Payment, error) {
	return s.repo.GetAll()
}

// GetByAppointmentID retrieves payments by appointment.
func (s *service) GetByAppointmentID(appointmentID uint) ([]Payment, error) {
	return s.repo.GetByAppointmentID(appointmentID)
}

// GetByPatientID retrieves payments by patient.
func (s *service) GetByPatientID(patientID uint) ([]Payment, error) {
	return s.repo.GetByPatientID(patientID)
}

// GetByDoctorID retrieves payments by doctor.
func (s *service) GetByDoctorID(doctorID uint) ([]Payment, error) {
	return s.repo.GetByDoctorID(doctorID)
}

// GetByHospitalID retrieves payments by hospital.
func (s *service) GetByHospitalID(hospitalID uint) ([]Payment, error) {
	return s.repo.GetByHospitalID(hospitalID)
}

// Update updates a payment.
func (s *service) Update(id uint, req UpdatePaymentRequest) (*Payment, error) {
	payment, err := s.repo.GetByID(id)
	if err != nil {
		return nil, ErrPaymentNotFound
	}

	if req.Status != "" {
		payment.Status = req.Status

		if req.Status == StatusPaid {
			now := time.Now()
			payment.PaidAt = &now
		}
	}

	if req.Method != "" {
		payment.Method = req.Method
	}

	if strings.TrimSpace(req.TransactionReference) != "" {
		payment.TransactionReference = strings.TrimSpace(req.TransactionReference)
	}

	if strings.TrimSpace(req.Description) != "" {
		payment.Description = strings.TrimSpace(req.Description)
	}

	if err := s.repo.Update(payment); err != nil {
		return nil, err
	}

	return payment, nil
}

// Delete removes a payment.
func (s *service) Delete(id uint) error {
	if _, err := s.repo.GetByID(id); err != nil {
		return ErrPaymentNotFound
	}

	return s.repo.Delete(id)
}

// GetSummary returns payment statistics.
func (s *service) GetSummary() (*PaymentSummaryResponse, error) {
	return s.repo.GetSummary()
}
