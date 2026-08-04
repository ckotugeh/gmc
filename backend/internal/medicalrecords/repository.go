package medicalrecords

import (
	"gorm.io/gorm"
)

// Repository defines the medical record repository contract.
type Repository interface {
	Create(record *MedicalRecord) error
	GetByID(id uint) (*MedicalRecord, error)
	GetAll() ([]MedicalRecord, error)
	GetByPatientID(patientID uint) ([]MedicalRecord, error)
	GetByDoctorID(doctorID uint) ([]MedicalRecord, error)
	Update(record *MedicalRecord) error
	Delete(id uint) error
}

type repository struct {
	db *gorm.DB
}

// NewRepository creates a new medical record repository.
func NewRepository(db *gorm.DB) Repository {
	return &repository{db: db}
}

// Create creates a new medical record.
func (r *repository) Create(record *MedicalRecord) error {
	return r.db.Create(record).Error
}

// GetByID retrieves a medical record by ID.
func (r *repository) GetByID(id uint) (*MedicalRecord, error) {
	var record MedicalRecord
	if err := r.db.First(&record, id).Error; err != nil {
		return nil, err
	}
	return &record, nil
}

// GetAll retrieves all medical records.
func (r *repository) GetAll() ([]MedicalRecord, error) {
	var records []MedicalRecord
	if err := r.db.Find(&records).Error; err != nil {
		return nil, err
	}
	return records, nil
}

// GetByPatientID retrieves all medical records for a patient.
func (r *repository) GetByPatientID(patientID uint) ([]MedicalRecord, error) {
	var records []MedicalRecord
	if err := r.db.Where("patient_id = ?", patientID).Find(&records).Error; err != nil {
		return nil, err
	}
	return records, nil
}

// GetByDoctorID retrieves all medical records created by a doctor.
func (r *repository) GetByDoctorID(doctorID uint) ([]MedicalRecord, error) {
	var records []MedicalRecord
	if err := r.db.Where("doctor_id = ?", doctorID).Find(&records).Error; err != nil {
		return nil, err
	}
	return records, nil
}

// Update updates a medical record.
func (r *repository) Update(record *MedicalRecord) error {
	return r.db.Save(record).Error
}

// Delete deletes a medical record.
func (r *repository) Delete(id uint) error {
	return r.db.Delete(&MedicalRecord{}, id).Error
}
