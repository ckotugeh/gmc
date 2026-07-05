package doctor_schedules

// CreateDoctorScheduleRequest represents a request to create a doctor's schedule.
type CreateDoctorScheduleRequest struct {
	DoctorID             uint    `json:"doctor_id" binding:"required"`
	Day                  Weekday `json:"day" binding:"required"`
	StartTime            string  `json:"start_time" binding:"required"` // HH:MM
	EndTime              string  `json:"end_time" binding:"required"`   // HH:MM
	BreakStart           *string `json:"break_start,omitempty"`         // HH:MM
	BreakEnd             *string `json:"break_end,omitempty"`           // HH:MM
	ConsultationDuration int     `json:"consultation_duration" binding:"required,gt=0"`
	MaxPatients          int     `json:"max_patients" binding:"required,gt=0"`
	IsActive             *bool   `json:"is_active,omitempty"`
}

// UpdateDoctorScheduleRequest represents a request to update a doctor's schedule.
type UpdateDoctorScheduleRequest struct {
	Day                  Weekday `json:"day"`
	StartTime            string  `json:"start_time"`
	EndTime              string  `json:"end_time"`
	BreakStart           *string `json:"break_start,omitempty"`
	BreakEnd             *string `json:"break_end,omitempty"`
	ConsultationDuration int     `json:"consultation_duration"`
	MaxPatients          int     `json:"max_patients"`
	IsActive             *bool   `json:"is_active,omitempty"`
}

// DoctorScheduleResponse represents a doctor's schedule returned to the client.
type DoctorScheduleResponse struct {
	ID                   uint    `json:"id"`
	DoctorID             uint    `json:"doctor_id"`
	Day                  Weekday `json:"day"`
	StartTime            string  `json:"start_time"`
	EndTime              string  `json:"end_time"`
	BreakStart           *string `json:"break_start,omitempty"`
	BreakEnd             *string `json:"break_end,omitempty"`
	ConsultationDuration int     `json:"consultation_duration"`
	MaxPatients          int     `json:"max_patients"`
	IsActive             bool    `json:"is_active"`
}

// DoctorScheduleFilterRequest represents schedule query filters.
type DoctorScheduleFilterRequest struct {
	DoctorID uint    `form:"doctor_id"`
	Day      Weekday `form:"day"`
	IsActive *bool   `form:"is_active"`
}
