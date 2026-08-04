package vitals

import "gorm.io/gorm"

// Repository defines the vitals repository contract.
type Repository interface {
	Create(vital *Vital) error
	GetByID(id uint) (*Vital, error)
	GetAll() ([]Vital, error)
	GetByPatientID(patientID uint) ([]Vital, error)
	GetByDoctorID(doctorID uint) ([]Vital, error)
	GetByAppointmentID(appointmentID uint) ([]Vital, error)
	Update(vital *Vital) error
	Delete(id uint) error
}

type repository struct {
	db *gorm.DB
}

// NewRepository creates a new vitals repository.
func NewRepository(db *gorm.DB) Repository {
	return &repository{db: db}
}

// Create creates a new vital record.
func (r *repository) Create(vital *Vital) error {
	return r.db.Create(vital).Error
}

// GetByID retrieves a vital by its ID.
func (r *repository) GetByID(id uint) (*Vital, error) {
	var vital Vital
	if err := r.db.First(&vital, id).Error; err != nil {
		return nil, err
	}
	return &vital, nil
}

// GetAll retrieves all vital records.
func (r *repository) GetAll() ([]Vital, error) {
	var vitals []Vital
	if err := r.db.Find(&vitals).Error; err != nil {
		return nil, err
	}
	return vitals, nil
}

// GetByPatientID retrieves all vitals for a patient.
func (r *repository) GetByPatientID(patientID uint) ([]Vital, error) {
	var vitals []Vital
	if err := r.db.Where("patient_id = ?", patientID).Find(&vitals).Error; err != nil {
		return nil, err
	}
	return vitals, nil
}

// GetByDoctorID retrieves all vitals recorded by a doctor.
func (r *repository) GetByDoctorID(doctorID uint) ([]Vital, error) {
	var vitals []Vital
	if err := r.db.Where("doctor_id = ?", doctorID).Find(&vitals).Error; err != nil {
		return nil, err
	}
	return vitals, nil
}

// GetByAppointmentID retrieves all vitals for an appointment.
func (r *repository) GetByAppointmentID(appointmentID uint) ([]Vital, error) {
	var vitals []Vital
	if err := r.db.Where("appointment_id = ?", appointmentID).Find(&vitals).Error; err != nil {
		return nil, err
	}
	return vitals, nil
}

// Update updates an existing vital record.
func (r *repository) Update(vital *Vital) error {
	return r.db.Save(vital).Error
}

// Delete deletes a vital record.
func (r *repository) Delete(id uint) error {
	return r.db.Delete(&Vital{}, id).Error
}
