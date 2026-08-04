package prescriptions

import (
	"errors"
	"strings"
	"time"
)

var (
	ErrPrescriptionNotFound = errors.New("prescription not found")
	ErrInvalidPrescription  = errors.New("invalid prescription")
)

// Service defines the prescription business logic.
type Service interface {
	Create(req CreatePrescriptionRequest) (*Prescription, error)
	GetByID(id uint) (*Prescription, error)
	GetAll() ([]Prescription, error)

	GetByDoctorID(doctorID uint) ([]Prescription, error)
	GetByPatientID(patientID uint) ([]Prescription, error)
	GetByAppointmentID(appointmentID uint) ([]Prescription, error)

	Update(id uint, req UpdatePrescriptionRequest) (*Prescription, error)
	Delete(id uint) error
}

type service struct {
	repo Repository
}

// NewService creates a new prescription service.
func NewService(repo Repository) Service {
	return &service{
		repo: repo,
	}
}

// Create creates a new prescription.
func (s *service) Create(req CreatePrescriptionRequest) (*Prescription, error) {
	if req.DoctorID == 0 ||
		req.PatientID == 0 ||
		req.AppointmentID == 0 ||
		strings.TrimSpace(req.Diagnosis) == "" ||
		len(req.Items) == 0 {
		return nil, ErrInvalidPrescription
	}

	prescription := &Prescription{
		DoctorID:      req.DoctorID,
		PatientID:     req.PatientID,
		AppointmentID: req.AppointmentID,
		Diagnosis:     strings.TrimSpace(req.Diagnosis),
		Notes:         strings.TrimSpace(req.Notes),
		Status:        StatusActive,
		IssuedAt:      time.Now(),
		ExpiresAt:     req.ExpiresAt,
	}

	for _, item := range req.Items {
		if strings.TrimSpace(item.MedicationName) == "" ||
			strings.TrimSpace(item.Dosage) == "" ||
			strings.TrimSpace(item.Frequency) == "" ||
			strings.TrimSpace(item.Duration) == "" ||
			item.Quantity <= 0 {
			return nil, ErrInvalidPrescription
		}

		prescription.Items = append(prescription.Items, PrescriptionItem{
			MedicationName: strings.TrimSpace(item.MedicationName),
			Dosage:         strings.TrimSpace(item.Dosage),
			Frequency:      strings.TrimSpace(item.Frequency),
			Duration:       strings.TrimSpace(item.Duration),
			Instructions:   strings.TrimSpace(item.Instructions),
			Quantity:       item.Quantity,
			Refills:        item.Refills,
		})
	}

	if err := s.repo.Create(prescription); err != nil {
		return nil, err
	}

	return prescription, nil
}

// GetByID retrieves a prescription by ID.
func (s *service) GetByID(id uint) (*Prescription, error) {
	prescription, err := s.repo.GetByID(id)
	if err != nil {
		return nil, ErrPrescriptionNotFound
	}

	return prescription, nil
}

// GetAll retrieves all prescriptions.
func (s *service) GetAll() ([]Prescription, error) {
	return s.repo.GetAll()
}

// GetByDoctorID retrieves prescriptions by doctor.
func (s *service) GetByDoctorID(doctorID uint) ([]Prescription, error) {
	return s.repo.GetByDoctorID(doctorID)
}

// GetByPatientID retrieves prescriptions by patient.
func (s *service) GetByPatientID(patientID uint) ([]Prescription, error) {
	return s.repo.GetByPatientID(patientID)
}

// GetByAppointmentID retrieves prescriptions by appointment.
func (s *service) GetByAppointmentID(appointmentID uint) ([]Prescription, error) {
	return s.repo.GetByAppointmentID(appointmentID)
}

// Update updates an existing prescription.
func (s *service) Update(id uint, req UpdatePrescriptionRequest) (*Prescription, error) {
	prescription, err := s.repo.GetByID(id)
	if err != nil {
		return nil, ErrPrescriptionNotFound
	}

	if diagnosis := strings.TrimSpace(req.Diagnosis); diagnosis != "" {
		prescription.Diagnosis = diagnosis
	}

	if notes := strings.TrimSpace(req.Notes); notes != "" {
		prescription.Notes = notes
	}

	if req.Status != "" {
		prescription.Status = req.Status
	}

	if req.ExpiresAt != nil {
		prescription.ExpiresAt = req.ExpiresAt
	}

	if len(req.Items) > 0 {
		items := make([]PrescriptionItem, 0, len(req.Items))

		for _, item := range req.Items {
			if strings.TrimSpace(item.MedicationName) == "" ||
				strings.TrimSpace(item.Dosage) == "" ||
				strings.TrimSpace(item.Frequency) == "" ||
				strings.TrimSpace(item.Duration) == "" ||
				item.Quantity <= 0 {
				return nil, ErrInvalidPrescription
			}

			items = append(items, PrescriptionItem{
				MedicationName: strings.TrimSpace(item.MedicationName),
				Dosage:         strings.TrimSpace(item.Dosage),
				Frequency:      strings.TrimSpace(item.Frequency),
				Duration:       strings.TrimSpace(item.Duration),
				Instructions:   strings.TrimSpace(item.Instructions),
				Quantity:       item.Quantity,
				Refills:        item.Refills,
			})
		}

		prescription.Items = items
	}

	if err := s.repo.Update(prescription); err != nil {
		return nil, err
	}

	return prescription, nil
}

// Delete removes a prescription.
func (s *service) Delete(id uint) error {
	if _, err := s.repo.GetByID(id); err != nil {
		return ErrPrescriptionNotFound
	}

	return s.repo.Delete(id)
}
