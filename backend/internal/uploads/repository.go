package uploads

import "gorm.io/gorm"

// Repository defines the upload repository contract.
type Repository interface {
	Create(upload *Upload) error
	GetByID(id uint) (*Upload, error)
	GetByUserID(userID uint) ([]Upload, error)
	GetByAppointmentID(appointmentID uint) ([]Upload, error)
	GetByMedicalRecordID(recordID uint) ([]Upload, error)
	GetByHospitalID(hospitalID uint) ([]Upload, error)
	Update(upload *Upload) error
	Delete(id uint) error
}

type repository struct {
	db *gorm.DB
}

// NewRepository creates a new upload repository.
func NewRepository(db *gorm.DB) Repository {
	return &repository{
		db: db,
	}
}

// Create stores a new upload.
func (r *repository) Create(upload *Upload) error {
	return r.db.Create(upload).Error
}

// GetByID retrieves an upload by ID.
func (r *repository) GetByID(id uint) (*Upload, error) {
	var upload Upload
	if err := r.db.First(&upload, id).Error; err != nil {
		return nil, err
	}
	return &upload, nil
}

// GetByUserID retrieves all uploads for a user.
func (r *repository) GetByUserID(userID uint) ([]Upload, error) {
	var uploads []Upload
	if err := r.db.Where("user_id = ?", userID).Find(&uploads).Error; err != nil {
		return nil, err
	}
	return uploads, nil
}

// GetByAppointmentID retrieves uploads for an appointment.
func (r *repository) GetByAppointmentID(appointmentID uint) ([]Upload, error) {
	var uploads []Upload
	if err := r.db.Where("appointment_id = ?", appointmentID).Find(&uploads).Error; err != nil {
		return nil, err
	}
	return uploads, nil
}

// GetByMedicalRecordID retrieves uploads for a medical record.
func (r *repository) GetByMedicalRecordID(recordID uint) ([]Upload, error) {
	var uploads []Upload
	if err := r.db.Where("medical_record_id = ?", recordID).Find(&uploads).Error; err != nil {
		return nil, err
	}
	return uploads, nil
}

// GetByHospitalID retrieves uploads for a hospital.
func (r *repository) GetByHospitalID(hospitalID uint) ([]Upload, error) {
	var uploads []Upload
	if err := r.db.Where("hospital_id = ?", hospitalID).Find(&uploads).Error; err != nil {
		return nil, err
	}
	return uploads, nil
}

// Update updates upload metadata.
func (r *repository) Update(upload *Upload) error {
	return r.db.Save(upload).Error
}

// Delete removes an upload.
func (r *repository) Delete(id uint) error {
	return r.db.Delete(&Upload{}, id).Error
}
