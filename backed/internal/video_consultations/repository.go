package video_consultations

import "gorm.io/gorm"

// Repository defines the video consultation repository contract.
type Repository interface {
	Create(consultation *VideoConsultation) error
	GetByID(id uint) (*VideoConsultation, error)
	GetByAppointmentID(appointmentID uint) (*VideoConsultation, error)
	GetByDoctorID(doctorID uint) ([]VideoConsultation, error)
	GetByPatientID(patientID uint) ([]VideoConsultation, error)
	GetByRoomID(roomID string) (*VideoConsultation, error)
	GetAll() ([]VideoConsultation, error)
	Update(consultation *VideoConsultation) error
	Delete(id uint) error
}

type repository struct {
	db *gorm.DB
}

// NewRepository creates a new video consultation repository.
func NewRepository(db *gorm.DB) Repository {
	return &repository{
		db: db,
	}
}

// Create stores a new video consultation.
func (r *repository) Create(consultation *VideoConsultation) error {
	return r.db.Create(consultation).Error
}

// GetByID retrieves a consultation by ID.
func (r *repository) GetByID(id uint) (*VideoConsultation, error) {
	var consultation VideoConsultation
	if err := r.db.First(&consultation, id).Error; err != nil {
		return nil, err
	}
	return &consultation, nil
}

// GetByAppointmentID retrieves a consultation by appointment ID.
func (r *repository) GetByAppointmentID(appointmentID uint) (*VideoConsultation, error) {
	var consultation VideoConsultation
	if err := r.db.Where("appointment_id = ?", appointmentID).First(&consultation).Error; err != nil {
		return nil, err
	}
	return &consultation, nil
}

// GetByDoctorID retrieves all consultations for a doctor.
func (r *repository) GetByDoctorID(doctorID uint) ([]VideoConsultation, error) {
	var consultations []VideoConsultation
	if err := r.db.Where("doctor_id = ?", doctorID).
		Order("scheduled_at DESC").
		Find(&consultations).Error; err != nil {
		return nil, err
	}
	return consultations, nil
}

// GetByPatientID retrieves all consultations for a patient.
func (r *repository) GetByPatientID(patientID uint) ([]VideoConsultation, error) {
	var consultations []VideoConsultation
	if err := r.db.Where("patient_id = ?", patientID).
		Order("scheduled_at DESC").
		Find(&consultations).Error; err != nil {
		return nil, err
	}
	return consultations, nil
}

// GetByRoomID retrieves a consultation by room ID.
func (r *repository) GetByRoomID(roomID string) (*VideoConsultation, error) {
	var consultation VideoConsultation
	if err := r.db.Where("room_id = ?", roomID).First(&consultation).Error; err != nil {
		return nil, err
	}
	return &consultation, nil
}

// GetAll retrieves all consultations.
func (r *repository) GetAll() ([]VideoConsultation, error) {
	var consultations []VideoConsultation
	if err := r.db.Order("scheduled_at DESC").Find(&consultations).Error; err != nil {
		return nil, err
	}
	return consultations, nil
}

// Update updates an existing consultation.
func (r *repository) Update(consultation *VideoConsultation) error {
	return r.db.Save(consultation).Error
}

// Delete removes a consultation.
func (r *repository) Delete(id uint) error {
	return r.db.Delete(&VideoConsultation{}, id).Error
}
