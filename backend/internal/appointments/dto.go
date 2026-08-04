package appointments

import "time"

// CreateAppointmentRequest represents the payload for booking an appointment.
type CreateAppointmentRequest struct {
	DoctorID        uint      `json:"doctor_id" binding:"required"`
	AppointmentTime time.Time `json:"appointment_time" binding:"required"`
	DurationMinutes int       `json:"duration_minutes"`
	Reason          string    `json:"reason" binding:"required"`
	MeetingLink     string    `json:"meeting_link,omitempty"`
}

// UpdateAppointmentRequest represents the payload for updating an appointment.
type UpdateAppointmentRequest struct {
	AppointmentTime *time.Time `json:"appointment_time,omitempty"`
	DurationMinutes *int       `json:"duration_minutes,omitempty"`
	Status          *string    `json:"status,omitempty"`
	Reason          *string    `json:"reason,omitempty"`
	Notes           *string    `json:"notes,omitempty"`
	MeetingLink     *string    `json:"meeting_link,omitempty"`
}

// AppointmentResponse represents an appointment returned to the client.
type AppointmentResponse struct {
	ID              uint      `json:"id"`
	DoctorID        uint      `json:"doctor_id"`
	PatientID       uint      `json:"patient_id"`
	AppointmentTime time.Time `json:"appointment_time"`
	DurationMinutes int       `json:"duration_minutes"`
	Status          string    `json:"status"`
	Reason          string    `json:"reason"`
	Notes           string    `json:"notes,omitempty"`
	MeetingLink     string    `json:"meeting_link,omitempty"`
	CreatedAt       time.Time `json:"created_at"`
	UpdatedAt       time.Time `json:"updated_at"`
}

// ToResponse converts an Appointment model into an AppointmentResponse.
func ToResponse(a *Appointment) AppointmentResponse {
	return AppointmentResponse{
		ID:              a.ID,
		DoctorID:        a.DoctorID,
		PatientID:       a.PatientID,
		AppointmentTime: a.AppointmentTime,
		DurationMinutes: a.DurationMinutes,
		Status:          a.Status,
		Reason:          a.Reason,
		Notes:           a.Notes,
		MeetingLink:     a.MeetingLink,
		CreatedAt:       a.CreatedAt,
		UpdatedAt:       a.UpdatedAt,
	}
}

// ToResponseList converts a slice of Appointment models into response DTOs.
func ToResponseList(appointments []Appointment) []AppointmentResponse {
	response := make([]AppointmentResponse, 0, len(appointments))

	for _, appointment := range appointments {
		response = append(response, ToResponse(&appointment))
	}

	return response
}
