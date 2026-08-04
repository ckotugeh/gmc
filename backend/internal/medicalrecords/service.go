package medicalrecords

import (
	"errors"
	"strings"
)

var (
	ErrMedicalRecordNotFound = errors.New("medical record not found")
	ErrInvalidMedicalRecord  = errors.New("invalid medical record")
)

// Service defines the medical record business logic.
type Service interface {
	Create(req CreateMedicalRecordRequest) (*MedicalRecord, error)
	GetByID(id uint) (*MedicalRecord, error)
	GetAll() ([]MedicalRecord, error)
	GetByPatientID(patientID uint) ([]MedicalRecord, error)
	GetByDoctorID(doctorID uint) ([]MedicalRecord, error)
	Update(id uint, req UpdateMedicalRecordRequest) (*MedicalRecord, error)
	Delete(id uint) error
}

type service struct {
	repo Repository
}

// NewService creates a new medical record service.
func NewService(repo Repository) Service {
	return &service{repo: repo}
}

// Create creates a new medical record.
func (s *service) Create(req CreateMedicalRecordRequest) (*MedicalRecord, error) {
	if req.PatientID == 0 ||
		req.DoctorID == 0 ||
		strings.TrimSpace(req.Diagnosis) == "" {
		return nil, ErrInvalidMedicalRecord
	}

	record := &MedicalRecord{
		PatientID:    req.PatientID,
		DoctorID:     req.DoctorID,
		Diagnosis:    req.Diagnosis,
		Symptoms:     req.Symptoms,
		Treatment:    req.Treatment,
		Prescription: req.Prescription,
		Notes:        req.Notes,
		FollowUpDate: req.FollowUpDate,
	}

	if err := s.repo.Create(record); err != nil {
		return nil, err
	}

	return record, nil
}

// GetByID retrieves a medical record by ID.
func (s *service) GetByID(id uint) (*MedicalRecord, error) {
	return s.repo.GetByID(id)
}

// GetAll retrieves all medical records.
func (s *service) GetAll() ([]MedicalRecord, error) {
	return s.repo.GetAll()
}

// GetByPatientID retrieves all records for a patient.
func (s *service) GetByPatientID(patientID uint) ([]MedicalRecord, error) {
	return s.repo.GetByPatientID(patientID)
}

// GetByDoctorID retrieves all records created by a doctor.
func (s *service) GetByDoctorID(doctorID uint) ([]MedicalRecord, error) {
	return s.repo.GetByDoctorID(doctorID)
}

// Update updates a medical record.
func (s *service) Update(id uint, req UpdateMedicalRecordRequest) (*MedicalRecord, error) {
	record, err := s.repo.GetByID(id)
	if err != nil {
		return nil, ErrMedicalRecordNotFound
	}

	if strings.TrimSpace(req.Diagnosis) != "" {
		record.Diagnosis = req.Diagnosis
	}
	if req.Symptoms != "" {
		record.Symptoms = req.Symptoms
	}
	if req.Treatment != "" {
		record.Treatment = req.Treatment
	}
	if req.Prescription != "" {
		record.Prescription = req.Prescription
	}
	if req.Notes != "" {
		record.Notes = req.Notes
	}
	if req.FollowUpDate != nil {
		record.FollowUpDate = req.FollowUpDate
	}

	if err := s.repo.Update(record); err != nil {
		return nil, err
	}

	return record, nil
}

// Delete removes a medical record.
func (s *service) Delete(id uint) error {
	if _, err := s.repo.GetByID(id); err != nil {
		return ErrMedicalRecordNotFound
	}

	return s.repo.Delete(id)
}
