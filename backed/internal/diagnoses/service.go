package diagnoses

import (
	"errors"
	"strings"
)

var (
	ErrDiagnosisNotFound = errors.New("diagnosis not found")
	ErrInvalidDiagnosis  = errors.New("invalid diagnosis")
)

// Service defines the diagnosis business logic.
type Service interface {
	Create(req CreateDiagnosisRequest) (*Diagnosis, error)
	GetByID(id uint) (*Diagnosis, error)
	GetByAppointmentID(appointmentID uint) ([]Diagnosis, error)
	GetByDoctorID(doctorID uint) ([]Diagnosis, error)
	GetByPatientID(patientID uint) ([]Diagnosis, error)
	GetByStatus(status string) ([]Diagnosis, error)
	GetAll() ([]Diagnosis, error)
	Update(id uint, req UpdateDiagnosisRequest) (*Diagnosis, error)
	Delete(id uint) error
}

type service struct {
	repo Repository
}

// NewService creates a new diagnosis service.
func NewService(repo Repository) Service {
	return &service{
		repo: repo,
	}
}

// Create creates a new diagnosis.
func (s *service) Create(req CreateDiagnosisRequest) (*Diagnosis, error) {
	if req.AppointmentID == 0 ||
		req.DoctorID == 0 ||
		req.PatientID == 0 ||
		strings.TrimSpace(req.Condition) == "" {
		return nil, ErrInvalidDiagnosis
	}

	status := strings.TrimSpace(req.Status)
	if status == "" {
		status = "Active"
	}

	diagnosis := &Diagnosis{
		AppointmentID: req.AppointmentID,
		DoctorID:      req.DoctorID,
		PatientID:     req.PatientID,
		DiagnosisCode: strings.TrimSpace(req.DiagnosisCode),
		Condition:     strings.TrimSpace(req.Condition),
		Description:   strings.TrimSpace(req.Description),
		Severity:      strings.TrimSpace(req.Severity),
		Status:        status,
		Notes:         strings.TrimSpace(req.Notes),
		FollowUpDate:  req.FollowUpDate,
	}

	if err := s.repo.Create(diagnosis); err != nil {
		return nil, err
	}

	return diagnosis, nil
}

// GetByID retrieves a diagnosis by ID.
func (s *service) GetByID(id uint) (*Diagnosis, error) {
	diagnosis, err := s.repo.GetByID(id)
	if err != nil {
		return nil, ErrDiagnosisNotFound
	}

	return diagnosis, nil
}

// GetByAppointmentID retrieves diagnoses for an appointment.
func (s *service) GetByAppointmentID(appointmentID uint) ([]Diagnosis, error) {
	return s.repo.GetByAppointmentID(appointmentID)
}

// GetByDoctorID retrieves diagnoses created by a doctor.
func (s *service) GetByDoctorID(doctorID uint) ([]Diagnosis, error) {
	return s.repo.GetByDoctorID(doctorID)
}

// GetByPatientID retrieves diagnoses for a patient.
func (s *service) GetByPatientID(patientID uint) ([]Diagnosis, error) {
	return s.repo.GetByPatientID(patientID)
}

// GetByStatus retrieves diagnoses by status.
func (s *service) GetByStatus(status string) ([]Diagnosis, error) {
	return s.repo.GetByStatus(status)
}

// GetAll retrieves all diagnoses.
func (s *service) GetAll() ([]Diagnosis, error) {
	return s.repo.GetAll()
}

// Update updates an existing diagnosis.
func (s *service) Update(id uint, req UpdateDiagnosisRequest) (*Diagnosis, error) {
	diagnosis, err := s.repo.GetByID(id)
	if err != nil {
		return nil, ErrDiagnosisNotFound
	}

	if req.DiagnosisCode != "" {
		diagnosis.DiagnosisCode = strings.TrimSpace(req.DiagnosisCode)
	}

	if req.Condition != "" {
		diagnosis.Condition = strings.TrimSpace(req.Condition)
	}

	if req.Description != "" {
		diagnosis.Description = strings.TrimSpace(req.Description)
	}

	if req.Severity != "" {
		diagnosis.Severity = strings.TrimSpace(req.Severity)
	}

	if req.Status != "" {
		diagnosis.Status = strings.TrimSpace(req.Status)
	}

	if req.Notes != "" {
		diagnosis.Notes = strings.TrimSpace(req.Notes)
	}

	if req.FollowUpDate != nil {
		diagnosis.FollowUpDate = req.FollowUpDate
	}

	if err := s.repo.Update(diagnosis); err != nil {
		return nil, err
	}

	return diagnosis, nil
}

// Delete removes a diagnosis.
func (s *service) Delete(id uint) error {
	if _, err := s.repo.GetByID(id); err != nil {
		return ErrDiagnosisNotFound
	}

	return s.repo.Delete(id)
}
