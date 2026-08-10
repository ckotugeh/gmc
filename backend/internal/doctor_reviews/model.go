package doctor_reviews

import (
	"time"

	"gorm.io/gorm"
)

// DoctorReview represents a patient's review of a doctor.
type DoctorReview struct {
	ID uint `gorm:"primaryKey" json:"id"`

	// Relationships
	DoctorID      uint `gorm:"not null;index" json:"doctor_id"`
	PatientID     uint `gorm:"not null;index" json:"patient_id"`
	AppointmentID uint `gorm:"not null;uniqueIndex" json:"appointment_id"`

	// Review
	Rating  int    `gorm:"not null;check:rating >= 1 AND rating <= 5" json:"rating"`
	Title   string `gorm:"size:150" json:"title"`
	Comment string `gorm:"type:text" json:"comment"`

	// Moderation
	IsAnonymous bool `gorm:"default:false" json:"is_anonymous"`
	IsPublished bool `gorm:"default:true" json:"is_published"`

	CreatedAt time.Time      `json:"created_at"`
	UpdatedAt time.Time      `json:"updated_at"`
	DeletedAt gorm.DeletedAt `gorm:"index" json:"-"`
}
