package lab_requests

import "gorm.io/gorm"

// Repository defines the lab request repository contract.
type Repository interface {
	Create(request *LabRequest) error
	GetByID(id uint) (*LabRequest, error)
	GetAll() ([]LabRequest, error)
	GetByPatientID(patientID uint) ([]LabRequest, error)
	GetByDoctorID(doctorID uint) ([]LabRequest, error)
	GetByAppointmentID(appointmentID uint) ([]LabRequest, error)
	GetByStatus(status string) ([]LabRequest, error)
	Update(request *LabRequest) error
	Delete(id uint) error
}

type repository struct {
	db *gorm.DB
}

// NewRepository creates a new lab request repository.
func NewRepository(db *gorm.DB) Repository {
	return &repository{
		db: db,
	}
}

// Create creates a new lab request.
func (r *repository) Create(request *LabRequest) error {
	return r.db.Create(request).Error
}

// GetByID retrieves a lab request by ID.
func (r *repository) GetByID(id uint) (*LabRequest, error) {
	var request LabRequest

	if err := r.db.First(&request, id).Error; err != nil {
		return nil, err
	}

	return &request, nil
}

// GetAll retrieves all lab requests.
func (r *repository) GetAll() ([]LabRequest, error) {
	var requests []LabRequest

	if err := r.db.Find(&requests).Error; err != nil {
		return nil, err
	}

	return requests, nil
}

// GetByPatientID retrieves lab requests for a patient.
func (r *repository) GetByPatientID(patientID uint) ([]LabRequest, error) {
	var requests []LabRequest

	if err := r.db.Where("patient_id = ?", patientID).
		Find(&requests).Error; err != nil {
		return nil, err
	}

	return requests, nil
}

// GetByDoctorID retrieves lab requests created by a doctor.
func (r *repository) GetByDoctorID(doctorID uint) ([]LabRequest, error) {
	var requests []LabRequest

	if err := r.db.Where("doctor_id = ?", doctorID).
		Find(&requests).Error; err != nil {
		return nil, err
	}

	return requests, nil
}

// GetByAppointmentID retrieves lab requests for an appointment.
func (r *repository) GetByAppointmentID(appointmentID uint) ([]LabRequest, error) {
	var requests []LabRequest

	if err := r.db.Where("appointment_id = ?", appointmentID).
		Find(&requests).Error; err != nil {
		return nil, err
	}

	return requests, nil
}

// GetByStatus retrieves lab requests by status.
func (r *repository) GetByStatus(status string) ([]LabRequest, error) {
	var requests []LabRequest

	if err := r.db.Where("status = ?", status).
		Find(&requests).Error; err != nil {
		return nil, err
	}

	return requests, nil
}

// Update updates a lab request.
func (r *repository) Update(request *LabRequest) error {
	return r.db.Save(request).Error
}

// Delete deletes a lab request.
func (r *repository) Delete(id uint) error {
	return r.db.Delete(&LabRequest{}, id).Error
}
