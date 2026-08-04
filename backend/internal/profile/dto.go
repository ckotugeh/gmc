package profile

type CreateProfileRequest struct {
	Specialization string `json:"specialization" binding:"required"`
	Hospital       string `json:"hospital"`

	Country string `json:"country"`
	City    string `json:"city"`

	YearsExperience int `json:"years_experience"`

	LicenseNumber string `json:"license_number"`

	Education string `json:"education"`

	Languages string `json:"languages"`

	Bio string `json:"bio"`

	ProfileImageURL string `json:"profile_image_url"`
}

type UpdateProfileRequest struct {
	Specialization string `json:"specialization"`
	Hospital       string `json:"hospital"`

	Country string `json:"country"`
	City    string `json:"city"`

	YearsExperience int `json:"years_experience"`

	LicenseNumber string `json:"license_number"`

	Education string `json:"education"`

	Languages string `json:"languages"`

	Bio string `json:"bio"`

	ProfileImageURL string `json:"profile_image_url"`
}

type ProfileResponse struct {
	ID uint `json:"id"`

	FullName string `json:"full_name"`

	Specialization string `json:"specialization"`

	Hospital string `json:"hospital"`

	Country string `json:"country"`

	City string `json:"city"`

	YearsExperience int `json:"years_experience"`

	LicenseNumber string `json:"license_number"`

	Education string `json:"education"`

	Languages string `json:"languages"`

	Bio string `json:"bio"`

	ProfileImageURL string `json:"profile_image_url"`
}
