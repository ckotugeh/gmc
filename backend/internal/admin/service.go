package admin

import (
	"errors"
	"strings"
)

var (
	ErrAdminNotFound = errors.New("admin action not found")
	ErrInvalidAdmin  = errors.New("invalid admin action")
)

// Service defines the admin business logic.
type Service interface {
	Create(req CreateAdminRequest, adminID uint) (*Admin, error)
	GetByID(id uint) (*Admin, error)
	GetAll() ([]Admin, error)
	GetByAdminID(adminID uint) ([]Admin, error)
	Update(id uint, req UpdateAdminRequest) (*Admin, error)
	Delete(id uint) error
	GetDashboardStats() (*DashboardStatsResponse, error)
}

type service struct {
	repo Repository
}

// NewService creates a new admin service.
func NewService(repo Repository) Service {
	return &service{
		repo: repo,
	}
}

// Create creates a new admin action.
func (s *service) Create(req CreateAdminRequest, adminID uint) (*Admin, error) {
	if adminID == 0 ||
		req.ResourceID == 0 ||
		strings.TrimSpace(string(req.Resource)) == "" ||
		strings.TrimSpace(string(req.Action)) == "" {
		return nil, ErrInvalidAdmin
	}

	admin := &Admin{
		AdminID:     adminID,
		ResourceID:  req.ResourceID,
		Resource:    req.Resource,
		Action:      req.Action,
		Description: strings.TrimSpace(req.Description),
	}

	if err := s.repo.Create(admin); err != nil {
		return nil, err
	}

	return admin, nil
}

// GetByID retrieves an admin action by ID.
func (s *service) GetByID(id uint) (*Admin, error) {
	admin, err := s.repo.GetByID(id)
	if err != nil {
		return nil, ErrAdminNotFound
	}

	return admin, nil
}

// GetAll retrieves all admin actions.
func (s *service) GetAll() ([]Admin, error) {
	return s.repo.GetAll()
}

// GetByAdminID retrieves actions performed by an administrator.
func (s *service) GetByAdminID(adminID uint) ([]Admin, error) {
	return s.repo.GetByAdminID(adminID)
}

// Update updates an existing admin action.
func (s *service) Update(id uint, req UpdateAdminRequest) (*Admin, error) {
	admin, err := s.repo.GetByID(id)
	if err != nil {
		return nil, ErrAdminNotFound
	}

	if strings.TrimSpace(string(req.Action)) != "" {
		admin.Action = req.Action
	}

	if strings.TrimSpace(req.Description) != "" {
		admin.Description = strings.TrimSpace(req.Description)
	}

	if err := s.repo.Update(admin); err != nil {
		return nil, err
	}

	return admin, nil
}

// Delete removes an admin action.
func (s *service) Delete(id uint) error {
	if _, err := s.repo.GetByID(id); err != nil {
		return ErrAdminNotFound
	}

	return s.repo.Delete(id)
}

// GetDashboardStats returns dashboard statistics.
func (s *service) GetDashboardStats() (*DashboardStatsResponse, error) {
	return s.repo.GetDashboardStats()
}
