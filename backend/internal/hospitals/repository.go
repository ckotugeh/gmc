package hospitals

import (
	"gorm.io/gorm"
)

// Repository defines the hospital repository contract.
type Repository interface {
	Create(hospital *Hospital) error
	GetByID(id uint) (*Hospital, error)
	GetAll() ([]Hospital, error)
	GetByEmail(email string) (*Hospital, error)
	GetByLicenseNumber(license string) (*Hospital, error)
	Update(hospital *Hospital) error
	Delete(id uint) error
}

type repository struct {
	db *gorm.DB
}

// NewRepository creates a new hospital repository.
func NewRepository(db *gorm.DB) Repository {
	return &repository{db: db}
}

// Create creates a new hospital.
func (r *repository) Create(hospital *Hospital) error {
	return r.db.Create(hospital).Error
}

// GetByID retrieves a hospital by ID.
func (r *repository) GetByID(id uint) (*Hospital, error) {
	var hospital Hospital
	if err := r.db.First(&hospital, id).Error; err != nil {
		return nil, err
	}
	return &hospital, nil
}

// GetAll retrieves all hospitals.
func (r *repository) GetAll() ([]Hospital, error) {
	var hospitals []Hospital
	if err := r.db.Find(&hospitals).Error; err != nil {
		return nil, err
	}
	return hospitals, nil
}

// GetByEmail retrieves a hospital by email.
func (r *repository) GetByEmail(email string) (*Hospital, error) {
	var hospital Hospital
	if err := r.db.Where("email = ?", email).First(&hospital).Error; err != nil {
		return nil, err
	}
	return &hospital, nil
}

// GetByLicenseNumber retrieves a hospital by license number.
func (r *repository) GetByLicenseNumber(license string) (*Hospital, error) {
	var hospital Hospital
	if err := r.db.Where("license_number = ?", license).First(&hospital).Error; err != nil {
		return nil, err
	}
	return &hospital, nil
}

// Update updates a hospital.
func (r *repository) Update(hospital *Hospital) error {
	return r.db.Save(hospital).Error
}

// Delete deletes a hospital.
func (r *repository) Delete(id uint) error {
	return r.db.Delete(&Hospital{}, id).Error
}
