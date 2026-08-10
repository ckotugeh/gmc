package prescriptions

import "gorm.io/gorm"

// Repository defines the prescription repository contract.
type Repository interface {
	Create(prescription *Prescription) error
	GetByID(id uint) (*Prescription, error)
	GetAll() ([]Prescription, error)

	GetByDoctorID(doctorID uint) ([]Prescription, error)
	GetByPatientID(patientID uint) ([]Prescription, error)
	GetByAppointmentID(appointmentID uint) ([]Prescription, error)

	Update(prescription *Prescription) error
	Delete(id uint) error
}

type repository struct {
	db *gorm.DB
}

// NewRepository creates a new prescription repository.
func NewRepository(db *gorm.DB) Repository {
	return &repository{
		db: db,
	}
}

// Create stores a prescription.
func (r *repository) Create(prescription *Prescription) error {
	return r.db.Create(prescription).Error
}

// GetByID retrieves a prescription by ID.
func (r *repository) GetByID(id uint) (*Prescription, error) {
	var prescription Prescription

	if err := r.db.
		Preload("Items").
		First(&prescription, id).Error; err != nil {
		return nil, err
	}

	return &prescription, nil
}

// GetAll retrieves all prescriptions.
func (r *repository) GetAll() ([]Prescription, error) {
	var prescriptions []Prescription

	if err := r.db.
		Preload("Items").
		Order("created_at DESC").
		Find(&prescriptions).Error; err != nil {
		return nil, err
	}

	return prescriptions, nil
}

// GetByDoctorID retrieves prescriptions by doctor.
func (r *repository) GetByDoctorID(doctorID uint) ([]Prescription, error) {
	var prescriptions []Prescription

	if err := r.db.
		Preload("Items").
		Where("doctor_id = ?", doctorID).
		Order("created_at DESC").
		Find(&prescriptions).Error; err != nil {
		return nil, err
	}

	return prescriptions, nil
}

// GetByPatientID retrieves prescriptions by patient.
func (r *repository) GetByPatientID(patientID uint) ([]Prescription, error) {
	var prescriptions []Prescription

	if err := r.db.
		Preload("Items").
		Where("patient_id = ?", patientID).
		Order("created_at DESC").
		Find(&prescriptions).Error; err != nil {
		return nil, err
	}

	return prescriptions, nil
}

// GetByAppointmentID retrieves prescriptions by appointment.
func (r *repository) GetByAppointmentID(appointmentID uint) ([]Prescription, error) {
	var prescriptions []Prescription

	if err := r.db.
		Preload("Items").
		Where("appointment_id = ?", appointmentID).
		Order("created_at DESC").
		Find(&prescriptions).Error; err != nil {
		return nil, err
	}

	return prescriptions, nil
}

// Update updates a prescription.
func (r *repository) Update(prescription *Prescription) error {
	return r.db.Session(&gorm.Session{FullSaveAssociations: true}).
		Updates(prescription).Error
}

// Delete removes a prescription.
func (r *repository) Delete(id uint) error {
	return r.db.Delete(&Prescription{}, id).Error
}
