package appointments

import "gorm.io/gorm"

// Repository defines appointment persistence operations.
type Repository interface {
	Create(appointment *Appointment) error
	GetByID(id uint) (*Appointment, error)
	GetAll() ([]Appointment, error)
	GetByDoctorID(doctorID uint) ([]Appointment, error)
	GetByPatientID(patientID uint) ([]Appointment, error)
	Update(appointment *Appointment) error
	Delete(id uint) error
}

// repository implements Repository.
type repository struct {
	db *gorm.DB
}

// NewRepository creates a new appointment repository.
func NewRepository(db *gorm.DB) Repository {
	return &repository{
		db: db,
	}
}

// Create saves a new appointment.
func (r *repository) Create(appointment *Appointment) error {
	return r.db.Create(appointment).Error
}

// GetByID retrieves an appointment by its ID.
func (r *repository) GetByID(id uint) (*Appointment, error) {
	var appointment Appointment

	if err := r.db.First(&appointment, id).Error; err != nil {
		return nil, err
	}

	return &appointment, nil
}

// GetAll retrieves all appointments.
func (r *repository) GetAll() ([]Appointment, error) {
	var appointments []Appointment

	if err := r.db.Find(&appointments).Error; err != nil {
		return nil, err
	}

	return appointments, nil
}

// GetByDoctorID retrieves appointments for a doctor.
func (r *repository) GetByDoctorID(doctorID uint) ([]Appointment, error) {
	var appointments []Appointment

	if err := r.db.Where("doctor_id = ?", doctorID).
		Find(&appointments).Error; err != nil {
		return nil, err
	}

	return appointments, nil
}

// GetByPatientID retrieves appointments for a patient.
func (r *repository) GetByPatientID(patientID uint) ([]Appointment, error) {
	var appointments []Appointment

	if err := r.db.Where("patient_id = ?", patientID).
		Find(&appointments).Error; err != nil {
		return nil, err
	}

	return appointments, nil
}

// Update updates an existing appointment.
func (r *repository) Update(appointment *Appointment) error {
	return r.db.Save(appointment).Error
}

// Delete removes an appointment.
func (r *repository) Delete(id uint) error {
	return r.db.Delete(&Appointment{}, id).Error
}
