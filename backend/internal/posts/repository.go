package posts

import (
	"doctor-platform/internal/database"
)

type Repository interface {
	Create(post *Post) error
	GetByID(id uint) (*Post, error)
	GetAll() ([]Post, error)
	GetByCommunityID(communityID uint) ([]Post, error)
	Update(post *Post) error
	Delete(id uint) error
}

type repository struct{}

func NewRepository() Repository {
	return &repository{}
}

// Create a new post
func (r *repository) Create(post *Post) error {
	return database.DB.Create(post).Error
}

// Get a single post with all related data
func (r *repository) GetByID(id uint) (*Post, error) {

	var post Post

	err := database.DB.
		Preload("Attachments").
		Preload("Tags").
		Preload("Polls").
		Preload("Polls.Options").
		First(&post, id).Error

	if err != nil {
		return nil, err
	}

	return &post, nil
}

// Get all posts belonging to one community
func (r *repository) GetByCommunityID(communityID uint) ([]Post, error) {

	var posts []Post

	err := database.DB.
		Where("community_id = ?", communityID).
		Preload("Attachments").
		Preload("Tags").
		Preload("Polls").
		Preload("Polls.Options").
		Order("created_at DESC").
		Find(&posts).Error

	if err != nil {
		return nil, err
	}

	return posts, nil
}

// Update a post
func (r *repository) Update(post *Post) error {
	return database.DB.Save(post).Error
}

// Delete a post
func (r *repository) Delete(id uint) error {
	return database.DB.Delete(&Post{}, id).Error
}

// Get all posts
func (r *repository) GetAll() ([]Post, error) {
	var posts []Post

	err := database.DB.
		Preload("Attachments").
		Preload("Tags").
		Preload("Polls").
		Preload("Polls.Options").
		Order("created_at DESC").
		Find(&posts).Error

	if err != nil {
		return []Post{}, err
	}

	if posts == nil {
		return []Post{}, nil
	}

	return posts, nil
}
