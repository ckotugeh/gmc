package presence

import (
	"errors"

	"gorm.io/gorm"
)

// Repository defines presence persistence methods.
type Repository interface {
	Create(presence *Presence) error
	Update(presence *Presence) error
	GetByID(id uint) (*Presence, error)
	GetByUserID(userID uint) (*Presence, error)
	GetOnlineUsers() ([]Presence, error)
	Delete(id uint) error
}

// repository implements Repository.
type repository struct {
	db *gorm.DB
}

// NewRepository creates a new presence repository.
func NewRepository(db *gorm.DB) Repository {
	return &repository{
		db: db,
	}
}

// Create creates a new presence record.
func (r *repository) Create(presence *Presence) error {
	return r.db.Create(presence).Error
}

// Update updates an existing presence record.
func (r *repository) Update(presence *Presence) error {
	return r.db.Save(presence).Error
}

// GetByID retrieves a presence record by ID.
func (r *repository) GetByID(id uint) (*Presence, error) {
	var presence Presence

	if err := r.db.First(&presence, id).Error; err != nil {
		return nil, err
	}

	return &presence, nil
}

// GetByUserID retrieves a presence record by user ID.
func (r *repository) GetByUserID(userID uint) (*Presence, error) {
	var presence Presence

	if err := r.db.Where("user_id = ?", userID).First(&presence).Error; err != nil {
		return nil, err
	}

	return &presence, nil
}

// GetOnlineUsers returns all currently online users.
func (r *repository) GetOnlineUsers() ([]Presence, error) {
	var presences []Presence

	if err := r.db.Where("is_online = ?", true).Find(&presences).Error; err != nil {
		return nil, err
	}

	return presences, nil
}

// Delete deletes a presence record.
func (r *repository) Delete(id uint) error {
	result := r.db.Delete(&Presence{}, id)

	if result.RowsAffected == 0 {
		return errors.New("presence not found")
	}

	return result.Error
}
