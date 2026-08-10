package video_consultations

import "time"

// CreateVideoConsultationRequest represents the payload for creating a video consultation.
type CreateVideoConsultationRequest struct {
	AppointmentID uint      `json:"appointment_id" binding:"required"`
	DoctorID      uint      `json:"doctor_id" binding:"required"`
	PatientID     uint      `json:"patient_id" binding:"required"`
	ScheduledAt   time.Time `json:"scheduled_at" binding:"required"`
	Notes         string    `json:"notes"`
}

// UpdateVideoConsultationRequest represents the payload for updating a video consultation.
type UpdateVideoConsultationRequest struct {
	ScheduledAt *time.Time          `json:"scheduled_at,omitempty"`
	Status      *ConsultationStatus `json:"status,omitempty"`
	Notes       *string             `json:"notes,omitempty"`
}

// VideoConsultationResponse represents a video consultation returned to the client.
type VideoConsultationResponse struct {
	ID            uint               `json:"id"`
	AppointmentID uint               `json:"appointment_id"`
	DoctorID      uint               `json:"doctor_id"`
	PatientID     uint               `json:"patient_id"`
	RoomID        string             `json:"room_id"`
	SessionKey    string             `json:"session_key,omitempty"`
	ScheduledAt   time.Time          `json:"scheduled_at"`
	StartedAt     *time.Time         `json:"started_at,omitempty"`
	EndedAt       *time.Time         `json:"ended_at,omitempty"`
	Status        ConsultationStatus `json:"status"`
	Notes         string             `json:"notes,omitempty"`
	CreatedAt     time.Time          `json:"created_at"`
	UpdatedAt     time.Time          `json:"updated_at"`
}
