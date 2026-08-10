package presence

import "errors"

// MockRepository is an in-memory implementation of Repository.
type MockRepository struct {
	Presences map[uint]*Presence
	NextID    uint
}

// NewMockRepository creates a new mock repository.
func NewMockRepository() *MockRepository {
	return &MockRepository{
		Presences: make(map[uint]*Presence),
		NextID:    1,
	}
}

// Create stores a new presence.
func (m *MockRepository) Create(presence *Presence) error {
	presence.ID = m.NextID
	m.Presences[presence.ID] = presence
	m.NextID++
	return nil
}

// Update updates an existing presence.
func (m *MockRepository) Update(presence *Presence) error {
	if _, ok := m.Presences[presence.ID]; !ok {
		return errors.New("presence not found")
	}

	m.Presences[presence.ID] = presence
	return nil
}

// GetByID retrieves a presence by ID.
func (m *MockRepository) GetByID(id uint) (*Presence, error) {
	presence, ok := m.Presences[id]
	if !ok {
		return nil, errors.New("presence not found")
	}

	return presence, nil
}

// GetByUserID retrieves a presence by user ID.
func (m *MockRepository) GetByUserID(userID uint) (*Presence, error) {
	for _, presence := range m.Presences {
		if presence.UserID == userID {
			return presence, nil
		}
	}

	return nil, errors.New("presence not found")
}

// GetOnlineUsers returns all online users.
func (m *MockRepository) GetOnlineUsers() ([]Presence, error) {
	var online []Presence

	for _, presence := range m.Presences {
		if presence.IsOnline {
			online = append(online, *presence)
		}
	}

	return online, nil
}

// Delete removes a presence.
func (m *MockRepository) Delete(id uint) error {
	if _, ok := m.Presences[id]; !ok {
		return errors.New("presence not found")
	}

	delete(m.Presences, id)
	return nil
}
