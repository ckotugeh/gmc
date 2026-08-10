package messages

import (
	"errors"
)

type MockRepository struct {
	messages map[uint]*Message
	nextID   uint
}

func NewMockRepository() Repository {
	return &MockRepository{
		messages: make(map[uint]*Message),
		nextID:   1,
	}
}

func (m *MockRepository) Create(message *Message) error {
	message.ID = m.nextID
	m.messages[m.nextID] = message
	m.nextID++

	return nil
}

func (m *MockRepository) GetByID(id uint) (*Message, error) {
	message, ok := m.messages[id]
	if !ok {
		return nil, errors.New("message not found")
	}

	return message, nil
}

func (m *MockRepository) GetConversation(user1ID, user2ID uint) ([]Message, error) {
	var conversation []Message

	for _, message := range m.messages {
		if (message.SenderID == user1ID && message.ReceiverID == user2ID) ||
			(message.SenderID == user2ID && message.ReceiverID == user1ID) {
			conversation = append(conversation, *message)
		}
	}

	return conversation, nil
}

func (m *MockRepository) GetUserMessages(userID uint) ([]Message, error) {
	var messages []Message

	for _, message := range m.messages {
		if message.SenderID == userID || message.ReceiverID == userID {
			messages = append(messages, *message)
		}
	}

	return messages, nil
}

func (m *MockRepository) Update(message *Message) error {
	if _, ok := m.messages[message.ID]; !ok {
		return errors.New("message not found")
	}

	m.messages[message.ID] = message
	return nil
}

func (m *MockRepository) Delete(id uint) error {
	if _, ok := m.messages[id]; !ok {
		return errors.New("message not found")
	}

	delete(m.messages, id)
	return nil
}

func (m *MockRepository) MarkAsRead(id uint) error {
	message, ok := m.messages[id]
	if !ok {
		return errors.New("message not found")
	}

	message.IsRead = true
	return nil
}
