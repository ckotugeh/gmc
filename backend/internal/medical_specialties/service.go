package medical_specialties

import (
	"errors"
	"strings"
)

var (
	ErrMedicalSpecialtyNotFound = errors.New("medical specialty not found")
	ErrInvalidMedicalSpecialty  = errors.New("invalid medical specialty")
	ErrDuplicateSpecialtyName   = errors.New("medical specialty name already exists")
	ErrDuplicateSpecialtyCode   = errors.New("medical specialty code already exists")
)

// Service defines the medical specialty business logic.
type Service interface {
	Create(req CreateMedicalSpecialtyRequest) (*MedicalSpecialty, error)
	GetByID(id uint) (*MedicalSpecialty, error)
	GetByName(name string) (*MedicalSpecialty, error)
	GetByCode(code string) (*MedicalSpecialty, error)
	GetAll() ([]MedicalSpecialty, error)
	GetActive() ([]MedicalSpecialty, error)
	Update(id uint, req UpdateMedicalSpecialtyRequest) (*MedicalSpecialty, error)
	Delete(id uint) error
}

type service struct {
	repo Repository
}

// NewService creates a new medical specialty service.
func NewService(repo Repository) Service {
	return &service{
		repo: repo,
	}
}

// Create creates a new medical specialty.
func (s *service) Create(req CreateMedicalSpecialtyRequest) (*MedicalSpecialty, error) {
	name := strings.TrimSpace(req.Name)
	code := strings.ToUpper(strings.TrimSpace(req.Code))

	if name == "" || code == "" {
		return nil, ErrInvalidMedicalSpecialty
	}

	if _, err := s.repo.GetByName(name); err == nil {
		return nil, ErrDuplicateSpecialtyName
	}

	if _, err := s.repo.GetByCode(code); err == nil {
		return nil, ErrDuplicateSpecialtyCode
	}

	active := true
	if req.IsActive != nil {
		active = *req.IsActive
	}

	specialty := &MedicalSpecialty{
		Name:        name,
		Code:        code,
		Description: strings.TrimSpace(req.Description),
		IsActive:    active,
	}

	if err := s.repo.Create(specialty); err != nil {
		return nil, err
	}

	return specialty, nil
}

// GetByID retrieves a specialty by ID.
func (s *service) GetByID(id uint) (*MedicalSpecialty, error) {
	specialty, err := s.repo.GetByID(id)
	if err != nil {
		return nil, ErrMedicalSpecialtyNotFound
	}

	return specialty, nil
}

// GetByName retrieves a specialty by name.
func (s *service) GetByName(name string) (*MedicalSpecialty, error) {
	name = strings.TrimSpace(name)
	if name == "" {
		return nil, ErrInvalidMedicalSpecialty
	}

	specialty, err := s.repo.GetByName(name)
	if err != nil {
		return nil, ErrMedicalSpecialtyNotFound
	}

	return specialty, nil
}

// GetByCode retrieves a specialty by code.
func (s *service) GetByCode(code string) (*MedicalSpecialty, error) {
	code = strings.ToUpper(strings.TrimSpace(code))
	if code == "" {
		return nil, ErrInvalidMedicalSpecialty
	}

	specialty, err := s.repo.GetByCode(code)
	if err != nil {
		return nil, ErrMedicalSpecialtyNotFound
	}

	return specialty, nil
}

// GetAll retrieves all specialties.
func (s *service) GetAll() ([]MedicalSpecialty, error) {
	return s.repo.GetAll()
}

// GetActive retrieves all active specialties.
func (s *service) GetActive() ([]MedicalSpecialty, error) {
	return s.repo.GetActive()
}

// Update updates a medical specialty.
func (s *service) Update(id uint, req UpdateMedicalSpecialtyRequest) (*MedicalSpecialty, error) {
	specialty, err := s.repo.GetByID(id)
	if err != nil {
		return nil, ErrMedicalSpecialtyNotFound
	}

	if req.Name != "" {
		specialty.Name = strings.TrimSpace(req.Name)
	}

	if req.Code != "" {
		specialty.Code = strings.ToUpper(strings.TrimSpace(req.Code))
	}

	if req.Description != "" {
		specialty.Description = strings.TrimSpace(req.Description)
	}

	if req.IsActive != nil {
		specialty.IsActive = *req.IsActive
	}

	if err := s.repo.Update(specialty); err != nil {
		return nil, err
	}

	return specialty, nil
}

// Delete removes a medical specialty.
func (s *service) Delete(id uint) error {
	if _, err := s.repo.GetByID(id); err != nil {
		return ErrMedicalSpecialtyNotFound
	}

	return s.repo.Delete(id)
}
