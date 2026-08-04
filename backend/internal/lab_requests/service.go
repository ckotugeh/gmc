package lab_requests

import (
	"errors"
	"strings"
	"time"
)

var (
	ErrLabRequestNotFound = errors.New("lab request not found")
	ErrInvalidLabRequest  = errors.New("invalid lab request")
)

// Service defines the business logic for lab requests.
type Service interface {
	Create(req CreateLabRequestRequest) (*LabRequest, error)
	GetByID(id uint) (*LabRequest, error)
	GetAll() ([]LabRequest, error)
	GetByPatientID(patientID uint) ([]LabRequest, error)
	GetByDoctorID(doctorID uint) ([]LabRequest, error)
	GetByAppointmentID(appointmentID uint) ([]LabRequest, error)
	GetByStatus(status string) ([]LabRequest, error)
	Update(id uint, req UpdateLabRequestRequest) (*LabRequest, error)
	Delete(id uint) error
}

type service struct {
	repo Repository
}

// NewService creates a new lab request service.
func NewService(repo Repository) Service {
	return &service{
		repo: repo,
	}
}

// Create creates a new lab request.
func (s *service) Create(req CreateLabRequestRequest) (*LabRequest, error) {
	if req.PatientID == 0 || req.DoctorID == 0 || strings.TrimSpace(req.TestName) == "" {
		return nil, ErrInvalidLabRequest
	}

	if req.RequestedAt.IsZero() {
		req.RequestedAt = time.Now()
	}

	if strings.TrimSpace(req.Priority) == "" {
		req.Priority = "Routine"
	}

	if strings.TrimSpace(req.Status) == "" {
		req.Status = "Pending"
	}

	request := &LabRequest{
		PatientID:     req.PatientID,
		DoctorID:      req.DoctorID,
		AppointmentID: req.AppointmentID,
		TestName:      strings.TrimSpace(req.TestName),
		Category:      strings.TrimSpace(req.Category),
		Priority:      req.Priority,
		ClinicalNotes: strings.TrimSpace(req.ClinicalNotes),
		Reason:        strings.TrimSpace(req.Reason),
		Status:        req.Status,
		RequestedAt:   req.RequestedAt,
	}

	if err := s.repo.Create(request); err != nil {
		return nil, err
	}

	return request, nil
}

// GetByID retrieves a lab request by ID.
func (s *service) GetByID(id uint) (*LabRequest, error) {
	request, err := s.repo.GetByID(id)
	if err != nil {
		return nil, ErrLabRequestNotFound
	}

	return request, nil
}

// GetAll retrieves all lab requests.
func (s *service) GetAll() ([]LabRequest, error) {
	return s.repo.GetAll()
}

// GetByPatientID retrieves lab requests for a patient.
func (s *service) GetByPatientID(patientID uint) ([]LabRequest, error) {
	return s.repo.GetByPatientID(patientID)
}

// GetByDoctorID retrieves lab requests by doctor.
func (s *service) GetByDoctorID(doctorID uint) ([]LabRequest, error) {
	return s.repo.GetByDoctorID(doctorID)
}

// GetByAppointmentID retrieves lab requests for an appointment.
func (s *service) GetByAppointmentID(appointmentID uint) ([]LabRequest, error) {
	return s.repo.GetByAppointmentID(appointmentID)
}

// GetByStatus retrieves lab requests by status.
func (s *service) GetByStatus(status string) ([]LabRequest, error) {
	return s.repo.GetByStatus(status)
}

// Update updates an existing lab request.
func (s *service) Update(id uint, req UpdateLabRequestRequest) (*LabRequest, error) {
	request, err := s.repo.GetByID(id)
	if err != nil {
		return nil, ErrLabRequestNotFound
	}

	if req.TestName != nil {
		request.TestName = strings.TrimSpace(*req.TestName)
	}

	if req.Category != nil {
		request.Category = strings.TrimSpace(*req.Category)
	}

	if req.Priority != nil {
		request.Priority = strings.TrimSpace(*req.Priority)
	}

	if req.ClinicalNotes != nil {
		request.ClinicalNotes = strings.TrimSpace(*req.ClinicalNotes)
	}

	if req.Reason != nil {
		request.Reason = strings.TrimSpace(*req.Reason)
	}

	if req.Status != nil {
		request.Status = strings.TrimSpace(*req.Status)
	}

	if req.RequestedAt != nil {
		request.RequestedAt = *req.RequestedAt
	}

	if err := s.repo.Update(request); err != nil {
		return nil, err
	}

	return request, nil
}

// Delete removes a lab request.
func (s *service) Delete(id uint) error {
	if _, err := s.repo.GetByID(id); err != nil {
		return ErrLabRequestNotFound
	}

	return s.repo.Delete(id)
}
