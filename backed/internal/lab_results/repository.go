package lab_results

import "gorm.io/gorm"

// Repository defines the lab result repository contract.
type Repository interface {
	Create(result *LabResult) error
	GetByID(id uint) (*LabResult, error)
	GetAll() ([]LabResult, error)
	GetByLabRequestID(labRequestID uint) ([]LabResult, error)
	GetByPatientID(patientID uint) ([]LabResult, error)
	GetByDoctorID(doctorID uint) ([]LabResult, error)
	GetByStatus(status string) ([]LabResult, error)
	Update(result *LabResult) error
	Delete(id uint) error
}

type repository struct {
	db *gorm.DB
}

// NewRepository creates a new lab result repository.
func NewRepository(db *gorm.DB) Repository {
	return &repository{
		db: db,
	}
}

// Create creates a new lab result.
func (r *repository) Create(result *LabResult) error {
	return r.db.Create(result).Error
}

// GetByID retrieves a lab result by ID.
func (r *repository) GetByID(id uint) (*LabResult, error) {
	var result LabResult

	if err := r.db.First(&result, id).Error; err != nil {
		return nil, err
	}

	return &result, nil
}

// GetAll retrieves all lab results.
func (r *repository) GetAll() ([]LabResult, error) {
	var results []LabResult

	if err := r.db.Find(&results).Error; err != nil {
		return nil, err
	}

	return results, nil
}

// GetByLabRequestID retrieves results for a lab request.
func (r *repository) GetByLabRequestID(labRequestID uint) ([]LabResult, error) {
	var results []LabResult

	if err := r.db.Where("lab_request_id = ?", labRequestID).
		Find(&results).Error; err != nil {
		return nil, err
	}

	return results, nil
}

// GetByPatientID retrieves lab results for a patient.
func (r *repository) GetByPatientID(patientID uint) ([]LabResult, error) {
	var results []LabResult

	if err := r.db.Where("patient_id = ?", patientID).
		Find(&results).Error; err != nil {
		return nil, err
	}

	return results, nil
}

// GetByDoctorID retrieves lab results recorded by a doctor.
func (r *repository) GetByDoctorID(doctorID uint) ([]LabResult, error) {
	var results []LabResult

	if err := r.db.Where("doctor_id = ?", doctorID).
		Find(&results).Error; err != nil {
		return nil, err
	}

	return results, nil
}

// GetByStatus retrieves lab results by status.
func (r *repository) GetByStatus(status string) ([]LabResult, error) {
	var results []LabResult

	if err := r.db.Where("status = ?", status).
		Find(&results).Error; err != nil {
		return nil, err
	}

	return results, nil
}

// Update updates a lab result.
func (r *repository) Update(result *LabResult) error {
	return r.db.Save(result).Error
}

// Delete deletes a lab result.
func (r *repository) Delete(id uint) error {
	return r.db.Delete(&LabResult{}, id).Error
}
