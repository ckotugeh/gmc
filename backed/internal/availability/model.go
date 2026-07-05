package availability

import (
	"time"

	"gorm.io/gorm"
)

// SlotStatus represents the availability status of a time slot.
type SlotStatus string

const (
	SlotAvailable SlotStatus = "available"
	SlotBooked    SlotStatus = "booked"
	SlotBlocked   SlotStatus = "blocked"
)

// Availability represents a doctor's available appointment slot.
type Availability struct {
	ID uint `gorm:"primaryKey" json:"id"`

	// Relationships
	DoctorID   uint `gorm:"not null;index" json:"doctor_id"`
	ScheduleID uint `gorm:"not null;index" json:"schedule_id"`

	// Slot details
	Date      time.Time `gorm:"type:date;not null;index" json:"date"`
	StartTime string    `gorm:"size:5;not null" json:"start_time"` // HH:MM
	EndTime   string    `gorm:"size:5;not null" json:"end_time"`   // HH:MM

	// Booking status
	Status SlotStatus `gorm:"type:varchar(20);default:'available'" json:"status"`

	// Optional appointment linked to this slot
	AppointmentID *uint `gorm:"index" json:"appointment_id,omitempty"`

	// Notes (e.g. "Doctor on leave", "Reserved")
	Notes string `gorm:"type:text" json:"notes,omitempty"`

	CreatedAt time.Time      `json:"created_at"`
	UpdatedAt time.Time      `json:"updated_at"`
	DeletedAt gorm.DeletedAt `gorm:"index" json:"-"`
}
