package doctor_schedules

import (
	"time"

	"gorm.io/gorm"
)

// Weekday represents a day of the week.
type Weekday string

const (
	Monday    Weekday = "monday"
	Tuesday   Weekday = "tuesday"
	Wednesday Weekday = "wednesday"
	Thursday  Weekday = "thursday"
	Friday    Weekday = "friday"
	Saturday  Weekday = "saturday"
	Sunday    Weekday = "sunday"
)

// DoctorSchedule represents a doctor's working schedule.
type DoctorSchedule struct {
	ID uint `gorm:"primaryKey" json:"id"`

	// Relationships
	DoctorID uint `gorm:"not null;index" json:"doctor_id"`

	// Schedule details
	Day       Weekday `gorm:"type:varchar(20);not null;index" json:"day"`
	StartTime string  `gorm:"size:5;not null" json:"start_time"` // HH:MM
	EndTime   string  `gorm:"size:5;not null" json:"end_time"`   // HH:MM

	// Optional break
	BreakStart *string `gorm:"size:5" json:"break_start,omitempty"` // HH:MM
	BreakEnd   *string `gorm:"size:5" json:"break_end,omitempty"`   // HH:MM

	// Appointment configuration
	ConsultationDuration int `gorm:"not null;default:30" json:"consultation_duration"` // minutes
	MaxPatients          int `gorm:"not null;default:20" json:"max_patients"`

	// Status
	IsActive bool `gorm:"default:true" json:"is_active"`

	CreatedAt time.Time      `json:"created_at"`
	UpdatedAt time.Time      `json:"updated_at"`
	DeletedAt gorm.DeletedAt `gorm:"index" json:"-"`
}
