package admin

import "gorm.io/gorm"

// Repository defines the admin repository contract.
type Repository interface {
	Create(admin *Admin) error
	GetByID(id uint) (*Admin, error)
	GetAll() ([]Admin, error)
	GetByAdminID(adminID uint) ([]Admin, error)
	Update(admin *Admin) error
	Delete(id uint) error
	GetDashboardStats() (*DashboardStatsResponse, error)
}

type repository struct {
	db *gorm.DB
}

// NewRepository creates a new admin repository.
func NewRepository(db *gorm.DB) Repository {
	return &repository{
		db: db,
	}
}

// Create stores an admin action.
func (r *repository) Create(admin *Admin) error {
	return r.db.Create(admin).Error
}

// GetByID retrieves an admin action by ID.
func (r *repository) GetByID(id uint) (*Admin, error) {
	var admin Admin

	if err := r.db.First(&admin, id).Error; err != nil {
		return nil, err
	}

	return &admin, nil
}

// GetAll retrieves all admin actions.
func (r *repository) GetAll() ([]Admin, error) {
	var admins []Admin

	if err := r.db.Order("created_at DESC").Find(&admins).Error; err != nil {
		return nil, err
	}

	return admins, nil
}

// GetByAdminID retrieves all actions performed by an administrator.
func (r *repository) GetByAdminID(adminID uint) ([]Admin, error) {
	var admins []Admin

	if err := r.db.Where("admin_id = ?", adminID).
		Order("created_at DESC").
		Find(&admins).Error; err != nil {
		return nil, err
	}

	return admins, nil
}

// Update updates an admin action.
func (r *repository) Update(admin *Admin) error {
	return r.db.Save(admin).Error
}

// Delete removes an admin action.
func (r *repository) Delete(id uint) error {
	return r.db.Delete(&Admin{}, id).Error
}

// GetDashboardStats returns platform statistics.
func (r *repository) GetDashboardStats() (*DashboardStatsResponse, error) {
	stats := &DashboardStatsResponse{}

	// These values will be populated once the corresponding
	// repositories/models are integrated.
	stats.TotalUsers = 0
	stats.TotalDoctors = 0
	stats.TotalPatients = 0
	stats.TotalHospitals = 0
	stats.TotalCommunities = 0
	stats.TotalPosts = 0
	stats.TotalAppointments = 0

	return stats, nil
}
