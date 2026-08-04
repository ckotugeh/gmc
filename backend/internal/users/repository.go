package users

import (
	"errors"

	"gorm.io/gorm"
)

// Repository defines user persistence methods.
type Repository interface {
	Create(user *User) error
	GetByID(id uint) (*User, error)
	GetByEmail(email string) (*User, error)
	GetAll() ([]User, error)
	GetDoctors() ([]User, error)
	GetPatients() ([]User, error)
	Update(user *User) error
	Delete(id uint) error
}

// repository implements Repository.
type repository struct {
	db *gorm.DB
}

// NewRepository creates a new user repository.
func NewRepository(db *gorm.DB) Repository {
	return &repository{
		db: db,
	}
}

// Create creates a new user.
func (r *repository) Create(user *User) error {
	return r.db.Create(user).Error
}

// GetByID retrieves a user by ID.
func (r *repository) GetByID(id uint) (*User, error) {
	var user User

	if err := r.db.First(&user, id).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, err
		}
		return nil, err
	}

	return &user, nil
}

// GetByEmail retrieves a user by email.
func (r *repository) GetByEmail(email string) (*User, error) {
	var user User

	if err := r.db.Where("email = ?", email).First(&user).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, err
		}
		return nil, err
	}

	return &user, nil
}

// GetAll returns all users.
func (r *repository) GetAll() ([]User, error) {
	var users []User

	if err := r.db.Find(&users).Error; err != nil {
		return nil, err
	}

	return users, nil
}

// GetDoctors returns all doctors.
func (r *repository) GetDoctors() ([]User, error) {
	var users []User

	if err := r.db.Where("role = ?", "doctor").Find(&users).Error; err != nil {
		return nil, err
	}

	return users, nil
}

// GetPatients returns all patients.
func (r *repository) GetPatients() ([]User, error) {
	var users []User

	if err := r.db.Where("role = ?", "patient").Find(&users).Error; err != nil {
		return nil, err
	}

	return users, nil
}

// Update updates a user.
func (r *repository) Update(user *User) error {
	return r.db.Save(user).Error
}

// Delete deletes a user.
func (r *repository) Delete(id uint) error {
	return r.db.Delete(&User{}, id).Error
}
