package hospitals

// MockRepository is an in-memory implementation of Repository for testing.
type MockRepository struct {
	hospitals []Hospital
	nextID    uint
}

// NewMockRepository creates a new mock repository.
func NewMockRepository() Repository {
	return &MockRepository{
		hospitals: []Hospital{},
		nextID:    1,
	}
}

// Create creates a new hospital.
func (m *MockRepository) Create(hospital *Hospital) error {
	hospital.ID = m.nextID
	m.nextID++
	m.hospitals = append(m.hospitals, *hospital)
	return nil
}

// GetByID retrieves a hospital by ID.
func (m *MockRepository) GetByID(id uint) (*Hospital, error) {
	for _, hospital := range m.hospitals {
		if hospital.ID == id {
			h := hospital
			return &h, nil
		}
	}
	return nil, ErrHospitalNotFound
}

// GetAll retrieves all hospitals.
func (m *MockRepository) GetAll() ([]Hospital, error) {
	return m.hospitals, nil
}

// GetByEmail retrieves a hospital by email.
func (m *MockRepository) GetByEmail(email string) (*Hospital, error) {
	for _, hospital := range m.hospitals {
		if hospital.Email == email {
			h := hospital
			return &h, nil
		}
	}
	return nil, ErrHospitalNotFound
}

// GetByLicenseNumber retrieves a hospital by license number.
func (m *MockRepository) GetByLicenseNumber(license string) (*Hospital, error) {
	for _, hospital := range m.hospitals {
		if hospital.LicenseNumber == license {
			h := hospital
			return &h, nil
		}
	}
	return nil, ErrHospitalNotFound
}

// Update updates a hospital.
func (m *MockRepository) Update(updated *Hospital) error {
	for i, hospital := range m.hospitals {
		if hospital.ID == updated.ID {
			m.hospitals[i] = *updated
			return nil
		}
	}
	return ErrHospitalNotFound
}

// Delete deletes a hospital.
func (m *MockRepository) Delete(id uint) error {
	for i, hospital := range m.hospitals {
		if hospital.ID == id {
			m.hospitals = append(m.hospitals[:i], m.hospitals[i+1:]...)
			return nil
		}
	}
	return ErrHospitalNotFound
}
