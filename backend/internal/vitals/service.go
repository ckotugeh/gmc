package vitals

import (
	"errors"
	"strings"
	"time"
)

var (
	ErrVitalNotFound = errors.New("vital record not found")
	ErrInvalidVital  = errors.New("invalid vital record")
)

// Service defines the vitals business logic.
type Service interface {
	Create(req CreateVitalRequest) (*Vital, error)
	GetByID(id uint) (*Vital, error)
	GetAll() ([]Vital, error)
	GetByPatientID(patientID uint) ([]Vital, error)
	GetByDoctorID(doctorID uint) ([]Vital, error)
	GetByAppointmentID(appointmentID uint) ([]Vital, error)
	Update(id uint, req UpdateVitalRequest) (*Vital, error)
	Delete(id uint) error
}

type service struct {
	repo Repository
}

// NewService creates a new vitals service.
func NewService(repo Repository) Service {
	return &service{
		repo: repo,
	}
}

// Create creates a new vital record.
func (s *service) Create(req CreateVitalRequest) (*Vital, error) {
	if req.PatientID == 0 || req.DoctorID == 0 {
		return nil, ErrInvalidVital
	}

	recordedAt := req.RecordedAt
	if recordedAt.IsZero() {
		recordedAt = time.Now()
	}

	vital := &Vital{
		PatientID:        req.PatientID,
		DoctorID:         req.DoctorID,
		AppointmentID:    req.AppointmentID,
		Temperature:      req.Temperature,
		HeartRate:        req.HeartRate,
		RespiratoryRate:  req.RespiratoryRate,
		SystolicBP:       req.SystolicBP,
		DiastolicBP:      req.DiastolicBP,
		OxygenSaturation: req.OxygenSaturation,
		Weight:           req.Weight,
		Height:           req.Height,
		BMI:              req.BMI,
		Notes:            strings.TrimSpace(req.Notes),
		RecordedAt:       recordedAt,
	}

	if err := s.repo.Create(vital); err != nil {
		return nil, err
	}

	return vital, nil
}

// GetByID retrieves a vital record by ID.
func (s *service) GetByID(id uint) (*Vital, error) {
	vital, err := s.repo.GetByID(id)
	if err != nil {
		return nil, ErrVitalNotFound
	}

	return vital, nil
}

// GetAll retrieves all vital records.
func (s *service) GetAll() ([]Vital, error) {
	return s.repo.GetAll()
}

// GetByPatientID retrieves all vitals for a patient.
func (s *service) GetByPatientID(patientID uint) ([]Vital, error) {
	return s.repo.GetByPatientID(patientID)
}

// GetByDoctorID retrieves all vitals recorded by a doctor.
func (s *service) GetByDoctorID(doctorID uint) ([]Vital, error) {
	return s.repo.GetByDoctorID(doctorID)
}

// GetByAppointmentID retrieves all vitals for an appointment.
func (s *service) GetByAppointmentID(appointmentID uint) ([]Vital, error) {
	return s.repo.GetByAppointmentID(appointmentID)
}

// Update updates an existing vital record.
func (s *service) Update(id uint, req UpdateVitalRequest) (*Vital, error) {
	vital, err := s.repo.GetByID(id)
	if err != nil {
		return nil, ErrVitalNotFound
	}

	if req.Temperature != nil {
		vital.Temperature = *req.Temperature
	}

	if req.HeartRate != nil {
		vital.HeartRate = *req.HeartRate
	}

	if req.RespiratoryRate != nil {
		vital.RespiratoryRate = *req.RespiratoryRate
	}

	if req.SystolicBP != nil {
		vital.SystolicBP = *req.SystolicBP
	}

	if req.DiastolicBP != nil {
		vital.DiastolicBP = *req.DiastolicBP
	}

	if req.OxygenSaturation != nil {
		vital.OxygenSaturation = *req.OxygenSaturation
	}

	if req.Weight != nil {
		vital.Weight = *req.Weight
	}

	if req.Height != nil {
		vital.Height = *req.Height
	}

	if req.BMI != nil {
		vital.BMI = *req.BMI
	}

	if req.Notes != nil {
		vital.Notes = strings.TrimSpace(*req.Notes)
	}

	if req.RecordedAt != nil {
		vital.RecordedAt = *req.RecordedAt
	}

	if err := s.repo.Update(vital); err != nil {
		return nil, err
	}

	return vital, nil
}

// Delete removes a vital record.
func (s *service) Delete(id uint) error {
	if _, err := s.repo.GetByID(id); err != nil {
		return ErrVitalNotFound
	}

	return s.repo.Delete(id)
}
