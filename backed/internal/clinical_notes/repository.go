package clinical_notes

import "gorm.io/gorm"

// Repository defines the clinical note repository contract.
type Repository interface {
	Create(note *ClinicalNote) error
	GetByID(id uint) (*ClinicalNote, error)
	GetByAppointmentID(appointmentID uint) ([]ClinicalNote, error)
	GetByDoctorID(doctorID uint) ([]ClinicalNote, error)
	GetByPatientID(patientID uint) ([]ClinicalNote, error)
	GetByDiagnosisID(diagnosisID uint) ([]ClinicalNote, error)
	GetConfidential() ([]ClinicalNote, error)
	GetAll() ([]ClinicalNote, error)
	Update(note *ClinicalNote) error
	Delete(id uint) error
}

type repository struct {
	db *gorm.DB
}

// NewRepository creates a new clinical note repository.
func NewRepository(db *gorm.DB) Repository {
	return &repository{
		db: db,
	}
}

// Create stores a clinical note.
func (r *repository) Create(note *ClinicalNote) error {
	return r.db.Create(note).Error
}

// GetByID retrieves a clinical note by ID.
func (r *repository) GetByID(id uint) (*ClinicalNote, error) {
	var note ClinicalNote

	if err := r.db.First(&note, id).Error; err != nil {
		return nil, err
	}

	return &note, nil
}

// GetByAppointmentID retrieves clinical notes for an appointment.
func (r *repository) GetByAppointmentID(appointmentID uint) ([]ClinicalNote, error) {
	var notes []ClinicalNote

	if err := r.db.
		Where("appointment_id = ?", appointmentID).
		Order("created_at DESC").
		Find(&notes).Error; err != nil {
		return nil, err
	}

	return notes, nil
}

// GetByDoctorID retrieves clinical notes created by a doctor.
func (r *repository) GetByDoctorID(doctorID uint) ([]ClinicalNote, error) {
	var notes []ClinicalNote

	if err := r.db.
		Where("doctor_id = ?", doctorID).
		Order("created_at DESC").
		Find(&notes).Error; err != nil {
		return nil, err
	}

	return notes, nil
}

// GetByPatientID retrieves clinical notes for a patient.
func (r *repository) GetByPatientID(patientID uint) ([]ClinicalNote, error) {
	var notes []ClinicalNote

	if err := r.db.
		Where("patient_id = ?", patientID).
		Order("created_at DESC").
		Find(&notes).Error; err != nil {
		return nil, err
	}

	return notes, nil
}

// GetByDiagnosisID retrieves clinical notes for a diagnosis.
func (r *repository) GetByDiagnosisID(diagnosisID uint) ([]ClinicalNote, error) {
	var notes []ClinicalNote

	if err := r.db.
		Where("diagnosis_id = ?", diagnosisID).
		Order("created_at DESC").
		Find(&notes).Error; err != nil {
		return nil, err
	}

	return notes, nil
}

// GetConfidential retrieves all confidential clinical notes.
func (r *repository) GetConfidential() ([]ClinicalNote, error) {
	var notes []ClinicalNote

	if err := r.db.
		Where("is_confidential = ?", true).
		Order("created_at DESC").
		Find(&notes).Error; err != nil {
		return nil, err
	}

	return notes, nil
}

// GetAll retrieves all clinical notes.
func (r *repository) GetAll() ([]ClinicalNote, error) {
	var notes []ClinicalNote

	if err := r.db.
		Order("created_at DESC").
		Find(&notes).Error; err != nil {
		return nil, err
	}

	return notes, nil
}

// Update updates a clinical note.
func (r *repository) Update(note *ClinicalNote) error {
	return r.db.Save(note).Error
}

// Delete removes a clinical note.
func (r *repository) Delete(id uint) error {
	return r.db.Delete(&ClinicalNote{}, id).Error
}
