package video_consultations

import (
	"errors"
	"fmt"
	"time"
)

var (
	ErrConsultationNotFound = errors.New("video consultation not found")
	ErrInvalidConsultation  = errors.New("invalid video consultation")
	ErrConsultationExists   = errors.New("video consultation already exists for this appointment")
)

// Service defines the video consultation business logic.
type Service interface {
	Create(req CreateVideoConsultationRequest) (*VideoConsultation, error)
	GetByID(id uint) (*VideoConsultation, error)
	GetByAppointmentID(appointmentID uint) (*VideoConsultation, error)
	GetByDoctorID(doctorID uint) ([]VideoConsultation, error)
	GetByPatientID(patientID uint) ([]VideoConsultation, error)
	GetAll() ([]VideoConsultation, error)
	Update(id uint, req UpdateVideoConsultationRequest) (*VideoConsultation, error)
	Delete(id uint) error
}

type service struct {
	repo Repository
}

// NewService creates a new video consultation service.
func NewService(repo Repository) Service {
	return &service{
		repo: repo,
	}
}

// Create creates a new video consultation.
func (s *service) Create(req CreateVideoConsultationRequest) (*VideoConsultation, error) {
	if req.AppointmentID == 0 ||
		req.DoctorID == 0 ||
		req.PatientID == 0 ||
		req.ScheduledAt.IsZero() {
		return nil, ErrInvalidConsultation
	}

	if existing, _ := s.repo.GetByAppointmentID(req.AppointmentID); existing != nil {
		return nil, ErrConsultationExists
	}

	roomID := fmt.Sprintf(
		"consultation-%d-%d",
		req.AppointmentID,
		time.Now().UnixNano(),
	)

	consultation := &VideoConsultation{
		AppointmentID: req.AppointmentID,
		DoctorID:      req.DoctorID,
		PatientID:     req.PatientID,
		RoomID:        roomID,
		ScheduledAt:   req.ScheduledAt,
		Status:        StatusScheduled,
		Notes:         req.Notes,
	}

	if err := s.repo.Create(consultation); err != nil {
		return nil, err
	}

	return consultation, nil
}

// GetByID returns a consultation by ID.
func (s *service) GetByID(id uint) (*VideoConsultation, error) {
	consultation, err := s.repo.GetByID(id)
	if err != nil {
		return nil, ErrConsultationNotFound
	}

	return consultation, nil
}

// GetByAppointmentID returns a consultation by appointment.
func (s *service) GetByAppointmentID(appointmentID uint) (*VideoConsultation, error) {
	consultation, err := s.repo.GetByAppointmentID(appointmentID)
	if err != nil {
		return nil, ErrConsultationNotFound
	}

	return consultation, nil
}

// GetByDoctorID returns all consultations for a doctor.
func (s *service) GetByDoctorID(doctorID uint) ([]VideoConsultation, error) {
	return s.repo.GetByDoctorID(doctorID)
}

// GetByPatientID returns all consultations for a patient.
func (s *service) GetByPatientID(patientID uint) ([]VideoConsultation, error) {
	return s.repo.GetByPatientID(patientID)
}

// GetAll returns all consultations.
func (s *service) GetAll() ([]VideoConsultation, error) {
	return s.repo.GetAll()
}

// Update updates a consultation.
func (s *service) Update(id uint, req UpdateVideoConsultationRequest) (*VideoConsultation, error) {
	consultation, err := s.repo.GetByID(id)
	if err != nil {
		return nil, ErrConsultationNotFound
	}

	if req.ScheduledAt != nil {
		consultation.ScheduledAt = *req.ScheduledAt
	}

	if req.Status != nil {
		consultation.Status = *req.Status

		now := time.Now()

		switch *req.Status {
		case StatusOngoing:
			if consultation.StartedAt == nil {
				consultation.StartedAt = &now
			}

		case StatusCompleted, StatusCancelled:
			if consultation.EndedAt == nil {
				consultation.EndedAt = &now
			}
		}
	}

	if req.Notes != nil {
		consultation.Notes = *req.Notes
	}

	if err := s.repo.Update(consultation); err != nil {
		return nil, err
	}

	return consultation, nil
}

// Delete removes a consultation.
func (s *service) Delete(id uint) error {
	if _, err := s.repo.GetByID(id); err != nil {
		return ErrConsultationNotFound
	}

	return s.repo.Delete(id)
}
