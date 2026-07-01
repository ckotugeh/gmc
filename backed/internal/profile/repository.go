package profile

import (
	"doctor-platform/internal/database"
)

type Repository interface {
	Create(profile *Profile) error
	GetByUserID(userID uint) (*Profile, error)
	GetByID(id uint) (*Profile, error)
	Update(profile *Profile) error
	Delete(id uint) error
}

type repository struct{}

func NewRepository() Repository {
	return &repository{}
}

func (r *repository) Create(profile *Profile) error {
	return database.DB.Create(profile).Error
}

func (r *repository) GetByUserID(userID uint) (*Profile, error) {
	var profile Profile

	err := database.DB.
		Where("user_id = ?", userID).
		First(&profile).Error

	if err != nil {
		return nil, err
	}

	return &profile, nil
}

func (r *repository) GetByID(id uint) (*Profile, error) {
	var profile Profile

	err := database.DB.
		First(&profile, id).Error

	if err != nil {
		return nil, err
	}

	return &profile, nil
}

func (r *repository) Update(profile *Profile) error {
	return database.DB.Save(profile).Error
}

func (r *repository) Delete(id uint) error {
	return database.DB.Delete(&Profile{}, id).Error
}
