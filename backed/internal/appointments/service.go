package appointments

import (
	"errors"
)

// Service defines appointment business logic.
type Service interface {
	CreateAppointment(patientID uint, req CreateAppointmentRequest) (*Appointment, error)
	GetAppointment(id uint) (*Appointment, error)
	GetAppointments() ([]Appointment, error)
	GetDoctorAppointments(doctorID uint) ([]Appointment, error)
	GetPatientAppointments(patientID uint) ([]Appointment, error)
	UpdateAppointment(id uint, req UpdateAppointmentRequest) (*Appointment, error)
	DeleteAppointment(id uint) error
}

type service struct {
	repository Repository
}

// NewService creates a new appointment service.
func NewService(repository Repository) Service {
	return &service{
		repository: repository,
	}
}

// CreateAppointment creates a new appointment.
func (s *service) CreateAppointment(patientID uint, req CreateAppointmentRequest) (*Appointment, error) {
	if req.DoctorID == 0 {
		return nil, errors.New("doctor ID is required")
	}

	if patientID == 0 {
		return nil, errors.New("patient ID is required")
	}

	if req.DoctorID == patientID {
		return nil, errors.New("doctor and patient cannot be the same user")
	}

	if req.Reason == "" {
		return nil, errors.New("reason is required")
	}

	duration := req.DurationMinutes
	if duration <= 0 {
		duration = 30
	}

	appointment := &Appointment{
		DoctorID:        req.DoctorID,
		PatientID:       patientID,
		AppointmentTime: req.AppointmentTime,
		DurationMinutes: duration,
		Status:          StatusPending,
		Reason:          req.Reason,
		MeetingLink:     req.MeetingLink,
	}

	if err := s.repository.Create(appointment); err != nil {
		return nil, err
	}

	return appointment, nil
}

// GetAppointment retrieves an appointment by ID.
func (s *service) GetAppointment(id uint) (*Appointment, error) {
	return s.repository.GetByID(id)
}

// GetAppointments retrieves all appointments.
func (s *service) GetAppointments() ([]Appointment, error) {
	return s.repository.GetAll()
}

// GetDoctorAppointments retrieves appointments for a doctor.
func (s *service) GetDoctorAppointments(doctorID uint) ([]Appointment, error) {
	return s.repository.GetByDoctorID(doctorID)
}

// GetPatientAppointments retrieves appointments for a patient.
func (s *service) GetPatientAppointments(patientID uint) ([]Appointment, error) {
	return s.repository.GetByPatientID(patientID)
}

// UpdateAppointment updates an appointment.
func (s *service) UpdateAppointment(id uint, req UpdateAppointmentRequest) (*Appointment, error) {
	appointment, err := s.repository.GetByID(id)
	if err != nil {
		return nil, err
	}

	if req.AppointmentTime != nil {
		appointment.AppointmentTime = *req.AppointmentTime
	}

	if req.DurationMinutes != nil {
		appointment.DurationMinutes = *req.DurationMinutes
	}

	if req.Status != nil {
		appointment.Status = *req.Status
	}

	if req.Reason != nil {
		appointment.Reason = *req.Reason
	}

	if req.Notes != nil {
		appointment.Notes = *req.Notes
	}

	if req.MeetingLink != nil {
		appointment.MeetingLink = *req.MeetingLink
	}

	if err := s.repository.Update(appointment); err != nil {
		return nil, err
	}

	return appointment, nil
}

// DeleteAppointment deletes an appointment.
func (s *service) DeleteAppointment(id uint) error {
	return s.repository.Delete(id)
}
