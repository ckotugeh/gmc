package notifications

import (
	"errors"
)

type MockRepository struct {
	notifications map[uint]*Notification
	nextID        uint
}

func NewMockRepository() Repository {
	return &MockRepository{
		notifications: make(map[uint]*Notification),
		nextID:        1,
	}
}

func (m *MockRepository) Create(notification *Notification) error {
	notification.ID = m.nextID
	m.notifications[m.nextID] = notification
	m.nextID++

	return nil
}

func (m *MockRepository) GetByID(id uint) (*Notification, error) {
	notification, ok := m.notifications[id]
	if !ok {
		return nil, errors.New("notification not found")
	}

	return notification, nil
}

func (m *MockRepository) GetUserNotifications(userID uint) ([]Notification, error) {
	var notifications []Notification

	for _, notification := range m.notifications {
		if notification.UserID == userID {
			notifications = append(notifications, *notification)
		}
	}

	return notifications, nil
}

func (m *MockRepository) GetUnreadNotifications(userID uint) ([]Notification, error) {
	var notifications []Notification

	for _, notification := range m.notifications {
		if notification.UserID == userID && !notification.IsRead {
			notifications = append(notifications, *notification)
		}
	}

	return notifications, nil
}

func (m *MockRepository) MarkAsRead(id uint) error {
	notification, ok := m.notifications[id]
	if !ok {
		return errors.New("notification not found")
	}

	notification.IsRead = true
	return nil
}

func (m *MockRepository) MarkAllAsRead(userID uint) error {
	for _, notification := range m.notifications {
		if notification.UserID == userID {
			notification.IsRead = true
		}
	}

	return nil
}

func (m *MockRepository) Delete(id uint) error {
	if _, ok := m.notifications[id]; !ok {
		return errors.New("notification not found")
	}

	delete(m.notifications, id)
	return nil
}
