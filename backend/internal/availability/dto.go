package availability

import "time"

// CreateAvailabilityRequest represents a request to create an availability slot.
type CreateAvailabilityRequest struct {
	DoctorID      uint       `json:"doctor_id" binding:"required"`
	ScheduleID    uint       `json:"schedule_id" binding:"required"`
	Date          time.Time  `json:"date" binding:"required"`
	StartTime     string     `json:"start_time" binding:"required"` // HH:MM
	EndTime       string     `json:"end_time" binding:"required"`   // HH:MM
	Status        SlotStatus `json:"status"`
	AppointmentID *uint      `json:"appointment_id,omitempty"`
	Notes         string     `json:"notes,omitempty"`
}

// UpdateAvailabilityRequest represents a request to update an availability slot.
type UpdateAvailabilityRequest struct {
	StartTime     string     `json:"start_time,omitempty"`
	EndTime       string     `json:"end_time,omitempty"`
	Status        SlotStatus `json:"status,omitempty"`
	AppointmentID *uint      `json:"appointment_id,omitempty"`
	Notes         string     `json:"notes,omitempty"`
}

// AvailabilityResponse represents an availability slot returned to the client.
type AvailabilityResponse struct {
	ID            uint       `json:"id"`
	DoctorID      uint       `json:"doctor_id"`
	ScheduleID    uint       `json:"schedule_id"`
	Date          time.Time  `json:"date"`
	StartTime     string     `json:"start_time"`
	EndTime       string     `json:"end_time"`
	Status        SlotStatus `json:"status"`
	AppointmentID *uint      `json:"appointment_id,omitempty"`
	Notes         string     `json:"notes,omitempty"`
}

// AvailabilityFilterRequest represents query parameters for filtering slots.
type AvailabilityFilterRequest struct {
	DoctorID uint       `form:"doctor_id"`
	Date     *time.Time `form:"date"`
	Status   SlotStatus `form:"status"`
}
