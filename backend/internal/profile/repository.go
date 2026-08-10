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

	ExistsByUserID(userID uint) (bool, error)
}

type repository struct{}

func (r *repository) ExistsByUserID(userID uint) (bool, error) {
	var count int64

	err := database.DB.
		Model(&Profile{}).
		Where("user_id = ?", userID).
		Count(&count).Error

	if err != nil {
		return false, err
	}

	return count > 0, nil
}

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
