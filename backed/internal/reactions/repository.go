package reactions

import (
	"doctor-platform/internal/database"
)

type Repository interface {
	Create(reaction *Reaction) error
	GetByID(id uint) (*Reaction, error)
	GetByPostAndUser(postID, userID uint) (*Reaction, error)
	GetByPost(postID uint) ([]Reaction, error)
	Update(reaction *Reaction) error
	Delete(id uint) error
}

type repository struct{}

func NewRepository() Repository {
	return &repository{}
}

func (r *repository) Create(reaction *Reaction) error {
	return database.DB.Create(reaction).Error
}

func (r *repository) GetByID(id uint) (*Reaction, error) {
	var reaction Reaction

	if err := database.DB.First(&reaction, id).Error; err != nil {
		return nil, err
	}

	return &reaction, nil
}

func (r *repository) GetByPostAndUser(postID, userID uint) (*Reaction, error) {
	var reaction Reaction

	if err := database.DB.
		Where("post_id = ? AND user_id = ?", postID, userID).
		First(&reaction).Error; err != nil {
		return nil, err
	}

	return &reaction, nil
}

func (r *repository) GetByPost(postID uint) ([]Reaction, error) {
	var reactions []Reaction

	if err := database.DB.
		Where("post_id = ?", postID).
		Find(&reactions).Error; err != nil {
		return nil, err
	}

	return reactions, nil
}

func (r *repository) Update(reaction *Reaction) error {
	return database.DB.Save(reaction).Error
}

func (r *repository) Delete(id uint) error {
	return database.DB.Delete(&Reaction{}, id).Error
}
