package doctor_reviews

import "gorm.io/gorm"

// Repository defines the doctor review repository contract.
type Repository interface {
	Create(review *DoctorReview) error
	GetByID(id uint) (*DoctorReview, error)
	GetByAppointmentID(appointmentID uint) (*DoctorReview, error)
	GetByDoctorID(doctorID uint) ([]DoctorReview, error)
	GetByPatientID(patientID uint) ([]DoctorReview, error)
	GetPublishedByDoctorID(doctorID uint) ([]DoctorReview, error)
	GetAll() ([]DoctorReview, error)
	Update(review *DoctorReview) error
	Delete(id uint) error
}

type repository struct {
	db *gorm.DB
}

// NewRepository creates a new doctor review repository.
func NewRepository(db *gorm.DB) Repository {
	return &repository{
		db: db,
	}
}

// Create stores a doctor review.
func (r *repository) Create(review *DoctorReview) error {
	return r.db.Create(review).Error
}

// GetByID retrieves a doctor review by ID.
func (r *repository) GetByID(id uint) (*DoctorReview, error) {
	var review DoctorReview

	if err := r.db.First(&review, id).Error; err != nil {
		return nil, err
	}

	return &review, nil
}

// GetByAppointmentID retrieves a review by appointment ID.
func (r *repository) GetByAppointmentID(appointmentID uint) (*DoctorReview, error) {
	var review DoctorReview

	if err := r.db.Where("appointment_id = ?", appointmentID).
		First(&review).Error; err != nil {
		return nil, err
	}

	return &review, nil
}

// GetByDoctorID retrieves all reviews for a doctor.
func (r *repository) GetByDoctorID(doctorID uint) ([]DoctorReview, error) {
	var reviews []DoctorReview

	if err := r.db.
		Where("doctor_id = ?", doctorID).
		Order("created_at DESC").
		Find(&reviews).Error; err != nil {
		return nil, err
	}

	return reviews, nil
}

// GetByPatientID retrieves all reviews written by a patient.
func (r *repository) GetByPatientID(patientID uint) ([]DoctorReview, error) {
	var reviews []DoctorReview

	if err := r.db.
		Where("patient_id = ?", patientID).
		Order("created_at DESC").
		Find(&reviews).Error; err != nil {
		return nil, err
	}

	return reviews, nil
}

// GetPublishedByDoctorID retrieves published reviews for a doctor.
func (r *repository) GetPublishedByDoctorID(doctorID uint) ([]DoctorReview, error) {
	var reviews []DoctorReview

	if err := r.db.
		Where("doctor_id = ? AND is_published = ?", doctorID, true).
		Order("created_at DESC").
		Find(&reviews).Error; err != nil {
		return nil, err
	}

	return reviews, nil
}

// GetAll retrieves all doctor reviews.
func (r *repository) GetAll() ([]DoctorReview, error) {
	var reviews []DoctorReview

	if err := r.db.
		Order("created_at DESC").
		Find(&reviews).Error; err != nil {
		return nil, err
	}

	return reviews, nil
}

// Update updates a doctor review.
func (r *repository) Update(review *DoctorReview) error {
	return r.db.Save(review).Error
}

// Delete removes a doctor review.
func (r *repository) Delete(id uint) error {
	return r.db.Delete(&DoctorReview{}, id).Error
}
