package notifications

import (
	"errors"
	"strings"
)

type Service interface {
	CreateNotification(req *CreateNotificationRequest) (*Notification, error)
	GetNotification(id, userID uint) (*Notification, error)
	GetUserNotifications(userID uint) ([]Notification, error)
	GetUnreadNotifications(userID uint) ([]Notification, error)
	MarkAsRead(id, userID uint) error
	MarkAllAsRead(userID uint) error
	DeleteNotification(id, userID uint) error
}

type service struct {
	repo Repository
}

func NewService(repo Repository) Service {
	return &service{
		repo: repo,
	}
}

func (s *service) CreateNotification(req *CreateNotificationRequest) (*Notification, error) {
	if req.UserID == 0 {
		return nil, errors.New("user ID is required")
	}

	if strings.TrimSpace(req.Type) == "" {
		return nil, errors.New("notification type is required")
	}

	if strings.TrimSpace(req.Title) == "" {
		return nil, errors.New("notification title is required")
	}

	if strings.TrimSpace(req.Message) == "" {
		return nil, errors.New("notification message is required")
	}

	notification := &Notification{
		UserID:        req.UserID,
		ActorID:       req.ActorID,
		Type:          strings.TrimSpace(req.Type),
		Title:         strings.TrimSpace(req.Title),
		Message:       strings.TrimSpace(req.Message),
		ReferenceID:   req.ReferenceID,
		ReferenceType: strings.TrimSpace(req.ReferenceType),
		ActionURL:     strings.TrimSpace(req.ActionURL),
		IsRead:        false,
	}

	if err := s.repo.Create(notification); err != nil {
		return nil, err
	}

	return notification, nil
}

func (s *service) GetNotification(id, userID uint) (*Notification, error) {
	notification, err := s.repo.GetByID(id)
	if err != nil {
		return nil, err
	}

	if notification.UserID != userID {
		return nil, errors.New("unauthorized")
	}

	return notification, nil
}

func (s *service) GetUserNotifications(userID uint) ([]Notification, error) {
	return s.repo.GetUserNotifications(userID)
}

func (s *service) GetUnreadNotifications(userID uint) ([]Notification, error) {
	return s.repo.GetUnreadNotifications(userID)
}

func (s *service) MarkAsRead(id, userID uint) error {
	notification, err := s.repo.GetByID(id)
	if err != nil {
		return err
	}

	if notification.UserID != userID {
		return errors.New("unauthorized")
	}

	if notification.IsRead {
		return nil
	}

	return s.repo.MarkAsRead(id)
}

func (s *service) MarkAllAsRead(userID uint) error {
	return s.repo.MarkAllAsRead(userID)
}

func (s *service) DeleteNotification(id, userID uint) error {
	notification, err := s.repo.GetByID(id)
	if err != nil {
		return err
	}

	if notification.UserID != userID {
		return errors.New("unauthorized")
	}

	return s.repo.Delete(id)
}
