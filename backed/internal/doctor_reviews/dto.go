package doctor_reviews

// CreateDoctorReviewRequest represents a request to create a doctor review.
type CreateDoctorReviewRequest struct {
	DoctorID      uint   `json:"doctor_id" binding:"required"`
	PatientID     uint   `json:"patient_id" binding:"required"`
	AppointmentID uint   `json:"appointment_id" binding:"required"`
	Rating        int    `json:"rating" binding:"required,min=1,max=5"`
	Title         string `json:"title"`
	Comment       string `json:"comment" binding:"required"`
	IsAnonymous   *bool  `json:"is_anonymous,omitempty"`
	IsPublished   *bool  `json:"is_published,omitempty"`
}

// UpdateDoctorReviewRequest represents a request to update a doctor review.
type UpdateDoctorReviewRequest struct {
	Rating      int    `json:"rating,omitempty"`
	Title       string `json:"title,omitempty"`
	Comment     string `json:"comment,omitempty"`
	IsAnonymous *bool  `json:"is_anonymous,omitempty"`
	IsPublished *bool  `json:"is_published,omitempty"`
}

// DoctorReviewResponse represents a doctor review returned to the client.
type DoctorReviewResponse struct {
	ID            uint   `json:"id"`
	DoctorID      uint   `json:"doctor_id"`
	PatientID     uint   `json:"patient_id"`
	AppointmentID uint   `json:"appointment_id"`
	Rating        int    `json:"rating"`
	Title         string `json:"title"`
	Comment       string `json:"comment"`
	IsAnonymous   bool   `json:"is_anonymous"`
	IsPublished   bool   `json:"is_published"`
}

// DoctorReviewFilterRequest represents query filters for doctor reviews.
type DoctorReviewFilterRequest struct {
	DoctorID    uint  `form:"doctor_id"`
	PatientID   uint  `form:"patient_id"`
	Rating      int   `form:"rating"`
	IsPublished *bool `form:"is_published"`
}
