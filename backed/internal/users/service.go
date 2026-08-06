package users

import (
	"errors"
	"strings"

	"doctor-platform/internal/utils"
)

// Service defines user business logic.
type Service interface {
	CreateUser(req CreateUserRequest) (*User, error)
	GetUser(id uint) (*User, error)
	GetAllUsers() ([]User, error)
	GetDoctors() ([]User, error)
	GetPatients() ([]User, error)
	UpdateUser(id uint, req UpdateUserRequest) (*User, error)
	DeleteUser(id uint) error
}

type service struct {
	repo Repository
}

// NewService creates a new user service.
func NewService(repo Repository) Service {
	return &service{
		repo: repo,
	}
}

// CreateUser creates a new user.
func (s *service) CreateUser(req CreateUserRequest) (*User, error) {
	req.Email = strings.TrimSpace(strings.ToLower(req.Email))

	if req.FirstName == "" {
		return nil, errors.New("first name is required")
	}

	if req.LastName == "" {
		return nil, errors.New("last name is required")
	}

	if req.Email == "" {
		return nil, errors.New("email is required")
	}

	if req.Password == "" {
		return nil, errors.New("password is required")
	}

	// Check if email already exists.
	if existing, _ := s.repo.GetByEmail(req.Email); existing != nil {
		return nil, errors.New("email already exists")
	}

	hashedPassword, err := utils.HashPassword(req.Password)
	if err != nil {
		return nil, err
	}

	role := req.Role
	if role == "" {
		role = "patient"
	}

	user := &User{
		FirstName: req.FirstName,
		LastName:  req.LastName,
		Email:     req.Email,
		Password:  hashedPassword,
		Phone:     req.Phone,
		Role:      role,
		IsActive:  true,
	}

	if err := s.repo.Create(user); err != nil {
		return nil, err
	}

	return user, nil
}

// GetUser returns a user by ID.
func (s *service) GetUser(id uint) (*User, error) {
	return s.repo.GetByID(id)
}

// GetAllUsers returns all users.
func (s *service) GetAllUsers() ([]User, error) {
	return s.repo.GetAll()
}

// GetDoctors returns all doctors.
func (s *service) GetDoctors() ([]User, error) {
	return s.repo.GetDoctors()
}

// GetPatients returns all patients.
func (s *service) GetPatients() ([]User, error) {
	return s.repo.GetPatients()
}

// UpdateUser updates an existing user.
func (s *service) UpdateUser(id uint, req UpdateUserRequest) (*User, error) {
	user, err := s.repo.GetByID(id)
	if err != nil {
		return nil, err
	}

	if req.FirstName != nil {
		user.FirstName = *req.FirstName
	}

	if req.LastName != nil {
		user.LastName = *req.LastName
	}

	if req.Phone != nil {
		user.Phone = *req.Phone
	}

	if req.Role != nil {
		user.Role = *req.Role
	}

	if req.IsActive != nil {
		user.IsActive = *req.IsActive
	}

	if err := s.repo.Update(user); err != nil {
		return nil, err
	}

	return user, nil
}

// DeleteUser deletes a user.
func (s *service) DeleteUser(id uint) error {
	_, err := s.repo.GetByID(id)
	if err != nil {
		return err
	}

	return s.repo.Delete(id)
}
