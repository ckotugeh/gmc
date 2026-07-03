package communities

import (
	"doctor-platform/internal/database"
)

type Repository interface {
	Create(community *Community) error
	GetByID(id uint) (*Community, error)
	GetByName(name string) (*Community, error)
	GetAll() ([]Community, error)
	Update(community *Community) error
	Delete(id uint) error
}

type repository struct{}

func NewRepository() Repository {
	return &repository{}
}

func (r *repository) Create(community *Community) error {
	return database.DB.Create(community).Error
}

func (r *repository) GetByID(id uint) (*Community, error) {
	var community Community

	if err := database.DB.First(&community, id).Error; err != nil {
		return nil, err
	}

	return &community, nil
}

func (r *repository) GetByName(name string) (*Community, error) {
	var community Community

	if err := database.DB.Where("name = ?", name).First(&community).Error; err != nil {
		return nil, err
	}

	return &community, nil
}

func (r *repository) GetAll() ([]Community, error) {
	var communities []Community

	if err := database.DB.Order("created_at DESC").Find(&communities).Error; err != nil {
		return nil, err
	}

	return communities, nil
}

func (r *repository) Update(community *Community) error {
	return database.DB.Save(community).Error
}

func (r *repository) Delete(id uint) error {
	return database.DB.Delete(&Community{}, id).Error
}