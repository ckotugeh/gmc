package appointments

import (
	"time"

	"gorm.io/gorm"
)

// Appointment status constants.
const (
	StatusPending   = "pending"
	StatusConfirmed = "confirmed"
	StatusCompleted = "completed"
	StatusCancelled = "cancelled"
)

// Appointment represents a doctor's appointment with a patient.
type Appointment struct {
	ID uint `gorm:"primaryKey" json:"id"`

	DoctorID  uint `gorm:"not null;index" json:"doctor_id"`
	PatientID uint `gorm:"not null;index" json:"patient_id"`

	AppointmentTime time.Time `gorm:"not null" json:"appointment_time"`
	DurationMinutes int       `gorm:"default:30" json:"duration_minutes"`

	Status string `gorm:"type:varchar(20);default:'pending'" json:"status"`

	Reason string `gorm:"type:text" json:"reason"`
	Notes  string `gorm:"type:text" json:"notes,omitempty"`

	MeetingLink string `gorm:"size:255" json:"meeting_link,omitempty"`

	CreatedAt time.Time      `json:"created_at"`
	UpdatedAt time.Time      `json:"updated_at"`
	DeletedAt gorm.DeletedAt `gorm:"index" json:"-"`
}
