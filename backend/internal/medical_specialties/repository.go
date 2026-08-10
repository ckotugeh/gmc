package medical_specialties

import "gorm.io/gorm"

// Repository defines the medical specialty repository contract.
type Repository interface {
	Create(specialty *MedicalSpecialty) error
	GetByID(id uint) (*MedicalSpecialty, error)
	GetByName(name string) (*MedicalSpecialty, error)
	GetByCode(code string) (*MedicalSpecialty, error)
	GetAll() ([]MedicalSpecialty, error)
	GetActive() ([]MedicalSpecialty, error)
	Update(specialty *MedicalSpecialty) error
	Delete(id uint) error
}

type repository struct {
	db *gorm.DB
}

// NewRepository creates a new medical specialty repository.
func NewRepository(db *gorm.DB) Repository {
	return &repository{
		db: db,
	}
}

// Create stores a medical specialty.
func (r *repository) Create(specialty *MedicalSpecialty) error {
	return r.db.Create(specialty).Error
}

// GetByID retrieves a medical specialty by ID.
func (r *repository) GetByID(id uint) (*MedicalSpecialty, error) {
	var specialty MedicalSpecialty

	if err := r.db.First(&specialty, id).Error; err != nil {
		return nil, err
	}

	return &specialty, nil
}

// GetByName retrieves a medical specialty by name.
func (r *repository) GetByName(name string) (*MedicalSpecialty, error) {
	var specialty MedicalSpecialty

	if err := r.db.Where("name = ?", name).First(&specialty).Error; err != nil {
		return nil, err
	}

	return &specialty, nil
}

// GetByCode retrieves a medical specialty by code.
func (r *repository) GetByCode(code string) (*MedicalSpecialty, error) {
	var specialty MedicalSpecialty

	if err := r.db.Where("code = ?", code).First(&specialty).Error; err != nil {
		return nil, err
	}

	return &specialty, nil
}

// GetAll retrieves all medical specialties.
func (r *repository) GetAll() ([]MedicalSpecialty, error) {
	var specialties []MedicalSpecialty

	if err := r.db.
		Order("name ASC").
		Find(&specialties).Error; err != nil {
		return nil, err
	}

	return specialties, nil
}

// GetActive retrieves all active medical specialties.
func (r *repository) GetActive() ([]MedicalSpecialty, error) {
	var specialties []MedicalSpecialty

	if err := r.db.
		Where("is_active = ?", true).
		Order("name ASC").
		Find(&specialties).Error; err != nil {
		return nil, err
	}

	return specialties, nil
}

// Update updates a medical specialty.
func (r *repository) Update(specialty *MedicalSpecialty) error {
	return r.db.Save(specialty).Error
}

// Delete removes a medical specialty.
func (r *repository) Delete(id uint) error {
	return r.db.Delete(&MedicalSpecialty{}, id).Error
}
