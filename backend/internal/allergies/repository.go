package allergies

import "gorm.io/gorm"

// Repository defines the allergy repository contract.
type Repository interface {
	Create(allergy *Allergy) error
	GetByID(id uint) (*Allergy, error)
	GetAll() ([]Allergy, error)
	GetByPatientID(patientID uint) ([]Allergy, error)
	GetByDoctorID(doctorID uint) ([]Allergy, error)
	GetBySeverity(severity string) ([]Allergy, error)
	GetActive() ([]Allergy, error)
	Update(allergy *Allergy) error
	Delete(id uint) error
}

type repository struct {
	db *gorm.DB
}

// NewRepository creates a new allergy repository.
func NewRepository(db *gorm.DB) Repository {
	return &repository{
		db: db,
	}
}

// Create creates a new allergy record.
func (r *repository) Create(allergy *Allergy) error {
	return r.db.Create(allergy).Error
}

// GetByID retrieves an allergy by ID.
func (r *repository) GetByID(id uint) (*Allergy, error) {
	var allergy Allergy

	if err := r.db.First(&allergy, id).Error; err != nil {
		return nil, err
	}

	return &allergy, nil
}

// GetAll retrieves all allergy records.
func (r *repository) GetAll() ([]Allergy, error) {
	var allergies []Allergy

	if err := r.db.Find(&allergies).Error; err != nil {
		return nil, err
	}

	return allergies, nil
}

// GetByPatientID retrieves allergies for a patient.
func (r *repository) GetByPatientID(patientID uint) ([]Allergy, error) {
	var allergies []Allergy

	if err := r.db.Where("patient_id = ?", patientID).
		Find(&allergies).Error; err != nil {
		return nil, err
	}

	return allergies, nil
}

// GetByDoctorID retrieves allergies recorded by a doctor.
func (r *repository) GetByDoctorID(doctorID uint) ([]Allergy, error) {
	var allergies []Allergy

	if err := r.db.Where("doctor_id = ?", doctorID).
		Find(&allergies).Error; err != nil {
		return nil, err
	}

	return allergies, nil
}

// GetBySeverity retrieves allergies by severity.
func (r *repository) GetBySeverity(severity string) ([]Allergy, error) {
	var allergies []Allergy

	if err := r.db.Where("severity = ?", severity).
		Find(&allergies).Error; err != nil {
		return nil, err
	}

	return allergies, nil
}

// GetActive retrieves all active allergies.
func (r *repository) GetActive() ([]Allergy, error) {
	var allergies []Allergy

	if err := r.db.Where("status = ?", "Active").
		Find(&allergies).Error; err != nil {
		return nil, err
	}

	return allergies, nil
}

// Update updates an allergy record.
func (r *repository) Update(allergy *Allergy) error {
	return r.db.Save(allergy).Error
}

// Delete deletes an allergy record.
func (r *repository) Delete(id uint) error {
	return r.db.Delete(&Allergy{}, id).Error
}
