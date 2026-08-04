package notifications

import (
	"doctor-platform/internal/database"

	"gorm.io/gorm"
)

type Repository interface {
	Create(notification *Notification) error
	GetByID(id uint) (*Notification, error)
	GetUserNotifications(userID uint) ([]Notification, error)
	GetUnreadNotifications(userID uint) ([]Notification, error)
	MarkAsRead(id uint) error
	MarkAllAsRead(userID uint) error
	Delete(id uint) error
}

type repository struct {
	db *gorm.DB
}

func NewRepository() Repository {
	return &repository{
		db: database.DB,
	}
}

func (r *repository) Create(notification *Notification) error {
	return r.db.Create(notification).Error
}

func (r *repository) GetByID(id uint) (*Notification, error) {
	var notification Notification

	if err := r.db.First(&notification, id).Error; err != nil {
		return nil, err
	}

	return &notification, nil
}

func (r *repository) GetUserNotifications(userID uint) ([]Notification, error) {
	var notifications []Notification

	err := r.db.
		Where("user_id = ?", userID).
		Order("created_at DESC").
		Find(&notifications).Error

	return notifications, err
}

func (r *repository) GetUnreadNotifications(userID uint) ([]Notification, error) {
	var notifications []Notification

	err := r.db.
		Where("user_id = ? AND is_read = ?", userID, false).
		Order("created_at DESC").
		Find(&notifications).Error

	return notifications, err
}

func (r *repository) MarkAsRead(id uint) error {
	return r.db.Model(&Notification{}).
		Where("id = ?", id).
		Update("is_read", true).Error
}

func (r *repository) MarkAllAsRead(userID uint) error {
	return r.db.Model(&Notification{}).
		Where("user_id = ?", userID).
		Update("is_read", true).Error
}

func (r *repository) Delete(id uint) error {
	return r.db.Delete(&Notification{}, id).Error
}
