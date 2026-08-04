package presence

import (
	"errors"
	"time"
)

// Service contains presence business logic.
type Service struct {
	repo Repository
}

// NewService creates a new presence service.
func NewService(repo Repository) *Service {
	return &Service{
		repo: repo,
	}
}

// CreatePresence creates a presence record for a user.
func (s *Service) CreatePresence(userID uint) (*Presence, error) {
	_, err := s.repo.GetByUserID(userID)
	if err == nil {
		return nil, errors.New("presence already exists")
	}

	presence := &Presence{
		UserID:   userID,
		IsOnline: false,
		LastSeen: time.Now(),
	}

	if err := s.repo.Create(presence); err != nil {
		return nil, err
	}

	return presence, nil
}

// GetPresence retrieves a presence by ID.
func (s *Service) GetPresence(id uint) (*Presence, error) {
	return s.repo.GetByID(id)
}

// GetPresenceByUserID retrieves a user's presence.
func (s *Service) GetPresenceByUserID(userID uint) (*Presence, error) {
	return s.repo.GetByUserID(userID)
}

// GetOnlineUsers returns all online users.
func (s *Service) GetOnlineUsers() ([]Presence, error) {
	return s.repo.GetOnlineUsers()
}

// UpdatePresence updates a user's online status.
func (s *Service) UpdatePresence(userID uint, req UpdatePresenceRequest) (*Presence, error) {
	presence, err := s.repo.GetByUserID(userID)
	if err != nil {
		return nil, err
	}

	presence.IsOnline = req.IsOnline

	if !req.IsOnline {
		presence.LastSeen = time.Now()
	}

	if err := s.repo.Update(presence); err != nil {
		return nil, err
	}

	return presence, nil
}

// SetConnectionID updates the user's websocket connection ID.
func (s *Service) SetConnectionID(userID uint, connectionID string) (*Presence, error) {
	presence, err := s.repo.GetByUserID(userID)
	if err != nil {
		return nil, err
	}

	presence.ConnectionID = connectionID

	if err := s.repo.Update(presence); err != nil {
		return nil, err
	}

	return presence, nil
}

// DeletePresence removes a presence record.
func (s *Service) DeletePresence(id uint) error {
	return s.repo.Delete(id)
}
