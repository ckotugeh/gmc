package websockets

// MockRepository is an in-memory implementation for testing.
type MockRepository struct {
	messages []Message
}

func NewMockRepository() *MockRepository {
	return &MockRepository{
		messages: []Message{},
	}
}

func (m *MockRepository) SaveMessage(msg Message) error {
	m.messages = append(m.messages, msg)
	return nil
}

func (m *MockRepository) GetMessages(userID uint) ([]Message, error) {
	return m.messages, nil
}
