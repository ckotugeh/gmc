package availability

import (
	"errors"
	"strings"
	"time"
)

var (
	ErrAvailabilityNotFound = errors.New("availability slot not found")
	ErrInvalidAvailability  = errors.New("invalid availability slot")
)

// Service defines the availability business logic.
type Service interface {
	Create(req CreateAvailabilityRequest) (*Availability, error)
	GetByID(id uint) (*Availability, error)
	GetAll() ([]Availability, error)

	GetByDoctorID(doctorID uint) ([]Availability, error)
	GetByScheduleID(scheduleID uint) ([]Availability, error)
	GetByDate(date time.Time) ([]Availability, error)
	GetByDoctorAndDate(doctorID uint, date time.Time) ([]Availability, error)

	Update(id uint, req UpdateAvailabilityRequest) (*Availability, error)
	Delete(id uint) error
}

type service struct {
	repo Repository
}

// NewService creates a new availability service.
func NewService(repo Repository) Service {
	return &service{
		repo: repo,
	}
}

// Create creates a new availability slot.
func (s *service) Create(req CreateAvailabilityRequest) (*Availability, error) {
	if req.DoctorID == 0 ||
		req.ScheduleID == 0 ||
		req.Date.IsZero() ||
		strings.TrimSpace(req.StartTime) == "" ||
		strings.TrimSpace(req.EndTime) == "" {
		return nil, ErrInvalidAvailability
	}

	status := req.Status
	if status == "" {
		status = SlotAvailable
	}

	slot := &Availability{
		DoctorID:      req.DoctorID,
		ScheduleID:    req.ScheduleID,
		Date:          req.Date,
		StartTime:     strings.TrimSpace(req.StartTime),
		EndTime:       strings.TrimSpace(req.EndTime),
		Status:        status,
		AppointmentID: req.AppointmentID,
		Notes:         strings.TrimSpace(req.Notes),
	}

	if err := s.repo.Create(slot); err != nil {
		return nil, err
	}

	return slot, nil
}

// GetByID retrieves an availability slot by ID.
func (s *service) GetByID(id uint) (*Availability, error) {
	slot, err := s.repo.GetByID(id)
	if err != nil {
		return nil, ErrAvailabilityNotFound
	}

	return slot, nil
}

// GetAll retrieves all availability slots.
func (s *service) GetAll() ([]Availability, error) {
	return s.repo.GetAll()
}

// GetByDoctorID retrieves availability slots for a doctor.
func (s *service) GetByDoctorID(doctorID uint) ([]Availability, error) {
	return s.repo.GetByDoctorID(doctorID)
}

// GetByScheduleID retrieves availability slots for a schedule.
func (s *service) GetByScheduleID(scheduleID uint) ([]Availability, error) {
	return s.repo.GetByScheduleID(scheduleID)
}

// GetByDate retrieves availability slots for a date.
func (s *service) GetByDate(date time.Time) ([]Availability, error) {
	if date.IsZero() {
		return nil, ErrInvalidAvailability
	}

	return s.repo.GetByDate(date)
}

// GetByDoctorAndDate retrieves availability slots for a doctor on a given date.
func (s *service) GetByDoctorAndDate(doctorID uint, date time.Time) ([]Availability, error) {
	if doctorID == 0 || date.IsZero() {
		return nil, ErrInvalidAvailability
	}

	return s.repo.GetByDoctorAndDate(doctorID, date)
}

// Update updates an availability slot.
func (s *service) Update(id uint, req UpdateAvailabilityRequest) (*Availability, error) {
	slot, err := s.repo.GetByID(id)
	if err != nil {
		return nil, ErrAvailabilityNotFound
	}

	if start := strings.TrimSpace(req.StartTime); start != "" {
		slot.StartTime = start
	}

	if end := strings.TrimSpace(req.EndTime); end != "" {
		slot.EndTime = end
	}

	if req.Status != "" {
		slot.Status = req.Status
	}

	if req.AppointmentID != nil {
		slot.AppointmentID = req.AppointmentID
	}

	if notes := strings.TrimSpace(req.Notes); notes != "" {
		slot.Notes = notes
	}

	if err := s.repo.Update(slot); err != nil {
		return nil, err
	}

	return slot, nil
}

// Delete removes an availability slot.
func (s *service) Delete(id uint) error {
	if _, err := s.repo.GetByID(id); err != nil {
		return ErrAvailabilityNotFound
	}

	return s.repo.Delete(id)
}
