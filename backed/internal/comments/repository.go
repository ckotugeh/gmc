package comments

import (
	"doctor-platform/internal/database"
)

type Repository interface {
	Create(comment *Comment) error
	GetByID(id uint) (*Comment, error)
	GetByPostID(postID uint) ([]Comment, error)
	Update(comment *Comment) error
	Delete(id uint) error
}

type repository struct{}

func NewRepository() Repository {
	return &repository{}
}

// Create a new comment
func (r *repository) Create(comment *Comment) error {
	return database.DB.Create(comment).Error
}

// Get a single comment
func (r *repository) GetByID(id uint) (*Comment, error) {
	var comment Comment

	if err := database.DB.First(&comment, id).Error; err != nil {
		return nil, err
	}

	return &comment, nil
}

// Get all comments for a post
func (r *repository) GetByPostID(postID uint) ([]Comment, error) {
	var comments []Comment

	if err := database.DB.
		Where("post_id = ?", postID).
		Order("created_at ASC").
		Find(&comments).Error; err != nil {
		return nil, err
	}

	return comments, nil
}

// Update a comment
func (r *repository) Update(comment *Comment) error {
	return database.DB.Save(comment).Error
}

// Delete a comment
func (r *repository) Delete(id uint) error {
	return database.DB.Delete(&Comment{}, id).Error
}
