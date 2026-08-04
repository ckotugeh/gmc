package video_consultations

import (
	"time"

	"gorm.io/gorm"
)

// ConsultationStatus represents the state of a consultation.
type ConsultationStatus string

const (
	StatusScheduled ConsultationStatus = "scheduled"
	StatusOngoing   ConsultationStatus = "ongoing"
	StatusCompleted ConsultationStatus = "completed"
	StatusCancelled ConsultationStatus = "cancelled"
)

// VideoConsultation represents a doctor-patient video consultation session.
type VideoConsultation struct {
	ID uint `gorm:"primaryKey" json:"id"`

	// Relationships
	AppointmentID uint `gorm:"not null;uniqueIndex" json:"appointment_id"`
	DoctorID      uint `gorm:"not null;index" json:"doctor_id"`
	PatientID     uint `gorm:"not null;index" json:"patient_id"`

	// Session information
	RoomID     string `gorm:"size:255;not null;uniqueIndex" json:"room_id"`
	SessionKey string `gorm:"size:255" json:"session_key,omitempty"`

	// Timing
	ScheduledAt time.Time  `json:"scheduled_at"`
	StartedAt   *time.Time `json:"started_at,omitempty"`
	EndedAt     *time.Time `json:"ended_at,omitempty"`

	// Status
	Status ConsultationStatus `gorm:"type:varchar(20);default:'scheduled'" json:"status"`

	// Optional notes
	Notes string `gorm:"type:text" json:"notes,omitempty"`

	CreatedAt time.Time      `json:"created_at"`
	UpdatedAt time.Time      `json:"updated_at"`
	DeletedAt gorm.DeletedAt `gorm:"index" json:"-"`
}
