package messages

import (
	"doctor-platform/internal/database"

	"gorm.io/gorm"
)

type Repository interface {
	Create(message *Message) error
	GetByID(id uint) (*Message, error)
	GetConversation(user1ID, user2ID uint) ([]Message, error)
	GetUserMessages(userID uint) ([]Message, error)
	Update(message *Message) error
	Delete(id uint) error
	MarkAsRead(id uint) error
}

type repository struct {
	db *gorm.DB
}

func NewRepository() Repository {
	return &repository{
		db: database.DB,
	}
}

func (r *repository) Create(message *Message) error {
	return r.db.Create(message).Error
}

func (r *repository) GetByID(id uint) (*Message, error) {
	var message Message

	if err := r.db.First(&message, id).Error; err != nil {
		return nil, err
	}

	return &message, nil
}

func (r *repository) GetConversation(user1ID, user2ID uint) ([]Message, error) {
	var messages []Message

	err := r.db.
		Where(
			"(sender_id = ? AND receiver_id = ?) OR (sender_id = ? AND receiver_id = ?)",
			user1ID, user2ID,
			user2ID, user1ID,
		).
		Order("created_at ASC").
		Find(&messages).Error

	return messages, err
}

func (r *repository) GetUserMessages(userID uint) ([]Message, error) {
	var messages []Message

	err := r.db.
		Where("sender_id = ? OR receiver_id = ?", userID, userID).
		Order("created_at DESC").
		Find(&messages).Error

	return messages, err
}

func (r *repository) Update(message *Message) error {
	return r.db.Save(message).Error
}

func (r *repository) Delete(id uint) error {
	return r.db.Delete(&Message{}, id).Error
}

func (r *repository) MarkAsRead(id uint) error {
	return r.db.Model(&Message{}).
		Where("id = ?", id).
		Update("is_read", true).Error
}
