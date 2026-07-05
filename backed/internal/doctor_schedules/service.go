package doctor_schedules

import (
	"errors"
	"strings"
)

var (
	ErrScheduleNotFound = errors.New("doctor schedule not found")
	ErrInvalidSchedule  = errors.New("invalid doctor schedule")
)

// Service defines the doctor schedule business logic.
type Service interface {
	Create(req CreateDoctorScheduleRequest) (*DoctorSchedule, error)
	GetByID(id uint) (*DoctorSchedule, error)
	GetAll() ([]DoctorSchedule, error)
	GetByDoctorID(doctorID uint) ([]DoctorSchedule, error)
	GetByDay(day Weekday) ([]DoctorSchedule, error)
	Update(id uint, req UpdateDoctorScheduleRequest) (*DoctorSchedule, error)
	Delete(id uint) error
}

type service struct {
	repo Repository
}

// NewService creates a new doctor schedule service.
func NewService(repo Repository) Service {
	return &service{
		repo: repo,
	}
}

// Create creates a new doctor schedule.
func (s *service) Create(req CreateDoctorScheduleRequest) (*DoctorSchedule, error) {
	if req.DoctorID == 0 ||
		req.Day == "" ||
		strings.TrimSpace(req.StartTime) == "" ||
		strings.TrimSpace(req.EndTime) == "" ||
		req.ConsultationDuration <= 0 ||
		req.MaxPatients <= 0 {
		return nil, ErrInvalidSchedule
	}

	schedule := &DoctorSchedule{
		DoctorID:             req.DoctorID,
		Day:                  req.Day,
		StartTime:            strings.TrimSpace(req.StartTime),
		EndTime:              strings.TrimSpace(req.EndTime),
		BreakStart:           req.BreakStart,
		BreakEnd:             req.BreakEnd,
		ConsultationDuration: req.ConsultationDuration,
		MaxPatients:          req.MaxPatients,
		IsActive:             true,
	}

	if req.IsActive != nil {
		schedule.IsActive = *req.IsActive
	}

	if err := s.repo.Create(schedule); err != nil {
		return nil, err
	}

	return schedule, nil
}

// GetByID retrieves a doctor schedule by ID.
func (s *service) GetByID(id uint) (*DoctorSchedule, error) {
	schedule, err := s.repo.GetByID(id)
	if err != nil {
		return nil, ErrScheduleNotFound
	}

	return schedule, nil
}

// GetAll retrieves all doctor schedules.
func (s *service) GetAll() ([]DoctorSchedule, error) {
	return s.repo.GetAll()
}

// GetByDoctorID retrieves schedules for a doctor.
func (s *service) GetByDoctorID(doctorID uint) ([]DoctorSchedule, error) {
	return s.repo.GetByDoctorID(doctorID)
}

// GetByDay retrieves schedules for a weekday.
func (s *service) GetByDay(day Weekday) ([]DoctorSchedule, error) {
	if day == "" {
		return nil, ErrInvalidSchedule
	}

	return s.repo.GetByDay(day)
}

// Update updates an existing doctor schedule.
func (s *service) Update(id uint, req UpdateDoctorScheduleRequest) (*DoctorSchedule, error) {
	schedule, err := s.repo.GetByID(id)
	if err != nil {
		return nil, ErrScheduleNotFound
	}

	if req.Day != "" {
		schedule.Day = req.Day
	}

	if start := strings.TrimSpace(req.StartTime); start != "" {
		schedule.StartTime = start
	}

	if end := strings.TrimSpace(req.EndTime); end != "" {
		schedule.EndTime = end
	}

	if req.BreakStart != nil {
		schedule.BreakStart = req.BreakStart
	}

	if req.BreakEnd != nil {
		schedule.BreakEnd = req.BreakEnd
	}

	if req.ConsultationDuration > 0 {
		schedule.ConsultationDuration = req.ConsultationDuration
	}

	if req.MaxPatients > 0 {
		schedule.MaxPatients = req.MaxPatients
	}

	if req.IsActive != nil {
		schedule.IsActive = *req.IsActive
	}

	if err := s.repo.Update(schedule); err != nil {
		return nil, err
	}

	return schedule, nil
}

// Delete removes a doctor schedule.
func (s *service) Delete(id uint) error {
	if _, err := s.repo.GetByID(id); err != nil {
		return ErrScheduleNotFound
	}

	return s.repo.Delete(id)
}
