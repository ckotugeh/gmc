package hospitals

// CreateHospitalRequest represents the payload for creating a hospital.
type CreateHospitalRequest struct {
	Name          string `json:"name" binding:"required"`
	Description   string `json:"description"`
	Email         string `json:"email" binding:"required,email"`
	Phone         string `json:"phone"`
	Website       string `json:"website"`
	Address       string `json:"address"`
	City          string `json:"city"`
	State         string `json:"state"`
	Country       string `json:"country"`
	ZipCode       string `json:"zip_code"`
	LicenseNumber string `json:"license_number" binding:"required"`
	IsActive      bool   `json:"is_active"`
}

// UpdateHospitalRequest represents the payload for updating a hospital.
type UpdateHospitalRequest struct {
	Name          string `json:"name"`
	Description   string `json:"description"`
	Email         string `json:"email"`
	Phone         string `json:"phone"`
	Website       string `json:"website"`
	Address       string `json:"address"`
	City          string `json:"city"`
	State         string `json:"state"`
	Country       string `json:"country"`
	ZipCode       string `json:"zip_code"`
	LicenseNumber string `json:"license_number"`
	IsActive      *bool  `json:"is_active"`
}

// HospitalResponse represents the hospital returned to clients.
type HospitalResponse struct {
	ID            uint   `json:"id"`
	Name          string `json:"name"`
	Description   string `json:"description"`
	Email         string `json:"email"`
	Phone         string `json:"phone"`
	Website       string `json:"website"`
	Address       string `json:"address"`
	City          string `json:"city"`
	State         string `json:"state"`
	Country       string `json:"country"`
	ZipCode       string `json:"zip_code"`
	LicenseNumber string `json:"license_number"`
	IsActive      bool   `json:"is_active"`
	CreatedAt     string `json:"created_at"`
	UpdatedAt     string `json:"updated_at"`
}
