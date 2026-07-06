package diagnoses

import "gorm.io/gorm"

// Repository defines the diagnosis repository contract.
type Repository interface {
	Create(diagnosis *Diagnosis) error
	GetByID(id uint) (*Diagnosis, error)
	GetByAppointmentID(appointmentID uint) ([]Diagnosis, error)
	GetByDoctorID(doctorID uint) ([]Diagnosis, error)
	GetByPatientID(patientID uint) ([]Diagnosis, error)
	GetByStatus(status string) ([]Diagnosis, error)
	GetAll() ([]Diagnosis, error)
	Update(diagnosis *Diagnosis) error
	Delete(id uint) error
}

type repository struct {
	db *gorm.DB
}

// NewRepository creates a new diagnosis repository.
func NewRepository(db *gorm.DB) Repository {
	return &repository{
		db: db,
	}
}

// Create stores a diagnosis.
func (r *repository) Create(diagnosis *Diagnosis) error {
	return r.db.Create(diagnosis).Error
}

// GetByID retrieves a diagnosis by ID.
func (r *repository) GetByID(id uint) (*Diagnosis, error) {
	var diagnosis Diagnosis

	if err := r.db.First(&diagnosis, id).Error; err != nil {
		return nil, err
	}

	return &diagnosis, nil
}

// GetByAppointmentID retrieves diagnoses for an appointment.
func (r *repository) GetByAppointmentID(appointmentID uint) ([]Diagnosis, error) {
	var diagnoses []Diagnosis

	if err := r.db.
		Where("appointment_id = ?", appointmentID).
		Order("created_at DESC").
		Find(&diagnoses).Error; err != nil {
		return nil, err
	}

	return diagnoses, nil
}

// GetByDoctorID retrieves diagnoses created by a doctor.
func (r *repository) GetByDoctorID(doctorID uint) ([]Diagnosis, error) {
	var diagnoses []Diagnosis

	if err := r.db.
		Where("doctor_id = ?", doctorID).
		Order("created_at DESC").
		Find(&diagnoses).Error; err != nil {
		return nil, err
	}

	return diagnoses, nil
}

// GetByPatientID retrieves diagnoses for a patient.
func (r *repository) GetByPatientID(patientID uint) ([]Diagnosis, error) {
	var diagnoses []Diagnosis

	if err := r.db.
		Where("patient_id = ?", patientID).
		Order("created_at DESC").
		Find(&diagnoses).Error; err != nil {
		return nil, err
	}

	return diagnoses, nil
}

// GetByStatus retrieves diagnoses by status.
func (r *repository) GetByStatus(status string) ([]Diagnosis, error) {
	var diagnoses []Diagnosis

	if err := r.db.
		Where("status = ?", status).
		Order("created_at DESC").
		Find(&diagnoses).Error; err != nil {
		return nil, err
	}

	return diagnoses, nil
}

// GetAll retrieves all diagnoses.
func (r *repository) GetAll() ([]Diagnosis, error) {
	var diagnoses []Diagnosis

	if err := r.db.
		Order("created_at DESC").
		Find(&diagnoses).Error; err != nil {
		return nil, err
	}

	return diagnoses, nil
}

// Update updates a diagnosis.
func (r *repository) Update(diagnosis *Diagnosis) error {
	return r.db.Save(diagnosis).Error
}

// Delete removes a diagnosis.
func (r *repository) Delete(id uint) error {
	return r.db.Delete(&Diagnosis{}, id).Error
}
