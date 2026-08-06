package doctor_reviews

import (
	"errors"
	"strings"
)

var (
	ErrDoctorReviewNotFound       = errors.New("doctor review not found")
	ErrInvalidDoctorReview        = errors.New("invalid doctor review")
	ErrDuplicateAppointmentReview = errors.New("review already exists for this appointment")
)

// Service defines the doctor review business logic.
type Service interface {
	Create(req CreateDoctorReviewRequest) (*DoctorReview, error)
	GetByID(id uint) (*DoctorReview, error)
	GetByAppointmentID(appointmentID uint) (*DoctorReview, error)
	GetByDoctorID(doctorID uint) ([]DoctorReview, error)
	GetByPatientID(patientID uint) ([]DoctorReview, error)
	GetPublishedByDoctorID(doctorID uint) ([]DoctorReview, error)
	GetAll() ([]DoctorReview, error)
	Update(id uint, req UpdateDoctorReviewRequest) (*DoctorReview, error)
	Delete(id uint) error
}

type service struct {
	repo Repository
}

// NewService creates a new doctor review service.
func NewService(repo Repository) Service {
	return &service{
		repo: repo,
	}
}

// Create creates a new doctor review.
func (s *service) Create(req CreateDoctorReviewRequest) (*DoctorReview, error) {
	if req.DoctorID == 0 ||
		req.PatientID == 0 ||
		req.AppointmentID == 0 ||
		req.Rating < 1 ||
		req.Rating > 5 ||
		strings.TrimSpace(req.Comment) == "" {
		return nil, ErrInvalidDoctorReview
	}

	if _, err := s.repo.GetByAppointmentID(req.AppointmentID); err == nil {
		return nil, ErrDuplicateAppointmentReview
	}

	isAnonymous := false
	if req.IsAnonymous != nil {
		isAnonymous = *req.IsAnonymous
	}

	isPublished := true
	if req.IsPublished != nil {
		isPublished = *req.IsPublished
	}

	review := &DoctorReview{
		DoctorID:      req.DoctorID,
		PatientID:     req.PatientID,
		AppointmentID: req.AppointmentID,
		Rating:        req.Rating,
		Title:         strings.TrimSpace(req.Title),
		Comment:       strings.TrimSpace(req.Comment),
		IsAnonymous:   isAnonymous,
		IsPublished:   isPublished,
	}

	if err := s.repo.Create(review); err != nil {
		return nil, err
	}

	return review, nil
}

// GetByID retrieves a review by ID.
func (s *service) GetByID(id uint) (*DoctorReview, error) {
	review, err := s.repo.GetByID(id)
	if err != nil {
		return nil, ErrDoctorReviewNotFound
	}

	return review, nil
}

// GetByAppointmentID retrieves a review by appointment ID.
func (s *service) GetByAppointmentID(appointmentID uint) (*DoctorReview, error) {
	review, err := s.repo.GetByAppointmentID(appointmentID)
	if err != nil {
		return nil, ErrDoctorReviewNotFound
	}

	return review, nil
}

// GetByDoctorID retrieves all reviews for a doctor.
func (s *service) GetByDoctorID(doctorID uint) ([]DoctorReview, error) {
	return s.repo.GetByDoctorID(doctorID)
}

// GetByPatientID retrieves all reviews written by a patient.
func (s *service) GetByPatientID(patientID uint) ([]DoctorReview, error) {
	return s.repo.GetByPatientID(patientID)
}

// GetPublishedByDoctorID retrieves all published reviews for a doctor.
func (s *service) GetPublishedByDoctorID(doctorID uint) ([]DoctorReview, error) {
	return s.repo.GetPublishedByDoctorID(doctorID)
}

// GetAll retrieves all reviews.
func (s *service) GetAll() ([]DoctorReview, error) {
	return s.repo.GetAll()
}

// Update updates an existing review.
func (s *service) Update(id uint, req UpdateDoctorReviewRequest) (*DoctorReview, error) {
	review, err := s.repo.GetByID(id)
	if err != nil {
		return nil, ErrDoctorReviewNotFound
	}

	if req.Rating != 0 {
		if req.Rating < 1 || req.Rating > 5 {
			return nil, ErrInvalidDoctorReview
		}
		review.Rating = req.Rating
	}

	if req.Title != "" {
		review.Title = strings.TrimSpace(req.Title)
	}

	if req.Comment != "" {
		review.Comment = strings.TrimSpace(req.Comment)
	}

	if req.IsAnonymous != nil {
		review.IsAnonymous = *req.IsAnonymous
	}

	if req.IsPublished != nil {
		review.IsPublished = *req.IsPublished
	}

	if err := s.repo.Update(review); err != nil {
		return nil, err
	}

	return review, nil
}

// Delete removes a doctor review.
func (s *service) Delete(id uint) error {
	if _, err := s.repo.GetByID(id); err != nil {
		return ErrDoctorReviewNotFound
	}

	return s.repo.Delete(id)
}
