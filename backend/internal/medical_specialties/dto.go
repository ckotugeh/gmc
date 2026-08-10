package medical_specialties

// CreateMedicalSpecialtyRequest represents a request to create a medical specialty.
type CreateMedicalSpecialtyRequest struct {
	Name        string `json:"name" binding:"required"`
	Code        string `json:"code" binding:"required"`
	Description string `json:"description"`
	IsActive    *bool  `json:"is_active,omitempty"`
}

// UpdateMedicalSpecialtyRequest represents a request to update a medical specialty.
type UpdateMedicalSpecialtyRequest struct {
	Name        string `json:"name"`
	Code        string `json:"code"`
	Description string `json:"description"`
	IsActive    *bool  `json:"is_active,omitempty"`
}

// MedicalSpecialtyResponse represents a medical specialty returned to the client.
type MedicalSpecialtyResponse struct {
	ID          uint   `json:"id"`
	Name        string `json:"name"`
	Code        string `json:"code"`
	Description string `json:"description"`
	IsActive    bool   `json:"is_active"`
}

// MedicalSpecialtyFilterRequest represents specialty query filters.
type MedicalSpecialtyFilterRequest struct {
	Name     string `form:"name"`
	Code     string `form:"code"`
	IsActive *bool  `form:"is_active"`
}
