package clinical_notes

import (
	"errors"
	"strings"
)

var (
	ErrClinicalNoteNotFound = errors.New("clinical note not found")
	ErrInvalidClinicalNote  = errors.New("invalid clinical note")
)

// Service defines the clinical note business logic.
type Service interface {
	Create(req CreateClinicalNoteRequest) (*ClinicalNote, error)
	GetByID(id uint) (*ClinicalNote, error)
	GetByAppointmentID(appointmentID uint) ([]ClinicalNote, error)
	GetByDoctorID(doctorID uint) ([]ClinicalNote, error)
	GetByPatientID(patientID uint) ([]ClinicalNote, error)
	GetByDiagnosisID(diagnosisID uint) ([]ClinicalNote, error)
	GetConfidential() ([]ClinicalNote, error)
	GetAll() ([]ClinicalNote, error)
	Update(id uint, req UpdateClinicalNoteRequest) (*ClinicalNote, error)
	Delete(id uint) error
}

type service struct {
	repo Repository
}

// NewService creates a new clinical note service.
func NewService(repo Repository) Service {
	return &service{
		repo: repo,
	}
}

// Create creates a new clinical note.
func (s *service) Create(req CreateClinicalNoteRequest) (*ClinicalNote, error) {
	if req.AppointmentID == 0 ||
		req.DoctorID == 0 ||
		req.PatientID == 0 ||
		strings.TrimSpace(req.Subject) == "" ||
		strings.TrimSpace(req.Note) == "" {
		return nil, ErrInvalidClinicalNote
	}

	isConfidential := true
	if req.IsConfidential != nil {
		isConfidential = *req.IsConfidential
	}

	note := &ClinicalNote{
		AppointmentID:  req.AppointmentID,
		DoctorID:       req.DoctorID,
		PatientID:      req.PatientID,
		DiagnosisID:    req.DiagnosisID,
		Subject:        strings.TrimSpace(req.Subject),
		Note:           strings.TrimSpace(req.Note),
		Assessment:     strings.TrimSpace(req.Assessment),
		Plan:           strings.TrimSpace(req.Plan),
		IsConfidential: isConfidential,
	}

	if err := s.repo.Create(note); err != nil {
		return nil, err
	}

	return note, nil
}

// GetByID retrieves a clinical note by ID.
func (s *service) GetByID(id uint) (*ClinicalNote, error) {
	note, err := s.repo.GetByID(id)
	if err != nil {
		return nil, ErrClinicalNoteNotFound
	}

	return note, nil
}

// GetByAppointmentID retrieves notes for an appointment.
func (s *service) GetByAppointmentID(appointmentID uint) ([]ClinicalNote, error) {
	return s.repo.GetByAppointmentID(appointmentID)
}

// GetByDoctorID retrieves notes created by a doctor.
func (s *service) GetByDoctorID(doctorID uint) ([]ClinicalNote, error) {
	return s.repo.GetByDoctorID(doctorID)
}

// GetByPatientID retrieves notes for a patient.
func (s *service) GetByPatientID(patientID uint) ([]ClinicalNote, error) {
	return s.repo.GetByPatientID(patientID)
}

// GetByDiagnosisID retrieves notes linked to a diagnosis.
func (s *service) GetByDiagnosisID(diagnosisID uint) ([]ClinicalNote, error) {
	return s.repo.GetByDiagnosisID(diagnosisID)
}

// GetConfidential retrieves all confidential notes.
func (s *service) GetConfidential() ([]ClinicalNote, error) {
	return s.repo.GetConfidential()
}

// GetAll retrieves all clinical notes.
func (s *service) GetAll() ([]ClinicalNote, error) {
	return s.repo.GetAll()
}

// Update updates an existing clinical note.
func (s *service) Update(id uint, req UpdateClinicalNoteRequest) (*ClinicalNote, error) {
	note, err := s.repo.GetByID(id)
	if err != nil {
		return nil, ErrClinicalNoteNotFound
	}

	if req.Subject != "" {
		note.Subject = strings.TrimSpace(req.Subject)
	}

	if req.Note != "" {
		note.Note = strings.TrimSpace(req.Note)
	}

	if req.Assessment != "" {
		note.Assessment = strings.TrimSpace(req.Assessment)
	}

	if req.Plan != "" {
		note.Plan = strings.TrimSpace(req.Plan)
	}

	if req.IsConfidential != nil {
		note.IsConfidential = *req.IsConfidential
	}

	if err := s.repo.Update(note); err != nil {
		return nil, err
	}

	return note, nil
}

// Delete removes a clinical note.
func (s *service) Delete(id uint) error {
	if _, err := s.repo.GetByID(id); err != nil {
		return ErrClinicalNoteNotFound
	}

	return s.repo.Delete(id)
}
