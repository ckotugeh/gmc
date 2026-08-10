package messages

import (
	"errors"
	"strings"
)

type Service interface {
	CreateMessage(senderID uint, req *CreateMessageRequest) (*Message, error)
	GetMessage(id uint) (*Message, error)
	GetConversation(userID, otherUserID uint) ([]Message, error)
	GetUserMessages(userID uint) ([]Message, error)
	UpdateMessage(id, userID uint, req *UpdateMessageRequest) (*Message, error)
	DeleteMessage(id, userID uint) error
	MarkAsRead(id, userID uint) error
}

type service struct {
	repo Repository
}

func NewService(repo Repository) Service {
	return &service{
		repo: repo,
	}
}

func (s *service) CreateMessage(senderID uint, req *CreateMessageRequest) (*Message, error) {
	if strings.TrimSpace(req.Content) == "" {
		return nil, errors.New("message content is required")
	}

	if req.ReceiverID == 0 {
		return nil, errors.New("receiver is required")
	}

	if senderID == req.ReceiverID {
		return nil, errors.New("cannot send a message to yourself")
	}

	if req.MessageType == "" {
		req.MessageType = MessageTypeText
	}

	message := &Message{
		SenderID:      senderID,
		ReceiverID:    req.ReceiverID,
		Content:       strings.TrimSpace(req.Content),
		MessageType:   req.MessageType,
		AttachmentURL: req.AttachmentURL,
		ReplyToID:     req.ReplyToID,
		IsRead:        false,
		IsEdited:      false,
	}

	if err := s.repo.Create(message); err != nil {
		return nil, err
	}

	return message, nil
}

func (s *service) GetMessage(id uint) (*Message, error) {
	return s.repo.GetByID(id)
}

func (s *service) GetConversation(userID, otherUserID uint) ([]Message, error) {
	return s.repo.GetConversation(userID, otherUserID)
}

func (s *service) GetUserMessages(userID uint) ([]Message, error) {
	return s.repo.GetUserMessages(userID)
}

func (s *service) UpdateMessage(id, userID uint, req *UpdateMessageRequest) (*Message, error) {
	message, err := s.repo.GetByID(id)
	if err != nil {
		return nil, err
	}

	if message.SenderID != userID {
		return nil, errors.New("unauthorized")
	}

	if strings.TrimSpace(req.Content) == "" {
		return nil, errors.New("message content is required")
	}

	message.Content = strings.TrimSpace(req.Content)
	message.AttachmentURL = req.AttachmentURL
	message.IsEdited = true

	if err := s.repo.Update(message); err != nil {
		return nil, err
	}

	return message, nil
}

func (s *service) DeleteMessage(id, userID uint) error {
	message, err := s.repo.GetByID(id)
	if err != nil {
		return err
	}

	if message.SenderID != userID {
		return errors.New("unauthorized")
	}

	return s.repo.Delete(id)
}

func (s *service) MarkAsRead(id, userID uint) error {
	message, err := s.repo.GetByID(id)
	if err != nil {
		return err
	}

	// Only the receiver can mark the message as read.
	if message.ReceiverID != userID {
		return errors.New("unauthorized")
	}

	if message.IsRead {
		return nil
	}

	return s.repo.MarkAsRead(id)
}
