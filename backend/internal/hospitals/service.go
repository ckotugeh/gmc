package hospitals

import (
	"errors"
	"strings"
)

var (
	ErrHospitalNotFound      = errors.New("hospital not found")
	ErrInvalidHospital       = errors.New("invalid hospital")
	ErrHospitalEmailExists   = errors.New("hospital email already exists")
	ErrHospitalLicenseExists = errors.New("hospital license number already exists")
)

// Service defines the hospital business logic.
type Service interface {
	Create(req CreateHospitalRequest) (*Hospital, error)
	GetByID(id uint) (*Hospital, error)
	GetAll() ([]Hospital, error)
	Update(id uint, req UpdateHospitalRequest) (*Hospital, error)
	Delete(id uint) error
}

type service struct {
	repo Repository
}

// NewService creates a new hospital service.
func NewService(repo Repository) Service {
	return &service{repo: repo}
}

// Create creates a new hospital.
func (s *service) Create(req CreateHospitalRequest) (*Hospital, error) {
	if strings.TrimSpace(req.Name) == "" ||
		strings.TrimSpace(req.Email) == "" ||
		strings.TrimSpace(req.LicenseNumber) == "" {
		return nil, ErrInvalidHospital
	}

	if _, err := s.repo.GetByEmail(req.Email); err == nil {
		return nil, ErrHospitalEmailExists
	}

	if _, err := s.repo.GetByLicenseNumber(req.LicenseNumber); err == nil {
		return nil, ErrHospitalLicenseExists
	}

	hospital := &Hospital{
		Name:          req.Name,
		Description:   req.Description,
		Email:         req.Email,
		Phone:         req.Phone,
		Website:       req.Website,
		Address:       req.Address,
		City:          req.City,
		State:         req.State,
		Country:       req.Country,
		ZipCode:       req.ZipCode,
		LicenseNumber: req.LicenseNumber,
		IsActive:      req.IsActive,
	}

	if err := s.repo.Create(hospital); err != nil {
		return nil, err
	}

	return hospital, nil
}

// GetByID retrieves a hospital by ID.
func (s *service) GetByID(id uint) (*Hospital, error) {
	return s.repo.GetByID(id)
}

// GetAll retrieves all hospitals.
func (s *service) GetAll() ([]Hospital, error) {
	return s.repo.GetAll()
}

// Update updates a hospital.
func (s *service) Update(id uint, req UpdateHospitalRequest) (*Hospital, error) {
	hospital, err := s.repo.GetByID(id)
	if err != nil {
		return nil, ErrHospitalNotFound
	}

	if req.Name != "" {
		hospital.Name = req.Name
	}
	if req.Description != "" {
		hospital.Description = req.Description
	}
	if req.Email != "" && req.Email != hospital.Email {
		if _, err := s.repo.GetByEmail(req.Email); err == nil {
			return nil, ErrHospitalEmailExists
		}
		hospital.Email = req.Email
	}
	if req.Phone != "" {
		hospital.Phone = req.Phone
	}
	if req.Website != "" {
		hospital.Website = req.Website
	}
	if req.Address != "" {
		hospital.Address = req.Address
	}
	if req.City != "" {
		hospital.City = req.City
	}
	if req.State != "" {
		hospital.State = req.State
	}
	if req.Country != "" {
		hospital.Country = req.Country
	}
	if req.ZipCode != "" {
		hospital.ZipCode = req.ZipCode
	}
	if req.LicenseNumber != "" && req.LicenseNumber != hospital.LicenseNumber {
		if _, err := s.repo.GetByLicenseNumber(req.LicenseNumber); err == nil {
			return nil, ErrHospitalLicenseExists
		}
		hospital.LicenseNumber = req.LicenseNumber
	}
	if req.IsActive != nil {
		hospital.IsActive = *req.IsActive
	}

	if err := s.repo.Update(hospital); err != nil {
		return nil, err
	}

	return hospital, nil
}

// Delete removes a hospital.
func (s *service) Delete(id uint) error {
	if _, err := s.repo.GetByID(id); err != nil {
		return ErrHospitalNotFound
	}

	return s.repo.Delete(id)
}
