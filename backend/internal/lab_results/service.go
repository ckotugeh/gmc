package lab_results

import (
	"errors"
	"strings"
	"time"
)

var (
	ErrLabResultNotFound = errors.New("lab result not found")
	ErrInvalidLabResult  = errors.New("invalid lab result")
)

// Service defines the lab result business logic.
type Service interface {
	Create(req CreateLabResultRequest) (*LabResult, error)
	GetByID(id uint) (*LabResult, error)
	GetAll() ([]LabResult, error)
	GetByLabRequestID(labRequestID uint) ([]LabResult, error)
	GetByPatientID(patientID uint) ([]LabResult, error)
	GetByDoctorID(doctorID uint) ([]LabResult, error)
	GetByStatus(status string) ([]LabResult, error)
	Update(id uint, req UpdateLabResultRequest) (*LabResult, error)
	Delete(id uint) error
}

type service struct {
	repo Repository
}

// NewService creates a new lab result service.
func NewService(repo Repository) Service {
	return &service{
		repo: repo,
	}
}

// Create creates a new lab result.
func (s *service) Create(req CreateLabResultRequest) (*LabResult, error) {
	if req.LabRequestID == 0 ||
		req.PatientID == 0 ||
		req.DoctorID == 0 ||
		strings.TrimSpace(req.TestName) == "" ||
		strings.TrimSpace(req.Result) == "" {
		return nil, ErrInvalidLabResult
	}

	performedAt := req.PerformedAt
	if performedAt.IsZero() {
		performedAt = time.Now()
	}

	status := strings.TrimSpace(req.Status)
	if status == "" {
		status = "Completed"
	}

	result := &LabResult{
		LabRequestID:   req.LabRequestID,
		PatientID:      req.PatientID,
		DoctorID:       req.DoctorID,
		TestName:       strings.TrimSpace(req.TestName),
		Result:         strings.TrimSpace(req.Result),
		ReferenceRange: strings.TrimSpace(req.ReferenceRange),
		Units:          strings.TrimSpace(req.Units),
		Interpretation: strings.TrimSpace(req.Interpretation),
		Status:         status,
		Remarks:        strings.TrimSpace(req.Remarks),
		PerformedAt:    performedAt,
		VerifiedAt:     req.VerifiedAt,
	}

	if err := s.repo.Create(result); err != nil {
		return nil, err
	}

	return result, nil
}

// GetByID retrieves a lab result by ID.
func (s *service) GetByID(id uint) (*LabResult, error) {
	result, err := s.repo.GetByID(id)
	if err != nil {
		return nil, ErrLabResultNotFound
	}

	return result, nil
}

// GetAll retrieves all lab results.
func (s *service) GetAll() ([]LabResult, error) {
	return s.repo.GetAll()
}

// GetByLabRequestID retrieves results for a lab request.
func (s *service) GetByLabRequestID(labRequestID uint) ([]LabResult, error) {
	return s.repo.GetByLabRequestID(labRequestID)
}

// GetByPatientID retrieves results for a patient.
func (s *service) GetByPatientID(patientID uint) ([]LabResult, error) {
	return s.repo.GetByPatientID(patientID)
}

// GetByDoctorID retrieves results created by a doctor.
func (s *service) GetByDoctorID(doctorID uint) ([]LabResult, error) {
	return s.repo.GetByDoctorID(doctorID)
}

// GetByStatus retrieves results by status.
func (s *service) GetByStatus(status string) ([]LabResult, error) {
	return s.repo.GetByStatus(status)
}

// Update updates an existing lab result.
func (s *service) Update(id uint, req UpdateLabResultRequest) (*LabResult, error) {
	result, err := s.repo.GetByID(id)
	if err != nil {
		return nil, ErrLabResultNotFound
	}

	if req.TestName != nil {
		result.TestName = strings.TrimSpace(*req.TestName)
	}

	if req.Result != nil {
		result.Result = strings.TrimSpace(*req.Result)
	}

	if req.ReferenceRange != nil {
		result.ReferenceRange = strings.TrimSpace(*req.ReferenceRange)
	}

	if req.Units != nil {
		result.Units = strings.TrimSpace(*req.Units)
	}

	if req.Interpretation != nil {
		result.Interpretation = strings.TrimSpace(*req.Interpretation)
	}

	if req.Status != nil {
		result.Status = strings.TrimSpace(*req.Status)
	}

	if req.Remarks != nil {
		result.Remarks = strings.TrimSpace(*req.Remarks)
	}

	if req.PerformedAt != nil {
		result.PerformedAt = *req.PerformedAt
	}

	if req.VerifiedAt != nil {
		result.VerifiedAt = req.VerifiedAt
	}

	if err := s.repo.Update(result); err != nil {
		return nil, err
	}

	return result, nil
}

// Delete removes a lab result.
func (s *service) Delete(id uint) error {
	if _, err := s.repo.GetByID(id); err != nil {
		return ErrLabResultNotFound
	}

	return s.repo.Delete(id)
}
