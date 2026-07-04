package uploads

import (
	"errors"
	"strings"
)

var (
	ErrUploadNotFound = errors.New("upload not found")
	ErrInvalidUpload  = errors.New("invalid upload")
)

// Service defines upload business logic.
type Service interface {
	Create(upload *Upload) (*Upload, error)
	GetByID(id uint) (*Upload, error)
	GetByUserID(userID uint) ([]Upload, error)
	GetByAppointmentID(appointmentID uint) ([]Upload, error)
	GetByMedicalRecordID(recordID uint) ([]Upload, error)
	GetByHospitalID(hospitalID uint) ([]Upload, error)
	Update(id uint, req UpdateUploadRequest) (*Upload, error)
	Delete(id uint) error
}

type service struct {
	repo Repository
}

// NewService creates a new upload service.
func NewService(repo Repository) Service {
	return &service{
		repo: repo,
	}
}

// Create creates a new upload record.
func (s *service) Create(upload *Upload) (*Upload, error) {
	if upload == nil {
		return nil, ErrInvalidUpload
	}

	if upload.UserID == 0 ||
		strings.TrimSpace(upload.FileName) == "" ||
		strings.TrimSpace(upload.OriginalName) == "" ||
		strings.TrimSpace(upload.FileType) == "" ||
		strings.TrimSpace(upload.FilePath) == "" ||
		upload.FileSize <= 0 {
		return nil, ErrInvalidUpload
	}

	if err := s.repo.Create(upload); err != nil {
		return nil, err
	}

	return upload, nil
}

// GetByID retrieves an upload by ID.
func (s *service) GetByID(id uint) (*Upload, error) {
	upload, err := s.repo.GetByID(id)
	if err != nil {
		return nil, ErrUploadNotFound
	}
	return upload, nil
}

// GetByUserID retrieves uploads belonging to a user.
func (s *service) GetByUserID(userID uint) ([]Upload, error) {
	return s.repo.GetByUserID(userID)
}

// GetByAppointmentID retrieves uploads for an appointment.
func (s *service) GetByAppointmentID(appointmentID uint) ([]Upload, error) {
	return s.repo.GetByAppointmentID(appointmentID)
}

// GetByMedicalRecordID retrieves uploads for a medical record.
func (s *service) GetByMedicalRecordID(recordID uint) ([]Upload, error) {
	return s.repo.GetByMedicalRecordID(recordID)
}

// GetByHospitalID retrieves uploads for a hospital.
func (s *service) GetByHospitalID(hospitalID uint) ([]Upload, error) {
	return s.repo.GetByHospitalID(hospitalID)
}

// Update updates upload metadata.
func (s *service) Update(id uint, req UpdateUploadRequest) (*Upload, error) {
	upload, err := s.repo.GetByID(id)
	if err != nil {
		return nil, ErrUploadNotFound
	}

	if strings.TrimSpace(req.Description) != "" {
		upload.Description = req.Description
	}

	if req.IsPublic != nil {
		upload.IsPublic = *req.IsPublic
	}

	if err := s.repo.Update(upload); err != nil {
		return nil, err
	}

	return upload, nil
}

// Delete removes an upload.
func (s *service) Delete(id uint) error {
	if _, err := s.repo.GetByID(id); err != nil {
		return ErrUploadNotFound
	}

	return s.repo.Delete(id)
}
