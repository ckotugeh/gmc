package allergies

import (
	"errors"
	"strings"
	"time"
)

var (
	ErrAllergyNotFound = errors.New("allergy record not found")
	ErrInvalidAllergy  = errors.New("invalid allergy record")
)

// Service defines the allergy business logic.
type Service interface {
	Create(req CreateAllergyRequest) (*Allergy, error)
	GetByID(id uint) (*Allergy, error)
	GetAll() ([]Allergy, error)
	GetByPatientID(patientID uint) ([]Allergy, error)
	GetByDoctorID(doctorID uint) ([]Allergy, error)
	GetBySeverity(severity string) ([]Allergy, error)
	GetActive() ([]Allergy, error)
	Update(id uint, req UpdateAllergyRequest) (*Allergy, error)
	Delete(id uint) error
}

type service struct {
	repo Repository
}

// NewService creates a new allergy service.
func NewService(repo Repository) Service {
	return &service{
		repo: repo,
	}
}

// Create creates a new allergy record.
func (s *service) Create(req CreateAllergyRequest) (*Allergy, error) {
	if req.PatientID == 0 ||
		req.DoctorID == 0 ||
		strings.TrimSpace(req.Allergen) == "" ||
		strings.TrimSpace(req.Type) == "" ||
		strings.TrimSpace(req.Severity) == "" {
		return nil, ErrInvalidAllergy
	}

	recordedAt := req.RecordedAt
	if recordedAt.IsZero() {
		recordedAt = time.Now()
	}

	status := strings.TrimSpace(req.Status)
	if status == "" {
		status = "Active"
	}

	allergy := &Allergy{
		PatientID:  req.PatientID,
		DoctorID:   req.DoctorID,
		Allergen:   strings.TrimSpace(req.Allergen),
		Type:       strings.TrimSpace(req.Type),
		Severity:   strings.TrimSpace(req.Severity),
		Reaction:   strings.TrimSpace(req.Reaction),
		Status:     status,
		Notes:      strings.TrimSpace(req.Notes),
		RecordedAt: recordedAt,
	}

	if err := s.repo.Create(allergy); err != nil {
		return nil, err
	}

	return allergy, nil
}

// GetByID retrieves an allergy by ID.
func (s *service) GetByID(id uint) (*Allergy, error) {
	allergy, err := s.repo.GetByID(id)
	if err != nil {
		return nil, ErrAllergyNotFound
	}

	return allergy, nil
}

// GetAll retrieves all allergy records.
func (s *service) GetAll() ([]Allergy, error) {
	return s.repo.GetAll()
}

// GetByPatientID retrieves allergies for a patient.
func (s *service) GetByPatientID(patientID uint) ([]Allergy, error) {
	return s.repo.GetByPatientID(patientID)
}

// GetByDoctorID retrieves allergies recorded by a doctor.
func (s *service) GetByDoctorID(doctorID uint) ([]Allergy, error) {
	return s.repo.GetByDoctorID(doctorID)
}

// GetBySeverity retrieves allergies by severity.
func (s *service) GetBySeverity(severity string) ([]Allergy, error) {
	return s.repo.GetBySeverity(severity)
}

// GetActive retrieves active allergies.
func (s *service) GetActive() ([]Allergy, error) {
	return s.repo.GetActive()
}

// Update updates an allergy record.
func (s *service) Update(id uint, req UpdateAllergyRequest) (*Allergy, error) {
	allergy, err := s.repo.GetByID(id)
	if err != nil {
		return nil, ErrAllergyNotFound
	}

	if req.Allergen != nil {
		allergy.Allergen = strings.TrimSpace(*req.Allergen)
	}

	if req.Type != nil {
		allergy.Type = strings.TrimSpace(*req.Type)
	}

	if req.Severity != nil {
		allergy.Severity = strings.TrimSpace(*req.Severity)
	}

	if req.Reaction != nil {
		allergy.Reaction = strings.TrimSpace(*req.Reaction)
	}

	if req.Status != nil {
		allergy.Status = strings.TrimSpace(*req.Status)
	}

	if req.Notes != nil {
		allergy.Notes = strings.TrimSpace(*req.Notes)
	}

	if req.RecordedAt != nil {
		allergy.RecordedAt = *req.RecordedAt
	}

	if err := s.repo.Update(allergy); err != nil {
		return nil, err
	}

	return allergy, nil
}

// Delete removes an allergy record.
func (s *service) Delete(id uint) error {
	if _, err := s.repo.GetByID(id); err != nil {
		return ErrAllergyNotFound
	}

	return s.repo.Delete(id)
}
