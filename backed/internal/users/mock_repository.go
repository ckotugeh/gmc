package users

import "errors"

// MockRepository is an in-memory implementation of Repository for testing.
type MockRepository struct {
	users  map[uint]*User
	nextID uint
}

// NewMockRepository creates a new mock repository.
func NewMockRepository() *MockRepository {
	return &MockRepository{
		users:  make(map[uint]*User),
		nextID: 1,
	}
}

// Create creates a new user.
func (m *MockRepository) Create(user *User) error {
	user.ID = m.nextID
	m.users[user.ID] = user
	m.nextID++
	return nil
}

// GetByID retrieves a user by ID.
func (m *MockRepository) GetByID(id uint) (*User, error) {
	user, exists := m.users[id]
	if !exists {
		return nil, errors.New("user not found")
	}
	return user, nil
}

// GetByEmail retrieves a user by email.
func (m *MockRepository) GetByEmail(email string) (*User, error) {
	for _, user := range m.users {
		if user.Email == email {
			return user, nil
		}
	}
	return nil, errors.New("user not found")
}

// GetAll returns all users.
func (m *MockRepository) GetAll() ([]User, error) {
	users := make([]User, 0, len(m.users))

	for _, user := range m.users {
		users = append(users, *user)
	}

	return users, nil
}

// GetDoctors returns all doctors.
func (m *MockRepository) GetDoctors() ([]User, error) {
	var doctors []User

	for _, user := range m.users {
		if user.Role == "doctor" {
			doctors = append(doctors, *user)
		}
	}

	return doctors, nil
}

// GetPatients returns all patients.
func (m *MockRepository) GetPatients() ([]User, error) {
	var patients []User

	for _, user := range m.users {
		if user.Role == "patient" {
			patients = append(patients, *user)
		}
	}

	return patients, nil
}

// Update updates a user.
func (m *MockRepository) Update(user *User) error {
	if _, exists := m.users[user.ID]; !exists {
		return errors.New("user not found")
	}

	m.users[user.ID] = user
	return nil
}

// Delete deletes a user.
func (m *MockRepository) Delete(id uint) error {
	if _, exists := m.users[id]; !exists {
		return errors.New("user not found")
	}

	delete(m.users, id)
	return nil
}
