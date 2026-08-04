package admin

import "time"

// CreateAdminRequest represents a request to create an admin action.
type CreateAdminRequest struct {
	ResourceID  uint         `json:"resource_id" binding:"required"`
	Resource    ResourceType `json:"resource" binding:"required"`
	Action      AdminAction  `json:"action" binding:"required"`
	Description string       `json:"description"`
}

// UpdateAdminRequest represents a request to update an admin action.
type UpdateAdminRequest struct {
	Action      AdminAction `json:"action"`
	Description string      `json:"description"`
}

// AdminResponse represents an admin action response.
type AdminResponse struct {
	ID          uint         `json:"id"`
	AdminID     uint         `json:"admin_id"`
	ResourceID  uint         `json:"resource_id"`
	Resource    ResourceType `json:"resource"`
	Action      AdminAction  `json:"action"`
	Description string       `json:"description"`
	CreatedAt   time.Time    `json:"created_at"`
	UpdatedAt   time.Time    `json:"updated_at"`
}

// DashboardStatsResponse represents statistics shown on the admin dashboard.
type DashboardStatsResponse struct {
	TotalUsers        int64 `json:"total_users"`
	TotalDoctors      int64 `json:"total_doctors"`
	TotalPatients     int64 `json:"total_patients"`
	TotalHospitals    int64 `json:"total_hospitals"`
	TotalCommunities  int64 `json:"total_communities"`
	TotalPosts        int64 `json:"total_posts"`
	TotalAppointments int64 `json:"total_appointments"`
}

// AdminFilterRequest represents query filters for admin actions.
type AdminFilterRequest struct {
	Resource ResourceType `form:"resource"`
	Action   AdminAction  `form:"action"`
	AdminID  uint         `form:"admin_id"`
}
