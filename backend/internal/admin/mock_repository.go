package admin

import "errors"

// MockRepository is an in-memory implementation of Repository.
type MockRepository struct {
	admins []Admin
	nextID uint
}

// NewMockRepository creates a new mock repository.
func NewMockRepository() Repository {
	return &MockRepository{
		admins: make([]Admin, 0),
		nextID: 1,
	}
}

// Create stores an admin action.
func (m *MockRepository) Create(admin *Admin) error {
	admin.ID = m.nextID
	m.nextID++
	m.admins = append(m.admins, *admin)
	return nil
}

// GetByID retrieves an admin action by ID.
func (m *MockRepository) GetByID(id uint) (*Admin, error) {
	for _, admin := range m.admins {
		if admin.ID == id {
			a := admin
			return &a, nil
		}
	}
	return nil, errors.New("admin action not found")
}

// GetAll retrieves all admin actions.
func (m *MockRepository) GetAll() ([]Admin, error) {
	return m.admins, nil
}

// GetByAdminID retrieves actions performed by an administrator.
func (m *MockRepository) GetByAdminID(adminID uint) ([]Admin, error) {
	var admins []Admin

	for _, admin := range m.admins {
		if admin.AdminID == adminID {
			admins = append(admins, admin)
		}
	}

	return admins, nil
}

// Update updates an existing admin action.
func (m *MockRepository) Update(updated *Admin) error {
	for i, admin := range m.admins {
		if admin.ID == updated.ID {
			m.admins[i] = *updated
			return nil
		}
	}

	return errors.New("admin action not found")
}

// Delete removes an admin action.
func (m *MockRepository) Delete(id uint) error {
	for i, admin := range m.admins {
		if admin.ID == id {
			m.admins = append(m.admins[:i], m.admins[i+1:]...)
			return nil
		}
	}

	return errors.New("admin action not found")
}

// GetDashboardStats returns mock dashboard statistics.
func (m *MockRepository) GetDashboardStats() (*DashboardStatsResponse, error) {
	stats := &DashboardStatsResponse{
		TotalUsers:        100,
		TotalDoctors:      25,
		TotalPatients:     75,
		TotalHospitals:    10,
		TotalCommunities:  15,
		TotalPosts:        200,
		TotalAppointments: 80,
	}

	return stats, nil
}
